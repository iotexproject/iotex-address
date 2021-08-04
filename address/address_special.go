package address

import "golang.org/x/crypto/sha3"

const (
	// ZeroAddress is the IoTeX address whose hash160 is all zero
	ZeroAddress = "io1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqd39ym7"

	// StakingBucketPoolAddr is the staking bucket pool address
	StakingBucketPoolAddr = "io000000000000000000000000stakingprotocol"

	// StakingCreateAddr is the special address for staking create actions
	StakingCreateAddr = "io000000000000000000000000web3stakecreate"

	// StakingAddDepositAddr is the special address for staking add deposit actions
	StakingAddDepositAddr = "io00000000000000000000web3stakeadddeposit"

	// StakingChangeCandidateAddr is the special address for staking change candidate actions
	StakingChangeCandidateAddr = "io000000000000000web3stakechangecandidate"

	// StakingUnstakeAddr is the special address for staking unstake actions
	StakingUnstakeAddr = "io00000000000000000000000web3stakeunstake"

	// StakingWithdrawAddr is the special address for staking withdraw actions
	StakingWithdrawAddr = "io0000000000000000000000web3stakewithdraw"

	// StakingRestakeAddr is the special address for staking restake actions
	StakingRestakeAddr = "io00000000000000000000000web3stakerestake"

	// StakingTransferStakeAddr is the special address for staking transfer stake actions
	StakingTransferStakeAddr = "io00000000000000000web3staketransferstake"

	// StakingCandidateRegisterAddr is the special address for staking candidate register actions
	StakingCandidateRegisterAddr = "io0000000000000web3stakecandidateregister"

	// StakingCandidateUpdateAddr is the special address for staking candidate update actions
	StakingCandidateUpdateAddr = "io000000000000000web3stakecandidateupdate"

	// RewardingPoolAddr is the rewarding pool address
	RewardingPoolAddr = "io0000000000000000000000rewardingprotocol"
)

// 20-byte protocol address hash
var (
	StakingProtocolAddrHash          = hash160b([]byte("staking"))
	RewardingProtocolAddrHash        = hash160b([]byte("rewarding"))
	StakingCreateAddrHash            = hash160b([]byte("stakingCreate"))
	StakingAddDepositAddrHash        = hash160b([]byte("stakingAddDeposit"))
	StakingChangeCandidateAddrHash   = hash160b([]byte("stakingChangeCandidate"))
	StakingUnstakeAddrHash           = hash160b([]byte("stakingUnstake"))
	StakingWithdrawAddrHash          = hash160b([]byte("stakingWithdraw"))
	StakingRestakeAddrHash           = hash160b([]byte("stakingRestake"))
	StakingTransferStakeAddrHash     = hash160b([]byte("stakingTransferStake"))
	StakingCandidateRegisterAddrHash = hash160b([]byte("stakingCandidateRegister"))
	StakingCandidateUpdateAddrHash   = hash160b([]byte("stakingCandidateUpdate"))
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
