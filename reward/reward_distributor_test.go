// Copyright 2019 The klaytn Authors
// This file is part of the klaytn library.
//
// The klaytn library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The klaytn library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the klaytn library. If not, see <http://www.gnu.org/licenses/>.

package reward

import (
	"math/big"
	"testing"

	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/params"
	"github.com/stretchr/testify/assert"
)

type testBalanceAdder struct {
	accounts map[common.Address]*big.Int
}

func newTestBalanceAdder() *testBalanceAdder {
	balanceAdder := &testBalanceAdder{}
	balanceAdder.accounts = make(map[common.Address]*big.Int)
	return balanceAdder
}

func (balanceAdder *testBalanceAdder) AddBalance(addr common.Address, v *big.Int) {
	balance, ok := balanceAdder.accounts[addr]
	if ok {
		balanceAdder.accounts[addr] = big.NewInt(0).Add(balance, v)
	} else {
		balanceAdder.accounts[addr] = v
	}
}

func (balanceAdder *testBalanceAdder) GetBalance(addr common.Address) *big.Int {
	balance, ok := balanceAdder.accounts[addr]
	if ok {
		return balance
	} else {
		return nil
	}
}

func Test_isEmptyAddress(t *testing.T) {
	testCases := []struct {
		address common.Address
		result  bool
	}{
		{
			common.Address{},
			true,
		},
		{
			common.HexToAddress("0x0000000000000000000000000000000000000000"),
			true,
		},
		{
			common.StringToAddress("0xA75Ed91f789BF9dc121DACB822849955ca3AB6aD"),
			false,
		},
		{
			common.StringToAddress("0x4bCDd8E3F9776d16056815E189EcB5A8bF8E4CBb"),
			false,
		},
	}
	for _, testCase := range testCases {
		assert.Equal(t, testCase.result, common.EmptyAddress(testCase.address))
	}
}

func TestRewardDistributor_getTotalTxFee(t *testing.T) {
	testCases := []struct {
		gasUsed            uint64
		unitPrice          *big.Int
		baseFee            *big.Int
		expectedTotalTxFee *big.Int
	}{
		// before magma hardfork
		{0, big.NewInt(25000000000), nil, big.NewInt(0)},
		{200000, big.NewInt(25000000000), nil, big.NewInt(5000000000000000)},
		{200000, big.NewInt(25000000000), nil, big.NewInt(5000000000000000)},
		{129346, big.NewInt(10000000000), nil, big.NewInt(1293460000000000)},
		{129346, big.NewInt(10000000000), nil, big.NewInt(1293460000000000)},
		{9236192, big.NewInt(50000), nil, big.NewInt(461809600000)},
		{9236192, big.NewInt(50000), nil, big.NewInt(461809600000)},
		{12936418927364923, big.NewInt(0), nil, big.NewInt(0)},
		// after magma hardfork, unitprice ignored
		{0, big.NewInt(25000000000), big.NewInt(25000000000), big.NewInt(0)},
		{200000, big.NewInt(25000000000), big.NewInt(25000000000), big.NewInt(5000000000000000)},
		{200000, big.NewInt(25000000000), big.NewInt(25000000000), big.NewInt(5000000000000000)},
		{129346, big.NewInt(25000000000), big.NewInt(10000000000), big.NewInt(1293460000000000)},
		{129346, big.NewInt(250), big.NewInt(10000000000), big.NewInt(1293460000000000)},
		{9236192, big.NewInt(9876), big.NewInt(50000), big.NewInt(461809600000)},
		{9236192, big.NewInt(25000000000), big.NewInt(50000), big.NewInt(461809600000)},
		{12936418927364923, big.NewInt(25000000000), big.NewInt(0), big.NewInt(0)},
	}
	rewardDistributor := NewRewardDistributor(newDefaultTestGovernance())
	rewardConfig := &rewardConfig{}

	header := &types.Header{}
	for _, testCase := range testCases {
		header.GasUsed = testCase.gasUsed
		header.BaseFee = testCase.baseFee
		rewardConfig.unitPrice = testCase.unitPrice

		result := rewardDistributor.getTotalTxFee(header, rewardConfig)
		assert.Equal(t, testCase.expectedTotalTxFee.Uint64(), result.Uint64())
	}
}

