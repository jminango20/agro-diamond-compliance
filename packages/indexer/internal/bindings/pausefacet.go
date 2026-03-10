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

// PauseFacetMetaData contains all meta data concerning the PauseFacet contract.
var PauseFacetMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"isAssetPaused\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isProtocolPaused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pauseAsset\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauseAsset\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauseProtocol\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AssetPaused\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"by\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AssetUnpaused\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"by\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EmergencyPause\",\"inputs\":[{\"name\":\"triggeredBy\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProtocolUnpaused\",\"inputs\":[{\"name\":\"by\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"LibDiamond__OnlyOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PauseFacet__AlreadyPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PauseFacet__AssetAlreadyPaused\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"PauseFacet__AssetNotPaused\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"PauseFacet__AssetNotRegistered\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"PauseFacet__NotPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PauseFacet__Unauthorized\",\"inputs\":[]}]",
}

// PauseFacetABI is the input ABI used to generate the binding from.
// Deprecated: Use PauseFacetMetaData.ABI instead.
var PauseFacetABI = PauseFacetMetaData.ABI

// PauseFacet is an auto generated Go binding around an Ethereum contract.
type PauseFacet struct {
	PauseFacetCaller     // Read-only binding to the contract
	PauseFacetTransactor // Write-only binding to the contract
	PauseFacetFilterer   // Log filterer for contract events
}

