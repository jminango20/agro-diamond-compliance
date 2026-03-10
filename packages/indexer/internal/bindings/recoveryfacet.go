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

// RecoveryFacetMetaData contains all meta data concerning the RecoveryFacet contract.
var RecoveryFacetMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"recoverWallet\",\"inputs\":[{\"name\":\"lostWallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newWallet\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"TokensRecovered\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"lostWallet\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newWallet\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"free\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"locked\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"custody\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"pendingSettlement\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WalletRecovered\",\"inputs\":[{\"name\":\"lostWallet\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newWallet\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"agent\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"RecoveryFacet__NewWalletAlreadyRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RecoveryFacet__SameAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RecoveryFacet__Unauthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RecoveryFacet__ZeroAddress\",\"inputs\":[]}]",
}

// RecoveryFacetABI is the input ABI used to generate the binding from.
// Deprecated: Use RecoveryFacetMetaData.ABI instead.
var RecoveryFacetABI = RecoveryFacetMetaData.ABI

// RecoveryFacet is an auto generated Go binding around an Ethereum contract.
type RecoveryFacet struct {
	RecoveryFacetCaller     // Read-only binding to the contract
	RecoveryFacetTransactor // Write-only binding to the contract
	RecoveryFacetFilterer   // Log filterer for contract events
}

// RecoveryFacetCaller is an auto generated read-only Go binding around an Ethereum contract.
type RecoveryFacetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecoveryFacetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RecoveryFacetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecoveryFacetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RecoveryFacetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecoveryFacetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RecoveryFacetSession struct {
	Contract     *RecoveryFacet    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RecoveryFacetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RecoveryFacetCallerSession struct {
	Contract *RecoveryFacetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// RecoveryFacetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RecoveryFacetTransactorSession struct {
	Contract     *RecoveryFacetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// RecoveryFacetRaw is an auto generated low-level Go binding around an Ethereum contract.
type RecoveryFacetRaw struct {
	Contract *RecoveryFacet // Generic contract binding to access the raw methods on
}

// RecoveryFacetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RecoveryFacetCallerRaw struct {
	Contract *RecoveryFacetCaller // Generic read-only contract binding to access the raw methods on
}

// RecoveryFacetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RecoveryFacetTransactorRaw struct {
	Contract *RecoveryFacetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRecoveryFacet creates a new instance of RecoveryFacet, bound to a specific deployed contract.
func NewRecoveryFacet(address common.Address, backend bind.ContractBackend) (*RecoveryFacet, error) {
	contract, err := bindRecoveryFacet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RecoveryFacet{RecoveryFacetCaller: RecoveryFacetCaller{contract: contract}, RecoveryFacetTransactor: RecoveryFacetTransactor{contract: contract}, RecoveryFacetFilterer: RecoveryFacetFilterer{contract: contract}}, nil
}

// NewRecoveryFacetCaller creates a new read-only instance of RecoveryFacet, bound to a specific deployed contract.
func NewRecoveryFacetCaller(address common.Address, caller bind.ContractCaller) (*RecoveryFacetCaller, error) {
	contract, err := bindRecoveryFacet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RecoveryFacetCaller{contract: contract}, nil
}

// NewRecoveryFacetTransactor creates a new write-only instance of RecoveryFacet, bound to a specific deployed contract.
func NewRecoveryFacetTransactor(address common.Address, transactor bind.ContractTransactor) (*RecoveryFacetTransactor, error) {
	contract, err := bindRecoveryFacet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RecoveryFacetTransactor{contract: contract}, nil
}

// NewRecoveryFacetFilterer creates a new log filterer instance of RecoveryFacet, bound to a specific deployed contract.
func NewRecoveryFacetFilterer(address common.Address, filterer bind.ContractFilterer) (*RecoveryFacetFilterer, error) {
	contract, err := bindRecoveryFacet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RecoveryFacetFilterer{contract: contract}, nil
}

// bindRecoveryFacet binds a generic wrapper to an already deployed contract.
func bindRecoveryFacet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RecoveryFacetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RecoveryFacet *RecoveryFacetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RecoveryFacet.Contract.RecoveryFacetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RecoveryFacet *RecoveryFacetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RecoveryFacet.Contract.RecoveryFacetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RecoveryFacet *RecoveryFacetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RecoveryFacet.Contract.RecoveryFacetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RecoveryFacet *RecoveryFacetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RecoveryFacet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RecoveryFacet *RecoveryFacetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RecoveryFacet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RecoveryFacet *RecoveryFacetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RecoveryFacet.Contract.contract.Transact(opts, method, params...)
}

