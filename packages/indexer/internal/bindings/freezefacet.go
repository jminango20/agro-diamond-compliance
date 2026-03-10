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

// FreezeFacetMetaData contains all meta data concerning the FreezeFacet contract.
var FreezeFacetMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getFrozenAmount\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLockupExpiry\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAssetWalletFrozen\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isWalletFrozen\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setAssetWalletFrozen\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"frozen\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFrozenAmount\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setLockupExpiry\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"expiry\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWalletFrozen\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"frozen\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AssetFrozen\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"wallet\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"frozen\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockupSet\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"wallet\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"expiry\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PartialFreeze\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"wallet\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WalletFrozen\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"frozen\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"FreezeFacet__Unauthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FreezeFacet__ZeroAddress\",\"inputs\":[]}]",
}

// FreezeFacetABI is the input ABI used to generate the binding from.
// Deprecated: Use FreezeFacetMetaData.ABI instead.
var FreezeFacetABI = FreezeFacetMetaData.ABI

// FreezeFacet is an auto generated Go binding around an Ethereum contract.
type FreezeFacet struct {
	FreezeFacetCaller     // Read-only binding to the contract
	FreezeFacetTransactor // Write-only binding to the contract
	FreezeFacetFilterer   // Log filterer for contract events
}

// FreezeFacetCaller is an auto generated read-only Go binding around an Ethereum contract.
type FreezeFacetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FreezeFacetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FreezeFacetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FreezeFacetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FreezeFacetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FreezeFacetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FreezeFacetSession struct {
	Contract     *FreezeFacet      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FreezeFacetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FreezeFacetCallerSession struct {
	Contract *FreezeFacetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// FreezeFacetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FreezeFacetTransactorSession struct {
	Contract     *FreezeFacetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// FreezeFacetRaw is an auto generated low-level Go binding around an Ethereum contract.
type FreezeFacetRaw struct {
	Contract *FreezeFacet // Generic contract binding to access the raw methods on
}

// FreezeFacetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FreezeFacetCallerRaw struct {
	Contract *FreezeFacetCaller // Generic read-only contract binding to access the raw methods on
}

// FreezeFacetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FreezeFacetTransactorRaw struct {
	Contract *FreezeFacetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFreezeFacet creates a new instance of FreezeFacet, bound to a specific deployed contract.
func NewFreezeFacet(address common.Address, backend bind.ContractBackend) (*FreezeFacet, error) {
	contract, err := bindFreezeFacet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FreezeFacet{FreezeFacetCaller: FreezeFacetCaller{contract: contract}, FreezeFacetTransactor: FreezeFacetTransactor{contract: contract}, FreezeFacetFilterer: FreezeFacetFilterer{contract: contract}}, nil
}

// NewFreezeFacetCaller creates a new read-only instance of FreezeFacet, bound to a specific deployed contract.
func NewFreezeFacetCaller(address common.Address, caller bind.ContractCaller) (*FreezeFacetCaller, error) {
	contract, err := bindFreezeFacet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FreezeFacetCaller{contract: contract}, nil
}

// NewFreezeFacetTransactor creates a new write-only instance of FreezeFacet, bound to a specific deployed contract.
func NewFreezeFacetTransactor(address common.Address, transactor bind.ContractTransactor) (*FreezeFacetTransactor, error) {
	contract, err := bindFreezeFacet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FreezeFacetTransactor{contract: contract}, nil
}

// NewFreezeFacetFilterer creates a new log filterer instance of FreezeFacet, bound to a specific deployed contract.
func NewFreezeFacetFilterer(address common.Address, filterer bind.ContractFilterer) (*FreezeFacetFilterer, error) {
	contract, err := bindFreezeFacet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FreezeFacetFilterer{contract: contract}, nil
}

// bindFreezeFacet binds a generic wrapper to an already deployed contract.
func bindFreezeFacet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FreezeFacetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FreezeFacet *FreezeFacetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FreezeFacet.Contract.FreezeFacetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FreezeFacet *FreezeFacetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreezeFacet.Contract.FreezeFacetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FreezeFacet *FreezeFacetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FreezeFacet.Contract.FreezeFacetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FreezeFacet *FreezeFacetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FreezeFacet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FreezeFacet *FreezeFacetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreezeFacet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FreezeFacet *FreezeFacetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FreezeFacet.Contract.contract.Transact(opts, method, params...)
}

