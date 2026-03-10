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

// AccessControlFacetMetaData contains all meta data concerning the AccessControlFacet contract.
var AccessControlFacetMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"CLAIM_ISSUER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"COMPLIANCE_ADMIN\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GOVERNANCE_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ISSUER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RECOVERY_AGENT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TRANSFER_AGENT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"adminRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdmin\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdmin\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlFacet__RoleAdminOnly\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AccessControlFacet__ZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LibDiamond__OnlyOwner\",\"inputs\":[]}]",
}

// AccessControlFacetABI is the input ABI used to generate the binding from.
// Deprecated: Use AccessControlFacetMetaData.ABI instead.
var AccessControlFacetABI = AccessControlFacetMetaData.ABI

// AccessControlFacet is an auto generated Go binding around an Ethereum contract.
type AccessControlFacet struct {
	AccessControlFacetCaller     // Read-only binding to the contract
	AccessControlFacetTransactor // Write-only binding to the contract
	AccessControlFacetFilterer   // Log filterer for contract events
}

// AccessControlFacetCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccessControlFacetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlFacetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccessControlFacetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlFacetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccessControlFacetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlFacetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccessControlFacetSession struct {
	Contract     *AccessControlFacet // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AccessControlFacetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccessControlFacetCallerSession struct {
	Contract *AccessControlFacetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// AccessControlFacetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccessControlFacetTransactorSession struct {
	Contract     *AccessControlFacetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// AccessControlFacetRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccessControlFacetRaw struct {
	Contract *AccessControlFacet // Generic contract binding to access the raw methods on
}

// AccessControlFacetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccessControlFacetCallerRaw struct {
	Contract *AccessControlFacetCaller // Generic read-only contract binding to access the raw methods on
}

// AccessControlFacetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccessControlFacetTransactorRaw struct {
	Contract *AccessControlFacetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccessControlFacet creates a new instance of AccessControlFacet, bound to a specific deployed contract.
func NewAccessControlFacet(address common.Address, backend bind.ContractBackend) (*AccessControlFacet, error) {
	contract, err := bindAccessControlFacet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccessControlFacet{AccessControlFacetCaller: AccessControlFacetCaller{contract: contract}, AccessControlFacetTransactor: AccessControlFacetTransactor{contract: contract}, AccessControlFacetFilterer: AccessControlFacetFilterer{contract: contract}}, nil
}

// NewAccessControlFacetCaller creates a new read-only instance of AccessControlFacet, bound to a specific deployed contract.
func NewAccessControlFacetCaller(address common.Address, caller bind.ContractCaller) (*AccessControlFacetCaller, error) {
	contract, err := bindAccessControlFacet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlFacetCaller{contract: contract}, nil
}

// NewAccessControlFacetTransactor creates a new write-only instance of AccessControlFacet, bound to a specific deployed contract.
func NewAccessControlFacetTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessControlFacetTransactor, error) {
	contract, err := bindAccessControlFacet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlFacetTransactor{contract: contract}, nil
}

// NewAccessControlFacetFilterer creates a new log filterer instance of AccessControlFacet, bound to a specific deployed contract.
func NewAccessControlFacetFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessControlFacetFilterer, error) {
	contract, err := bindAccessControlFacet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessControlFacetFilterer{contract: contract}, nil
}

// bindAccessControlFacet binds a generic wrapper to an already deployed contract.
func bindAccessControlFacet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AccessControlFacetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControlFacet *AccessControlFacetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControlFacet.Contract.AccessControlFacetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControlFacet *AccessControlFacetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControlFacet.Contract.AccessControlFacetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControlFacet *AccessControlFacetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControlFacet.Contract.AccessControlFacetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControlFacet *AccessControlFacetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControlFacet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControlFacet *AccessControlFacetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControlFacet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControlFacet *AccessControlFacetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControlFacet.Contract.contract.Transact(opts, method, params...)
}