// RecoverWallet is a paid mutator transaction binding the contract method 0x78228c1f.
//
// Solidity: function recoverWallet(address lostWallet, address newWallet) returns()
func (_RecoveryFacet *RecoveryFacetTransactor) RecoverWallet(opts *bind.TransactOpts, lostWallet common.Address, newWallet common.Address) (*types.Transaction, error) {
	return _RecoveryFacet.contract.Transact(opts, "recoverWallet", lostWallet, newWallet)
}

// RecoverWallet is a paid mutator transaction binding the contract method 0x78228c1f.
//
// Solidity: function recoverWallet(address lostWallet, address newWallet) returns()
func (_RecoveryFacet *RecoveryFacetSession) RecoverWallet(lostWallet common.Address, newWallet common.Address) (*types.Transaction, error) {
	return _RecoveryFacet.Contract.RecoverWallet(&_RecoveryFacet.TransactOpts, lostWallet, newWallet)
}

// RecoverWallet is a paid mutator transaction binding the contract method 0x78228c1f.
//
// Solidity: function recoverWallet(address lostWallet, address newWallet) returns()
func (_RecoveryFacet *RecoveryFacetTransactorSession) RecoverWallet(lostWallet common.Address, newWallet common.Address) (*types.Transaction, error) {
	return _RecoveryFacet.Contract.RecoverWallet(&_RecoveryFacet.TransactOpts, lostWallet, newWallet)
}

