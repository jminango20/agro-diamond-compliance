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

// IdentityRegistryFacetMetaData contains all meta data concerning the IdentityRegistryFacet contract.
var IdentityRegistryFacetMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"batchRegisterIdentity\",\"inputs\":[{\"name\":\"wallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"identities\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"countries\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"contains\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deleteIdentity\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCountry\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getIdentity\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isVerified\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"profileId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"verified\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerIdentity\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"identity\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"country\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateCountry\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"country\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateIdentity\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"identity\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"IdentityBound\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"identity\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"country\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IdentityUnbound\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VerificationCacheInvalidated\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"profileId\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"IdentityRegistryFacet__AlreadyRegistered\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"IdentityRegistryFacet__ArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IdentityRegistryFacet__NotRegistered\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"IdentityRegistryFacet__Unauthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IdentityRegistryFacet__ZeroAddress\",\"inputs\":[]}]",
}

// IdentityRegistryFacetABI is the input ABI used to generate the binding from.
// Deprecated: Use IdentityRegistryFacetMetaData.ABI instead.
var IdentityRegistryFacetABI = IdentityRegistryFacetMetaData.ABI

// IdentityRegistryFacet is an auto generated Go binding around an Ethereum contract.
type IdentityRegistryFacet struct {
	IdentityRegistryFacetCaller     // Read-only binding to the contract
	IdentityRegistryFacetTransactor // Write-only binding to the contract
	IdentityRegistryFacetFilterer   // Log filterer for contract events
}

// IdentityRegistryFacetCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdentityRegistryFacetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityRegistryFacetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdentityRegistryFacetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityRegistryFacetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdentityRegistryFacetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityRegistryFacetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdentityRegistryFacetSession struct {
	Contract     *IdentityRegistryFacet // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IdentityRegistryFacetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdentityRegistryFacetCallerSession struct {
	Contract *IdentityRegistryFacetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// IdentityRegistryFacetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdentityRegistryFacetTransactorSession struct {
	Contract     *IdentityRegistryFacetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// IdentityRegistryFacetRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdentityRegistryFacetRaw struct {
	Contract *IdentityRegistryFacet // Generic contract binding to access the raw methods on
}

// IdentityRegistryFacetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdentityRegistryFacetCallerRaw struct {
	Contract *IdentityRegistryFacetCaller // Generic read-only contract binding to access the raw methods on
}

// IdentityRegistryFacetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdentityRegistryFacetTransactorRaw struct {
	Contract *IdentityRegistryFacetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdentityRegistryFacet creates a new instance of IdentityRegistryFacet, bound to a specific deployed contract.
func NewIdentityRegistryFacet(address common.Address, backend bind.ContractBackend) (*IdentityRegistryFacet, error) {
	contract, err := bindIdentityRegistryFacet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryFacet{IdentityRegistryFacetCaller: IdentityRegistryFacetCaller{contract: contract}, IdentityRegistryFacetTransactor: IdentityRegistryFacetTransactor{contract: contract}, IdentityRegistryFacetFilterer: IdentityRegistryFacetFilterer{contract: contract}}, nil
}

// NewIdentityRegistryFacetCaller creates a new read-only instance of IdentityRegistryFacet, bound to a specific deployed contract.
func NewIdentityRegistryFacetCaller(address common.Address, caller bind.ContractCaller) (*IdentityRegistryFacetCaller, error) {
	contract, err := bindIdentityRegistryFacet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryFacetCaller{contract: contract}, nil
}

// NewIdentityRegistryFacetTransactor creates a new write-only instance of IdentityRegistryFacet, bound to a specific deployed contract.
func NewIdentityRegistryFacetTransactor(address common.Address, transactor bind.ContractTransactor) (*IdentityRegistryFacetTransactor, error) {
	contract, err := bindIdentityRegistryFacet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryFacetTransactor{contract: contract}, nil
}

// NewIdentityRegistryFacetFilterer creates a new log filterer instance of IdentityRegistryFacet, bound to a specific deployed contract.
func NewIdentityRegistryFacetFilterer(address common.Address, filterer bind.ContractFilterer) (*IdentityRegistryFacetFilterer, error) {
	contract, err := bindIdentityRegistryFacet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryFacetFilterer{contract: contract}, nil
}

