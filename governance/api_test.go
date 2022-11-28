package governance

import (
	"math/big"
	"testing"

	"github.com/klaytn/klaytn/blockchain/state"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/consensus"
	"github.com/klaytn/klaytn/log"
	"github.com/klaytn/klaytn/networks/rpc"
	"github.com/klaytn/klaytn/params"
	"github.com/klaytn/klaytn/reward"
	"github.com/klaytn/klaytn/storage/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type blockchainTest struct {
	num uint64
}

func newTestBlockchain() *blockchainTest {
	return &blockchainTest{}
}

func newTestGovernanceApi() *PublicGovernanceAPI {
	config := params.CypressChainConfig
	config.Governance.KIP71 = params.GetDefaultKIP71Config()
	govApi := NewGovernanceAPI(NewMixedEngine(config, database.NewMemoryDBManager()))
	govApi.governance.SetNodeAddress(common.HexToAddress("0x52d41ca72af615a1ac3301b0a93efa222ecc7541"))
	return govApi
}

func TestUpperBoundBaseFeeSet(t *testing.T) {
	govApi := newTestGovernanceApi()

	curLowerBoundBaseFee := govApi.governance.Params().LowerBoundBaseFee()
	// unexpected case : upperboundbasefee < lowerboundbasefee
	invalidUpperBoundBaseFee := curLowerBoundBaseFee - 100
	_, err := govApi.Vote("kip71.upperboundbasefee", invalidUpperBoundBaseFee)
	assert.Equal(t, err, errInvalidUpperBound)
}

func TestLowerBoundFeeSet(t *testing.T) {
	govApi := newTestGovernanceApi()

	curUpperBoundBaseFee := govApi.governance.Params().UpperBoundBaseFee()
	// unexpected case : upperboundbasefee < lowerboundbasefee
	invalidLowerBoundBaseFee := curUpperBoundBaseFee + 100
	_, err := govApi.Vote("kip71.lowerboundbasefee", invalidLowerBoundBaseFee)
	assert.Equal(t, err, errInvalidLowerBound)
}

func TestGetRewards(t *testing.T) {
	type db = map[string]interface{}
	type expected = map[int]uint64
	type testcase struct {
		length   int // total number of blocks to simulate
		votes    map[int]db
		expected expected
	}

	mintAmount := uint64(1)
	koreBlock := uint64(9)
	epoch := 3

	testcases := []testcase{
		{
			12,
			map[int]db{
				6: {"reward.mintingamount": "2"},
				9: {"reward.mintingamount": "3"},
			},
			map[int]uint64{
				1:  1,
				2:  1,
				3:  1,
				4:  1,
				5:  1,
				6:  1,
				7:  2, // 2 is minted from now
				8:  2,
				9:  3, // 3 is minted from now
				10: 3,
				11: 3,
				12: 3,
				13: 3,
			},
		},
	}

	log.EnableLogForTest(log.LvlCrit, log.LvlDebug)

	config := getTestConfig()
	config.Governance.Reward.MintingAmount = new(big.Int).SetUint64(mintAmount)
	config.Istanbul.Epoch = uint64(epoch)
	config.KoreCompatibleBlock = new(big.Int).SetUint64(koreBlock)

	bc := newTestBlockchain()
	dbm := database.NewMemoryDBManager()
	pset, err := params.NewGovParamSetChainConfig(config)
	dbm.WriteGovernance(pset.StrMap(), 0)
	e := NewMixedEngine(config, dbm)
	e.SetBlockchain(bc)
	e.UpdateParams()
	govKlayApi := NewGovernanceKlayAPI(e, bc)
	latestNum := rpc.BlockNumber(-1)
	proposer := common.HexToAddress("0x0000000000000000000000000000000000000000")

	for _, tc := range testcases {
		// Place a vote if a vote is scheduled in upcoming block
		// Note that we're building (head+1)'th block here.
		for num := 0; num <= tc.length; num++ {
			for k, v := range tc.votes[num+1] {
				dbm.WriteGovernance(pset.StrMap(), 0)
			}

			rewardSpec, err := govKlayApi.GetRewards(&latestNum)
			assert.Nil(t, err)

			minted := new(big.Int).SetUint64(tc.expected[num+1])
			expectedRewardSpec := &reward.RewardSpec{
				Minted:   minted,
				TotalFee: common.Big0,
				BurntFee: common.Big0,
				Proposer: minted,
				Rewards: map[common.Address]*big.Int{
					proposer: minted,
				},
			}
			require.Equal(t, expectedRewardSpec, rewardSpec, "wrong at block %d", num+1)
			_ = proposer
		}
	}
}

func (bc *blockchainTest) Engine() consensus.Engine                         { return nil }
func (bc *blockchainTest) GetHeader(common.Hash, uint64) *types.Header      { return nil }
func (bc *blockchainTest) GetHeaderByNumber(val uint64) *types.Header       { return nil }
func (bc *blockchainTest) GetBlockByNumber(num uint64) *types.Block         { return nil }
func (bc *blockchainTest) StateAt(root common.Hash) (*state.StateDB, error) { return nil, nil }
func (bc *blockchainTest) Config() *params.ChainConfig                      { return nil }
func (bc *blockchainTest) CurrentHeader() *types.Header {
	return &types.Header{
		Number: new(big.Int).SetUint64(bc.num),
	}
}

func (bc *blockchainTest) SetBlockNum(num uint64) {
	bc.num = num
}

func (gov *govTest) Config() *params.ChainConfig { return nil }
