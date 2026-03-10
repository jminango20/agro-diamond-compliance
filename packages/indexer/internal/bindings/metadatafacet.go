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

// MetadataFacetMetaData contains all meta data concerning the MetadataFacet contract.
var MetadataFacetMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"allowedCountries\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"issuer\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supplyCap\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenInfo\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"name_\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol_\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"uri_\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"totalSupply_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"supplyCap_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"holderCount_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"issuer_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"paused_\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"uri\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"URI\",\"inputs\":[{\"name\":\"value\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// MetadataFacetABI is the input ABI used to generate the binding from.
// Deprecated: Use MetadataFacetMetaData.ABI instead.
var MetadataFacetABI = MetadataFacetMetaData.ABI

// MetadataFacet is an auto generated Go binding around an Ethereum contract.
type MetadataFacet struct {
	MetadataFacetCaller     // Read-only binding to the contract
	MetadataFacetTransactor // Write-only binding to the contract
	MetadataFacetFilterer   // Log filterer for contract events
}

// MetadataFacetCaller is an auto generated read-only Go binding around an Ethereum contract.
type MetadataFacetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetadataFacetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MetadataFacetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetadataFacetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MetadataFacetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetadataFacetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MetadataFacetSession struct {
	Contract     *MetadataFacet    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MetadataFacetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MetadataFacetCallerSession struct {
	Contract *MetadataFacetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MetadataFacetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MetadataFacetTransactorSession struct {
	Contract     *MetadataFacetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MetadataFacetRaw is an auto generated low-level Go binding around an Ethereum contract.
type MetadataFacetRaw struct {
	Contract *MetadataFacet // Generic contract binding to access the raw methods on
}

// MetadataFacetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MetadataFacetCallerRaw struct {
	Contract *MetadataFacetCaller // Generic read-only contract binding to access the raw methods on
}

// MetadataFacetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MetadataFacetTransactorRaw struct {
	Contract *MetadataFacetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMetadataFacet creates a new instance of MetadataFacet, bound to a specific deployed contract.
func NewMetadataFacet(address common.Address, backend bind.ContractBackend) (*MetadataFacet, error) {
	contract, err := bindMetadataFacet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MetadataFacet{MetadataFacetCaller: MetadataFacetCaller{contract: contract}, MetadataFacetTransactor: MetadataFacetTransactor{contract: contract}, MetadataFacetFilterer: MetadataFacetFilterer{contract: contract}}, nil
}

// NewMetadataFacetCaller creates a new read-only instance of MetadataFacet, bound to a specific deployed contract.
func NewMetadataFacetCaller(address common.Address, caller bind.ContractCaller) (*MetadataFacetCaller, error) {
	contract, err := bindMetadataFacet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MetadataFacetCaller{contract: contract}, nil
}

// NewMetadataFacetTransactor creates a new write-only instance of MetadataFacet, bound to a specific deployed contract.
func NewMetadataFacetTransactor(address common.Address, transactor bind.ContractTransactor) (*MetadataFacetTransactor, error) {
	contract, err := bindMetadataFacet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MetadataFacetTransactor{contract: contract}, nil
}

// NewMetadataFacetFilterer creates a new log filterer instance of MetadataFacet, bound to a specific deployed contract.
func NewMetadataFacetFilterer(address common.Address, filterer bind.ContractFilterer) (*MetadataFacetFilterer, error) {
	contract, err := bindMetadataFacet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MetadataFacetFilterer{contract: contract}, nil
}

// bindMetadataFacet binds a generic wrapper to an already deployed contract.
func bindMetadataFacet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MetadataFacetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetadataFacet *MetadataFacetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MetadataFacet.Contract.MetadataFacetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetadataFacet *MetadataFacetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetadataFacet.Contract.MetadataFacetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetadataFacet *MetadataFacetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetadataFacet.Contract.MetadataFacetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetadataFacet *MetadataFacetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MetadataFacet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetadataFacet *MetadataFacetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetadataFacet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetadataFacet *MetadataFacetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetadataFacet.Contract.contract.Transact(opts, method, params...)
}

// AllowedCountries is a free data retrieval call binding the contract method 0x42181b8e.
//
// Solidity: function allowedCountries(uint256 tokenId) view returns(uint16[])
func (_MetadataFacet *MetadataFacetCaller) AllowedCountries(opts *bind.CallOpts, tokenId *big.Int) ([]uint16, error) {
	var out []interface{}
	err := _MetadataFacet.contract.Call(opts, &out, "allowedCountries", tokenId)

	if err != nil {
		return *new([]uint16), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint16)).(*[]uint16)

	return out0, err

}

