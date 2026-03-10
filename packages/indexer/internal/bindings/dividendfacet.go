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

// DividendFacetMetaData contains all meta data concerning the DividendFacet contract.
var DividendFacetMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"claimDividend\",\"inputs\":[{\"name\":\"dividendId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimableAmount\",\"inputs\":[{\"name\":\"dividendId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"holder\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createDividend\",\"inputs\":[{\"name\":\"snapshotId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymentToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"dividendId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getDividend\",\"inputs\":[{\"name\":\"dividendId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"snapshotId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymentToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"claimedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"createdAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenDividends\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasClaimed\",\"inputs\":[{\"name\":\"dividendId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"holder\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DividendClaimed\",\"inputs\":[{\"name\":\"dividendId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"holder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DividendCreated\",\"inputs\":[{\"name\":\"dividendId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"snapshotId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"totalAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"paymentToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"DividendFacet__AlreadyClaimed\",\"inputs\":[{\"name\":\"dividendId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"holder\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"DividendFacet__DividendNotFound\",\"inputs\":[{\"name\":\"dividendId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"DividendFacet__HolderNotRecorded\",\"inputs\":[{\"name\":\"dividendId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"holder\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"DividendFacet__InsufficientETH\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DividendFacet__NothingToClaim\",\"inputs\":[{\"name\":\"dividendId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"holder\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"DividendFacet__SnapshotNotFound\",\"inputs\":[{\"name\":\"snapshotId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"DividendFacet__TransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DividendFacet__Unauthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DividendFacet__ZeroAmount\",\"inputs\":[]}]",
}

// DividendFacetABI is the input ABI used to generate the binding from.
// Deprecated: Use DividendFacetMetaData.ABI instead.
var DividendFacetABI = DividendFacetMetaData.ABI

// DividendFacet is an auto generated Go binding around an Ethereum contract.
type DividendFacet struct {
	DividendFacetCaller     // Read-only binding to the contract
	DividendFacetTransactor // Write-only binding to the contract
	DividendFacetFilterer   // Log filterer for contract events
}

// DividendFacetCaller is an auto generated read-only Go binding around an Ethereum contract.
type DividendFacetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DividendFacetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DividendFacetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DividendFacetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DividendFacetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DividendFacetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DividendFacetSession struct {
	Contract     *DividendFacet    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DividendFacetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DividendFacetCallerSession struct {
	Contract *DividendFacetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// DividendFacetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DividendFacetTransactorSession struct {
	Contract     *DividendFacetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// DividendFacetRaw is an auto generated low-level Go binding around an Ethereum contract.
type DividendFacetRaw struct {
	Contract *DividendFacet // Generic contract binding to access the raw methods on
}

// DividendFacetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DividendFacetCallerRaw struct {
	Contract *DividendFacetCaller // Generic read-only contract binding to access the raw methods on
}

// DividendFacetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DividendFacetTransactorRaw struct {
	Contract *DividendFacetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDividendFacet creates a new instance of DividendFacet, bound to a specific deployed contract.
func NewDividendFacet(address common.Address, backend bind.ContractBackend) (*DividendFacet, error) {
	contract, err := bindDividendFacet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DividendFacet{DividendFacetCaller: DividendFacetCaller{contract: contract}, DividendFacetTransactor: DividendFacetTransactor{contract: contract}, DividendFacetFilterer: DividendFacetFilterer{contract: contract}}, nil
}

// NewDividendFacetCaller creates a new read-only instance of DividendFacet, bound to a specific deployed contract.
func NewDividendFacetCaller(address common.Address, caller bind.ContractCaller) (*DividendFacetCaller, error) {
	contract, err := bindDividendFacet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DividendFacetCaller{contract: contract}, nil
}

// NewDividendFacetTransactor creates a new write-only instance of DividendFacet, bound to a specific deployed contract.
func NewDividendFacetTransactor(address common.Address, transactor bind.ContractTransactor) (*DividendFacetTransactor, error) {
	contract, err := bindDividendFacet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DividendFacetTransactor{contract: contract}, nil
}

// NewDividendFacetFilterer creates a new log filterer instance of DividendFacet, bound to a specific deployed contract.
func NewDividendFacetFilterer(address common.Address, filterer bind.ContractFilterer) (*DividendFacetFilterer, error) {
	contract, err := bindDividendFacet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DividendFacetFilterer{contract: contract}, nil
}

// bindDividendFacet binds a generic wrapper to an already deployed contract.
func bindDividendFacet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DividendFacetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DividendFacet *DividendFacetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DividendFacet.Contract.DividendFacetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DividendFacet *DividendFacetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DividendFacet.Contract.DividendFacetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DividendFacet *DividendFacetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DividendFacet.Contract.DividendFacetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DividendFacet *DividendFacetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DividendFacet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DividendFacet *DividendFacetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DividendFacet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DividendFacet *DividendFacetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DividendFacet.Contract.contract.Transact(opts, method, params...)
}

