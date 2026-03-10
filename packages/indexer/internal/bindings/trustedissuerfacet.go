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

// TrustedIssuerFacetMetaData contains all meta data concerning the TrustedIssuerFacet contract.
var TrustedIssuerFacetMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"addTrustedIssuer\",\"inputs\":[{\"name\":\"profileId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isTrustedIssuer\",\"inputs\":[{\"name\":\"profileId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeTrustedIssuer\",\"inputs\":[{\"name\":\"profileId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"TrustedIssuerAdded\",\"inputs\":[{\"name\":\"profileId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"issuer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TrustedIssuerRemoved\",\"inputs\":[{\"name\":\"profileId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"issuer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VerificationCacheInvalidated\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"profileId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"TrustedIssuerFacet__AlreadyTrusted\",\"inputs\":[{\"name\":\"profileId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TrustedIssuerFacet__NotTrusted\",\"inputs\":[{\"name\":\"profileId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"issuer\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TrustedIssuerFacet__ProfileNotFound\",\"inputs\":[{\"name\":\"profileId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"TrustedIssuerFacet__Unauthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TrustedIssuerFacet__ZeroAddress\",\"inputs\":[]}]",
}

// TrustedIssuerFacetABI is the input ABI used to generate the binding from.
// Deprecated: Use TrustedIssuerFacetMetaData.ABI instead.
var TrustedIssuerFacetABI = TrustedIssuerFacetMetaData.ABI

// TrustedIssuerFacet is an auto generated Go binding around an Ethereum contract.
type TrustedIssuerFacet struct {
	TrustedIssuerFacetCaller     // Read-only binding to the contract
	TrustedIssuerFacetTransactor // Write-only binding to the contract
	TrustedIssuerFacetFilterer   // Log filterer for contract events
}

// TrustedIssuerFacetCaller is an auto generated read-only Go binding around an Ethereum contract.
type TrustedIssuerFacetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrustedIssuerFacetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TrustedIssuerFacetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrustedIssuerFacetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TrustedIssuerFacetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrustedIssuerFacetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TrustedIssuerFacetSession struct {
	Contract     *TrustedIssuerFacet // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TrustedIssuerFacetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TrustedIssuerFacetCallerSession struct {
	Contract *TrustedIssuerFacetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// TrustedIssuerFacetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TrustedIssuerFacetTransactorSession struct {
	Contract     *TrustedIssuerFacetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// TrustedIssuerFacetRaw is an auto generated low-level Go binding around an Ethereum contract.
type TrustedIssuerFacetRaw struct {
	Contract *TrustedIssuerFacet // Generic contract binding to access the raw methods on
}

// TrustedIssuerFacetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TrustedIssuerFacetCallerRaw struct {
	Contract *TrustedIssuerFacetCaller // Generic read-only contract binding to access the raw methods on
}

// TrustedIssuerFacetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TrustedIssuerFacetTransactorRaw struct {
	Contract *TrustedIssuerFacetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrustedIssuerFacet creates a new instance of TrustedIssuerFacet, bound to a specific deployed contract.
func NewTrustedIssuerFacet(address common.Address, backend bind.ContractBackend) (*TrustedIssuerFacet, error) {
	contract, err := bindTrustedIssuerFacet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TrustedIssuerFacet{TrustedIssuerFacetCaller: TrustedIssuerFacetCaller{contract: contract}, TrustedIssuerFacetTransactor: TrustedIssuerFacetTransactor{contract: contract}, TrustedIssuerFacetFilterer: TrustedIssuerFacetFilterer{contract: contract}}, nil
}

// NewTrustedIssuerFacetCaller creates a new read-only instance of TrustedIssuerFacet, bound to a specific deployed contract.
func NewTrustedIssuerFacetCaller(address common.Address, caller bind.ContractCaller) (*TrustedIssuerFacetCaller, error) {
	contract, err := bindTrustedIssuerFacet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TrustedIssuerFacetCaller{contract: contract}, nil
}

