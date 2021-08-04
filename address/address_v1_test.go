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
		require.True(Equal(addr1, addr2))
		_, err = FromHex(addrHex[:len(addrHex)-2])
		require.Error(err)
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

	addr1, err := _v1.FromBytes(StakingProtocolAddrHash[:])
	require.NoError(err)
	require.Equal("io1qnpz47hx5q6r3w876axtrn6yz95d70cjl35r53", addr1.String())
	addr1, err = _v1.FromBytes(RewardingProtocolAddrHash[:])
	require.NoError(err)
	require.Equal("io154mvzs09vkgn0hw6gg3ayzw5w39jzp47f8py9v", addr1.String())

	// special address has same length
	require.Equal(41, len(StakingBucketPoolAddr))
	require.Equal(41, len(RewardingPoolAddr))

	// but cannot decode
	addr1, err = _v1.FromString(StakingBucketPoolAddr)
	require.Error(err)
	addr1, err = _v1.FromString(RewardingPoolAddr)
	require.Error(err)

	// special address for staking actions
	tests := []struct {
		hashByte Hash160
		addr     string
	}{
		{StakingCreateAddrHash, "io1qqqqqqqqqqq8xarpdd5kue6rwfjkzar9k0wk6t"},
		{StakingAddDepositAddrHash, "io1qqqqqum5v94kjmn8g9jxg3r9wphhx6t58x7tye"},
		{StakingChangeCandAddrHash, "io1qqqqqum5v94kjmn8gd5xzmn8v4pkzmnye5v3fh"},
		{StakingUnstakeAddrHash, "io1qqqqqqqqqpehgcttd9hxw4twwd6xz6m9pl4r27"},
		{StakingWithdrawAddrHash, "io1qqqqqqqqwd6xz6mfden4w6t5dpj8ycthwsq5ng"},
		{StakingRestakeAddrHash, "io1qqqqqqqqqpehgcttd9hxw5n9wd6xz6m995w4zm"},
		{StakingTransferAddrHash, "io1qqqqqqqqwd6xz6mfden4gunpdeekvetjzwh99y"},
		{StakingRegisterCandAddrHash, "io1qpehgcttd9hxw5n9va5hxar9wfpkzmnyahxhjk"},
		{StakingUpdateCandAddrHash, "io1qqqqqum5v94kjmn824cxgct5v4pkzmnyxxj98n"},
	}
	for _, test := range tests {
		addr, _ := _v1.FromBytes(test.hashByte[:])
		require.Equal(test.addr, addr.String())
	}
}
