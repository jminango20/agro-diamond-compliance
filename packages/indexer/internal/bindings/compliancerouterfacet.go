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

// ComplianceRouterFacetMetaData contains all meta data concerning the ComplianceRouterFacet contract.
var ComplianceRouterFacetMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"burned\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"canTransfer\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"ok\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"reason\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getComplianceModule\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"module\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minted\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferred\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"error\",\"name\":\"ComplianceRouterFacet__AssetNotRegistered\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
}

// ComplianceRouterFacetABI is the input ABI used to generate the binding from.
// Deprecated: Use ComplianceRouterFacetMetaData.ABI instead.
var ComplianceRouterFacetABI = ComplianceRouterFacetMetaData.ABI

// ComplianceRouterFacet is an auto generated Go binding around an Ethereum contract.
type ComplianceRouterFacet struct {
	ComplianceRouterFacetCaller     // Read-only binding to the contract
	ComplianceRouterFacetTransactor // Write-only binding to the contract
	ComplianceRouterFacetFilterer   // Log filterer for contract events
}

// ComplianceRouterFacetCaller is an auto generated read-only Go binding around an Ethereum contract.
type ComplianceRouterFacetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComplianceRouterFacetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ComplianceRouterFacetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComplianceRouterFacetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ComplianceRouterFacetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComplianceRouterFacetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ComplianceRouterFacetSession struct {
	Contract     *ComplianceRouterFacet // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ComplianceRouterFacetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ComplianceRouterFacetCallerSession struct {
	Contract *ComplianceRouterFacetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ComplianceRouterFacetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ComplianceRouterFacetTransactorSession struct {
	Contract     *ComplianceRouterFacetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ComplianceRouterFacetRaw is an auto generated low-level Go binding around an Ethereum contract.
type ComplianceRouterFacetRaw struct {
	Contract *ComplianceRouterFacet // Generic contract binding to access the raw methods on
}

// ComplianceRouterFacetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ComplianceRouterFacetCallerRaw struct {
	Contract *ComplianceRouterFacetCaller // Generic read-only contract binding to access the raw methods on
}

// ComplianceRouterFacetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ComplianceRouterFacetTransactorRaw struct {
	Contract *ComplianceRouterFacetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewComplianceRouterFacet creates a new instance of ComplianceRouterFacet, bound to a specific deployed contract.
func NewComplianceRouterFacet(address common.Address, backend bind.ContractBackend) (*ComplianceRouterFacet, error) {
	contract, err := bindComplianceRouterFacet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ComplianceRouterFacet{ComplianceRouterFacetCaller: ComplianceRouterFacetCaller{contract: contract}, ComplianceRouterFacetTransactor: ComplianceRouterFacetTransactor{contract: contract}, ComplianceRouterFacetFilterer: ComplianceRouterFacetFilterer{contract: contract}}, nil
}

// NewComplianceRouterFacetCaller creates a new read-only instance of ComplianceRouterFacet, bound to a specific deployed contract.
func NewComplianceRouterFacetCaller(address common.Address, caller bind.ContractCaller) (*ComplianceRouterFacetCaller, error) {
	contract, err := bindComplianceRouterFacet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ComplianceRouterFacetCaller{contract: contract}, nil
}

// NewComplianceRouterFacetTransactor creates a new write-only instance of ComplianceRouterFacet, bound to a specific deployed contract.
func NewComplianceRouterFacetTransactor(address common.Address, transactor bind.ContractTransactor) (*ComplianceRouterFacetTransactor, error) {
	contract, err := bindComplianceRouterFacet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ComplianceRouterFacetTransactor{contract: contract}, nil
}

// NewComplianceRouterFacetFilterer creates a new log filterer instance of ComplianceRouterFacet, bound to a specific deployed contract.
func NewComplianceRouterFacetFilterer(address common.Address, filterer bind.ContractFilterer) (*ComplianceRouterFacetFilterer, error) {
	contract, err := bindComplianceRouterFacet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ComplianceRouterFacetFilterer{contract: contract}, nil
}

// bindComplianceRouterFacet binds a generic wrapper to an already deployed contract.
func bindComplianceRouterFacet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ComplianceRouterFacetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ComplianceRouterFacet *ComplianceRouterFacetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ComplianceRouterFacet.Contract.ComplianceRouterFacetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ComplianceRouterFacet *ComplianceRouterFacetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ComplianceRouterFacet.Contract.ComplianceRouterFacetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ComplianceRouterFacet *ComplianceRouterFacetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ComplianceRouterFacet.Contract.ComplianceRouterFacetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ComplianceRouterFacet *ComplianceRouterFacetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ComplianceRouterFacet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ComplianceRouterFacet *ComplianceRouterFacetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ComplianceRouterFacet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ComplianceRouterFacet *ComplianceRouterFacetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ComplianceRouterFacet.Contract.contract.Transact(opts, method, params...)
}

