// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package buy

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
)

// BuyMetaData contains all meta data concerning the Buy contract.
var BuyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_path0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_path1\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_slippage\",\"type\":\"uint256\"}],\"name\":\"buyNewToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"__router\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_path0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_path1\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_slippage\",\"type\":\"uint256\"}],\"name\":\"simpleBuy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_path0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_path1\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"}],\"name\":\"simpleBuyv1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"_router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BuyABI is the input ABI used to generate the binding from.
// Deprecated: Use BuyMetaData.ABI instead.
var BuyABI = BuyMetaData.ABI

// Buy is an auto generated Go binding around an Ethereum contract.
type Buy struct {
	BuyCaller     // Read-only binding to the contract
	BuyTransactor // Write-only binding to the contract
	BuyFilterer   // Log filterer for contract events
}

// BuyCaller is an auto generated read-only Go binding around an Ethereum contract.
type BuyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BuyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BuyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BuySession struct {
	Contract     *Buy              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BuyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BuyCallerSession struct {
	Contract *BuyCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BuyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BuyTransactorSession struct {
	Contract     *BuyTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BuyRaw is an auto generated low-level Go binding around an Ethereum contract.
type BuyRaw struct {
	Contract *Buy // Generic contract binding to access the raw methods on
}

// BuyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BuyCallerRaw struct {
	Contract *BuyCaller // Generic read-only contract binding to access the raw methods on
}

// BuyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BuyTransactorRaw struct {
	Contract *BuyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBuy creates a new instance of Buy, bound to a specific deployed contract.
func NewBuy(address common.Address, backend bind.ContractBackend) (*Buy, error) {
	contract, err := bindBuy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Buy{BuyCaller: BuyCaller{contract: contract}, BuyTransactor: BuyTransactor{contract: contract}, BuyFilterer: BuyFilterer{contract: contract}}, nil
}

// NewBuyCaller creates a new read-only instance of Buy, bound to a specific deployed contract.
func NewBuyCaller(address common.Address, caller bind.ContractCaller) (*BuyCaller, error) {
	contract, err := bindBuy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BuyCaller{contract: contract}, nil
}

// NewBuyTransactor creates a new write-only instance of Buy, bound to a specific deployed contract.
func NewBuyTransactor(address common.Address, transactor bind.ContractTransactor) (*BuyTransactor, error) {
	contract, err := bindBuy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BuyTransactor{contract: contract}, nil
}

// NewBuyFilterer creates a new log filterer instance of Buy, bound to a specific deployed contract.
func NewBuyFilterer(address common.Address, filterer bind.ContractFilterer) (*BuyFilterer, error) {
	contract, err := bindBuy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BuyFilterer{contract: contract}, nil
}

// bindBuy binds a generic wrapper to an already deployed contract.
func bindBuy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BuyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Buy *BuyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Buy.Contract.BuyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Buy *BuyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Buy.Contract.BuyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Buy *BuyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Buy.Contract.BuyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Buy *BuyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Buy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Buy *BuyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Buy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Buy *BuyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Buy.Contract.contract.Transact(opts, method, params...)
}

// Router is a free data retrieval call binding the contract method 0xedae876f.
//
// Solidity: function _router() view returns(address)
func (_Buy *BuyCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "_router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xedae876f.
//
// Solidity: function _router() view returns(address)
func (_Buy *BuySession) Router() (common.Address, error) {
	return _Buy.Contract.Router(&_Buy.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xedae876f.
//
// Solidity: function _router() view returns(address)
func (_Buy *BuyCallerSession) Router() (common.Address, error) {
	return _Buy.Contract.Router(&_Buy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Buy *BuyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Buy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Buy *BuySession) Owner() (common.Address, error) {
	return _Buy.Contract.Owner(&_Buy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Buy *BuyCallerSession) Owner() (common.Address, error) {
	return _Buy.Contract.Owner(&_Buy.CallOpts)
}

// BuyNewToken is a paid mutator transaction binding the contract method 0x55c6fc8d.
//
// Solidity: function buyNewToken(address _path0, address _path1, uint256 _amountIn, uint256 _slippage) returns()
func (_Buy *BuyTransactor) BuyNewToken(opts *bind.TransactOpts, _path0 common.Address, _path1 common.Address, _amountIn *big.Int, _slippage *big.Int) (*types.Transaction, error) {
	return _Buy.contract.Transact(opts, "buyNewToken", _path0, _path1, _amountIn, _slippage)
}

// BuyNewToken is a paid mutator transaction binding the contract method 0x55c6fc8d.
//
// Solidity: function buyNewToken(address _path0, address _path1, uint256 _amountIn, uint256 _slippage) returns()
func (_Buy *BuySession) BuyNewToken(_path0 common.Address, _path1 common.Address, _amountIn *big.Int, _slippage *big.Int) (*types.Transaction, error) {
	return _Buy.Contract.BuyNewToken(&_Buy.TransactOpts, _path0, _path1, _amountIn, _slippage)
}

// BuyNewToken is a paid mutator transaction binding the contract method 0x55c6fc8d.
//
// Solidity: function buyNewToken(address _path0, address _path1, uint256 _amountIn, uint256 _slippage) returns()
func (_Buy *BuyTransactorSession) BuyNewToken(_path0 common.Address, _path1 common.Address, _amountIn *big.Int, _slippage *big.Int) (*types.Transaction, error) {
	return _Buy.Contract.BuyNewToken(&_Buy.TransactOpts, _path0, _path1, _amountIn, _slippage)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Buy *BuyTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Buy.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Buy *BuySession) RenounceOwnership() (*types.Transaction, error) {
	return _Buy.Contract.RenounceOwnership(&_Buy.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Buy *BuyTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Buy.Contract.RenounceOwnership(&_Buy.TransactOpts)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address __router) returns()
func (_Buy *BuyTransactor) SetRouter(opts *bind.TransactOpts, __router common.Address) (*types.Transaction, error) {
	return _Buy.contract.Transact(opts, "setRouter", __router)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address __router) returns()
func (_Buy *BuySession) SetRouter(__router common.Address) (*types.Transaction, error) {
	return _Buy.Contract.SetRouter(&_Buy.TransactOpts, __router)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address __router) returns()
func (_Buy *BuyTransactorSession) SetRouter(__router common.Address) (*types.Transaction, error) {
	return _Buy.Contract.SetRouter(&_Buy.TransactOpts, __router)
}

// SimpleBuy is a paid mutator transaction binding the contract method 0x08543e40.
//
// Solidity: function simpleBuy(address _path0, address _path1, uint256 _amountIn, uint256 _slippage) returns()
func (_Buy *BuyTransactor) SimpleBuy(opts *bind.TransactOpts, _path0 common.Address, _path1 common.Address, _amountIn *big.Int, _slippage *big.Int) (*types.Transaction, error) {
	return _Buy.contract.Transact(opts, "simpleBuy", _path0, _path1, _amountIn, _slippage)
}

// SimpleBuy is a paid mutator transaction binding the contract method 0x08543e40.
//
// Solidity: function simpleBuy(address _path0, address _path1, uint256 _amountIn, uint256 _slippage) returns()
func (_Buy *BuySession) SimpleBuy(_path0 common.Address, _path1 common.Address, _amountIn *big.Int, _slippage *big.Int) (*types.Transaction, error) {
	return _Buy.Contract.SimpleBuy(&_Buy.TransactOpts, _path0, _path1, _amountIn, _slippage)
}

// SimpleBuy is a paid mutator transaction binding the contract method 0x08543e40.
//
// Solidity: function simpleBuy(address _path0, address _path1, uint256 _amountIn, uint256 _slippage) returns()
func (_Buy *BuyTransactorSession) SimpleBuy(_path0 common.Address, _path1 common.Address, _amountIn *big.Int, _slippage *big.Int) (*types.Transaction, error) {
	return _Buy.Contract.SimpleBuy(&_Buy.TransactOpts, _path0, _path1, _amountIn, _slippage)
}

// SimpleBuyv1 is a paid mutator transaction binding the contract method 0x0648f126.
//
// Solidity: function simpleBuyv1(address _path0, address _path1, uint256 _amountIn) returns()
func (_Buy *BuyTransactor) SimpleBuyv1(opts *bind.TransactOpts, _path0 common.Address, _path1 common.Address, _amountIn *big.Int) (*types.Transaction, error) {
	return _Buy.contract.Transact(opts, "simpleBuyv1", _path0, _path1, _amountIn)
}

// SimpleBuyv1 is a paid mutator transaction binding the contract method 0x0648f126.
//
// Solidity: function simpleBuyv1(address _path0, address _path1, uint256 _amountIn) returns()
func (_Buy *BuySession) SimpleBuyv1(_path0 common.Address, _path1 common.Address, _amountIn *big.Int) (*types.Transaction, error) {
	return _Buy.Contract.SimpleBuyv1(&_Buy.TransactOpts, _path0, _path1, _amountIn)
}

// SimpleBuyv1 is a paid mutator transaction binding the contract method 0x0648f126.
//
// Solidity: function simpleBuyv1(address _path0, address _path1, uint256 _amountIn) returns()
func (_Buy *BuyTransactorSession) SimpleBuyv1(_path0 common.Address, _path1 common.Address, _amountIn *big.Int) (*types.Transaction, error) {
	return _Buy.Contract.SimpleBuyv1(&_Buy.TransactOpts, _path0, _path1, _amountIn)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Buy *BuyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Buy.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Buy *BuySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Buy.Contract.TransferOwnership(&_Buy.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Buy *BuyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Buy.Contract.TransferOwnership(&_Buy.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address token) returns()
func (_Buy *BuyTransactor) Withdraw(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Buy.contract.Transact(opts, "withdraw", token)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address token) returns()
func (_Buy *BuySession) Withdraw(token common.Address) (*types.Transaction, error) {
	return _Buy.Contract.Withdraw(&_Buy.TransactOpts, token)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address token) returns()
func (_Buy *BuyTransactorSession) Withdraw(token common.Address) (*types.Transaction, error) {
	return _Buy.Contract.Withdraw(&_Buy.TransactOpts, token)
}

// WithdrawAmount is a paid mutator transaction binding the contract method 0x736fe565.
//
// Solidity: function withdrawAmount(address token, uint256 amount) returns()
func (_Buy *BuyTransactor) WithdrawAmount(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Buy.contract.Transact(opts, "withdrawAmount", token, amount)
}

// WithdrawAmount is a paid mutator transaction binding the contract method 0x736fe565.
//
// Solidity: function withdrawAmount(address token, uint256 amount) returns()
func (_Buy *BuySession) WithdrawAmount(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Buy.Contract.WithdrawAmount(&_Buy.TransactOpts, token, amount)
}

// WithdrawAmount is a paid mutator transaction binding the contract method 0x736fe565.
//
// Solidity: function withdrawAmount(address token, uint256 amount) returns()
func (_Buy *BuyTransactorSession) WithdrawAmount(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Buy.Contract.WithdrawAmount(&_Buy.TransactOpts, token, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xe086e5ec.
//
// Solidity: function withdrawETH() payable returns()
func (_Buy *BuyTransactor) WithdrawETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Buy.contract.Transact(opts, "withdrawETH")
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xe086e5ec.
//
// Solidity: function withdrawETH() payable returns()
func (_Buy *BuySession) WithdrawETH() (*types.Transaction, error) {
	return _Buy.Contract.WithdrawETH(&_Buy.TransactOpts)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xe086e5ec.
//
// Solidity: function withdrawETH() payable returns()
func (_Buy *BuyTransactorSession) WithdrawETH() (*types.Transaction, error) {
	return _Buy.Contract.WithdrawETH(&_Buy.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Buy *BuyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Buy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Buy *BuySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Buy.Contract.Fallback(&_Buy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Buy *BuyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Buy.Contract.Fallback(&_Buy.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Buy *BuyTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Buy.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Buy *BuySession) Receive() (*types.Transaction, error) {
	return _Buy.Contract.Receive(&_Buy.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Buy *BuyTransactorSession) Receive() (*types.Transaction, error) {
	return _Buy.Contract.Receive(&_Buy.TransactOpts)
}

// BuyOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Buy contract.
type BuyOwnershipTransferredIterator struct {
	Event *BuyOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BuyOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuyOwnershipTransferred)
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
		it.Event = new(BuyOwnershipTransferred)
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
func (it *BuyOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuyOwnershipTransferred represents a OwnershipTransferred event raised by the Buy contract.
type BuyOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Buy *BuyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BuyOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Buy.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BuyOwnershipTransferredIterator{contract: _Buy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Buy *BuyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BuyOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Buy.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuyOwnershipTransferred)
				if err := _Buy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Buy *BuyFilterer) ParseOwnershipTransferred(log types.Log) (*BuyOwnershipTransferred, error) {
	event := new(BuyOwnershipTransferred)
	if err := _Buy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
