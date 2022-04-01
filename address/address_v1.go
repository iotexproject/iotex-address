// Copyright (c) 2018 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package address

import (
	"encoding/hex"
	"log"

	"github.com/pkg/errors"

	"github.com/iotexproject/iotex-address/address/bech32"
)

// V1AddressStringLength is the length of v1 address string
const V1AddressStringLength = 41

// _v1 is a singleton and defines V1 address metadata
var _v1 = v1{
	AddressLength: 20,
}

type v1 struct {
	// AddressLength indicates the byte length of an address
	AddressLength int
}

// FromString decodes an encoded address string into an address struct
func (v *v1) FromString(encodedAddr string) (Address, error) {
	if IsAddrV1Special(encodedAddr) {
		return newAddrV1Special(encodedAddr), nil
	}
	if len(encodedAddr) != V1AddressStringLength {
		return nil, errors.Wrapf(ErrInvalidAddr, "address length = %d, expecting 41", len(encodedAddr))
	}
	payload, err := v.decodeBech32(encodedAddr)
	if err != nil {
		return nil, err
	}
	return v.FromBytes(payload)
}

// FromStringLeacy decodes an encoded address string into an address struct
func (v *v1) FromStringLegacy(encodedAddr string) (Address, error) {
	if IsAddrV1Special(encodedAddr) {
		return newAddrV1Special(encodedAddr), nil
	}
	payload, err := v.decodeBech32Legacy(encodedAddr)
	if err != nil {
		return nil, err
	}
	return v.FromBytes(payload)
}

// FromBytes converts a byte array into an address struct
// If b is larger than v.AddressLength, b will be cropped from the left
// otherwise, b will be left-padded with 0
func (v *v1) FromBytes(b []byte) (Address, error) {
	if len(b) > v.AddressLength {
		b = b[len(b)-v.AddressLength:]
	}
	addr := AddrV1{}
	copy(addr.payload[v.AddressLength-len(b):], b)
	return &addr, nil
}

// FromHex converts a hex-encoded string into an address struct
func (v *v1) FromHex(s string) (Address, error) {
	if len(s) > 1 {
		if s[0] == '0' && (s[1] == 'x' || s[1] == 'X') {
			s = s[2:]
		}
	}
	if len(s)%2 == 1 {
		s = "0" + s
	}
	bytes, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return v.FromBytes(bytes)
}

func (v *v1) decodeBech32(encodedAddr string) ([]byte, error) {
	hrp, grouped, err := bech32.Decode(encodedAddr)
	if err != nil {
		return nil, errors.Wrapf(ErrInvalidAddr, err.Error())
	}
	if hrp != prefix() {
		return nil, errors.Wrapf(ErrInvalidAddr, "hrp %s and address prefix %s don't match", hrp, prefix())
	}
	// Group the payload into 8 bit groups.
	payload, err := bech32.ConvertBits(grouped, 5, 8, false)
	if err != nil {
		return nil, errors.Wrapf(ErrInvalidAddr, err.Error())
	}
	return payload, nil
}

func (v *v1) decodeBech32Legacy(encodedAddr string) ([]byte, error) {
	hrp, grouped, err := bech32.Decode(encodedAddr)
	if hrp != prefix() {
		return nil, errors.Wrapf(err, "hrp %s and address prefix %s don't match", hrp, prefix())
	}
	// Group the payload into 8 bit groups.
	payload, err := bech32.ConvertBits(grouped, 5, 8, false)
	if err != nil {
		return nil, errors.Wrapf(err, "error when converting 5 bit groups into the payload")
	}
	return payload, nil
}

type (
	// Hash160 for 160-bit hash used for account and smart contract address
	Hash160 [20]byte

	// AddrV1 is V1 address format to be used on IoTeX blockchain and subchains
	// It is composed of a 20-byte hash derived from the the public key
	AddrV1 struct {
		payload Hash160
	}
)

// String encodes an address struct into a a String encoded address string
// The encoded address string will start with "io" for mainnet, and with "it" for testnet
func (addr *AddrV1) String() string {
	payload := addr.payload[:]
	// Group the payload into 5 bit groups.
	grouped, err := bech32.ConvertBits(payload, 8, 5, true)
	if err != nil {
		log.Panic("Error when grouping the payload into 5 bit groups." + err.Error())
		return ""
	}
	encodedAddr, err := bech32.Encode(prefix(), grouped)
	if err != nil {
		log.Panic("Error when encoding bytes into a base32 string." + err.Error())
		return ""
	}
	return encodedAddr
}

// Bytes converts an address struct into a byte array
func (addr *AddrV1) Bytes() []byte {
	return addr.payload[:]
}

// Hex is the hex-encoding of Bytes, prefixed with "0x"
func (addr *AddrV1) Hex() string {
	return "0x" + hex.EncodeToString(addr.payload[:])
}