// CanTransfer is a free data retrieval call binding the contract method 0x765ccf4f.
//
// Solidity: function canTransfer(uint256 tokenId, address from, address to, uint256 amount, bytes data) view returns(bool ok, bytes32 reason)
func (_ComplianceRouterFacet *ComplianceRouterFacetCaller) CanTransfer(opts *bind.CallOpts, tokenId *big.Int, from common.Address, to common.Address, amount *big.Int, data []byte) (struct {
	Ok     bool
	Reason [32]byte
}, error) {
	var out []interface{}
	err := _ComplianceRouterFacet.contract.Call(opts, &out, "canTransfer", tokenId, from, to, amount, data)

	outstruct := new(struct {
		Ok     bool
		Reason [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ok = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Reason = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// CanTransfer is a free data retrieval call binding the contract method 0x765ccf4f.
//
// Solidity: function canTransfer(uint256 tokenId, address from, address to, uint256 amount, bytes data) view returns(bool ok, bytes32 reason)
func (_ComplianceRouterFacet *ComplianceRouterFacetSession) CanTransfer(tokenId *big.Int, from common.Address, to common.Address, amount *big.Int, data []byte) (struct {
	Ok     bool
	Reason [32]byte
}, error) {
	return _ComplianceRouterFacet.Contract.CanTransfer(&_ComplianceRouterFacet.CallOpts, tokenId, from, to, amount, data)
}

// CanTransfer is a free data retrieval call binding the contract method 0x765ccf4f.
//
// Solidity: function canTransfer(uint256 tokenId, address from, address to, uint256 amount, bytes data) view returns(bool ok, bytes32 reason)
func (_ComplianceRouterFacet *ComplianceRouterFacetCallerSession) CanTransfer(tokenId *big.Int, from common.Address, to common.Address, amount *big.Int, data []byte) (struct {
	Ok     bool
	Reason [32]byte
}, error) {
	return _ComplianceRouterFacet.Contract.CanTransfer(&_ComplianceRouterFacet.CallOpts, tokenId, from, to, amount, data)
}

// GetComplianceModule is a free data retrieval call binding the contract method 0x52dc0250.
//
// Solidity: function getComplianceModule(uint256 tokenId) view returns(address module)
func (_ComplianceRouterFacet *ComplianceRouterFacetCaller) GetComplianceModule(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ComplianceRouterFacet.contract.Call(opts, &out, "getComplianceModule", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetComplianceModule is a free data retrieval call binding the contract method 0x52dc0250.
//
// Solidity: function getComplianceModule(uint256 tokenId) view returns(address module)
func (_ComplianceRouterFacet *ComplianceRouterFacetSession) GetComplianceModule(tokenId *big.Int) (common.Address, error) {
	return _ComplianceRouterFacet.Contract.GetComplianceModule(&_ComplianceRouterFacet.CallOpts, tokenId)
}

// GetComplianceModule is a free data retrieval call binding the contract method 0x52dc0250.
//
// Solidity: function getComplianceModule(uint256 tokenId) view returns(address module)
func (_ComplianceRouterFacet *ComplianceRouterFacetCallerSession) GetComplianceModule(tokenId *big.Int) (common.Address, error) {
	return _ComplianceRouterFacet.Contract.GetComplianceModule(&_ComplianceRouterFacet.CallOpts, tokenId)
}

// Burned is a paid mutator transaction binding the contract method 0xae7d103e.
//
// Solidity: function burned(uint256 tokenId, address from, uint256 amount) returns()
func (_ComplianceRouterFacet *ComplianceRouterFacetTransactor) Burned(opts *bind.TransactOpts, tokenId *big.Int, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ComplianceRouterFacet.contract.Transact(opts, "burned", tokenId, from, amount)
}

// Burned is a paid mutator transaction binding the contract method 0xae7d103e.
//
// Solidity: function burned(uint256 tokenId, address from, uint256 amount) returns()
func (_ComplianceRouterFacet *ComplianceRouterFacetSession) Burned(tokenId *big.Int, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ComplianceRouterFacet.Contract.Burned(&_ComplianceRouterFacet.TransactOpts, tokenId, from, amount)
}

// Burned is a paid mutator transaction binding the contract method 0xae7d103e.
//
// Solidity: function burned(uint256 tokenId, address from, uint256 amount) returns()
func (_ComplianceRouterFacet *ComplianceRouterFacetTransactorSession) Burned(tokenId *big.Int, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ComplianceRouterFacet.Contract.Burned(&_ComplianceRouterFacet.TransactOpts, tokenId, from, amount)
}

// Minted is a paid mutator transaction binding the contract method 0x23204584.
//
// Solidity: function minted(uint256 tokenId, address to, uint256 amount) returns()
func (_ComplianceRouterFacet *ComplianceRouterFacetTransactor) Minted(opts *bind.TransactOpts, tokenId *big.Int, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ComplianceRouterFacet.contract.Transact(opts, "minted", tokenId, to, amount)
}

// Minted is a paid mutator transaction binding the contract method 0x23204584.
//
// Solidity: function minted(uint256 tokenId, address to, uint256 amount) returns()
func (_ComplianceRouterFacet *ComplianceRouterFacetSession) Minted(tokenId *big.Int, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ComplianceRouterFacet.Contract.Minted(&_ComplianceRouterFacet.TransactOpts, tokenId, to, amount)
}

// Minted is a paid mutator transaction binding the contract method 0x23204584.
//
// Solidity: function minted(uint256 tokenId, address to, uint256 amount) returns()
func (_ComplianceRouterFacet *ComplianceRouterFacetTransactorSession) Minted(tokenId *big.Int, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ComplianceRouterFacet.Contract.Minted(&_ComplianceRouterFacet.TransactOpts, tokenId, to, amount)
}

// Transferred is a paid mutator transaction binding the contract method 0x63a56d3a.
//
// Solidity: function transferred(uint256 tokenId, address from, address to, uint256 amount) returns()
func (_ComplianceRouterFacet *ComplianceRouterFacetTransactor) Transferred(opts *bind.TransactOpts, tokenId *big.Int, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ComplianceRouterFacet.contract.Transact(opts, "transferred", tokenId, from, to, amount)
}

// Transferred is a paid mutator transaction binding the contract method 0x63a56d3a.
//
// Solidity: function transferred(uint256 tokenId, address from, address to, uint256 amount) returns()
func (_ComplianceRouterFacet *ComplianceRouterFacetSession) Transferred(tokenId *big.Int, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ComplianceRouterFacet.Contract.Transferred(&_ComplianceRouterFacet.TransactOpts, tokenId, from, to, amount)
}

// Transferred is a paid mutator transaction binding the contract method 0x63a56d3a.
//
// Solidity: function transferred(uint256 tokenId, address from, address to, uint256 amount) returns()
func (_ComplianceRouterFacet *ComplianceRouterFacetTransactorSession) Transferred(tokenId *big.Int, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ComplianceRouterFacet.Contract.Transferred(&_ComplianceRouterFacet.TransactOpts, tokenId, from, to, amount)
}
