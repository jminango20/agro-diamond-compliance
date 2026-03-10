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

// SupplyFacetMetaData contains all meta data concerning the SupplyFacet contract.
var SupplyFacetMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"batchMint\",\"inputs\":[{\"name\":\"tokenIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"recipients\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"amounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forcedTransfer\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reasonCode\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"holderCount\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isHolder\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Burned\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ForcedTransfer\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"reasonCode\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Minted\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"SupplyFacet__ArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SupplyFacet__AssetNotRegistered\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SupplyFacet__AssetPaused\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SupplyFacet__BurnFromZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SupplyFacet__InsufficientFreeBalance\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"required\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SupplyFacet__MintToZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SupplyFacet__ProtocolPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SupplyFacet__SupplyCapExceeded\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"currentSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"cap\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SupplyFacet__TransferToZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SupplyFacet__Unauthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SupplyFacet__WalletFrozenAsset\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SupplyFacet__WalletFrozenGlobal\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// SupplyFacetABI is the input ABI used to generate the binding from.
// Deprecated: Use SupplyFacetMetaData.ABI instead.
var SupplyFacetABI = SupplyFacetMetaData.ABI

// SupplyFacet is an auto generated Go binding around an Ethereum contract.
type SupplyFacet struct {
	SupplyFacetCaller     // Read-only binding to the contract
	SupplyFacetTransactor // Write-only binding to the contract
	SupplyFacetFilterer   // Log filterer for contract events
}

