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

	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
)

func (rd *RewardDistributor) distributeNewBlockReward(b BalanceAdder, header *types.Header, totalTxFee *big.Int, rewardConfig *rewardConfig, stakingInfo *StakingInfo) {
	proposer := header.Rewardbase

	basic, stake, kgf, kir, splitRemaining := rd.splitReward(rewardConfig, totalTxFee)
	minStaking, err := rd.gh.GetMinimumStakingAtNumber(header.Number.Uint64())
	if err != nil {
		logger.Warn("distributeBlockReward failed", "err", err)
		return
	}
	shares, shareRemaining := rd.calcStakeShares(stake, minStaking, rewardConfig, stakingInfo)

	kgf.Add(kgf, splitRemaining)
	kgf.Add(kgf, shareRemaining)

	b.AddBalance(proposer, basic)
	b.AddBalance(stakingInfo.PoCAddr, kgf)
	b.AddBalance(stakingInfo.KIRAddr, kir)
	for rewardee, amount := range shares {
		b.AddBalance(rewardee, amount)
	}
}

func (rd *RewardDistributor) splitReward(config *rewardConfig, totalTxFee *big.Int) (basic, stake, kgf, kir, remaining *big.Int) {
	minted := config.mintingAmount

	tmpInt := big.NewInt(0)

	// 1 - split minted
	tmpInt.Mul(minted, config.cnRatio)
	cn := big.NewInt(0).Div(tmpInt, config.totalRatio)

	tmpInt.Mul(minted, config.pocRatio)
	kgf = big.NewInt(0).Div(tmpInt, config.totalRatio)

	tmpInt.Mul(minted, config.kirRatio)
	kir = big.NewInt(0).Div(tmpInt, config.totalRatio)

	// 2 - split cn
	tmpInt.Mul(cn, config.kip82.basicRatio)
	basic = big.NewInt(0).Div(tmpInt, config.kip82.totalRatio)

	tmpInt.Mul(cn, config.kip82.stakeRatio)
	stake = big.NewInt(0).Div(tmpInt, config.kip82.totalRatio)

	// remaining = minted - basic - stake - kgf - kir
	remaining = tmpInt.Sub(minted, basic)
	remaining = tmpInt.Sub(remaining, stake)
	remaining = tmpInt.Sub(remaining, kgf)
	remaining = tmpInt.Sub(remaining, kir)

	if totalTxFee.Cmp(basic) > 0 {
		basic = totalTxFee
	}

	logger.Info("splitReward", "basic", basic.Uint64(),
		"stake", stake.Uint64(),
		"kgf", kgf.Uint64(),
		"kir", kir.Uint64(),
		"remaining", remaining.Uint64())

	return
}

func (rd *RewardDistributor) calcStakeShares(stake *big.Int, minStaking uint64, config *rewardConfig, stakingInfo *StakingInfo) (map[common.Address]*big.Int, *big.Int) {
	cns := stakingInfo.GetConsolidatedStakingInfo()
	totalStake := big.NewInt(0)

	for _, node := range cns.GetAllNodes() {
		if node.StakingAmount > minStaking {
			diff := node.StakingAmount - minStaking
			diffBig := big.NewInt(0).SetUint64(diff)
			totalStake.Add(totalStake, diffBig)
		}
	}

	remaining := new(big.Int)
	remaining.Set(stake)

	shares := make(map[common.Address]*big.Int)

	for _, node := range cns.GetAllNodes() {
		if node.StakingAmount > minStaking {
			diff := node.StakingAmount - minStaking
			diffBig := big.NewInt(0).SetUint64(diff)

			tmpInt := big.NewInt(0).Mul(stake, diffBig)
			reward := big.NewInt(0).Div(tmpInt, totalStake)
			shares[node.RewardAddr] = reward
			remaining.Sub(remaining, reward)
		}
	}
	logger.Info("calcStakeShares", "minStaking", minStaking,
		"stake", stake.Uint64(),
		"remaining", remaining.Uint64(),
		"shares", shares)

	return shares, remaining
}
