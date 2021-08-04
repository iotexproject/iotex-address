package address

import "golang.org/x/crypto/sha3"

const (
	// ZeroAddress is the IoTeX address whose hash160 is all zero
	ZeroAddress = "io1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqd39ym7"

	// StakingBucketPoolAddr is the staking bucket pool address
	StakingBucketPoolAddr = "io000000000000000000000000stakingprotocol"

	// StakingCreateAddr is the special address for staking create actions
	StakingCreateActionAddr = "io000000000000000000000000web3stakecreate"

	// StakingAddDepositActionAddr is the special address for staking add deposit actions
	StakingAddDepositActionAddr = "io00000000000000000000web3stakeadddeposit"

	// StakingChangeCandidateActionAddr is the special address for staking change candidate actions
	StakingChangeCandidateActionAddr = "io000000000000000web3stakechangecandidate"

	// StakingReclaimActionAddr is the special address for staking reclaim actions
	StakingReclaimActionAddr = "io00000000000000000000000web3stakereclaim"

	// StakingRestakeActionAddr is the special address for staking restake actions
	StakingRestakeActionAddr = "io00000000000000000000000web3stakerestake"

	// StakingTransferStakeAddr is the special address for staking transfer stake actions
	StakingTransferStakeAddr = "io00000000000000000web3staketransferstake"

	// RewardingPoolAddr is the rewarding pool address
	RewardingPoolAddr = "io0000000000000000000000rewardingprotocol"
)

// 20-byte protocol address hash
var (
	StakingProtocolAddrHash              = hash160b([]byte("staking"))
	RewardingProtocolAddrHash            = hash160b([]byte("rewarding"))
	StakingCreateActionAddrHash          = hash160b([]byte("stakingCreateAct"))
	StakingAddDepositActionAddrHash      = hash160b([]byte("stakingAddDepositAct"))
	StakingChangeCandidateActionAddrHash = hash160b([]byte("stakingChangeCandidateAct"))
	StakingReclaimActionAddrHash         = hash160b([]byte("stakingReclaimAct"))
	StakingRestakeActionAddrHash         = hash160b([]byte("stakingRestakeAct"))
	StakingTransferStakeActionAddrHash   = hash160b([]byte("stakingTransferStakeAct"))
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
