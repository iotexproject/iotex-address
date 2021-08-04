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

	tests := []struct {
		specialAddr string
		hashByte    Hash160
		hash        string
	}{
		{StakingBucketPoolAddr, StakingProtocolAddrHash, "io1qnpz47hx5q6r3w876axtrn6yz95d70cjl35r53"},
		{StakingCreateAddr, StakingCreateAddrHash, "io16tmc6yd9kn0z90axeack8r0sd9dsr0p8lve54d"},
		{StakingAddDepositAddr, StakingAddDepositAddrHash, "io1l08h9492aphd5p7e6mc22e352dwpc62qwj43za"},
		{StakingChangeCandidateAddr, StakingChangeCandidateAddrHash, "io1332snsr9jnpdlsmhprmxsktdq6q3a3yh0wu8uf"},
		{StakingUnstakeAddr, StakingUnstakeAddrHash, "io1mdr5s62d0pww8pn0zjx4xne7u34z66ml53du45"},
		{StakingWithdrawAddr, StakingWithdrawAddrHash, "io1sa8nv8s0u4h0ngvhavl0lzqnrq88a6dymkuxv7"},
		{StakingRestakeAddr, StakingRestakeAddrHash, "io1fg8hwmdgkm4ytrx4p7mlx0sak0wmq6az96lnlk"},
		{StakingTransferStakeAddr, StakingTransferStakeAddrHash, "io12v9rhxl8xr0neh9x2uwyu5xhyl043048ctdhs7"},
		{StakingCandidateRegisterAddr, StakingCandidateRegisterAddrHash, "io1yrwqura2ww2gt00gmf79jdkuh8hq0kufufdsss"},
		{StakingCandidateUpdateAddr, StakingCandidateUpdateAddrHash, "io15q33lxah8u5g97nh02r5x3c2x22pj5p0zavtd3"},
		{RewardingPoolAddr, RewardingProtocolAddrHash, "io154mvzs09vkgn0hw6gg3ayzw5w39jzp47f8py9v"},
	}
	const ADDRLENGTH int = 41

	for _, test := range tests {
		require.Equal(ADDRLENGTH, len(test.specialAddr))
		_, err := _v1.FromString(test.specialAddr)
		require.Error(err)
		addr, _ := _v1.FromBytes(test.hashByte[:])
		require.Equal(test.hash, addr.String())
	}
}