// GetFrozenAmount is a free data retrieval call binding the contract method 0x3143fcd4.
//
// Solidity: function getFrozenAmount(uint256 tokenId, address wallet) view returns(uint256)
func (_FreezeFacet *FreezeFacetCaller) GetFrozenAmount(opts *bind.CallOpts, tokenId *big.Int, wallet common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FreezeFacet.contract.Call(opts, &out, "getFrozenAmount", tokenId, wallet)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFrozenAmount is a free data retrieval call binding the contract method 0x3143fcd4.
//
// Solidity: function getFrozenAmount(uint256 tokenId, address wallet) view returns(uint256)
func (_FreezeFacet *FreezeFacetSession) GetFrozenAmount(tokenId *big.Int, wallet common.Address) (*big.Int, error) {
	return _FreezeFacet.Contract.GetFrozenAmount(&_FreezeFacet.CallOpts, tokenId, wallet)
}

// GetFrozenAmount is a free data retrieval call binding the contract method 0x3143fcd4.
//
// Solidity: function getFrozenAmount(uint256 tokenId, address wallet) view returns(uint256)
func (_FreezeFacet *FreezeFacetCallerSession) GetFrozenAmount(tokenId *big.Int, wallet common.Address) (*big.Int, error) {
	return _FreezeFacet.Contract.GetFrozenAmount(&_FreezeFacet.CallOpts, tokenId, wallet)
}

// GetLockupExpiry is a free data retrieval call binding the contract method 0x512d2f40.
//
// Solidity: function getLockupExpiry(uint256 tokenId, address wallet) view returns(uint64)
func (_FreezeFacet *FreezeFacetCaller) GetLockupExpiry(opts *bind.CallOpts, tokenId *big.Int, wallet common.Address) (uint64, error) {
	var out []interface{}
	err := _FreezeFacet.contract.Call(opts, &out, "getLockupExpiry", tokenId, wallet)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetLockupExpiry is a free data retrieval call binding the contract method 0x512d2f40.
//
// Solidity: function getLockupExpiry(uint256 tokenId, address wallet) view returns(uint64)
func (_FreezeFacet *FreezeFacetSession) GetLockupExpiry(tokenId *big.Int, wallet common.Address) (uint64, error) {
	return _FreezeFacet.Contract.GetLockupExpiry(&_FreezeFacet.CallOpts, tokenId, wallet)
}

// GetLockupExpiry is a free data retrieval call binding the contract method 0x512d2f40.
//
// Solidity: function getLockupExpiry(uint256 tokenId, address wallet) view returns(uint64)
func (_FreezeFacet *FreezeFacetCallerSession) GetLockupExpiry(tokenId *big.Int, wallet common.Address) (uint64, error) {
	return _FreezeFacet.Contract.GetLockupExpiry(&_FreezeFacet.CallOpts, tokenId, wallet)
}

// IsAssetWalletFrozen is a free data retrieval call binding the contract method 0x9dba96d5.
//
// Solidity: function isAssetWalletFrozen(uint256 tokenId, address wallet) view returns(bool)
func (_FreezeFacet *FreezeFacetCaller) IsAssetWalletFrozen(opts *bind.CallOpts, tokenId *big.Int, wallet common.Address) (bool, error) {
	var out []interface{}
	err := _FreezeFacet.contract.Call(opts, &out, "isAssetWalletFrozen", tokenId, wallet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAssetWalletFrozen is a free data retrieval call binding the contract method 0x9dba96d5.
//
// Solidity: function isAssetWalletFrozen(uint256 tokenId, address wallet) view returns(bool)
func (_FreezeFacet *FreezeFacetSession) IsAssetWalletFrozen(tokenId *big.Int, wallet common.Address) (bool, error) {
	return _FreezeFacet.Contract.IsAssetWalletFrozen(&_FreezeFacet.CallOpts, tokenId, wallet)
}

// IsAssetWalletFrozen is a free data retrieval call binding the contract method 0x9dba96d5.
//
// Solidity: function isAssetWalletFrozen(uint256 tokenId, address wallet) view returns(bool)
func (_FreezeFacet *FreezeFacetCallerSession) IsAssetWalletFrozen(tokenId *big.Int, wallet common.Address) (bool, error) {
	return _FreezeFacet.Contract.IsAssetWalletFrozen(&_FreezeFacet.CallOpts, tokenId, wallet)
}

// IsWalletFrozen is a free data retrieval call binding the contract method 0x95583b07.
//
// Solidity: function isWalletFrozen(address wallet) view returns(bool)
func (_FreezeFacet *FreezeFacetCaller) IsWalletFrozen(opts *bind.CallOpts, wallet common.Address) (bool, error) {
	var out []interface{}
	err := _FreezeFacet.contract.Call(opts, &out, "isWalletFrozen", wallet)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWalletFrozen is a free data retrieval call binding the contract method 0x95583b07.
//
// Solidity: function isWalletFrozen(address wallet) view returns(bool)
func (_FreezeFacet *FreezeFacetSession) IsWalletFrozen(wallet common.Address) (bool, error) {
	return _FreezeFacet.Contract.IsWalletFrozen(&_FreezeFacet.CallOpts, wallet)
}

// IsWalletFrozen is a free data retrieval call binding the contract method 0x95583b07.
//
// Solidity: function isWalletFrozen(address wallet) view returns(bool)
func (_FreezeFacet *FreezeFacetCallerSession) IsWalletFrozen(wallet common.Address) (bool, error) {
	return _FreezeFacet.Contract.IsWalletFrozen(&_FreezeFacet.CallOpts, wallet)
}

// SetAssetWalletFrozen is a paid mutator transaction binding the contract method 0xd25e9ce4.
//
// Solidity: function setAssetWalletFrozen(uint256 tokenId, address wallet, bool frozen) returns()
func (_FreezeFacet *FreezeFacetTransactor) SetAssetWalletFrozen(opts *bind.TransactOpts, tokenId *big.Int, wallet common.Address, frozen bool) (*types.Transaction, error) {
	return _FreezeFacet.contract.Transact(opts, "setAssetWalletFrozen", tokenId, wallet, frozen)
}

// SetAssetWalletFrozen is a paid mutator transaction binding the contract method 0xd25e9ce4.
//
// Solidity: function setAssetWalletFrozen(uint256 tokenId, address wallet, bool frozen) returns()
func (_FreezeFacet *FreezeFacetSession) SetAssetWalletFrozen(tokenId *big.Int, wallet common.Address, frozen bool) (*types.Transaction, error) {
	return _FreezeFacet.Contract.SetAssetWalletFrozen(&_FreezeFacet.TransactOpts, tokenId, wallet, frozen)
}

// SetAssetWalletFrozen is a paid mutator transaction binding the contract method 0xd25e9ce4.
//
// Solidity: function setAssetWalletFrozen(uint256 tokenId, address wallet, bool frozen) returns()
func (_FreezeFacet *FreezeFacetTransactorSession) SetAssetWalletFrozen(tokenId *big.Int, wallet common.Address, frozen bool) (*types.Transaction, error) {
	return _FreezeFacet.Contract.SetAssetWalletFrozen(&_FreezeFacet.TransactOpts, tokenId, wallet, frozen)
}

// SetFrozenAmount is a paid mutator transaction binding the contract method 0xca386caf.
//
// Solidity: function setFrozenAmount(uint256 tokenId, address wallet, uint256 amount) returns()
func (_FreezeFacet *FreezeFacetTransactor) SetFrozenAmount(opts *bind.TransactOpts, tokenId *big.Int, wallet common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FreezeFacet.contract.Transact(opts, "setFrozenAmount", tokenId, wallet, amount)
}

// SetFrozenAmount is a paid mutator transaction binding the contract method 0xca386caf.
//
// Solidity: function setFrozenAmount(uint256 tokenId, address wallet, uint256 amount) returns()
func (_FreezeFacet *FreezeFacetSession) SetFrozenAmount(tokenId *big.Int, wallet common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FreezeFacet.Contract.SetFrozenAmount(&_FreezeFacet.TransactOpts, tokenId, wallet, amount)
}

// SetFrozenAmount is a paid mutator transaction binding the contract method 0xca386caf.
//
// Solidity: function setFrozenAmount(uint256 tokenId, address wallet, uint256 amount) returns()
func (_FreezeFacet *FreezeFacetTransactorSession) SetFrozenAmount(tokenId *big.Int, wallet common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FreezeFacet.Contract.SetFrozenAmount(&_FreezeFacet.TransactOpts, tokenId, wallet, amount)
}

// SetLockupExpiry is a paid mutator transaction binding the contract method 0x3ee59787.
//
// Solidity: function setLockupExpiry(uint256 tokenId, address wallet, uint64 expiry) returns()
func (_FreezeFacet *FreezeFacetTransactor) SetLockupExpiry(opts *bind.TransactOpts, tokenId *big.Int, wallet common.Address, expiry uint64) (*types.Transaction, error) {
	return _FreezeFacet.contract.Transact(opts, "setLockupExpiry", tokenId, wallet, expiry)
}

// SetLockupExpiry is a paid mutator transaction binding the contract method 0x3ee59787.
//
// Solidity: function setLockupExpiry(uint256 tokenId, address wallet, uint64 expiry) returns()
func (_FreezeFacet *FreezeFacetSession) SetLockupExpiry(tokenId *big.Int, wallet common.Address, expiry uint64) (*types.Transaction, error) {
	return _FreezeFacet.Contract.SetLockupExpiry(&_FreezeFacet.TransactOpts, tokenId, wallet, expiry)
}

// SetLockupExpiry is a paid mutator transaction binding the contract method 0x3ee59787.
//
// Solidity: function setLockupExpiry(uint256 tokenId, address wallet, uint64 expiry) returns()
func (_FreezeFacet *FreezeFacetTransactorSession) SetLockupExpiry(tokenId *big.Int, wallet common.Address, expiry uint64) (*types.Transaction, error) {
	return _FreezeFacet.Contract.SetLockupExpiry(&_FreezeFacet.TransactOpts, tokenId, wallet, expiry)
}

// SetWalletFrozen is a paid mutator transaction binding the contract method 0xf57cff1b.
//
// Solidity: function setWalletFrozen(address wallet, bool frozen) returns()
func (_FreezeFacet *FreezeFacetTransactor) SetWalletFrozen(opts *bind.TransactOpts, wallet common.Address, frozen bool) (*types.Transaction, error) {
	return _FreezeFacet.contract.Transact(opts, "setWalletFrozen", wallet, frozen)
}

// SetWalletFrozen is a paid mutator transaction binding the contract method 0xf57cff1b.
//
// Solidity: function setWalletFrozen(address wallet, bool frozen) returns()
func (_FreezeFacet *FreezeFacetSession) SetWalletFrozen(wallet common.Address, frozen bool) (*types.Transaction, error) {
	return _FreezeFacet.Contract.SetWalletFrozen(&_FreezeFacet.TransactOpts, wallet, frozen)
}

// SetWalletFrozen is a paid mutator transaction binding the contract method 0xf57cff1b.
//
// Solidity: function setWalletFrozen(address wallet, bool frozen) returns()
func (_FreezeFacet *FreezeFacetTransactorSession) SetWalletFrozen(wallet common.Address, frozen bool) (*types.Transaction, error) {
	return _FreezeFacet.Contract.SetWalletFrozen(&_FreezeFacet.TransactOpts, wallet, frozen)
}

// FreezeFacetAssetFrozenIterator is returned from FilterAssetFrozen and is used to iterate over the raw logs and unpacked data for AssetFrozen events raised by the FreezeFacet contract.
type FreezeFacetAssetFrozenIterator struct {
	Event *FreezeFacetAssetFrozen // Event containing the contract specifics and raw log

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
func (it *FreezeFacetAssetFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FreezeFacetAssetFrozen)
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
		it.Event = new(FreezeFacetAssetFrozen)
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
func (it *FreezeFacetAssetFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FreezeFacetAssetFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FreezeFacetAssetFrozen represents a AssetFrozen event raised by the FreezeFacet contract.
type FreezeFacetAssetFrozen struct {
	TokenId *big.Int
	Wallet  common.Address
	Frozen  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAssetFrozen is a free log retrieval operation binding the contract event 0x114395d7a412c3817dc7b3b1b49ea50fbe7e1f46c18c14bc5ef60c1e17091441.
//
// Solidity: event AssetFrozen(uint256 indexed tokenId, address indexed wallet, bool frozen)
func (_FreezeFacet *FreezeFacetFilterer) FilterAssetFrozen(opts *bind.FilterOpts, tokenId []*big.Int, wallet []common.Address) (*FreezeFacetAssetFrozenIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _FreezeFacet.contract.FilterLogs(opts, "AssetFrozen", tokenIdRule, walletRule)
	if err != nil {
		return nil, err
	}
	return &FreezeFacetAssetFrozenIterator{contract: _FreezeFacet.contract, event: "AssetFrozen", logs: logs, sub: sub}, nil
}

// WatchAssetFrozen is a free log subscription operation binding the contract event 0x114395d7a412c3817dc7b3b1b49ea50fbe7e1f46c18c14bc5ef60c1e17091441.
//
// Solidity: event AssetFrozen(uint256 indexed tokenId, address indexed wallet, bool frozen)
func (_FreezeFacet *FreezeFacetFilterer) WatchAssetFrozen(opts *bind.WatchOpts, sink chan<- *FreezeFacetAssetFrozen, tokenId []*big.Int, wallet []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _FreezeFacet.contract.WatchLogs(opts, "AssetFrozen", tokenIdRule, walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FreezeFacetAssetFrozen)
				if err := _FreezeFacet.contract.UnpackLog(event, "AssetFrozen", log); err != nil {
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

// ParseAssetFrozen is a log parse operation binding the contract event 0x114395d7a412c3817dc7b3b1b49ea50fbe7e1f46c18c14bc5ef60c1e17091441.
//
// Solidity: event AssetFrozen(uint256 indexed tokenId, address indexed wallet, bool frozen)
func (_FreezeFacet *FreezeFacetFilterer) ParseAssetFrozen(log types.Log) (*FreezeFacetAssetFrozen, error) {
	event := new(FreezeFacetAssetFrozen)
	if err := _FreezeFacet.contract.UnpackLog(event, "AssetFrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FreezeFacetLockupSetIterator is returned from FilterLockupSet and is used to iterate over the raw logs and unpacked data for LockupSet events raised by the FreezeFacet contract.
type FreezeFacetLockupSetIterator struct {
	Event *FreezeFacetLockupSet // Event containing the contract specifics and raw log

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
func (it *FreezeFacetLockupSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FreezeFacetLockupSet)
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
		it.Event = new(FreezeFacetLockupSet)
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
func (it *FreezeFacetLockupSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FreezeFacetLockupSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FreezeFacetLockupSet represents a LockupSet event raised by the FreezeFacet contract.
type FreezeFacetLockupSet struct {
	TokenId *big.Int
	Wallet  common.Address
	Expiry  uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLockupSet is a free log retrieval operation binding the contract event 0x37139ab567e12fbeaa4db4ec0eb1013a532c5b3a1dc207832482689076f4f62e.
//
// Solidity: event LockupSet(uint256 indexed tokenId, address indexed wallet, uint64 expiry)
func (_FreezeFacet *FreezeFacetFilterer) FilterLockupSet(opts *bind.FilterOpts, tokenId []*big.Int, wallet []common.Address) (*FreezeFacetLockupSetIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _FreezeFacet.contract.FilterLogs(opts, "LockupSet", tokenIdRule, walletRule)
	if err != nil {
		return nil, err
	}
	return &FreezeFacetLockupSetIterator{contract: _FreezeFacet.contract, event: "LockupSet", logs: logs, sub: sub}, nil
}

// WatchLockupSet is a free log subscription operation binding the contract event 0x37139ab567e12fbeaa4db4ec0eb1013a532c5b3a1dc207832482689076f4f62e.
//
// Solidity: event LockupSet(uint256 indexed tokenId, address indexed wallet, uint64 expiry)
func (_FreezeFacet *FreezeFacetFilterer) WatchLockupSet(opts *bind.WatchOpts, sink chan<- *FreezeFacetLockupSet, tokenId []*big.Int, wallet []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _FreezeFacet.contract.WatchLogs(opts, "LockupSet", tokenIdRule, walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FreezeFacetLockupSet)
				if err := _FreezeFacet.contract.UnpackLog(event, "LockupSet", log); err != nil {
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

// ParseLockupSet is a log parse operation binding the contract event 0x37139ab567e12fbeaa4db4ec0eb1013a532c5b3a1dc207832482689076f4f62e.
//
// Solidity: event LockupSet(uint256 indexed tokenId, address indexed wallet, uint64 expiry)
func (_FreezeFacet *FreezeFacetFilterer) ParseLockupSet(log types.Log) (*FreezeFacetLockupSet, error) {
	event := new(FreezeFacetLockupSet)
	if err := _FreezeFacet.contract.UnpackLog(event, "LockupSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FreezeFacetPartialFreezeIterator is returned from FilterPartialFreeze and is used to iterate over the raw logs and unpacked data for PartialFreeze events raised by the FreezeFacet contract.
type FreezeFacetPartialFreezeIterator struct {
	Event *FreezeFacetPartialFreeze // Event containing the contract specifics and raw log

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
func (it *FreezeFacetPartialFreezeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FreezeFacetPartialFreeze)
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
		it.Event = new(FreezeFacetPartialFreeze)
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
func (it *FreezeFacetPartialFreezeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FreezeFacetPartialFreezeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FreezeFacetPartialFreeze represents a PartialFreeze event raised by the FreezeFacet contract.
type FreezeFacetPartialFreeze struct {
	TokenId *big.Int
	Wallet  common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPartialFreeze is a free log retrieval operation binding the contract event 0xec6ff36aa6f763ce67ee07c976be81c4ba56033079e067878ba01f7994abf78d.
//
// Solidity: event PartialFreeze(uint256 indexed tokenId, address indexed wallet, uint256 amount)
func (_FreezeFacet *FreezeFacetFilterer) FilterPartialFreeze(opts *bind.FilterOpts, tokenId []*big.Int, wallet []common.Address) (*FreezeFacetPartialFreezeIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _FreezeFacet.contract.FilterLogs(opts, "PartialFreeze", tokenIdRule, walletRule)
	if err != nil {
		return nil, err
	}
	return &FreezeFacetPartialFreezeIterator{contract: _FreezeFacet.contract, event: "PartialFreeze", logs: logs, sub: sub}, nil
}

// WatchPartialFreeze is a free log subscription operation binding the contract event 0xec6ff36aa6f763ce67ee07c976be81c4ba56033079e067878ba01f7994abf78d.
//
// Solidity: event PartialFreeze(uint256 indexed tokenId, address indexed wallet, uint256 amount)
func (_FreezeFacet *FreezeFacetFilterer) WatchPartialFreeze(opts *bind.WatchOpts, sink chan<- *FreezeFacetPartialFreeze, tokenId []*big.Int, wallet []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _FreezeFacet.contract.WatchLogs(opts, "PartialFreeze", tokenIdRule, walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FreezeFacetPartialFreeze)
				if err := _FreezeFacet.contract.UnpackLog(event, "PartialFreeze", log); err != nil {
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

// ParsePartialFreeze is a log parse operation binding the contract event 0xec6ff36aa6f763ce67ee07c976be81c4ba56033079e067878ba01f7994abf78d.
//
// Solidity: event PartialFreeze(uint256 indexed tokenId, address indexed wallet, uint256 amount)
func (_FreezeFacet *FreezeFacetFilterer) ParsePartialFreeze(log types.Log) (*FreezeFacetPartialFreeze, error) {
	event := new(FreezeFacetPartialFreeze)
	if err := _FreezeFacet.contract.UnpackLog(event, "PartialFreeze", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FreezeFacetWalletFrozenIterator is returned from FilterWalletFrozen and is used to iterate over the raw logs and unpacked data for WalletFrozen events raised by the FreezeFacet contract.
type FreezeFacetWalletFrozenIterator struct {
	Event *FreezeFacetWalletFrozen // Event containing the contract specifics and raw log

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
func (it *FreezeFacetWalletFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FreezeFacetWalletFrozen)
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
		it.Event = new(FreezeFacetWalletFrozen)
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
func (it *FreezeFacetWalletFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FreezeFacetWalletFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FreezeFacetWalletFrozen represents a WalletFrozen event raised by the FreezeFacet contract.
type FreezeFacetWalletFrozen struct {
	Wallet common.Address
	Frozen bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWalletFrozen is a free log retrieval operation binding the contract event 0xd17e29b9185a00800171e680c627983f284cdf937b14ea98ef64d9b44a72f2fc.
//
// Solidity: event WalletFrozen(address indexed wallet, bool frozen)
func (_FreezeFacet *FreezeFacetFilterer) FilterWalletFrozen(opts *bind.FilterOpts, wallet []common.Address) (*FreezeFacetWalletFrozenIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _FreezeFacet.contract.FilterLogs(opts, "WalletFrozen", walletRule)
	if err != nil {
		return nil, err
	}
	return &FreezeFacetWalletFrozenIterator{contract: _FreezeFacet.contract, event: "WalletFrozen", logs: logs, sub: sub}, nil
}

// WatchWalletFrozen is a free log subscription operation binding the contract event 0xd17e29b9185a00800171e680c627983f284cdf937b14ea98ef64d9b44a72f2fc.
//
// Solidity: event WalletFrozen(address indexed wallet, bool frozen)
func (_FreezeFacet *FreezeFacetFilterer) WatchWalletFrozen(opts *bind.WatchOpts, sink chan<- *FreezeFacetWalletFrozen, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _FreezeFacet.contract.WatchLogs(opts, "WalletFrozen", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FreezeFacetWalletFrozen)
				if err := _FreezeFacet.contract.UnpackLog(event, "WalletFrozen", log); err != nil {
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

// ParseWalletFrozen is a log parse operation binding the contract event 0xd17e29b9185a00800171e680c627983f284cdf937b14ea98ef64d9b44a72f2fc.
//
// Solidity: event WalletFrozen(address indexed wallet, bool frozen)
func (_FreezeFacet *FreezeFacetFilterer) ParseWalletFrozen(log types.Log) (*FreezeFacetWalletFrozen, error) {
	event := new(FreezeFacetWalletFrozen)
	if err := _FreezeFacet.contract.UnpackLog(event, "WalletFrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
