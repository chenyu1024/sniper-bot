// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vmtStorage

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

// VmtSMetaData contains all meta data concerning the VmtS contract.
var VmtSMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"sendVMTToOther\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VmtSABI is the input ABI used to generate the binding from.
// Deprecated: Use VmtSMetaData.ABI instead.
var VmtSABI = VmtSMetaData.ABI

// VmtS is an auto generated Go binding around an Ethereum contract.
type VmtS struct {
	VmtSCaller     // Read-only binding to the contract
	VmtSTransactor // Write-only binding to the contract
	VmtSFilterer   // Log filterer for contract events
}

// VmtSCaller is an auto generated read-only Go binding around an Ethereum contract.
type VmtSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VmtSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VmtSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VmtSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VmtSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VmtSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VmtSSession struct {
	Contract     *VmtS             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VmtSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VmtSCallerSession struct {
	Contract *VmtSCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VmtSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VmtSTransactorSession struct {
	Contract     *VmtSTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VmtSRaw is an auto generated low-level Go binding around an Ethereum contract.
type VmtSRaw struct {
	Contract *VmtS // Generic contract binding to access the raw methods on
}

// VmtSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VmtSCallerRaw struct {
	Contract *VmtSCaller // Generic read-only contract binding to access the raw methods on
}

// VmtSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VmtSTransactorRaw struct {
	Contract *VmtSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVmtS creates a new instance of VmtS, bound to a specific deployed contract.
func NewVmtS(address common.Address, backend bind.ContractBackend) (*VmtS, error) {
	contract, err := bindVmtS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VmtS{VmtSCaller: VmtSCaller{contract: contract}, VmtSTransactor: VmtSTransactor{contract: contract}, VmtSFilterer: VmtSFilterer{contract: contract}}, nil
}

// NewVmtSCaller creates a new read-only instance of VmtS, bound to a specific deployed contract.
func NewVmtSCaller(address common.Address, caller bind.ContractCaller) (*VmtSCaller, error) {
	contract, err := bindVmtS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VmtSCaller{contract: contract}, nil
}

// NewVmtSTransactor creates a new write-only instance of VmtS, bound to a specific deployed contract.
func NewVmtSTransactor(address common.Address, transactor bind.ContractTransactor) (*VmtSTransactor, error) {
	contract, err := bindVmtS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VmtSTransactor{contract: contract}, nil
}

// NewVmtSFilterer creates a new log filterer instance of VmtS, bound to a specific deployed contract.
func NewVmtSFilterer(address common.Address, filterer bind.ContractFilterer) (*VmtSFilterer, error) {
	contract, err := bindVmtS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VmtSFilterer{contract: contract}, nil
}

// bindVmtS binds a generic wrapper to an already deployed contract.
func bindVmtS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VmtSABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VmtS *VmtSRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VmtS.Contract.VmtSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VmtS *VmtSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VmtS.Contract.VmtSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VmtS *VmtSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VmtS.Contract.VmtSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VmtS *VmtSCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VmtS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VmtS *VmtSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VmtS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VmtS *VmtSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VmtS.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_VmtS *VmtSCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VmtS.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_VmtS *VmtSSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _VmtS.Contract.Allowance(&_VmtS.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_VmtS *VmtSCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _VmtS.Contract.Allowance(&_VmtS.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_VmtS *VmtSCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VmtS.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_VmtS *VmtSSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _VmtS.Contract.BalanceOf(&_VmtS.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_VmtS *VmtSCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _VmtS.Contract.BalanceOf(&_VmtS.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_VmtS *VmtSCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _VmtS.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_VmtS *VmtSSession) Decimals() (uint8, error) {
	return _VmtS.Contract.Decimals(&_VmtS.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_VmtS *VmtSCallerSession) Decimals() (uint8, error) {
	return _VmtS.Contract.Decimals(&_VmtS.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VmtS *VmtSCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VmtS.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VmtS *VmtSSession) Name() (string, error) {
	return _VmtS.Contract.Name(&_VmtS.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_VmtS *VmtSCallerSession) Name() (string, error) {
	return _VmtS.Contract.Name(&_VmtS.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VmtS *VmtSCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VmtS.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VmtS *VmtSSession) Owner() (common.Address, error) {
	return _VmtS.Contract.Owner(&_VmtS.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VmtS *VmtSCallerSession) Owner() (common.Address, error) {
	return _VmtS.Contract.Owner(&_VmtS.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_VmtS *VmtSCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _VmtS.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_VmtS *VmtSSession) Symbol() (string, error) {
	return _VmtS.Contract.Symbol(&_VmtS.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_VmtS *VmtSCallerSession) Symbol() (string, error) {
	return _VmtS.Contract.Symbol(&_VmtS.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VmtS *VmtSCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VmtS.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VmtS *VmtSSession) TotalSupply() (*big.Int, error) {
	return _VmtS.Contract.TotalSupply(&_VmtS.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VmtS *VmtSCallerSession) TotalSupply() (*big.Int, error) {
	return _VmtS.Contract.TotalSupply(&_VmtS.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_VmtS *VmtSTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VmtS.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_VmtS *VmtSSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VmtS.Contract.Approve(&_VmtS.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_VmtS *VmtSTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VmtS.Contract.Approve(&_VmtS.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_VmtS *VmtSTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _VmtS.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_VmtS *VmtSSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _VmtS.Contract.DecreaseAllowance(&_VmtS.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_VmtS *VmtSTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _VmtS.Contract.DecreaseAllowance(&_VmtS.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_VmtS *VmtSTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _VmtS.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_VmtS *VmtSSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _VmtS.Contract.IncreaseAllowance(&_VmtS.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_VmtS *VmtSTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _VmtS.Contract.IncreaseAllowance(&_VmtS.TransactOpts, spender, addedValue)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VmtS *VmtSTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VmtS.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VmtS *VmtSSession) RenounceOwnership() (*types.Transaction, error) {
	return _VmtS.Contract.RenounceOwnership(&_VmtS.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VmtS *VmtSTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VmtS.Contract.RenounceOwnership(&_VmtS.TransactOpts)
}

// SendVMTToOther is a paid mutator transaction binding the contract method 0x068c034a.
//
// Solidity: function sendVMTToOther(address token, address _to, uint256 _amount) returns()
func (_VmtS *VmtSTransactor) SendVMTToOther(opts *bind.TransactOpts, token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VmtS.contract.Transact(opts, "sendVMTToOther", token, _to, _amount)
}

// SendVMTToOther is a paid mutator transaction binding the contract method 0x068c034a.
//
// Solidity: function sendVMTToOther(address token, address _to, uint256 _amount) returns()
func (_VmtS *VmtSSession) SendVMTToOther(token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VmtS.Contract.SendVMTToOther(&_VmtS.TransactOpts, token, _to, _amount)
}

// SendVMTToOther is a paid mutator transaction binding the contract method 0x068c034a.
//
// Solidity: function sendVMTToOther(address token, address _to, uint256 _amount) returns()
func (_VmtS *VmtSTransactorSession) SendVMTToOther(token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VmtS.Contract.SendVMTToOther(&_VmtS.TransactOpts, token, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_VmtS *VmtSTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VmtS.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_VmtS *VmtSSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VmtS.Contract.Transfer(&_VmtS.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_VmtS *VmtSTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VmtS.Contract.Transfer(&_VmtS.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_VmtS *VmtSTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VmtS.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_VmtS *VmtSSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VmtS.Contract.TransferFrom(&_VmtS.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_VmtS *VmtSTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _VmtS.Contract.TransferFrom(&_VmtS.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VmtS *VmtSTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VmtS.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VmtS *VmtSSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VmtS.Contract.TransferOwnership(&_VmtS.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VmtS *VmtSTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VmtS.Contract.TransferOwnership(&_VmtS.TransactOpts, newOwner)
}

// VmtSApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the VmtS contract.
type VmtSApprovalIterator struct {
	Event *VmtSApproval // Event containing the contract specifics and raw log

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
func (it *VmtSApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VmtSApproval)
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
		it.Event = new(VmtSApproval)
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
func (it *VmtSApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VmtSApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VmtSApproval represents a Approval event raised by the VmtS contract.
type VmtSApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_VmtS *VmtSFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*VmtSApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _VmtS.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &VmtSApprovalIterator{contract: _VmtS.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_VmtS *VmtSFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *VmtSApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _VmtS.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VmtSApproval)
				if err := _VmtS.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_VmtS *VmtSFilterer) ParseApproval(log types.Log) (*VmtSApproval, error) {
	event := new(VmtSApproval)
	if err := _VmtS.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VmtSOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VmtS contract.
type VmtSOwnershipTransferredIterator struct {
	Event *VmtSOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VmtSOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VmtSOwnershipTransferred)
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
		it.Event = new(VmtSOwnershipTransferred)
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
func (it *VmtSOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VmtSOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VmtSOwnershipTransferred represents a OwnershipTransferred event raised by the VmtS contract.
type VmtSOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VmtS *VmtSFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VmtSOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VmtS.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VmtSOwnershipTransferredIterator{contract: _VmtS.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VmtS *VmtSFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VmtSOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VmtS.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VmtSOwnershipTransferred)
				if err := _VmtS.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VmtS *VmtSFilterer) ParseOwnershipTransferred(log types.Log) (*VmtSOwnershipTransferred, error) {
	event := new(VmtSOwnershipTransferred)
	if err := _VmtS.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VmtSTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the VmtS contract.
type VmtSTransferIterator struct {
	Event *VmtSTransfer // Event containing the contract specifics and raw log

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
func (it *VmtSTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VmtSTransfer)
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
		it.Event = new(VmtSTransfer)
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
func (it *VmtSTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VmtSTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VmtSTransfer represents a Transfer event raised by the VmtS contract.
type VmtSTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_VmtS *VmtSFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VmtSTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VmtS.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VmtSTransferIterator{contract: _VmtS.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_VmtS *VmtSFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *VmtSTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VmtS.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VmtSTransfer)
				if err := _VmtS.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_VmtS *VmtSFilterer) ParseTransfer(log types.Log) (*VmtSTransfer, error) {
	event := new(VmtSTransfer)
	if err := _VmtS.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
