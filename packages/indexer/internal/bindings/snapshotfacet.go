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

// SnapshotFacetMetaData contains all meta data concerning the SnapshotFacet contract.
var SnapshotFacetMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"createSnapshot\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"snapshotId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getLatestSnapshotId\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSnapshot\",\"inputs\":[{\"name\":\"snapshotId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"holderCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSnapshotBalance\",\"inputs\":[{\"name\":\"snapshotId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"holder\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recorded\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenSnapshots\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextSnapshotId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recordHolder\",\"inputs\":[{\"name\":\"snapshotId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"holder\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"recordHoldersBatch\",\"inputs\":[{\"name\":\"snapshotId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"holders\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"HolderRecorded\",\"inputs\":[{\"name\":\"snapshotId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"holder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SnapshotCreated\",\"inputs\":[{\"name\":\"snapshotId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"totalSupply\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"SnapshotFacet__AssetNotRegistered\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SnapshotFacet__HolderAlreadyRecorded\",\"inputs\":[{\"name\":\"snapshotId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"holder\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SnapshotFacet__SnapshotNotFound\",\"inputs\":[{\"name\":\"snapshotId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SnapshotFacet__Unauthorized\",\"inputs\":[]}]",
}

// SnapshotFacetABI is the input ABI used to generate the binding from.
// Deprecated: Use SnapshotFacetMetaData.ABI instead.
var SnapshotFacetABI = SnapshotFacetMetaData.ABI

// SnapshotFacet is an auto generated Go binding around an Ethereum contract.
type SnapshotFacet struct {
	SnapshotFacetCaller     // Read-only binding to the contract
	SnapshotFacetTransactor // Write-only binding to the contract
	SnapshotFacetFilterer   // Log filterer for contract events
}

