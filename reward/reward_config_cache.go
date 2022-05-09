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
	"errors"
	"math/big"
	"strconv"
	"strings"

	lru "github.com/hashicorp/golang-lru"
)

var (
	errFailGettingConfigure = errors.New("fail to get configure from governance")
	errInvalidFormat        = errors.New("invalid format")
	errParsingRatio         = errors.New("parsing ratio fail")
)

const (
	maxRewardConfigCache = 3
)

type rewardConfig struct {
	blockNum      uint64
	mintingAmount *big.Int
	cnRatio       *big.Int
	pocRatio      *big.Int
	kirRatio      *big.Int
	totalRatio    *big.Int
	unitPrice     *big.Int
}

// Cache for parsed reward parameters from governance
type rewardConfigCache struct {
	cache            *lru.ARCCache
	governanceHelper governanceHelper
}

func newRewardConfigCache(governanceHelper governanceHelper) *rewardConfigCache {
	cache, _ := lru.NewARC(maxRewardConfigCache)
	return &rewardConfigCache{
		cache:            cache,
		governanceHelper: governanceHelper,
	}
}

func (rewardConfigCache *rewardConfigCache) get(blockNumber uint64) (*rewardConfig, error) {
	govParams, err := rewardConfigCache.governanceHelper.ParamsAt(blockNumber)
	if err != nil {
		return nil, err
	}

	epoch := govParams.Epoch()
	remainder := blockNumber % epoch
	if remainder == 0 {
		blockNumber -= epoch
	} else {
		blockNumber -= remainder
	}

	config, ok := rewardConfigCache.cache.Get(blockNumber)
	if ok {
		return config.(*rewardConfig), nil
	}

	newConfig, err := rewardConfigCache.newRewardConfig(blockNumber)
	if err != nil {
		return nil, err
	}

	rewardConfigCache.add(blockNumber, newConfig)
	return newConfig, nil
}

func (rewardConfigCache *rewardConfigCache) newRewardConfig(blockNumber uint64) (*rewardConfig, error) {
	govParams, err := rewardConfigCache.governanceHelper.ParamsAt(blockNumber)
	if err != nil {
		return nil, err
	}

	mintingAmount := govParams.MintingAmountBig()

	ratio := govParams.Ratio()
	cn, poc, kir, parsingError := rewardConfigCache.parseRewardRatio(ratio)
	if parsingError != nil {
		return nil, parsingError
	}
	cnRatio := big.NewInt(int64(cn))
	pocRatio := big.NewInt(int64(poc))
	kirRatio := big.NewInt(int64(kir))
	totalRatio := big.NewInt(int64(cn + poc + kir))

	unitPriceUint64 := govParams.UnitPrice()
	unitPrice := new(big.Int).SetUint64(unitPriceUint64)

	rewardConfig := &rewardConfig{
		blockNum:      blockNumber,
		mintingAmount: mintingAmount,
		cnRatio:       cnRatio,
		pocRatio:      pocRatio,
		kirRatio:      kirRatio,
		totalRatio:    totalRatio,
		unitPrice:     unitPrice,
	}
	return rewardConfig, nil
}

func (rewardConfigCache *rewardConfigCache) add(blockNumber uint64, config *rewardConfig) {
	rewardConfigCache.cache.Add(blockNumber, config)
}

func (rewardConfigCache *rewardConfigCache) parseRewardRatio(ratio string) (int, int, int, error) {
	s := strings.Split(ratio, "/")
	if len(s) != 3 {
		return 0, 0, 0, errInvalidFormat
	}
	cn, err1 := strconv.Atoi(s[0])
	poc, err2 := strconv.Atoi(s[1])
	kir, err3 := strconv.Atoi(s[2])

	if err1 != nil || err2 != nil || err3 != nil {
		return 0, 0, 0, errParsingRatio
	}
	return cn, poc, kir, nil
}
