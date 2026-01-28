// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package excusestore

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

// LoadingstoreMetaData contains all meta data concerning the Loadingstore contract.
var LoadingstoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"ItemSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"items\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"setItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LoadingstoreABI is the input ABI used to generate the binding from.
// Deprecated: Use LoadingstoreMetaData.ABI instead.
var LoadingstoreABI = LoadingstoreMetaData.ABI

// Loadingstore is an auto generated Go binding around an Ethereum contract.
type Loadingstore struct {
	LoadingstoreCaller     // Read-only binding to the contract
	LoadingstoreTransactor // Write-only binding to the contract
	LoadingstoreFilterer   // Log filterer for contract events
}

// LoadingstoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type LoadingstoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoadingstoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LoadingstoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoadingstoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LoadingstoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoadingstoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LoadingstoreSession struct {
	Contract     *Loadingstore     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LoadingstoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LoadingstoreCallerSession struct {
	Contract *LoadingstoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// LoadingstoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LoadingstoreTransactorSession struct {
	Contract     *LoadingstoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// LoadingstoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type LoadingstoreRaw struct {
	Contract *Loadingstore // Generic contract binding to access the raw methods on
}

// LoadingstoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LoadingstoreCallerRaw struct {
	Contract *LoadingstoreCaller // Generic read-only contract binding to access the raw methods on
}

// LoadingstoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LoadingstoreTransactorRaw struct {
	Contract *LoadingstoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLoadingstore creates a new instance of Loadingstore, bound to a specific deployed contract.
func NewLoadingstore(address common.Address, backend bind.ContractBackend) (*Loadingstore, error) {
	contract, err := bindLoadingstore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Loadingstore{LoadingstoreCaller: LoadingstoreCaller{contract: contract}, LoadingstoreTransactor: LoadingstoreTransactor{contract: contract}, LoadingstoreFilterer: LoadingstoreFilterer{contract: contract}}, nil
}

// NewLoadingstoreCaller creates a new read-only instance of Loadingstore, bound to a specific deployed contract.
func NewLoadingstoreCaller(address common.Address, caller bind.ContractCaller) (*LoadingstoreCaller, error) {
	contract, err := bindLoadingstore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LoadingstoreCaller{contract: contract}, nil
}

// NewLoadingstoreTransactor creates a new write-only instance of Loadingstore, bound to a specific deployed contract.
func NewLoadingstoreTransactor(address common.Address, transactor bind.ContractTransactor) (*LoadingstoreTransactor, error) {
	contract, err := bindLoadingstore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LoadingstoreTransactor{contract: contract}, nil
}

// NewLoadingstoreFilterer creates a new log filterer instance of Loadingstore, bound to a specific deployed contract.
func NewLoadingstoreFilterer(address common.Address, filterer bind.ContractFilterer) (*LoadingstoreFilterer, error) {
	contract, err := bindLoadingstore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LoadingstoreFilterer{contract: contract}, nil
}

// bindLoadingstore binds a generic wrapper to an already deployed contract.
func bindLoadingstore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LoadingstoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Loadingstore *LoadingstoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Loadingstore.Contract.LoadingstoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Loadingstore *LoadingstoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Loadingstore.Contract.LoadingstoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Loadingstore *LoadingstoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Loadingstore.Contract.LoadingstoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Loadingstore *LoadingstoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Loadingstore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Loadingstore *LoadingstoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Loadingstore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Loadingstore *LoadingstoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Loadingstore.Contract.contract.Transact(opts, method, params...)
}