// SupplyFacetCaller is an auto generated read-only Go binding around an Ethereum contract.
type SupplyFacetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SupplyFacetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SupplyFacetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SupplyFacetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SupplyFacetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SupplyFacetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SupplyFacetSession struct {
	Contract     *SupplyFacet      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SupplyFacetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SupplyFacetCallerSession struct {
	Contract *SupplyFacetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SupplyFacetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SupplyFacetTransactorSession struct {
	Contract     *SupplyFacetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SupplyFacetRaw is an auto generated low-level Go binding around an Ethereum contract.
type SupplyFacetRaw struct {
	Contract *SupplyFacet // Generic contract binding to access the raw methods on
}

// SupplyFacetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SupplyFacetCallerRaw struct {
	Contract *SupplyFacetCaller // Generic read-only contract binding to access the raw methods on
}

// SupplyFacetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SupplyFacetTransactorRaw struct {
	Contract *SupplyFacetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSupplyFacet creates a new instance of SupplyFacet, bound to a specific deployed contract.
func NewSupplyFacet(address common.Address, backend bind.ContractBackend) (*SupplyFacet, error) {
	contract, err := bindSupplyFacet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SupplyFacet{SupplyFacetCaller: SupplyFacetCaller{contract: contract}, SupplyFacetTransactor: SupplyFacetTransactor{contract: contract}, SupplyFacetFilterer: SupplyFacetFilterer{contract: contract}}, nil
}

// NewSupplyFacetCaller creates a new read-only instance of SupplyFacet, bound to a specific deployed contract.
func NewSupplyFacetCaller(address common.Address, caller bind.ContractCaller) (*SupplyFacetCaller, error) {
	contract, err := bindSupplyFacet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SupplyFacetCaller{contract: contract}, nil
}

// NewSupplyFacetTransactor creates a new write-only instance of SupplyFacet, bound to a specific deployed contract.
func NewSupplyFacetTransactor(address common.Address, transactor bind.ContractTransactor) (*SupplyFacetTransactor, error) {
	contract, err := bindSupplyFacet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SupplyFacetTransactor{contract: contract}, nil
}

// NewSupplyFacetFilterer creates a new log filterer instance of SupplyFacet, bound to a specific deployed contract.
func NewSupplyFacetFilterer(address common.Address, filterer bind.ContractFilterer) (*SupplyFacetFilterer, error) {
	contract, err := bindSupplyFacet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SupplyFacetFilterer{contract: contract}, nil
}

// bindSupplyFacet binds a generic wrapper to an already deployed contract.
func bindSupplyFacet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SupplyFacetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SupplyFacet *SupplyFacetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SupplyFacet.Contract.SupplyFacetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SupplyFacet *SupplyFacetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SupplyFacet.Contract.SupplyFacetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SupplyFacet *SupplyFacetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SupplyFacet.Contract.SupplyFacetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SupplyFacet *SupplyFacetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SupplyFacet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SupplyFacet *SupplyFacetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SupplyFacet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SupplyFacet *SupplyFacetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SupplyFacet.Contract.contract.Transact(opts, method, params...)
}

// HolderCount is a free data retrieval call binding the contract method 0x87a58d80.
//
// Solidity: function holderCount(uint256 tokenId) view returns(uint256)
func (_SupplyFacet *SupplyFacetCaller) HolderCount(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SupplyFacet.contract.Call(opts, &out, "holderCount", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HolderCount is a free data retrieval call binding the contract method 0x87a58d80.
//
// Solidity: function holderCount(uint256 tokenId) view returns(uint256)
func (_SupplyFacet *SupplyFacetSession) HolderCount(tokenId *big.Int) (*big.Int, error) {
	return _SupplyFacet.Contract.HolderCount(&_SupplyFacet.CallOpts, tokenId)
}

// HolderCount is a free data retrieval call binding the contract method 0x87a58d80.
//
// Solidity: function holderCount(uint256 tokenId) view returns(uint256)
func (_SupplyFacet *SupplyFacetCallerSession) HolderCount(tokenId *big.Int) (*big.Int, error) {
	return _SupplyFacet.Contract.HolderCount(&_SupplyFacet.CallOpts, tokenId)
}

// IsHolder is a free data retrieval call binding the contract method 0x36f732a1.
//
// Solidity: function isHolder(uint256 tokenId, address account) view returns(bool)
func (_SupplyFacet *SupplyFacetCaller) IsHolder(opts *bind.CallOpts, tokenId *big.Int, account common.Address) (bool, error) {
	var out []interface{}
	err := _SupplyFacet.contract.Call(opts, &out, "isHolder", tokenId, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsHolder is a free data retrieval call binding the contract method 0x36f732a1.
//
// Solidity: function isHolder(uint256 tokenId, address account) view returns(bool)
func (_SupplyFacet *SupplyFacetSession) IsHolder(tokenId *big.Int, account common.Address) (bool, error) {
	return _SupplyFacet.Contract.IsHolder(&_SupplyFacet.CallOpts, tokenId, account)
}

// IsHolder is a free data retrieval call binding the contract method 0x36f732a1.
//
// Solidity: function isHolder(uint256 tokenId, address account) view returns(bool)
func (_SupplyFacet *SupplyFacetCallerSession) IsHolder(tokenId *big.Int, account common.Address) (bool, error) {
	return _SupplyFacet.Contract.IsHolder(&_SupplyFacet.CallOpts, tokenId, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 tokenId) view returns(uint256)
func (_SupplyFacet *SupplyFacetCaller) TotalSupply(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SupplyFacet.contract.Call(opts, &out, "totalSupply", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 tokenId) view returns(uint256)
func (_SupplyFacet *SupplyFacetSession) TotalSupply(tokenId *big.Int) (*big.Int, error) {
	return _SupplyFacet.Contract.TotalSupply(&_SupplyFacet.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 tokenId) view returns(uint256)
func (_SupplyFacet *SupplyFacetCallerSession) TotalSupply(tokenId *big.Int) (*big.Int, error) {
	return _SupplyFacet.Contract.TotalSupply(&_SupplyFacet.CallOpts, tokenId)
}

// BatchMint is a paid mutator transaction binding the contract method 0x97514f69.
//
// Solidity: function batchMint(uint256[] tokenIds, address[] recipients, uint256[] amounts) returns()
func (_SupplyFacet *SupplyFacetTransactor) BatchMint(opts *bind.TransactOpts, tokenIds []*big.Int, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _SupplyFacet.contract.Transact(opts, "batchMint", tokenIds, recipients, amounts)
}

// BatchMint is a paid mutator transaction binding the contract method 0x97514f69.
//
// Solidity: function batchMint(uint256[] tokenIds, address[] recipients, uint256[] amounts) returns()
func (_SupplyFacet *SupplyFacetSession) BatchMint(tokenIds []*big.Int, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _SupplyFacet.Contract.BatchMint(&_SupplyFacet.TransactOpts, tokenIds, recipients, amounts)
}

// BatchMint is a paid mutator transaction binding the contract method 0x97514f69.
//
// Solidity: function batchMint(uint256[] tokenIds, address[] recipients, uint256[] amounts) returns()
func (_SupplyFacet *SupplyFacetTransactorSession) BatchMint(tokenIds []*big.Int, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _SupplyFacet.Contract.BatchMint(&_SupplyFacet.TransactOpts, tokenIds, recipients, amounts)
}

// Burn is a paid mutator transaction binding the contract method 0x9eea5f66.
//
// Solidity: function burn(uint256 tokenId, address from, uint256 amount) returns()
func (_SupplyFacet *SupplyFacetTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SupplyFacet.contract.Transact(opts, "burn", tokenId, from, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9eea5f66.
//
// Solidity: function burn(uint256 tokenId, address from, uint256 amount) returns()
func (_SupplyFacet *SupplyFacetSession) Burn(tokenId *big.Int, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SupplyFacet.Contract.Burn(&_SupplyFacet.TransactOpts, tokenId, from, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9eea5f66.
//
// Solidity: function burn(uint256 tokenId, address from, uint256 amount) returns()
func (_SupplyFacet *SupplyFacetTransactorSession) Burn(tokenId *big.Int, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SupplyFacet.Contract.Burn(&_SupplyFacet.TransactOpts, tokenId, from, amount)
}

// ForcedTransfer is a paid mutator transaction binding the contract method 0x65c1b4dc.
//
// Solidity: function forcedTransfer(uint256 tokenId, address from, address to, uint256 amount, bytes32 reasonCode) returns()
func (_SupplyFacet *SupplyFacetTransactor) ForcedTransfer(opts *bind.TransactOpts, tokenId *big.Int, from common.Address, to common.Address, amount *big.Int, reasonCode [32]byte) (*types.Transaction, error) {
	return _SupplyFacet.contract.Transact(opts, "forcedTransfer", tokenId, from, to, amount, reasonCode)
}

// ForcedTransfer is a paid mutator transaction binding the contract method 0x65c1b4dc.
//
// Solidity: function forcedTransfer(uint256 tokenId, address from, address to, uint256 amount, bytes32 reasonCode) returns()
func (_SupplyFacet *SupplyFacetSession) ForcedTransfer(tokenId *big.Int, from common.Address, to common.Address, amount *big.Int, reasonCode [32]byte) (*types.Transaction, error) {
	return _SupplyFacet.Contract.ForcedTransfer(&_SupplyFacet.TransactOpts, tokenId, from, to, amount, reasonCode)
}

// ForcedTransfer is a paid mutator transaction binding the contract method 0x65c1b4dc.
//
// Solidity: function forcedTransfer(uint256 tokenId, address from, address to, uint256 amount, bytes32 reasonCode) returns()
func (_SupplyFacet *SupplyFacetTransactorSession) ForcedTransfer(tokenId *big.Int, from common.Address, to common.Address, amount *big.Int, reasonCode [32]byte) (*types.Transaction, error) {
	return _SupplyFacet.Contract.ForcedTransfer(&_SupplyFacet.TransactOpts, tokenId, from, to, amount, reasonCode)
}

// Mint is a paid mutator transaction binding the contract method 0x836a1040.
//
// Solidity: function mint(uint256 tokenId, address to, uint256 amount) returns()
func (_SupplyFacet *SupplyFacetTransactor) Mint(opts *bind.TransactOpts, tokenId *big.Int, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SupplyFacet.contract.Transact(opts, "mint", tokenId, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x836a1040.
//
// Solidity: function mint(uint256 tokenId, address to, uint256 amount) returns()
func (_SupplyFacet *SupplyFacetSession) Mint(tokenId *big.Int, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SupplyFacet.Contract.Mint(&_SupplyFacet.TransactOpts, tokenId, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x836a1040.
//
// Solidity: function mint(uint256 tokenId, address to, uint256 amount) returns()
func (_SupplyFacet *SupplyFacetTransactorSession) Mint(tokenId *big.Int, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SupplyFacet.Contract.Mint(&_SupplyFacet.TransactOpts, tokenId, to, amount)
}

// SupplyFacetBurnedIterator is returned from FilterBurned and is used to iterate over the raw logs and unpacked data for Burned events raised by the SupplyFacet contract.
type SupplyFacetBurnedIterator struct {
	Event *SupplyFacetBurned // Event containing the contract specifics and raw log

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
func (it *SupplyFacetBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SupplyFacetBurned)
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
		it.Event = new(SupplyFacetBurned)
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
func (it *SupplyFacetBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SupplyFacetBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SupplyFacetBurned represents a Burned event raised by the SupplyFacet contract.
type SupplyFacetBurned struct {
	TokenId *big.Int
	From    common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurned is a free log retrieval operation binding the contract event 0x7a6396f9141e42bbd82eddb43e30077ef07aaafcd4ee3dfbd6adb1dca8f2445a.
//
// Solidity: event Burned(uint256 indexed tokenId, address indexed from, uint256 amount)
func (_SupplyFacet *SupplyFacetFilterer) FilterBurned(opts *bind.FilterOpts, tokenId []*big.Int, from []common.Address) (*SupplyFacetBurnedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _SupplyFacet.contract.FilterLogs(opts, "Burned", tokenIdRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &SupplyFacetBurnedIterator{contract: _SupplyFacet.contract, event: "Burned", logs: logs, sub: sub}, nil
}

// WatchBurned is a free log subscription operation binding the contract event 0x7a6396f9141e42bbd82eddb43e30077ef07aaafcd4ee3dfbd6adb1dca8f2445a.
//
// Solidity: event Burned(uint256 indexed tokenId, address indexed from, uint256 amount)
func (_SupplyFacet *SupplyFacetFilterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *SupplyFacetBurned, tokenId []*big.Int, from []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _SupplyFacet.contract.WatchLogs(opts, "Burned", tokenIdRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SupplyFacetBurned)
				if err := _SupplyFacet.contract.UnpackLog(event, "Burned", log); err != nil {
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

// ParseBurned is a log parse operation binding the contract event 0x7a6396f9141e42bbd82eddb43e30077ef07aaafcd4ee3dfbd6adb1dca8f2445a.
//
// Solidity: event Burned(uint256 indexed tokenId, address indexed from, uint256 amount)
func (_SupplyFacet *SupplyFacetFilterer) ParseBurned(log types.Log) (*SupplyFacetBurned, error) {
	event := new(SupplyFacetBurned)
	if err := _SupplyFacet.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SupplyFacetForcedTransferIterator is returned from FilterForcedTransfer and is used to iterate over the raw logs and unpacked data for ForcedTransfer events raised by the SupplyFacet contract.
type SupplyFacetForcedTransferIterator struct {
	Event *SupplyFacetForcedTransfer // Event containing the contract specifics and raw log

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
func (it *SupplyFacetForcedTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SupplyFacetForcedTransfer)
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
		it.Event = new(SupplyFacetForcedTransfer)
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
func (it *SupplyFacetForcedTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SupplyFacetForcedTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SupplyFacetForcedTransfer represents a ForcedTransfer event raised by the SupplyFacet contract.
type SupplyFacetForcedTransfer struct {
	TokenId    *big.Int
	From       common.Address
	To         common.Address
	Amount     *big.Int
	ReasonCode [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterForcedTransfer is a free log retrieval operation binding the contract event 0x17eb2ee0ff281058b0e8f1344438f0efd429d6cdd8bcece987490365e477a9be.
//
// Solidity: event ForcedTransfer(uint256 indexed tokenId, address indexed from, address indexed to, uint256 amount, bytes32 reasonCode)
func (_SupplyFacet *SupplyFacetFilterer) FilterForcedTransfer(opts *bind.FilterOpts, tokenId []*big.Int, from []common.Address, to []common.Address) (*SupplyFacetForcedTransferIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SupplyFacet.contract.FilterLogs(opts, "ForcedTransfer", tokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SupplyFacetForcedTransferIterator{contract: _SupplyFacet.contract, event: "ForcedTransfer", logs: logs, sub: sub}, nil
}

// WatchForcedTransfer is a free log subscription operation binding the contract event 0x17eb2ee0ff281058b0e8f1344438f0efd429d6cdd8bcece987490365e477a9be.
//
// Solidity: event ForcedTransfer(uint256 indexed tokenId, address indexed from, address indexed to, uint256 amount, bytes32 reasonCode)
func (_SupplyFacet *SupplyFacetFilterer) WatchForcedTransfer(opts *bind.WatchOpts, sink chan<- *SupplyFacetForcedTransfer, tokenId []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SupplyFacet.contract.WatchLogs(opts, "ForcedTransfer", tokenIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SupplyFacetForcedTransfer)
				if err := _SupplyFacet.contract.UnpackLog(event, "ForcedTransfer", log); err != nil {
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

// ParseForcedTransfer is a log parse operation binding the contract event 0x17eb2ee0ff281058b0e8f1344438f0efd429d6cdd8bcece987490365e477a9be.
//
// Solidity: event ForcedTransfer(uint256 indexed tokenId, address indexed from, address indexed to, uint256 amount, bytes32 reasonCode)
func (_SupplyFacet *SupplyFacetFilterer) ParseForcedTransfer(log types.Log) (*SupplyFacetForcedTransfer, error) {
	event := new(SupplyFacetForcedTransfer)
	if err := _SupplyFacet.contract.UnpackLog(event, "ForcedTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SupplyFacetMintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the SupplyFacet contract.
type SupplyFacetMintedIterator struct {
	Event *SupplyFacetMinted // Event containing the contract specifics and raw log

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
func (it *SupplyFacetMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SupplyFacetMinted)
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
		it.Event = new(SupplyFacetMinted)
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
func (it *SupplyFacetMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SupplyFacetMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SupplyFacetMinted represents a Minted event raised by the SupplyFacet contract.
type SupplyFacetMinted struct {
	TokenId *big.Int
	To      common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0xc9d0543a84d3510329c0783b91576878ceb484e8699944cb5610c3436b3b8e39.
//
// Solidity: event Minted(uint256 indexed tokenId, address indexed to, uint256 amount)
func (_SupplyFacet *SupplyFacetFilterer) FilterMinted(opts *bind.FilterOpts, tokenId []*big.Int, to []common.Address) (*SupplyFacetMintedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SupplyFacet.contract.FilterLogs(opts, "Minted", tokenIdRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SupplyFacetMintedIterator{contract: _SupplyFacet.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0xc9d0543a84d3510329c0783b91576878ceb484e8699944cb5610c3436b3b8e39.
//
// Solidity: event Minted(uint256 indexed tokenId, address indexed to, uint256 amount)
func (_SupplyFacet *SupplyFacetFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *SupplyFacetMinted, tokenId []*big.Int, to []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SupplyFacet.contract.WatchLogs(opts, "Minted", tokenIdRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SupplyFacetMinted)
				if err := _SupplyFacet.contract.UnpackLog(event, "Minted", log); err != nil {
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

// ParseMinted is a log parse operation binding the contract event 0xc9d0543a84d3510329c0783b91576878ceb484e8699944cb5610c3436b3b8e39.
//
// Solidity: event Minted(uint256 indexed tokenId, address indexed to, uint256 amount)
func (_SupplyFacet *SupplyFacetFilterer) ParseMinted(log types.Log) (*SupplyFacetMinted, error) {
	event := new(SupplyFacetMinted)
	if err := _SupplyFacet.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
