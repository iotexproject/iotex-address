package address

import "golang.org/x/crypto/sha3"

const (
	// ZeroAddress is the IoTeX address whose hash160 is all zero
	ZeroAddress = "io1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqd39ym7"

	// StakingBucketPoolAddr is the staking bucket pool address
	StakingBucketPoolAddr = "io000000000000000000000000stakingprotocol"

	// RewardingPoolAddr is the rewarding pool address
	RewardingPoolAddr = "io0000000000000000000000rewardingprotocol"

	// StakingProtocolAddr is the staking protocol address
	StakingProtocolAddr = "io1qnpz47hx5q6r3w876axtrn6yz95d70cjl35r53"

	// RewardingProtocol is the rewarding protocol address
	RewardingProtocol = "io154mvzs09vkgn0hw6gg3ayzw5w39jzp47f8py9v"

	// StakingCreateAddr is the staking create address
	StakingCreateAddr = "io1qqqqqqqqqqq8xarpdd5kue6rwfjkzar9k0wk6t"

	// StakingAddDepositAddr is the staking add deposit address
	StakingAddDepositAddr = "io1qqqqqum5v94kjmn8g9jxg3r9wphhx6t58x7tye"

	// StakingChangeCandAddr is the staking change candidate address
	StakingChangeCandAddr = "io1qqqqqum5v94kjmn8gd5xzmn8v4pkzmnye5v3fh"

	// StakingUnstakeAddr is the staking unstake address
	StakingUnstakeAddr = "io1qqqqqqqqqpehgcttd9hxw4twwd6xz6m9pl4r27"

	// StakingWithdrawAddr is the staking withdraw address
	StakingWithdrawAddr = "io1qqqqqqqqwd6xz6mfden4w6t5dpj8ycthwsq5ng"

	// StakingRestakeAddr is the staking restake address
	StakingRestakeAddr = "io1qqqqqqqqqpehgcttd9hxw5n9wd6xz6m995w4zm"

	// StakingTransferAddr is the staking transfer address
	StakingTransferAddr = "io1qqqqqqqqwd6xz6mfden4gunpdeekvetjzwh99y"

	// StakingRegisterCandAddr is the staking register candidate address
	StakingRegisterCandAddr = "io1qpehgcttd9hxw5n9va5hxar9wfpkzmnyahxhjk"

	// StakingUpdateCandAddr is the staking update candidate address
	StakingUpdateCandAddr = "io1qqqqqum5v94kjmn824cxgct5v4pkzmnyxxj98n"
)

// 20-byte protocol address hash
var (
	StakingProtocolAddrHash     = hash160b([]byte("staking"))
	RewardingProtocolAddrHash   = hash160b([]byte("rewarding"))
	StakingCreateAddrHash       = bytesToHash160([]byte("stakingCreate"))
	StakingAddDepositAddrHash   = bytesToHash160([]byte("stakingAddDeposit"))
	StakingChangeCandAddrHash   = bytesToHash160([]byte("stakingChangeCand"))
	StakingUnstakeAddrHash      = bytesToHash160([]byte("stakingUnstake"))
	StakingWithdrawAddrHash     = bytesToHash160([]byte("stakingWithdraw"))
	StakingRestakeAddrHash      = bytesToHash160([]byte("stakingRestake"))
	StakingTransferAddrHash     = bytesToHash160([]byte("stakingTransfer"))
	StakingRegisterCandAddrHash = bytesToHash160([]byte("stakingRegisterCand"))
	StakingUpdateCandAddrHash   = bytesToHash160([]byte("stakingUpdateCand"))
)

// hash160b returns 160-bit (20-byte) hash of input
func hash160b(input []byte) Hash160 {
	// use keccak algorithm
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(input)
	var hash Hash160
	copy(hash[:], hasher.Sum(nil)[12:])
	return hash
}

// bytesToHash160 copies the byte slice into hash
func bytesToHash160(b []byte) Hash160 {
	var h Hash160
	if len(b) > 20 {
		b = b[len(b)-20:]
	}
	copy(h[20-len(b):], b)
	return h
}

// AddrV1Special is address that consists of special text
// it is NOT a valid bech32 encoding of the 20-bytes hash
type AddrV1Special struct {
	addr string
}

func newAddrV1Special(s string) *AddrV1Special {
	return &AddrV1Special{
		addr: s,
	}
}

// IsAddrV1Special returns true for special address
func IsAddrV1Special(s string) bool {
	switch s {
	case RewardingPoolAddr, StakingBucketPoolAddr:
		return true
	default:
		return false
	}
}

// String returns the special-text address
func (addr *AddrV1Special) String() string { return addr.addr }

// Bytes panics since it is NOT a valid bech32 encoding
func (addr *AddrV1Special) Bytes() []byte {
	panic("Bytes() does not apply for special address")
}

// Hex panics since it is NOT a valid bech32 encoding
func (addr *AddrV1Special) Hex() string {
	panic("Hex() does not apply for special address")
}