// AllowedCountries is a free data retrieval call binding the contract method 0x42181b8e.
//
// Solidity: function allowedCountries(uint256 tokenId) view returns(uint16[])
func (_MetadataFacet *MetadataFacetSession) AllowedCountries(tokenId *big.Int) ([]uint16, error) {
	return _MetadataFacet.Contract.AllowedCountries(&_MetadataFacet.CallOpts, tokenId)
}

// AllowedCountries is a free data retrieval call binding the contract method 0x42181b8e.
//
// Solidity: function allowedCountries(uint256 tokenId) view returns(uint16[])
func (_MetadataFacet *MetadataFacetCallerSession) AllowedCountries(tokenId *big.Int) ([]uint16, error) {
	return _MetadataFacet.Contract.AllowedCountries(&_MetadataFacet.CallOpts, tokenId)
}

// Issuer is a free data retrieval call binding the contract method 0x0068fe8b.
//
// Solidity: function issuer(uint256 tokenId) view returns(address)
func (_MetadataFacet *MetadataFacetCaller) Issuer(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MetadataFacet.contract.Call(opts, &out, "issuer", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Issuer is a free data retrieval call binding the contract method 0x0068fe8b.
//
// Solidity: function issuer(uint256 tokenId) view returns(address)
func (_MetadataFacet *MetadataFacetSession) Issuer(tokenId *big.Int) (common.Address, error) {
	return _MetadataFacet.Contract.Issuer(&_MetadataFacet.CallOpts, tokenId)
}

// Issuer is a free data retrieval call binding the contract method 0x0068fe8b.
//
// Solidity: function issuer(uint256 tokenId) view returns(address)
func (_MetadataFacet *MetadataFacetCallerSession) Issuer(tokenId *big.Int) (common.Address, error) {
	return _MetadataFacet.Contract.Issuer(&_MetadataFacet.CallOpts, tokenId)
}

// Name is a free data retrieval call binding the contract method 0x00ad800c.
//
// Solidity: function name(uint256 tokenId) view returns(string)
func (_MetadataFacet *MetadataFacetCaller) Name(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _MetadataFacet.contract.Call(opts, &out, "name", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x00ad800c.
//
// Solidity: function name(uint256 tokenId) view returns(string)
func (_MetadataFacet *MetadataFacetSession) Name(tokenId *big.Int) (string, error) {
	return _MetadataFacet.Contract.Name(&_MetadataFacet.CallOpts, tokenId)
}

// Name is a free data retrieval call binding the contract method 0x00ad800c.
//
// Solidity: function name(uint256 tokenId) view returns(string)
func (_MetadataFacet *MetadataFacetCallerSession) Name(tokenId *big.Int) (string, error) {
	return _MetadataFacet.Contract.Name(&_MetadataFacet.CallOpts, tokenId)
}

// SupplyCap is a free data retrieval call binding the contract method 0x40b71c40.
//
// Solidity: function supplyCap(uint256 tokenId) view returns(uint256)
func (_MetadataFacet *MetadataFacetCaller) SupplyCap(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MetadataFacet.contract.Call(opts, &out, "supplyCap", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyCap is a free data retrieval call binding the contract method 0x40b71c40.
//
// Solidity: function supplyCap(uint256 tokenId) view returns(uint256)
func (_MetadataFacet *MetadataFacetSession) SupplyCap(tokenId *big.Int) (*big.Int, error) {
	return _MetadataFacet.Contract.SupplyCap(&_MetadataFacet.CallOpts, tokenId)
}

// SupplyCap is a free data retrieval call binding the contract method 0x40b71c40.
//
// Solidity: function supplyCap(uint256 tokenId) view returns(uint256)
func (_MetadataFacet *MetadataFacetCallerSession) SupplyCap(tokenId *big.Int) (*big.Int, error) {
	return _MetadataFacet.Contract.SupplyCap(&_MetadataFacet.CallOpts, tokenId)
}

// Symbol is a free data retrieval call binding the contract method 0x4e41a1fb.
//
// Solidity: function symbol(uint256 tokenId) view returns(string)
func (_MetadataFacet *MetadataFacetCaller) Symbol(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _MetadataFacet.contract.Call(opts, &out, "symbol", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x4e41a1fb.
//
// Solidity: function symbol(uint256 tokenId) view returns(string)
func (_MetadataFacet *MetadataFacetSession) Symbol(tokenId *big.Int) (string, error) {
	return _MetadataFacet.Contract.Symbol(&_MetadataFacet.CallOpts, tokenId)
}

// Symbol is a free data retrieval call binding the contract method 0x4e41a1fb.
//
// Solidity: function symbol(uint256 tokenId) view returns(string)
func (_MetadataFacet *MetadataFacetCallerSession) Symbol(tokenId *big.Int) (string, error) {
	return _MetadataFacet.Contract.Symbol(&_MetadataFacet.CallOpts, tokenId)
}

// TokenInfo is a free data retrieval call binding the contract method 0xcc33c875.
//
// Solidity: function tokenInfo(uint256 tokenId) view returns(string name_, string symbol_, string uri_, uint256 totalSupply_, uint256 supplyCap_, uint256 holderCount_, address issuer_, bool paused_)
func (_MetadataFacet *MetadataFacetCaller) TokenInfo(opts *bind.CallOpts, tokenId *big.Int) (struct {
	Name        string
	Symbol      string
	Uri         string
	TotalSupply *big.Int
	SupplyCap   *big.Int
	HolderCount *big.Int
	Issuer      common.Address
	Paused      bool
}, error) {
	var out []interface{}
	err := _MetadataFacet.contract.Call(opts, &out, "tokenInfo", tokenId)

	outstruct := new(struct {
		Name        string
		Symbol      string
		Uri         string
		TotalSupply *big.Int
		SupplyCap   *big.Int
		HolderCount *big.Int
		Issuer      common.Address
		Paused      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Symbol = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Uri = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.TotalSupply = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.SupplyCap = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.HolderCount = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Issuer = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.Paused = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// TokenInfo is a free data retrieval call binding the contract method 0xcc33c875.
//
// Solidity: function tokenInfo(uint256 tokenId) view returns(string name_, string symbol_, string uri_, uint256 totalSupply_, uint256 supplyCap_, uint256 holderCount_, address issuer_, bool paused_)
func (_MetadataFacet *MetadataFacetSession) TokenInfo(tokenId *big.Int) (struct {
	Name        string
	Symbol      string
	Uri         string
	TotalSupply *big.Int
	SupplyCap   *big.Int
	HolderCount *big.Int
	Issuer      common.Address
	Paused      bool
}, error) {
	return _MetadataFacet.Contract.TokenInfo(&_MetadataFacet.CallOpts, tokenId)
}

// TokenInfo is a free data retrieval call binding the contract method 0xcc33c875.
//
// Solidity: function tokenInfo(uint256 tokenId) view returns(string name_, string symbol_, string uri_, uint256 totalSupply_, uint256 supplyCap_, uint256 holderCount_, address issuer_, bool paused_)
func (_MetadataFacet *MetadataFacetCallerSession) TokenInfo(tokenId *big.Int) (struct {
	Name        string
	Symbol      string
	Uri         string
	TotalSupply *big.Int
	SupplyCap   *big.Int
	HolderCount *big.Int
	Issuer      common.Address
	Paused      bool
}, error) {
	return _MetadataFacet.Contract.TokenInfo(&_MetadataFacet.CallOpts, tokenId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 id) view returns(string)
func (_MetadataFacet *MetadataFacetCaller) Uri(opts *bind.CallOpts, id *big.Int) (string, error) {
	var out []interface{}
	err := _MetadataFacet.contract.Call(opts, &out, "uri", id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 id) view returns(string)
func (_MetadataFacet *MetadataFacetSession) Uri(id *big.Int) (string, error) {
	return _MetadataFacet.Contract.Uri(&_MetadataFacet.CallOpts, id)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 id) view returns(string)
func (_MetadataFacet *MetadataFacetCallerSession) Uri(id *big.Int) (string, error) {
	return _MetadataFacet.Contract.Uri(&_MetadataFacet.CallOpts, id)
}

// MetadataFacetURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the MetadataFacet contract.
type MetadataFacetURIIterator struct {
	Event *MetadataFacetURI // Event containing the contract specifics and raw log

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
func (it *MetadataFacetURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MetadataFacetURI)
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
		it.Event = new(MetadataFacetURI)
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
func (it *MetadataFacetURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MetadataFacetURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MetadataFacetURI represents a URI event raised by the MetadataFacet contract.
type MetadataFacetURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_MetadataFacet *MetadataFacetFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*MetadataFacetURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MetadataFacet.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &MetadataFacetURIIterator{contract: _MetadataFacet.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_MetadataFacet *MetadataFacetFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *MetadataFacetURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _MetadataFacet.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MetadataFacetURI)
				if err := _MetadataFacet.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_MetadataFacet *MetadataFacetFilterer) ParseURI(log types.Log) (*MetadataFacetURI, error) {
	event := new(MetadataFacetURI)
	if err := _MetadataFacet.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