func TestRewardDistributor_TxFeeBurning(t *testing.T) {
	testCases := []struct {
		gasUsed            uint64
		unitPrice          *big.Int
		baseFee            *big.Int
		expectedTotalTxFee *big.Int
	}{
		{0, nil, big.NewInt(25000000000), big.NewInt(0)},
		{200000, nil, big.NewInt(25000000000), big.NewInt(5000000000000000 / 2)},
		{200000, nil, big.NewInt(25000000000), big.NewInt(5000000000000000 / 2)},
		{129346, nil, big.NewInt(10000000000), big.NewInt(1293460000000000 / 2)},
		{129346, nil, big.NewInt(10000000000), big.NewInt(1293460000000000 / 2)},
		{9236192, nil, big.NewInt(50000), big.NewInt(461809600000 / 2)},
		{9236192, nil, big.NewInt(50000), big.NewInt(461809600000 / 2)},
		{12936418927364923, nil, big.NewInt(0), big.NewInt(0)},
	}
	rewardDistributor := NewRewardDistributor(newDefaultTestGovernance())
	rewardConfig := &rewardConfig{}

	header := &types.Header{}

	for _, testCase := range testCases {
		header.GasUsed = testCase.gasUsed
		header.BaseFee = testCase.baseFee
		rewardConfig.unitPrice = testCase.baseFee

		txFee := rewardDistributor.getTotalTxFee(header, rewardConfig)
		burnedTxFee := rewardDistributor.txFeeBurning(txFee)
		assert.Equal(t, testCase.expectedTotalTxFee.Uint64(), burnedTxFee.Uint64())
	}
}

func TestRewardDistributor_MintKLAY(t *testing.T) {
	BalanceAdder := newTestBalanceAdder()
	header := &types.Header{}
	header.Number = big.NewInt(0)
	header.BaseFee = big.NewInt(30000000000)
	header.Rewardbase = common.StringToAddress("0x1552F52D459B713E0C4558e66C8c773a75615FA8")
	governance := newDefaultTestGovernance()
	rewardDistributor := NewRewardDistributor(governance)

	err := rewardDistributor.MintKLAY(BalanceAdder, header)
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	assert.NotNil(t, BalanceAdder.GetBalance(header.Rewardbase).Int64())
	assert.Equal(t, governance.Params().MintingAmountStr(), BalanceAdder.GetBalance(header.Rewardbase).String())
}

func TestRewardDistributor_distributeBlockReward(t *testing.T) {
	testCases := []struct {
		totalTxFee         *big.Int
		rewardConfig       *rewardConfig
		expectedCnBalance  *big.Int
		expectedPocBalance *big.Int
		expectedKirBalance *big.Int
	}{
		{
			totalTxFee: big.NewInt(0),
			rewardConfig: &rewardConfig{
				mintingAmount: big.NewInt(0).SetUint64(9600000000000000000),
				cnRatio:       big.NewInt(0).SetInt64(34),
				pocRatio:      big.NewInt(0).SetInt64(54),
				kirRatio:      big.NewInt(0).SetInt64(12),
				totalRatio:    big.NewInt(0).SetInt64(100),
				unitPrice:     big.NewInt(0).SetInt64(25000000000),
			},
			expectedCnBalance:  big.NewInt(0).SetUint64(3264000000000000000),
			expectedPocBalance: big.NewInt(0).SetUint64(5184000000000000000),
			expectedKirBalance: big.NewInt(0).SetUint64(1152000000000000000),
		},
		{
			totalTxFee: big.NewInt(1000000),
			rewardConfig: &rewardConfig{
				mintingAmount: big.NewInt(0).SetUint64(10000000000),
				cnRatio:       big.NewInt(0).SetInt64(60),
				pocRatio:      big.NewInt(0).SetInt64(30),
				kirRatio:      big.NewInt(0).SetInt64(10),
				totalRatio:    big.NewInt(0).SetInt64(100),
				unitPrice:     big.NewInt(0).SetInt64(25000000000),
			},
			expectedCnBalance:  big.NewInt(0).SetUint64(6000600000),
			expectedPocBalance: big.NewInt(0).SetUint64(3000300000),
			expectedKirBalance: big.NewInt(0).SetUint64(1000100000),
		},
	}

	header := &types.Header{}
	header.BaseFee = big.NewInt(30000000000)
	header.Number = big.NewInt(0)
	header.Rewardbase = common.StringToAddress("0x1552F52D459B713E0C4558e66C8c773a75615FA8")
	pocAddress := common.StringToAddress("0x4bCDd8E3F9776d16056815E189EcB5A8bF8E4CBb")
	kirAddress := common.StringToAddress("0xd38A08AD21B44681f5e75D0a3CA4793f3E6c03e7")
	governance := newDefaultTestGovernance()

	for _, testCase := range testCases {
		BalanceAdder := newTestBalanceAdder()
		rewardDistributor := NewRewardDistributor(governance)
		rewardDistributor.distributeBlockReward(BalanceAdder, header, testCase.totalTxFee, testCase.rewardConfig, pocAddress, kirAddress)

		assert.Equal(t, testCase.expectedCnBalance.Uint64(), BalanceAdder.GetBalance(header.Rewardbase).Uint64())
		assert.Equal(t, testCase.expectedPocBalance.Uint64(), BalanceAdder.GetBalance(pocAddress).Uint64())
		assert.Equal(t, testCase.expectedKirBalance.Uint64(), BalanceAdder.GetBalance(kirAddress).Uint64())
	}
}