// bindIdentityRegistryFacet binds a generic wrapper to an already deployed contract.
func bindIdentityRegistryFacet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IdentityRegistryFacetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityRegistryFacet *IdentityRegistryFacetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IdentityRegistryFacet.Contract.IdentityRegistryFacetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityRegistryFacet *IdentityRegistryFacetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityRegistryFacet.Contract.IdentityRegistryFacetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityRegistryFacet *IdentityRegistryFacetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityRegistryFacet.Contract.IdentityRegistryFacetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IdentityRegistryFacet *IdentityRegistryFacetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IdentityRegistryFacet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IdentityRegistryFacet *IdentityRegistryFacetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IdentityRegistryFacet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IdentityRegistryFacet *IdentityRegistryFacetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IdentityRegistryFacet.Contract.contract.Transact(opts, method, params...)
}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address wallet) view returns(bool)
func (_IdentityRegistryFacet *IdentityRegistryFacetCaller) Contains(opts *bind.CallOpts, wallet common.Address) (bool, error) {
	var out []interface{}
	err := _IdentityRegistryFacet.contract.Call(opts, &out, "contains", wallet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address wallet) view returns(bool)
func (_IdentityRegistryFacet *IdentityRegistryFacetSession) Contains(wallet common.Address) (bool, error) {
	return _IdentityRegistryFacet.Contract.Contains(&_IdentityRegistryFacet.CallOpts, wallet)
}

// Contains is a free data retrieval call binding the contract method 0x5dbe47e8.
//
// Solidity: function contains(address wallet) view returns(bool)
func (_IdentityRegistryFacet *IdentityRegistryFacetCallerSession) Contains(wallet common.Address) (bool, error) {
	return _IdentityRegistryFacet.Contract.Contains(&_IdentityRegistryFacet.CallOpts, wallet)
}

// GetCountry is a free data retrieval call binding the contract method 0xd821f92d.
//
// Solidity: function getCountry(address wallet) view returns(uint16)
func (_IdentityRegistryFacet *IdentityRegistryFacetCaller) GetCountry(opts *bind.CallOpts, wallet common.Address) (uint16, error) {
	var out []interface{}
	err := _IdentityRegistryFacet.contract.Call(opts, &out, "getCountry", wallet)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetCountry is a free data retrieval call binding the contract method 0xd821f92d.
//
// Solidity: function getCountry(address wallet) view returns(uint16)
func (_IdentityRegistryFacet *IdentityRegistryFacetSession) GetCountry(wallet common.Address) (uint16, error) {
	return _IdentityRegistryFacet.Contract.GetCountry(&_IdentityRegistryFacet.CallOpts, wallet)
}

// GetCountry is a free data retrieval call binding the contract method 0xd821f92d.
//
// Solidity: function getCountry(address wallet) view returns(uint16)
func (_IdentityRegistryFacet *IdentityRegistryFacetCallerSession) GetCountry(wallet common.Address) (uint16, error) {
	return _IdentityRegistryFacet.Contract.GetCountry(&_IdentityRegistryFacet.CallOpts, wallet)
}

// GetIdentity is a free data retrieval call binding the contract method 0x2fea7b81.
//
// Solidity: function getIdentity(address wallet) view returns(address)
func (_IdentityRegistryFacet *IdentityRegistryFacetCaller) GetIdentity(opts *bind.CallOpts, wallet common.Address) (common.Address, error) {
	var out []interface{}
	err := _IdentityRegistryFacet.contract.Call(opts, &out, "getIdentity", wallet)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetIdentity is a free data retrieval call binding the contract method 0x2fea7b81.
//
// Solidity: function getIdentity(address wallet) view returns(address)
func (_IdentityRegistryFacet *IdentityRegistryFacetSession) GetIdentity(wallet common.Address) (common.Address, error) {
	return _IdentityRegistryFacet.Contract.GetIdentity(&_IdentityRegistryFacet.CallOpts, wallet)
}

// GetIdentity is a free data retrieval call binding the contract method 0x2fea7b81.
//
// Solidity: function getIdentity(address wallet) view returns(address)
func (_IdentityRegistryFacet *IdentityRegistryFacetCallerSession) GetIdentity(wallet common.Address) (common.Address, error) {
	return _IdentityRegistryFacet.Contract.GetIdentity(&_IdentityRegistryFacet.CallOpts, wallet)
}

// BatchRegisterIdentity is a paid mutator transaction binding the contract method 0x653dc9f1.
//
// Solidity: function batchRegisterIdentity(address[] wallets, address[] identities, uint16[] countries) returns()
func (_IdentityRegistryFacet *IdentityRegistryFacetTransactor) BatchRegisterIdentity(opts *bind.TransactOpts, wallets []common.Address, identities []common.Address, countries []uint16) (*types.Transaction, error) {
	return _IdentityRegistryFacet.contract.Transact(opts, "batchRegisterIdentity", wallets, identities, countries)
}

// BatchRegisterIdentity is a paid mutator transaction binding the contract method 0x653dc9f1.
//
// Solidity: function batchRegisterIdentity(address[] wallets, address[] identities, uint16[] countries) returns()
func (_IdentityRegistryFacet *IdentityRegistryFacetSession) BatchRegisterIdentity(wallets []common.Address, identities []common.Address, countries []uint16) (*types.Transaction, error) {
	return _IdentityRegistryFacet.Contract.BatchRegisterIdentity(&_IdentityRegistryFacet.TransactOpts, wallets, identities, countries)
}

// BatchRegisterIdentity is a paid mutator transaction binding the contract method 0x653dc9f1.
//
// Solidity: function batchRegisterIdentity(address[] wallets, address[] identities, uint16[] countries) returns()
func (_IdentityRegistryFacet *IdentityRegistryFacetTransactorSession) BatchRegisterIdentity(wallets []common.Address, identities []common.Address, countries []uint16) (*types.Transaction, error) {
	return _IdentityRegistryFacet.Contract.BatchRegisterIdentity(&_IdentityRegistryFacet.TransactOpts, wallets, identities, countries)
}

// DeleteIdentity is a paid mutator transaction binding the contract method 0xa8d29d1d.
//
// Solidity: function deleteIdentity(address wallet) returns()
func (_IdentityRegistryFacet *IdentityRegistryFacetTransactor) DeleteIdentity(opts *bind.TransactOpts, wallet common.Address) (*types.Transaction, error) {
	return _IdentityRegistryFacet.contract.Transact(opts, "deleteIdentity", wallet)
}

// DeleteIdentity is a paid mutator transaction binding the contract method 0xa8d29d1d.
//
// Solidity: function deleteIdentity(address wallet) returns()
func (_IdentityRegistryFacet *IdentityRegistryFacetSession) DeleteIdentity(wallet common.Address) (*types.Transaction, error) {
	return _IdentityRegistryFacet.Contract.DeleteIdentity(&_IdentityRegistryFacet.TransactOpts, wallet)
}

// DeleteIdentity is a paid mutator transaction binding the contract method 0xa8d29d1d.
//
// Solidity: function deleteIdentity(address wallet) returns()
func (_IdentityRegistryFacet *IdentityRegistryFacetTransactorSession) DeleteIdentity(wallet common.Address) (*types.Transaction, error) {
	return _IdentityRegistryFacet.Contract.DeleteIdentity(&_IdentityRegistryFacet.TransactOpts, wallet)
}

// IsVerified is a paid mutator transaction binding the contract method 0xd8094892.
//
// Solidity: function isVerified(address wallet, uint32 profileId) returns(bool verified)
func (_IdentityRegistryFacet *IdentityRegistryFacetTransactor) IsVerified(opts *bind.TransactOpts, wallet common.Address, profileId uint32) (*types.Transaction, error) {
	return _IdentityRegistryFacet.contract.Transact(opts, "isVerified", wallet, profileId)
}

// IsVerified is a paid mutator transaction binding the contract method 0xd8094892.
//
// Solidity: function isVerified(address wallet, uint32 profileId) returns(bool verified)
func (_IdentityRegistryFacet *IdentityRegistryFacetSession) IsVerified(wallet common.Address, profileId uint32) (*types.Transaction, error) {
	return _IdentityRegistryFacet.Contract.IsVerified(&_IdentityRegistryFacet.TransactOpts, wallet, profileId)
}

// IsVerified is a paid mutator transaction binding the contract method 0xd8094892.
//
// Solidity: function isVerified(address wallet, uint32 profileId) returns(bool verified)
func (_IdentityRegistryFacet *IdentityRegistryFacetTransactorSession) IsVerified(wallet common.Address, profileId uint32) (*types.Transaction, error) {
	return _IdentityRegistryFacet.Contract.IsVerified(&_IdentityRegistryFacet.TransactOpts, wallet, profileId)
}

// RegisterIdentity is a paid mutator transaction binding the contract method 0x454a03e0.
//
// Solidity: function registerIdentity(address wallet, address identity, uint16 country) returns()
func (_IdentityRegistryFacet *IdentityRegistryFacetTransactor) RegisterIdentity(opts *bind.TransactOpts, wallet common.Address, identity common.Address, country uint16) (*types.Transaction, error) {
	return _IdentityRegistryFacet.contract.Transact(opts, "registerIdentity", wallet, identity, country)
}

// RegisterIdentity is a paid mutator transaction binding the contract method 0x454a03e0.
//
// Solidity: function registerIdentity(address wallet, address identity, uint16 country) returns()
func (_IdentityRegistryFacet *IdentityRegistryFacetSession) RegisterIdentity(wallet common.Address, identity common.Address, country uint16) (*types.Transaction, error) {
	return _IdentityRegistryFacet.Contract.RegisterIdentity(&_IdentityRegistryFacet.TransactOpts, wallet, identity, country)
}

// RegisterIdentity is a paid mutator transaction binding the contract method 0x454a03e0.
//
// Solidity: function registerIdentity(address wallet, address identity, uint16 country) returns()
func (_IdentityRegistryFacet *IdentityRegistryFacetTransactorSession) RegisterIdentity(wallet common.Address, identity common.Address, country uint16) (*types.Transaction, error) {
	return _IdentityRegistryFacet.Contract.RegisterIdentity(&_IdentityRegistryFacet.TransactOpts, wallet, identity, country)
}

// UpdateCountry is a paid mutator transaction binding the contract method 0x3b239a7f.
//
// Solidity: function updateCountry(address wallet, uint16 country) returns()
func (_IdentityRegistryFacet *IdentityRegistryFacetTransactor) UpdateCountry(opts *bind.TransactOpts, wallet common.Address, country uint16) (*types.Transaction, error) {
	return _IdentityRegistryFacet.contract.Transact(opts, "updateCountry", wallet, country)
}

// UpdateCountry is a paid mutator transaction binding the contract method 0x3b239a7f.
//
// Solidity: function updateCountry(address wallet, uint16 country) returns()
func (_IdentityRegistryFacet *IdentityRegistryFacetSession) UpdateCountry(wallet common.Address, country uint16) (*types.Transaction, error) {
	return _IdentityRegistryFacet.Contract.UpdateCountry(&_IdentityRegistryFacet.TransactOpts, wallet, country)
}

// UpdateCountry is a paid mutator transaction binding the contract method 0x3b239a7f.
//
// Solidity: function updateCountry(address wallet, uint16 country) returns()
func (_IdentityRegistryFacet *IdentityRegistryFacetTransactorSession) UpdateCountry(wallet common.Address, country uint16) (*types.Transaction, error) {
	return _IdentityRegistryFacet.Contract.UpdateCountry(&_IdentityRegistryFacet.TransactOpts, wallet, country)
}

// UpdateIdentity is a paid mutator transaction binding the contract method 0x8e098ca1.
//
// Solidity: function updateIdentity(address wallet, address identity) returns()
func (_IdentityRegistryFacet *IdentityRegistryFacetTransactor) UpdateIdentity(opts *bind.TransactOpts, wallet common.Address, identity common.Address) (*types.Transaction, error) {
	return _IdentityRegistryFacet.contract.Transact(opts, "updateIdentity", wallet, identity)
}

// UpdateIdentity is a paid mutator transaction binding the contract method 0x8e098ca1.
//
// Solidity: function updateIdentity(address wallet, address identity) returns()
func (_IdentityRegistryFacet *IdentityRegistryFacetSession) UpdateIdentity(wallet common.Address, identity common.Address) (*types.Transaction, error) {
	return _IdentityRegistryFacet.Contract.UpdateIdentity(&_IdentityRegistryFacet.TransactOpts, wallet, identity)
}

// UpdateIdentity is a paid mutator transaction binding the contract method 0x8e098ca1.
//
// Solidity: function updateIdentity(address wallet, address identity) returns()
func (_IdentityRegistryFacet *IdentityRegistryFacetTransactorSession) UpdateIdentity(wallet common.Address, identity common.Address) (*types.Transaction, error) {
	return _IdentityRegistryFacet.Contract.UpdateIdentity(&_IdentityRegistryFacet.TransactOpts, wallet, identity)
}

// IdentityRegistryFacetIdentityBoundIterator is returned from FilterIdentityBound and is used to iterate over the raw logs and unpacked data for IdentityBound events raised by the IdentityRegistryFacet contract.
type IdentityRegistryFacetIdentityBoundIterator struct {
	Event *IdentityRegistryFacetIdentityBound // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryFacetIdentityBoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryFacetIdentityBound)
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
		it.Event = new(IdentityRegistryFacetIdentityBound)
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
func (it *IdentityRegistryFacetIdentityBoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryFacetIdentityBoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryFacetIdentityBound represents a IdentityBound event raised by the IdentityRegistryFacet contract.
type IdentityRegistryFacetIdentityBound struct {
	Wallet   common.Address
	Identity common.Address
	Country  uint16
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterIdentityBound is a free log retrieval operation binding the contract event 0x9dfdf88e26a2c4bded105443cc053a166bbd2c705cccf502c56e785a270501f4.
//
// Solidity: event IdentityBound(address indexed wallet, address indexed identity, uint16 country)
func (_IdentityRegistryFacet *IdentityRegistryFacetFilterer) FilterIdentityBound(opts *bind.FilterOpts, wallet []common.Address, identity []common.Address) (*IdentityRegistryFacetIdentityBoundIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _IdentityRegistryFacet.contract.FilterLogs(opts, "IdentityBound", walletRule, identityRule)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryFacetIdentityBoundIterator{contract: _IdentityRegistryFacet.contract, event: "IdentityBound", logs: logs, sub: sub}, nil
}

// WatchIdentityBound is a free log subscription operation binding the contract event 0x9dfdf88e26a2c4bded105443cc053a166bbd2c705cccf502c56e785a270501f4.
//
// Solidity: event IdentityBound(address indexed wallet, address indexed identity, uint16 country)
func (_IdentityRegistryFacet *IdentityRegistryFacetFilterer) WatchIdentityBound(opts *bind.WatchOpts, sink chan<- *IdentityRegistryFacetIdentityBound, wallet []common.Address, identity []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}

	logs, sub, err := _IdentityRegistryFacet.contract.WatchLogs(opts, "IdentityBound", walletRule, identityRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryFacetIdentityBound)
				if err := _IdentityRegistryFacet.contract.UnpackLog(event, "IdentityBound", log); err != nil {
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

// ParseIdentityBound is a log parse operation binding the contract event 0x9dfdf88e26a2c4bded105443cc053a166bbd2c705cccf502c56e785a270501f4.
//
// Solidity: event IdentityBound(address indexed wallet, address indexed identity, uint16 country)
func (_IdentityRegistryFacet *IdentityRegistryFacetFilterer) ParseIdentityBound(log types.Log) (*IdentityRegistryFacetIdentityBound, error) {
	event := new(IdentityRegistryFacetIdentityBound)
	if err := _IdentityRegistryFacet.contract.UnpackLog(event, "IdentityBound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityRegistryFacetIdentityUnboundIterator is returned from FilterIdentityUnbound and is used to iterate over the raw logs and unpacked data for IdentityUnbound events raised by the IdentityRegistryFacet contract.
type IdentityRegistryFacetIdentityUnboundIterator struct {
	Event *IdentityRegistryFacetIdentityUnbound // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryFacetIdentityUnboundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryFacetIdentityUnbound)
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
		it.Event = new(IdentityRegistryFacetIdentityUnbound)
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
func (it *IdentityRegistryFacetIdentityUnboundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryFacetIdentityUnboundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryFacetIdentityUnbound represents a IdentityUnbound event raised by the IdentityRegistryFacet contract.
type IdentityRegistryFacetIdentityUnbound struct {
	Wallet common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIdentityUnbound is a free log retrieval operation binding the contract event 0x60e6e5156777f14df461ae3bf407d96d1e2314215c879ffd3904740e269afebc.
//
// Solidity: event IdentityUnbound(address indexed wallet)
func (_IdentityRegistryFacet *IdentityRegistryFacetFilterer) FilterIdentityUnbound(opts *bind.FilterOpts, wallet []common.Address) (*IdentityRegistryFacetIdentityUnboundIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _IdentityRegistryFacet.contract.FilterLogs(opts, "IdentityUnbound", walletRule)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryFacetIdentityUnboundIterator{contract: _IdentityRegistryFacet.contract, event: "IdentityUnbound", logs: logs, sub: sub}, nil
}

// WatchIdentityUnbound is a free log subscription operation binding the contract event 0x60e6e5156777f14df461ae3bf407d96d1e2314215c879ffd3904740e269afebc.
//
// Solidity: event IdentityUnbound(address indexed wallet)
func (_IdentityRegistryFacet *IdentityRegistryFacetFilterer) WatchIdentityUnbound(opts *bind.WatchOpts, sink chan<- *IdentityRegistryFacetIdentityUnbound, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _IdentityRegistryFacet.contract.WatchLogs(opts, "IdentityUnbound", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryFacetIdentityUnbound)
				if err := _IdentityRegistryFacet.contract.UnpackLog(event, "IdentityUnbound", log); err != nil {
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

// ParseIdentityUnbound is a log parse operation binding the contract event 0x60e6e5156777f14df461ae3bf407d96d1e2314215c879ffd3904740e269afebc.
//
// Solidity: event IdentityUnbound(address indexed wallet)
func (_IdentityRegistryFacet *IdentityRegistryFacetFilterer) ParseIdentityUnbound(log types.Log) (*IdentityRegistryFacetIdentityUnbound, error) {
	event := new(IdentityRegistryFacetIdentityUnbound)
	if err := _IdentityRegistryFacet.contract.UnpackLog(event, "IdentityUnbound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityRegistryFacetVerificationCacheInvalidatedIterator is returned from FilterVerificationCacheInvalidated and is used to iterate over the raw logs and unpacked data for VerificationCacheInvalidated events raised by the IdentityRegistryFacet contract.
type IdentityRegistryFacetVerificationCacheInvalidatedIterator struct {
	Event *IdentityRegistryFacetVerificationCacheInvalidated // Event containing the contract specifics and raw log

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
func (it *IdentityRegistryFacetVerificationCacheInvalidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityRegistryFacetVerificationCacheInvalidated)
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
		it.Event = new(IdentityRegistryFacetVerificationCacheInvalidated)
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
func (it *IdentityRegistryFacetVerificationCacheInvalidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityRegistryFacetVerificationCacheInvalidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityRegistryFacetVerificationCacheInvalidated represents a VerificationCacheInvalidated event raised by the IdentityRegistryFacet contract.
type IdentityRegistryFacetVerificationCacheInvalidated struct {
	Wallet    common.Address
	ProfileId uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVerificationCacheInvalidated is a free log retrieval operation binding the contract event 0xce6b309b9fa71743315f14bf7e1733dc69ae99047b2497c3c07fd2a3ca46b9a1.
//
// Solidity: event VerificationCacheInvalidated(address indexed wallet, uint32 indexed profileId)
func (_IdentityRegistryFacet *IdentityRegistryFacetFilterer) FilterVerificationCacheInvalidated(opts *bind.FilterOpts, wallet []common.Address, profileId []uint32) (*IdentityRegistryFacetVerificationCacheInvalidatedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _IdentityRegistryFacet.contract.FilterLogs(opts, "VerificationCacheInvalidated", walletRule, profileIdRule)
	if err != nil {
		return nil, err
	}
	return &IdentityRegistryFacetVerificationCacheInvalidatedIterator{contract: _IdentityRegistryFacet.contract, event: "VerificationCacheInvalidated", logs: logs, sub: sub}, nil
}

// WatchVerificationCacheInvalidated is a free log subscription operation binding the contract event 0xce6b309b9fa71743315f14bf7e1733dc69ae99047b2497c3c07fd2a3ca46b9a1.
//
// Solidity: event VerificationCacheInvalidated(address indexed wallet, uint32 indexed profileId)
func (_IdentityRegistryFacet *IdentityRegistryFacetFilterer) WatchVerificationCacheInvalidated(opts *bind.WatchOpts, sink chan<- *IdentityRegistryFacetVerificationCacheInvalidated, wallet []common.Address, profileId []uint32) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}
	var profileIdRule []interface{}
	for _, profileIdItem := range profileId {
		profileIdRule = append(profileIdRule, profileIdItem)
	}

	logs, sub, err := _IdentityRegistryFacet.contract.WatchLogs(opts, "VerificationCacheInvalidated", walletRule, profileIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityRegistryFacetVerificationCacheInvalidated)
				if err := _IdentityRegistryFacet.contract.UnpackLog(event, "VerificationCacheInvalidated", log); err != nil {
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
func (_IdentityRegistryFacet *IdentityRegistryFacetFilterer) ParseVerificationCacheInvalidated(log types.Log) (*IdentityRegistryFacetVerificationCacheInvalidated, error) {
	event := new(IdentityRegistryFacetVerificationCacheInvalidated)
	if err := _IdentityRegistryFacet.contract.UnpackLog(event, "VerificationCacheInvalidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