// Items is a free data retrieval call binding the contract method 0x48f343f3.
//
// Solidity: function items(bytes32 ) view returns(bytes32)
func (_Loadingstore *LoadingstoreCaller) Items(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Loadingstore.contract.Call(opts, &out, "items", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Items is a free data retrieval call binding the contract method 0x48f343f3.
//
// Solidity: function items(bytes32 ) view returns(bytes32)
func (_Loadingstore *LoadingstoreSession) Items(arg0 [32]byte) ([32]byte, error) {
	return _Loadingstore.Contract.Items(&_Loadingstore.CallOpts, arg0)
}

// Items is a free data retrieval call binding the contract method 0x48f343f3.
//
// Solidity: function items(bytes32 ) view returns(bytes32)
func (_Loadingstore *LoadingstoreCallerSession) Items(arg0 [32]byte) ([32]byte, error) {
	return _Loadingstore.Contract.Items(&_Loadingstore.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Loadingstore *LoadingstoreCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Loadingstore.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Loadingstore *LoadingstoreSession) Version() (string, error) {
	return _Loadingstore.Contract.Version(&_Loadingstore.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Loadingstore *LoadingstoreCallerSession) Version() (string, error) {
	return _Loadingstore.Contract.Version(&_Loadingstore.CallOpts)
}

// SetItem is a paid mutator transaction binding the contract method 0xf56256c7.
//
// Solidity: function setItem(bytes32 key, bytes32 value) returns()
func (_Loadingstore *LoadingstoreTransactor) SetItem(opts *bind.TransactOpts, key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Loadingstore.contract.Transact(opts, "setItem", key, value)
}

// SetItem is a paid mutator transaction binding the contract method 0xf56256c7.
//
// Solidity: function setItem(bytes32 key, bytes32 value) returns()
func (_Loadingstore *LoadingstoreSession) SetItem(key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Loadingstore.Contract.SetItem(&_Loadingstore.TransactOpts, key, value)
}

// SetItem is a paid mutator transaction binding the contract method 0xf56256c7.
//
// Solidity: function setItem(bytes32 key, bytes32 value) returns()
func (_Loadingstore *LoadingstoreTransactorSession) SetItem(key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Loadingstore.Contract.SetItem(&_Loadingstore.TransactOpts, key, value)
}

// LoadingstoreItemSetIterator is returned from FilterItemSet and is used to iterate over the raw logs and unpacked data for ItemSet events raised by the Loadingstore contract.
type LoadingstoreItemSetIterator struct {
	Event *LoadingstoreItemSet // Event containing the contract specifics and raw log

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
func (it *LoadingstoreItemSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoadingstoreItemSet)
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
		it.Event = new(LoadingstoreItemSet)
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
func (it *LoadingstoreItemSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoadingstoreItemSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoadingstoreItemSet represents a ItemSet event raised by the Loadingstore contract.
type LoadingstoreItemSet struct {
	Key   [32]byte
	Value [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterItemSet is a free log retrieval operation binding the contract event 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4.
//
// Solidity: event ItemSet(bytes32 key, bytes32 value)
func (_Loadingstore *LoadingstoreFilterer) FilterItemSet(opts *bind.FilterOpts) (*LoadingstoreItemSetIterator, error) {

	logs, sub, err := _Loadingstore.contract.FilterLogs(opts, "ItemSet")
	if err != nil {
		return nil, err
	}
	return &LoadingstoreItemSetIterator{contract: _Loadingstore.contract, event: "ItemSet", logs: logs, sub: sub}, nil
}

// WatchItemSet is a free log subscription operation binding the contract event 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4.
//
// Solidity: event ItemSet(bytes32 key, bytes32 value)
func (_Loadingstore *LoadingstoreFilterer) WatchItemSet(opts *bind.WatchOpts, sink chan<- *LoadingstoreItemSet) (event.Subscription, error) {

	logs, sub, err := _Loadingstore.contract.WatchLogs(opts, "ItemSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoadingstoreItemSet)
				if err := _Loadingstore.contract.UnpackLog(event, "ItemSet", log); err != nil {
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

// ParseItemSet is a log parse operation binding the contract event 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4.
//
// Solidity: event ItemSet(bytes32 key, bytes32 value)
func (_Loadingstore *LoadingstoreFilterer) ParseItemSet(log types.Log) (*LoadingstoreItemSet, error) {
	event := new(LoadingstoreItemSet)
	if err := _Loadingstore.contract.UnpackLog(event, "ItemSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
