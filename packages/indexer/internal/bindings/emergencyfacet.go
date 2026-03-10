// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// EmergencyFacetMetaData contains all meta data concerning the EmergencyFacet contract.
var EmergencyFacetMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"emergencyPause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isEmergencyPaused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"EmergencyPause\",\"inputs\":[{\"name\":\"triggeredBy\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"EmergencyFacet__AlreadyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmergencyFacet__Unauthorized\",\"inputs\":[]}]",
}

// EmergencyFacetABI is the input ABI used to generate the binding from.
// Deprecated: Use EmergencyFacetMetaData.ABI instead.
var EmergencyFacetABI = EmergencyFacetMetaData.ABI

// EmergencyFacet is an auto generated Go binding around an Ethereum contract.
type EmergencyFacet struct {
	EmergencyFacetCaller     // Read-only binding to the contract
	EmergencyFacetTransactor // Write-only binding to the contract
	EmergencyFacetFilterer   // Log filterer for contract events
}

// EmergencyFacetCaller is an auto generated read-only Go binding around an Ethereum contract.
type EmergencyFacetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EmergencyFacetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EmergencyFacetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EmergencyFacetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EmergencyFacetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EmergencyFacetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EmergencyFacetSession struct {
	Contract     *EmergencyFacet   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EmergencyFacetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EmergencyFacetCallerSession struct {
	Contract *EmergencyFacetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// EmergencyFacetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EmergencyFacetTransactorSession struct {
	Contract     *EmergencyFacetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// EmergencyFacetRaw is an auto generated low-level Go binding around an Ethereum contract.
type EmergencyFacetRaw struct {
	Contract *EmergencyFacet // Generic contract binding to access the raw methods on
}

// EmergencyFacetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EmergencyFacetCallerRaw struct {
	Contract *EmergencyFacetCaller // Generic read-only contract binding to access the raw methods on
}

// EmergencyFacetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EmergencyFacetTransactorRaw struct {
	Contract *EmergencyFacetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEmergencyFacet creates a new instance of EmergencyFacet, bound to a specific deployed contract.
func NewEmergencyFacet(address common.Address, backend bind.ContractBackend) (*EmergencyFacet, error) {
	contract, err := bindEmergencyFacet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EmergencyFacet{EmergencyFacetCaller: EmergencyFacetCaller{contract: contract}, EmergencyFacetTransactor: EmergencyFacetTransactor{contract: contract}, EmergencyFacetFilterer: EmergencyFacetFilterer{contract: contract}}, nil
}

// NewEmergencyFacetCaller creates a new read-only instance of EmergencyFacet, bound to a specific deployed contract.
func NewEmergencyFacetCaller(address common.Address, caller bind.ContractCaller) (*EmergencyFacetCaller, error) {
	contract, err := bindEmergencyFacet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EmergencyFacetCaller{contract: contract}, nil
}

// NewEmergencyFacetTransactor creates a new write-only instance of EmergencyFacet, bound to a specific deployed contract.
func NewEmergencyFacetTransactor(address common.Address, transactor bind.ContractTransactor) (*EmergencyFacetTransactor, error) {
	contract, err := bindEmergencyFacet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EmergencyFacetTransactor{contract: contract}, nil
}

// NewEmergencyFacetFilterer creates a new log filterer instance of EmergencyFacet, bound to a specific deployed contract.
func NewEmergencyFacetFilterer(address common.Address, filterer bind.ContractFilterer) (*EmergencyFacetFilterer, error) {
	contract, err := bindEmergencyFacet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EmergencyFacetFilterer{contract: contract}, nil
}

// bindEmergencyFacet binds a generic wrapper to an already deployed contract.
func bindEmergencyFacet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EmergencyFacetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EmergencyFacet *EmergencyFacetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EmergencyFacet.Contract.EmergencyFacetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EmergencyFacet *EmergencyFacetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EmergencyFacet.Contract.EmergencyFacetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EmergencyFacet *EmergencyFacetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EmergencyFacet.Contract.EmergencyFacetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EmergencyFacet *EmergencyFacetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EmergencyFacet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EmergencyFacet *EmergencyFacetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EmergencyFacet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EmergencyFacet *EmergencyFacetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EmergencyFacet.Contract.contract.Transact(opts, method, params...)
}

