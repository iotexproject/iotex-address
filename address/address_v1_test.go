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
		{StakingCreateAddrHash, StakingCreateAddr},
		{StakingAddDepositAddrHash, StakingAddDepositAddr},
		{StakingChangeCandAddrHash, StakingChangeCandAddr},
		{StakingUnstakeAddrHash, StakingUnstakeAddr},
		{StakingWithdrawAddrHash, StakingWithdrawAddr},
		{StakingRestakeAddrHash, StakingRestakeAddr},
		{StakingTransferAddrHash, StakingTransferAddr},
		{StakingRegisterCandAddrHash, StakingRegisterCandAddr},
		{StakingUpdateCandAddrHash, StakingUpdateCandAddr},
	}
	for _, test := range tests {
		require.False(IsAddrV1Special(test.addr))
		addr, _ := _v1.FromBytes(test.hashByte[:])
		require.Equal(test.addr, addr.String())
		addr, _ = _v1.FromString(test.addr)
		require.Equal(test.hashByte[:], addr.Bytes())
	}
}
