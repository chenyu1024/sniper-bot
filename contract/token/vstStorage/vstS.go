// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vstStorage

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

// VstSMetaData contains all meta data concerning the VstS contract.
var VstSMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"sendTORunBa\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VstSABI is the input ABI used to generate the binding from.
// Deprecated: Use VstSMetaData.ABI instead.
var VstSABI = VstSMetaData.ABI

// VstS is an auto generated Go binding around an Ethereum contract.
type VstS struct {
	VstSCaller     // Read-only binding to the contract
	VstSTransactor // Write-only binding to the contract
	VstSFilterer   // Log filterer for contract events
}

// VstSCaller is an auto generated read-only Go binding around an Ethereum contract.
type VstSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VstSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VstSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VstSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VstSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VstSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VstSSession struct {
	Contract     *VstS             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VstSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VstSCallerSession struct {
	Contract *VstSCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VstSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VstSTransactorSession struct {
	Contract     *VstSTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VstSRaw is an auto generated low-level Go binding around an Ethereum contract.
type VstSRaw struct {
	Contract *VstS // Generic contract binding to access the raw methods on
}

// VstSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VstSCallerRaw struct {
	Contract *VstSCaller // Generic read-only contract binding to access the raw methods on
}

// VstSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VstSTransactorRaw struct {
	Contract *VstSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVstS creates a new instance of VstS, bound to a specific deployed contract.
func NewVstS(address common.Address, backend bind.ContractBackend) (*VstS, error) {
	contract, err := bindVstS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VstS{VstSCaller: VstSCaller{contract: contract}, VstSTransactor: VstSTransactor{contract: contract}, VstSFilterer: VstSFilterer{contract: contract}}, nil
}

// NewVstSCaller creates a new read-only instance of VstS, bound to a specific deployed contract.
func NewVstSCaller(address common.Address, caller bind.ContractCaller) (*VstSCaller, error) {
	contract, err := bindVstS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VstSCaller{contract: contract}, nil
}

// NewVstSTransactor creates a new write-only instance of VstS, bound to a specific deployed contract.
func NewVstSTransactor(address common.Address, transactor bind.ContractTransactor) (*VstSTransactor, error) {
	contract, err := bindVstS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VstSTransactor{contract: contract}, nil
}

// NewVstSFilterer creates a new log filterer instance of VstS, bound to a specific deployed contract.
func NewVstSFilterer(address common.Address, filterer bind.ContractFilterer) (*VstSFilterer, error) {
	contract, err := bindVstS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VstSFilterer{contract: contract}, nil
}

// bindVstS binds a generic wrapper to an already deployed contract.
func bindVstS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VstSABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VstS *VstSRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VstS.Contract.VstSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VstS *VstSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VstS.Contract.VstSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VstS *VstSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VstS.Contract.VstSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VstS *VstSCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VstS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VstS *VstSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VstS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VstS *VstSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VstS.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_VstS *VstSCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VstS.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_VstS *VstSSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _VstS.Contract.Allowance(&_VstS.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_VstS *VstSCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _VstS.Contract.Allowance(&_VstS.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_VstS *VstSCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VstS.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_VstS *VstSSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _VstS.Contract.BalanceOf(&_VstS.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_VstS *VstSCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _VstS.Contract.BalanceOf(&_VstS.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_VstS *VstSCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _VstS.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_VstS *VstSSession) Decimals() (uint8, error) {
	return _VstS.Contract.Decimals(&_VstS.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_VstS *VstSCallerSession) Decimals() (uint8, error) {
	return _VstS.Contract.Decimals(&_VstS.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VstS *VstSCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VstS.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VstS *VstSSession) Name() (string, error) {
	return _VstS.Contract.Name(&_VstS.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VstS *VstSCallerSession) Name() (string, error) {
	return _VstS.Contract.Name(&_VstS.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VstS *VstSCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VstS.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VstS *VstSSession) Owner() (common.Address, error) {
	return _VstS.Contract.Owner(&_VstS.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VstS *VstSCallerSession) Owner() (common.Address, error) {
	return _VstS.Contract.Owner(&_VstS.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_VstS *VstSCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VstS.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_VstS *VstSSession) Symbol() (string, error) {
	return _VstS.Contract.Symbol(&_VstS.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_VstS *VstSCallerSession) Symbol() (string, error) {
	return _VstS.Contract.Symbol(&_VstS.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VstS *VstSCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VstS.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VstS *VstSSession) TotalSupply() (*big.Int, error) {
	return _VstS.Contract.TotalSupply(&_VstS.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VstS *VstSCallerSession) TotalSupply() (*big.Int, error) {
	return _VstS.Contract.TotalSupply(&_VstS.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_VstS *VstSTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VstS.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_VstS *VstSSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VstS.Contract.Approve(&_VstS.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_VstS *VstSTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VstS.Contract.Approve(&_VstS.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_VstS *VstSTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _VstS.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_VstS *VstSSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _VstS.Contract.DecreaseAllowance(&_VstS.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_VstS *VstSTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _VstS.Contract.DecreaseAllowance(&_VstS.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_VstS *VstSTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _VstS.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_VstS *VstSSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _VstS.Contract.IncreaseAllowance(&_VstS.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_VstS *VstSTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _VstS.Contract.IncreaseAllowance(&_VstS.TransactOpts, spender, addedValue)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VstS *VstSTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VstS.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VstS *VstSSession) RenounceOwnership() (*types.Transaction, error) {
	return _VstS.Contract.RenounceOwnership(&_VstS.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VstS *VstSTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VstS.Contract.RenounceOwnership(&_VstS.TransactOpts)
}

// SendTORunBa is a paid mutator transaction binding the contract method 0xc9f3935e.
//
// Solidity: function sendTORunBa(address token, address _to, uint256 _amount) returns()
func (_VstS *VstSTransactor) SendTORunBa(opts *bind.TransactOpts, token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VstS.contract.Transact(opts, "sendTORunBa", token, _to, _amount)
}

// SendTORunBa is a paid mutator transaction binding the contract method 0xc9f3935e.
//
// Solidity: function sendTORunBa(address token, address _to, uint256 _amount) returns()
func (_VstS *VstSSession) SendTORunBa(token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VstS.Contract.SendTORunBa(&_VstS.TransactOpts, token, _to, _amount)
}

// SendTORunBa is a paid mutator transaction binding the contract method 0xc9f3935e.
//
// Solidity: function sendTORunBa(address token, address _to, uint256 _amount) returns()
func (_VstS *VstSTransactorSession) SendTORunBa(token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VstS.Contract.SendTORunBa(&_VstS.TransactOpts, token, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_VstS *VstSTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VstS.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_VstS *VstSSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VstS.Contract.Transfer(&_VstS.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_VstS *VstSTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VstS.Contract.Transfer(&_VstS.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_VstS *VstSTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VstS.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_VstS *VstSSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VstS.Contract.TransferFrom(&_VstS.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_VstS *VstSTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VstS.Contract.TransferFrom(&_VstS.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VstS *VstSTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VstS.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VstS *VstSSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VstS.Contract.TransferOwnership(&_VstS.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VstS *VstSTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VstS.Contract.TransferOwnership(&_VstS.TransactOpts, newOwner)
}

// VstSApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the VstS contract.
type VstSApprovalIterator struct {
	Event *VstSApproval // Event containing the contract specifics and raw log

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
func (it *VstSApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VstSApproval)
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
		it.Event = new(VstSApproval)
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
func (it *VstSApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VstSApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VstSApproval represents a Approval event raised by the VstS contract.
type VstSApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_VstS *VstSFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*VstSApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _VstS.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &VstSApprovalIterator{contract: _VstS.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_VstS *VstSFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *VstSApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _VstS.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VstSApproval)
				if err := _VstS.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_VstS *VstSFilterer) ParseApproval(log types.Log) (*VstSApproval, error) {
	event := new(VstSApproval)
	if err := _VstS.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VstSOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VstS contract.
type VstSOwnershipTransferredIterator struct {
	Event *VstSOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VstSOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VstSOwnershipTransferred)
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
		it.Event = new(VstSOwnershipTransferred)
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
func (it *VstSOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VstSOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VstSOwnershipTransferred represents a OwnershipTransferred event raised by the VstS contract.
type VstSOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VstS *VstSFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VstSOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VstS.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VstSOwnershipTransferredIterator{contract: _VstS.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VstS *VstSFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VstSOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VstS.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VstSOwnershipTransferred)
				if err := _VstS.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VstS *VstSFilterer) ParseOwnershipTransferred(log types.Log) (*VstSOwnershipTransferred, error) {
	event := new(VstSOwnershipTransferred)
	if err := _VstS.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VstSTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the VstS contract.
type VstSTransferIterator struct {
	Event *VstSTransfer // Event containing the contract specifics and raw log

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
func (it *VstSTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VstSTransfer)
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
		it.Event = new(VstSTransfer)
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
func (it *VstSTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VstSTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VstSTransfer represents a Transfer event raised by the VstS contract.
type VstSTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_VstS *VstSFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VstSTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VstS.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VstSTransferIterator{contract: _VstS.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_VstS *VstSFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *VstSTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VstS.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VstSTransfer)
				if err := _VstS.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_VstS *VstSFilterer) ParseTransfer(log types.Log) (*VstSTransfer, error) {
	event := new(VstSTransfer)
	if err := _VstS.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
