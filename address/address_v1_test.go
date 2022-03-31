// Copyright (c) 2018 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package address

import (
	"crypto/rand"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddress(t *testing.T) {
	require := require.New(t)

	// require v1 length = 20 bytes
	require.Equal(20, _v1.AddressLength)

	runTest := func(t *testing.T) {
		pkHash := make([]byte, _v1.AddressLength)
		n, err := io.ReadFull(rand.Reader, pkHash)
		require.NoError(err)
		require.Equal(_v1.AddressLength, n)

		addr1, err := FromBytes(pkHash)
		require.NoError(err)
		require.Equal(pkHash, addr1.Bytes())

		encodedAddr := addr1.String()
		if isTestNet {
			require.True(strings.HasPrefix(encodedAddr, TestnetPrefix))
		} else {
			require.True(strings.HasPrefix(encodedAddr, MainnetPrefix))
		}
		addr2, err := FromString(encodedAddr)
		require.NoError(err)
		require.True(Equal(addr1, addr2))

		addrHex := addr2.Hex()
		require.Equal(42, len(addrHex))
		require.Equal("0x", addrHex[:2])
		addr2, err = FromHex(addrHex)
		require.NoError(err)
		require.Equal(addr1, addr2)
		addr2, err = FromHex(addrHex[2:])
		require.NoError(err)
		require.Equal(addr1, addr2)
		// remove the last byte
		addr2, err = FromHex(addrHex[:len(addrHex)-2])
		require.NoError(err)
		require.Equal(addr2.Hex()[4:], addrHex[2:len(addrHex)-2])
	}
	t.Run("testnet", func(t *testing.T) {
		require.NoError(os.Setenv("IOTEX_NETWORK_TYPE", "testnet"))
		runTest(t)
	})
	t.Run("mainnet", func(t *testing.T) {
		require.NoError(os.Setenv("IOTEX_NETWORK_TYPE", "mainnet"))
		runTest(t)
	})
}

func TestAddressError(t *testing.T) {
	require := require.New(t)

	pkHash := make([]byte, _v1.AddressLength)
	n, err := io.ReadFull(rand.Reader, pkHash)
	require.NoError(err)
	require.Equal(_v1.AddressLength, n)

	addr1, err := _v1.FromBytes(pkHash)
	require.NoError(err)

	encodedAddr := addr1.String()
	encodedAddrBytes := []byte(encodedAddr)
	encodedAddrBytes[len(encodedAddrBytes)-1] = 'o'
	addr2, err := _v1.FromString(string(encodedAddrBytes))
	require.Nil(addr2)
	require.Error(err)
}

func TestSpecialAddress(t *testing.T) {
	require := require.New(t)

	for _, v := range []string{
		StakingBucketPoolAddr,
		RewardingPoolAddr,
	} {
		require.Equal(41, len(v))
		require.True(IsAddrV1Special(v))
		addr, err := _v1.FromString(v)
		require.NoError(err)
		require.Equal(v, addr.String())
		require.Panics(func() {
			addr.Bytes()
		})
		require.Panics(func() {
			addr.Hex()
		})
	}

	// special address for staking actions
	tests := []struct {
		hashByte Hash160
		addr     string
	}{
		{Hash160{}, ZeroAddress},
		{StakingProtocolAddrHash, StakingProtocolAddr},
		{RewardingProtocolAddrHash, RewardingProtocol},
	}
	for _, test := range tests {
		require.False(IsAddrV1Special(test.addr))
		addr, _ := _v1.FromBytes(test.hashByte[:])
		require.Equal(test.addr, addr.String())
		addr, _ = _v1.FromString(test.addr)
		require.Equal(test.hashByte[:], addr.Bytes())
	}
}

func TestLegacyFormat(t *testing.T) {
	r := require.New(t)

	for _, v := range []struct {
		addr           string
		errLegacy, err string
		nominal        string
	}{
		{"iota1qp3mxh8gx8fkqmss9c6jsm979wuv6qpm0waw6vhxt0dwzze8xxzkqzy3lxu", // wrong hrp, wrong checksum, long size
			"checksum failed: Expected anqr4d", "address length = 64",
			""},
		{"iota1qp3mxh8gx8fkqmss9c6jsm979wuv6qpm0waw6vhxt0dwzze8xxzkqanqr4d", // wrong hrp, right checksum, long size
			"", "address length = 64",
			ZeroAddress},
		{"iota1qp3mxh8gx8fkqmss9c6jsm979wuv6qpm0w", // wrong hrp, wrong checksum, short size
			"checksum failed: Expected 5a73lu", "address length = 39",
			""},
		{"iota1qp3mxh8gx8fkqmss9c6jsm979wuv5a73lu", // wrong hrp, right checksum, short size
			"", "address length = 39",
			ZeroAddress},
		{"iota1qp3mxh8gx8fkqmss9c6jsm979wuv6qpm0waw", // wrong hrp, wrong checksum, right size
			"checksum failed: Expected 06dmq2", "checksum failed: Expected 06dmq2",
			""},
		{"iota1qp3mxh8gx8fkqmss9c6jsm979wuv6q06dmq2", // wrong hrp, right checksum, right size
			"", "hrp iota and address prefix io don't match",
			ZeroAddress},
		{"io1qp3mxh8gx8fkqmss9c6jsm979wuv6qpm0waw6vhxt0dwzze8xxzkqanqr4d", // right hrp, wrong checksum, long size
			"checksum failed: Expected zy3lxu", "address length = 62",
			""},
		{"io1qp3mxh8gx8fkqmss9c6jsm979wuv6qpm0waw6vhxt0dwzze8xxzkqzy3lxu", // right hrp, right checksum, long size
			"", "address length = 62",
			"io1djlzhwxdqqahhwhdxtn9hkhppvnnrptqtwf2h5"},
		{"io1djlzhwxdqqahhwhdxtn9hkhppvnnrptqtwfh", // right hrp, wrong checksum, short size
			"checksum failed: Expected 726csn", "address length = 39",
			""},
		{"io1djlzhwxdqqahhwhdxtn9hkhppvnnrp726csn", // right hrp, right checksum, short size
			"invalid incomplete group", "address length = 39",
			""},
		{"io1djlzhwxdqqahhwhdxtn9hkhppvnnrptqzy3lxu", // right hrp, wrong checksum, right size
			"checksum failed: Expected twf2h5", "checksum failed: Expected twf2h5",
			""},
		{"io1djlzhwxdqqahhwhdxtn9hkhppvnnrptqtwf2h5", // right hrp, right checksum, right size
			"", "",
			"io1djlzhwxdqqahhwhdxtn9hkhppvnnrptqtwf2h5"},
	} {
		a, err := FromStringLegacy(v.addr)
		if v.errLegacy != "" {
			r.Contains(err.Error(), v.errLegacy)
		} else {
			r.Equal(v.nominal, a.String())
		}
		a, err = FromString(v.addr)
		if v.err != "" {
			r.Contains(err.Error(), v.err)
		} else {
			r.Equal(v.nominal, a.String())
		}
	}
}