func TestRewardDistributor_DistributeBlockReward(t *testing.T) {
	testCases := []struct {
		gasUsed            uint64
		baseFee            *big.Int
		config             map[int]interface{}
		expectedCnBalance  *big.Int
		expectedPocBalance *big.Int
		expectedKirBalance *big.Int
	}{
		{
			gasUsed: 100,
			baseFee: big.NewInt(500),
			config: map[int]interface{}{
				params.Epoch:         30,
				params.MintingAmount: "50000",
				params.Ratio:         "40/50/10",
				params.UnitPrice:     25000000000,
				params.UseGiniCoeff:  true,
				params.DeferredTxFee: true,
			},
			expectedCnBalance:  big.NewInt(0).SetUint64(30000),
			expectedPocBalance: big.NewInt(0).SetUint64(37500),
			expectedKirBalance: big.NewInt(0).SetUint64(7500),
		},
		{
			gasUsed: 100,
			config: map[int]interface{}{
				params.Epoch:         30,
				params.MintingAmount: "50000",
				params.Ratio:         "40/50/10",
				params.UnitPrice:     25000000000,
				params.UseGiniCoeff:  true,
				params.DeferredTxFee: true,
			},
			expectedCnBalance:  big.NewInt(0).SetUint64(250000005000 * 4),
			expectedPocBalance: big.NewInt(0).SetUint64(250000005000 * 5),
			expectedKirBalance: big.NewInt(0).SetUint64(250000005000),
		},
		{
			gasUsed: 0,
			baseFee: big.NewInt(25000000000),
			config: map[int]interface{}{
				params.Epoch:         604800,
				params.MintingAmount: "9600000000000000000",
				params.Ratio:         "34/54/12",
				params.UnitPrice:     25000000000,
				params.UseGiniCoeff:  true,
				params.DeferredTxFee: true,
			},
			expectedCnBalance:  big.NewInt(0).SetUint64(3264000000000000000),
			expectedPocBalance: big.NewInt(0).SetUint64(5184000000000000000),
			expectedKirBalance: big.NewInt(0).SetUint64(1152000000000000000),
		},
		{
			gasUsed: 0,
			config: map[int]interface{}{
				params.Epoch:         3600,
				params.MintingAmount: "0",
				params.Ratio:         "100/0/0",
				params.UnitPrice:     25000000000,
				params.UseGiniCoeff:  true,
				params.DeferredTxFee: true,
			},
			expectedCnBalance:  big.NewInt(0).SetUint64(0),
			expectedPocBalance: big.NewInt(0).SetUint64(0),
			expectedKirBalance: big.NewInt(0).SetUint64(0),
		},
	}

	header := &types.Header{}
	header.Number = big.NewInt(0)
	header.Rewardbase = common.StringToAddress("0x1552F52D459B713E0C4558e66C8c773a75615FA8")
	pocAddress := common.StringToAddress("0x4bCDd8E3F9776d16056815E189EcB5A8bF8E4CBb")
	kirAddress := common.StringToAddress("0xd38A08AD21B44681f5e75D0a3CA4793f3E6c03e7")
	governance := newDefaultTestGovernance()

	for _, testCase := range testCases {
		BalanceAdder := newTestBalanceAdder()
		governance.setTestGovernance(testCase.config)
		header.GasUsed = testCase.gasUsed
		header.BaseFee = testCase.baseFee
		rewardDistributor := NewRewardDistributor(governance)

		err := rewardDistributor.DistributeBlockReward(BalanceAdder, header, pocAddress, kirAddress)
		if !assert.NoError(t, err) {
			t.FailNow()
		}

		assert.NotNil(t, BalanceAdder.GetBalance(header.Rewardbase).Int64())
		assert.Equal(t, testCase.expectedCnBalance.Uint64(), BalanceAdder.GetBalance(header.Rewardbase).Uint64())
		assert.Equal(t, testCase.expectedPocBalance.Uint64(), BalanceAdder.GetBalance(pocAddress).Uint64())
		assert.Equal(t, testCase.expectedKirBalance.Uint64(), BalanceAdder.GetBalance(kirAddress).Uint64())
	}
}