// SnapshotFacetCaller is an auto generated read-only Go binding around an Ethereum contract.
type SnapshotFacetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SnapshotFacetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SnapshotFacetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SnapshotFacetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SnapshotFacetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SnapshotFacetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SnapshotFacetSession struct {
	Contract     *SnapshotFacet    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SnapshotFacetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SnapshotFacetCallerSession struct {
	Contract *SnapshotFacetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SnapshotFacetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SnapshotFacetTransactorSession struct {
	Contract     *SnapshotFacetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SnapshotFacetRaw is an auto generated low-level Go binding around an Ethereum contract.
type SnapshotFacetRaw struct {
	Contract *SnapshotFacet // Generic contract binding to access the raw methods on
}

// SnapshotFacetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SnapshotFacetCallerRaw struct {
	Contract *SnapshotFacetCaller // Generic read-only contract binding to access the raw methods on
}

// SnapshotFacetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SnapshotFacetTransactorRaw struct {
	Contract *SnapshotFacetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSnapshotFacet creates a new instance of SnapshotFacet, bound to a specific deployed contract.
func NewSnapshotFacet(address common.Address, backend bind.ContractBackend) (*SnapshotFacet, error) {
	contract, err := bindSnapshotFacet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SnapshotFacet{SnapshotFacetCaller: SnapshotFacetCaller{contract: contract}, SnapshotFacetTransactor: SnapshotFacetTransactor{contract: contract}, SnapshotFacetFilterer: SnapshotFacetFilterer{contract: contract}}, nil
}

// NewSnapshotFacetCaller creates a new read-only instance of SnapshotFacet, bound to a specific deployed contract.
func NewSnapshotFacetCaller(address common.Address, caller bind.ContractCaller) (*SnapshotFacetCaller, error) {
	contract, err := bindSnapshotFacet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SnapshotFacetCaller{contract: contract}, nil
}

// NewSnapshotFacetTransactor creates a new write-only instance of SnapshotFacet, bound to a specific deployed contract.
func NewSnapshotFacetTransactor(address common.Address, transactor bind.ContractTransactor) (*SnapshotFacetTransactor, error) {
	contract, err := bindSnapshotFacet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SnapshotFacetTransactor{contract: contract}, nil
}

// NewSnapshotFacetFilterer creates a new log filterer instance of SnapshotFacet, bound to a specific deployed contract.
func NewSnapshotFacetFilterer(address common.Address, filterer bind.ContractFilterer) (*SnapshotFacetFilterer, error) {
	contract, err := bindSnapshotFacet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SnapshotFacetFilterer{contract: contract}, nil
}

// bindSnapshotFacet binds a generic wrapper to an already deployed contract.
func bindSnapshotFacet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SnapshotFacetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SnapshotFacet *SnapshotFacetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SnapshotFacet.Contract.SnapshotFacetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SnapshotFacet *SnapshotFacetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SnapshotFacet.Contract.SnapshotFacetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SnapshotFacet *SnapshotFacetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SnapshotFacet.Contract.SnapshotFacetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SnapshotFacet *SnapshotFacetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SnapshotFacet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SnapshotFacet *SnapshotFacetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SnapshotFacet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SnapshotFacet *SnapshotFacetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SnapshotFacet.Contract.contract.Transact(opts, method, params...)
}

// GetLatestSnapshotId is a free data retrieval call binding the contract method 0x3d4f8df3.
//
// Solidity: function getLatestSnapshotId(uint256 tokenId) view returns(uint256)
func (_SnapshotFacet *SnapshotFacetCaller) GetLatestSnapshotId(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SnapshotFacet.contract.Call(opts, &out, "getLatestSnapshotId", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestSnapshotId is a free data retrieval call binding the contract method 0x3d4f8df3.
//
// Solidity: function getLatestSnapshotId(uint256 tokenId) view returns(uint256)
func (_SnapshotFacet *SnapshotFacetSession) GetLatestSnapshotId(tokenId *big.Int) (*big.Int, error) {
	return _SnapshotFacet.Contract.GetLatestSnapshotId(&_SnapshotFacet.CallOpts, tokenId)
}

// GetLatestSnapshotId is a free data retrieval call binding the contract method 0x3d4f8df3.
//
// Solidity: function getLatestSnapshotId(uint256 tokenId) view returns(uint256)
func (_SnapshotFacet *SnapshotFacetCallerSession) GetLatestSnapshotId(tokenId *big.Int) (*big.Int, error) {
	return _SnapshotFacet.Contract.GetLatestSnapshotId(&_SnapshotFacet.CallOpts, tokenId)
}

// GetSnapshot is a free data retrieval call binding the contract method 0x76f10ad0.
//
// Solidity: function getSnapshot(uint256 snapshotId) view returns(uint256 tokenId, uint256 totalSupply, uint64 timestamp, uint256 holderCount)
func (_SnapshotFacet *SnapshotFacetCaller) GetSnapshot(opts *bind.CallOpts, snapshotId *big.Int) (struct {
	TokenId     *big.Int
	TotalSupply *big.Int
	Timestamp   uint64
	HolderCount *big.Int
}, error) {
	var out []interface{}
	err := _SnapshotFacet.contract.Call(opts, &out, "getSnapshot", snapshotId)

	outstruct := new(struct {
		TokenId     *big.Int
		TotalSupply *big.Int
		Timestamp   uint64
		HolderCount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalSupply = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.HolderCount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetSnapshot is a free data retrieval call binding the contract method 0x76f10ad0.
//
// Solidity: function getSnapshot(uint256 snapshotId) view returns(uint256 tokenId, uint256 totalSupply, uint64 timestamp, uint256 holderCount)
func (_SnapshotFacet *SnapshotFacetSession) GetSnapshot(snapshotId *big.Int) (struct {
	TokenId     *big.Int
	TotalSupply *big.Int
	Timestamp   uint64
	HolderCount *big.Int
}, error) {
	return _SnapshotFacet.Contract.GetSnapshot(&_SnapshotFacet.CallOpts, snapshotId)
}

// GetSnapshot is a free data retrieval call binding the contract method 0x76f10ad0.
//
// Solidity: function getSnapshot(uint256 snapshotId) view returns(uint256 tokenId, uint256 totalSupply, uint64 timestamp, uint256 holderCount)
func (_SnapshotFacet *SnapshotFacetCallerSession) GetSnapshot(snapshotId *big.Int) (struct {
	TokenId     *big.Int
	TotalSupply *big.Int
	Timestamp   uint64
	HolderCount *big.Int
}, error) {
	return _SnapshotFacet.Contract.GetSnapshot(&_SnapshotFacet.CallOpts, snapshotId)
}

// GetSnapshotBalance is a free data retrieval call binding the contract method 0x65e789a7.
//
// Solidity: function getSnapshotBalance(uint256 snapshotId, address holder) view returns(uint256 balance, bool recorded)
func (_SnapshotFacet *SnapshotFacetCaller) GetSnapshotBalance(opts *bind.CallOpts, snapshotId *big.Int, holder common.Address) (struct {
	Balance  *big.Int
	Recorded bool
}, error) {
	var out []interface{}
	err := _SnapshotFacet.contract.Call(opts, &out, "getSnapshotBalance", snapshotId, holder)

	outstruct := new(struct {
		Balance  *big.Int
		Recorded bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Balance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Recorded = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetSnapshotBalance is a free data retrieval call binding the contract method 0x65e789a7.
//
// Solidity: function getSnapshotBalance(uint256 snapshotId, address holder) view returns(uint256 balance, bool recorded)
func (_SnapshotFacet *SnapshotFacetSession) GetSnapshotBalance(snapshotId *big.Int, holder common.Address) (struct {
	Balance  *big.Int
	Recorded bool
}, error) {
	return _SnapshotFacet.Contract.GetSnapshotBalance(&_SnapshotFacet.CallOpts, snapshotId, holder)
}

// GetSnapshotBalance is a free data retrieval call binding the contract method 0x65e789a7.
//
// Solidity: function getSnapshotBalance(uint256 snapshotId, address holder) view returns(uint256 balance, bool recorded)
func (_SnapshotFacet *SnapshotFacetCallerSession) GetSnapshotBalance(snapshotId *big.Int, holder common.Address) (struct {
	Balance  *big.Int
	Recorded bool
}, error) {
	return _SnapshotFacet.Contract.GetSnapshotBalance(&_SnapshotFacet.CallOpts, snapshotId, holder)
}

// GetTokenSnapshots is a free data retrieval call binding the contract method 0x2b085c27.
//
// Solidity: function getTokenSnapshots(uint256 tokenId) view returns(uint256[])
func (_SnapshotFacet *SnapshotFacetCaller) GetTokenSnapshots(opts *bind.CallOpts, tokenId *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _SnapshotFacet.contract.Call(opts, &out, "getTokenSnapshots", tokenId)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTokenSnapshots is a free data retrieval call binding the contract method 0x2b085c27.
//
// Solidity: function getTokenSnapshots(uint256 tokenId) view returns(uint256[])
func (_SnapshotFacet *SnapshotFacetSession) GetTokenSnapshots(tokenId *big.Int) ([]*big.Int, error) {
	return _SnapshotFacet.Contract.GetTokenSnapshots(&_SnapshotFacet.CallOpts, tokenId)
}

// GetTokenSnapshots is a free data retrieval call binding the contract method 0x2b085c27.
//
// Solidity: function getTokenSnapshots(uint256 tokenId) view returns(uint256[])
func (_SnapshotFacet *SnapshotFacetCallerSession) GetTokenSnapshots(tokenId *big.Int) ([]*big.Int, error) {
	return _SnapshotFacet.Contract.GetTokenSnapshots(&_SnapshotFacet.CallOpts, tokenId)
}

// NextSnapshotId is a free data retrieval call binding the contract method 0xed2655cf.
//
// Solidity: function nextSnapshotId() view returns(uint256)
func (_SnapshotFacet *SnapshotFacetCaller) NextSnapshotId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SnapshotFacet.contract.Call(opts, &out, "nextSnapshotId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextSnapshotId is a free data retrieval call binding the contract method 0xed2655cf.
//
// Solidity: function nextSnapshotId() view returns(uint256)
func (_SnapshotFacet *SnapshotFacetSession) NextSnapshotId() (*big.Int, error) {
	return _SnapshotFacet.Contract.NextSnapshotId(&_SnapshotFacet.CallOpts)
}

// NextSnapshotId is a free data retrieval call binding the contract method 0xed2655cf.
//
// Solidity: function nextSnapshotId() view returns(uint256)
func (_SnapshotFacet *SnapshotFacetCallerSession) NextSnapshotId() (*big.Int, error) {
	return _SnapshotFacet.Contract.NextSnapshotId(&_SnapshotFacet.CallOpts)
}

// CreateSnapshot is a paid mutator transaction binding the contract method 0x964f5517.
//
// Solidity: function createSnapshot(uint256 tokenId) returns(uint256 snapshotId)
func (_SnapshotFacet *SnapshotFacetTransactor) CreateSnapshot(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _SnapshotFacet.contract.Transact(opts, "createSnapshot", tokenId)
}

// CreateSnapshot is a paid mutator transaction binding the contract method 0x964f5517.
//
// Solidity: function createSnapshot(uint256 tokenId) returns(uint256 snapshotId)
func (_SnapshotFacet *SnapshotFacetSession) CreateSnapshot(tokenId *big.Int) (*types.Transaction, error) {
	return _SnapshotFacet.Contract.CreateSnapshot(&_SnapshotFacet.TransactOpts, tokenId)
}

// CreateSnapshot is a paid mutator transaction binding the contract method 0x964f5517.
//
// Solidity: function createSnapshot(uint256 tokenId) returns(uint256 snapshotId)
func (_SnapshotFacet *SnapshotFacetTransactorSession) CreateSnapshot(tokenId *big.Int) (*types.Transaction, error) {
	return _SnapshotFacet.Contract.CreateSnapshot(&_SnapshotFacet.TransactOpts, tokenId)
}

// RecordHolder is a paid mutator transaction binding the contract method 0xb7dbafcd.
//
// Solidity: function recordHolder(uint256 snapshotId, address holder) returns()
func (_SnapshotFacet *SnapshotFacetTransactor) RecordHolder(opts *bind.TransactOpts, snapshotId *big.Int, holder common.Address) (*types.Transaction, error) {
	return _SnapshotFacet.contract.Transact(opts, "recordHolder", snapshotId, holder)
}

// RecordHolder is a paid mutator transaction binding the contract method 0xb7dbafcd.
//
// Solidity: function recordHolder(uint256 snapshotId, address holder) returns()
func (_SnapshotFacet *SnapshotFacetSession) RecordHolder(snapshotId *big.Int, holder common.Address) (*types.Transaction, error) {
	return _SnapshotFacet.Contract.RecordHolder(&_SnapshotFacet.TransactOpts, snapshotId, holder)
}

// RecordHolder is a paid mutator transaction binding the contract method 0xb7dbafcd.
//
// Solidity: function recordHolder(uint256 snapshotId, address holder) returns()
func (_SnapshotFacet *SnapshotFacetTransactorSession) RecordHolder(snapshotId *big.Int, holder common.Address) (*types.Transaction, error) {
	return _SnapshotFacet.Contract.RecordHolder(&_SnapshotFacet.TransactOpts, snapshotId, holder)
}

// RecordHoldersBatch is a paid mutator transaction binding the contract method 0x6a54ea31.
//
// Solidity: function recordHoldersBatch(uint256 snapshotId, address[] holders) returns()
func (_SnapshotFacet *SnapshotFacetTransactor) RecordHoldersBatch(opts *bind.TransactOpts, snapshotId *big.Int, holders []common.Address) (*types.Transaction, error) {
	return _SnapshotFacet.contract.Transact(opts, "recordHoldersBatch", snapshotId, holders)
}

// RecordHoldersBatch is a paid mutator transaction binding the contract method 0x6a54ea31.
//
// Solidity: function recordHoldersBatch(uint256 snapshotId, address[] holders) returns()
func (_SnapshotFacet *SnapshotFacetSession) RecordHoldersBatch(snapshotId *big.Int, holders []common.Address) (*types.Transaction, error) {
	return _SnapshotFacet.Contract.RecordHoldersBatch(&_SnapshotFacet.TransactOpts, snapshotId, holders)
}

// RecordHoldersBatch is a paid mutator transaction binding the contract method 0x6a54ea31.
//
// Solidity: function recordHoldersBatch(uint256 snapshotId, address[] holders) returns()
func (_SnapshotFacet *SnapshotFacetTransactorSession) RecordHoldersBatch(snapshotId *big.Int, holders []common.Address) (*types.Transaction, error) {
	return _SnapshotFacet.Contract.RecordHoldersBatch(&_SnapshotFacet.TransactOpts, snapshotId, holders)
}

// SnapshotFacetHolderRecordedIterator is returned from FilterHolderRecorded and is used to iterate over the raw logs and unpacked data for HolderRecorded events raised by the SnapshotFacet contract.
type SnapshotFacetHolderRecordedIterator struct {
	Event *SnapshotFacetHolderRecorded // Event containing the contract specifics and raw log

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
func (it *SnapshotFacetHolderRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SnapshotFacetHolderRecorded)
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
		it.Event = new(SnapshotFacetHolderRecorded)
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
func (it *SnapshotFacetHolderRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SnapshotFacetHolderRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SnapshotFacetHolderRecorded represents a HolderRecorded event raised by the SnapshotFacet contract.
type SnapshotFacetHolderRecorded struct {
	SnapshotId *big.Int
	Holder     common.Address
	Balance    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterHolderRecorded is a free log retrieval operation binding the contract event 0x7a707151249269489a5eadbbcdd3d39af9743c6393ef8a65ec6f4b44700560c4.
//
// Solidity: event HolderRecorded(uint256 indexed snapshotId, address indexed holder, uint256 balance)
func (_SnapshotFacet *SnapshotFacetFilterer) FilterHolderRecorded(opts *bind.FilterOpts, snapshotId []*big.Int, holder []common.Address) (*SnapshotFacetHolderRecordedIterator, error) {

	var snapshotIdRule []interface{}
	for _, snapshotIdItem := range snapshotId {
		snapshotIdRule = append(snapshotIdRule, snapshotIdItem)
	}
	var holderRule []interface{}
	for _, holderItem := range holder {
		holderRule = append(holderRule, holderItem)
	}

	logs, sub, err := _SnapshotFacet.contract.FilterLogs(opts, "HolderRecorded", snapshotIdRule, holderRule)
	if err != nil {
		return nil, err
	}
	return &SnapshotFacetHolderRecordedIterator{contract: _SnapshotFacet.contract, event: "HolderRecorded", logs: logs, sub: sub}, nil
}

// WatchHolderRecorded is a free log subscription operation binding the contract event 0x7a707151249269489a5eadbbcdd3d39af9743c6393ef8a65ec6f4b44700560c4.
//
// Solidity: event HolderRecorded(uint256 indexed snapshotId, address indexed holder, uint256 balance)
func (_SnapshotFacet *SnapshotFacetFilterer) WatchHolderRecorded(opts *bind.WatchOpts, sink chan<- *SnapshotFacetHolderRecorded, snapshotId []*big.Int, holder []common.Address) (event.Subscription, error) {

	var snapshotIdRule []interface{}
	for _, snapshotIdItem := range snapshotId {
		snapshotIdRule = append(snapshotIdRule, snapshotIdItem)
	}
	var holderRule []interface{}
	for _, holderItem := range holder {
		holderRule = append(holderRule, holderItem)
	}

	logs, sub, err := _SnapshotFacet.contract.WatchLogs(opts, "HolderRecorded", snapshotIdRule, holderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SnapshotFacetHolderRecorded)
				if err := _SnapshotFacet.contract.UnpackLog(event, "HolderRecorded", log); err != nil {
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

// ParseHolderRecorded is a log parse operation binding the contract event 0x7a707151249269489a5eadbbcdd3d39af9743c6393ef8a65ec6f4b44700560c4.
//
// Solidity: event HolderRecorded(uint256 indexed snapshotId, address indexed holder, uint256 balance)
func (_SnapshotFacet *SnapshotFacetFilterer) ParseHolderRecorded(log types.Log) (*SnapshotFacetHolderRecorded, error) {
	event := new(SnapshotFacetHolderRecorded)
	if err := _SnapshotFacet.contract.UnpackLog(event, "HolderRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SnapshotFacetSnapshotCreatedIterator is returned from FilterSnapshotCreated and is used to iterate over the raw logs and unpacked data for SnapshotCreated events raised by the SnapshotFacet contract.
type SnapshotFacetSnapshotCreatedIterator struct {
	Event *SnapshotFacetSnapshotCreated // Event containing the contract specifics and raw log

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
func (it *SnapshotFacetSnapshotCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SnapshotFacetSnapshotCreated)
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
		it.Event = new(SnapshotFacetSnapshotCreated)
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
func (it *SnapshotFacetSnapshotCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SnapshotFacetSnapshotCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SnapshotFacetSnapshotCreated represents a SnapshotCreated event raised by the SnapshotFacet contract.
type SnapshotFacetSnapshotCreated struct {
	SnapshotId  *big.Int
	TokenId     *big.Int
	TotalSupply *big.Int
	Timestamp   uint64
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSnapshotCreated is a free log retrieval operation binding the contract event 0x83aee848076ecf93ac8243b1c2e4ee4ab5a117741b182bfafc64a17c2faa6437.
//
// Solidity: event SnapshotCreated(uint256 indexed snapshotId, uint256 indexed tokenId, uint256 totalSupply, uint64 timestamp)
func (_SnapshotFacet *SnapshotFacetFilterer) FilterSnapshotCreated(opts *bind.FilterOpts, snapshotId []*big.Int, tokenId []*big.Int) (*SnapshotFacetSnapshotCreatedIterator, error) {

	var snapshotIdRule []interface{}
	for _, snapshotIdItem := range snapshotId {
		snapshotIdRule = append(snapshotIdRule, snapshotIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SnapshotFacet.contract.FilterLogs(opts, "SnapshotCreated", snapshotIdRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SnapshotFacetSnapshotCreatedIterator{contract: _SnapshotFacet.contract, event: "SnapshotCreated", logs: logs, sub: sub}, nil
}

// WatchSnapshotCreated is a free log subscription operation binding the contract event 0x83aee848076ecf93ac8243b1c2e4ee4ab5a117741b182bfafc64a17c2faa6437.
//
// Solidity: event SnapshotCreated(uint256 indexed snapshotId, uint256 indexed tokenId, uint256 totalSupply, uint64 timestamp)
func (_SnapshotFacet *SnapshotFacetFilterer) WatchSnapshotCreated(opts *bind.WatchOpts, sink chan<- *SnapshotFacetSnapshotCreated, snapshotId []*big.Int, tokenId []*big.Int) (event.Subscription, error) {

	var snapshotIdRule []interface{}
	for _, snapshotIdItem := range snapshotId {
		snapshotIdRule = append(snapshotIdRule, snapshotIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SnapshotFacet.contract.WatchLogs(opts, "SnapshotCreated", snapshotIdRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SnapshotFacetSnapshotCreated)
				if err := _SnapshotFacet.contract.UnpackLog(event, "SnapshotCreated", log); err != nil {
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

// ParseSnapshotCreated is a log parse operation binding the contract event 0x83aee848076ecf93ac8243b1c2e4ee4ab5a117741b182bfafc64a17c2faa6437.
//
// Solidity: event SnapshotCreated(uint256 indexed snapshotId, uint256 indexed tokenId, uint256 totalSupply, uint64 timestamp)
func (_SnapshotFacet *SnapshotFacetFilterer) ParseSnapshotCreated(log types.Log) (*SnapshotFacetSnapshotCreated, error) {
	event := new(SnapshotFacetSnapshotCreated)
	if err := _SnapshotFacet.contract.UnpackLog(event, "SnapshotCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
