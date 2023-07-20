// Modifications Copyright 2018 The klaytn Authors
// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.
//
// This file is derived from quorum/consensus/istanbul/core/commit.go (2018/06/04).
// Modified and improved for the klaytn development.
package core

import (
	"math/big"
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
	vrankPrepreparedTime      = time.Now()
	vrankLateThreshold, _     = time.ParseDuration(vrankDefaultLateThreshold)
	vrankLateCommitView       = istanbul.View{
		Sequence: big.NewInt(0),
		Round:    big.NewInt(0),
	}
	vrankCommitArrivalTimeMap = make(map[common.Address]time.Duration)
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

func vrankAtPreprepare(view *istanbul.View) {
	/*
			lateCommits = filter CommitArrivalTimeMap whose value makes isLateCommittedSeal true
		    if round is 0: // last proposal was finalized
		        encode lateCommits into log format
		        logger.Info("VRank", "bitmap[committesizebit] bitmap[committesizebit] {500 340 600 350 ...}")

		    lastCommit = find the maximum value of CommitArrivalTimeMap
		    LastCommitArrivalTimeMetrics.Update(lastCommit)
	*/

	logger.Info("VRank", "vrankCommitArrivalTimeMap", vrankCommitArrivalTimeMap)
	lateCommits := filterLateCommits(vrankCommitArrivalTimeMap)
	if view.Round.Cmp(common.Big0) == 0 {
		// TODO-VRANK: encode
		logger.Info("VRank", "lateCommits", lateCommits)
	}

	lastCommit := time.Duration(0)
	for _, v := range lateCommits {
		if v < lastCommit {
			lastCommit = v
		}
	}
	vrankLastCommitArrivalTimeGauge.Update(int64(lastCommit))

	vrankPrepreparedTime = time.Now()
	vrankLateThreshold, _ = time.ParseDuration(vrankDefaultLateThreshold)
	vrankLateCommitView = *view
	vrankCommitArrivalTimeMap = make(map[common.Address]time.Duration)
}

func vrankAtCommit() {
	committedTime := timeSincePreprepare()
	if vrankLateThreshold > committedTime {
		vrankLateThreshold = committedTime
	}
	logger.Info("VRank", "threshold", vrankLateThreshold)
	vrankQuorumCommitArrivalTimeGauge.Update(1)          // TODO-VRANK: fix number
	vrankAvgCommitArrivalTimeWithinQuorumGauge.Update(1) // TODO-VRANK: fix number
}
