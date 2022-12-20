// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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

// AbiMetaData contains all meta data concerning the Abi contract.
var AbiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"ClaimBox\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"OpenBox\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"SetBoxURI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"SetSneakerURI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"UpdateTokenId\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_SUPPLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_leaderList\",\"type\":\"address[]\"}],\"name\":\"addLeader\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIdList\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"_toList\",\"type\":\"address[]\"}],\"name\":\"airdrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIdList\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"batchApprove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_leader\",\"type\":\"address\"}],\"name\":\"deleteLeader\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_leader\",\"type\":\"address\"}],\"name\":\"isLeader\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"isOpen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIdList\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"_tokenUriList\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"_boxUriList\",\"type\":\"string[]\"}],\"name\":\"mintBatch200\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"openBox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_defaultUri\",\"type\":\"string\"}],\"name\":\"setDefaultUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_footUri\",\"type\":\"string\"}],\"name\":\"setFootUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_headerUri\",\"type\":\"string\"}],\"name\":\"setHeaderUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_tokenIdUri\",\"type\":\"string\"}],\"name\":\"updateTokenUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenOwner\",\"type\":\"address\"}],\"name\":\"userNftList\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AbiABI is the input ABI used to generate the binding from.
// Deprecated: Use AbiMetaData.ABI instead.
var AbiABI = AbiMetaData.ABI

// Abi is an auto generated Go binding around an Ethereum contract.
type Abi struct {
	AbiCaller     // Read-only binding to the contract
	AbiTransactor // Write-only binding to the contract
	AbiFilterer   // Log filterer for contract events
}

// AbiCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbiSession struct {
	Contract     *Abi              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbiCallerSession struct {
	Contract *AbiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AbiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbiTransactorSession struct {
	Contract     *AbiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbiRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbiRaw struct {
	Contract *Abi // Generic contract binding to access the raw methods on
}

// AbiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbiCallerRaw struct {
	Contract *AbiCaller // Generic read-only contract binding to access the raw methods on
}

// AbiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbiTransactorRaw struct {
	Contract *AbiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbi creates a new instance of Abi, bound to a specific deployed contract.
func NewAbi(address common.Address, backend bind.ContractBackend) (*Abi, error) {
	contract, err := bindAbi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Abi{AbiCaller: AbiCaller{contract: contract}, AbiTransactor: AbiTransactor{contract: contract}, AbiFilterer: AbiFilterer{contract: contract}}, nil
}

// NewAbiCaller creates a new read-only instance of Abi, bound to a specific deployed contract.
func NewAbiCaller(address common.Address, caller bind.ContractCaller) (*AbiCaller, error) {
	contract, err := bindAbi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbiCaller{contract: contract}, nil
}

// NewAbiTransactor creates a new write-only instance of Abi, bound to a specific deployed contract.
func NewAbiTransactor(address common.Address, transactor bind.ContractTransactor) (*AbiTransactor, error) {
	contract, err := bindAbi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbiTransactor{contract: contract}, nil
}

// NewAbiFilterer creates a new log filterer instance of Abi, bound to a specific deployed contract.
func NewAbiFilterer(address common.Address, filterer bind.ContractFilterer) (*AbiFilterer, error) {
	contract, err := bindAbi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbiFilterer{contract: contract}, nil
}

// bindAbi binds a generic wrapper to an already deployed contract.
func bindAbi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AbiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abi *AbiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abi.Contract.AbiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abi *AbiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.Contract.AbiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abi *AbiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abi.Contract.AbiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abi *AbiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abi *AbiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abi *AbiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abi.Contract.contract.Transact(opts, method, params...)
}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint256)
func (_Abi *AbiCaller) MAXSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "MAX_SUPPLY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint256)
func (_Abi *AbiSession) MAXSUPPLY() (*big.Int, error) {
	return _Abi.Contract.MAXSUPPLY(&_Abi.CallOpts)
}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint256)
func (_Abi *AbiCallerSession) MAXSUPPLY() (*big.Int, error) {
	return _Abi.Contract.MAXSUPPLY(&_Abi.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Abi *AbiCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Abi *AbiSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Abi.Contract.BalanceOf(&_Abi.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Abi *AbiCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Abi.Contract.BalanceOf(&_Abi.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Abi *AbiCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Abi *AbiSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Abi.Contract.GetApproved(&_Abi.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Abi *AbiCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Abi.Contract.GetApproved(&_Abi.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Abi *AbiCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Abi *AbiSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Abi.Contract.IsApprovedForAll(&_Abi.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Abi *AbiCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Abi.Contract.IsApprovedForAll(&_Abi.CallOpts, owner, operator)
}

// IsLeader is a free data retrieval call binding the contract method 0xdb512e85.
//
// Solidity: function isLeader(address _leader) view returns(bool)
func (_Abi *AbiCaller) IsLeader(opts *bind.CallOpts, _leader common.Address) (bool, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "isLeader", _leader)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLeader is a free data retrieval call binding the contract method 0xdb512e85.
//
// Solidity: function isLeader(address _leader) view returns(bool)
func (_Abi *AbiSession) IsLeader(_leader common.Address) (bool, error) {
	return _Abi.Contract.IsLeader(&_Abi.CallOpts, _leader)
}

// IsLeader is a free data retrieval call binding the contract method 0xdb512e85.
//
// Solidity: function isLeader(address _leader) view returns(bool)
func (_Abi *AbiCallerSession) IsLeader(_leader common.Address) (bool, error) {
	return _Abi.Contract.IsLeader(&_Abi.CallOpts, _leader)
}

// IsOpen is a free data retrieval call binding the contract method 0x4d6861a6.
//
// Solidity: function isOpen(uint256 _tokenId) view returns(bool)
func (_Abi *AbiCaller) IsOpen(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "isOpen", _tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOpen is a free data retrieval call binding the contract method 0x4d6861a6.
//
// Solidity: function isOpen(uint256 _tokenId) view returns(bool)
func (_Abi *AbiSession) IsOpen(_tokenId *big.Int) (bool, error) {
	return _Abi.Contract.IsOpen(&_Abi.CallOpts, _tokenId)
}

// IsOpen is a free data retrieval call binding the contract method 0x4d6861a6.
//
// Solidity: function isOpen(uint256 _tokenId) view returns(bool)
func (_Abi *AbiCallerSession) IsOpen(_tokenId *big.Int) (bool, error) {
	return _Abi.Contract.IsOpen(&_Abi.CallOpts, _tokenId)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Abi *AbiCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Abi *AbiSession) Name() (string, error) {
	return _Abi.Contract.Name(&_Abi.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Abi *AbiCallerSession) Name() (string, error) {
	return _Abi.Contract.Name(&_Abi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Abi *AbiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Abi *AbiSession) Owner() (common.Address, error) {
	return _Abi.Contract.Owner(&_Abi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Abi *AbiCallerSession) Owner() (common.Address, error) {
	return _Abi.Contract.Owner(&_Abi.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Abi *AbiCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Abi *AbiSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Abi.Contract.OwnerOf(&_Abi.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Abi *AbiCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Abi.Contract.OwnerOf(&_Abi.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Abi *AbiCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Abi *AbiSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Abi.Contract.SupportsInterface(&_Abi.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Abi *AbiCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Abi.Contract.SupportsInterface(&_Abi.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Abi *AbiCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Abi *AbiSession) Symbol() (string, error) {
	return _Abi.Contract.Symbol(&_Abi.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Abi *AbiCallerSession) Symbol() (string, error) {
	return _Abi.Contract.Symbol(&_Abi.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Abi *AbiCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Abi *AbiSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Abi.Contract.TokenByIndex(&_Abi.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Abi *AbiCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Abi.Contract.TokenByIndex(&_Abi.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Abi *AbiCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Abi *AbiSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Abi.Contract.TokenOfOwnerByIndex(&_Abi.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Abi *AbiCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Abi.Contract.TokenOfOwnerByIndex(&_Abi.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Abi *AbiCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "tokenURI", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Abi *AbiSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Abi.Contract.TokenURI(&_Abi.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Abi *AbiCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Abi.Contract.TokenURI(&_Abi.CallOpts, _tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Abi *AbiCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Abi *AbiSession) TotalSupply() (*big.Int, error) {
	return _Abi.Contract.TotalSupply(&_Abi.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Abi *AbiCallerSession) TotalSupply() (*big.Int, error) {
	return _Abi.Contract.TotalSupply(&_Abi.CallOpts)
}

// UserNftList is a free data retrieval call binding the contract method 0x436eb490.
//
// Solidity: function userNftList(address _contractAddress, address _tokenOwner) view returns(uint256[])
func (_Abi *AbiCaller) UserNftList(opts *bind.CallOpts, _contractAddress common.Address, _tokenOwner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Abi.contract.Call(opts, &out, "userNftList", _contractAddress, _tokenOwner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// UserNftList is a free data retrieval call binding the contract method 0x436eb490.
//
// Solidity: function userNftList(address _contractAddress, address _tokenOwner) view returns(uint256[])
func (_Abi *AbiSession) UserNftList(_contractAddress common.Address, _tokenOwner common.Address) ([]*big.Int, error) {
	return _Abi.Contract.UserNftList(&_Abi.CallOpts, _contractAddress, _tokenOwner)
}

// UserNftList is a free data retrieval call binding the contract method 0x436eb490.
//
// Solidity: function userNftList(address _contractAddress, address _tokenOwner) view returns(uint256[])
func (_Abi *AbiCallerSession) UserNftList(_contractAddress common.Address, _tokenOwner common.Address) ([]*big.Int, error) {
	return _Abi.Contract.UserNftList(&_Abi.CallOpts, _contractAddress, _tokenOwner)
}

// AddLeader is a paid mutator transaction binding the contract method 0x8ff6f76d.
//
// Solidity: function addLeader(address[] _leaderList) returns()
func (_Abi *AbiTransactor) AddLeader(opts *bind.TransactOpts, _leaderList []common.Address) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "addLeader", _leaderList)
}

// AddLeader is a paid mutator transaction binding the contract method 0x8ff6f76d.
//
// Solidity: function addLeader(address[] _leaderList) returns()
func (_Abi *AbiSession) AddLeader(_leaderList []common.Address) (*types.Transaction, error) {
	return _Abi.Contract.AddLeader(&_Abi.TransactOpts, _leaderList)
}

// AddLeader is a paid mutator transaction binding the contract method 0x8ff6f76d.
//
// Solidity: function addLeader(address[] _leaderList) returns()
func (_Abi *AbiTransactorSession) AddLeader(_leaderList []common.Address) (*types.Transaction, error) {
	return _Abi.Contract.AddLeader(&_Abi.TransactOpts, _leaderList)
}

// Airdrop is a paid mutator transaction binding the contract method 0x6673c4c2.
//
// Solidity: function airdrop(uint256[] _tokenIdList, address[] _toList) returns()
func (_Abi *AbiTransactor) Airdrop(opts *bind.TransactOpts, _tokenIdList []*big.Int, _toList []common.Address) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "airdrop", _tokenIdList, _toList)
}

// Airdrop is a paid mutator transaction binding the contract method 0x6673c4c2.
//
// Solidity: function airdrop(uint256[] _tokenIdList, address[] _toList) returns()
func (_Abi *AbiSession) Airdrop(_tokenIdList []*big.Int, _toList []common.Address) (*types.Transaction, error) {
	return _Abi.Contract.Airdrop(&_Abi.TransactOpts, _tokenIdList, _toList)
}

// Airdrop is a paid mutator transaction binding the contract method 0x6673c4c2.
//
// Solidity: function airdrop(uint256[] _tokenIdList, address[] _toList) returns()
func (_Abi *AbiTransactorSession) Airdrop(_tokenIdList []*big.Int, _toList []common.Address) (*types.Transaction, error) {
	return _Abi.Contract.Airdrop(&_Abi.TransactOpts, _tokenIdList, _toList)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Abi *AbiTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Abi *AbiSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.Approve(&_Abi.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Abi *AbiTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.Approve(&_Abi.TransactOpts, to, tokenId)
}

// BatchApprove is a paid mutator transaction binding the contract method 0x38752e58.
//
// Solidity: function batchApprove(uint256[] _tokenIdList, address _to) returns()
func (_Abi *AbiTransactor) BatchApprove(opts *bind.TransactOpts, _tokenIdList []*big.Int, _to common.Address) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "batchApprove", _tokenIdList, _to)
}

// BatchApprove is a paid mutator transaction binding the contract method 0x38752e58.
//
// Solidity: function batchApprove(uint256[] _tokenIdList, address _to) returns()
func (_Abi *AbiSession) BatchApprove(_tokenIdList []*big.Int, _to common.Address) (*types.Transaction, error) {
	return _Abi.Contract.BatchApprove(&_Abi.TransactOpts, _tokenIdList, _to)
}

// BatchApprove is a paid mutator transaction binding the contract method 0x38752e58.
//
// Solidity: function batchApprove(uint256[] _tokenIdList, address _to) returns()
func (_Abi *AbiTransactorSession) BatchApprove(_tokenIdList []*big.Int, _to common.Address) (*types.Transaction, error) {
	return _Abi.Contract.BatchApprove(&_Abi.TransactOpts, _tokenIdList, _to)
}

// DeleteLeader is a paid mutator transaction binding the contract method 0x27efcc7c.
//
// Solidity: function deleteLeader(address _leader) returns()
func (_Abi *AbiTransactor) DeleteLeader(opts *bind.TransactOpts, _leader common.Address) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "deleteLeader", _leader)
}

// DeleteLeader is a paid mutator transaction binding the contract method 0x27efcc7c.
//
// Solidity: function deleteLeader(address _leader) returns()
func (_Abi *AbiSession) DeleteLeader(_leader common.Address) (*types.Transaction, error) {
	return _Abi.Contract.DeleteLeader(&_Abi.TransactOpts, _leader)
}

// DeleteLeader is a paid mutator transaction binding the contract method 0x27efcc7c.
//
// Solidity: function deleteLeader(address _leader) returns()
func (_Abi *AbiTransactorSession) DeleteLeader(_leader common.Address) (*types.Transaction, error) {
	return _Abi.Contract.DeleteLeader(&_Abi.TransactOpts, _leader)
}

// MintBatch200 is a paid mutator transaction binding the contract method 0x2de19971.
//
// Solidity: function mintBatch200(address _to, uint256[] _tokenIdList, string[] _tokenUriList, string[] _boxUriList) returns()
func (_Abi *AbiTransactor) MintBatch200(opts *bind.TransactOpts, _to common.Address, _tokenIdList []*big.Int, _tokenUriList []string, _boxUriList []string) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "mintBatch200", _to, _tokenIdList, _tokenUriList, _boxUriList)
}

// MintBatch200 is a paid mutator transaction binding the contract method 0x2de19971.
//
// Solidity: function mintBatch200(address _to, uint256[] _tokenIdList, string[] _tokenUriList, string[] _boxUriList) returns()
func (_Abi *AbiSession) MintBatch200(_to common.Address, _tokenIdList []*big.Int, _tokenUriList []string, _boxUriList []string) (*types.Transaction, error) {
	return _Abi.Contract.MintBatch200(&_Abi.TransactOpts, _to, _tokenIdList, _tokenUriList, _boxUriList)
}

// MintBatch200 is a paid mutator transaction binding the contract method 0x2de19971.
//
// Solidity: function mintBatch200(address _to, uint256[] _tokenIdList, string[] _tokenUriList, string[] _boxUriList) returns()
func (_Abi *AbiTransactorSession) MintBatch200(_to common.Address, _tokenIdList []*big.Int, _tokenUriList []string, _boxUriList []string) (*types.Transaction, error) {
	return _Abi.Contract.MintBatch200(&_Abi.TransactOpts, _to, _tokenIdList, _tokenUriList, _boxUriList)
}

// OpenBox is a paid mutator transaction binding the contract method 0xb1e5e2b7.
//
// Solidity: function openBox(uint256 _tokenId) returns()
func (_Abi *AbiTransactor) OpenBox(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "openBox", _tokenId)
}

// OpenBox is a paid mutator transaction binding the contract method 0xb1e5e2b7.
//
// Solidity: function openBox(uint256 _tokenId) returns()
func (_Abi *AbiSession) OpenBox(_tokenId *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.OpenBox(&_Abi.TransactOpts, _tokenId)
}

// OpenBox is a paid mutator transaction binding the contract method 0xb1e5e2b7.
//
// Solidity: function openBox(uint256 _tokenId) returns()
func (_Abi *AbiTransactorSession) OpenBox(_tokenId *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.OpenBox(&_Abi.TransactOpts, _tokenId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Abi *AbiTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Abi *AbiSession) RenounceOwnership() (*types.Transaction, error) {
	return _Abi.Contract.RenounceOwnership(&_Abi.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Abi *AbiTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Abi.Contract.RenounceOwnership(&_Abi.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Abi *AbiTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Abi *AbiSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.SafeTransferFrom(&_Abi.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Abi *AbiTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.SafeTransferFrom(&_Abi.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Abi *AbiTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Abi *AbiSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Abi.Contract.SafeTransferFrom0(&_Abi.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Abi *AbiTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Abi.Contract.SafeTransferFrom0(&_Abi.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Abi *AbiTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Abi *AbiSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Abi.Contract.SetApprovalForAll(&_Abi.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Abi *AbiTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Abi.Contract.SetApprovalForAll(&_Abi.TransactOpts, operator, approved)
}

// SetDefaultUri is a paid mutator transaction binding the contract method 0x466a18de.
//
// Solidity: function setDefaultUri(string _defaultUri) returns()
func (_Abi *AbiTransactor) SetDefaultUri(opts *bind.TransactOpts, _defaultUri string) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "setDefaultUri", _defaultUri)
}

// SetDefaultUri is a paid mutator transaction binding the contract method 0x466a18de.
//
// Solidity: function setDefaultUri(string _defaultUri) returns()
func (_Abi *AbiSession) SetDefaultUri(_defaultUri string) (*types.Transaction, error) {
	return _Abi.Contract.SetDefaultUri(&_Abi.TransactOpts, _defaultUri)
}

// SetDefaultUri is a paid mutator transaction binding the contract method 0x466a18de.
//
// Solidity: function setDefaultUri(string _defaultUri) returns()
func (_Abi *AbiTransactorSession) SetDefaultUri(_defaultUri string) (*types.Transaction, error) {
	return _Abi.Contract.SetDefaultUri(&_Abi.TransactOpts, _defaultUri)
}

// SetFootUri is a paid mutator transaction binding the contract method 0x663fdd23.
//
// Solidity: function setFootUri(string _footUri) returns()
func (_Abi *AbiTransactor) SetFootUri(opts *bind.TransactOpts, _footUri string) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "setFootUri", _footUri)
}

// SetFootUri is a paid mutator transaction binding the contract method 0x663fdd23.
//
// Solidity: function setFootUri(string _footUri) returns()
func (_Abi *AbiSession) SetFootUri(_footUri string) (*types.Transaction, error) {
	return _Abi.Contract.SetFootUri(&_Abi.TransactOpts, _footUri)
}

// SetFootUri is a paid mutator transaction binding the contract method 0x663fdd23.
//
// Solidity: function setFootUri(string _footUri) returns()
func (_Abi *AbiTransactorSession) SetFootUri(_footUri string) (*types.Transaction, error) {
	return _Abi.Contract.SetFootUri(&_Abi.TransactOpts, _footUri)
}

// SetHeaderUri is a paid mutator transaction binding the contract method 0xa4c8d209.
//
// Solidity: function setHeaderUri(string _headerUri) returns()
func (_Abi *AbiTransactor) SetHeaderUri(opts *bind.TransactOpts, _headerUri string) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "setHeaderUri", _headerUri)
}

// SetHeaderUri is a paid mutator transaction binding the contract method 0xa4c8d209.
//
// Solidity: function setHeaderUri(string _headerUri) returns()
func (_Abi *AbiSession) SetHeaderUri(_headerUri string) (*types.Transaction, error) {
	return _Abi.Contract.SetHeaderUri(&_Abi.TransactOpts, _headerUri)
}

// SetHeaderUri is a paid mutator transaction binding the contract method 0xa4c8d209.
//
// Solidity: function setHeaderUri(string _headerUri) returns()
func (_Abi *AbiTransactorSession) SetHeaderUri(_headerUri string) (*types.Transaction, error) {
	return _Abi.Contract.SetHeaderUri(&_Abi.TransactOpts, _headerUri)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Abi *AbiTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Abi *AbiSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.TransferFrom(&_Abi.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Abi *AbiTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Abi.Contract.TransferFrom(&_Abi.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Abi *AbiTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Abi *AbiSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Abi.Contract.TransferOwnership(&_Abi.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Abi *AbiTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Abi.Contract.TransferOwnership(&_Abi.TransactOpts, newOwner)
}

// UpdateTokenUri is a paid mutator transaction binding the contract method 0xd31af484.
//
// Solidity: function updateTokenUri(uint256 _tokenId, string _tokenIdUri) returns()
func (_Abi *AbiTransactor) UpdateTokenUri(opts *bind.TransactOpts, _tokenId *big.Int, _tokenIdUri string) (*types.Transaction, error) {
	return _Abi.contract.Transact(opts, "updateTokenUri", _tokenId, _tokenIdUri)
}

// UpdateTokenUri is a paid mutator transaction binding the contract method 0xd31af484.
//
// Solidity: function updateTokenUri(uint256 _tokenId, string _tokenIdUri) returns()
func (_Abi *AbiSession) UpdateTokenUri(_tokenId *big.Int, _tokenIdUri string) (*types.Transaction, error) {
	return _Abi.Contract.UpdateTokenUri(&_Abi.TransactOpts, _tokenId, _tokenIdUri)
}

// UpdateTokenUri is a paid mutator transaction binding the contract method 0xd31af484.
//
// Solidity: function updateTokenUri(uint256 _tokenId, string _tokenIdUri) returns()
func (_Abi *AbiTransactorSession) UpdateTokenUri(_tokenId *big.Int, _tokenIdUri string) (*types.Transaction, error) {
	return _Abi.Contract.UpdateTokenUri(&_Abi.TransactOpts, _tokenId, _tokenIdUri)
}

// AbiApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Abi contract.
type AbiApprovalIterator struct {
	Event *AbiApproval // Event containing the contract specifics and raw log

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
func (it *AbiApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiApproval)
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
		it.Event = new(AbiApproval)
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
func (it *AbiApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiApproval represents a Approval event raised by the Abi contract.
type AbiApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Abi *AbiFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*AbiApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &AbiApprovalIterator{contract: _Abi.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Abi *AbiFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AbiApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiApproval)
				if err := _Abi.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Abi *AbiFilterer) ParseApproval(log types.Log) (*AbiApproval, error) {
	event := new(AbiApproval)
	if err := _Abi.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Abi contract.
type AbiApprovalForAllIterator struct {
	Event *AbiApprovalForAll // Event containing the contract specifics and raw log

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
func (it *AbiApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiApprovalForAll)
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
		it.Event = new(AbiApprovalForAll)
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
func (it *AbiApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiApprovalForAll represents a ApprovalForAll event raised by the Abi contract.
type AbiApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Abi *AbiFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*AbiApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &AbiApprovalForAllIterator{contract: _Abi.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Abi *AbiFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *AbiApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiApprovalForAll)
				if err := _Abi.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Abi *AbiFilterer) ParseApprovalForAll(log types.Log) (*AbiApprovalForAll, error) {
	event := new(AbiApprovalForAll)
	if err := _Abi.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiClaimBoxIterator is returned from FilterClaimBox and is used to iterate over the raw logs and unpacked data for ClaimBox events raised by the Abi contract.
type AbiClaimBoxIterator struct {
	Event *AbiClaimBox // Event containing the contract specifics and raw log

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
func (it *AbiClaimBoxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiClaimBox)
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
		it.Event = new(AbiClaimBox)
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
func (it *AbiClaimBoxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiClaimBoxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiClaimBox represents a ClaimBox event raised by the Abi contract.
type AbiClaimBox struct {
	TokenId  *big.Int
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterClaimBox is a free log retrieval operation binding the contract event 0xef3fe58350bbd6cecb2eee87220741c24603a7a5f0ed19286e3e9c1cb7a75df9.
//
// Solidity: event ClaimBox(uint256 _tokenId, address _operator)
func (_Abi *AbiFilterer) FilterClaimBox(opts *bind.FilterOpts) (*AbiClaimBoxIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "ClaimBox")
	if err != nil {
		return nil, err
	}
	return &AbiClaimBoxIterator{contract: _Abi.contract, event: "ClaimBox", logs: logs, sub: sub}, nil
}

// WatchClaimBox is a free log subscription operation binding the contract event 0xef3fe58350bbd6cecb2eee87220741c24603a7a5f0ed19286e3e9c1cb7a75df9.
//
// Solidity: event ClaimBox(uint256 _tokenId, address _operator)
func (_Abi *AbiFilterer) WatchClaimBox(opts *bind.WatchOpts, sink chan<- *AbiClaimBox) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "ClaimBox")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiClaimBox)
				if err := _Abi.contract.UnpackLog(event, "ClaimBox", log); err != nil {
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

// ParseClaimBox is a log parse operation binding the contract event 0xef3fe58350bbd6cecb2eee87220741c24603a7a5f0ed19286e3e9c1cb7a75df9.
//
// Solidity: event ClaimBox(uint256 _tokenId, address _operator)
func (_Abi *AbiFilterer) ParseClaimBox(log types.Log) (*AbiClaimBox, error) {
	event := new(AbiClaimBox)
	if err := _Abi.contract.UnpackLog(event, "ClaimBox", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiOpenBoxIterator is returned from FilterOpenBox and is used to iterate over the raw logs and unpacked data for OpenBox events raised by the Abi contract.
type AbiOpenBoxIterator struct {
	Event *AbiOpenBox // Event containing the contract specifics and raw log

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
func (it *AbiOpenBoxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiOpenBox)
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
		it.Event = new(AbiOpenBox)
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
func (it *AbiOpenBoxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiOpenBoxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiOpenBox represents a OpenBox event raised by the Abi contract.
type AbiOpenBox struct {
	TokenId  *big.Int
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOpenBox is a free log retrieval operation binding the contract event 0x1363abdc7b7584f1c5f0c1032e066b036050a3f41ea052a527297a92a34a452c.
//
// Solidity: event OpenBox(uint256 _tokenId, address _operator)
func (_Abi *AbiFilterer) FilterOpenBox(opts *bind.FilterOpts) (*AbiOpenBoxIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "OpenBox")
	if err != nil {
		return nil, err
	}
	return &AbiOpenBoxIterator{contract: _Abi.contract, event: "OpenBox", logs: logs, sub: sub}, nil
}

// WatchOpenBox is a free log subscription operation binding the contract event 0x1363abdc7b7584f1c5f0c1032e066b036050a3f41ea052a527297a92a34a452c.
//
// Solidity: event OpenBox(uint256 _tokenId, address _operator)
func (_Abi *AbiFilterer) WatchOpenBox(opts *bind.WatchOpts, sink chan<- *AbiOpenBox) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "OpenBox")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiOpenBox)
				if err := _Abi.contract.UnpackLog(event, "OpenBox", log); err != nil {
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

// ParseOpenBox is a log parse operation binding the contract event 0x1363abdc7b7584f1c5f0c1032e066b036050a3f41ea052a527297a92a34a452c.
//
// Solidity: event OpenBox(uint256 _tokenId, address _operator)
func (_Abi *AbiFilterer) ParseOpenBox(log types.Log) (*AbiOpenBox, error) {
	event := new(AbiOpenBox)
	if err := _Abi.contract.UnpackLog(event, "OpenBox", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Abi contract.
type AbiOwnershipTransferredIterator struct {
	Event *AbiOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AbiOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiOwnershipTransferred)
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
		it.Event = new(AbiOwnershipTransferred)
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
func (it *AbiOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiOwnershipTransferred represents a OwnershipTransferred event raised by the Abi contract.
type AbiOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Abi *AbiFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AbiOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AbiOwnershipTransferredIterator{contract: _Abi.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Abi *AbiFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AbiOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiOwnershipTransferred)
				if err := _Abi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Abi *AbiFilterer) ParseOwnershipTransferred(log types.Log) (*AbiOwnershipTransferred, error) {
	event := new(AbiOwnershipTransferred)
	if err := _Abi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiSetBoxURIIterator is returned from FilterSetBoxURI and is used to iterate over the raw logs and unpacked data for SetBoxURI events raised by the Abi contract.
type AbiSetBoxURIIterator struct {
	Event *AbiSetBoxURI // Event containing the contract specifics and raw log

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
func (it *AbiSetBoxURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiSetBoxURI)
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
		it.Event = new(AbiSetBoxURI)
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
func (it *AbiSetBoxURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiSetBoxURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiSetBoxURI represents a SetBoxURI event raised by the Abi contract.
type AbiSetBoxURI struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetBoxURI is a free log retrieval operation binding the contract event 0xb30eea395822d1634bee3d4815af251d02b453d82d259c3ccf42b4a1e65199a7.
//
// Solidity: event SetBoxURI(uint256 _tokenId)
func (_Abi *AbiFilterer) FilterSetBoxURI(opts *bind.FilterOpts) (*AbiSetBoxURIIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "SetBoxURI")
	if err != nil {
		return nil, err
	}
	return &AbiSetBoxURIIterator{contract: _Abi.contract, event: "SetBoxURI", logs: logs, sub: sub}, nil
}

// WatchSetBoxURI is a free log subscription operation binding the contract event 0xb30eea395822d1634bee3d4815af251d02b453d82d259c3ccf42b4a1e65199a7.
//
// Solidity: event SetBoxURI(uint256 _tokenId)
func (_Abi *AbiFilterer) WatchSetBoxURI(opts *bind.WatchOpts, sink chan<- *AbiSetBoxURI) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "SetBoxURI")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiSetBoxURI)
				if err := _Abi.contract.UnpackLog(event, "SetBoxURI", log); err != nil {
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

// ParseSetBoxURI is a log parse operation binding the contract event 0xb30eea395822d1634bee3d4815af251d02b453d82d259c3ccf42b4a1e65199a7.
//
// Solidity: event SetBoxURI(uint256 _tokenId)
func (_Abi *AbiFilterer) ParseSetBoxURI(log types.Log) (*AbiSetBoxURI, error) {
	event := new(AbiSetBoxURI)
	if err := _Abi.contract.UnpackLog(event, "SetBoxURI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiSetSneakerURIIterator is returned from FilterSetSneakerURI and is used to iterate over the raw logs and unpacked data for SetSneakerURI events raised by the Abi contract.
type AbiSetSneakerURIIterator struct {
	Event *AbiSetSneakerURI // Event containing the contract specifics and raw log

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
func (it *AbiSetSneakerURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiSetSneakerURI)
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
		it.Event = new(AbiSetSneakerURI)
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
func (it *AbiSetSneakerURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiSetSneakerURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiSetSneakerURI represents a SetSneakerURI event raised by the Abi contract.
type AbiSetSneakerURI struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetSneakerURI is a free log retrieval operation binding the contract event 0x73c44c51395bcfbc6a4aac8ffb72347e66ca24b2d3f3b57c0c69690a6ecf45fa.
//
// Solidity: event SetSneakerURI(uint256 _tokenId)
func (_Abi *AbiFilterer) FilterSetSneakerURI(opts *bind.FilterOpts) (*AbiSetSneakerURIIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "SetSneakerURI")
	if err != nil {
		return nil, err
	}
	return &AbiSetSneakerURIIterator{contract: _Abi.contract, event: "SetSneakerURI", logs: logs, sub: sub}, nil
}

// WatchSetSneakerURI is a free log subscription operation binding the contract event 0x73c44c51395bcfbc6a4aac8ffb72347e66ca24b2d3f3b57c0c69690a6ecf45fa.
//
// Solidity: event SetSneakerURI(uint256 _tokenId)
func (_Abi *AbiFilterer) WatchSetSneakerURI(opts *bind.WatchOpts, sink chan<- *AbiSetSneakerURI) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "SetSneakerURI")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiSetSneakerURI)
				if err := _Abi.contract.UnpackLog(event, "SetSneakerURI", log); err != nil {
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

// ParseSetSneakerURI is a log parse operation binding the contract event 0x73c44c51395bcfbc6a4aac8ffb72347e66ca24b2d3f3b57c0c69690a6ecf45fa.
//
// Solidity: event SetSneakerURI(uint256 _tokenId)
func (_Abi *AbiFilterer) ParseSetSneakerURI(log types.Log) (*AbiSetSneakerURI, error) {
	event := new(AbiSetSneakerURI)
	if err := _Abi.contract.UnpackLog(event, "SetSneakerURI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Abi contract.
type AbiTransferIterator struct {
	Event *AbiTransfer // Event containing the contract specifics and raw log

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
func (it *AbiTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiTransfer)
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
		it.Event = new(AbiTransfer)
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
func (it *AbiTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiTransfer represents a Transfer event raised by the Abi contract.
type AbiTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Abi *AbiFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*AbiTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Abi.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &AbiTransferIterator{contract: _Abi.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Abi *AbiFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AbiTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Abi.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiTransfer)
				if err := _Abi.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Abi *AbiFilterer) ParseTransfer(log types.Log) (*AbiTransfer, error) {
	event := new(AbiTransfer)
	if err := _Abi.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiUpdateTokenIdIterator is returned from FilterUpdateTokenId and is used to iterate over the raw logs and unpacked data for UpdateTokenId events raised by the Abi contract.
type AbiUpdateTokenIdIterator struct {
	Event *AbiUpdateTokenId // Event containing the contract specifics and raw log

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
func (it *AbiUpdateTokenIdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiUpdateTokenId)
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
		it.Event = new(AbiUpdateTokenId)
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
func (it *AbiUpdateTokenIdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiUpdateTokenIdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiUpdateTokenId represents a UpdateTokenId event raised by the Abi contract.
type AbiUpdateTokenId struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateTokenId is a free log retrieval operation binding the contract event 0x4515a5ddb22c0a35420d783b66f71e6c8b3111ae5d5891f269eb07a2a9619a68.
//
// Solidity: event UpdateTokenId(uint256 _tokenId)
func (_Abi *AbiFilterer) FilterUpdateTokenId(opts *bind.FilterOpts) (*AbiUpdateTokenIdIterator, error) {

	logs, sub, err := _Abi.contract.FilterLogs(opts, "UpdateTokenId")
	if err != nil {
		return nil, err
	}
	return &AbiUpdateTokenIdIterator{contract: _Abi.contract, event: "UpdateTokenId", logs: logs, sub: sub}, nil
}

// WatchUpdateTokenId is a free log subscription operation binding the contract event 0x4515a5ddb22c0a35420d783b66f71e6c8b3111ae5d5891f269eb07a2a9619a68.
//
// Solidity: event UpdateTokenId(uint256 _tokenId)
func (_Abi *AbiFilterer) WatchUpdateTokenId(opts *bind.WatchOpts, sink chan<- *AbiUpdateTokenId) (event.Subscription, error) {

	logs, sub, err := _Abi.contract.WatchLogs(opts, "UpdateTokenId")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiUpdateTokenId)
				if err := _Abi.contract.UnpackLog(event, "UpdateTokenId", log); err != nil {
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

// ParseUpdateTokenId is a log parse operation binding the contract event 0x4515a5ddb22c0a35420d783b66f71e6c8b3111ae5d5891f269eb07a2a9619a68.
//
// Solidity: event UpdateTokenId(uint256 _tokenId)
func (_Abi *AbiFilterer) ParseUpdateTokenId(log types.Log) (*AbiUpdateTokenId, error) {
	event := new(AbiUpdateTokenId)
	if err := _Abi.contract.UnpackLog(event, "UpdateTokenId", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
