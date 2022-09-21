package reward

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/log"
	"gotest.tools/assert"
)

var (
	cnBaseAddr     = 500
	stakeBaseAddr  = 600
	rewardBaseAddr = 700
	minStaking     = uint64(2000000) // changing this value will not change the governance's min staking
)

// rewardOverride[i] = j means rewards[i] = rewards[j]
func genStakingInfo(num int, rewardOverride map[int]int, amountOverride map[int]uint64) *StakingInfo {
	cns := make([]common.Address, 0)
	stakes := make([]common.Address, 0)
	rewards := make([]common.Address, 0)
	amounts := make([]uint64, 0)

	for i := 0; i < num; i++ {
		cns = append(cns, common.HexToAddress(fmt.Sprintf("0x%040d", cnBaseAddr+i)))
		stakes = append(stakes, common.HexToAddress(fmt.Sprintf("0x%040d", stakeBaseAddr+i)))
		rewards = append(rewards, common.HexToAddress(fmt.Sprintf("0x%040d", rewardBaseAddr+i)))
		amounts = append(amounts, minStaking)
	}

	for i := range rewardOverride {
		rewards[i] = rewards[rewardOverride[i]]
	}

	for i := range amountOverride {
		amounts[i] = amountOverride[i]
	}

	return &StakingInfo{
		BlockNum:              0,
		CouncilNodeAddrs:      cns,
		CouncilStakingAddrs:   stakes,
		CouncilRewardAddrs:    rewards,
		KIRAddr:               common.HexToAddress("0x0000000000000000000000000000000000001000"),
		PoCAddr:               common.HexToAddress("0x0000000000000000000000000000000000002000"),
		UseGini:               false,
		CouncilStakingAmounts: amounts,
	}
}

func TestRewardDistributor_distributeNewBlockReward(t *testing.T) {
	log.EnableLogForTest(log.LvlCrit, log.LvlInfo)

	rewardConfig := &rewardConfig{
		blockNum:      1,
		mintingAmount: big.NewInt(0).SetUint64(10000),
		cnRatio:       big.NewInt(0).SetInt64(34),
		pocRatio:      big.NewInt(0).SetInt64(54),
		kirRatio:      big.NewInt(0).SetInt64(12),
		totalRatio:    big.NewInt(0).SetInt64(100),
		unitPrice:     big.NewInt(0).SetInt64(25e9),
		kipxx: &KIPxxRewardConfig{
			basicRatio: big.NewInt(0).SetInt64(20),
			stakeRatio: big.NewInt(0).SetInt64(80),
			totalRatio: big.NewInt(0).SetInt64(100),
		},
	}

	testCases := []struct {
		totalTxFee           *big.Int
		stakingInfo          *StakingInfo
		expectedBasicBalance *big.Int
		expectedStakeBalance *big.Int
		expectedPocBalance   *big.Int
		expectedKirBalance   *big.Int
	}{
		{
			totalTxFee:           big.NewInt(500),
			stakingInfo:          genStakingInfo(1, nil, nil),
			expectedBasicBalance: big.NewInt(0).SetUint64(680),
			expectedKirBalance:   big.NewInt(0).SetUint64(1200),
			expectedPocBalance:   big.NewInt(0).SetUint64(8120), // mint = 5400, remainingStake = 2720
		},
		{
			totalTxFee: big.NewInt(500),
			stakingInfo: genStakingInfo(2,
				map[int]int{1: 0}, nil),
			expectedBasicBalance: big.NewInt(0).SetUint64(3400),
			expectedKirBalance:   big.NewInt(0).SetUint64(1200),
			expectedPocBalance:   big.NewInt(0).SetUint64(5400),
		},
		{
			totalTxFee: big.NewInt(500),
			stakingInfo: genStakingInfo(3,
				map[int]int{1: 0},
				map[int]uint64{2: minStaking * 2}),
			expectedBasicBalance: big.NewInt(0).SetUint64(2040),
			expectedKirBalance:   big.NewInt(0).SetUint64(1200),
			expectedPocBalance:   big.NewInt(0).SetUint64(5400),
		},
	}

	header := &types.Header{}
	header.BaseFee = big.NewInt(30000000000)
	header.Number = big.NewInt(0)
	header.Rewardbase = common.HexToAddress(fmt.Sprintf("0x%040d", rewardBaseAddr))

	governance := newDefaultTestGovernance()

	for _, testCase := range testCases {
		BalanceAdder := newTestBalanceAdder()
		rewardDistributor := NewRewardDistributor(governance)
		rewardDistributor.distributeNewBlockReward(BalanceAdder, header, testCase.totalTxFee, rewardConfig, testCase.stakingInfo)

		assert.Equal(t, testCase.expectedBasicBalance.Uint64(), BalanceAdder.GetBalance(header.Rewardbase).Uint64())
		assert.Equal(t, testCase.expectedKirBalance.Uint64(), BalanceAdder.GetBalance(testCase.stakingInfo.KIRAddr).Uint64())
		assert.Equal(t, testCase.expectedPocBalance.Uint64(), BalanceAdder.GetBalance(testCase.stakingInfo.PoCAddr).Uint64())
	}
}

