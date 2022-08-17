// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gov

import (
	"math/big"
	"strings"

	"github.com/klaytn/klaytn"
	"github.com/klaytn/klaytn/accounts/abi"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = klaytn.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IGovParamParam is an auto generated low-level Go binding around an user-defined struct.
type IGovParamParam struct {
	Activation uint64
	Votable    bool
	Prev       IGovParamParamState
	Next       IGovParamParamState
}

// IGovParamParamState is an auto generated low-level Go binding around an user-defined struct.
type IGovParamParamState struct {
	Value  []byte
	Exists bool
}

// AddressUpgradeableABI is the input ABI used to generate the binding from.
const AddressUpgradeableABI = "[]"

// AddressUpgradeableBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const AddressUpgradeableBinRuntime = `73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212207f8a515105bc81ab5cb407dd40722a02f1e686f694ae06880950f0ce28ba512b64736f6c634300080e0033`

// AddressUpgradeableBin is the compiled bytecode used for deploying new contracts.
var AddressUpgradeableBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212207f8a515105bc81ab5cb407dd40722a02f1e686f694ae06880950f0ce28ba512b64736f6c634300080e0033"

// DeployAddressUpgradeable deploys a new Klaytn contract, binding an instance of AddressUpgradeable to it.
func DeployAddressUpgradeable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AddressUpgradeable, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressUpgradeableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AddressUpgradeableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AddressUpgradeable{AddressUpgradeableCaller: AddressUpgradeableCaller{contract: contract}, AddressUpgradeableTransactor: AddressUpgradeableTransactor{contract: contract}, AddressUpgradeableFilterer: AddressUpgradeableFilterer{contract: contract}}, nil
}

// AddressUpgradeable is an auto generated Go binding around a Klaytn contract.
type AddressUpgradeable struct {
	AddressUpgradeableCaller     // Read-only binding to the contract
	AddressUpgradeableTransactor // Write-only binding to the contract
	AddressUpgradeableFilterer   // Log filterer for contract events
}

// AddressUpgradeableCaller is an auto generated read-only Go binding around a Klaytn contract.
type AddressUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUpgradeableTransactor is an auto generated write-only Go binding around a Klaytn contract.
type AddressUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUpgradeableFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type AddressUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUpgradeableSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type AddressUpgradeableSession struct {
	Contract     *AddressUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AddressUpgradeableCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type AddressUpgradeableCallerSession struct {
	Contract *AddressUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// AddressUpgradeableTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type AddressUpgradeableTransactorSession struct {
	Contract     *AddressUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// AddressUpgradeableRaw is an auto generated low-level Go binding around a Klaytn contract.
type AddressUpgradeableRaw struct {
	Contract *AddressUpgradeable // Generic contract binding to access the raw methods on
}

// AddressUpgradeableCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type AddressUpgradeableCallerRaw struct {
	Contract *AddressUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// AddressUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type AddressUpgradeableTransactorRaw struct {
	Contract *AddressUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressUpgradeable creates a new instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeable(address common.Address, backend bind.ContractBackend) (*AddressUpgradeable, error) {
	contract, err := bindAddressUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeable{AddressUpgradeableCaller: AddressUpgradeableCaller{contract: contract}, AddressUpgradeableTransactor: AddressUpgradeableTransactor{contract: contract}, AddressUpgradeableFilterer: AddressUpgradeableFilterer{contract: contract}}, nil
}

// NewAddressUpgradeableCaller creates a new read-only instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*AddressUpgradeableCaller, error) {
	contract, err := bindAddressUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeableCaller{contract: contract}, nil
}

// NewAddressUpgradeableTransactor creates a new write-only instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressUpgradeableTransactor, error) {
	contract, err := bindAddressUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeableTransactor{contract: contract}, nil
}

// NewAddressUpgradeableFilterer creates a new log filterer instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressUpgradeableFilterer, error) {
	contract, err := bindAddressUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeableFilterer{contract: contract}, nil
}

// bindAddressUpgradeable binds a generic wrapper to an already deployed contract.
func bindAddressUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressUpgradeable *AddressUpgradeableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AddressUpgradeable.Contract.AddressUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressUpgradeable *AddressUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.AddressUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressUpgradeable *AddressUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.AddressUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressUpgradeable *AddressUpgradeableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AddressUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressUpgradeable *AddressUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressUpgradeable *AddressUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// ContextUpgradeableABI is the input ABI used to generate the binding from.
const ContextUpgradeableABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"}]"

// ContextUpgradeableBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const ContextUpgradeableBinRuntime = ``

// ContextUpgradeable is an auto generated Go binding around a Klaytn contract.
type ContextUpgradeable struct {
	ContextUpgradeableCaller     // Read-only binding to the contract
	ContextUpgradeableTransactor // Write-only binding to the contract
	ContextUpgradeableFilterer   // Log filterer for contract events
}

// ContextUpgradeableCaller is an auto generated read-only Go binding around a Klaytn contract.
type ContextUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextUpgradeableTransactor is an auto generated write-only Go binding around a Klaytn contract.
type ContextUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextUpgradeableFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type ContextUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextUpgradeableSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type ContextUpgradeableSession struct {
	Contract     *ContextUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContextUpgradeableCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type ContextUpgradeableCallerSession struct {
	Contract *ContextUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ContextUpgradeableTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type ContextUpgradeableTransactorSession struct {
	Contract     *ContextUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ContextUpgradeableRaw is an auto generated low-level Go binding around a Klaytn contract.
type ContextUpgradeableRaw struct {
	Contract *ContextUpgradeable // Generic contract binding to access the raw methods on
}

// ContextUpgradeableCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type ContextUpgradeableCallerRaw struct {
	Contract *ContextUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// ContextUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type ContextUpgradeableTransactorRaw struct {
	Contract *ContextUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContextUpgradeable creates a new instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeable(address common.Address, backend bind.ContractBackend) (*ContextUpgradeable, error) {
	contract, err := bindContextUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeable{ContextUpgradeableCaller: ContextUpgradeableCaller{contract: contract}, ContextUpgradeableTransactor: ContextUpgradeableTransactor{contract: contract}, ContextUpgradeableFilterer: ContextUpgradeableFilterer{contract: contract}}, nil
}

// NewContextUpgradeableCaller creates a new read-only instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*ContextUpgradeableCaller, error) {
	contract, err := bindContextUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeableCaller{contract: contract}, nil
}

// NewContextUpgradeableTransactor creates a new write-only instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextUpgradeableTransactor, error) {
	contract, err := bindContextUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeableTransactor{contract: contract}, nil
}

// NewContextUpgradeableFilterer creates a new log filterer instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextUpgradeableFilterer, error) {
	contract, err := bindContextUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeableFilterer{contract: contract}, nil
}

// bindContextUpgradeable binds a generic wrapper to an already deployed contract.
func bindContextUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContextUpgradeable *ContextUpgradeableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ContextUpgradeable.Contract.ContextUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContextUpgradeable *ContextUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.ContextUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContextUpgradeable *ContextUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.ContextUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContextUpgradeable *ContextUpgradeableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ContextUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContextUpgradeable *ContextUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContextUpgradeable *ContextUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// ContextUpgradeableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContextUpgradeable contract.
type ContextUpgradeableInitializedIterator struct {
	Event *ContextUpgradeableInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContextUpgradeableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContextUpgradeableInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContextUpgradeableInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContextUpgradeableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContextUpgradeableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContextUpgradeableInitialized represents a Initialized event raised by the ContextUpgradeable contract.
type ContextUpgradeableInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContextUpgradeable *ContextUpgradeableFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContextUpgradeableInitializedIterator, error) {
	logs, sub, err := _ContextUpgradeable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeableInitializedIterator{contract: _ContextUpgradeable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContextUpgradeable *ContextUpgradeableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContextUpgradeableInitialized) (event.Subscription, error) {
	logs, sub, err := _ContextUpgradeable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContextUpgradeableInitialized)
				if err := _ContextUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContextUpgradeable *ContextUpgradeableFilterer) ParseInitialized(log types.Log) (*ContextUpgradeableInitialized, error) {
	event := new(ContextUpgradeableInitialized)
	if err := _ContextUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC1967UpgradeUpgradeableABI is the input ABI used to generate the binding from.
const ERC1967UpgradeUpgradeableABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"}]"

// ERC1967UpgradeUpgradeableBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const ERC1967UpgradeUpgradeableBinRuntime = ``

// ERC1967UpgradeUpgradeable is an auto generated Go binding around a Klaytn contract.
type ERC1967UpgradeUpgradeable struct {
	ERC1967UpgradeUpgradeableCaller     // Read-only binding to the contract
	ERC1967UpgradeUpgradeableTransactor // Write-only binding to the contract
	ERC1967UpgradeUpgradeableFilterer   // Log filterer for contract events
}

// ERC1967UpgradeUpgradeableCaller is an auto generated read-only Go binding around a Klaytn contract.
type ERC1967UpgradeUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1967UpgradeUpgradeableTransactor is an auto generated write-only Go binding around a Klaytn contract.
type ERC1967UpgradeUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1967UpgradeUpgradeableFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type ERC1967UpgradeUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1967UpgradeUpgradeableSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type ERC1967UpgradeUpgradeableSession struct {
	Contract     *ERC1967UpgradeUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ERC1967UpgradeUpgradeableCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type ERC1967UpgradeUpgradeableCallerSession struct {
	Contract *ERC1967UpgradeUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// ERC1967UpgradeUpgradeableTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type ERC1967UpgradeUpgradeableTransactorSession struct {
	Contract     *ERC1967UpgradeUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// ERC1967UpgradeUpgradeableRaw is an auto generated low-level Go binding around a Klaytn contract.
type ERC1967UpgradeUpgradeableRaw struct {
	Contract *ERC1967UpgradeUpgradeable // Generic contract binding to access the raw methods on
}

// ERC1967UpgradeUpgradeableCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type ERC1967UpgradeUpgradeableCallerRaw struct {
	Contract *ERC1967UpgradeUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// ERC1967UpgradeUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type ERC1967UpgradeUpgradeableTransactorRaw struct {
	Contract *ERC1967UpgradeUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC1967UpgradeUpgradeable creates a new instance of ERC1967UpgradeUpgradeable, bound to a specific deployed contract.
func NewERC1967UpgradeUpgradeable(address common.Address, backend bind.ContractBackend) (*ERC1967UpgradeUpgradeable, error) {
	contract, err := bindERC1967UpgradeUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeUpgradeable{ERC1967UpgradeUpgradeableCaller: ERC1967UpgradeUpgradeableCaller{contract: contract}, ERC1967UpgradeUpgradeableTransactor: ERC1967UpgradeUpgradeableTransactor{contract: contract}, ERC1967UpgradeUpgradeableFilterer: ERC1967UpgradeUpgradeableFilterer{contract: contract}}, nil
}

// NewERC1967UpgradeUpgradeableCaller creates a new read-only instance of ERC1967UpgradeUpgradeable, bound to a specific deployed contract.
func NewERC1967UpgradeUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*ERC1967UpgradeUpgradeableCaller, error) {
	contract, err := bindERC1967UpgradeUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeUpgradeableCaller{contract: contract}, nil
}

// NewERC1967UpgradeUpgradeableTransactor creates a new write-only instance of ERC1967UpgradeUpgradeable, bound to a specific deployed contract.
func NewERC1967UpgradeUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC1967UpgradeUpgradeableTransactor, error) {
	contract, err := bindERC1967UpgradeUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeUpgradeableTransactor{contract: contract}, nil
}

// NewERC1967UpgradeUpgradeableFilterer creates a new log filterer instance of ERC1967UpgradeUpgradeable, bound to a specific deployed contract.
func NewERC1967UpgradeUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC1967UpgradeUpgradeableFilterer, error) {
	contract, err := bindERC1967UpgradeUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeUpgradeableFilterer{contract: contract}, nil
}

