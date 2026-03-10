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

// ClaimTopicsFacetMetaData contains all meta data concerning the ClaimTopicsFacet contract.
var ClaimTopicsFacetMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"createProfile\",\"inputs\":[{\"name\":\"claimTopics\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"profileId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getProfileClaimTopics\",\"inputs\":[{\"name\":\"profileId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getProfileVersion\",\"inputs\":[{\"name\":\"profileId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"profileExists\",\"inputs\":[{\"name\":\"profileId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setProfileClaimTopics\",\"inputs\":[{\"name\":\"profileId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"claimTopics\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ProfileUpdated\",\"inputs\":[{\"name\":\"profileId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ClaimTopicsFacet__EmptyClaimTopics\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ClaimTopicsFacet__ProfileNotFound\",\"inputs\":[{\"name\":\"profileId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"type\":\"error\",\"name\":\"ClaimTopicsFacet__Unauthorized\",\"inputs\":[]}]",
}

// ClaimTopicsFacetABI is the input ABI used to generate the binding from.
// Deprecated: Use ClaimTopicsFacetMetaData.ABI instead.
var ClaimTopicsFacetABI = ClaimTopicsFacetMetaData.ABI

// ClaimTopicsFacet is an auto generated Go binding around an Ethereum contract.
type ClaimTopicsFacet struct {
	ClaimTopicsFacetCaller     // Read-only binding to the contract
	ClaimTopicsFacetTransactor // Write-only binding to the contract
	ClaimTopicsFacetFilterer   // Log filterer for contract events
}

// ClaimTopicsFacetCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClaimTopicsFacetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimTopicsFacetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClaimTopicsFacetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimTopicsFacetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClaimTopicsFacetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimTopicsFacetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClaimTopicsFacetSession struct {
	Contract     *ClaimTopicsFacet // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClaimTopicsFacetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClaimTopicsFacetCallerSession struct {
	Contract *ClaimTopicsFacetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ClaimTopicsFacetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClaimTopicsFacetTransactorSession struct {
	Contract     *ClaimTopicsFacetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ClaimTopicsFacetRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClaimTopicsFacetRaw struct {
	Contract *ClaimTopicsFacet // Generic contract binding to access the raw methods on
}

// ClaimTopicsFacetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClaimTopicsFacetCallerRaw struct {
	Contract *ClaimTopicsFacetCaller // Generic read-only contract binding to access the raw methods on
}

// ClaimTopicsFacetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClaimTopicsFacetTransactorRaw struct {
	Contract *ClaimTopicsFacetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClaimTopicsFacet creates a new instance of ClaimTopicsFacet, bound to a specific deployed contract.
func NewClaimTopicsFacet(address common.Address, backend bind.ContractBackend) (*ClaimTopicsFacet, error) {
	contract, err := bindClaimTopicsFacet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ClaimTopicsFacet{ClaimTopicsFacetCaller: ClaimTopicsFacetCaller{contract: contract}, ClaimTopicsFacetTransactor: ClaimTopicsFacetTransactor{contract: contract}, ClaimTopicsFacetFilterer: ClaimTopicsFacetFilterer{contract: contract}}, nil
}

// NewClaimTopicsFacetCaller creates a new read-only instance of ClaimTopicsFacet, bound to a specific deployed contract.
func NewClaimTopicsFacetCaller(address common.Address, caller bind.ContractCaller) (*ClaimTopicsFacetCaller, error) {
	contract, err := bindClaimTopicsFacet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimTopicsFacetCaller{contract: contract}, nil
}

// NewClaimTopicsFacetTransactor creates a new write-only instance of ClaimTopicsFacet, bound to a specific deployed contract.
func NewClaimTopicsFacetTransactor(address common.Address, transactor bind.ContractTransactor) (*ClaimTopicsFacetTransactor, error) {
	contract, err := bindClaimTopicsFacet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimTopicsFacetTransactor{contract: contract}, nil
}

// NewClaimTopicsFacetFilterer creates a new log filterer instance of ClaimTopicsFacet, bound to a specific deployed contract.
func NewClaimTopicsFacetFilterer(address common.Address, filterer bind.ContractFilterer) (*ClaimTopicsFacetFilterer, error) {
	contract, err := bindClaimTopicsFacet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClaimTopicsFacetFilterer{contract: contract}, nil
}

// bindClaimTopicsFacet binds a generic wrapper to an already deployed contract.
func bindClaimTopicsFacet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ClaimTopicsFacetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClaimTopicsFacet *ClaimTopicsFacetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClaimTopicsFacet.Contract.ClaimTopicsFacetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClaimTopicsFacet *ClaimTopicsFacetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimTopicsFacet.Contract.ClaimTopicsFacetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClaimTopicsFacet *ClaimTopicsFacetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClaimTopicsFacet.Contract.ClaimTopicsFacetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClaimTopicsFacet *ClaimTopicsFacetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClaimTopicsFacet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClaimTopicsFacet *ClaimTopicsFacetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimTopicsFacet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClaimTopicsFacet *ClaimTopicsFacetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClaimTopicsFacet.Contract.contract.Transact(opts, method, params...)
}

