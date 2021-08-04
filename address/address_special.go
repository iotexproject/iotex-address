package address

import "golang.org/x/crypto/sha3"

const (
	// ZeroAddress is the IoTeX address whose hash160 is all zero
	ZeroAddress = "io1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqd39ym7"

	// StakingBucketPoolAddr is the staking bucket pool address
	StakingBucketPoolAddr = "io000000000000000000000000stakingprotocol"

	// RewardingPoolAddr is the rewarding pool address
	RewardingPoolAddr = "io0000000000000000000000rewardingprotocol"
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

type (
	// Hash160 for 160-bit hash used for account and smart contract address
	Hash160 [20]byte
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