// ClaimableAmount is a free data retrieval call binding the contract method 0xf311df8e.
//
// Solidity: function claimableAmount(uint256 dividendId, address holder) view returns(uint256)
func (_DividendFacet *DividendFacetCaller) ClaimableAmount(opts *bind.CallOpts, dividendId *big.Int, holder common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DividendFacet.contract.Call(opts, &out, "claimableAmount", dividendId, holder)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimableAmount is a free data retrieval call binding the contract method 0xf311df8e.
//
// Solidity: function claimableAmount(uint256 dividendId, address holder) view returns(uint256)
func (_DividendFacet *DividendFacetSession) ClaimableAmount(dividendId *big.Int, holder common.Address) (*big.Int, error) {
	return _DividendFacet.Contract.ClaimableAmount(&_DividendFacet.CallOpts, dividendId, holder)
}

// ClaimableAmount is a free data retrieval call binding the contract method 0xf311df8e.
//
// Solidity: function claimableAmount(uint256 dividendId, address holder) view returns(uint256)
func (_DividendFacet *DividendFacetCallerSession) ClaimableAmount(dividendId *big.Int, holder common.Address) (*big.Int, error) {
	return _DividendFacet.Contract.ClaimableAmount(&_DividendFacet.CallOpts, dividendId, holder)
}

// GetDividend is a free data retrieval call binding the contract method 0x0ecfcaa4.
//
// Solidity: function getDividend(uint256 dividendId) view returns(uint256 snapshotId, uint256 tokenId, uint256 totalAmount, address paymentToken, uint256 claimedAmount, uint64 createdAt)
func (_DividendFacet *DividendFacetCaller) GetDividend(opts *bind.CallOpts, dividendId *big.Int) (struct {
	SnapshotId    *big.Int
	TokenId       *big.Int
	TotalAmount   *big.Int
	PaymentToken  common.Address
	ClaimedAmount *big.Int
	CreatedAt     uint64
}, error) {
	var out []interface{}
	err := _DividendFacet.contract.Call(opts, &out, "getDividend", dividendId)

	outstruct := new(struct {
		SnapshotId    *big.Int
		TokenId       *big.Int
		TotalAmount   *big.Int
		PaymentToken  common.Address
		ClaimedAmount *big.Int
		CreatedAt     uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SnapshotId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TokenId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.PaymentToken = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.ClaimedAmount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.CreatedAt = *abi.ConvertType(out[5], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetDividend is a free data retrieval call binding the contract method 0x0ecfcaa4.
//
// Solidity: function getDividend(uint256 dividendId) view returns(uint256 snapshotId, uint256 tokenId, uint256 totalAmount, address paymentToken, uint256 claimedAmount, uint64 createdAt)
func (_DividendFacet *DividendFacetSession) GetDividend(dividendId *big.Int) (struct {
	SnapshotId    *big.Int
	TokenId       *big.Int
	TotalAmount   *big.Int
	PaymentToken  common.Address
	ClaimedAmount *big.Int
	CreatedAt     uint64
}, error) {
	return _DividendFacet.Contract.GetDividend(&_DividendFacet.CallOpts, dividendId)
}

// GetDividend is a free data retrieval call binding the contract method 0x0ecfcaa4.
//
// Solidity: function getDividend(uint256 dividendId) view returns(uint256 snapshotId, uint256 tokenId, uint256 totalAmount, address paymentToken, uint256 claimedAmount, uint64 createdAt)
func (_DividendFacet *DividendFacetCallerSession) GetDividend(dividendId *big.Int) (struct {
	SnapshotId    *big.Int
	TokenId       *big.Int
	TotalAmount   *big.Int
	PaymentToken  common.Address
	ClaimedAmount *big.Int
	CreatedAt     uint64
}, error) {
	return _DividendFacet.Contract.GetDividend(&_DividendFacet.CallOpts, dividendId)
}

// GetTokenDividends is a free data retrieval call binding the contract method 0xe013da05.
//
// Solidity: function getTokenDividends(uint256 tokenId) view returns(uint256[])
func (_DividendFacet *DividendFacetCaller) GetTokenDividends(opts *bind.CallOpts, tokenId *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _DividendFacet.contract.Call(opts, &out, "getTokenDividends", tokenId)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTokenDividends is a free data retrieval call binding the contract method 0xe013da05.
//
// Solidity: function getTokenDividends(uint256 tokenId) view returns(uint256[])
func (_DividendFacet *DividendFacetSession) GetTokenDividends(tokenId *big.Int) ([]*big.Int, error) {
	return _DividendFacet.Contract.GetTokenDividends(&_DividendFacet.CallOpts, tokenId)
}

// GetTokenDividends is a free data retrieval call binding the contract method 0xe013da05.
//
// Solidity: function getTokenDividends(uint256 tokenId) view returns(uint256[])
func (_DividendFacet *DividendFacetCallerSession) GetTokenDividends(tokenId *big.Int) ([]*big.Int, error) {
	return _DividendFacet.Contract.GetTokenDividends(&_DividendFacet.CallOpts, tokenId)
}

// HasClaimed is a free data retrieval call binding the contract method 0x873f6f9e.
//
// Solidity: function hasClaimed(uint256 dividendId, address holder) view returns(bool)
func (_DividendFacet *DividendFacetCaller) HasClaimed(opts *bind.CallOpts, dividendId *big.Int, holder common.Address) (bool, error) {
	var out []interface{}
	err := _DividendFacet.contract.Call(opts, &out, "hasClaimed", dividendId, holder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasClaimed is a free data retrieval call binding the contract method 0x873f6f9e.
//
// Solidity: function hasClaimed(uint256 dividendId, address holder) view returns(bool)
func (_DividendFacet *DividendFacetSession) HasClaimed(dividendId *big.Int, holder common.Address) (bool, error) {
	return _DividendFacet.Contract.HasClaimed(&_DividendFacet.CallOpts, dividendId, holder)
}

// HasClaimed is a free data retrieval call binding the contract method 0x873f6f9e.
//
// Solidity: function hasClaimed(uint256 dividendId, address holder) view returns(bool)
func (_DividendFacet *DividendFacetCallerSession) HasClaimed(dividendId *big.Int, holder common.Address) (bool, error) {
	return _DividendFacet.Contract.HasClaimed(&_DividendFacet.CallOpts, dividendId, holder)
}

// ClaimDividend is a paid mutator transaction binding the contract method 0x9abd3572.
//
// Solidity: function claimDividend(uint256 dividendId) returns()
func (_DividendFacet *DividendFacetTransactor) ClaimDividend(opts *bind.TransactOpts, dividendId *big.Int) (*types.Transaction, error) {
	return _DividendFacet.contract.Transact(opts, "claimDividend", dividendId)
}

// ClaimDividend is a paid mutator transaction binding the contract method 0x9abd3572.
//
// Solidity: function claimDividend(uint256 dividendId) returns()
func (_DividendFacet *DividendFacetSession) ClaimDividend(dividendId *big.Int) (*types.Transaction, error) {
	return _DividendFacet.Contract.ClaimDividend(&_DividendFacet.TransactOpts, dividendId)
}

// ClaimDividend is a paid mutator transaction binding the contract method 0x9abd3572.
//
// Solidity: function claimDividend(uint256 dividendId) returns()
func (_DividendFacet *DividendFacetTransactorSession) ClaimDividend(dividendId *big.Int) (*types.Transaction, error) {
	return _DividendFacet.Contract.ClaimDividend(&_DividendFacet.TransactOpts, dividendId)
}

// CreateDividend is a paid mutator transaction binding the contract method 0xf3b9eccd.
//
// Solidity: function createDividend(uint256 snapshotId, uint256 totalAmount, address paymentToken) payable returns(uint256 dividendId)
func (_DividendFacet *DividendFacetTransactor) CreateDividend(opts *bind.TransactOpts, snapshotId *big.Int, totalAmount *big.Int, paymentToken common.Address) (*types.Transaction, error) {
	return _DividendFacet.contract.Transact(opts, "createDividend", snapshotId, totalAmount, paymentToken)
}

// CreateDividend is a paid mutator transaction binding the contract method 0xf3b9eccd.
//
// Solidity: function createDividend(uint256 snapshotId, uint256 totalAmount, address paymentToken) payable returns(uint256 dividendId)
func (_DividendFacet *DividendFacetSession) CreateDividend(snapshotId *big.Int, totalAmount *big.Int, paymentToken common.Address) (*types.Transaction, error) {
	return _DividendFacet.Contract.CreateDividend(&_DividendFacet.TransactOpts, snapshotId, totalAmount, paymentToken)
}

// CreateDividend is a paid mutator transaction binding the contract method 0xf3b9eccd.
//
// Solidity: function createDividend(uint256 snapshotId, uint256 totalAmount, address paymentToken) payable returns(uint256 dividendId)
func (_DividendFacet *DividendFacetTransactorSession) CreateDividend(snapshotId *big.Int, totalAmount *big.Int, paymentToken common.Address) (*types.Transaction, error) {
	return _DividendFacet.Contract.CreateDividend(&_DividendFacet.TransactOpts, snapshotId, totalAmount, paymentToken)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DividendFacet *DividendFacetTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DividendFacet.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DividendFacet *DividendFacetSession) Receive() (*types.Transaction, error) {
	return _DividendFacet.Contract.Receive(&_DividendFacet.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DividendFacet *DividendFacetTransactorSession) Receive() (*types.Transaction, error) {
	return _DividendFacet.Contract.Receive(&_DividendFacet.TransactOpts)
}

// DividendFacetDividendClaimedIterator is returned from FilterDividendClaimed and is used to iterate over the raw logs and unpacked data for DividendClaimed events raised by the DividendFacet contract.
type DividendFacetDividendClaimedIterator struct {
	Event *DividendFacetDividendClaimed // Event containing the contract specifics and raw log

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
func (it *DividendFacetDividendClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DividendFacetDividendClaimed)
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
		it.Event = new(DividendFacetDividendClaimed)
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
func (it *DividendFacetDividendClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DividendFacetDividendClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DividendFacetDividendClaimed represents a DividendClaimed event raised by the DividendFacet contract.
type DividendFacetDividendClaimed struct {
	DividendId *big.Int
	Holder     common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDividendClaimed is a free log retrieval operation binding the contract event 0xa1594d215a577c1905bcb0b5b186a40a0104714277441d4b2ac428c89bf7f4b1.
//
// Solidity: event DividendClaimed(uint256 indexed dividendId, address indexed holder, uint256 amount)
func (_DividendFacet *DividendFacetFilterer) FilterDividendClaimed(opts *bind.FilterOpts, dividendId []*big.Int, holder []common.Address) (*DividendFacetDividendClaimedIterator, error) {

	var dividendIdRule []interface{}
	for _, dividendIdItem := range dividendId {
		dividendIdRule = append(dividendIdRule, dividendIdItem)
	}
	var holderRule []interface{}
	for _, holderItem := range holder {
		holderRule = append(holderRule, holderItem)
	}

	logs, sub, err := _DividendFacet.contract.FilterLogs(opts, "DividendClaimed", dividendIdRule, holderRule)
	if err != nil {
		return nil, err
	}
	return &DividendFacetDividendClaimedIterator{contract: _DividendFacet.contract, event: "DividendClaimed", logs: logs, sub: sub}, nil
}

// WatchDividendClaimed is a free log subscription operation binding the contract event 0xa1594d215a577c1905bcb0b5b186a40a0104714277441d4b2ac428c89bf7f4b1.
//
// Solidity: event DividendClaimed(uint256 indexed dividendId, address indexed holder, uint256 amount)
func (_DividendFacet *DividendFacetFilterer) WatchDividendClaimed(opts *bind.WatchOpts, sink chan<- *DividendFacetDividendClaimed, dividendId []*big.Int, holder []common.Address) (event.Subscription, error) {

	var dividendIdRule []interface{}
	for _, dividendIdItem := range dividendId {
		dividendIdRule = append(dividendIdRule, dividendIdItem)
	}
	var holderRule []interface{}
	for _, holderItem := range holder {
		holderRule = append(holderRule, holderItem)
	}

	logs, sub, err := _DividendFacet.contract.WatchLogs(opts, "DividendClaimed", dividendIdRule, holderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DividendFacetDividendClaimed)
				if err := _DividendFacet.contract.UnpackLog(event, "DividendClaimed", log); err != nil {
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

// ParseDividendClaimed is a log parse operation binding the contract event 0xa1594d215a577c1905bcb0b5b186a40a0104714277441d4b2ac428c89bf7f4b1.
//
// Solidity: event DividendClaimed(uint256 indexed dividendId, address indexed holder, uint256 amount)
func (_DividendFacet *DividendFacetFilterer) ParseDividendClaimed(log types.Log) (*DividendFacetDividendClaimed, error) {
	event := new(DividendFacetDividendClaimed)
	if err := _DividendFacet.contract.UnpackLog(event, "DividendClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DividendFacetDividendCreatedIterator is returned from FilterDividendCreated and is used to iterate over the raw logs and unpacked data for DividendCreated events raised by the DividendFacet contract.
type DividendFacetDividendCreatedIterator struct {
	Event *DividendFacetDividendCreated // Event containing the contract specifics and raw log

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
func (it *DividendFacetDividendCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DividendFacetDividendCreated)
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
		it.Event = new(DividendFacetDividendCreated)
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
func (it *DividendFacetDividendCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DividendFacetDividendCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DividendFacetDividendCreated represents a DividendCreated event raised by the DividendFacet contract.
type DividendFacetDividendCreated struct {
	DividendId   *big.Int
	TokenId      *big.Int
	SnapshotId   *big.Int
	TotalAmount  *big.Int
	PaymentToken common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDividendCreated is a free log retrieval operation binding the contract event 0x56e70de6ebdfdae1873bfa666f655b7d9e62aed0c54aac6f7546a34e423c3e9e.
//
// Solidity: event DividendCreated(uint256 indexed dividendId, uint256 indexed tokenId, uint256 indexed snapshotId, uint256 totalAmount, address paymentToken)
func (_DividendFacet *DividendFacetFilterer) FilterDividendCreated(opts *bind.FilterOpts, dividendId []*big.Int, tokenId []*big.Int, snapshotId []*big.Int) (*DividendFacetDividendCreatedIterator, error) {

	var dividendIdRule []interface{}
	for _, dividendIdItem := range dividendId {
		dividendIdRule = append(dividendIdRule, dividendIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var snapshotIdRule []interface{}
	for _, snapshotIdItem := range snapshotId {
		snapshotIdRule = append(snapshotIdRule, snapshotIdItem)
	}

	logs, sub, err := _DividendFacet.contract.FilterLogs(opts, "DividendCreated", dividendIdRule, tokenIdRule, snapshotIdRule)
	if err != nil {
		return nil, err
	}
	return &DividendFacetDividendCreatedIterator{contract: _DividendFacet.contract, event: "DividendCreated", logs: logs, sub: sub}, nil
}

// WatchDividendCreated is a free log subscription operation binding the contract event 0x56e70de6ebdfdae1873bfa666f655b7d9e62aed0c54aac6f7546a34e423c3e9e.
//
// Solidity: event DividendCreated(uint256 indexed dividendId, uint256 indexed tokenId, uint256 indexed snapshotId, uint256 totalAmount, address paymentToken)
func (_DividendFacet *DividendFacetFilterer) WatchDividendCreated(opts *bind.WatchOpts, sink chan<- *DividendFacetDividendCreated, dividendId []*big.Int, tokenId []*big.Int, snapshotId []*big.Int) (event.Subscription, error) {

	var dividendIdRule []interface{}
	for _, dividendIdItem := range dividendId {
		dividendIdRule = append(dividendIdRule, dividendIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var snapshotIdRule []interface{}
	for _, snapshotIdItem := range snapshotId {
		snapshotIdRule = append(snapshotIdRule, snapshotIdItem)
	}

	logs, sub, err := _DividendFacet.contract.WatchLogs(opts, "DividendCreated", dividendIdRule, tokenIdRule, snapshotIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DividendFacetDividendCreated)
				if err := _DividendFacet.contract.UnpackLog(event, "DividendCreated", log); err != nil {
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

// ParseDividendCreated is a log parse operation binding the contract event 0x56e70de6ebdfdae1873bfa666f655b7d9e62aed0c54aac6f7546a34e423c3e9e.
//
// Solidity: event DividendCreated(uint256 indexed dividendId, uint256 indexed tokenId, uint256 indexed snapshotId, uint256 totalAmount, address paymentToken)
func (_DividendFacet *DividendFacetFilterer) ParseDividendCreated(log types.Log) (*DividendFacetDividendCreated, error) {
	event := new(DividendFacetDividendCreated)
	if err := _DividendFacet.contract.UnpackLog(event, "DividendCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