// NewTrustedIssuerFacetTransactor creates a new write-only instance of TrustedIssuerFacet, bound to a specific deployed contract.
func NewTrustedIssuerFacetTransactor(address common.Address, transactor bind.ContractTransactor) (*TrustedIssuerFacetTransactor, error) {
	contract, err := bindTrustedIssuerFacet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TrustedIssuerFacetTransactor{contract: contract}, nil
}

// NewTrustedIssuerFacetFilterer creates a new log filterer instance of TrustedIssuerFacet, bound to a specific deployed contract.
func NewTrustedIssuerFacetFilterer(address common.Address, filterer bind.ContractFilterer) (*TrustedIssuerFacetFilterer, error) {
	contract, err := bindTrustedIssuerFacet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TrustedIssuerFacetFilterer{contract: contract}, nil
}

// bindTrustedIssuerFacet binds a generic wrapper to an already deployed contract.
func bindTrustedIssuerFacet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TrustedIssuerFacetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TrustedIssuerFacet *TrustedIssuerFacetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TrustedIssuerFacet.Contract.TrustedIssuerFacetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TrustedIssuerFacet *TrustedIssuerFacetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrustedIssuerFacet.Contract.TrustedIssuerFacetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TrustedIssuerFacet *TrustedIssuerFacetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TrustedIssuerFacet.Contract.TrustedIssuerFacetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TrustedIssuerFacet *TrustedIssuerFacetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TrustedIssuerFacet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TrustedIssuerFacet *TrustedIssuerFacetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrustedIssuerFacet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TrustedIssuerFacet *TrustedIssuerFacetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TrustedIssuerFacet.Contract.contract.Transact(opts, method, params...)
}