// bindERC1967UpgradeUpgradeable binds a generic wrapper to an already deployed contract.
func bindERC1967UpgradeUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC1967UpgradeUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC1967UpgradeUpgradeable.Contract.ERC1967UpgradeUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1967UpgradeUpgradeable.Contract.ERC1967UpgradeUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1967UpgradeUpgradeable.Contract.ERC1967UpgradeUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC1967UpgradeUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1967UpgradeUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1967UpgradeUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// ERC1967UpgradeUpgradeableAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the ERC1967UpgradeUpgradeable contract.
type ERC1967UpgradeUpgradeableAdminChangedIterator struct {
	Event *ERC1967UpgradeUpgradeableAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC1967UpgradeUpgradeableAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1967UpgradeUpgradeableAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC1967UpgradeUpgradeableAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC1967UpgradeUpgradeableAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1967UpgradeUpgradeableAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1967UpgradeUpgradeableAdminChanged represents a AdminChanged event raised by the ERC1967UpgradeUpgradeable contract.
type ERC1967UpgradeUpgradeableAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*ERC1967UpgradeUpgradeableAdminChangedIterator, error) {
	logs, sub, err := _ERC1967UpgradeUpgradeable.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeUpgradeableAdminChangedIterator{contract: _ERC1967UpgradeUpgradeable.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *ERC1967UpgradeUpgradeableAdminChanged) (event.Subscription, error) {
	logs, sub, err := _ERC1967UpgradeUpgradeable.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1967UpgradeUpgradeableAdminChanged)
				if err := _ERC1967UpgradeUpgradeable.contract.UnpackLog(event, "AdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) ParseAdminChanged(log types.Log) (*ERC1967UpgradeUpgradeableAdminChanged, error) {
	event := new(ERC1967UpgradeUpgradeableAdminChanged)
	if err := _ERC1967UpgradeUpgradeable.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC1967UpgradeUpgradeableBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the ERC1967UpgradeUpgradeable contract.
type ERC1967UpgradeUpgradeableBeaconUpgradedIterator struct {
	Event *ERC1967UpgradeUpgradeableBeaconUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC1967UpgradeUpgradeableBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1967UpgradeUpgradeableBeaconUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC1967UpgradeUpgradeableBeaconUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC1967UpgradeUpgradeableBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1967UpgradeUpgradeableBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1967UpgradeUpgradeableBeaconUpgraded represents a BeaconUpgraded event raised by the ERC1967UpgradeUpgradeable contract.
type ERC1967UpgradeUpgradeableBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*ERC1967UpgradeUpgradeableBeaconUpgradedIterator, error) {
	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ERC1967UpgradeUpgradeable.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeUpgradeableBeaconUpgradedIterator{contract: _ERC1967UpgradeUpgradeable.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *ERC1967UpgradeUpgradeableBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {
	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ERC1967UpgradeUpgradeable.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1967UpgradeUpgradeableBeaconUpgraded)
				if err := _ERC1967UpgradeUpgradeable.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) ParseBeaconUpgraded(log types.Log) (*ERC1967UpgradeUpgradeableBeaconUpgraded, error) {
	event := new(ERC1967UpgradeUpgradeableBeaconUpgraded)
	if err := _ERC1967UpgradeUpgradeable.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC1967UpgradeUpgradeableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ERC1967UpgradeUpgradeable contract.
type ERC1967UpgradeUpgradeableInitializedIterator struct {
	Event *ERC1967UpgradeUpgradeableInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC1967UpgradeUpgradeableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1967UpgradeUpgradeableInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC1967UpgradeUpgradeableInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC1967UpgradeUpgradeableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1967UpgradeUpgradeableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1967UpgradeUpgradeableInitialized represents a Initialized event raised by the ERC1967UpgradeUpgradeable contract.
type ERC1967UpgradeUpgradeableInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) FilterInitialized(opts *bind.FilterOpts) (*ERC1967UpgradeUpgradeableInitializedIterator, error) {
	logs, sub, err := _ERC1967UpgradeUpgradeable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeUpgradeableInitializedIterator{contract: _ERC1967UpgradeUpgradeable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ERC1967UpgradeUpgradeableInitialized) (event.Subscription, error) {
	logs, sub, err := _ERC1967UpgradeUpgradeable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1967UpgradeUpgradeableInitialized)
				if err := _ERC1967UpgradeUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) ParseInitialized(log types.Log) (*ERC1967UpgradeUpgradeableInitialized, error) {
	event := new(ERC1967UpgradeUpgradeableInitialized)
	if err := _ERC1967UpgradeUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC1967UpgradeUpgradeableUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ERC1967UpgradeUpgradeable contract.
type ERC1967UpgradeUpgradeableUpgradedIterator struct {
	Event *ERC1967UpgradeUpgradeableUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC1967UpgradeUpgradeableUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1967UpgradeUpgradeableUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC1967UpgradeUpgradeableUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC1967UpgradeUpgradeableUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1967UpgradeUpgradeableUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1967UpgradeUpgradeableUpgraded represents a Upgraded event raised by the ERC1967UpgradeUpgradeable contract.
type ERC1967UpgradeUpgradeableUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ERC1967UpgradeUpgradeableUpgradedIterator, error) {
	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ERC1967UpgradeUpgradeable.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeUpgradeableUpgradedIterator{contract: _ERC1967UpgradeUpgradeable.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ERC1967UpgradeUpgradeableUpgraded, implementation []common.Address) (event.Subscription, error) {
	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ERC1967UpgradeUpgradeable.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1967UpgradeUpgradeableUpgraded)
				if err := _ERC1967UpgradeUpgradeable.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) ParseUpgraded(log types.Log) (*ERC1967UpgradeUpgradeableUpgraded, error) {
	event := new(ERC1967UpgradeUpgradeableUpgraded)
	if err := _ERC1967UpgradeUpgradeable.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovParamABI is the input ABI used to generate the binding from.
const GovParamABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"AddParam\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"DeleteParam\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"SetParam\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"SetParamVotable\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"addParam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"deleteParam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllParams\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllStructParams\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"activation\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"votable\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"internalType\":\"structIGovParam.ParamState\",\"name\":\"prev\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"internalType\":\"structIGovParam.ParamState\",\"name\":\"next\",\"type\":\"tuple\"}],\"internalType\":\"structIGovParam.Param[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getParam\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paramExistingNames\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paramNames\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"params\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"activation\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"votable\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"internalType\":\"structIGovParam.ParamState\",\"name\":\"prev\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"internalType\":\"structIGovParam.ParamState\",\"name\":\"next\",\"type\":\"tuple\"}],\"internalType\":\"structIGovParam.Param\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"activation\",\"type\":\"uint64\"}],\"name\":\"setParam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"votable\",\"type\":\"bool\"}],\"name\":\"setParamVotable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// GovParamBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const GovParamBinRuntime = `6080604052600436106100fe5760003560e01c80635d4f71d41161009557806399cb2e111161006457806399cb2e11146102a0578063a170052e146102b5578063c4d66de8146102d8578063ede7d2b4146102f8578063f2fde38b1461031857600080fd5b80635d4f71d414610216578063715018a6146102435780638841067a146102585780638da5cb5b1461027857600080fd5b80633659cfe6116100d15780633659cfe6146101a057806342a3f759146101c05780634f1ef286146101e057806352d1902d146101f357600080fd5b80630a6281fc146101035780630df2cc041461012f578063145f822b146101515780631ba1220d1461017e575b600080fd5b34801561010f57600080fd5b50610118610338565b604051610126929190611ca6565b60405180910390f35b34801561013b57600080fd5b50610144610406565b6040516101269190611d18565b34801561015d57600080fd5b5061017161016c366004611dd6565b610543565b6040516101269190611e12565b34801561018a57600080fd5b5061019e610199366004611e6d565b610710565b005b3480156101ac57600080fd5b5061019e6101bb366004611ec5565b6108d6565b3480156101cc57600080fd5b5061019e6101db366004611ee0565b6109b5565b61019e6101ee366004611f36565b610aa0565b3480156101ff57600080fd5b50610208610b70565b604051908152602001610126565b34801561022257600080fd5b50610236610231366004611dd6565b610c23565b6040516101269190611f97565b34801561024f57600080fd5b5061019e610c56565b34801561026457600080fd5b5061019e610273366004611faa565b610c8c565b34801561028457600080fd5b506097546040516001600160a01b039091168152602001610126565b3480156102ac57600080fd5b50610144610ea5565b3480156102c157600080fd5b506102ca610f7e565b604051610126929190612033565b3480156102e457600080fd5b5061019e6102f3366004611ec5565b611040565b34801561030457600080fd5b5061019e610313366004612096565b6110c6565b34801561032457600080fd5b5061019e610333366004611ec5565b6112f5565b6060806000610345610ea5565b9050600081516001600160401b0381111561036257610362611d2b565b60405190808252806020026020018201604052801561039b57816020015b6103886119a2565b8152602001906001900390816103805790505b50905060005b82518110156103fc576103cc8382815181106103bf576103bf612101565b6020026020010151610543565b8282815181106103de576103de612101565b602002602001018190525080806103f49061212d565b9150506103a1565b5090939092509050565b60606000610412610ea5565b90506000805b82518110156104685761044383828151811061043657610436612101565b602002602001015161138d565b1561045657816104528161212d565b9250505b806104608161212d565b915050610418565b506000816001600160401b0381111561048357610483611d2b565b6040519080825280602002602001820160405280156104b657816020015b60608152602001906001900390816104a15790505b5090506000915060005b835181101561053b576104de84828151811061043657610436612101565b15610529578381815181106104f5576104f5612101565b602002602001015182848151811061050f5761050f612101565b602002602001018190525082806105259061212d565b9350505b806105338161212d565b9150506104c0565b509392505050565b61054b6119a2565b60ca8260405161055b9190612146565b908152604080516020928190038301812060808201835280546001600160401b0381168352600160401b900460ff1615159382019390935281518083018352600184018054929493850192829082906105b390612162565b80601f01602080910402602001604051908101604052809291908181526020018280546105df90612162565b801561062c5780601f106106015761010080835404028352916020019161062c565b820191906000526020600020905b81548152906001019060200180831161060f57829003601f168201915b505050505081526020016001820160009054906101000a900460ff16151515158152505081526020016003820160405180604001604052908160008201805461067490612162565b80601f01602080910402602001604051908101604052809291908181526020018280546106a090612162565b80156106ed5780601f106106c2576101008083540402835291602001916106ed565b820191906000526020600020905b8154815290600101906020018083116106d057829003601f168201915b50505091835250506001919091015460ff16151560209091015290525092915050565b6097546001600160a01b031633146107435760405162461bcd60e51b815260040161073a9061219c565b60405180910390fd5b806107605760405162461bcd60e51b815260040161073a906121d1565b61079f82828080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061138d92505050565b6107bb5760405162461bcd60e51b815260040161073a90612208565b6107fa82828080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506113a292505050565b600060ca838360405161080e92919061224a565b90815260408051918290036020908101832090830191829052600092839052600301925061083e918391906119f9565b506001808201805460ff1916905561085790439061225a565b60ca848460405161086992919061224a565b90815260405190819003602001812080546001600160401b039390931667ffffffffffffffff19909316929092179091557fb888810dc5f09b858ca878b61b517f1a3925eb93a2051546d50021db9170fe57906108c9908590859061229b565b60405180910390a1505050565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016300361091e5760405162461bcd60e51b815260040161073a906122af565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610967600080516020612493833981519152546001600160a01b031690565b6001600160a01b03161461098d5760405162461bcd60e51b815260040161073a906122fb565b6109968161143b565b604080516000808252602082019092526109b291839190611465565b50565b6097546001600160a01b031633146109df5760405162461bcd60e51b815260040161073a9061219c565b6000825111610a005760405162461bcd60e51b815260040161073a906121d1565b610a098261138d565b610a255760405162461bcd60e51b815260040161073a90612208565b8060ca83604051610a369190612146565b9081526040519081900360200181208054921515600160401b0268ff000000000000000019909316929092179091557f064a476da825aa99a31c682cc490b34f462457006480bda5cb27151a201213f590610a949084908490612347565b60405180910390a15050565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163003610ae85760405162461bcd60e51b815260040161073a906122af565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610b31600080516020612493833981519152546001600160a01b031690565b6001600160a01b031614610b575760405162461bcd60e51b815260040161073a906122fb565b610b608261143b565b610b6c82826001611465565b5050565b6000306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610c105760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c0000000000000000606482015260840161073a565b5060008051602061249383398151915290565b6060610c2e8261138d565b610c4657505060408051602081019091526000815290565b610c4f826115d5565b5192915050565b6097546001600160a01b03163314610c805760405162461bcd60e51b815260040161073a9061219c565b610c8a6000611629565b565b6097546001600160a01b03163314610cb65760405162461bcd60e51b815260040161073a9061219c565b83610cd35760405162461bcd60e51b815260040161073a906121d1565b610d1285858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061138d92505050565b610d2e5760405162461bcd60e51b815260040161073a90612208565b43816001600160401b031611610d975760405162461bcd60e51b815260206004820152602860248201527f476f76506172616d3a2061637469766174696f6e206d75737420626520696e20604482015267612066757475726560c01b606482015260840161073a565b610dd685858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506113a292505050565b600060ca8686604051610dea92919061224a565b9081526040519081900360200190206003019050610e09818585611a7d565b506001818101805460ff19169091179055604051829060ca90610e2f908990899061224a565b90815260405190819003602001812080546001600160401b039390931667ffffffffffffffff19909316929092179091557ff8d4c4710beeb9d48e0bc6888dfcaf21467794f5da39c3bfb2874a2ac3a73c5690610e95908890889088908890889061236b565b60405180910390a1505050505050565b606060c9805480602002602001604051908101604052809291908181526020016000905b82821015610f75578382906000526020600020018054610ee890612162565b80601f0160208091040260200160405190810160405280929190818152602001828054610f1490612162565b8015610f615780601f10610f3657610100808354040283529160200191610f61565b820191906000526020600020905b815481529060010190602001808311610f4457829003601f168201915b505050505081526020019060010190610ec9565b50505050905090565b6060806000610f8b610406565b9050600081516001600160401b03811115610fa857610fa8611d2b565b604051908082528060200260200182016040528015610fdb57816020015b6060815260200190600190039081610fc65790505b50905060005b82518110156103fc5761100c838281518110610fff57610fff612101565b60200260200101516115d5565b6000015182828151811061102257611022612101565b602002602001018190525080806110389061212d565b915050610fe1565b600061104c600161167b565b90508015611064576000805461ff0019166101001790555b61106c611708565b6001600160a01b038216156110845761108482611629565b8015610b6c576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602001610a94565b6097546001600160a01b031633146110f05760405162461bcd60e51b815260040161073a9061219c565b8261110d5760405162461bcd60e51b815260040161073a906121d1565b61114c84848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061138d92505050565b156111a45760405162461bcd60e51b815260206004820152602260248201527f476f76506172616d3a20706172616d6574657220616c72656164792065786973604482015261747360f01b606482015260840161073a565b6111e384848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506113a292505050565b600060ca85856040516111f792919061224a565b9081526040519081900360200190206003019050611216818484611a7d565b506001818101805460ff19168217905561123190439061225a565b60ca868660405161124392919061224a565b90815260405190819003602001902080546001600160401b039290921667ffffffffffffffff1990921691909117905560c980546001810182556000919091526112b0907f66be4f155c5ef2ebd3772b228f2f00681e4ed5826cdb3b1943cc11ad15ad1d28018686611a7d565b507f83878da4b3d5b833aa14a193f9b57d482ab60c1e18ef91dce48f4584f85b8f72858585856040516112e694939291906123ae565b60405180910390a15050505050565b6097546001600160a01b0316331461131f5760405162461bcd60e51b815260040161073a9061219c565b6001600160a01b0381166113845760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161073a565b6109b281611629565b6000611398826115d5565b6020015192915050565b6113ab81610543565b516001600160401b031643106109b25760ca816040516113cb9190612146565b908152602001604051809103902060030160ca826040516113ec9190612146565b9081526020016040518091039020600101600082018160000190805461141190612162565b61141c929190611af1565b506001918201549101805460ff191660ff909216151591909117905550565b6097546001600160a01b031633146109b25760405162461bcd60e51b815260040161073a9061219c565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff161561149d5761149883611737565b505050565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156114f7575060408051601f3d908101601f191682019092526114f4918101906123e0565b60015b61155a5760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b606482015260840161073a565b60008051602061249383398151915281146115c95760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b606482015260840161073a565b506114988383836117d3565b6040805180820190915260608152600060208201526115f382610543565b516001600160401b031643106116165761160c82610543565b6060015192915050565b61161f82610543565b6040015192915050565b609780546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60008054610100900460ff16156116c2578160ff16600114801561169e5750303b155b6116ba5760405162461bcd60e51b815260040161073a906123f9565b506000919050565b60005460ff8084169116106116e95760405162461bcd60e51b815260040161073a906123f9565b506000805460ff191660ff92909216919091179055600190565b919050565b600054610100900460ff1661172f5760405162461bcd60e51b815260040161073a90612447565b610c8a6117fe565b6001600160a01b0381163b6117a45760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b606482015260840161073a565b60008051602061249383398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6117dc8361182e565b6000825111806117e95750805b15611498576117f8838361186e565b50505050565b600054610100900460ff166118255760405162461bcd60e51b815260040161073a90612447565b610c8a33611629565b61183781611737565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b60606001600160a01b0383163b6118d65760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b606482015260840161073a565b600080846001600160a01b0316846040516118f19190612146565b600060405180830381855af49150503d806000811461192c576040519150601f19603f3d011682016040523d82523d6000602084013e611931565b606091505b509150915061195982826040518060600160405280602781526020016124b360279139611962565b95945050505050565b6060831561197157508161199b565b8251156119815782518084602001fd5b8160405162461bcd60e51b815260040161073a9190611f97565b9392505050565b6040805160808101825260008082526020808301829052835180850185526060815290810191909152909182019081526020016119f46040518060400160405280606081526020016000151581525090565b905290565b828054611a0590612162565b90600052602060002090601f016020900481019282611a275760008555611a6d565b82601f10611a4057805160ff1916838001178555611a6d565b82800160010185558215611a6d579182015b82811115611a6d578251825591602001919060010190611a52565b50611a79929150611b6c565b5090565b828054611a8990612162565b90600052602060002090601f016020900481019282611aab5760008555611a6d565b82601f10611ac45782800160ff19823516178555611a6d565b82800160010185558215611a6d579182015b82811115611a6d578235825591602001919060010190611ad6565b828054611afd90612162565b90600052602060002090601f016020900481019282611b1f5760008555611a6d565b82601f10611b305780548555611a6d565b82800160010185558215611a6d57600052602060002091601f016020900482015b82811115611a6d578254825591600101919060010190611b51565b5b80821115611a795760008155600101611b6d565b60005b83811015611b9c578181015183820152602001611b84565b838111156117f85750506000910152565b60008151808452611bc5816020860160208601611b81565b601f01601f19169290920160200192915050565b600081518084526020808501808196508360051b8101915082860160005b85811015611c21578284038952611c0f848351611bad565b98850198935090840190600101611bf7565b5091979650505050505050565b6000815160408452611c436040850182611bad565b6020938401511515949093019390935250919050565b6001600160401b0381511682526020810151151560208301526000604082015160806040850152611c8d6080850182611c2e565b9050606083015184820360608601526119598282611c2e565b604081526000611cb96040830185611bd9565b6020838203818501528185518084528284019150828160051b85010183880160005b83811015611d0957601f19878403018552611cf7838351611c59565b94860194925090850190600101611cdb565b50909998505050505050505050565b60208152600061199b6020830184611bd9565b634e487b7160e01b600052604160045260246000fd5b60006001600160401b0380841115611d5b57611d5b611d2b565b604051601f8501601f19908116603f01168101908282118183101715611d8357611d83611d2b565b81604052809350858152868686011115611d9c57600080fd5b858560208301376000602087830101525050509392505050565b600082601f830112611dc757600080fd5b61199b83833560208501611d41565b600060208284031215611de857600080fd5b81356001600160401b03811115611dfe57600080fd5b611e0a84828501611db6565b949350505050565b60208152600061199b6020830184611c59565b60008083601f840112611e3757600080fd5b5081356001600160401b03811115611e4e57600080fd5b602083019150836020828501011115611e6657600080fd5b9250929050565b60008060208385031215611e8057600080fd5b82356001600160401b03811115611e9657600080fd5b611ea285828601611e25565b90969095509350505050565b80356001600160a01b038116811461170357600080fd5b600060208284031215611ed757600080fd5b61199b82611eae565b60008060408385031215611ef357600080fd5b82356001600160401b03811115611f0957600080fd5b611f1585828601611db6565b92505060208301358015158114611f2b57600080fd5b809150509250929050565b60008060408385031215611f4957600080fd5b611f5283611eae565b915060208301356001600160401b03811115611f6d57600080fd5b8301601f81018513611f7e57600080fd5b611f8d85823560208401611d41565b9150509250929050565b60208152600061199b6020830184611bad565b600080600080600060608688031215611fc257600080fd5b85356001600160401b0380821115611fd957600080fd5b611fe589838a01611e25565b90975095506020880135915080821115611ffe57600080fd5b61200a89838a01611e25565b909550935060408801359150808216821461202457600080fd5b50809150509295509295909350565b6040815260006120466040830185611bd9565b6020838203818501528185518084528284019150828160051b85010183880160005b83811015611d0957601f19878403018552612084838351611bad565b94860194925090850190600101612068565b600080600080604085870312156120ac57600080fd5b84356001600160401b03808211156120c357600080fd5b6120cf88838901611e25565b909650945060208701359150808211156120e857600080fd5b506120f587828801611e25565b95989497509550505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60006001820161213f5761213f612117565b5060010190565b60008251612158818460208701611b81565b9190910192915050565b600181811c9082168061217657607f821691505b60208210810361219657634e487b7160e01b600052602260045260246000fd5b50919050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6020808252601e908201527f476f76506172616d3a206e616d652063616e6e6f7420626520656d7074790000604082015260600190565b60208082526022908201527f476f76506172616d3a20706172616d6574657220646f6573206e6f74206578696040820152611cdd60f21b606082015260800190565b8183823760009101908152919050565b6000821982111561226d5761226d612117565b500190565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b602081526000611e0a602083018486612272565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b60408152600061235a6040830185611bad565b905082151560208301529392505050565b60608152600061237f606083018789612272565b8281036020840152612392818688612272565b9150506001600160401b03831660408301529695505050505050565b6040815260006123c2604083018688612272565b82810360208401526123d5818587612272565b979650505050505050565b6000602082840312156123f257600080fd5b5051919050565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b60608201526080019056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220f9458b9e41d9ccc92ee60ab0d42e458e6c58bcbe84ac4bc8bd70003558acf3ec64736f6c634300080e0033`

// GovParamFuncSigs maps the 4-byte function signature to its string representation.
var GovParamFuncSigs = map[string]string{
	"ede7d2b4": "addParam(string,bytes)",
	"1ba1220d": "deleteParam(string)",
	"a170052e": "getAllParams()",
	"0a6281fc": "getAllStructParams()",
	"5d4f71d4": "getParam(string)",
	"c4d66de8": "initialize(address)",
	"8da5cb5b": "owner()",
	"0df2cc04": "paramExistingNames()",
	"99cb2e11": "paramNames()",
	"145f822b": "params(string)",
	"52d1902d": "proxiableUUID()",
	"715018a6": "renounceOwnership()",
	"8841067a": "setParam(string,bytes,uint64)",
	"42a3f759": "setParamVotable(string,bool)",
	"f2fde38b": "transferOwnership(address)",
	"3659cfe6": "upgradeTo(address)",
	"4f1ef286": "upgradeToAndCall(address,bytes)",
}

// GovParamBin is the compiled bytecode used for deploying new contracts.
var GovParamBin = "0x60a06040523060805234801561001457600080fd5b5060805161250f61004c600039600081816108e00152818161092001528181610aaa01528181610aea0152610b7d015261250f6000f3fe6080604052600436106100fe5760003560e01c80635d4f71d41161009557806399cb2e111161006457806399cb2e11146102a0578063a170052e146102b5578063c4d66de8146102d8578063ede7d2b4146102f8578063f2fde38b1461031857600080fd5b80635d4f71d414610216578063715018a6146102435780638841067a146102585780638da5cb5b1461027857600080fd5b80633659cfe6116100d15780633659cfe6146101a057806342a3f759146101c05780634f1ef286146101e057806352d1902d146101f357600080fd5b80630a6281fc146101035780630df2cc041461012f578063145f822b146101515780631ba1220d1461017e575b600080fd5b34801561010f57600080fd5b50610118610338565b604051610126929190611ca6565b60405180910390f35b34801561013b57600080fd5b50610144610406565b6040516101269190611d18565b34801561015d57600080fd5b5061017161016c366004611dd6565b610543565b6040516101269190611e12565b34801561018a57600080fd5b5061019e610199366004611e6d565b610710565b005b3480156101ac57600080fd5b5061019e6101bb366004611ec5565b6108d6565b3480156101cc57600080fd5b5061019e6101db366004611ee0565b6109b5565b61019e6101ee366004611f36565b610aa0565b3480156101ff57600080fd5b50610208610b70565b604051908152602001610126565b34801561022257600080fd5b50610236610231366004611dd6565b610c23565b6040516101269190611f97565b34801561024f57600080fd5b5061019e610c56565b34801561026457600080fd5b5061019e610273366004611faa565b610c8c565b34801561028457600080fd5b506097546040516001600160a01b039091168152602001610126565b3480156102ac57600080fd5b50610144610ea5565b3480156102c157600080fd5b506102ca610f7e565b604051610126929190612033565b3480156102e457600080fd5b5061019e6102f3366004611ec5565b611040565b34801561030457600080fd5b5061019e610313366004612096565b6110c6565b34801561032457600080fd5b5061019e610333366004611ec5565b6112f5565b6060806000610345610ea5565b9050600081516001600160401b0381111561036257610362611d2b565b60405190808252806020026020018201604052801561039b57816020015b6103886119a2565b8152602001906001900390816103805790505b50905060005b82518110156103fc576103cc8382815181106103bf576103bf612101565b6020026020010151610543565b8282815181106103de576103de612101565b602002602001018190525080806103f49061212d565b9150506103a1565b5090939092509050565b60606000610412610ea5565b90506000805b82518110156104685761044383828151811061043657610436612101565b602002602001015161138d565b1561045657816104528161212d565b9250505b806104608161212d565b915050610418565b506000816001600160401b0381111561048357610483611d2b565b6040519080825280602002602001820160405280156104b657816020015b60608152602001906001900390816104a15790505b5090506000915060005b835181101561053b576104de84828151811061043657610436612101565b15610529578381815181106104f5576104f5612101565b602002602001015182848151811061050f5761050f612101565b602002602001018190525082806105259061212d565b9350505b806105338161212d565b9150506104c0565b509392505050565b61054b6119a2565b60ca8260405161055b9190612146565b908152604080516020928190038301812060808201835280546001600160401b0381168352600160401b900460ff1615159382019390935281518083018352600184018054929493850192829082906105b390612162565b80601f01602080910402602001604051908101604052809291908181526020018280546105df90612162565b801561062c5780601f106106015761010080835404028352916020019161062c565b820191906000526020600020905b81548152906001019060200180831161060f57829003601f168201915b505050505081526020016001820160009054906101000a900460ff16151515158152505081526020016003820160405180604001604052908160008201805461067490612162565b80601f01602080910402602001604051908101604052809291908181526020018280546106a090612162565b80156106ed5780601f106106c2576101008083540402835291602001916106ed565b820191906000526020600020905b8154815290600101906020018083116106d057829003601f168201915b50505091835250506001919091015460ff16151560209091015290525092915050565b6097546001600160a01b031633146107435760405162461bcd60e51b815260040161073a9061219c565b60405180910390fd5b806107605760405162461bcd60e51b815260040161073a906121d1565b61079f82828080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061138d92505050565b6107bb5760405162461bcd60e51b815260040161073a90612208565b6107fa82828080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506113a292505050565b600060ca838360405161080e92919061224a565b90815260408051918290036020908101832090830191829052600092839052600301925061083e918391906119f9565b506001808201805460ff1916905561085790439061225a565b60ca848460405161086992919061224a565b90815260405190819003602001812080546001600160401b039390931667ffffffffffffffff19909316929092179091557fb888810dc5f09b858ca878b61b517f1a3925eb93a2051546d50021db9170fe57906108c9908590859061229b565b60405180910390a1505050565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016300361091e5760405162461bcd60e51b815260040161073a906122af565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610967600080516020612493833981519152546001600160a01b031690565b6001600160a01b03161461098d5760405162461bcd60e51b815260040161073a906122fb565b6109968161143b565b604080516000808252602082019092526109b291839190611465565b50565b6097546001600160a01b031633146109df5760405162461bcd60e51b815260040161073a9061219c565b6000825111610a005760405162461bcd60e51b815260040161073a906121d1565b610a098261138d565b610a255760405162461bcd60e51b815260040161073a90612208565b8060ca83604051610a369190612146565b9081526040519081900360200181208054921515600160401b0268ff000000000000000019909316929092179091557f064a476da825aa99a31c682cc490b34f462457006480bda5cb27151a201213f590610a949084908490612347565b60405180910390a15050565b6001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000163003610ae85760405162461bcd60e51b815260040161073a906122af565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610b31600080516020612493833981519152546001600160a01b031690565b6001600160a01b031614610b575760405162461bcd60e51b815260040161073a906122fb565b610b608261143b565b610b6c82826001611465565b5050565b6000306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610c105760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c0000000000000000606482015260840161073a565b5060008051602061249383398151915290565b6060610c2e8261138d565b610c4657505060408051602081019091526000815290565b610c4f826115d5565b5192915050565b6097546001600160a01b03163314610c805760405162461bcd60e51b815260040161073a9061219c565b610c8a6000611629565b565b6097546001600160a01b03163314610cb65760405162461bcd60e51b815260040161073a9061219c565b83610cd35760405162461bcd60e51b815260040161073a906121d1565b610d1285858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061138d92505050565b610d2e5760405162461bcd60e51b815260040161073a90612208565b43816001600160401b031611610d975760405162461bcd60e51b815260206004820152602860248201527f476f76506172616d3a2061637469766174696f6e206d75737420626520696e20604482015267612066757475726560c01b606482015260840161073a565b610dd685858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506113a292505050565b600060ca8686604051610dea92919061224a565b9081526040519081900360200190206003019050610e09818585611a7d565b506001818101805460ff19169091179055604051829060ca90610e2f908990899061224a565b90815260405190819003602001812080546001600160401b039390931667ffffffffffffffff19909316929092179091557ff8d4c4710beeb9d48e0bc6888dfcaf21467794f5da39c3bfb2874a2ac3a73c5690610e95908890889088908890889061236b565b60405180910390a1505050505050565b606060c9805480602002602001604051908101604052809291908181526020016000905b82821015610f75578382906000526020600020018054610ee890612162565b80601f0160208091040260200160405190810160405280929190818152602001828054610f1490612162565b8015610f615780601f10610f3657610100808354040283529160200191610f61565b820191906000526020600020905b815481529060010190602001808311610f4457829003601f168201915b505050505081526020019060010190610ec9565b50505050905090565b6060806000610f8b610406565b9050600081516001600160401b03811115610fa857610fa8611d2b565b604051908082528060200260200182016040528015610fdb57816020015b6060815260200190600190039081610fc65790505b50905060005b82518110156103fc5761100c838281518110610fff57610fff612101565b60200260200101516115d5565b6000015182828151811061102257611022612101565b602002602001018190525080806110389061212d565b915050610fe1565b600061104c600161167b565b90508015611064576000805461ff0019166101001790555b61106c611708565b6001600160a01b038216156110845761108482611629565b8015610b6c576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602001610a94565b6097546001600160a01b031633146110f05760405162461bcd60e51b815260040161073a9061219c565b8261110d5760405162461bcd60e51b815260040161073a906121d1565b61114c84848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061138d92505050565b156111a45760405162461bcd60e51b815260206004820152602260248201527f476f76506172616d3a20706172616d6574657220616c72656164792065786973604482015261747360f01b606482015260840161073a565b6111e384848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506113a292505050565b600060ca85856040516111f792919061224a565b9081526040519081900360200190206003019050611216818484611a7d565b506001818101805460ff19168217905561123190439061225a565b60ca868660405161124392919061224a565b90815260405190819003602001902080546001600160401b039290921667ffffffffffffffff1990921691909117905560c980546001810182556000919091526112b0907f66be4f155c5ef2ebd3772b228f2f00681e4ed5826cdb3b1943cc11ad15ad1d28018686611a7d565b507f83878da4b3d5b833aa14a193f9b57d482ab60c1e18ef91dce48f4584f85b8f72858585856040516112e694939291906123ae565b60405180910390a15050505050565b6097546001600160a01b0316331461131f5760405162461bcd60e51b815260040161073a9061219c565b6001600160a01b0381166113845760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840161073a565b6109b281611629565b6000611398826115d5565b6020015192915050565b6113ab81610543565b516001600160401b031643106109b25760ca816040516113cb9190612146565b908152602001604051809103902060030160ca826040516113ec9190612146565b9081526020016040518091039020600101600082018160000190805461141190612162565b61141c929190611af1565b506001918201549101805460ff191660ff909216151591909117905550565b6097546001600160a01b031633146109b25760405162461bcd60e51b815260040161073a9061219c565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff161561149d5761149883611737565b505050565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156114f7575060408051601f3d908101601f191682019092526114f4918101906123e0565b60015b61155a5760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b606482015260840161073a565b60008051602061249383398151915281146115c95760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b606482015260840161073a565b506114988383836117d3565b6040805180820190915260608152600060208201526115f382610543565b516001600160401b031643106116165761160c82610543565b6060015192915050565b61161f82610543565b6040015192915050565b609780546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60008054610100900460ff16156116c2578160ff16600114801561169e5750303b155b6116ba5760405162461bcd60e51b815260040161073a906123f9565b506000919050565b60005460ff8084169116106116e95760405162461bcd60e51b815260040161073a906123f9565b506000805460ff191660ff92909216919091179055600190565b919050565b600054610100900460ff1661172f5760405162461bcd60e51b815260040161073a90612447565b610c8a6117fe565b6001600160a01b0381163b6117a45760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b606482015260840161073a565b60008051602061249383398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6117dc8361182e565b6000825111806117e95750805b15611498576117f8838361186e565b50505050565b600054610100900460ff166118255760405162461bcd60e51b815260040161073a90612447565b610c8a33611629565b61183781611737565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b60606001600160a01b0383163b6118d65760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b606482015260840161073a565b600080846001600160a01b0316846040516118f19190612146565b600060405180830381855af49150503d806000811461192c576040519150601f19603f3d011682016040523d82523d6000602084013e611931565b606091505b509150915061195982826040518060600160405280602781526020016124b360279139611962565b95945050505050565b6060831561197157508161199b565b8251156119815782518084602001fd5b8160405162461bcd60e51b815260040161073a9190611f97565b9392505050565b6040805160808101825260008082526020808301829052835180850185526060815290810191909152909182019081526020016119f46040518060400160405280606081526020016000151581525090565b905290565b828054611a0590612162565b90600052602060002090601f016020900481019282611a275760008555611a6d565b82601f10611a4057805160ff1916838001178555611a6d565b82800160010185558215611a6d579182015b82811115611a6d578251825591602001919060010190611a52565b50611a79929150611b6c565b5090565b828054611a8990612162565b90600052602060002090601f016020900481019282611aab5760008555611a6d565b82601f10611ac45782800160ff19823516178555611a6d565b82800160010185558215611a6d579182015b82811115611a6d578235825591602001919060010190611ad6565b828054611afd90612162565b90600052602060002090601f016020900481019282611b1f5760008555611a6d565b82601f10611b305780548555611a6d565b82800160010185558215611a6d57600052602060002091601f016020900482015b82811115611a6d578254825591600101919060010190611b51565b5b80821115611a795760008155600101611b6d565b60005b83811015611b9c578181015183820152602001611b84565b838111156117f85750506000910152565b60008151808452611bc5816020860160208601611b81565b601f01601f19169290920160200192915050565b600081518084526020808501808196508360051b8101915082860160005b85811015611c21578284038952611c0f848351611bad565b98850198935090840190600101611bf7565b5091979650505050505050565b6000815160408452611c436040850182611bad565b6020938401511515949093019390935250919050565b6001600160401b0381511682526020810151151560208301526000604082015160806040850152611c8d6080850182611c2e565b9050606083015184820360608601526119598282611c2e565b604081526000611cb96040830185611bd9565b6020838203818501528185518084528284019150828160051b85010183880160005b83811015611d0957601f19878403018552611cf7838351611c59565b94860194925090850190600101611cdb565b50909998505050505050505050565b60208152600061199b6020830184611bd9565b634e487b7160e01b600052604160045260246000fd5b60006001600160401b0380841115611d5b57611d5b611d2b565b604051601f8501601f19908116603f01168101908282118183101715611d8357611d83611d2b565b81604052809350858152868686011115611d9c57600080fd5b858560208301376000602087830101525050509392505050565b600082601f830112611dc757600080fd5b61199b83833560208501611d41565b600060208284031215611de857600080fd5b81356001600160401b03811115611dfe57600080fd5b611e0a84828501611db6565b949350505050565b60208152600061199b6020830184611c59565b60008083601f840112611e3757600080fd5b5081356001600160401b03811115611e4e57600080fd5b602083019150836020828501011115611e6657600080fd5b9250929050565b60008060208385031215611e8057600080fd5b82356001600160401b03811115611e9657600080fd5b611ea285828601611e25565b90969095509350505050565b80356001600160a01b038116811461170357600080fd5b600060208284031215611ed757600080fd5b61199b82611eae565b60008060408385031215611ef357600080fd5b82356001600160401b03811115611f0957600080fd5b611f1585828601611db6565b92505060208301358015158114611f2b57600080fd5b809150509250929050565b60008060408385031215611f4957600080fd5b611f5283611eae565b915060208301356001600160401b03811115611f6d57600080fd5b8301601f81018513611f7e57600080fd5b611f8d85823560208401611d41565b9150509250929050565b60208152600061199b6020830184611bad565b600080600080600060608688031215611fc257600080fd5b85356001600160401b0380821115611fd957600080fd5b611fe589838a01611e25565b90975095506020880135915080821115611ffe57600080fd5b61200a89838a01611e25565b909550935060408801359150808216821461202457600080fd5b50809150509295509295909350565b6040815260006120466040830185611bd9565b6020838203818501528185518084528284019150828160051b85010183880160005b83811015611d0957601f19878403018552612084838351611bad565b94860194925090850190600101612068565b600080600080604085870312156120ac57600080fd5b84356001600160401b03808211156120c357600080fd5b6120cf88838901611e25565b909650945060208701359150808211156120e857600080fd5b506120f587828801611e25565b95989497509550505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60006001820161213f5761213f612117565b5060010190565b60008251612158818460208701611b81565b9190910192915050565b600181811c9082168061217657607f821691505b60208210810361219657634e487b7160e01b600052602260045260246000fd5b50919050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6020808252601e908201527f476f76506172616d3a206e616d652063616e6e6f7420626520656d7074790000604082015260600190565b60208082526022908201527f476f76506172616d3a20706172616d6574657220646f6573206e6f74206578696040820152611cdd60f21b606082015260800190565b8183823760009101908152919050565b6000821982111561226d5761226d612117565b500190565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b602081526000611e0a602083018486612272565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b60408152600061235a6040830185611bad565b905082151560208301529392505050565b60608152600061237f606083018789612272565b8281036020840152612392818688612272565b9150506001600160401b03831660408301529695505050505050565b6040815260006123c2604083018688612272565b82810360208401526123d5818587612272565b979650505050505050565b6000602082840312156123f257600080fd5b5051919050565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b60608201526080019056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220f9458b9e41d9ccc92ee60ab0d42e458e6c58bcbe84ac4bc8bd70003558acf3ec64736f6c634300080e0033"

// DeployGovParam deploys a new Klaytn contract, binding an instance of GovParam to it.
func DeployGovParam(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GovParam, error) {
	parsed, err := abi.JSON(strings.NewReader(GovParamABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GovParamBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GovParam{GovParamCaller: GovParamCaller{contract: contract}, GovParamTransactor: GovParamTransactor{contract: contract}, GovParamFilterer: GovParamFilterer{contract: contract}}, nil
}

// GovParam is an auto generated Go binding around a Klaytn contract.
type GovParam struct {
	GovParamCaller     // Read-only binding to the contract
	GovParamTransactor // Write-only binding to the contract
	GovParamFilterer   // Log filterer for contract events
}

// GovParamCaller is an auto generated read-only Go binding around a Klaytn contract.
type GovParamCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovParamTransactor is an auto generated write-only Go binding around a Klaytn contract.
type GovParamTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovParamFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type GovParamFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovParamSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type GovParamSession struct {
	Contract     *GovParam         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovParamCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type GovParamCallerSession struct {
	Contract *GovParamCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// GovParamTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type GovParamTransactorSession struct {
	Contract     *GovParamTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// GovParamRaw is an auto generated low-level Go binding around a Klaytn contract.
type GovParamRaw struct {
	Contract *GovParam // Generic contract binding to access the raw methods on
}

// GovParamCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type GovParamCallerRaw struct {
	Contract *GovParamCaller // Generic read-only contract binding to access the raw methods on
}

// GovParamTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type GovParamTransactorRaw struct {
	Contract *GovParamTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGovParam creates a new instance of GovParam, bound to a specific deployed contract.
func NewGovParam(address common.Address, backend bind.ContractBackend) (*GovParam, error) {
	contract, err := bindGovParam(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GovParam{GovParamCaller: GovParamCaller{contract: contract}, GovParamTransactor: GovParamTransactor{contract: contract}, GovParamFilterer: GovParamFilterer{contract: contract}}, nil
}

// NewGovParamCaller creates a new read-only instance of GovParam, bound to a specific deployed contract.
func NewGovParamCaller(address common.Address, caller bind.ContractCaller) (*GovParamCaller, error) {
	contract, err := bindGovParam(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovParamCaller{contract: contract}, nil
}

// NewGovParamTransactor creates a new write-only instance of GovParam, bound to a specific deployed contract.
func NewGovParamTransactor(address common.Address, transactor bind.ContractTransactor) (*GovParamTransactor, error) {
	contract, err := bindGovParam(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovParamTransactor{contract: contract}, nil
}

// NewGovParamFilterer creates a new log filterer instance of GovParam, bound to a specific deployed contract.
func NewGovParamFilterer(address common.Address, filterer bind.ContractFilterer) (*GovParamFilterer, error) {
	contract, err := bindGovParam(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovParamFilterer{contract: contract}, nil
}

// bindGovParam binds a generic wrapper to an already deployed contract.
func bindGovParam(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GovParamABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GovParam *GovParamRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GovParam.Contract.GovParamCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GovParam *GovParamRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovParam.Contract.GovParamTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GovParam *GovParamRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GovParam.Contract.GovParamTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GovParam *GovParamCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GovParam.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GovParam *GovParamTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovParam.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GovParam *GovParamTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GovParam.Contract.contract.Transact(opts, method, params...)
}

// GetAllParams is a free data retrieval call binding the contract method 0xa170052e.
//
// Solidity: function getAllParams() view returns(string[], bytes[])
func (_GovParam *GovParamCaller) GetAllParams(opts *bind.CallOpts) ([]string, [][]byte, error) {
	var (
		ret0 = new([]string)
		ret1 = new([][]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _GovParam.contract.Call(opts, out, "getAllParams")
	return *ret0, *ret1, err
}

// GetAllParams is a free data retrieval call binding the contract method 0xa170052e.
//
// Solidity: function getAllParams() view returns(string[], bytes[])
func (_GovParam *GovParamSession) GetAllParams() ([]string, [][]byte, error) {
	return _GovParam.Contract.GetAllParams(&_GovParam.CallOpts)
}

// GetAllParams is a free data retrieval call binding the contract method 0xa170052e.
//
// Solidity: function getAllParams() view returns(string[], bytes[])
func (_GovParam *GovParamCallerSession) GetAllParams() ([]string, [][]byte, error) {
	return _GovParam.Contract.GetAllParams(&_GovParam.CallOpts)
}

// GetAllStructParams is a free data retrieval call binding the contract method 0x0a6281fc.
//
// Solidity: function getAllStructParams() view returns(string[], (uint64,bool,(bytes,bool),(bytes,bool))[])
func (_GovParam *GovParamCaller) GetAllStructParams(opts *bind.CallOpts) ([]string, []IGovParamParam, error) {
	var (
		ret0 = new([]string)
		ret1 = new([]IGovParamParam)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _GovParam.contract.Call(opts, out, "getAllStructParams")
	return *ret0, *ret1, err
}

// GetAllStructParams is a free data retrieval call binding the contract method 0x0a6281fc.
//
// Solidity: function getAllStructParams() view returns(string[], (uint64,bool,(bytes,bool),(bytes,bool))[])
func (_GovParam *GovParamSession) GetAllStructParams() ([]string, []IGovParamParam, error) {
	return _GovParam.Contract.GetAllStructParams(&_GovParam.CallOpts)
}

// GetAllStructParams is a free data retrieval call binding the contract method 0x0a6281fc.
//
// Solidity: function getAllStructParams() view returns(string[], (uint64,bool,(bytes,bool),(bytes,bool))[])
func (_GovParam *GovParamCallerSession) GetAllStructParams() ([]string, []IGovParamParam, error) {
	return _GovParam.Contract.GetAllStructParams(&_GovParam.CallOpts)
}

// GetParam is a free data retrieval call binding the contract method 0x5d4f71d4.
//
// Solidity: function getParam(string name) view returns(bytes)
func (_GovParam *GovParamCaller) GetParam(opts *bind.CallOpts, name string) ([]byte, error) {
	ret0 := new([]byte)
	out := ret0
	err := _GovParam.contract.Call(opts, out, "getParam", name)
	return *ret0, err
}

// GetParam is a free data retrieval call binding the contract method 0x5d4f71d4.
//
// Solidity: function getParam(string name) view returns(bytes)
func (_GovParam *GovParamSession) GetParam(name string) ([]byte, error) {
	return _GovParam.Contract.GetParam(&_GovParam.CallOpts, name)
}

// GetParam is a free data retrieval call binding the contract method 0x5d4f71d4.
//
// Solidity: function getParam(string name) view returns(bytes)
func (_GovParam *GovParamCallerSession) GetParam(name string) ([]byte, error) {
	return _GovParam.Contract.GetParam(&_GovParam.CallOpts, name)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GovParam *GovParamCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	ret0 := new(common.Address)
	out := ret0
	err := _GovParam.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GovParam *GovParamSession) Owner() (common.Address, error) {
	return _GovParam.Contract.Owner(&_GovParam.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GovParam *GovParamCallerSession) Owner() (common.Address, error) {
	return _GovParam.Contract.Owner(&_GovParam.CallOpts)
}

// ParamExistingNames is a free data retrieval call binding the contract method 0x0df2cc04.
//
// Solidity: function paramExistingNames() view returns(string[])
func (_GovParam *GovParamCaller) ParamExistingNames(opts *bind.CallOpts) ([]string, error) {
	ret0 := new([]string)
	out := ret0
	err := _GovParam.contract.Call(opts, out, "paramExistingNames")
	return *ret0, err
}

// ParamExistingNames is a free data retrieval call binding the contract method 0x0df2cc04.
//
// Solidity: function paramExistingNames() view returns(string[])
func (_GovParam *GovParamSession) ParamExistingNames() ([]string, error) {
	return _GovParam.Contract.ParamExistingNames(&_GovParam.CallOpts)
}

// ParamExistingNames is a free data retrieval call binding the contract method 0x0df2cc04.
//
// Solidity: function paramExistingNames() view returns(string[])
func (_GovParam *GovParamCallerSession) ParamExistingNames() ([]string, error) {
	return _GovParam.Contract.ParamExistingNames(&_GovParam.CallOpts)
}

// ParamNames is a free data retrieval call binding the contract method 0x99cb2e11.
//
// Solidity: function paramNames() view returns(string[])
func (_GovParam *GovParamCaller) ParamNames(opts *bind.CallOpts) ([]string, error) {
	ret0 := new([]string)
	out := ret0
	err := _GovParam.contract.Call(opts, out, "paramNames")
	return *ret0, err
}

// ParamNames is a free data retrieval call binding the contract method 0x99cb2e11.
//
// Solidity: function paramNames() view returns(string[])
func (_GovParam *GovParamSession) ParamNames() ([]string, error) {
	return _GovParam.Contract.ParamNames(&_GovParam.CallOpts)
}

// ParamNames is a free data retrieval call binding the contract method 0x99cb2e11.
//
// Solidity: function paramNames() view returns(string[])
func (_GovParam *GovParamCallerSession) ParamNames() ([]string, error) {
	return _GovParam.Contract.ParamNames(&_GovParam.CallOpts)
}

// Params is a free data retrieval call binding the contract method 0x145f822b.
//
// Solidity: function params(string name) view returns((uint64,bool,(bytes,bool),(bytes,bool)))
func (_GovParam *GovParamCaller) Params(opts *bind.CallOpts, name string) (IGovParamParam, error) {
	ret0 := new(IGovParamParam)
	out := ret0
	err := _GovParam.contract.Call(opts, out, "params", name)
	return *ret0, err
}

// Params is a free data retrieval call binding the contract method 0x145f822b.
//
// Solidity: function params(string name) view returns((uint64,bool,(bytes,bool),(bytes,bool)))
func (_GovParam *GovParamSession) Params(name string) (IGovParamParam, error) {
	return _GovParam.Contract.Params(&_GovParam.CallOpts, name)
}

// Params is a free data retrieval call binding the contract method 0x145f822b.
//
// Solidity: function params(string name) view returns((uint64,bool,(bytes,bool),(bytes,bool)))
func (_GovParam *GovParamCallerSession) Params(name string) (IGovParamParam, error) {
	return _GovParam.Contract.Params(&_GovParam.CallOpts, name)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GovParam *GovParamCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	ret0 := new([32]byte)
	out := ret0
	err := _GovParam.contract.Call(opts, out, "proxiableUUID")
	return *ret0, err
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GovParam *GovParamSession) ProxiableUUID() ([32]byte, error) {
	return _GovParam.Contract.ProxiableUUID(&_GovParam.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_GovParam *GovParamCallerSession) ProxiableUUID() ([32]byte, error) {
	return _GovParam.Contract.ProxiableUUID(&_GovParam.CallOpts)
}

// AddParam is a paid mutator transaction binding the contract method 0xede7d2b4.
//
// Solidity: function addParam(string name, bytes value) returns()
func (_GovParam *GovParamTransactor) AddParam(opts *bind.TransactOpts, name string, value []byte) (*types.Transaction, error) {
	return _GovParam.contract.Transact(opts, "addParam", name, value)
}

// AddParam is a paid mutator transaction binding the contract method 0xede7d2b4.
//
// Solidity: function addParam(string name, bytes value) returns()
func (_GovParam *GovParamSession) AddParam(name string, value []byte) (*types.Transaction, error) {
	return _GovParam.Contract.AddParam(&_GovParam.TransactOpts, name, value)
}

// AddParam is a paid mutator transaction binding the contract method 0xede7d2b4.
//
// Solidity: function addParam(string name, bytes value) returns()
func (_GovParam *GovParamTransactorSession) AddParam(name string, value []byte) (*types.Transaction, error) {
	return _GovParam.Contract.AddParam(&_GovParam.TransactOpts, name, value)
}

// DeleteParam is a paid mutator transaction binding the contract method 0x1ba1220d.
//
// Solidity: function deleteParam(string name) returns()
func (_GovParam *GovParamTransactor) DeleteParam(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _GovParam.contract.Transact(opts, "deleteParam", name)
}

// DeleteParam is a paid mutator transaction binding the contract method 0x1ba1220d.
//
// Solidity: function deleteParam(string name) returns()
func (_GovParam *GovParamSession) DeleteParam(name string) (*types.Transaction, error) {
	return _GovParam.Contract.DeleteParam(&_GovParam.TransactOpts, name)
}

// DeleteParam is a paid mutator transaction binding the contract method 0x1ba1220d.
//
// Solidity: function deleteParam(string name) returns()
func (_GovParam *GovParamTransactorSession) DeleteParam(name string) (*types.Transaction, error) {
	return _GovParam.Contract.DeleteParam(&_GovParam.TransactOpts, name)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner) returns()
func (_GovParam *GovParamTransactor) Initialize(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _GovParam.contract.Transact(opts, "initialize", owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner) returns()
func (_GovParam *GovParamSession) Initialize(owner common.Address) (*types.Transaction, error) {
	return _GovParam.Contract.Initialize(&_GovParam.TransactOpts, owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner) returns()
func (_GovParam *GovParamTransactorSession) Initialize(owner common.Address) (*types.Transaction, error) {
	return _GovParam.Contract.Initialize(&_GovParam.TransactOpts, owner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GovParam *GovParamTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovParam.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GovParam *GovParamSession) RenounceOwnership() (*types.Transaction, error) {
	return _GovParam.Contract.RenounceOwnership(&_GovParam.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GovParam *GovParamTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GovParam.Contract.RenounceOwnership(&_GovParam.TransactOpts)
}

// SetParam is a paid mutator transaction binding the contract method 0x8841067a.
//
// Solidity: function setParam(string name, bytes value, uint64 activation) returns()
func (_GovParam *GovParamTransactor) SetParam(opts *bind.TransactOpts, name string, value []byte, activation uint64) (*types.Transaction, error) {
	return _GovParam.contract.Transact(opts, "setParam", name, value, activation)
}

// SetParam is a paid mutator transaction binding the contract method 0x8841067a.
//
// Solidity: function setParam(string name, bytes value, uint64 activation) returns()
func (_GovParam *GovParamSession) SetParam(name string, value []byte, activation uint64) (*types.Transaction, error) {
	return _GovParam.Contract.SetParam(&_GovParam.TransactOpts, name, value, activation)
}

// SetParam is a paid mutator transaction binding the contract method 0x8841067a.
//
// Solidity: function setParam(string name, bytes value, uint64 activation) returns()
func (_GovParam *GovParamTransactorSession) SetParam(name string, value []byte, activation uint64) (*types.Transaction, error) {
	return _GovParam.Contract.SetParam(&_GovParam.TransactOpts, name, value, activation)
}

// SetParamVotable is a paid mutator transaction binding the contract method 0x42a3f759.
//
// Solidity: function setParamVotable(string name, bool votable) returns()
func (_GovParam *GovParamTransactor) SetParamVotable(opts *bind.TransactOpts, name string, votable bool) (*types.Transaction, error) {
	return _GovParam.contract.Transact(opts, "setParamVotable", name, votable)
}

// SetParamVotable is a paid mutator transaction binding the contract method 0x42a3f759.
//
// Solidity: function setParamVotable(string name, bool votable) returns()
func (_GovParam *GovParamSession) SetParamVotable(name string, votable bool) (*types.Transaction, error) {
	return _GovParam.Contract.SetParamVotable(&_GovParam.TransactOpts, name, votable)
}

// SetParamVotable is a paid mutator transaction binding the contract method 0x42a3f759.
//
// Solidity: function setParamVotable(string name, bool votable) returns()
func (_GovParam *GovParamTransactorSession) SetParamVotable(name string, votable bool) (*types.Transaction, error) {
	return _GovParam.Contract.SetParamVotable(&_GovParam.TransactOpts, name, votable)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GovParam *GovParamTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GovParam.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GovParam *GovParamSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GovParam.Contract.TransferOwnership(&_GovParam.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GovParam *GovParamTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GovParam.Contract.TransferOwnership(&_GovParam.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_GovParam *GovParamTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _GovParam.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_GovParam *GovParamSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _GovParam.Contract.UpgradeTo(&_GovParam.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_GovParam *GovParamTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _GovParam.Contract.UpgradeTo(&_GovParam.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GovParam *GovParamTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GovParam.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GovParam *GovParamSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GovParam.Contract.UpgradeToAndCall(&_GovParam.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_GovParam *GovParamTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _GovParam.Contract.UpgradeToAndCall(&_GovParam.TransactOpts, newImplementation, data)
}

// GovParamAddParamIterator is returned from FilterAddParam and is used to iterate over the raw logs and unpacked data for AddParam events raised by the GovParam contract.
type GovParamAddParamIterator struct {
	Event *GovParamAddParam // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovParamAddParamIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovParamAddParam)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovParamAddParam)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovParamAddParamIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovParamAddParamIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovParamAddParam represents a AddParam event raised by the GovParam contract.
type GovParamAddParam struct {
	Arg0 string
	Arg1 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddParam is a free log retrieval operation binding the contract event 0x83878da4b3d5b833aa14a193f9b57d482ab60c1e18ef91dce48f4584f85b8f72.
//
// Solidity: event AddParam(string arg0, bytes arg1)
func (_GovParam *GovParamFilterer) FilterAddParam(opts *bind.FilterOpts) (*GovParamAddParamIterator, error) {
	logs, sub, err := _GovParam.contract.FilterLogs(opts, "AddParam")
	if err != nil {
		return nil, err
	}
	return &GovParamAddParamIterator{contract: _GovParam.contract, event: "AddParam", logs: logs, sub: sub}, nil
}

// WatchAddParam is a free log subscription operation binding the contract event 0x83878da4b3d5b833aa14a193f9b57d482ab60c1e18ef91dce48f4584f85b8f72.
//
// Solidity: event AddParam(string arg0, bytes arg1)
func (_GovParam *GovParamFilterer) WatchAddParam(opts *bind.WatchOpts, sink chan<- *GovParamAddParam) (event.Subscription, error) {
	logs, sub, err := _GovParam.contract.WatchLogs(opts, "AddParam")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovParamAddParam)
				if err := _GovParam.contract.UnpackLog(event, "AddParam", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddParam is a log parse operation binding the contract event 0x83878da4b3d5b833aa14a193f9b57d482ab60c1e18ef91dce48f4584f85b8f72.
//
// Solidity: event AddParam(string arg0, bytes arg1)
func (_GovParam *GovParamFilterer) ParseAddParam(log types.Log) (*GovParamAddParam, error) {
	event := new(GovParamAddParam)
	if err := _GovParam.contract.UnpackLog(event, "AddParam", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovParamAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the GovParam contract.
type GovParamAdminChangedIterator struct {
	Event *GovParamAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovParamAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovParamAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovParamAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovParamAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovParamAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovParamAdminChanged represents a AdminChanged event raised by the GovParam contract.
type GovParamAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_GovParam *GovParamFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*GovParamAdminChangedIterator, error) {
	logs, sub, err := _GovParam.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &GovParamAdminChangedIterator{contract: _GovParam.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_GovParam *GovParamFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *GovParamAdminChanged) (event.Subscription, error) {
	logs, sub, err := _GovParam.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovParamAdminChanged)
				if err := _GovParam.contract.UnpackLog(event, "AdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_GovParam *GovParamFilterer) ParseAdminChanged(log types.Log) (*GovParamAdminChanged, error) {
	event := new(GovParamAdminChanged)
	if err := _GovParam.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovParamBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the GovParam contract.
type GovParamBeaconUpgradedIterator struct {
	Event *GovParamBeaconUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovParamBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovParamBeaconUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovParamBeaconUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovParamBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovParamBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovParamBeaconUpgraded represents a BeaconUpgraded event raised by the GovParam contract.
type GovParamBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_GovParam *GovParamFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*GovParamBeaconUpgradedIterator, error) {
	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _GovParam.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &GovParamBeaconUpgradedIterator{contract: _GovParam.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_GovParam *GovParamFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *GovParamBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {
	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _GovParam.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovParamBeaconUpgraded)
				if err := _GovParam.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_GovParam *GovParamFilterer) ParseBeaconUpgraded(log types.Log) (*GovParamBeaconUpgraded, error) {
	event := new(GovParamBeaconUpgraded)
	if err := _GovParam.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovParamDeleteParamIterator is returned from FilterDeleteParam and is used to iterate over the raw logs and unpacked data for DeleteParam events raised by the GovParam contract.
type GovParamDeleteParamIterator struct {
	Event *GovParamDeleteParam // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovParamDeleteParamIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovParamDeleteParam)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovParamDeleteParam)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovParamDeleteParamIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovParamDeleteParamIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovParamDeleteParam represents a DeleteParam event raised by the GovParam contract.
type GovParamDeleteParam struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDeleteParam is a free log retrieval operation binding the contract event 0xb888810dc5f09b858ca878b61b517f1a3925eb93a2051546d50021db9170fe57.
//
// Solidity: event DeleteParam(string arg0)
func (_GovParam *GovParamFilterer) FilterDeleteParam(opts *bind.FilterOpts) (*GovParamDeleteParamIterator, error) {
	logs, sub, err := _GovParam.contract.FilterLogs(opts, "DeleteParam")
	if err != nil {
		return nil, err
	}
	return &GovParamDeleteParamIterator{contract: _GovParam.contract, event: "DeleteParam", logs: logs, sub: sub}, nil
}

// WatchDeleteParam is a free log subscription operation binding the contract event 0xb888810dc5f09b858ca878b61b517f1a3925eb93a2051546d50021db9170fe57.
//
// Solidity: event DeleteParam(string arg0)
func (_GovParam *GovParamFilterer) WatchDeleteParam(opts *bind.WatchOpts, sink chan<- *GovParamDeleteParam) (event.Subscription, error) {
	logs, sub, err := _GovParam.contract.WatchLogs(opts, "DeleteParam")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovParamDeleteParam)
				if err := _GovParam.contract.UnpackLog(event, "DeleteParam", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeleteParam is a log parse operation binding the contract event 0xb888810dc5f09b858ca878b61b517f1a3925eb93a2051546d50021db9170fe57.
//
// Solidity: event DeleteParam(string arg0)
func (_GovParam *GovParamFilterer) ParseDeleteParam(log types.Log) (*GovParamDeleteParam, error) {
	event := new(GovParamDeleteParam)
	if err := _GovParam.contract.UnpackLog(event, "DeleteParam", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovParamInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the GovParam contract.
type GovParamInitializedIterator struct {
	Event *GovParamInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovParamInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovParamInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovParamInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovParamInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovParamInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovParamInitialized represents a Initialized event raised by the GovParam contract.
type GovParamInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GovParam *GovParamFilterer) FilterInitialized(opts *bind.FilterOpts) (*GovParamInitializedIterator, error) {
	logs, sub, err := _GovParam.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GovParamInitializedIterator{contract: _GovParam.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GovParam *GovParamFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GovParamInitialized) (event.Subscription, error) {
	logs, sub, err := _GovParam.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovParamInitialized)
				if err := _GovParam.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GovParam *GovParamFilterer) ParseInitialized(log types.Log) (*GovParamInitialized, error) {
	event := new(GovParamInitialized)
	if err := _GovParam.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovParamOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GovParam contract.
type GovParamOwnershipTransferredIterator struct {
	Event *GovParamOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovParamOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovParamOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovParamOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovParamOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovParamOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovParamOwnershipTransferred represents a OwnershipTransferred event raised by the GovParam contract.
type GovParamOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GovParam *GovParamFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GovParamOwnershipTransferredIterator, error) {
	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GovParam.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GovParamOwnershipTransferredIterator{contract: _GovParam.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GovParam *GovParamFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GovParamOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {
	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GovParam.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovParamOwnershipTransferred)
				if err := _GovParam.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GovParam *GovParamFilterer) ParseOwnershipTransferred(log types.Log) (*GovParamOwnershipTransferred, error) {
	event := new(GovParamOwnershipTransferred)
	if err := _GovParam.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovParamSetParamIterator is returned from FilterSetParam and is used to iterate over the raw logs and unpacked data for SetParam events raised by the GovParam contract.
type GovParamSetParamIterator struct {
	Event *GovParamSetParam // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovParamSetParamIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovParamSetParam)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovParamSetParam)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovParamSetParamIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovParamSetParamIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovParamSetParam represents a SetParam event raised by the GovParam contract.
type GovParamSetParam struct {
	Arg0 string
	Arg1 []byte
	Arg2 uint64
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetParam is a free log retrieval operation binding the contract event 0xf8d4c4710beeb9d48e0bc6888dfcaf21467794f5da39c3bfb2874a2ac3a73c56.
//
// Solidity: event SetParam(string arg0, bytes arg1, uint64 arg2)
func (_GovParam *GovParamFilterer) FilterSetParam(opts *bind.FilterOpts) (*GovParamSetParamIterator, error) {
	logs, sub, err := _GovParam.contract.FilterLogs(opts, "SetParam")
	if err != nil {
		return nil, err
	}
	return &GovParamSetParamIterator{contract: _GovParam.contract, event: "SetParam", logs: logs, sub: sub}, nil
}

// WatchSetParam is a free log subscription operation binding the contract event 0xf8d4c4710beeb9d48e0bc6888dfcaf21467794f5da39c3bfb2874a2ac3a73c56.
//
// Solidity: event SetParam(string arg0, bytes arg1, uint64 arg2)
func (_GovParam *GovParamFilterer) WatchSetParam(opts *bind.WatchOpts, sink chan<- *GovParamSetParam) (event.Subscription, error) {
	logs, sub, err := _GovParam.contract.WatchLogs(opts, "SetParam")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovParamSetParam)
				if err := _GovParam.contract.UnpackLog(event, "SetParam", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetParam is a log parse operation binding the contract event 0xf8d4c4710beeb9d48e0bc6888dfcaf21467794f5da39c3bfb2874a2ac3a73c56.
//
// Solidity: event SetParam(string arg0, bytes arg1, uint64 arg2)
func (_GovParam *GovParamFilterer) ParseSetParam(log types.Log) (*GovParamSetParam, error) {
	event := new(GovParamSetParam)
	if err := _GovParam.contract.UnpackLog(event, "SetParam", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovParamSetParamVotableIterator is returned from FilterSetParamVotable and is used to iterate over the raw logs and unpacked data for SetParamVotable events raised by the GovParam contract.
type GovParamSetParamVotableIterator struct {
	Event *GovParamSetParamVotable // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovParamSetParamVotableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovParamSetParamVotable)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovParamSetParamVotable)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovParamSetParamVotableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovParamSetParamVotableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovParamSetParamVotable represents a SetParamVotable event raised by the GovParam contract.
type GovParamSetParamVotable struct {
	Arg0 string
	Arg1 bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetParamVotable is a free log retrieval operation binding the contract event 0x064a476da825aa99a31c682cc490b34f462457006480bda5cb27151a201213f5.
//
// Solidity: event SetParamVotable(string arg0, bool arg1)
func (_GovParam *GovParamFilterer) FilterSetParamVotable(opts *bind.FilterOpts) (*GovParamSetParamVotableIterator, error) {
	logs, sub, err := _GovParam.contract.FilterLogs(opts, "SetParamVotable")
	if err != nil {
		return nil, err
	}
	return &GovParamSetParamVotableIterator{contract: _GovParam.contract, event: "SetParamVotable", logs: logs, sub: sub}, nil
}

// WatchSetParamVotable is a free log subscription operation binding the contract event 0x064a476da825aa99a31c682cc490b34f462457006480bda5cb27151a201213f5.
//
// Solidity: event SetParamVotable(string arg0, bool arg1)
func (_GovParam *GovParamFilterer) WatchSetParamVotable(opts *bind.WatchOpts, sink chan<- *GovParamSetParamVotable) (event.Subscription, error) {
	logs, sub, err := _GovParam.contract.WatchLogs(opts, "SetParamVotable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovParamSetParamVotable)
				if err := _GovParam.contract.UnpackLog(event, "SetParamVotable", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetParamVotable is a log parse operation binding the contract event 0x064a476da825aa99a31c682cc490b34f462457006480bda5cb27151a201213f5.
//
// Solidity: event SetParamVotable(string arg0, bool arg1)
func (_GovParam *GovParamFilterer) ParseSetParamVotable(log types.Log) (*GovParamSetParamVotable, error) {
	event := new(GovParamSetParamVotable)
	if err := _GovParam.contract.UnpackLog(event, "SetParamVotable", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovParamUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the GovParam contract.
type GovParamUpgradedIterator struct {
	Event *GovParamUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GovParamUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovParamUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GovParamUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GovParamUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovParamUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovParamUpgraded represents a Upgraded event raised by the GovParam contract.
type GovParamUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GovParam *GovParamFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*GovParamUpgradedIterator, error) {
	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GovParam.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &GovParamUpgradedIterator{contract: _GovParam.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GovParam *GovParamFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *GovParamUpgraded, implementation []common.Address) (event.Subscription, error) {
	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _GovParam.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovParamUpgraded)
				if err := _GovParam.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_GovParam *GovParamFilterer) ParseUpgraded(log types.Log) (*GovParamUpgraded, error) {
	event := new(GovParamUpgraded)
	if err := _GovParam.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IBeaconUpgradeableABI is the input ABI used to generate the binding from.
const IBeaconUpgradeableABI = "[{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IBeaconUpgradeableBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const IBeaconUpgradeableBinRuntime = ``

// IBeaconUpgradeableFuncSigs maps the 4-byte function signature to its string representation.
var IBeaconUpgradeableFuncSigs = map[string]string{
	"5c60da1b": "implementation()",
}

// IBeaconUpgradeable is an auto generated Go binding around a Klaytn contract.
type IBeaconUpgradeable struct {
	IBeaconUpgradeableCaller     // Read-only binding to the contract
	IBeaconUpgradeableTransactor // Write-only binding to the contract
	IBeaconUpgradeableFilterer   // Log filterer for contract events
}

// IBeaconUpgradeableCaller is an auto generated read-only Go binding around a Klaytn contract.
type IBeaconUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBeaconUpgradeableTransactor is an auto generated write-only Go binding around a Klaytn contract.
type IBeaconUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBeaconUpgradeableFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type IBeaconUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBeaconUpgradeableSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type IBeaconUpgradeableSession struct {
	Contract     *IBeaconUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IBeaconUpgradeableCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type IBeaconUpgradeableCallerSession struct {
	Contract *IBeaconUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IBeaconUpgradeableTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type IBeaconUpgradeableTransactorSession struct {
	Contract     *IBeaconUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IBeaconUpgradeableRaw is an auto generated low-level Go binding around a Klaytn contract.
type IBeaconUpgradeableRaw struct {
	Contract *IBeaconUpgradeable // Generic contract binding to access the raw methods on
}

// IBeaconUpgradeableCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type IBeaconUpgradeableCallerRaw struct {
	Contract *IBeaconUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// IBeaconUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type IBeaconUpgradeableTransactorRaw struct {
	Contract *IBeaconUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBeaconUpgradeable creates a new instance of IBeaconUpgradeable, bound to a specific deployed contract.
func NewIBeaconUpgradeable(address common.Address, backend bind.ContractBackend) (*IBeaconUpgradeable, error) {
	contract, err := bindIBeaconUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBeaconUpgradeable{IBeaconUpgradeableCaller: IBeaconUpgradeableCaller{contract: contract}, IBeaconUpgradeableTransactor: IBeaconUpgradeableTransactor{contract: contract}, IBeaconUpgradeableFilterer: IBeaconUpgradeableFilterer{contract: contract}}, nil
}

// NewIBeaconUpgradeableCaller creates a new read-only instance of IBeaconUpgradeable, bound to a specific deployed contract.
func NewIBeaconUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*IBeaconUpgradeableCaller, error) {
	contract, err := bindIBeaconUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBeaconUpgradeableCaller{contract: contract}, nil
}

// NewIBeaconUpgradeableTransactor creates a new write-only instance of IBeaconUpgradeable, bound to a specific deployed contract.
func NewIBeaconUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*IBeaconUpgradeableTransactor, error) {
	contract, err := bindIBeaconUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBeaconUpgradeableTransactor{contract: contract}, nil
}

// NewIBeaconUpgradeableFilterer creates a new log filterer instance of IBeaconUpgradeable, bound to a specific deployed contract.
func NewIBeaconUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*IBeaconUpgradeableFilterer, error) {
	contract, err := bindIBeaconUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBeaconUpgradeableFilterer{contract: contract}, nil
}

// bindIBeaconUpgradeable binds a generic wrapper to an already deployed contract.
func bindIBeaconUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IBeaconUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBeaconUpgradeable *IBeaconUpgradeableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IBeaconUpgradeable.Contract.IBeaconUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBeaconUpgradeable *IBeaconUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBeaconUpgradeable.Contract.IBeaconUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBeaconUpgradeable *IBeaconUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBeaconUpgradeable.Contract.IBeaconUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBeaconUpgradeable *IBeaconUpgradeableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IBeaconUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBeaconUpgradeable *IBeaconUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBeaconUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBeaconUpgradeable *IBeaconUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBeaconUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_IBeaconUpgradeable *IBeaconUpgradeableCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	ret0 := new(common.Address)
	out := ret0
	err := _IBeaconUpgradeable.contract.Call(opts, out, "implementation")
	return *ret0, err
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_IBeaconUpgradeable *IBeaconUpgradeableSession) Implementation() (common.Address, error) {
	return _IBeaconUpgradeable.Contract.Implementation(&_IBeaconUpgradeable.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_IBeaconUpgradeable *IBeaconUpgradeableCallerSession) Implementation() (common.Address, error) {
	return _IBeaconUpgradeable.Contract.Implementation(&_IBeaconUpgradeable.CallOpts)
}

// IERC1822ProxiableUpgradeableABI is the input ABI used to generate the binding from.
const IERC1822ProxiableUpgradeableABI = "[{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IERC1822ProxiableUpgradeableBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const IERC1822ProxiableUpgradeableBinRuntime = ``

// IERC1822ProxiableUpgradeableFuncSigs maps the 4-byte function signature to its string representation.
var IERC1822ProxiableUpgradeableFuncSigs = map[string]string{
	"52d1902d": "proxiableUUID()",
}

// IERC1822ProxiableUpgradeable is an auto generated Go binding around a Klaytn contract.
type IERC1822ProxiableUpgradeable struct {
	IERC1822ProxiableUpgradeableCaller     // Read-only binding to the contract
	IERC1822ProxiableUpgradeableTransactor // Write-only binding to the contract
	IERC1822ProxiableUpgradeableFilterer   // Log filterer for contract events
}

// IERC1822ProxiableUpgradeableCaller is an auto generated read-only Go binding around a Klaytn contract.
type IERC1822ProxiableUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1822ProxiableUpgradeableTransactor is an auto generated write-only Go binding around a Klaytn contract.
type IERC1822ProxiableUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1822ProxiableUpgradeableFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type IERC1822ProxiableUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1822ProxiableUpgradeableSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type IERC1822ProxiableUpgradeableSession struct {
	Contract     *IERC1822ProxiableUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IERC1822ProxiableUpgradeableCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type IERC1822ProxiableUpgradeableCallerSession struct {
	Contract *IERC1822ProxiableUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// IERC1822ProxiableUpgradeableTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type IERC1822ProxiableUpgradeableTransactorSession struct {
	Contract     *IERC1822ProxiableUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// IERC1822ProxiableUpgradeableRaw is an auto generated low-level Go binding around a Klaytn contract.
type IERC1822ProxiableUpgradeableRaw struct {
	Contract *IERC1822ProxiableUpgradeable // Generic contract binding to access the raw methods on
}

// IERC1822ProxiableUpgradeableCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type IERC1822ProxiableUpgradeableCallerRaw struct {
	Contract *IERC1822ProxiableUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// IERC1822ProxiableUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type IERC1822ProxiableUpgradeableTransactorRaw struct {
	Contract *IERC1822ProxiableUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC1822ProxiableUpgradeable creates a new instance of IERC1822ProxiableUpgradeable, bound to a specific deployed contract.
func NewIERC1822ProxiableUpgradeable(address common.Address, backend bind.ContractBackend) (*IERC1822ProxiableUpgradeable, error) {
	contract, err := bindIERC1822ProxiableUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC1822ProxiableUpgradeable{IERC1822ProxiableUpgradeableCaller: IERC1822ProxiableUpgradeableCaller{contract: contract}, IERC1822ProxiableUpgradeableTransactor: IERC1822ProxiableUpgradeableTransactor{contract: contract}, IERC1822ProxiableUpgradeableFilterer: IERC1822ProxiableUpgradeableFilterer{contract: contract}}, nil
}

// NewIERC1822ProxiableUpgradeableCaller creates a new read-only instance of IERC1822ProxiableUpgradeable, bound to a specific deployed contract.
func NewIERC1822ProxiableUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*IERC1822ProxiableUpgradeableCaller, error) {
	contract, err := bindIERC1822ProxiableUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1822ProxiableUpgradeableCaller{contract: contract}, nil
}

// NewIERC1822ProxiableUpgradeableTransactor creates a new write-only instance of IERC1822ProxiableUpgradeable, bound to a specific deployed contract.
func NewIERC1822ProxiableUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC1822ProxiableUpgradeableTransactor, error) {
	contract, err := bindIERC1822ProxiableUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1822ProxiableUpgradeableTransactor{contract: contract}, nil
}

// NewIERC1822ProxiableUpgradeableFilterer creates a new log filterer instance of IERC1822ProxiableUpgradeable, bound to a specific deployed contract.
func NewIERC1822ProxiableUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC1822ProxiableUpgradeableFilterer, error) {
	contract, err := bindIERC1822ProxiableUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC1822ProxiableUpgradeableFilterer{contract: contract}, nil
}

// bindIERC1822ProxiableUpgradeable binds a generic wrapper to an already deployed contract.
func bindIERC1822ProxiableUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC1822ProxiableUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC1822ProxiableUpgradeable *IERC1822ProxiableUpgradeableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC1822ProxiableUpgradeable.Contract.IERC1822ProxiableUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC1822ProxiableUpgradeable *IERC1822ProxiableUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC1822ProxiableUpgradeable.Contract.IERC1822ProxiableUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC1822ProxiableUpgradeable *IERC1822ProxiableUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC1822ProxiableUpgradeable.Contract.IERC1822ProxiableUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC1822ProxiableUpgradeable *IERC1822ProxiableUpgradeableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC1822ProxiableUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC1822ProxiableUpgradeable *IERC1822ProxiableUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC1822ProxiableUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC1822ProxiableUpgradeable *IERC1822ProxiableUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC1822ProxiableUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_IERC1822ProxiableUpgradeable *IERC1822ProxiableUpgradeableCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	ret0 := new([32]byte)
	out := ret0
	err := _IERC1822ProxiableUpgradeable.contract.Call(opts, out, "proxiableUUID")
	return *ret0, err
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_IERC1822ProxiableUpgradeable *IERC1822ProxiableUpgradeableSession) ProxiableUUID() ([32]byte, error) {
	return _IERC1822ProxiableUpgradeable.Contract.ProxiableUUID(&_IERC1822ProxiableUpgradeable.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_IERC1822ProxiableUpgradeable *IERC1822ProxiableUpgradeableCallerSession) ProxiableUUID() ([32]byte, error) {
	return _IERC1822ProxiableUpgradeable.Contract.ProxiableUUID(&_IERC1822ProxiableUpgradeable.CallOpts)
}

// IGovParamABI is the input ABI used to generate the binding from.
const IGovParamABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"AddParam\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"DeleteParam\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"SetParam\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"SetParamVotable\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"addParam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"deleteParam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllParams\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllStructParams\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"activation\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"votable\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"internalType\":\"structIGovParam.ParamState\",\"name\":\"prev\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"internalType\":\"structIGovParam.ParamState\",\"name\":\"next\",\"type\":\"tuple\"}],\"internalType\":\"structIGovParam.Param[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getParam\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"activation\",\"type\":\"uint64\"}],\"name\":\"setParam\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"votable\",\"type\":\"bool\"}],\"name\":\"setParamVotable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IGovParamBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const IGovParamBinRuntime = ``

// IGovParamFuncSigs maps the 4-byte function signature to its string representation.
var IGovParamFuncSigs = map[string]string{
	"ede7d2b4": "addParam(string,bytes)",
	"1ba1220d": "deleteParam(string)",
	"a170052e": "getAllParams()",
	"0a6281fc": "getAllStructParams()",
	"5d4f71d4": "getParam(string)",
	"8841067a": "setParam(string,bytes,uint64)",
	"42a3f759": "setParamVotable(string,bool)",
}

// IGovParam is an auto generated Go binding around a Klaytn contract.
type IGovParam struct {
	IGovParamCaller     // Read-only binding to the contract
	IGovParamTransactor // Write-only binding to the contract
	IGovParamFilterer   // Log filterer for contract events
}

// IGovParamCaller is an auto generated read-only Go binding around a Klaytn contract.
type IGovParamCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGovParamTransactor is an auto generated write-only Go binding around a Klaytn contract.
type IGovParamTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGovParamFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type IGovParamFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGovParamSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type IGovParamSession struct {
	Contract     *IGovParam        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGovParamCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type IGovParamCallerSession struct {
	Contract *IGovParamCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IGovParamTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type IGovParamTransactorSession struct {
	Contract     *IGovParamTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IGovParamRaw is an auto generated low-level Go binding around a Klaytn contract.
type IGovParamRaw struct {
	Contract *IGovParam // Generic contract binding to access the raw methods on
}

// IGovParamCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type IGovParamCallerRaw struct {
	Contract *IGovParamCaller // Generic read-only contract binding to access the raw methods on
}

// IGovParamTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type IGovParamTransactorRaw struct {
	Contract *IGovParamTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGovParam creates a new instance of IGovParam, bound to a specific deployed contract.
func NewIGovParam(address common.Address, backend bind.ContractBackend) (*IGovParam, error) {
	contract, err := bindIGovParam(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGovParam{IGovParamCaller: IGovParamCaller{contract: contract}, IGovParamTransactor: IGovParamTransactor{contract: contract}, IGovParamFilterer: IGovParamFilterer{contract: contract}}, nil
}

// NewIGovParamCaller creates a new read-only instance of IGovParam, bound to a specific deployed contract.
func NewIGovParamCaller(address common.Address, caller bind.ContractCaller) (*IGovParamCaller, error) {
	contract, err := bindIGovParam(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGovParamCaller{contract: contract}, nil
}

// NewIGovParamTransactor creates a new write-only instance of IGovParam, bound to a specific deployed contract.
func NewIGovParamTransactor(address common.Address, transactor bind.ContractTransactor) (*IGovParamTransactor, error) {
	contract, err := bindIGovParam(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGovParamTransactor{contract: contract}, nil
}

// NewIGovParamFilterer creates a new log filterer instance of IGovParam, bound to a specific deployed contract.
func NewIGovParamFilterer(address common.Address, filterer bind.ContractFilterer) (*IGovParamFilterer, error) {
	contract, err := bindIGovParam(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGovParamFilterer{contract: contract}, nil
}

// bindIGovParam binds a generic wrapper to an already deployed contract.
func bindIGovParam(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IGovParamABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGovParam *IGovParamRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IGovParam.Contract.IGovParamCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGovParam *IGovParamRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGovParam.Contract.IGovParamTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGovParam *IGovParamRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGovParam.Contract.IGovParamTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGovParam *IGovParamCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IGovParam.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGovParam *IGovParamTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGovParam.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGovParam *IGovParamTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGovParam.Contract.contract.Transact(opts, method, params...)
}

// GetAllParams is a free data retrieval call binding the contract method 0xa170052e.
//
// Solidity: function getAllParams() view returns(string[], bytes[])
func (_IGovParam *IGovParamCaller) GetAllParams(opts *bind.CallOpts) ([]string, [][]byte, error) {
	var (
		ret0 = new([]string)
		ret1 = new([][]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _IGovParam.contract.Call(opts, out, "getAllParams")
	return *ret0, *ret1, err
}

// GetAllParams is a free data retrieval call binding the contract method 0xa170052e.
//
// Solidity: function getAllParams() view returns(string[], bytes[])
func (_IGovParam *IGovParamSession) GetAllParams() ([]string, [][]byte, error) {
	return _IGovParam.Contract.GetAllParams(&_IGovParam.CallOpts)
}

// GetAllParams is a free data retrieval call binding the contract method 0xa170052e.
//
// Solidity: function getAllParams() view returns(string[], bytes[])
func (_IGovParam *IGovParamCallerSession) GetAllParams() ([]string, [][]byte, error) {
	return _IGovParam.Contract.GetAllParams(&_IGovParam.CallOpts)
}

// GetAllStructParams is a free data retrieval call binding the contract method 0x0a6281fc.
//
// Solidity: function getAllStructParams() view returns(string[], (uint64,bool,(bytes,bool),(bytes,bool))[])
func (_IGovParam *IGovParamCaller) GetAllStructParams(opts *bind.CallOpts) ([]string, []IGovParamParam, error) {
	var (
		ret0 = new([]string)
		ret1 = new([]IGovParamParam)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _IGovParam.contract.Call(opts, out, "getAllStructParams")
	return *ret0, *ret1, err
}

// GetAllStructParams is a free data retrieval call binding the contract method 0x0a6281fc.
//
// Solidity: function getAllStructParams() view returns(string[], (uint64,bool,(bytes,bool),(bytes,bool))[])
func (_IGovParam *IGovParamSession) GetAllStructParams() ([]string, []IGovParamParam, error) {
	return _IGovParam.Contract.GetAllStructParams(&_IGovParam.CallOpts)
}

// GetAllStructParams is a free data retrieval call binding the contract method 0x0a6281fc.
//
// Solidity: function getAllStructParams() view returns(string[], (uint64,bool,(bytes,bool),(bytes,bool))[])
func (_IGovParam *IGovParamCallerSession) GetAllStructParams() ([]string, []IGovParamParam, error) {
	return _IGovParam.Contract.GetAllStructParams(&_IGovParam.CallOpts)
}

// GetParam is a free data retrieval call binding the contract method 0x5d4f71d4.
//
// Solidity: function getParam(string name) view returns(bytes)
func (_IGovParam *IGovParamCaller) GetParam(opts *bind.CallOpts, name string) ([]byte, error) {
	ret0 := new([]byte)
	out := ret0
	err := _IGovParam.contract.Call(opts, out, "getParam", name)
	return *ret0, err
}

// GetParam is a free data retrieval call binding the contract method 0x5d4f71d4.
//
// Solidity: function getParam(string name) view returns(bytes)
func (_IGovParam *IGovParamSession) GetParam(name string) ([]byte, error) {
	return _IGovParam.Contract.GetParam(&_IGovParam.CallOpts, name)
}

// GetParam is a free data retrieval call binding the contract method 0x5d4f71d4.
//
// Solidity: function getParam(string name) view returns(bytes)
func (_IGovParam *IGovParamCallerSession) GetParam(name string) ([]byte, error) {
	return _IGovParam.Contract.GetParam(&_IGovParam.CallOpts, name)
}

// AddParam is a paid mutator transaction binding the contract method 0xede7d2b4.
//
// Solidity: function addParam(string name, bytes value) returns()
func (_IGovParam *IGovParamTransactor) AddParam(opts *bind.TransactOpts, name string, value []byte) (*types.Transaction, error) {
	return _IGovParam.contract.Transact(opts, "addParam", name, value)
}

// AddParam is a paid mutator transaction binding the contract method 0xede7d2b4.
//
// Solidity: function addParam(string name, bytes value) returns()
func (_IGovParam *IGovParamSession) AddParam(name string, value []byte) (*types.Transaction, error) {
	return _IGovParam.Contract.AddParam(&_IGovParam.TransactOpts, name, value)
}

// AddParam is a paid mutator transaction binding the contract method 0xede7d2b4.
//
// Solidity: function addParam(string name, bytes value) returns()
func (_IGovParam *IGovParamTransactorSession) AddParam(name string, value []byte) (*types.Transaction, error) {
	return _IGovParam.Contract.AddParam(&_IGovParam.TransactOpts, name, value)
}

// DeleteParam is a paid mutator transaction binding the contract method 0x1ba1220d.
//
// Solidity: function deleteParam(string name) returns()
func (_IGovParam *IGovParamTransactor) DeleteParam(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _IGovParam.contract.Transact(opts, "deleteParam", name)
}

// DeleteParam is a paid mutator transaction binding the contract method 0x1ba1220d.
//
// Solidity: function deleteParam(string name) returns()
func (_IGovParam *IGovParamSession) DeleteParam(name string) (*types.Transaction, error) {
	return _IGovParam.Contract.DeleteParam(&_IGovParam.TransactOpts, name)
}

// DeleteParam is a paid mutator transaction binding the contract method 0x1ba1220d.
//
// Solidity: function deleteParam(string name) returns()
func (_IGovParam *IGovParamTransactorSession) DeleteParam(name string) (*types.Transaction, error) {
	return _IGovParam.Contract.DeleteParam(&_IGovParam.TransactOpts, name)
}

// SetParam is a paid mutator transaction binding the contract method 0x8841067a.
//
// Solidity: function setParam(string name, bytes value, uint64 activation) returns()
func (_IGovParam *IGovParamTransactor) SetParam(opts *bind.TransactOpts, name string, value []byte, activation uint64) (*types.Transaction, error) {
	return _IGovParam.contract.Transact(opts, "setParam", name, value, activation)
}

// SetParam is a paid mutator transaction binding the contract method 0x8841067a.
//
// Solidity: function setParam(string name, bytes value, uint64 activation) returns()
func (_IGovParam *IGovParamSession) SetParam(name string, value []byte, activation uint64) (*types.Transaction, error) {
	return _IGovParam.Contract.SetParam(&_IGovParam.TransactOpts, name, value, activation)
}

// SetParam is a paid mutator transaction binding the contract method 0x8841067a.
//
// Solidity: function setParam(string name, bytes value, uint64 activation) returns()
func (_IGovParam *IGovParamTransactorSession) SetParam(name string, value []byte, activation uint64) (*types.Transaction, error) {
	return _IGovParam.Contract.SetParam(&_IGovParam.TransactOpts, name, value, activation)
}

// SetParamVotable is a paid mutator transaction binding the contract method 0x42a3f759.
//
// Solidity: function setParamVotable(string name, bool votable) returns()
func (_IGovParam *IGovParamTransactor) SetParamVotable(opts *bind.TransactOpts, name string, votable bool) (*types.Transaction, error) {
	return _IGovParam.contract.Transact(opts, "setParamVotable", name, votable)
}

// SetParamVotable is a paid mutator transaction binding the contract method 0x42a3f759.
//
// Solidity: function setParamVotable(string name, bool votable) returns()
func (_IGovParam *IGovParamSession) SetParamVotable(name string, votable bool) (*types.Transaction, error) {
	return _IGovParam.Contract.SetParamVotable(&_IGovParam.TransactOpts, name, votable)
}

// SetParamVotable is a paid mutator transaction binding the contract method 0x42a3f759.
//
// Solidity: function setParamVotable(string name, bool votable) returns()
func (_IGovParam *IGovParamTransactorSession) SetParamVotable(name string, votable bool) (*types.Transaction, error) {
	return _IGovParam.Contract.SetParamVotable(&_IGovParam.TransactOpts, name, votable)
}

// IGovParamAddParamIterator is returned from FilterAddParam and is used to iterate over the raw logs and unpacked data for AddParam events raised by the IGovParam contract.
type IGovParamAddParamIterator struct {
	Event *IGovParamAddParam // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IGovParamAddParamIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGovParamAddParam)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IGovParamAddParam)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IGovParamAddParamIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGovParamAddParamIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGovParamAddParam represents a AddParam event raised by the IGovParam contract.
type IGovParamAddParam struct {
	Arg0 string
	Arg1 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddParam is a free log retrieval operation binding the contract event 0x83878da4b3d5b833aa14a193f9b57d482ab60c1e18ef91dce48f4584f85b8f72.
//
// Solidity: event AddParam(string arg0, bytes arg1)
func (_IGovParam *IGovParamFilterer) FilterAddParam(opts *bind.FilterOpts) (*IGovParamAddParamIterator, error) {
	logs, sub, err := _IGovParam.contract.FilterLogs(opts, "AddParam")
	if err != nil {
		return nil, err
	}
	return &IGovParamAddParamIterator{contract: _IGovParam.contract, event: "AddParam", logs: logs, sub: sub}, nil
}

// WatchAddParam is a free log subscription operation binding the contract event 0x83878da4b3d5b833aa14a193f9b57d482ab60c1e18ef91dce48f4584f85b8f72.
//
// Solidity: event AddParam(string arg0, bytes arg1)
func (_IGovParam *IGovParamFilterer) WatchAddParam(opts *bind.WatchOpts, sink chan<- *IGovParamAddParam) (event.Subscription, error) {
	logs, sub, err := _IGovParam.contract.WatchLogs(opts, "AddParam")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGovParamAddParam)
				if err := _IGovParam.contract.UnpackLog(event, "AddParam", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddParam is a log parse operation binding the contract event 0x83878da4b3d5b833aa14a193f9b57d482ab60c1e18ef91dce48f4584f85b8f72.
//
// Solidity: event AddParam(string arg0, bytes arg1)
func (_IGovParam *IGovParamFilterer) ParseAddParam(log types.Log) (*IGovParamAddParam, error) {
	event := new(IGovParamAddParam)
	if err := _IGovParam.contract.UnpackLog(event, "AddParam", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IGovParamDeleteParamIterator is returned from FilterDeleteParam and is used to iterate over the raw logs and unpacked data for DeleteParam events raised by the IGovParam contract.
type IGovParamDeleteParamIterator struct {
	Event *IGovParamDeleteParam // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IGovParamDeleteParamIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGovParamDeleteParam)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IGovParamDeleteParam)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IGovParamDeleteParamIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGovParamDeleteParamIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGovParamDeleteParam represents a DeleteParam event raised by the IGovParam contract.
type IGovParamDeleteParam struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDeleteParam is a free log retrieval operation binding the contract event 0xb888810dc5f09b858ca878b61b517f1a3925eb93a2051546d50021db9170fe57.
//
// Solidity: event DeleteParam(string arg0)
func (_IGovParam *IGovParamFilterer) FilterDeleteParam(opts *bind.FilterOpts) (*IGovParamDeleteParamIterator, error) {
	logs, sub, err := _IGovParam.contract.FilterLogs(opts, "DeleteParam")
	if err != nil {
		return nil, err
	}
	return &IGovParamDeleteParamIterator{contract: _IGovParam.contract, event: "DeleteParam", logs: logs, sub: sub}, nil
}

// WatchDeleteParam is a free log subscription operation binding the contract event 0xb888810dc5f09b858ca878b61b517f1a3925eb93a2051546d50021db9170fe57.
//
// Solidity: event DeleteParam(string arg0)
func (_IGovParam *IGovParamFilterer) WatchDeleteParam(opts *bind.WatchOpts, sink chan<- *IGovParamDeleteParam) (event.Subscription, error) {
	logs, sub, err := _IGovParam.contract.WatchLogs(opts, "DeleteParam")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGovParamDeleteParam)
				if err := _IGovParam.contract.UnpackLog(event, "DeleteParam", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeleteParam is a log parse operation binding the contract event 0xb888810dc5f09b858ca878b61b517f1a3925eb93a2051546d50021db9170fe57.
//
// Solidity: event DeleteParam(string arg0)
func (_IGovParam *IGovParamFilterer) ParseDeleteParam(log types.Log) (*IGovParamDeleteParam, error) {
	event := new(IGovParamDeleteParam)
	if err := _IGovParam.contract.UnpackLog(event, "DeleteParam", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IGovParamSetParamIterator is returned from FilterSetParam and is used to iterate over the raw logs and unpacked data for SetParam events raised by the IGovParam contract.
type IGovParamSetParamIterator struct {
	Event *IGovParamSetParam // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IGovParamSetParamIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGovParamSetParam)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IGovParamSetParam)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IGovParamSetParamIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGovParamSetParamIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGovParamSetParam represents a SetParam event raised by the IGovParam contract.
type IGovParamSetParam struct {
	Arg0 string
	Arg1 []byte
	Arg2 uint64
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetParam is a free log retrieval operation binding the contract event 0xf8d4c4710beeb9d48e0bc6888dfcaf21467794f5da39c3bfb2874a2ac3a73c56.
//
// Solidity: event SetParam(string arg0, bytes arg1, uint64 arg2)
func (_IGovParam *IGovParamFilterer) FilterSetParam(opts *bind.FilterOpts) (*IGovParamSetParamIterator, error) {
	logs, sub, err := _IGovParam.contract.FilterLogs(opts, "SetParam")
	if err != nil {
		return nil, err
	}
	return &IGovParamSetParamIterator{contract: _IGovParam.contract, event: "SetParam", logs: logs, sub: sub}, nil
}

// WatchSetParam is a free log subscription operation binding the contract event 0xf8d4c4710beeb9d48e0bc6888dfcaf21467794f5da39c3bfb2874a2ac3a73c56.
//
// Solidity: event SetParam(string arg0, bytes arg1, uint64 arg2)
func (_IGovParam *IGovParamFilterer) WatchSetParam(opts *bind.WatchOpts, sink chan<- *IGovParamSetParam) (event.Subscription, error) {
	logs, sub, err := _IGovParam.contract.WatchLogs(opts, "SetParam")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGovParamSetParam)
				if err := _IGovParam.contract.UnpackLog(event, "SetParam", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetParam is a log parse operation binding the contract event 0xf8d4c4710beeb9d48e0bc6888dfcaf21467794f5da39c3bfb2874a2ac3a73c56.
//
// Solidity: event SetParam(string arg0, bytes arg1, uint64 arg2)
func (_IGovParam *IGovParamFilterer) ParseSetParam(log types.Log) (*IGovParamSetParam, error) {
	event := new(IGovParamSetParam)
	if err := _IGovParam.contract.UnpackLog(event, "SetParam", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IGovParamSetParamVotableIterator is returned from FilterSetParamVotable and is used to iterate over the raw logs and unpacked data for SetParamVotable events raised by the IGovParam contract.
type IGovParamSetParamVotableIterator struct {
	Event *IGovParamSetParamVotable // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IGovParamSetParamVotableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGovParamSetParamVotable)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IGovParamSetParamVotable)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IGovParamSetParamVotableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGovParamSetParamVotableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGovParamSetParamVotable represents a SetParamVotable event raised by the IGovParam contract.
type IGovParamSetParamVotable struct {
	Arg0 string
	Arg1 bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetParamVotable is a free log retrieval operation binding the contract event 0x064a476da825aa99a31c682cc490b34f462457006480bda5cb27151a201213f5.
//
// Solidity: event SetParamVotable(string arg0, bool arg1)
func (_IGovParam *IGovParamFilterer) FilterSetParamVotable(opts *bind.FilterOpts) (*IGovParamSetParamVotableIterator, error) {
	logs, sub, err := _IGovParam.contract.FilterLogs(opts, "SetParamVotable")
	if err != nil {
		return nil, err
	}
	return &IGovParamSetParamVotableIterator{contract: _IGovParam.contract, event: "SetParamVotable", logs: logs, sub: sub}, nil
}

// WatchSetParamVotable is a free log subscription operation binding the contract event 0x064a476da825aa99a31c682cc490b34f462457006480bda5cb27151a201213f5.
//
// Solidity: event SetParamVotable(string arg0, bool arg1)
func (_IGovParam *IGovParamFilterer) WatchSetParamVotable(opts *bind.WatchOpts, sink chan<- *IGovParamSetParamVotable) (event.Subscription, error) {
	logs, sub, err := _IGovParam.contract.WatchLogs(opts, "SetParamVotable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGovParamSetParamVotable)
				if err := _IGovParam.contract.UnpackLog(event, "SetParamVotable", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetParamVotable is a log parse operation binding the contract event 0x064a476da825aa99a31c682cc490b34f462457006480bda5cb27151a201213f5.
//
// Solidity: event SetParamVotable(string arg0, bool arg1)
func (_IGovParam *IGovParamFilterer) ParseSetParamVotable(log types.Log) (*IGovParamSetParamVotable, error) {
	event := new(IGovParamSetParamVotable)
	if err := _IGovParam.contract.UnpackLog(event, "SetParamVotable", log); err != nil {
		return nil, err
	}
	return event, nil
}

// InitializableABI is the input ABI used to generate the binding from.
const InitializableABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"}]"

// InitializableBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const InitializableBinRuntime = ``

// Initializable is an auto generated Go binding around a Klaytn contract.
type Initializable struct {
	InitializableCaller     // Read-only binding to the contract
	InitializableTransactor // Write-only binding to the contract
	InitializableFilterer   // Log filterer for contract events
}

// InitializableCaller is an auto generated read-only Go binding around a Klaytn contract.
type InitializableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableTransactor is an auto generated write-only Go binding around a Klaytn contract.
type InitializableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type InitializableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type InitializableSession struct {
	Contract     *Initializable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InitializableCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type InitializableCallerSession struct {
	Contract *InitializableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// InitializableTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type InitializableTransactorSession struct {
	Contract     *InitializableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// InitializableRaw is an auto generated low-level Go binding around a Klaytn contract.
type InitializableRaw struct {
	Contract *Initializable // Generic contract binding to access the raw methods on
}

// InitializableCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type InitializableCallerRaw struct {
	Contract *InitializableCaller // Generic read-only contract binding to access the raw methods on
}

// InitializableTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type InitializableTransactorRaw struct {
	Contract *InitializableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInitializable creates a new instance of Initializable, bound to a specific deployed contract.
func NewInitializable(address common.Address, backend bind.ContractBackend) (*Initializable, error) {
	contract, err := bindInitializable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Initializable{InitializableCaller: InitializableCaller{contract: contract}, InitializableTransactor: InitializableTransactor{contract: contract}, InitializableFilterer: InitializableFilterer{contract: contract}}, nil
}

// NewInitializableCaller creates a new read-only instance of Initializable, bound to a specific deployed contract.
func NewInitializableCaller(address common.Address, caller bind.ContractCaller) (*InitializableCaller, error) {
	contract, err := bindInitializable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InitializableCaller{contract: contract}, nil
}

// NewInitializableTransactor creates a new write-only instance of Initializable, bound to a specific deployed contract.
func NewInitializableTransactor(address common.Address, transactor bind.ContractTransactor) (*InitializableTransactor, error) {
	contract, err := bindInitializable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InitializableTransactor{contract: contract}, nil
}

// NewInitializableFilterer creates a new log filterer instance of Initializable, bound to a specific deployed contract.
func NewInitializableFilterer(address common.Address, filterer bind.ContractFilterer) (*InitializableFilterer, error) {
	contract, err := bindInitializable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InitializableFilterer{contract: contract}, nil
}

// bindInitializable binds a generic wrapper to an already deployed contract.
func bindInitializable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InitializableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Initializable *InitializableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Initializable.Contract.InitializableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Initializable *InitializableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Initializable.Contract.InitializableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Initializable *InitializableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Initializable.Contract.InitializableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Initializable *InitializableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Initializable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Initializable *InitializableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Initializable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Initializable *InitializableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Initializable.Contract.contract.Transact(opts, method, params...)
}

// InitializableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Initializable contract.
type InitializableInitializedIterator struct {
	Event *InitializableInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *InitializableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InitializableInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(InitializableInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *InitializableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InitializableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InitializableInitialized represents a Initialized event raised by the Initializable contract.
type InitializableInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Initializable *InitializableFilterer) FilterInitialized(opts *bind.FilterOpts) (*InitializableInitializedIterator, error) {
	logs, sub, err := _Initializable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &InitializableInitializedIterator{contract: _Initializable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Initializable *InitializableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *InitializableInitialized) (event.Subscription, error) {
	logs, sub, err := _Initializable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InitializableInitialized)
				if err := _Initializable.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Initializable *InitializableFilterer) ParseInitialized(log types.Log) (*InitializableInitialized, error) {
	event := new(InitializableInitialized)
	if err := _Initializable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OwnableUpgradeableABI is the input ABI used to generate the binding from.
const OwnableUpgradeableABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OwnableUpgradeableBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const OwnableUpgradeableBinRuntime = ``

// OwnableUpgradeableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableUpgradeableFuncSigs = map[string]string{
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
}

// OwnableUpgradeable is an auto generated Go binding around a Klaytn contract.
type OwnableUpgradeable struct {
	OwnableUpgradeableCaller     // Read-only binding to the contract
	OwnableUpgradeableTransactor // Write-only binding to the contract
	OwnableUpgradeableFilterer   // Log filterer for contract events
}

// OwnableUpgradeableCaller is an auto generated read-only Go binding around a Klaytn contract.
type OwnableUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableUpgradeableTransactor is an auto generated write-only Go binding around a Klaytn contract.
type OwnableUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableUpgradeableFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type OwnableUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableUpgradeableSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type OwnableUpgradeableSession struct {
	Contract     *OwnableUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// OwnableUpgradeableCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type OwnableUpgradeableCallerSession struct {
	Contract *OwnableUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// OwnableUpgradeableTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type OwnableUpgradeableTransactorSession struct {
	Contract     *OwnableUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// OwnableUpgradeableRaw is an auto generated low-level Go binding around a Klaytn contract.
type OwnableUpgradeableRaw struct {
	Contract *OwnableUpgradeable // Generic contract binding to access the raw methods on
}

// OwnableUpgradeableCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type OwnableUpgradeableCallerRaw struct {
	Contract *OwnableUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type OwnableUpgradeableTransactorRaw struct {
	Contract *OwnableUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnableUpgradeable creates a new instance of OwnableUpgradeable, bound to a specific deployed contract.
func NewOwnableUpgradeable(address common.Address, backend bind.ContractBackend) (*OwnableUpgradeable, error) {
	contract, err := bindOwnableUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeable{OwnableUpgradeableCaller: OwnableUpgradeableCaller{contract: contract}, OwnableUpgradeableTransactor: OwnableUpgradeableTransactor{contract: contract}, OwnableUpgradeableFilterer: OwnableUpgradeableFilterer{contract: contract}}, nil
}

// NewOwnableUpgradeableCaller creates a new read-only instance of OwnableUpgradeable, bound to a specific deployed contract.
func NewOwnableUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*OwnableUpgradeableCaller, error) {
	contract, err := bindOwnableUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeableCaller{contract: contract}, nil
}

// NewOwnableUpgradeableTransactor creates a new write-only instance of OwnableUpgradeable, bound to a specific deployed contract.
func NewOwnableUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableUpgradeableTransactor, error) {
	contract, err := bindOwnableUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeableTransactor{contract: contract}, nil
}

// NewOwnableUpgradeableFilterer creates a new log filterer instance of OwnableUpgradeable, bound to a specific deployed contract.
func NewOwnableUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableUpgradeableFilterer, error) {
	contract, err := bindOwnableUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeableFilterer{contract: contract}, nil
}

// bindOwnableUpgradeable binds a generic wrapper to an already deployed contract.
func bindOwnableUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnableUpgradeable *OwnableUpgradeableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OwnableUpgradeable.Contract.OwnableUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnableUpgradeable *OwnableUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.OwnableUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnableUpgradeable *OwnableUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.OwnableUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnableUpgradeable *OwnableUpgradeableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OwnableUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnableUpgradeable *OwnableUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnableUpgradeable *OwnableUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnableUpgradeable *OwnableUpgradeableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	ret0 := new(common.Address)
	out := ret0
	err := _OwnableUpgradeable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnableUpgradeable *OwnableUpgradeableSession) Owner() (common.Address, error) {
	return _OwnableUpgradeable.Contract.Owner(&_OwnableUpgradeable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnableUpgradeable *OwnableUpgradeableCallerSession) Owner() (common.Address, error) {
	return _OwnableUpgradeable.Contract.Owner(&_OwnableUpgradeable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OwnableUpgradeable *OwnableUpgradeableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableUpgradeable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OwnableUpgradeable *OwnableUpgradeableSession) RenounceOwnership() (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.RenounceOwnership(&_OwnableUpgradeable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OwnableUpgradeable *OwnableUpgradeableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.RenounceOwnership(&_OwnableUpgradeable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OwnableUpgradeable *OwnableUpgradeableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OwnableUpgradeable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OwnableUpgradeable *OwnableUpgradeableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.TransferOwnership(&_OwnableUpgradeable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OwnableUpgradeable *OwnableUpgradeableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.TransferOwnership(&_OwnableUpgradeable.TransactOpts, newOwner)
}

// OwnableUpgradeableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the OwnableUpgradeable contract.
type OwnableUpgradeableInitializedIterator struct {
	Event *OwnableUpgradeableInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OwnableUpgradeableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableUpgradeableInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OwnableUpgradeableInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OwnableUpgradeableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableUpgradeableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableUpgradeableInitialized represents a Initialized event raised by the OwnableUpgradeable contract.
type OwnableUpgradeableInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OwnableUpgradeable *OwnableUpgradeableFilterer) FilterInitialized(opts *bind.FilterOpts) (*OwnableUpgradeableInitializedIterator, error) {
	logs, sub, err := _OwnableUpgradeable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeableInitializedIterator{contract: _OwnableUpgradeable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OwnableUpgradeable *OwnableUpgradeableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OwnableUpgradeableInitialized) (event.Subscription, error) {
	logs, sub, err := _OwnableUpgradeable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableUpgradeableInitialized)
				if err := _OwnableUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OwnableUpgradeable *OwnableUpgradeableFilterer) ParseInitialized(log types.Log) (*OwnableUpgradeableInitialized, error) {
	event := new(OwnableUpgradeableInitialized)
	if err := _OwnableUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OwnableUpgradeableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OwnableUpgradeable contract.
type OwnableUpgradeableOwnershipTransferredIterator struct {
	Event *OwnableUpgradeableOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OwnableUpgradeableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableUpgradeableOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OwnableUpgradeableOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OwnableUpgradeableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableUpgradeableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableUpgradeableOwnershipTransferred represents a OwnershipTransferred event raised by the OwnableUpgradeable contract.
type OwnableUpgradeableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OwnableUpgradeable *OwnableUpgradeableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableUpgradeableOwnershipTransferredIterator, error) {
	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OwnableUpgradeable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeableOwnershipTransferredIterator{contract: _OwnableUpgradeable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OwnableUpgradeable *OwnableUpgradeableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableUpgradeableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {
	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OwnableUpgradeable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableUpgradeableOwnershipTransferred)
				if err := _OwnableUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OwnableUpgradeable *OwnableUpgradeableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableUpgradeableOwnershipTransferred, error) {
	event := new(OwnableUpgradeableOwnershipTransferred)
	if err := _OwnableUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StorageSlotUpgradeableABI is the input ABI used to generate the binding from.
const StorageSlotUpgradeableABI = "[]"

// StorageSlotUpgradeableBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const StorageSlotUpgradeableBinRuntime = `73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220128228e2f9819bb6e6e7b58d55976d23c176dadd48636cafe6403f5a0c1b9f4f64736f6c634300080e0033`

// StorageSlotUpgradeableBin is the compiled bytecode used for deploying new contracts.
var StorageSlotUpgradeableBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220128228e2f9819bb6e6e7b58d55976d23c176dadd48636cafe6403f5a0c1b9f4f64736f6c634300080e0033"

// DeployStorageSlotUpgradeable deploys a new Klaytn contract, binding an instance of StorageSlotUpgradeable to it.
func DeployStorageSlotUpgradeable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StorageSlotUpgradeable, error) {
	parsed, err := abi.JSON(strings.NewReader(StorageSlotUpgradeableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StorageSlotUpgradeableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StorageSlotUpgradeable{StorageSlotUpgradeableCaller: StorageSlotUpgradeableCaller{contract: contract}, StorageSlotUpgradeableTransactor: StorageSlotUpgradeableTransactor{contract: contract}, StorageSlotUpgradeableFilterer: StorageSlotUpgradeableFilterer{contract: contract}}, nil
}

// StorageSlotUpgradeable is an auto generated Go binding around a Klaytn contract.
type StorageSlotUpgradeable struct {
	StorageSlotUpgradeableCaller     // Read-only binding to the contract
	StorageSlotUpgradeableTransactor // Write-only binding to the contract
	StorageSlotUpgradeableFilterer   // Log filterer for contract events
}

// StorageSlotUpgradeableCaller is an auto generated read-only Go binding around a Klaytn contract.
type StorageSlotUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSlotUpgradeableTransactor is an auto generated write-only Go binding around a Klaytn contract.
type StorageSlotUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSlotUpgradeableFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type StorageSlotUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSlotUpgradeableSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type StorageSlotUpgradeableSession struct {
	Contract     *StorageSlotUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StorageSlotUpgradeableCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type StorageSlotUpgradeableCallerSession struct {
	Contract *StorageSlotUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// StorageSlotUpgradeableTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type StorageSlotUpgradeableTransactorSession struct {
	Contract     *StorageSlotUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// StorageSlotUpgradeableRaw is an auto generated low-level Go binding around a Klaytn contract.
type StorageSlotUpgradeableRaw struct {
	Contract *StorageSlotUpgradeable // Generic contract binding to access the raw methods on
}

// StorageSlotUpgradeableCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type StorageSlotUpgradeableCallerRaw struct {
	Contract *StorageSlotUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// StorageSlotUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type StorageSlotUpgradeableTransactorRaw struct {
	Contract *StorageSlotUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorageSlotUpgradeable creates a new instance of StorageSlotUpgradeable, bound to a specific deployed contract.
func NewStorageSlotUpgradeable(address common.Address, backend bind.ContractBackend) (*StorageSlotUpgradeable, error) {
	contract, err := bindStorageSlotUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StorageSlotUpgradeable{StorageSlotUpgradeableCaller: StorageSlotUpgradeableCaller{contract: contract}, StorageSlotUpgradeableTransactor: StorageSlotUpgradeableTransactor{contract: contract}, StorageSlotUpgradeableFilterer: StorageSlotUpgradeableFilterer{contract: contract}}, nil
}

// NewStorageSlotUpgradeableCaller creates a new read-only instance of StorageSlotUpgradeable, bound to a specific deployed contract.
func NewStorageSlotUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*StorageSlotUpgradeableCaller, error) {
	contract, err := bindStorageSlotUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageSlotUpgradeableCaller{contract: contract}, nil
}

// NewStorageSlotUpgradeableTransactor creates a new write-only instance of StorageSlotUpgradeable, bound to a specific deployed contract.
func NewStorageSlotUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageSlotUpgradeableTransactor, error) {
	contract, err := bindStorageSlotUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageSlotUpgradeableTransactor{contract: contract}, nil
}

// NewStorageSlotUpgradeableFilterer creates a new log filterer instance of StorageSlotUpgradeable, bound to a specific deployed contract.
func NewStorageSlotUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageSlotUpgradeableFilterer, error) {
	contract, err := bindStorageSlotUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageSlotUpgradeableFilterer{contract: contract}, nil
}

// bindStorageSlotUpgradeable binds a generic wrapper to an already deployed contract.
func bindStorageSlotUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StorageSlotUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageSlotUpgradeable *StorageSlotUpgradeableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StorageSlotUpgradeable.Contract.StorageSlotUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageSlotUpgradeable *StorageSlotUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageSlotUpgradeable.Contract.StorageSlotUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageSlotUpgradeable *StorageSlotUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageSlotUpgradeable.Contract.StorageSlotUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageSlotUpgradeable *StorageSlotUpgradeableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StorageSlotUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageSlotUpgradeable *StorageSlotUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageSlotUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageSlotUpgradeable *StorageSlotUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageSlotUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// UUPSUpgradeableABI is the input ABI used to generate the binding from.
const UUPSUpgradeableABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// UUPSUpgradeableBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const UUPSUpgradeableBinRuntime = ``

// UUPSUpgradeableFuncSigs maps the 4-byte function signature to its string representation.
var UUPSUpgradeableFuncSigs = map[string]string{
	"52d1902d": "proxiableUUID()",
	"3659cfe6": "upgradeTo(address)",
	"4f1ef286": "upgradeToAndCall(address,bytes)",
}

// UUPSUpgradeable is an auto generated Go binding around a Klaytn contract.
type UUPSUpgradeable struct {
	UUPSUpgradeableCaller     // Read-only binding to the contract
	UUPSUpgradeableTransactor // Write-only binding to the contract
	UUPSUpgradeableFilterer   // Log filterer for contract events
}

// UUPSUpgradeableCaller is an auto generated read-only Go binding around a Klaytn contract.
type UUPSUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UUPSUpgradeableTransactor is an auto generated write-only Go binding around a Klaytn contract.
type UUPSUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UUPSUpgradeableFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type UUPSUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UUPSUpgradeableSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type UUPSUpgradeableSession struct {
	Contract     *UUPSUpgradeable  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UUPSUpgradeableCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type UUPSUpgradeableCallerSession struct {
	Contract *UUPSUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// UUPSUpgradeableTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type UUPSUpgradeableTransactorSession struct {
	Contract     *UUPSUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// UUPSUpgradeableRaw is an auto generated low-level Go binding around a Klaytn contract.
type UUPSUpgradeableRaw struct {
	Contract *UUPSUpgradeable // Generic contract binding to access the raw methods on
}

// UUPSUpgradeableCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type UUPSUpgradeableCallerRaw struct {
	Contract *UUPSUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// UUPSUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type UUPSUpgradeableTransactorRaw struct {
	Contract *UUPSUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUUPSUpgradeable creates a new instance of UUPSUpgradeable, bound to a specific deployed contract.
func NewUUPSUpgradeable(address common.Address, backend bind.ContractBackend) (*UUPSUpgradeable, error) {
	contract, err := bindUUPSUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UUPSUpgradeable{UUPSUpgradeableCaller: UUPSUpgradeableCaller{contract: contract}, UUPSUpgradeableTransactor: UUPSUpgradeableTransactor{contract: contract}, UUPSUpgradeableFilterer: UUPSUpgradeableFilterer{contract: contract}}, nil
}

// NewUUPSUpgradeableCaller creates a new read-only instance of UUPSUpgradeable, bound to a specific deployed contract.
func NewUUPSUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*UUPSUpgradeableCaller, error) {
	contract, err := bindUUPSUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UUPSUpgradeableCaller{contract: contract}, nil
}

// NewUUPSUpgradeableTransactor creates a new write-only instance of UUPSUpgradeable, bound to a specific deployed contract.
func NewUUPSUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*UUPSUpgradeableTransactor, error) {
	contract, err := bindUUPSUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UUPSUpgradeableTransactor{contract: contract}, nil
}

// NewUUPSUpgradeableFilterer creates a new log filterer instance of UUPSUpgradeable, bound to a specific deployed contract.
func NewUUPSUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*UUPSUpgradeableFilterer, error) {
	contract, err := bindUUPSUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UUPSUpgradeableFilterer{contract: contract}, nil
}

// bindUUPSUpgradeable binds a generic wrapper to an already deployed contract.
func bindUUPSUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UUPSUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UUPSUpgradeable *UUPSUpgradeableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UUPSUpgradeable.Contract.UUPSUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UUPSUpgradeable *UUPSUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UUPSUpgradeable.Contract.UUPSUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UUPSUpgradeable *UUPSUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UUPSUpgradeable.Contract.UUPSUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UUPSUpgradeable *UUPSUpgradeableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UUPSUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UUPSUpgradeable *UUPSUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UUPSUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UUPSUpgradeable *UUPSUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UUPSUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_UUPSUpgradeable *UUPSUpgradeableCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	ret0 := new([32]byte)
	out := ret0
	err := _UUPSUpgradeable.contract.Call(opts, out, "proxiableUUID")
	return *ret0, err
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_UUPSUpgradeable *UUPSUpgradeableSession) ProxiableUUID() ([32]byte, error) {
	return _UUPSUpgradeable.Contract.ProxiableUUID(&_UUPSUpgradeable.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_UUPSUpgradeable *UUPSUpgradeableCallerSession) ProxiableUUID() ([32]byte, error) {
	return _UUPSUpgradeable.Contract.ProxiableUUID(&_UUPSUpgradeable.CallOpts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_UUPSUpgradeable *UUPSUpgradeableTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _UUPSUpgradeable.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_UUPSUpgradeable *UUPSUpgradeableSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _UUPSUpgradeable.Contract.UpgradeTo(&_UUPSUpgradeable.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_UUPSUpgradeable *UUPSUpgradeableTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _UUPSUpgradeable.Contract.UpgradeTo(&_UUPSUpgradeable.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_UUPSUpgradeable *UUPSUpgradeableTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _UUPSUpgradeable.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_UUPSUpgradeable *UUPSUpgradeableSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _UUPSUpgradeable.Contract.UpgradeToAndCall(&_UUPSUpgradeable.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_UUPSUpgradeable *UUPSUpgradeableTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _UUPSUpgradeable.Contract.UpgradeToAndCall(&_UUPSUpgradeable.TransactOpts, newImplementation, data)
}

// UUPSUpgradeableAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the UUPSUpgradeable contract.
type UUPSUpgradeableAdminChangedIterator struct {
	Event *UUPSUpgradeableAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UUPSUpgradeableAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UUPSUpgradeableAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UUPSUpgradeableAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UUPSUpgradeableAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UUPSUpgradeableAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UUPSUpgradeableAdminChanged represents a AdminChanged event raised by the UUPSUpgradeable contract.
type UUPSUpgradeableAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*UUPSUpgradeableAdminChangedIterator, error) {
	logs, sub, err := _UUPSUpgradeable.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &UUPSUpgradeableAdminChangedIterator{contract: _UUPSUpgradeable.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *UUPSUpgradeableAdminChanged) (event.Subscription, error) {
	logs, sub, err := _UUPSUpgradeable.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UUPSUpgradeableAdminChanged)
				if err := _UUPSUpgradeable.contract.UnpackLog(event, "AdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) ParseAdminChanged(log types.Log) (*UUPSUpgradeableAdminChanged, error) {
	event := new(UUPSUpgradeableAdminChanged)
	if err := _UUPSUpgradeable.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// UUPSUpgradeableBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the UUPSUpgradeable contract.
type UUPSUpgradeableBeaconUpgradedIterator struct {
	Event *UUPSUpgradeableBeaconUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UUPSUpgradeableBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UUPSUpgradeableBeaconUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UUPSUpgradeableBeaconUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UUPSUpgradeableBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UUPSUpgradeableBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UUPSUpgradeableBeaconUpgraded represents a BeaconUpgraded event raised by the UUPSUpgradeable contract.
type UUPSUpgradeableBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*UUPSUpgradeableBeaconUpgradedIterator, error) {
	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _UUPSUpgradeable.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &UUPSUpgradeableBeaconUpgradedIterator{contract: _UUPSUpgradeable.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *UUPSUpgradeableBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {
	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _UUPSUpgradeable.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UUPSUpgradeableBeaconUpgraded)
				if err := _UUPSUpgradeable.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) ParseBeaconUpgraded(log types.Log) (*UUPSUpgradeableBeaconUpgraded, error) {
	event := new(UUPSUpgradeableBeaconUpgraded)
	if err := _UUPSUpgradeable.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// UUPSUpgradeableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the UUPSUpgradeable contract.
type UUPSUpgradeableInitializedIterator struct {
	Event *UUPSUpgradeableInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UUPSUpgradeableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UUPSUpgradeableInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UUPSUpgradeableInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UUPSUpgradeableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UUPSUpgradeableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UUPSUpgradeableInitialized represents a Initialized event raised by the UUPSUpgradeable contract.
type UUPSUpgradeableInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) FilterInitialized(opts *bind.FilterOpts) (*UUPSUpgradeableInitializedIterator, error) {
	logs, sub, err := _UUPSUpgradeable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &UUPSUpgradeableInitializedIterator{contract: _UUPSUpgradeable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *UUPSUpgradeableInitialized) (event.Subscription, error) {
	logs, sub, err := _UUPSUpgradeable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UUPSUpgradeableInitialized)
				if err := _UUPSUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) ParseInitialized(log types.Log) (*UUPSUpgradeableInitialized, error) {
	event := new(UUPSUpgradeableInitialized)
	if err := _UUPSUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	return event, nil
}

// UUPSUpgradeableUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the UUPSUpgradeable contract.
type UUPSUpgradeableUpgradedIterator struct {
	Event *UUPSUpgradeableUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UUPSUpgradeableUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UUPSUpgradeableUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UUPSUpgradeableUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UUPSUpgradeableUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UUPSUpgradeableUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UUPSUpgradeableUpgraded represents a Upgraded event raised by the UUPSUpgradeable contract.
type UUPSUpgradeableUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*UUPSUpgradeableUpgradedIterator, error) {
	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _UUPSUpgradeable.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &UUPSUpgradeableUpgradedIterator{contract: _UUPSUpgradeable.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *UUPSUpgradeableUpgraded, implementation []common.Address) (event.Subscription, error) {
	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _UUPSUpgradeable.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UUPSUpgradeableUpgraded)
				if err := _UUPSUpgradeable.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) ParseUpgraded(log types.Log) (*UUPSUpgradeableUpgraded, error) {
	event := new(UUPSUpgradeableUpgraded)
	if err := _UUPSUpgradeable.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ConsoleABI is the input ABI used to generate the binding from.
const ConsoleABI = "[]"

// ConsoleBinRuntime is the compiled bytecode used for adding genesis block without deploying code.
const ConsoleBinRuntime = `73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122053ae8df758f712ef022bb04e216542c7df35cfa91695669f4e286b095ee3b32b64736f6c634300080e0033`

// ConsoleBin is the compiled bytecode used for deploying new contracts.
var ConsoleBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122053ae8df758f712ef022bb04e216542c7df35cfa91695669f4e286b095ee3b32b64736f6c634300080e0033"

// DeployConsole deploys a new Klaytn contract, binding an instance of Console to it.
func DeployConsole(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Console, error) {
	parsed, err := abi.JSON(strings.NewReader(ConsoleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ConsoleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Console{ConsoleCaller: ConsoleCaller{contract: contract}, ConsoleTransactor: ConsoleTransactor{contract: contract}, ConsoleFilterer: ConsoleFilterer{contract: contract}}, nil
}

// Console is an auto generated Go binding around a Klaytn contract.
type Console struct {
	ConsoleCaller     // Read-only binding to the contract
	ConsoleTransactor // Write-only binding to the contract
	ConsoleFilterer   // Log filterer for contract events
}

// ConsoleCaller is an auto generated read-only Go binding around a Klaytn contract.
type ConsoleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsoleTransactor is an auto generated write-only Go binding around a Klaytn contract.
type ConsoleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsoleFilterer is an auto generated log filtering Go binding around a Klaytn contract events.
type ConsoleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConsoleSession is an auto generated Go binding around a Klaytn contract,
// with pre-set call and transact options.
type ConsoleSession struct {
	Contract     *Console          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConsoleCallerSession is an auto generated read-only Go binding around a Klaytn contract,
// with pre-set call options.
type ConsoleCallerSession struct {
	Contract *ConsoleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ConsoleTransactorSession is an auto generated write-only Go binding around a Klaytn contract,
// with pre-set transact options.
type ConsoleTransactorSession struct {
	Contract     *ConsoleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ConsoleRaw is an auto generated low-level Go binding around a Klaytn contract.
type ConsoleRaw struct {
	Contract *Console // Generic contract binding to access the raw methods on
}

// ConsoleCallerRaw is an auto generated low-level read-only Go binding around a Klaytn contract.
type ConsoleCallerRaw struct {
	Contract *ConsoleCaller // Generic read-only contract binding to access the raw methods on
}

// ConsoleTransactorRaw is an auto generated low-level write-only Go binding around a Klaytn contract.
type ConsoleTransactorRaw struct {
	Contract *ConsoleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConsole creates a new instance of Console, bound to a specific deployed contract.
func NewConsole(address common.Address, backend bind.ContractBackend) (*Console, error) {
	contract, err := bindConsole(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Console{ConsoleCaller: ConsoleCaller{contract: contract}, ConsoleTransactor: ConsoleTransactor{contract: contract}, ConsoleFilterer: ConsoleFilterer{contract: contract}}, nil
}

// NewConsoleCaller creates a new read-only instance of Console, bound to a specific deployed contract.
func NewConsoleCaller(address common.Address, caller bind.ContractCaller) (*ConsoleCaller, error) {
	contract, err := bindConsole(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConsoleCaller{contract: contract}, nil
}

// NewConsoleTransactor creates a new write-only instance of Console, bound to a specific deployed contract.
func NewConsoleTransactor(address common.Address, transactor bind.ContractTransactor) (*ConsoleTransactor, error) {
	contract, err := bindConsole(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConsoleTransactor{contract: contract}, nil
}

// NewConsoleFilterer creates a new log filterer instance of Console, bound to a specific deployed contract.
func NewConsoleFilterer(address common.Address, filterer bind.ContractFilterer) (*ConsoleFilterer, error) {
	contract, err := bindConsole(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConsoleFilterer{contract: contract}, nil
}

// bindConsole binds a generic wrapper to an already deployed contract.
func bindConsole(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConsoleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Console *ConsoleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Console.Contract.ConsoleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Console *ConsoleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Console.Contract.ConsoleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Console *ConsoleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Console.Contract.ConsoleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Console *ConsoleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Console.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Console *ConsoleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Console.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Console *ConsoleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Console.Contract.contract.Transact(opts, method, params...)
}