// CLAIMISSUERROLE is a free data retrieval call binding the contract method 0xa044e70f.
//
// Solidity: function CLAIM_ISSUER_ROLE() view returns(bytes32)
func (_AccessControlFacet *AccessControlFacetCaller) CLAIMISSUERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AccessControlFacet.contract.Call(opts, &out, "CLAIM_ISSUER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CLAIMISSUERROLE is a free data retrieval call binding the contract method 0xa044e70f.
//
// Solidity: function CLAIM_ISSUER_ROLE() view returns(bytes32)
func (_AccessControlFacet *AccessControlFacetSession) CLAIMISSUERROLE() ([32]byte, error) {
	return _AccessControlFacet.Contract.CLAIMISSUERROLE(&_AccessControlFacet.CallOpts)
}

// CLAIMISSUERROLE is a free data retrieval call binding the contract method 0xa044e70f.
//
// Solidity: function CLAIM_ISSUER_ROLE() view returns(bytes32)
func (_AccessControlFacet *AccessControlFacetCallerSession) CLAIMISSUERROLE() ([32]byte, error) {
	return _AccessControlFacet.Contract.CLAIMISSUERROLE(&_AccessControlFacet.CallOpts)
}

// COMPLIANCEADMIN is a free data retrieval call binding the contract method 0x8b0cafd6.
//
// Solidity: function COMPLIANCE_ADMIN() view returns(bytes32)
func (_AccessControlFacet *AccessControlFacetCaller) COMPLIANCEADMIN(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AccessControlFacet.contract.Call(opts, &out, "COMPLIANCE_ADMIN")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// COMPLIANCEADMIN is a free data retrieval call binding the contract method 0x8b0cafd6.
//
// Solidity: function COMPLIANCE_ADMIN() view returns(bytes32)
func (_AccessControlFacet *AccessControlFacetSession) COMPLIANCEADMIN() ([32]byte, error) {
	return _AccessControlFacet.Contract.COMPLIANCEADMIN(&_AccessControlFacet.CallOpts)
}

// COMPLIANCEADMIN is a free data retrieval call binding the contract method 0x8b0cafd6.
//
// Solidity: function COMPLIANCE_ADMIN() view returns(bytes32)
func (_AccessControlFacet *AccessControlFacetCallerSession) COMPLIANCEADMIN() ([32]byte, error) {
	return _AccessControlFacet.Contract.COMPLIANCEADMIN(&_AccessControlFacet.CallOpts)
}

// GOVERNANCEROLE is a free data retrieval call binding the contract method 0xf36c8f5c.
//
// Solidity: function GOVERNANCE_ROLE() view returns(bytes32)
func (_AccessControlFacet *AccessControlFacetCaller) GOVERNANCEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AccessControlFacet.contract.Call(opts, &out, "GOVERNANCE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GOVERNANCEROLE is a free data retrieval call binding the contract method 0xf36c8f5c.
//
// Solidity: function GOVERNANCE_ROLE() view returns(bytes32)
func (_AccessControlFacet *AccessControlFacetSession) GOVERNANCEROLE() ([32]byte, error) {
	return _AccessControlFacet.Contract.GOVERNANCEROLE(&_AccessControlFacet.CallOpts)
}

// GOVERNANCEROLE is a free data retrieval call binding the contract method 0xf36c8f5c.
//
// Solidity: function GOVERNANCE_ROLE() view returns(bytes32)
func (_AccessControlFacet *AccessControlFacetCallerSession) GOVERNANCEROLE() ([32]byte, error) {
	return _AccessControlFacet.Contract.GOVERNANCEROLE(&_AccessControlFacet.CallOpts)
}

// ISSUERROLE is a free data retrieval call binding the contract method 0x82aefa24.
//
// Solidity: function ISSUER_ROLE() view returns(bytes32)
func (_AccessControlFacet *AccessControlFacetCaller) ISSUERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AccessControlFacet.contract.Call(opts, &out, "ISSUER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ISSUERROLE is a free data retrieval call binding the contract method 0x82aefa24.
//
// Solidity: function ISSUER_ROLE() view returns(bytes32)
func (_AccessControlFacet *AccessControlFacetSession) ISSUERROLE() ([32]byte, error) {
	return _AccessControlFacet.Contract.ISSUERROLE(&_AccessControlFacet.CallOpts)
}

// ISSUERROLE is a free data retrieval call binding the contract method 0x82aefa24.
//
// Solidity: function ISSUER_ROLE() view returns(bytes32)
func (_AccessControlFacet *AccessControlFacetCallerSession) ISSUERROLE() ([32]byte, error) {
	return _AccessControlFacet.Contract.ISSUERROLE(&_AccessControlFacet.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_AccessControlFacet *AccessControlFacetCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AccessControlFacet.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_AccessControlFacet *AccessControlFacetSession) PAUSERROLE() ([32]byte, error) {
	return _AccessControlFacet.Contract.PAUSERROLE(&_AccessControlFacet.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_AccessControlFacet *AccessControlFacetCallerSession) PAUSERROLE() ([32]byte, error) {
	return _AccessControlFacet.Contract.PAUSERROLE(&_AccessControlFacet.CallOpts)
}

// RECOVERYAGENT is a free data retrieval call binding the contract method 0x88776d3a.
//
// Solidity: function RECOVERY_AGENT() view returns(bytes32)
func (_AccessControlFacet *AccessControlFacetCaller) RECOVERYAGENT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AccessControlFacet.contract.Call(opts, &out, "RECOVERY_AGENT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RECOVERYAGENT is a free data retrieval call binding the contract method 0x88776d3a.
//
// Solidity: function RECOVERY_AGENT() view returns(bytes32)
func (_AccessControlFacet *AccessControlFacetSession) RECOVERYAGENT() ([32]byte, error) {
	return _AccessControlFacet.Contract.RECOVERYAGENT(&_AccessControlFacet.CallOpts)
}

// RECOVERYAGENT is a free data retrieval call binding the contract method 0x88776d3a.
//
// Solidity: function RECOVERY_AGENT() view returns(bytes32)
func (_AccessControlFacet *AccessControlFacetCallerSession) RECOVERYAGENT() ([32]byte, error) {
	return _AccessControlFacet.Contract.RECOVERYAGENT(&_AccessControlFacet.CallOpts)
}

// TRANSFERAGENT is a free data retrieval call binding the contract method 0x6a4141d8.
//
// Solidity: function TRANSFER_AGENT() view returns(bytes32)
func (_AccessControlFacet *AccessControlFacetCaller) TRANSFERAGENT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AccessControlFacet.contract.Call(opts, &out, "TRANSFER_AGENT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TRANSFERAGENT is a free data retrieval call binding the contract method 0x6a4141d8.
//
// Solidity: function TRANSFER_AGENT() view returns(bytes32)
func (_AccessControlFacet *AccessControlFacetSession) TRANSFERAGENT() ([32]byte, error) {
	return _AccessControlFacet.Contract.TRANSFERAGENT(&_AccessControlFacet.CallOpts)
}

// TRANSFERAGENT is a free data retrieval call binding the contract method 0x6a4141d8.
//
// Solidity: function TRANSFER_AGENT() view returns(bytes32)
func (_AccessControlFacet *AccessControlFacetCallerSession) TRANSFERAGENT() ([32]byte, error) {
	return _AccessControlFacet.Contract.TRANSFERAGENT(&_AccessControlFacet.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_AccessControlFacet *AccessControlFacetCaller) UPGRADERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AccessControlFacet.contract.Call(opts, &out, "UPGRADER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_AccessControlFacet *AccessControlFacetSession) UPGRADERROLE() ([32]byte, error) {
	return _AccessControlFacet.Contract.UPGRADERROLE(&_AccessControlFacet.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_AccessControlFacet *AccessControlFacetCallerSession) UPGRADERROLE() ([32]byte, error) {
	return _AccessControlFacet.Contract.UPGRADERROLE(&_AccessControlFacet.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessControlFacet *AccessControlFacetCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AccessControlFacet.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessControlFacet *AccessControlFacetSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AccessControlFacet.Contract.GetRoleAdmin(&_AccessControlFacet.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessControlFacet *AccessControlFacetCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AccessControlFacet.Contract.GetRoleAdmin(&_AccessControlFacet.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControlFacet *AccessControlFacetCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AccessControlFacet.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControlFacet *AccessControlFacetSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccessControlFacet.Contract.HasRole(&_AccessControlFacet.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControlFacet *AccessControlFacetCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccessControlFacet.Contract.HasRole(&_AccessControlFacet.CallOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControlFacet *AccessControlFacetTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlFacet.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControlFacet *AccessControlFacetSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlFacet.Contract.GrantRole(&_AccessControlFacet.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControlFacet *AccessControlFacetTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlFacet.Contract.GrantRole(&_AccessControlFacet.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x8bb9c5bf.
//
// Solidity: function renounceRole(bytes32 role) returns()
func (_AccessControlFacet *AccessControlFacetTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte) (*types.Transaction, error) {
	return _AccessControlFacet.contract.Transact(opts, "renounceRole", role)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x8bb9c5bf.
//
// Solidity: function renounceRole(bytes32 role) returns()
func (_AccessControlFacet *AccessControlFacetSession) RenounceRole(role [32]byte) (*types.Transaction, error) {
	return _AccessControlFacet.Contract.RenounceRole(&_AccessControlFacet.TransactOpts, role)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x8bb9c5bf.
//
// Solidity: function renounceRole(bytes32 role) returns()
func (_AccessControlFacet *AccessControlFacetTransactorSession) RenounceRole(role [32]byte) (*types.Transaction, error) {
	return _AccessControlFacet.Contract.RenounceRole(&_AccessControlFacet.TransactOpts, role)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControlFacet *AccessControlFacetTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlFacet.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControlFacet *AccessControlFacetSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlFacet.Contract.RevokeRole(&_AccessControlFacet.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControlFacet *AccessControlFacetTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlFacet.Contract.RevokeRole(&_AccessControlFacet.TransactOpts, role, account)
}

// SetRoleAdmin is a paid mutator transaction binding the contract method 0x1e4e0091.
//
// Solidity: function setRoleAdmin(bytes32 role, bytes32 adminRole) returns()
func (_AccessControlFacet *AccessControlFacetTransactor) SetRoleAdmin(opts *bind.TransactOpts, role [32]byte, adminRole [32]byte) (*types.Transaction, error) {
	return _AccessControlFacet.contract.Transact(opts, "setRoleAdmin", role, adminRole)
}

// SetRoleAdmin is a paid mutator transaction binding the contract method 0x1e4e0091.
//
// Solidity: function setRoleAdmin(bytes32 role, bytes32 adminRole) returns()
func (_AccessControlFacet *AccessControlFacetSession) SetRoleAdmin(role [32]byte, adminRole [32]byte) (*types.Transaction, error) {
	return _AccessControlFacet.Contract.SetRoleAdmin(&_AccessControlFacet.TransactOpts, role, adminRole)
}

// SetRoleAdmin is a paid mutator transaction binding the contract method 0x1e4e0091.
//
// Solidity: function setRoleAdmin(bytes32 role, bytes32 adminRole) returns()
func (_AccessControlFacet *AccessControlFacetTransactorSession) SetRoleAdmin(role [32]byte, adminRole [32]byte) (*types.Transaction, error) {
	return _AccessControlFacet.Contract.SetRoleAdmin(&_AccessControlFacet.TransactOpts, role, adminRole)
}

// AccessControlFacetRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AccessControlFacet contract.
type AccessControlFacetRoleAdminChangedIterator struct {
	Event *AccessControlFacetRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AccessControlFacetRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlFacetRoleAdminChanged)
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
		it.Event = new(AccessControlFacetRoleAdminChanged)
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
func (it *AccessControlFacetRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlFacetRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlFacetRoleAdminChanged represents a RoleAdminChanged event raised by the AccessControlFacet contract.
type AccessControlFacetRoleAdminChanged struct {
	Role          [32]byte
	PreviousAdmin [32]byte
	NewAdmin      [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdmin, bytes32 indexed newAdmin)
func (_AccessControlFacet *AccessControlFacetFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdmin [][32]byte, newAdmin [][32]byte) (*AccessControlFacetRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRule []interface{}
	for _, previousAdminItem := range previousAdmin {
		previousAdminRule = append(previousAdminRule, previousAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _AccessControlFacet.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlFacetRoleAdminChangedIterator{contract: _AccessControlFacet.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdmin, bytes32 indexed newAdmin)
func (_AccessControlFacet *AccessControlFacetFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AccessControlFacetRoleAdminChanged, role [][32]byte, previousAdmin [][32]byte, newAdmin [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRule []interface{}
	for _, previousAdminItem := range previousAdmin {
		previousAdminRule = append(previousAdminRule, previousAdminItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _AccessControlFacet.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlFacetRoleAdminChanged)
				if err := _AccessControlFacet.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdmin, bytes32 indexed newAdmin)
func (_AccessControlFacet *AccessControlFacetFilterer) ParseRoleAdminChanged(log types.Log) (*AccessControlFacetRoleAdminChanged, error) {
	event := new(AccessControlFacetRoleAdminChanged)
	if err := _AccessControlFacet.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlFacetRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AccessControlFacet contract.
type AccessControlFacetRoleGrantedIterator struct {
	Event *AccessControlFacetRoleGranted // Event containing the contract specifics and raw log

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
func (it *AccessControlFacetRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlFacetRoleGranted)
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
		it.Event = new(AccessControlFacetRoleGranted)
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
func (it *AccessControlFacetRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlFacetRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlFacetRoleGranted represents a RoleGranted event raised by the AccessControlFacet contract.
type AccessControlFacetRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControlFacet *AccessControlFacetFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AccessControlFacetRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AccessControlFacet.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlFacetRoleGrantedIterator{contract: _AccessControlFacet.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControlFacet *AccessControlFacetFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AccessControlFacetRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AccessControlFacet.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlFacetRoleGranted)
				if err := _AccessControlFacet.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControlFacet *AccessControlFacetFilterer) ParseRoleGranted(log types.Log) (*AccessControlFacetRoleGranted, error) {
	event := new(AccessControlFacetRoleGranted)
	if err := _AccessControlFacet.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlFacetRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AccessControlFacet contract.
type AccessControlFacetRoleRevokedIterator struct {
	Event *AccessControlFacetRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AccessControlFacetRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlFacetRoleRevoked)
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
		it.Event = new(AccessControlFacetRoleRevoked)
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
func (it *AccessControlFacetRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlFacetRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlFacetRoleRevoked represents a RoleRevoked event raised by the AccessControlFacet contract.
type AccessControlFacetRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControlFacet *AccessControlFacetFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AccessControlFacetRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AccessControlFacet.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlFacetRoleRevokedIterator{contract: _AccessControlFacet.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControlFacet *AccessControlFacetFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AccessControlFacetRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AccessControlFacet.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlFacetRoleRevoked)
				if err := _AccessControlFacet.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControlFacet *AccessControlFacetFilterer) ParseRoleRevoked(log types.Log) (*AccessControlFacetRoleRevoked, error) {
	event := new(AccessControlFacetRoleRevoked)
	if err := _AccessControlFacet.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