// IsTrustedIssuer is a free data retrieval call binding the contract method 0xa0567d37.
//
// Solidity: function isTrustedIssuer(uint32 profileId, address issuer) view returns(bool)
func (_TrustedIssuerFacet *TrustedIssuerFacetCaller) IsTrustedIssuer(opts *bind.CallOpts, profileId uint32, issuer common.Address) (bool, error) {
	var out []interface{}
	err := _TrustedIssuerFacet.contract.Call(opts, &out, "isTrustedIssuer", profileId, issuer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedIssuer is a free data retrieval call binding the contract method 0xa0567d37.
//
// Solidity: function isTrustedIssuer(uint32 profileId, address issuer) view returns(bool)
func (_TrustedIssuerFacet *TrustedIssuerFacetSession) IsTrustedIssuer(profileId uint32, issuer common.Address) (bool, error) {
	return _TrustedIssuerFacet.Contract.IsTrustedIssuer(&_TrustedIssuerFacet.CallOpts, profileId, issuer)
}

// IsTrustedIssuer is a free data retrieval call binding the contract method 0xa0567d37.
//
// Solidity: function isTrustedIssuer(uint32 profileId, address issuer) view returns(bool)
func (_TrustedIssuerFacet *TrustedIssuerFacetCallerSession) IsTrustedIssuer(profileId uint32, issuer common.Address) (bool, error) {
	return _TrustedIssuerFacet.Contract.IsTrustedIssuer(&_TrustedIssuerFacet.CallOpts, profileId, issuer)
}

// AddTrustedIssuer is a paid mutator transaction binding the contract method 0xca84950e.
//
// Solidity: function addTrustedIssuer(uint32 profileId, address issuer) returns()
func (_TrustedIssuerFacet *TrustedIssuerFacetTransactor) AddTrustedIssuer(opts *bind.TransactOpts, profileId uint32, issuer common.Address) (*types.Transaction, error) {
	return _TrustedIssuerFacet.contract.Transact(opts, "addTrustedIssuer", profileId, issuer)
}

// AddTrustedIssuer is a paid mutator transaction binding the contract method 0xca84950e.
//
// Solidity: function addTrustedIssuer(uint32 profileId, address issuer) returns()
func (_TrustedIssuerFacet *TrustedIssuerFacetSession) AddTrustedIssuer(profileId uint32, issuer common.Address) (*types.Transaction, error) {
	return _TrustedIssuerFacet.Contract.AddTrustedIssuer(&_TrustedIssuerFacet.TransactOpts, profileId, issuer)
}

// AddTrustedIssuer is a paid mutator transaction binding the contract method 0xca84950e.
//
// Solidity: function addTrustedIssuer(uint32 profileId, address issuer) returns()
func (_TrustedIssuerFacet *TrustedIssuerFacetTransactorSession) AddTrustedIssuer(profileId uint32, issuer common.Address) (*types.Transaction, error) {
	return _TrustedIssuerFacet.Contract.AddTrustedIssuer(&_TrustedIssuerFacet.TransactOpts, profileId, issuer)
}

// RemoveTrustedIssuer is a paid mutator transaction binding the contract method 0xedaa5e1d.
//
// Solidity: function removeTrustedIssuer(uint32 profileId, address issuer) returns()
func (_TrustedIssuerFacet *TrustedIssuerFacetTransactor) RemoveTrustedIssuer(opts *bind.TransactOpts, profileId uint32, issuer common.Address) (*types.Transaction, error) {
	return _TrustedIssuerFacet.contract.Transact(opts, "removeTrustedIssuer", profileId, issuer)
}

// RemoveTrustedIssuer is a paid mutator transaction binding the contract method 0xedaa5e1d.
//
// Solidity: function removeTrustedIssuer(uint32 profileId, address issuer) returns()
func (_TrustedIssuerFacet *TrustedIssuerFacetSession) RemoveTrustedIssuer(profileId uint32, issuer common.Address) (*types.Transaction, error) {
	return _TrustedIssuerFacet.Contract.RemoveTrustedIssuer(&_TrustedIssuerFacet.TransactOpts, profileId, issuer)
}

// RemoveTrustedIssuer is a paid mutator transaction binding the contract method 0xedaa5e1d.
//
// Solidity: function removeTrustedIssuer(uint32 profileId, address issuer) returns()
func (_TrustedIssuerFacet *TrustedIssuerFacetTransactorSession) RemoveTrustedIssuer(profileId uint32, issuer common.Address) (*types.Transaction, error) {
	return _TrustedIssuerFacet.Contract.RemoveTrustedIssuer(&_TrustedIssuerFacet.TransactOpts, profileId, issuer)
}

// TrustedIssuerFacetTrustedIssuerAddedIterator is returned from FilterTrustedIssuerAdded and is used to iterate over the raw logs and unpacked data for TrustedIssuerAdded events raised by the TrustedIssuerFacet contract.
type TrustedIssuerFacetTrustedIssuerAddedIterator struct {
	Event *TrustedIssuerFacetTrustedIssuerAdded // Event containing the contract specifics and raw log

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
func (it *TrustedIssuerFacetTrustedIssuerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustedIssuerFacetTrustedIssuerAdded)
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
		it.Event = new(TrustedIssuerFacetTrustedIssuerAdded)
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
func (it *TrustedIssuerFacetTrustedIssuerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustedIssuerFacetTrustedIssuerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustedIssuerFacetTrustedIssuerAdded represents a TrustedIssuerAdded event raised by the TrustedIssuerFacet contract.
type TrustedIssuerFacetTrustedIssuerAdded struct {
	ProfileId uint32
	Issuer    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTrustedIssuerAdded is a free log retrieval operation binding the contract event 0xc63754e1356043a8c7c940e381e868fd59d965e8fa2a4cc61d945437a002e316.
//
// Solidity: event TrustedIssuerAdded(uint32 indexed profileId, address indexed issuer)
func (_TrustedIssuerFacet *TrustedIssuerFacetFilterer) FilterTrustedIssuerAdded(opts *bind.FilterOpts, profileId []uint32, issuer []common.Address) (*TrustedIssuerFacetTrustedIssuerAddedIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _TrustedIssuerFacet.contract.FilterLogs(opts, "TrustedIssuerAdded", profileIdRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &TrustedIssuerFacetTrustedIssuerAddedIterator{contract: _TrustedIssuerFacet.contract, event: "TrustedIssuerAdded", logs: logs, sub: sub}, nil
}

// WatchTrustedIssuerAdded is a free log subscription operation binding the contract event 0xc63754e1356043a8c7c940e381e868fd59d965e8fa2a4cc61d945437a002e316.
//
// Solidity: event TrustedIssuerAdded(uint32 indexed profileId, address indexed issuer)
func (_TrustedIssuerFacet *TrustedIssuerFacetFilterer) WatchTrustedIssuerAdded(opts *bind.WatchOpts, sink chan<- *TrustedIssuerFacetTrustedIssuerAdded, profileId []uint32, issuer []common.Address) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _TrustedIssuerFacet.contract.WatchLogs(opts, "TrustedIssuerAdded", profileIdRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustedIssuerFacetTrustedIssuerAdded)
				if err := _TrustedIssuerFacet.contract.UnpackLog(event, "TrustedIssuerAdded", log); err != nil {
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

// ParseTrustedIssuerAdded is a log parse operation binding the contract event 0xc63754e1356043a8c7c940e381e868fd59d965e8fa2a4cc61d945437a002e316.
//
// Solidity: event TrustedIssuerAdded(uint32 indexed profileId, address indexed issuer)
func (_TrustedIssuerFacet *TrustedIssuerFacetFilterer) ParseTrustedIssuerAdded(log types.Log) (*TrustedIssuerFacetTrustedIssuerAdded, error) {
	event := new(TrustedIssuerFacetTrustedIssuerAdded)
	if err := _TrustedIssuerFacet.contract.UnpackLog(event, "TrustedIssuerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrustedIssuerFacetTrustedIssuerRemovedIterator is returned from FilterTrustedIssuerRemoved and is used to iterate over the raw logs and unpacked data for TrustedIssuerRemoved events raised by the TrustedIssuerFacet contract.
type TrustedIssuerFacetTrustedIssuerRemovedIterator struct {
	Event *TrustedIssuerFacetTrustedIssuerRemoved // Event containing the contract specifics and raw log

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
func (it *TrustedIssuerFacetTrustedIssuerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustedIssuerFacetTrustedIssuerRemoved)
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
		it.Event = new(TrustedIssuerFacetTrustedIssuerRemoved)
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
func (it *TrustedIssuerFacetTrustedIssuerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustedIssuerFacetTrustedIssuerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustedIssuerFacetTrustedIssuerRemoved represents a TrustedIssuerRemoved event raised by the TrustedIssuerFacet contract.
type TrustedIssuerFacetTrustedIssuerRemoved struct {
	ProfileId uint32
	Issuer    common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTrustedIssuerRemoved is a free log retrieval operation binding the contract event 0x4a1cd5800972055cf83566a08f8ad366bd66268bcf0dbecd5feb3dc4fdfab4ff.
//
// Solidity: event TrustedIssuerRemoved(uint32 indexed profileId, address indexed issuer)
func (_TrustedIssuerFacet *TrustedIssuerFacetFilterer) FilterTrustedIssuerRemoved(opts *bind.FilterOpts, profileId []uint32, issuer []common.Address) (*TrustedIssuerFacetTrustedIssuerRemovedIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _TrustedIssuerFacet.contract.FilterLogs(opts, "TrustedIssuerRemoved", profileIdRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &TrustedIssuerFacetTrustedIssuerRemovedIterator{contract: _TrustedIssuerFacet.contract, event: "TrustedIssuerRemoved", logs: logs, sub: sub}, nil
}

// WatchTrustedIssuerRemoved is a free log subscription operation binding the contract event 0x4a1cd5800972055cf83566a08f8ad366bd66268bcf0dbecd5feb3dc4fdfab4ff.
//
// Solidity: event TrustedIssuerRemoved(uint32 indexed profileId, address indexed issuer)
func (_TrustedIssuerFacet *TrustedIssuerFacetFilterer) WatchTrustedIssuerRemoved(opts *bind.WatchOpts, sink chan<- *TrustedIssuerFacetTrustedIssuerRemoved, profileId []uint32, issuer []common.Address) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _TrustedIssuerFacet.contract.WatchLogs(opts, "TrustedIssuerRemoved", profileIdRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustedIssuerFacetTrustedIssuerRemoved)
				if err := _TrustedIssuerFacet.contract.UnpackLog(event, "TrustedIssuerRemoved", log); err != nil {
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

// ParseTrustedIssuerRemoved is a log parse operation binding the contract event 0x4a1cd5800972055cf83566a08f8ad366bd66268bcf0dbecd5feb3dc4fdfab4ff.
//
// Solidity: event TrustedIssuerRemoved(uint32 indexed profileId, address indexed issuer)
func (_TrustedIssuerFacet *TrustedIssuerFacetFilterer) ParseTrustedIssuerRemoved(log types.Log) (*TrustedIssuerFacetTrustedIssuerRemoved, error) {
	event := new(TrustedIssuerFacetTrustedIssuerRemoved)
	if err := _TrustedIssuerFacet.contract.UnpackLog(event, "TrustedIssuerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrustedIssuerFacetVerificationCacheInvalidatedIterator is returned from FilterVerificationCacheInvalidated and is used to iterate over the raw logs and unpacked data for VerificationCacheInvalidated events raised by the TrustedIssuerFacet contract.
type TrustedIssuerFacetVerificationCacheInvalidatedIterator struct {
	Event *TrustedIssuerFacetVerificationCacheInvalidated // Event containing the contract specifics and raw log

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
func (it *TrustedIssuerFacetVerificationCacheInvalidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustedIssuerFacetVerificationCacheInvalidated)
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
		it.Event = new(TrustedIssuerFacetVerificationCacheInvalidated)
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
func (it *TrustedIssuerFacetVerificationCacheInvalidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustedIssuerFacetVerificationCacheInvalidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustedIssuerFacetVerificationCacheInvalidated represents a VerificationCacheInvalidated event raised by the TrustedIssuerFacet contract.
type TrustedIssuerFacetVerificationCacheInvalidated struct {
	Wallet    common.Address
	ProfileId uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVerificationCacheInvalidated is a free log retrieval operation binding the contract event 0xce6b309b9fa71743315f14bf7e1733dc69ae99047b2497c3c07fd2a3ca46b9a1.
//
// Solidity: event VerificationCacheInvalidated(address indexed wallet, uint32 indexed profileId)
func (_TrustedIssuerFacet *TrustedIssuerFacetFilterer) FilterVerificationCacheInvalidated(opts *bind.FilterOpts, wallet []common.Address, profileId []uint32) (*TrustedIssuerFacetVerificationCacheInvalidatedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _TrustedIssuerFacet.contract.FilterLogs(opts, "VerificationCacheInvalidated", walletRule, profileIdRule)
	if err != nil {
		return nil, err
	}
	return &TrustedIssuerFacetVerificationCacheInvalidatedIterator{contract: _TrustedIssuerFacet.contract, event: "VerificationCacheInvalidated", logs: logs, sub: sub}, nil
}

// WatchVerificationCacheInvalidated is a free log subscription operation binding the contract event 0xce6b309b9fa71743315f14bf7e1733dc69ae99047b2497c3c07fd2a3ca46b9a1.
//
// Solidity: event VerificationCacheInvalidated(address indexed wallet, uint32 indexed profileId)
func (_TrustedIssuerFacet *TrustedIssuerFacetFilterer) WatchVerificationCacheInvalidated(opts *bind.WatchOpts, sink chan<- *TrustedIssuerFacetVerificationCacheInvalidated, wallet []common.Address, profileId []uint32) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _TrustedIssuerFacet.contract.WatchLogs(opts, "VerificationCacheInvalidated", walletRule, profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustedIssuerFacetVerificationCacheInvalidated)
				if err := _TrustedIssuerFacet.contract.UnpackLog(event, "VerificationCacheInvalidated", log); err != nil {
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

// ParseVerificationCacheInvalidated is a log parse operation binding the contract event 0xce6b309b9fa71743315f14bf7e1733dc69ae99047b2497c3c07fd2a3ca46b9a1.
//
// Solidity: event VerificationCacheInvalidated(address indexed wallet, uint32 indexed profileId)
func (_TrustedIssuerFacet *TrustedIssuerFacetFilterer) ParseVerificationCacheInvalidated(log types.Log) (*TrustedIssuerFacetVerificationCacheInvalidated, error) {
	event := new(TrustedIssuerFacetVerificationCacheInvalidated)
	if err := _TrustedIssuerFacet.contract.UnpackLog(event, "VerificationCacheInvalidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