func Benchmark_distributeNewBlockReward(b *testing.B) {
	// in the worst case, distribute stake shares among N
	amounts := make(map[int]uint64)
	N := 50
	for i := 0; i < N; i++ {
		amounts[i] = minStaking * 2
	}
	stakingInfo := genStakingInfo(N, nil, amounts)

	rewardConfig := &rewardConfig{
		blockNum:      1,
		mintingAmount: big.NewInt(0).SetUint64(10000),
		cnRatio:       big.NewInt(0).SetInt64(34),
		pocRatio:      big.NewInt(0).SetInt64(54),
		kirRatio:      big.NewInt(0).SetInt64(12),
		totalRatio:    big.NewInt(0).SetInt64(100),
		unitPrice:     big.NewInt(0).SetInt64(25e9),
		kipxx: &KIPxxRewardConfig{
			basicRatio: big.NewInt(0).SetInt64(20),
			stakeRatio: big.NewInt(0).SetInt64(80),
			totalRatio: big.NewInt(0).SetInt64(100),
		},
	}

	header := &types.Header{}
	header.BaseFee = big.NewInt(30000000000)
	header.Number = big.NewInt(0)
	header.Rewardbase = common.HexToAddress(fmt.Sprintf("0x%040d", rewardBaseAddr))
	totalTxFee := big.NewInt(30000000000)

	governance := newDefaultTestGovernance()

	BalanceAdder := newTestBalanceAdder()
	rewardDistributor := NewRewardDistributor(governance)

	b.ResetTimer()
	rewardDistributor.distributeNewBlockReward(BalanceAdder, header, totalTxFee, rewardConfig, stakingInfo)
}

func Benchmark_distributeBlockReward(b *testing.B) {
	// in the worst case, distribute stake shares among N
	amounts := make(map[int]uint64)
	N := 50
	for i := 0; i < N; i++ {
		amounts[i] = minStaking * 2
	}
	stakingInfo := genStakingInfo(N, nil, amounts)

	rewardConfig := &rewardConfig{
		blockNum:      1,
		mintingAmount: big.NewInt(0).SetUint64(10000),
		cnRatio:       big.NewInt(0).SetInt64(34),
		pocRatio:      big.NewInt(0).SetInt64(54),
		kirRatio:      big.NewInt(0).SetInt64(12),
		totalRatio:    big.NewInt(0).SetInt64(100),
		unitPrice:     big.NewInt(0).SetInt64(25e9),
		kipxx: &KIPxxRewardConfig{
			basicRatio: big.NewInt(0).SetInt64(20),
			stakeRatio: big.NewInt(0).SetInt64(80),
			totalRatio: big.NewInt(0).SetInt64(100),
		},
	}

	header := &types.Header{}
	header.BaseFee = big.NewInt(30000000000)
	header.Number = big.NewInt(0)
	header.Rewardbase = common.HexToAddress(fmt.Sprintf("0x%040d", rewardBaseAddr))
	totalTxFee := big.NewInt(30000000000)

	governance := newDefaultTestGovernance()

	BalanceAdder := newTestBalanceAdder()
	rewardDistributor := NewRewardDistributor(governance)

	b.ResetTimer()
	rewardDistributor.distributeBlockReward(BalanceAdder, header, totalTxFee, rewardConfig, stakingInfo.PoCAddr, stakingInfo.KIRAddr)
}
