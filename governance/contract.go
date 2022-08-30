package governance

import (
	"errors"
	"math/big"

	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/contracts/reward/contract"
	"github.com/klaytn/klaytn/params"
)

var GOVPARAM_NAME = [32]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x47, 0x6f, 0x76, 0x50, 0x61, 0x72, 0x61, 0x6d}

var errContractEngineNotReady = errors.New("ContractEngine is not ready")

type ContractEngine struct {
	config        *params.ChainConfig
	currentParams *params.GovParamSet

	chain blockChain // To access the contract state DB
}

func NewContractEngine(config *params.ChainConfig) *ContractEngine {
	e := &ContractEngine{
		config:        config,
		currentParams: params.NewGovParamSet(),
	}

	return e
}

func (e *ContractEngine) SetBlockchain(chain blockChain) {
	e.chain = chain
}

// Params effective at upcoming block (head+1)
func (e *ContractEngine) Params() *params.GovParamSet {
	return e.currentParams
}

// Params effective at requested block (num)
func (e *ContractEngine) ParamsAt(num uint64) (*params.GovParamSet, error) {
	if e.chain == nil {
		logger.Error("Invoked ParamsAt() before SetBlockchain", "num", num)
		return params.NewGovParamSet(), errContractEngineNotReady
	}

	head := e.chain.CurrentHeader().Number.Uint64()
	if num > head {
		// Sometimes future blocks are requested.
		// ex) reward distributor in istanbul.engine.Finalize() requests ParamsAt(head+1)
		// ex) governance_itemsAt(num) API requests arbitrary num
		// In those cases we refer to the head block.
		num = head + 1
	}

	pset, err := e.contractGetAllParams(num)
	if err != nil {
		return params.NewGovParamSet(), err
	}
	return pset, nil
}

// if UpdateParam fails, leave currentParams as-is
func (e *ContractEngine) UpdateParams() error {
	if e.chain == nil {
		logger.Error("Invoked UpdateParams() before SetBlockchain")
		return errContractEngineNotReady
	}

	head := e.chain.CurrentHeader().Number.Uint64()
	pset, err := e.contractGetAllParams(head + 1)
	if err != nil {
		return err
	}

	e.currentParams = pset
	return nil
}

func (e *ContractEngine) contractGetAllParams(num uint64) (*params.GovParamSet, error) {
	if e.chain == nil {
		logger.Error("Invoked ContractEngine before SetBlockchain")
		return nil, errContractEngineNotReady
	}

	addr := e.contractAddrAt(num)
	if common.EmptyAddress(addr) {
		logger.Error("Invoked ContractEngine but address is empty", "num", num)
		return nil, errContractEngineNotReady
	}

	caller := &contractCaller{
		chainConfig:  e.config,
		chain:        e.chain,
		contractAddr: addr,
	}
	if num > 0 {
		num -= 1
	}
	return caller.getAllParams(new(big.Int).SetUint64(num))
}

// Return the GovernanceContract address effective at given block number
func (e *ContractEngine) contractAddrAt(num uint64) common.Address {
	if !e.config.IsContractGovForkEnabled(new(big.Int).SetUint64(num)) {
		return common.Address{}
	}

	caller := &contractCaller{
		chainConfig:  e.config,
		chain:        e.chain,
		contractAddr: common.HexToAddress(contract.AddressBookContractAddress),
	}
	if num > 0 {
		num -= 1
	}

	regAddr, err := caller.getRegistryAt(new(big.Int).SetUint64(num))
	if err != nil {
		return common.Address{}
	}

	caller = &contractCaller{
		chainConfig:  e.config,
		chain:        e.chain,
		contractAddr: regAddr,
	}

	gpAddr, err := caller.getAddressAt(new(big.Int).SetUint64(num), GOVPARAM_NAME)
	if err != nil {
		return common.Address{}
	}
	return gpAddr
}
