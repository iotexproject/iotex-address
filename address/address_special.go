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
)

// 20-byte protocol address hash
var (
	StakingProtocolAddrHash   = hash160b([]byte("staking"))
	RewardingProtocolAddrHash = hash160b([]byte("rewarding"))
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

// IsAddrV1Special returns true for special address
func IsAddrV1Special(s string) bool {
	switch s {
	case RewardingPoolAddr, StakingBucketPoolAddr:
		return true
	default:
		return false
	}
}