// PauseFacetCaller is an auto generated read-only Go binding around an Ethereum contract.
type PauseFacetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PauseFacetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PauseFacetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PauseFacetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PauseFacetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PauseFacetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PauseFacetSession struct {
	Contract     *PauseFacet       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PauseFacetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PauseFacetCallerSession struct {
	Contract *PauseFacetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PauseFacetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PauseFacetTransactorSession struct {
	Contract     *PauseFacetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PauseFacetRaw is an auto generated low-level Go binding around an Ethereum contract.
type PauseFacetRaw struct {
	Contract *PauseFacet // Generic contract binding to access the raw methods on
}

// PauseFacetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PauseFacetCallerRaw struct {
	Contract *PauseFacetCaller // Generic read-only contract binding to access the raw methods on
}

// PauseFacetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PauseFacetTransactorRaw struct {
	Contract *PauseFacetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPauseFacet creates a new instance of PauseFacet, bound to a specific deployed contract.
func NewPauseFacet(address common.Address, backend bind.ContractBackend) (*PauseFacet, error) {
	contract, err := bindPauseFacet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PauseFacet{PauseFacetCaller: PauseFacetCaller{contract: contract}, PauseFacetTransactor: PauseFacetTransactor{contract: contract}, PauseFacetFilterer: PauseFacetFilterer{contract: contract}}, nil
}

// NewPauseFacetCaller creates a new read-only instance of PauseFacet, bound to a specific deployed contract.
func NewPauseFacetCaller(address common.Address, caller bind.ContractCaller) (*PauseFacetCaller, error) {
	contract, err := bindPauseFacet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PauseFacetCaller{contract: contract}, nil
}

// NewPauseFacetTransactor creates a new write-only instance of PauseFacet, bound to a specific deployed contract.
func NewPauseFacetTransactor(address common.Address, transactor bind.ContractTransactor) (*PauseFacetTransactor, error) {
	contract, err := bindPauseFacet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PauseFacetTransactor{contract: contract}, nil
}

// NewPauseFacetFilterer creates a new log filterer instance of PauseFacet, bound to a specific deployed contract.
func NewPauseFacetFilterer(address common.Address, filterer bind.ContractFilterer) (*PauseFacetFilterer, error) {
	contract, err := bindPauseFacet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PauseFacetFilterer{contract: contract}, nil
}

// bindPauseFacet binds a generic wrapper to an already deployed contract.
func bindPauseFacet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PauseFacetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PauseFacet *PauseFacetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PauseFacet.Contract.PauseFacetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PauseFacet *PauseFacetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PauseFacet.Contract.PauseFacetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PauseFacet *PauseFacetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PauseFacet.Contract.PauseFacetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PauseFacet *PauseFacetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PauseFacet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PauseFacet *PauseFacetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PauseFacet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PauseFacet *PauseFacetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PauseFacet.Contract.contract.Transact(opts, method, params...)
}

// IsAssetPaused is a free data retrieval call binding the contract method 0x1f50b2cb.
//
// Solidity: function isAssetPaused(uint256 tokenId) view returns(bool)
func (_PauseFacet *PauseFacetCaller) IsAssetPaused(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _PauseFacet.contract.Call(opts, &out, "isAssetPaused", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAssetPaused is a free data retrieval call binding the contract method 0x1f50b2cb.
//
// Solidity: function isAssetPaused(uint256 tokenId) view returns(bool)
func (_PauseFacet *PauseFacetSession) IsAssetPaused(tokenId *big.Int) (bool, error) {
	return _PauseFacet.Contract.IsAssetPaused(&_PauseFacet.CallOpts, tokenId)
}

// IsAssetPaused is a free data retrieval call binding the contract method 0x1f50b2cb.
//
// Solidity: function isAssetPaused(uint256 tokenId) view returns(bool)
func (_PauseFacet *PauseFacetCallerSession) IsAssetPaused(tokenId *big.Int) (bool, error) {
	return _PauseFacet.Contract.IsAssetPaused(&_PauseFacet.CallOpts, tokenId)
}

// IsProtocolPaused is a free data retrieval call binding the contract method 0xdac88561.
//
// Solidity: function isProtocolPaused() view returns(bool)
func (_PauseFacet *PauseFacetCaller) IsProtocolPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PauseFacet.contract.Call(opts, &out, "isProtocolPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsProtocolPaused is a free data retrieval call binding the contract method 0xdac88561.
//
// Solidity: function isProtocolPaused() view returns(bool)
func (_PauseFacet *PauseFacetSession) IsProtocolPaused() (bool, error) {
	return _PauseFacet.Contract.IsProtocolPaused(&_PauseFacet.CallOpts)
}

// IsProtocolPaused is a free data retrieval call binding the contract method 0xdac88561.
//
// Solidity: function isProtocolPaused() view returns(bool)
func (_PauseFacet *PauseFacetCallerSession) IsProtocolPaused() (bool, error) {
	return _PauseFacet.Contract.IsProtocolPaused(&_PauseFacet.CallOpts)
}

// PauseAsset is a paid mutator transaction binding the contract method 0x6cdbf1f9.
//
// Solidity: function pauseAsset(uint256 tokenId) returns()
func (_PauseFacet *PauseFacetTransactor) PauseAsset(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _PauseFacet.contract.Transact(opts, "pauseAsset", tokenId)
}

// PauseAsset is a paid mutator transaction binding the contract method 0x6cdbf1f9.
//
// Solidity: function pauseAsset(uint256 tokenId) returns()
func (_PauseFacet *PauseFacetSession) PauseAsset(tokenId *big.Int) (*types.Transaction, error) {
	return _PauseFacet.Contract.PauseAsset(&_PauseFacet.TransactOpts, tokenId)
}

// PauseAsset is a paid mutator transaction binding the contract method 0x6cdbf1f9.
//
// Solidity: function pauseAsset(uint256 tokenId) returns()
func (_PauseFacet *PauseFacetTransactorSession) PauseAsset(tokenId *big.Int) (*types.Transaction, error) {
	return _PauseFacet.Contract.PauseAsset(&_PauseFacet.TransactOpts, tokenId)
}

// PauseProtocol is a paid mutator transaction binding the contract method 0xdbf62489.
//
// Solidity: function pauseProtocol() returns()
func (_PauseFacet *PauseFacetTransactor) PauseProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PauseFacet.contract.Transact(opts, "pauseProtocol")
}

// PauseProtocol is a paid mutator transaction binding the contract method 0xdbf62489.
//
// Solidity: function pauseProtocol() returns()
func (_PauseFacet *PauseFacetSession) PauseProtocol() (*types.Transaction, error) {
	return _PauseFacet.Contract.PauseProtocol(&_PauseFacet.TransactOpts)
}

// PauseProtocol is a paid mutator transaction binding the contract method 0xdbf62489.
//
// Solidity: function pauseProtocol() returns()
func (_PauseFacet *PauseFacetTransactorSession) PauseProtocol() (*types.Transaction, error) {
	return _PauseFacet.Contract.PauseProtocol(&_PauseFacet.TransactOpts)
}

// UnpauseAsset is a paid mutator transaction binding the contract method 0x27312771.
//
// Solidity: function unpauseAsset(uint256 tokenId) returns()
func (_PauseFacet *PauseFacetTransactor) UnpauseAsset(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _PauseFacet.contract.Transact(opts, "unpauseAsset", tokenId)
}

// UnpauseAsset is a paid mutator transaction binding the contract method 0x27312771.
//
// Solidity: function unpauseAsset(uint256 tokenId) returns()
func (_PauseFacet *PauseFacetSession) UnpauseAsset(tokenId *big.Int) (*types.Transaction, error) {
	return _PauseFacet.Contract.UnpauseAsset(&_PauseFacet.TransactOpts, tokenId)
}

// UnpauseAsset is a paid mutator transaction binding the contract method 0x27312771.
//
// Solidity: function unpauseAsset(uint256 tokenId) returns()
func (_PauseFacet *PauseFacetTransactorSession) UnpauseAsset(tokenId *big.Int) (*types.Transaction, error) {
	return _PauseFacet.Contract.UnpauseAsset(&_PauseFacet.TransactOpts, tokenId)
}

// UnpauseProtocol is a paid mutator transaction binding the contract method 0xf93b6be5.
//
// Solidity: function unpauseProtocol() returns()
func (_PauseFacet *PauseFacetTransactor) UnpauseProtocol(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PauseFacet.contract.Transact(opts, "unpauseProtocol")
}

// UnpauseProtocol is a paid mutator transaction binding the contract method 0xf93b6be5.
//
// Solidity: function unpauseProtocol() returns()
func (_PauseFacet *PauseFacetSession) UnpauseProtocol() (*types.Transaction, error) {
	return _PauseFacet.Contract.UnpauseProtocol(&_PauseFacet.TransactOpts)
}

// UnpauseProtocol is a paid mutator transaction binding the contract method 0xf93b6be5.
//
// Solidity: function unpauseProtocol() returns()
func (_PauseFacet *PauseFacetTransactorSession) UnpauseProtocol() (*types.Transaction, error) {
	return _PauseFacet.Contract.UnpauseProtocol(&_PauseFacet.TransactOpts)
}

// PauseFacetAssetPausedIterator is returned from FilterAssetPaused and is used to iterate over the raw logs and unpacked data for AssetPaused events raised by the PauseFacet contract.
type PauseFacetAssetPausedIterator struct {
	Event *PauseFacetAssetPaused // Event containing the contract specifics and raw log

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
func (it *PauseFacetAssetPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PauseFacetAssetPaused)
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
		it.Event = new(PauseFacetAssetPaused)
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
func (it *PauseFacetAssetPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PauseFacetAssetPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PauseFacetAssetPaused represents a AssetPaused event raised by the PauseFacet contract.
type PauseFacetAssetPaused struct {
	TokenId *big.Int
	By      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAssetPaused is a free log retrieval operation binding the contract event 0xf9cc4af735657d0f63ceb0d850883a0608e74937c3815446bd72d81e0230d5d7.
//
// Solidity: event AssetPaused(uint256 indexed tokenId, address indexed by)
func (_PauseFacet *PauseFacetFilterer) FilterAssetPaused(opts *bind.FilterOpts, tokenId []*big.Int, by []common.Address) (*PauseFacetAssetPausedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _PauseFacet.contract.FilterLogs(opts, "AssetPaused", tokenIdRule, byRule)
	if err != nil {
		return nil, err
	}
	return &PauseFacetAssetPausedIterator{contract: _PauseFacet.contract, event: "AssetPaused", logs: logs, sub: sub}, nil
}

// WatchAssetPaused is a free log subscription operation binding the contract event 0xf9cc4af735657d0f63ceb0d850883a0608e74937c3815446bd72d81e0230d5d7.
//
// Solidity: event AssetPaused(uint256 indexed tokenId, address indexed by)
func (_PauseFacet *PauseFacetFilterer) WatchAssetPaused(opts *bind.WatchOpts, sink chan<- *PauseFacetAssetPaused, tokenId []*big.Int, by []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _PauseFacet.contract.WatchLogs(opts, "AssetPaused", tokenIdRule, byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PauseFacetAssetPaused)
				if err := _PauseFacet.contract.UnpackLog(event, "AssetPaused", log); err != nil {
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

// ParseAssetPaused is a log parse operation binding the contract event 0xf9cc4af735657d0f63ceb0d850883a0608e74937c3815446bd72d81e0230d5d7.
//
// Solidity: event AssetPaused(uint256 indexed tokenId, address indexed by)
func (_PauseFacet *PauseFacetFilterer) ParseAssetPaused(log types.Log) (*PauseFacetAssetPaused, error) {
	event := new(PauseFacetAssetPaused)
	if err := _PauseFacet.contract.UnpackLog(event, "AssetPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PauseFacetAssetUnpausedIterator is returned from FilterAssetUnpaused and is used to iterate over the raw logs and unpacked data for AssetUnpaused events raised by the PauseFacet contract.
type PauseFacetAssetUnpausedIterator struct {
	Event *PauseFacetAssetUnpaused // Event containing the contract specifics and raw log

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
func (it *PauseFacetAssetUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PauseFacetAssetUnpaused)
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
		it.Event = new(PauseFacetAssetUnpaused)
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
func (it *PauseFacetAssetUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PauseFacetAssetUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PauseFacetAssetUnpaused represents a AssetUnpaused event raised by the PauseFacet contract.
type PauseFacetAssetUnpaused struct {
	TokenId *big.Int
	By      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAssetUnpaused is a free log retrieval operation binding the contract event 0x7e88a201c7bdfee0f59dbb3ec2d36ad546c0df20d7e0876af7c95f3912808426.
//
// Solidity: event AssetUnpaused(uint256 indexed tokenId, address indexed by)
func (_PauseFacet *PauseFacetFilterer) FilterAssetUnpaused(opts *bind.FilterOpts, tokenId []*big.Int, by []common.Address) (*PauseFacetAssetUnpausedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _PauseFacet.contract.FilterLogs(opts, "AssetUnpaused", tokenIdRule, byRule)
	if err != nil {
		return nil, err
	}
	return &PauseFacetAssetUnpausedIterator{contract: _PauseFacet.contract, event: "AssetUnpaused", logs: logs, sub: sub}, nil
}

// WatchAssetUnpaused is a free log subscription operation binding the contract event 0x7e88a201c7bdfee0f59dbb3ec2d36ad546c0df20d7e0876af7c95f3912808426.
//
// Solidity: event AssetUnpaused(uint256 indexed tokenId, address indexed by)
func (_PauseFacet *PauseFacetFilterer) WatchAssetUnpaused(opts *bind.WatchOpts, sink chan<- *PauseFacetAssetUnpaused, tokenId []*big.Int, by []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _PauseFacet.contract.WatchLogs(opts, "AssetUnpaused", tokenIdRule, byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PauseFacetAssetUnpaused)
				if err := _PauseFacet.contract.UnpackLog(event, "AssetUnpaused", log); err != nil {
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

// ParseAssetUnpaused is a log parse operation binding the contract event 0x7e88a201c7bdfee0f59dbb3ec2d36ad546c0df20d7e0876af7c95f3912808426.
//
// Solidity: event AssetUnpaused(uint256 indexed tokenId, address indexed by)
func (_PauseFacet *PauseFacetFilterer) ParseAssetUnpaused(log types.Log) (*PauseFacetAssetUnpaused, error) {
	event := new(PauseFacetAssetUnpaused)
	if err := _PauseFacet.contract.UnpackLog(event, "AssetUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PauseFacetEmergencyPauseIterator is returned from FilterEmergencyPause and is used to iterate over the raw logs and unpacked data for EmergencyPause events raised by the PauseFacet contract.
type PauseFacetEmergencyPauseIterator struct {
	Event *PauseFacetEmergencyPause // Event containing the contract specifics and raw log

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
func (it *PauseFacetEmergencyPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PauseFacetEmergencyPause)
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
		it.Event = new(PauseFacetEmergencyPause)
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
func (it *PauseFacetEmergencyPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PauseFacetEmergencyPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PauseFacetEmergencyPause represents a EmergencyPause event raised by the PauseFacet contract.
type PauseFacetEmergencyPause struct {
	TriggeredBy common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEmergencyPause is a free log retrieval operation binding the contract event 0x7c83004a7e59a8ea03b200186c4dda29a4e144d9844d63dbc1a09acf7dfcd485.
//
// Solidity: event EmergencyPause(address indexed triggeredBy)
func (_PauseFacet *PauseFacetFilterer) FilterEmergencyPause(opts *bind.FilterOpts, triggeredBy []common.Address) (*PauseFacetEmergencyPauseIterator, error) {

	var triggeredByRule []interface{}
	for _, triggeredByItem := range triggeredBy {
		triggeredByRule = append(triggeredByRule, triggeredByItem)
	}

	logs, sub, err := _PauseFacet.contract.FilterLogs(opts, "EmergencyPause", triggeredByRule)
	if err != nil {
		return nil, err
	}
	return &PauseFacetEmergencyPauseIterator{contract: _PauseFacet.contract, event: "EmergencyPause", logs: logs, sub: sub}, nil
}

// WatchEmergencyPause is a free log subscription operation binding the contract event 0x7c83004a7e59a8ea03b200186c4dda29a4e144d9844d63dbc1a09acf7dfcd485.
//
// Solidity: event EmergencyPause(address indexed triggeredBy)
func (_PauseFacet *PauseFacetFilterer) WatchEmergencyPause(opts *bind.WatchOpts, sink chan<- *PauseFacetEmergencyPause, triggeredBy []common.Address) (event.Subscription, error) {

	var triggeredByRule []interface{}
	for _, triggeredByItem := range triggeredBy {
		triggeredByRule = append(triggeredByRule, triggeredByItem)
	}

	logs, sub, err := _PauseFacet.contract.WatchLogs(opts, "EmergencyPause", triggeredByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PauseFacetEmergencyPause)
				if err := _PauseFacet.contract.UnpackLog(event, "EmergencyPause", log); err != nil {
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

// ParseEmergencyPause is a log parse operation binding the contract event 0x7c83004a7e59a8ea03b200186c4dda29a4e144d9844d63dbc1a09acf7dfcd485.
//
// Solidity: event EmergencyPause(address indexed triggeredBy)
func (_PauseFacet *PauseFacetFilterer) ParseEmergencyPause(log types.Log) (*PauseFacetEmergencyPause, error) {
	event := new(PauseFacetEmergencyPause)
	if err := _PauseFacet.contract.UnpackLog(event, "EmergencyPause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PauseFacetProtocolUnpausedIterator is returned from FilterProtocolUnpaused and is used to iterate over the raw logs and unpacked data for ProtocolUnpaused events raised by the PauseFacet contract.
type PauseFacetProtocolUnpausedIterator struct {
	Event *PauseFacetProtocolUnpaused // Event containing the contract specifics and raw log

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
func (it *PauseFacetProtocolUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PauseFacetProtocolUnpaused)
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
		it.Event = new(PauseFacetProtocolUnpaused)
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
func (it *PauseFacetProtocolUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PauseFacetProtocolUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PauseFacetProtocolUnpaused represents a ProtocolUnpaused event raised by the PauseFacet contract.
type PauseFacetProtocolUnpaused struct {
	By  common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterProtocolUnpaused is a free log retrieval operation binding the contract event 0x2d553fb579456929d17c4cf0dc43b2f2e8e5a8f917cadcc8a5f18e9c4010b75a.
//
// Solidity: event ProtocolUnpaused(address indexed by)
func (_PauseFacet *PauseFacetFilterer) FilterProtocolUnpaused(opts *bind.FilterOpts, by []common.Address) (*PauseFacetProtocolUnpausedIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _PauseFacet.contract.FilterLogs(opts, "ProtocolUnpaused", byRule)
	if err != nil {
		return nil, err
	}
	return &PauseFacetProtocolUnpausedIterator{contract: _PauseFacet.contract, event: "ProtocolUnpaused", logs: logs, sub: sub}, nil
}

// WatchProtocolUnpaused is a free log subscription operation binding the contract event 0x2d553fb579456929d17c4cf0dc43b2f2e8e5a8f917cadcc8a5f18e9c4010b75a.
//
// Solidity: event ProtocolUnpaused(address indexed by)
func (_PauseFacet *PauseFacetFilterer) WatchProtocolUnpaused(opts *bind.WatchOpts, sink chan<- *PauseFacetProtocolUnpaused, by []common.Address) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _PauseFacet.contract.WatchLogs(opts, "ProtocolUnpaused", byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PauseFacetProtocolUnpaused)
				if err := _PauseFacet.contract.UnpackLog(event, "ProtocolUnpaused", log); err != nil {
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

// ParseProtocolUnpaused is a log parse operation binding the contract event 0x2d553fb579456929d17c4cf0dc43b2f2e8e5a8f917cadcc8a5f18e9c4010b75a.
//
// Solidity: event ProtocolUnpaused(address indexed by)
func (_PauseFacet *PauseFacetFilterer) ParseProtocolUnpaused(log types.Log) (*PauseFacetProtocolUnpaused, error) {
	event := new(PauseFacetProtocolUnpaused)
	if err := _PauseFacet.contract.UnpackLog(event, "ProtocolUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
