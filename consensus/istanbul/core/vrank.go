// Copyright 2023 The klaytn Authors
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
package core

import (
	"encoding/hex"
	"math/big"
	"sort"
	"time"

	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/consensus/istanbul"
	"github.com/rcrowley/go-metrics"
)

var (
	// VRank metrics
	vrankFirstCommitArrivalTimeGauge           = metrics.NewRegisteredGauge("vrank/first_commit", nil)
	vrankQuorumCommitArrivalTimeGauge          = metrics.NewRegisteredGauge("vrank/quorum_commit", nil)
	vrankAvgCommitArrivalTimeWithinQuorumGauge = metrics.NewRegisteredGauge("vrank/avg_commit_within_quorum", nil)
	vrankLastCommitArrivalTimeGauge            = metrics.NewRegisteredGauge("vrank/last_commit", nil)

	vrankDefaultLateThreshold = "300ms"

	vrankPrepreparedTime  = time.Now()
	vrankCommittee        = istanbul.Validators{}
	vrankLateThreshold, _ = time.ParseDuration(vrankDefaultLateThreshold)
	vrankLateCommitView   = istanbul.View{
		Sequence: big.NewInt(0),
		Round:    big.NewInt(0),
	}
	vrankCommitArrivalTimeMap = make(map[common.Address]time.Duration)
)

const (
	vrankArrivedEarly = iota
	vrankArrivedLate
	vrankArrivalDidNotReceive
)

func isVrankTargetCommit(msg *istanbul.Subject, src istanbul.Validator) bool {
	if msg.View.Cmp(&vrankLateCommitView) != 0 {
		logger.Warn("isVrankTargetCommit=false because view does not match")
		return false
	}
	_, ok := vrankCommitArrivalTimeMap[src.Address()]
	if ok {
		logger.Warn("isVrankTargetCommit=false because already exists")
		return false
	}
	return true
}

func isVrankLateCommit(t time.Duration) bool {
	if t <= vrankLateThreshold {
		return false
	}
	return true
}

func timeSincePreprepare() time.Duration {
	return time.Now().Sub(vrankPrepreparedTime)
}

func filterLateCommits(src map[common.Address]time.Duration) map[common.Address]time.Duration {
	ret := make(map[common.Address]time.Duration)
	for k, v := range src {
		if isVrankLateCommit(v) {
			ret[k] = v
		}
	}
	return ret
}

func vrankCategorizeArrivalTimeMap(src map[common.Address]time.Duration) map[common.Address]int {
	kindList := make(map[common.Address]int, len(vrankCommittee))
	for _, validator := range vrankCommittee {
		time, ok := src[validator.Address()]
		var kind int
		if !ok {
			kind = vrankArrivalDidNotReceive
		} else {
			if isVrankLateCommit(time) {
				kind = vrankArrivedLate
			} else {
				kind = vrankArrivedEarly
			}
		}
		kindList[validator.Address()] = kind
	}
	return kindList
}

func vrankSerialize(valSet istanbul.Validators, m map[common.Address]int) []int {
	var sorted istanbul.Validators
	copy(sorted[:], valSet[:])
	sort.Sort(sorted)

	serialized := []int{len(m)}
	for i, v := range vrankCommittee {
		serialized[i] = m[v.Address()]
	}
	return serialized
}

func compressSerializedArrivals(arr []int) []byte {
	zip := func(a, b, c, d int) byte {
		a &= 0b11
		b &= 0b11
		c &= 0b11
		d &= 0b11
		return byte(a<<6 | b<<4 | c<<2 | d<<0)
	}

	// pad zero to make len(arr)%4 == 0
	switch len(arr) % 4 {
	case 1:
		arr = append(arr, []int{0, 0, 0}...)
	case 2:
		arr = append(arr, []int{0, 0}...)
	case 3:
		arr = append(arr, []int{0}...)
	}

	ret := make([]byte, 0)

	for i := 0; i < len(arr)/4; i++ {
		chunk := arr[4*i : 4*(i+1)]
		ret = append(ret, zip(chunk[0], chunk[1], chunk[2], chunk[3]))
	}
	return ret
}

func vrankLog() {
	categorized := vrankCategorizeArrivalTimeMap(vrankCommitArrivalTimeMap)
	serialized := vrankSerialize(vrankCommittee, categorized)
	bytes := compressSerializedArrivals(serialized)
	bitmap := hex.EncodeToString(bytes)
	logger.Info("VRank", "lateCommits", bitmap)
}

func vrankAtPreprepare(view *istanbul.View, committee istanbul.Validators) {
	/*
			lateCommits = filter CommitArrivalTimeMap whose value makes isLateCommittedSeal true
		    if round is 0: // last proposal was finalized
		        encode lateCommits into log format
		        logger.Info("VRank", "bitmap[committesizebit] bitmap[committesizebit] {500 340 600 350 ...}")
	*/

	vrankLog()

	lastCommit := time.Duration(0)
	lateCommits := filterLateCommits(vrankCommitArrivalTimeMap)
	for _, v := range lateCommits {
		if v < lastCommit {
			lastCommit = v
		}
	}
	vrankLastCommitArrivalTimeGauge.Update(int64(lastCommit))

	// Restart measure
	vrankPrepreparedTime = time.Now()
	vrankCommittee = committee
	vrankLateThreshold, _ = time.ParseDuration(vrankDefaultLateThreshold)
	vrankLateCommitView = *view
	vrankCommitArrivalTimeMap = make(map[common.Address]time.Duration)
}

func vrankAtCommit(blockNum *big.Int) {
	if vrankLateCommitView.Sequence.Cmp(blockNum) != 0 {
		// not expecting this block
		return
	}

	committedTime := timeSincePreprepare()
	if vrankLateThreshold > committedTime {
		vrankLateThreshold = committedTime
	}
	vrankQuorumCommitArrivalTimeGauge.Update(int64(committedTime))
	sum := int64(0)
	for _, v := range vrankCommitArrivalTimeMap {
		sum += int64(v)
	}
	avg := sum / int64(len(vrankCommitArrivalTimeMap))
	vrankAvgCommitArrivalTimeWithinQuorumGauge.Update(avg)
}