// RecoveryFacetTokensRecoveredIterator is returned from FilterTokensRecovered and is used to iterate over the raw logs and unpacked data for TokensRecovered events raised by the RecoveryFacet contract.
type RecoveryFacetTokensRecoveredIterator struct {
	Event *RecoveryFacetTokensRecovered // Event containing the contract specifics and raw log

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
func (it *RecoveryFacetTokensRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RecoveryFacetTokensRecovered)
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
		it.Event = new(RecoveryFacetTokensRecovered)
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
func (it *RecoveryFacetTokensRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RecoveryFacetTokensRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RecoveryFacetTokensRecovered represents a TokensRecovered event raised by the RecoveryFacet contract.
type RecoveryFacetTokensRecovered struct {
	TokenId           *big.Int
	LostWallet        common.Address
	NewWallet         common.Address
	Free              *big.Int
	Locked            *big.Int
	Custody           *big.Int
	PendingSettlement *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTokensRecovered is a free log retrieval operation binding the contract event 0xf9ad4a15e0a86ed17820bafbd9fe303b805c379a638639b48839e5fbb856314f.
//
// Solidity: event TokensRecovered(uint256 indexed tokenId, address indexed lostWallet, address indexed newWallet, uint256 free, uint256 locked, uint256 custody, uint256 pendingSettlement)
func (_RecoveryFacet *RecoveryFacetFilterer) FilterTokensRecovered(opts *bind.FilterOpts, tokenId []*big.Int, lostWallet []common.Address, newWallet []common.Address) (*RecoveryFacetTokensRecoveredIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var lostWalletRule []interface{}
	for _, lostWalletItem := range lostWallet {
		lostWalletRule = append(lostWalletRule, lostWalletItem)
	}
	var newWalletRule []interface{}
	for _, newWalletItem := range newWallet {
		newWalletRule = append(newWalletRule, newWalletItem)
	}

	logs, sub, err := _RecoveryFacet.contract.FilterLogs(opts, "TokensRecovered", tokenIdRule, lostWalletRule, newWalletRule)
	if err != nil {
		return nil, err
	}
	return &RecoveryFacetTokensRecoveredIterator{contract: _RecoveryFacet.contract, event: "TokensRecovered", logs: logs, sub: sub}, nil
}

// WatchTokensRecovered is a free log subscription operation binding the contract event 0xf9ad4a15e0a86ed17820bafbd9fe303b805c379a638639b48839e5fbb856314f.
//
// Solidity: event TokensRecovered(uint256 indexed tokenId, address indexed lostWallet, address indexed newWallet, uint256 free, uint256 locked, uint256 custody, uint256 pendingSettlement)
func (_RecoveryFacet *RecoveryFacetFilterer) WatchTokensRecovered(opts *bind.WatchOpts, sink chan<- *RecoveryFacetTokensRecovered, tokenId []*big.Int, lostWallet []common.Address, newWallet []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var lostWalletRule []interface{}
	for _, lostWalletItem := range lostWallet {
		lostWalletRule = append(lostWalletRule, lostWalletItem)
	}
	var newWalletRule []interface{}
	for _, newWalletItem := range newWallet {
		newWalletRule = append(newWalletRule, newWalletItem)
	}

	logs, sub, err := _RecoveryFacet.contract.WatchLogs(opts, "TokensRecovered", tokenIdRule, lostWalletRule, newWalletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RecoveryFacetTokensRecovered)
				if err := _RecoveryFacet.contract.UnpackLog(event, "TokensRecovered", log); err != nil {
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

// ParseTokensRecovered is a log parse operation binding the contract event 0xf9ad4a15e0a86ed17820bafbd9fe303b805c379a638639b48839e5fbb856314f.
//
// Solidity: event TokensRecovered(uint256 indexed tokenId, address indexed lostWallet, address indexed newWallet, uint256 free, uint256 locked, uint256 custody, uint256 pendingSettlement)
func (_RecoveryFacet *RecoveryFacetFilterer) ParseTokensRecovered(log types.Log) (*RecoveryFacetTokensRecovered, error) {
	event := new(RecoveryFacetTokensRecovered)
	if err := _RecoveryFacet.contract.UnpackLog(event, "TokensRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RecoveryFacetWalletRecoveredIterator is returned from FilterWalletRecovered and is used to iterate over the raw logs and unpacked data for WalletRecovered events raised by the RecoveryFacet contract.
type RecoveryFacetWalletRecoveredIterator struct {
	Event *RecoveryFacetWalletRecovered // Event containing the contract specifics and raw log

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
func (it *RecoveryFacetWalletRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RecoveryFacetWalletRecovered)
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
		it.Event = new(RecoveryFacetWalletRecovered)
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
func (it *RecoveryFacetWalletRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RecoveryFacetWalletRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RecoveryFacetWalletRecovered represents a WalletRecovered event raised by the RecoveryFacet contract.
type RecoveryFacetWalletRecovered struct {
	LostWallet common.Address
	NewWallet  common.Address
	Agent      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWalletRecovered is a free log retrieval operation binding the contract event 0x45f7fbac61d7ff60d8103f50b3c66a224166a4f20d04ec58b7e5316bfbd0d3d5.
//
// Solidity: event WalletRecovered(address indexed lostWallet, address indexed newWallet, address indexed agent)
func (_RecoveryFacet *RecoveryFacetFilterer) FilterWalletRecovered(opts *bind.FilterOpts, lostWallet []common.Address, newWallet []common.Address, agent []common.Address) (*RecoveryFacetWalletRecoveredIterator, error) {

	var lostWalletRule []interface{}
	for _, lostWalletItem := range lostWallet {
		lostWalletRule = append(lostWalletRule, lostWalletItem)
	}
	var newWalletRule []interface{}
	for _, newWalletItem := range newWallet {
		newWalletRule = append(newWalletRule, newWalletItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _RecoveryFacet.contract.FilterLogs(opts, "WalletRecovered", lostWalletRule, newWalletRule, agentRule)
	if err != nil {
		return nil, err
	}
	return &RecoveryFacetWalletRecoveredIterator{contract: _RecoveryFacet.contract, event: "WalletRecovered", logs: logs, sub: sub}, nil
}

// WatchWalletRecovered is a free log subscription operation binding the contract event 0x45f7fbac61d7ff60d8103f50b3c66a224166a4f20d04ec58b7e5316bfbd0d3d5.
//
// Solidity: event WalletRecovered(address indexed lostWallet, address indexed newWallet, address indexed agent)
func (_RecoveryFacet *RecoveryFacetFilterer) WatchWalletRecovered(opts *bind.WatchOpts, sink chan<- *RecoveryFacetWalletRecovered, lostWallet []common.Address, newWallet []common.Address, agent []common.Address) (event.Subscription, error) {

	var lostWalletRule []interface{}
	for _, lostWalletItem := range lostWallet {
		lostWalletRule = append(lostWalletRule, lostWalletItem)
	}
	var newWalletRule []interface{}
	for _, newWalletItem := range newWallet {
		newWalletRule = append(newWalletRule, newWalletItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _RecoveryFacet.contract.WatchLogs(opts, "WalletRecovered", lostWalletRule, newWalletRule, agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RecoveryFacetWalletRecovered)
				if err := _RecoveryFacet.contract.UnpackLog(event, "WalletRecovered", log); err != nil {
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

// ParseWalletRecovered is a log parse operation binding the contract event 0x45f7fbac61d7ff60d8103f50b3c66a224166a4f20d04ec58b7e5316bfbd0d3d5.
//
// Solidity: event WalletRecovered(address indexed lostWallet, address indexed newWallet, address indexed agent)
func (_RecoveryFacet *RecoveryFacetFilterer) ParseWalletRecovered(log types.Log) (*RecoveryFacetWalletRecovered, error) {
	event := new(RecoveryFacetWalletRecovered)
	if err := _RecoveryFacet.contract.UnpackLog(event, "WalletRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