// GetProfileClaimTopics is a free data retrieval call binding the contract method 0xfe8be7c3.
//
// Solidity: function getProfileClaimTopics(uint32 profileId) view returns(uint256[])
func (_ClaimTopicsFacet *ClaimTopicsFacetCaller) GetProfileClaimTopics(opts *bind.CallOpts, profileId uint32) ([]*big.Int, error) {
	var out []interface{}
	err := _ClaimTopicsFacet.contract.Call(opts, &out, "getProfileClaimTopics", profileId)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetProfileClaimTopics is a free data retrieval call binding the contract method 0xfe8be7c3.
//
// Solidity: function getProfileClaimTopics(uint32 profileId) view returns(uint256[])
func (_ClaimTopicsFacet *ClaimTopicsFacetSession) GetProfileClaimTopics(profileId uint32) ([]*big.Int, error) {
	return _ClaimTopicsFacet.Contract.GetProfileClaimTopics(&_ClaimTopicsFacet.CallOpts, profileId)
}

// GetProfileClaimTopics is a free data retrieval call binding the contract method 0xfe8be7c3.
//
// Solidity: function getProfileClaimTopics(uint32 profileId) view returns(uint256[])
func (_ClaimTopicsFacet *ClaimTopicsFacetCallerSession) GetProfileClaimTopics(profileId uint32) ([]*big.Int, error) {
	return _ClaimTopicsFacet.Contract.GetProfileClaimTopics(&_ClaimTopicsFacet.CallOpts, profileId)
}

// GetProfileVersion is a free data retrieval call binding the contract method 0xc0917f73.
//
// Solidity: function getProfileVersion(uint32 profileId) view returns(uint64)
func (_ClaimTopicsFacet *ClaimTopicsFacetCaller) GetProfileVersion(opts *bind.CallOpts, profileId uint32) (uint64, error) {
	var out []interface{}
	err := _ClaimTopicsFacet.contract.Call(opts, &out, "getProfileVersion", profileId)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetProfileVersion is a free data retrieval call binding the contract method 0xc0917f73.
//
// Solidity: function getProfileVersion(uint32 profileId) view returns(uint64)
func (_ClaimTopicsFacet *ClaimTopicsFacetSession) GetProfileVersion(profileId uint32) (uint64, error) {
	return _ClaimTopicsFacet.Contract.GetProfileVersion(&_ClaimTopicsFacet.CallOpts, profileId)
}

// GetProfileVersion is a free data retrieval call binding the contract method 0xc0917f73.
//
// Solidity: function getProfileVersion(uint32 profileId) view returns(uint64)
func (_ClaimTopicsFacet *ClaimTopicsFacetCallerSession) GetProfileVersion(profileId uint32) (uint64, error) {
	return _ClaimTopicsFacet.Contract.GetProfileVersion(&_ClaimTopicsFacet.CallOpts, profileId)
}

// ProfileExists is a free data retrieval call binding the contract method 0x1cc410fd.
//
// Solidity: function profileExists(uint32 profileId) view returns(bool)
func (_ClaimTopicsFacet *ClaimTopicsFacetCaller) ProfileExists(opts *bind.CallOpts, profileId uint32) (bool, error) {
	var out []interface{}
	err := _ClaimTopicsFacet.contract.Call(opts, &out, "profileExists", profileId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProfileExists is a free data retrieval call binding the contract method 0x1cc410fd.
//
// Solidity: function profileExists(uint32 profileId) view returns(bool)
func (_ClaimTopicsFacet *ClaimTopicsFacetSession) ProfileExists(profileId uint32) (bool, error) {
	return _ClaimTopicsFacet.Contract.ProfileExists(&_ClaimTopicsFacet.CallOpts, profileId)
}

// ProfileExists is a free data retrieval call binding the contract method 0x1cc410fd.
//
// Solidity: function profileExists(uint32 profileId) view returns(bool)
func (_ClaimTopicsFacet *ClaimTopicsFacetCallerSession) ProfileExists(profileId uint32) (bool, error) {
	return _ClaimTopicsFacet.Contract.ProfileExists(&_ClaimTopicsFacet.CallOpts, profileId)
}

// CreateProfile is a paid mutator transaction binding the contract method 0xf2a6e5fd.
//
// Solidity: function createProfile(uint256[] claimTopics) returns(uint32 profileId)
func (_ClaimTopicsFacet *ClaimTopicsFacetTransactor) CreateProfile(opts *bind.TransactOpts, claimTopics []*big.Int) (*types.Transaction, error) {
	return _ClaimTopicsFacet.contract.Transact(opts, "createProfile", claimTopics)
}

// CreateProfile is a paid mutator transaction binding the contract method 0xf2a6e5fd.
//
// Solidity: function createProfile(uint256[] claimTopics) returns(uint32 profileId)
func (_ClaimTopicsFacet *ClaimTopicsFacetSession) CreateProfile(claimTopics []*big.Int) (*types.Transaction, error) {
	return _ClaimTopicsFacet.Contract.CreateProfile(&_ClaimTopicsFacet.TransactOpts, claimTopics)
}

// CreateProfile is a paid mutator transaction binding the contract method 0xf2a6e5fd.
//
// Solidity: function createProfile(uint256[] claimTopics) returns(uint32 profileId)
func (_ClaimTopicsFacet *ClaimTopicsFacetTransactorSession) CreateProfile(claimTopics []*big.Int) (*types.Transaction, error) {
	return _ClaimTopicsFacet.Contract.CreateProfile(&_ClaimTopicsFacet.TransactOpts, claimTopics)
}

// SetProfileClaimTopics is a paid mutator transaction binding the contract method 0xd7f11a62.
//
// Solidity: function setProfileClaimTopics(uint32 profileId, uint256[] claimTopics) returns()
func (_ClaimTopicsFacet *ClaimTopicsFacetTransactor) SetProfileClaimTopics(opts *bind.TransactOpts, profileId uint32, claimTopics []*big.Int) (*types.Transaction, error) {
	return _ClaimTopicsFacet.contract.Transact(opts, "setProfileClaimTopics", profileId, claimTopics)
}

// SetProfileClaimTopics is a paid mutator transaction binding the contract method 0xd7f11a62.
//
// Solidity: function setProfileClaimTopics(uint32 profileId, uint256[] claimTopics) returns()
func (_ClaimTopicsFacet *ClaimTopicsFacetSession) SetProfileClaimTopics(profileId uint32, claimTopics []*big.Int) (*types.Transaction, error) {
	return _ClaimTopicsFacet.Contract.SetProfileClaimTopics(&_ClaimTopicsFacet.TransactOpts, profileId, claimTopics)
}

// SetProfileClaimTopics is a paid mutator transaction binding the contract method 0xd7f11a62.
//
// Solidity: function setProfileClaimTopics(uint32 profileId, uint256[] claimTopics) returns()
func (_ClaimTopicsFacet *ClaimTopicsFacetTransactorSession) SetProfileClaimTopics(profileId uint32, claimTopics []*big.Int) (*types.Transaction, error) {
	return _ClaimTopicsFacet.Contract.SetProfileClaimTopics(&_ClaimTopicsFacet.TransactOpts, profileId, claimTopics)
}

// ClaimTopicsFacetProfileUpdatedIterator is returned from FilterProfileUpdated and is used to iterate over the raw logs and unpacked data for ProfileUpdated events raised by the ClaimTopicsFacet contract.
type ClaimTopicsFacetProfileUpdatedIterator struct {
	Event *ClaimTopicsFacetProfileUpdated // Event containing the contract specifics and raw log

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
func (it *ClaimTopicsFacetProfileUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimTopicsFacetProfileUpdated)
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
		it.Event = new(ClaimTopicsFacetProfileUpdated)
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
func (it *ClaimTopicsFacetProfileUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimTopicsFacetProfileUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimTopicsFacetProfileUpdated represents a ProfileUpdated event raised by the ClaimTopicsFacet contract.
type ClaimTopicsFacetProfileUpdated struct {
	ProfileId uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProfileUpdated is a free log retrieval operation binding the contract event 0x2403aa6a3f2c5d774a0a87e24356fea79ae169a2821dfe1e6a48c8642834f482.
//
// Solidity: event ProfileUpdated(uint32 indexed profileId)
func (_ClaimTopicsFacet *ClaimTopicsFacetFilterer) FilterProfileUpdated(opts *bind.FilterOpts, profileId []uint32) (*ClaimTopicsFacetProfileUpdatedIterator, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _ClaimTopicsFacet.contract.FilterLogs(opts, "ProfileUpdated", profileIdRule)
	if err != nil {
		return nil, err
	}
	return &ClaimTopicsFacetProfileUpdatedIterator{contract: _ClaimTopicsFacet.contract, event: "ProfileUpdated", logs: logs, sub: sub}, nil
}

// WatchProfileUpdated is a free log subscription operation binding the contract event 0x2403aa6a3f2c5d774a0a87e24356fea79ae169a2821dfe1e6a48c8642834f482.
//
// Solidity: event ProfileUpdated(uint32 indexed profileId)
func (_ClaimTopicsFacet *ClaimTopicsFacetFilterer) WatchProfileUpdated(opts *bind.WatchOpts, sink chan<- *ClaimTopicsFacetProfileUpdated, profileId []uint32) (event.Subscription, error) {

	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _ClaimTopicsFacet.contract.WatchLogs(opts, "ProfileUpdated", profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimTopicsFacetProfileUpdated)
				if err := _ClaimTopicsFacet.contract.UnpackLog(event, "ProfileUpdated", log); err != nil {
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

// ParseProfileUpdated is a log parse operation binding the contract event 0x2403aa6a3f2c5d774a0a87e24356fea79ae169a2821dfe1e6a48c8642834f482.
//
// Solidity: event ProfileUpdated(uint32 indexed profileId)
func (_ClaimTopicsFacet *ClaimTopicsFacetFilterer) ParseProfileUpdated(log types.Log) (*ClaimTopicsFacetProfileUpdated, error) {
	event := new(ClaimTopicsFacetProfileUpdated)
	if err := _ClaimTopicsFacet.contract.UnpackLog(event, "ProfileUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