// IsEmergencyPaused is a free data retrieval call binding the contract method 0x290d10c4.
//
// Solidity: function isEmergencyPaused() view returns(bool)
func (_EmergencyFacet *EmergencyFacetCaller) IsEmergencyPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EmergencyFacet.contract.Call(opts, &out, "isEmergencyPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEmergencyPaused is a free data retrieval call binding the contract method 0x290d10c4.
//
// Solidity: function isEmergencyPaused() view returns(bool)
func (_EmergencyFacet *EmergencyFacetSession) IsEmergencyPaused() (bool, error) {
	return _EmergencyFacet.Contract.IsEmergencyPaused(&_EmergencyFacet.CallOpts)
}

// IsEmergencyPaused is a free data retrieval call binding the contract method 0x290d10c4.
//
// Solidity: function isEmergencyPaused() view returns(bool)
func (_EmergencyFacet *EmergencyFacetCallerSession) IsEmergencyPaused() (bool, error) {
	return _EmergencyFacet.Contract.IsEmergencyPaused(&_EmergencyFacet.CallOpts)
}

// EmergencyPause is a paid mutator transaction binding the contract method 0x51858e27.
//
// Solidity: function emergencyPause() returns()
func (_EmergencyFacet *EmergencyFacetTransactor) EmergencyPause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EmergencyFacet.contract.Transact(opts, "emergencyPause")
}

// EmergencyPause is a paid mutator transaction binding the contract method 0x51858e27.
//
// Solidity: function emergencyPause() returns()
func (_EmergencyFacet *EmergencyFacetSession) EmergencyPause() (*types.Transaction, error) {
	return _EmergencyFacet.Contract.EmergencyPause(&_EmergencyFacet.TransactOpts)
}

// EmergencyPause is a paid mutator transaction binding the contract method 0x51858e27.
//
// Solidity: function emergencyPause() returns()
func (_EmergencyFacet *EmergencyFacetTransactorSession) EmergencyPause() (*types.Transaction, error) {
	return _EmergencyFacet.Contract.EmergencyPause(&_EmergencyFacet.TransactOpts)
}

// EmergencyFacetEmergencyPauseIterator is returned from FilterEmergencyPause and is used to iterate over the raw logs and unpacked data for EmergencyPause events raised by the EmergencyFacet contract.
type EmergencyFacetEmergencyPauseIterator struct {
	Event *EmergencyFacetEmergencyPause // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EmergencyFacetEmergencyPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmergencyFacetEmergencyPause)
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
		it.Event = new(EmergencyFacetEmergencyPause)
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
func (it *EmergencyFacetEmergencyPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmergencyFacetEmergencyPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmergencyFacetEmergencyPause represents a EmergencyPause event raised by the EmergencyFacet contract.
type EmergencyFacetEmergencyPause struct {
	TriggeredBy common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEmergencyPause is a free log retrieval operation binding the contract event 0x7c83004a7e59a8ea03b200186c4dda29a4e144d9844d63dbc1a09acf7dfcd485.
//
// Solidity: event EmergencyPause(address indexed triggeredBy)
func (_EmergencyFacet *EmergencyFacetFilterer) FilterEmergencyPause(opts *bind.FilterOpts, triggeredBy []common.Address) (*EmergencyFacetEmergencyPauseIterator, error) {

	var triggeredByRule []interface{}
	for _, triggeredByItem := range triggeredBy {
		triggeredByRule = append(triggeredByRule, triggeredByItem)
	}

	logs, sub, err := _EmergencyFacet.contract.FilterLogs(opts, "EmergencyPause", triggeredByRule)
	if err != nil {
		return nil, err
	}
	return &EmergencyFacetEmergencyPauseIterator{contract: _EmergencyFacet.contract, event: "EmergencyPause", logs: logs, sub: sub}, nil
}

// WatchEmergencyPause is a free log subscription operation binding the contract event 0x7c83004a7e59a8ea03b200186c4dda29a4e144d9844d63dbc1a09acf7dfcd485.
//
// Solidity: event EmergencyPause(address indexed triggeredBy)
func (_EmergencyFacet *EmergencyFacetFilterer) WatchEmergencyPause(opts *bind.WatchOpts, sink chan<- *EmergencyFacetEmergencyPause, triggeredBy []common.Address) (event.Subscription, error) {

	var triggeredByRule []interface{}
	for _, triggeredByItem := range triggeredBy {
		triggeredByRule = append(triggeredByRule, triggeredByItem)
	}

	logs, sub, err := _EmergencyFacet.contract.WatchLogs(opts, "EmergencyPause", triggeredByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmergencyFacetEmergencyPause)
				if err := _EmergencyFacet.contract.UnpackLog(event, "EmergencyPause", log); err != nil {
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

// ParseEmergencyPause is a log parse operation binding the contract event 0x7c83004a7e59a8ea03b200186c4dda29a4e144d9844d63dbc1a09acf7dfcd485.
//
// Solidity: event EmergencyPause(address indexed triggeredBy)
func (_EmergencyFacet *EmergencyFacetFilterer) ParseEmergencyPause(log types.Log) (*EmergencyFacetEmergencyPause, error) {
	event := new(EmergencyFacetEmergencyPause)
	if err := _EmergencyFacet.contract.UnpackLog(event, "EmergencyPause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
