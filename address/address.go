// Copyright (c) 2018 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package address

import (
	"bytes"
	"errors"
	"os"
	"strings"
)

// init reads IOTEX_NETWORK_TYPE environment variable. If it exists and the value is equal to "testnet" with case
// ignored,  the global variable isTestNet is set to true for the whole runtime
func init() {
	isTestNet = strings.EqualFold(os.Getenv("IOTEX_NETWORK_TYPE"), "testnet")
}

const (
	// MainnetPrefix is the prefix added to the human readable address of mainnet
	MainnetPrefix = "io"
	// TestnetPrefix is the prefix added to the human readable address of testnet
	TestnetPrefix = "it"
)

// ErrInvalidAddr indicates the invalid address error
var ErrInvalidAddr = errors.New("invalid address")

var isTestNet bool

// Address defines the interface of the blockchain address
type Address interface {
	// String encodes the address into a string using bech32 encoding
	String() string

	// Bytes returns the underlying 20-byte public key hash
	Bytes() []byte

	// Hex is the hex-encoding of Bytes, prefixed with "0x"
	Hex() string
}

// FromString decodes an encoded address string into an address struct
func FromString(encodedAddr string) (Address, error) { return _v1.FromString(encodedAddr) }

// FromStringLegacy decodes an encoded address string into an address struct
func FromStringLegacy(encodedAddr string) (Address, error) { return _v1.FromStringLegacy(encodedAddr) }

// FromBytes converts a byte array into an address struct
func FromBytes(bytes []byte) (Address, error) { return _v1.FromBytes(bytes) }

// FromHex converts a hex-encoded string into an address struct
func FromHex(s string) (Address, error) { return _v1.FromHex(s) }

// prefix returns the current prefix
func prefix() string {
	prefix := MainnetPrefix
	if isTestNet {
		prefix = TestnetPrefix
	}
	return prefix
}

// Equal determine if two addresses are equal
func Equal(addr1 Address, addr2 Address) bool {
	if addr1 == nil && addr2 == nil {
		return true
	}
	if addr1 != nil && addr2 == nil || addr1 == nil && addr2 != nil {
		return false
	}
	return bytes.Equal(addr1.Bytes(), addr2.Bytes())
}
