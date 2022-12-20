// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package runbaTest

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

// RunbaTestMetaData contains all meta data concerning the RunbaTest contract.
var RunbaTestMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_leaderList\",\"type\":\"address[]\"}],\"name\":\"addLeader\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIdList\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"_toList\",\"type\":\"address[]\"}],\"name\":\"airdrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIdList\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"batchApprove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"ClaimBox\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_leader\",\"type\":\"address\"}],\"name\":\"deleteLeader\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIdList\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"_tokenUriList\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"_boxUriList\",\"type\":\"string[]\"}],\"name\":\"mintBatch200\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"openBox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"OpenBox\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"SetBoxURI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_defaultUri\",\"type\":\"string\"}],\"name\":\"setDefaultUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_footUri\",\"type\":\"string\"}],\"name\":\"setFootUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_headerUri\",\"type\":\"string\"}],\"name\":\"setHeaderUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"SetSneakerURI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"UpdateTokenId\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_tokenIdUri\",\"type\":\"string\"}],\"name\":\"updateTokenUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_leader\",\"type\":\"address\"}],\"name\":\"isLeader\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"isOpen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_SUPPLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenOwner\",\"type\":\"address\"}],\"name\":\"userNftList\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// RunbaTestABI is the input ABI used to generate the binding from.
// Deprecated: Use RunbaTestMetaData.ABI instead.
var RunbaTestABI = RunbaTestMetaData.ABI

// RunbaTest is an auto generated Go binding around an Ethereum contract.
type RunbaTest struct {
	RunbaTestCaller     // Read-only binding to the contract
	RunbaTestTransactor // Write-only binding to the contract
	RunbaTestFilterer   // Log filterer for contract events
}

// RunbaTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type RunbaTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RunbaTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RunbaTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RunbaTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RunbaTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RunbaTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RunbaTestSession struct {
	Contract     *RunbaTest        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RunbaTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RunbaTestCallerSession struct {
	Contract *RunbaTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RunbaTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RunbaTestTransactorSession struct {
	Contract     *RunbaTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RunbaTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type RunbaTestRaw struct {
	Contract *RunbaTest // Generic contract binding to access the raw methods on
}

// RunbaTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RunbaTestCallerRaw struct {
	Contract *RunbaTestCaller // Generic read-only contract binding to access the raw methods on
}

// RunbaTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RunbaTestTransactorRaw struct {
	Contract *RunbaTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRunbaTest creates a new instance of RunbaTest, bound to a specific deployed contract.
func NewRunbaTest(address common.Address, backend bind.ContractBackend) (*RunbaTest, error) {
	contract, err := bindRunbaTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RunbaTest{RunbaTestCaller: RunbaTestCaller{contract: contract}, RunbaTestTransactor: RunbaTestTransactor{contract: contract}, RunbaTestFilterer: RunbaTestFilterer{contract: contract}}, nil
}

// NewRunbaTestCaller creates a new read-only instance of RunbaTest, bound to a specific deployed contract.
func NewRunbaTestCaller(address common.Address, caller bind.ContractCaller) (*RunbaTestCaller, error) {
	contract, err := bindRunbaTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RunbaTestCaller{contract: contract}, nil
}

// NewRunbaTestTransactor creates a new write-only instance of RunbaTest, bound to a specific deployed contract.
func NewRunbaTestTransactor(address common.Address, transactor bind.ContractTransactor) (*RunbaTestTransactor, error) {
	contract, err := bindRunbaTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RunbaTestTransactor{contract: contract}, nil
}

// NewRunbaTestFilterer creates a new log filterer instance of RunbaTest, bound to a specific deployed contract.
func NewRunbaTestFilterer(address common.Address, filterer bind.ContractFilterer) (*RunbaTestFilterer, error) {
	contract, err := bindRunbaTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RunbaTestFilterer{contract: contract}, nil
}

// bindRunbaTest binds a generic wrapper to an already deployed contract.
func bindRunbaTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RunbaTestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RunbaTest *RunbaTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RunbaTest.Contract.RunbaTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RunbaTest *RunbaTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RunbaTest.Contract.RunbaTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RunbaTest *RunbaTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RunbaTest.Contract.RunbaTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RunbaTest *RunbaTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RunbaTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RunbaTest *RunbaTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RunbaTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RunbaTest *RunbaTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RunbaTest.Contract.contract.Transact(opts, method, params...)
}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint256)
func (_RunbaTest *RunbaTestCaller) MAXSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RunbaTest.contract.Call(opts, &out, "MAX_SUPPLY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint256)
func (_RunbaTest *RunbaTestSession) MAXSUPPLY() (*big.Int, error) {
	return _RunbaTest.Contract.MAXSUPPLY(&_RunbaTest.CallOpts)
}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint256)
func (_RunbaTest *RunbaTestCallerSession) MAXSUPPLY() (*big.Int, error) {
	return _RunbaTest.Contract.MAXSUPPLY(&_RunbaTest.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_RunbaTest *RunbaTestCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RunbaTest.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_RunbaTest *RunbaTestSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _RunbaTest.Contract.BalanceOf(&_RunbaTest.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_RunbaTest *RunbaTestCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _RunbaTest.Contract.BalanceOf(&_RunbaTest.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_RunbaTest *RunbaTestCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RunbaTest.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_RunbaTest *RunbaTestSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _RunbaTest.Contract.GetApproved(&_RunbaTest.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_RunbaTest *RunbaTestCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _RunbaTest.Contract.GetApproved(&_RunbaTest.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_RunbaTest *RunbaTestCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _RunbaTest.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_RunbaTest *RunbaTestSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _RunbaTest.Contract.IsApprovedForAll(&_RunbaTest.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_RunbaTest *RunbaTestCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _RunbaTest.Contract.IsApprovedForAll(&_RunbaTest.CallOpts, owner, operator)
}

// IsLeader is a free data retrieval call binding the contract method 0xdb512e85.
//
// Solidity: function isLeader(address _leader) view returns(bool)
func (_RunbaTest *RunbaTestCaller) IsLeader(opts *bind.CallOpts, _leader common.Address) (bool, error) {
	var out []interface{}
	err := _RunbaTest.contract.Call(opts, &out, "isLeader", _leader)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLeader is a free data retrieval call binding the contract method 0xdb512e85.
//
// Solidity: function isLeader(address _leader) view returns(bool)
func (_RunbaTest *RunbaTestSession) IsLeader(_leader common.Address) (bool, error) {
	return _RunbaTest.Contract.IsLeader(&_RunbaTest.CallOpts, _leader)
}

// IsLeader is a free data retrieval call binding the contract method 0xdb512e85.
//
// Solidity: function isLeader(address _leader) view returns(bool)
func (_RunbaTest *RunbaTestCallerSession) IsLeader(_leader common.Address) (bool, error) {
	return _RunbaTest.Contract.IsLeader(&_RunbaTest.CallOpts, _leader)
}

// IsOpen is a free data retrieval call binding the contract method 0x4d6861a6.
//
// Solidity: function isOpen(uint256 _tokenId) view returns(bool)
func (_RunbaTest *RunbaTestCaller) IsOpen(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _RunbaTest.contract.Call(opts, &out, "isOpen", _tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOpen is a free data retrieval call binding the contract method 0x4d6861a6.
//
// Solidity: function isOpen(uint256 _tokenId) view returns(bool)
func (_RunbaTest *RunbaTestSession) IsOpen(_tokenId *big.Int) (bool, error) {
	return _RunbaTest.Contract.IsOpen(&_RunbaTest.CallOpts, _tokenId)
}

// IsOpen is a free data retrieval call binding the contract method 0x4d6861a6.
//
// Solidity: function isOpen(uint256 _tokenId) view returns(bool)
func (_RunbaTest *RunbaTestCallerSession) IsOpen(_tokenId *big.Int) (bool, error) {
	return _RunbaTest.Contract.IsOpen(&_RunbaTest.CallOpts, _tokenId)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RunbaTest *RunbaTestCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RunbaTest.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RunbaTest *RunbaTestSession) Name() (string, error) {
	return _RunbaTest.Contract.Name(&_RunbaTest.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RunbaTest *RunbaTestCallerSession) Name() (string, error) {
	return _RunbaTest.Contract.Name(&_RunbaTest.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RunbaTest *RunbaTestCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RunbaTest.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RunbaTest *RunbaTestSession) Owner() (common.Address, error) {
	return _RunbaTest.Contract.Owner(&_RunbaTest.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RunbaTest *RunbaTestCallerSession) Owner() (common.Address, error) {
	return _RunbaTest.Contract.Owner(&_RunbaTest.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_RunbaTest *RunbaTestCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RunbaTest.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_RunbaTest *RunbaTestSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _RunbaTest.Contract.OwnerOf(&_RunbaTest.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_RunbaTest *RunbaTestCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _RunbaTest.Contract.OwnerOf(&_RunbaTest.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RunbaTest *RunbaTestCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _RunbaTest.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RunbaTest *RunbaTestSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RunbaTest.Contract.SupportsInterface(&_RunbaTest.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RunbaTest *RunbaTestCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RunbaTest.Contract.SupportsInterface(&_RunbaTest.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RunbaTest *RunbaTestCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RunbaTest.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RunbaTest *RunbaTestSession) Symbol() (string, error) {
	return _RunbaTest.Contract.Symbol(&_RunbaTest.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RunbaTest *RunbaTestCallerSession) Symbol() (string, error) {
	return _RunbaTest.Contract.Symbol(&_RunbaTest.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_RunbaTest *RunbaTestCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RunbaTest.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_RunbaTest *RunbaTestSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _RunbaTest.Contract.TokenByIndex(&_RunbaTest.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_RunbaTest *RunbaTestCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _RunbaTest.Contract.TokenByIndex(&_RunbaTest.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_RunbaTest *RunbaTestCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RunbaTest.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_RunbaTest *RunbaTestSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _RunbaTest.Contract.TokenOfOwnerByIndex(&_RunbaTest.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_RunbaTest *RunbaTestCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _RunbaTest.Contract.TokenOfOwnerByIndex(&_RunbaTest.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_RunbaTest *RunbaTestCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _RunbaTest.contract.Call(opts, &out, "tokenURI", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_RunbaTest *RunbaTestSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _RunbaTest.Contract.TokenURI(&_RunbaTest.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_RunbaTest *RunbaTestCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _RunbaTest.Contract.TokenURI(&_RunbaTest.CallOpts, _tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RunbaTest *RunbaTestCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RunbaTest.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RunbaTest *RunbaTestSession) TotalSupply() (*big.Int, error) {
	return _RunbaTest.Contract.TotalSupply(&_RunbaTest.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RunbaTest *RunbaTestCallerSession) TotalSupply() (*big.Int, error) {
	return _RunbaTest.Contract.TotalSupply(&_RunbaTest.CallOpts)
}

// UserNftList is a free data retrieval call binding the contract method 0x436eb490.
//
// Solidity: function userNftList(address _contractAddress, address _tokenOwner) view returns(uint256[])
func (_RunbaTest *RunbaTestCaller) UserNftList(opts *bind.CallOpts, _contractAddress common.Address, _tokenOwner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _RunbaTest.contract.Call(opts, &out, "userNftList", _contractAddress, _tokenOwner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// UserNftList is a free data retrieval call binding the contract method 0x436eb490.
//
// Solidity: function userNftList(address _contractAddress, address _tokenOwner) view returns(uint256[])
func (_RunbaTest *RunbaTestSession) UserNftList(_contractAddress common.Address, _tokenOwner common.Address) ([]*big.Int, error) {
	return _RunbaTest.Contract.UserNftList(&_RunbaTest.CallOpts, _contractAddress, _tokenOwner)
}

// UserNftList is a free data retrieval call binding the contract method 0x436eb490.
//
// Solidity: function userNftList(address _contractAddress, address _tokenOwner) view returns(uint256[])
func (_RunbaTest *RunbaTestCallerSession) UserNftList(_contractAddress common.Address, _tokenOwner common.Address) ([]*big.Int, error) {
	return _RunbaTest.Contract.UserNftList(&_RunbaTest.CallOpts, _contractAddress, _tokenOwner)
}

// AddLeader is a paid mutator transaction binding the contract method 0x8ff6f76d.
//
// Solidity: function addLeader(address[] _leaderList) returns()
func (_RunbaTest *RunbaTestTransactor) AddLeader(opts *bind.TransactOpts, _leaderList []common.Address) (*types.Transaction, error) {
	return _RunbaTest.contract.Transact(opts, "addLeader", _leaderList)
}

// AddLeader is a paid mutator transaction binding the contract method 0x8ff6f76d.
//
// Solidity: function addLeader(address[] _leaderList) returns()
func (_RunbaTest *RunbaTestSession) AddLeader(_leaderList []common.Address) (*types.Transaction, error) {
	return _RunbaTest.Contract.AddLeader(&_RunbaTest.TransactOpts, _leaderList)
}

// AddLeader is a paid mutator transaction binding the contract method 0x8ff6f76d.
//
// Solidity: function addLeader(address[] _leaderList) returns()
func (_RunbaTest *RunbaTestTransactorSession) AddLeader(_leaderList []common.Address) (*types.Transaction, error) {
	return _RunbaTest.Contract.AddLeader(&_RunbaTest.TransactOpts, _leaderList)
}

// Airdrop is a paid mutator transaction binding the contract method 0x6673c4c2.
//
// Solidity: function airdrop(uint256[] _tokenIdList, address[] _toList) returns()
func (_RunbaTest *RunbaTestTransactor) Airdrop(opts *bind.TransactOpts, _tokenIdList []*big.Int, _toList []common.Address) (*types.Transaction, error) {
	return _RunbaTest.contract.Transact(opts, "airdrop", _tokenIdList, _toList)
}

// Airdrop is a paid mutator transaction binding the contract method 0x6673c4c2.
//
// Solidity: function airdrop(uint256[] _tokenIdList, address[] _toList) returns()
func (_RunbaTest *RunbaTestSession) Airdrop(_tokenIdList []*big.Int, _toList []common.Address) (*types.Transaction, error) {
	return _RunbaTest.Contract.Airdrop(&_RunbaTest.TransactOpts, _tokenIdList, _toList)
}

// Airdrop is a paid mutator transaction binding the contract method 0x6673c4c2.
//
// Solidity: function airdrop(uint256[] _tokenIdList, address[] _toList) returns()
func (_RunbaTest *RunbaTestTransactorSession) Airdrop(_tokenIdList []*big.Int, _toList []common.Address) (*types.Transaction, error) {
	return _RunbaTest.Contract.Airdrop(&_RunbaTest.TransactOpts, _tokenIdList, _toList)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_RunbaTest *RunbaTestTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RunbaTest.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_RunbaTest *RunbaTestSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RunbaTest.Contract.Approve(&_RunbaTest.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_RunbaTest *RunbaTestTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RunbaTest.Contract.Approve(&_RunbaTest.TransactOpts, to, tokenId)
}

// BatchApprove is a paid mutator transaction binding the contract method 0x38752e58.
//
// Solidity: function batchApprove(uint256[] _tokenIdList, address _to) returns()
func (_RunbaTest *RunbaTestTransactor) BatchApprove(opts *bind.TransactOpts, _tokenIdList []*big.Int, _to common.Address) (*types.Transaction, error) {
	return _RunbaTest.contract.Transact(opts, "batchApprove", _tokenIdList, _to)
}

// BatchApprove is a paid mutator transaction binding the contract method 0x38752e58.
//
// Solidity: function batchApprove(uint256[] _tokenIdList, address _to) returns()
func (_RunbaTest *RunbaTestSession) BatchApprove(_tokenIdList []*big.Int, _to common.Address) (*types.Transaction, error) {
	return _RunbaTest.Contract.BatchApprove(&_RunbaTest.TransactOpts, _tokenIdList, _to)
}

// BatchApprove is a paid mutator transaction binding the contract method 0x38752e58.
//
// Solidity: function batchApprove(uint256[] _tokenIdList, address _to) returns()
func (_RunbaTest *RunbaTestTransactorSession) BatchApprove(_tokenIdList []*big.Int, _to common.Address) (*types.Transaction, error) {
	return _RunbaTest.Contract.BatchApprove(&_RunbaTest.TransactOpts, _tokenIdList, _to)
}

// DeleteLeader is a paid mutator transaction binding the contract method 0x27efcc7c.
//
// Solidity: function deleteLeader(address _leader) returns()
func (_RunbaTest *RunbaTestTransactor) DeleteLeader(opts *bind.TransactOpts, _leader common.Address) (*types.Transaction, error) {
	return _RunbaTest.contract.Transact(opts, "deleteLeader", _leader)
}

// DeleteLeader is a paid mutator transaction binding the contract method 0x27efcc7c.
//
// Solidity: function deleteLeader(address _leader) returns()
func (_RunbaTest *RunbaTestSession) DeleteLeader(_leader common.Address) (*types.Transaction, error) {
	return _RunbaTest.Contract.DeleteLeader(&_RunbaTest.TransactOpts, _leader)
}

// DeleteLeader is a paid mutator transaction binding the contract method 0x27efcc7c.
//
// Solidity: function deleteLeader(address _leader) returns()
func (_RunbaTest *RunbaTestTransactorSession) DeleteLeader(_leader common.Address) (*types.Transaction, error) {
	return _RunbaTest.Contract.DeleteLeader(&_RunbaTest.TransactOpts, _leader)
}

// MintBatch200 is a paid mutator transaction binding the contract method 0x2de19971.
//
// Solidity: function mintBatch200(address _to, uint256[] _tokenIdList, string[] _tokenUriList, string[] _boxUriList) returns()
func (_RunbaTest *RunbaTestTransactor) MintBatch200(opts *bind.TransactOpts, _to common.Address, _tokenIdList []*big.Int, _tokenUriList []string, _boxUriList []string) (*types.Transaction, error) {
	return _RunbaTest.contract.Transact(opts, "mintBatch200", _to, _tokenIdList, _tokenUriList, _boxUriList)
}

// MintBatch200 is a paid mutator transaction binding the contract method 0x2de19971.
//
// Solidity: function mintBatch200(address _to, uint256[] _tokenIdList, string[] _tokenUriList, string[] _boxUriList) returns()
func (_RunbaTest *RunbaTestSession) MintBatch200(_to common.Address, _tokenIdList []*big.Int, _tokenUriList []string, _boxUriList []string) (*types.Transaction, error) {
	return _RunbaTest.Contract.MintBatch200(&_RunbaTest.TransactOpts, _to, _tokenIdList, _tokenUriList, _boxUriList)
}

// MintBatch200 is a paid mutator transaction binding the contract method 0x2de19971.
//
// Solidity: function mintBatch200(address _to, uint256[] _tokenIdList, string[] _tokenUriList, string[] _boxUriList) returns()
func (_RunbaTest *RunbaTestTransactorSession) MintBatch200(_to common.Address, _tokenIdList []*big.Int, _tokenUriList []string, _boxUriList []string) (*types.Transaction, error) {
	return _RunbaTest.Contract.MintBatch200(&_RunbaTest.TransactOpts, _to, _tokenIdList, _tokenUriList, _boxUriList)
}

// OpenBox is a paid mutator transaction binding the contract method 0xb1e5e2b7.
//
// Solidity: function openBox(uint256 _tokenId) returns()
func (_RunbaTest *RunbaTestTransactor) OpenBox(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _RunbaTest.contract.Transact(opts, "openBox", _tokenId)
}

// OpenBox is a paid mutator transaction binding the contract method 0xb1e5e2b7.
//
// Solidity: function openBox(uint256 _tokenId) returns()
func (_RunbaTest *RunbaTestSession) OpenBox(_tokenId *big.Int) (*types.Transaction, error) {
	return _RunbaTest.Contract.OpenBox(&_RunbaTest.TransactOpts, _tokenId)
}

// OpenBox is a paid mutator transaction binding the contract method 0xb1e5e2b7.
//
// Solidity: function openBox(uint256 _tokenId) returns()
func (_RunbaTest *RunbaTestTransactorSession) OpenBox(_tokenId *big.Int) (*types.Transaction, error) {
	return _RunbaTest.Contract.OpenBox(&_RunbaTest.TransactOpts, _tokenId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RunbaTest *RunbaTestTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RunbaTest.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RunbaTest *RunbaTestSession) RenounceOwnership() (*types.Transaction, error) {
	return _RunbaTest.Contract.RenounceOwnership(&_RunbaTest.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RunbaTest *RunbaTestTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RunbaTest.Contract.RenounceOwnership(&_RunbaTest.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_RunbaTest *RunbaTestTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RunbaTest.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_RunbaTest *RunbaTestSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RunbaTest.Contract.SafeTransferFrom(&_RunbaTest.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_RunbaTest *RunbaTestTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RunbaTest.Contract.SafeTransferFrom(&_RunbaTest.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_RunbaTest *RunbaTestTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _RunbaTest.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_RunbaTest *RunbaTestSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _RunbaTest.Contract.SafeTransferFrom0(&_RunbaTest.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_RunbaTest *RunbaTestTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _RunbaTest.Contract.SafeTransferFrom0(&_RunbaTest.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_RunbaTest *RunbaTestTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _RunbaTest.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_RunbaTest *RunbaTestSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _RunbaTest.Contract.SetApprovalForAll(&_RunbaTest.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_RunbaTest *RunbaTestTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _RunbaTest.Contract.SetApprovalForAll(&_RunbaTest.TransactOpts, operator, approved)
}

// SetDefaultUri is a paid mutator transaction binding the contract method 0x466a18de.
//
// Solidity: function setDefaultUri(string _defaultUri) returns()
func (_RunbaTest *RunbaTestTransactor) SetDefaultUri(opts *bind.TransactOpts, _defaultUri string) (*types.Transaction, error) {
	return _RunbaTest.contract.Transact(opts, "setDefaultUri", _defaultUri)
}

// SetDefaultUri is a paid mutator transaction binding the contract method 0x466a18de.
//
// Solidity: function setDefaultUri(string _defaultUri) returns()
func (_RunbaTest *RunbaTestSession) SetDefaultUri(_defaultUri string) (*types.Transaction, error) {
	return _RunbaTest.Contract.SetDefaultUri(&_RunbaTest.TransactOpts, _defaultUri)
}

// SetDefaultUri is a paid mutator transaction binding the contract method 0x466a18de.
//
// Solidity: function setDefaultUri(string _defaultUri) returns()
func (_RunbaTest *RunbaTestTransactorSession) SetDefaultUri(_defaultUri string) (*types.Transaction, error) {
	return _RunbaTest.Contract.SetDefaultUri(&_RunbaTest.TransactOpts, _defaultUri)
}

// SetFootUri is a paid mutator transaction binding the contract method 0x663fdd23.
//
// Solidity: function setFootUri(string _footUri) returns()
func (_RunbaTest *RunbaTestTransactor) SetFootUri(opts *bind.TransactOpts, _footUri string) (*types.Transaction, error) {
	return _RunbaTest.contract.Transact(opts, "setFootUri", _footUri)
}

// SetFootUri is a paid mutator transaction binding the contract method 0x663fdd23.
//
// Solidity: function setFootUri(string _footUri) returns()
func (_RunbaTest *RunbaTestSession) SetFootUri(_footUri string) (*types.Transaction, error) {
	return _RunbaTest.Contract.SetFootUri(&_RunbaTest.TransactOpts, _footUri)
}

// SetFootUri is a paid mutator transaction binding the contract method 0x663fdd23.
//
// Solidity: function setFootUri(string _footUri) returns()
func (_RunbaTest *RunbaTestTransactorSession) SetFootUri(_footUri string) (*types.Transaction, error) {
	return _RunbaTest.Contract.SetFootUri(&_RunbaTest.TransactOpts, _footUri)
}

// SetHeaderUri is a paid mutator transaction binding the contract method 0xa4c8d209.
//
// Solidity: function setHeaderUri(string _headerUri) returns()
func (_RunbaTest *RunbaTestTransactor) SetHeaderUri(opts *bind.TransactOpts, _headerUri string) (*types.Transaction, error) {
	return _RunbaTest.contract.Transact(opts, "setHeaderUri", _headerUri)
}

// SetHeaderUri is a paid mutator transaction binding the contract method 0xa4c8d209.
//
// Solidity: function setHeaderUri(string _headerUri) returns()
func (_RunbaTest *RunbaTestSession) SetHeaderUri(_headerUri string) (*types.Transaction, error) {
	return _RunbaTest.Contract.SetHeaderUri(&_RunbaTest.TransactOpts, _headerUri)
}

// SetHeaderUri is a paid mutator transaction binding the contract method 0xa4c8d209.
//
// Solidity: function setHeaderUri(string _headerUri) returns()
func (_RunbaTest *RunbaTestTransactorSession) SetHeaderUri(_headerUri string) (*types.Transaction, error) {
	return _RunbaTest.Contract.SetHeaderUri(&_RunbaTest.TransactOpts, _headerUri)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_RunbaTest *RunbaTestTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RunbaTest.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_RunbaTest *RunbaTestSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RunbaTest.Contract.TransferFrom(&_RunbaTest.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_RunbaTest *RunbaTestTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RunbaTest.Contract.TransferFrom(&_RunbaTest.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RunbaTest *RunbaTestTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RunbaTest.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RunbaTest *RunbaTestSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RunbaTest.Contract.TransferOwnership(&_RunbaTest.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RunbaTest *RunbaTestTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RunbaTest.Contract.TransferOwnership(&_RunbaTest.TransactOpts, newOwner)
}

// UpdateTokenUri is a paid mutator transaction binding the contract method 0xd31af484.
//
// Solidity: function updateTokenUri(uint256 _tokenId, string _tokenIdUri) returns()
func (_RunbaTest *RunbaTestTransactor) UpdateTokenUri(opts *bind.TransactOpts, _tokenId *big.Int, _tokenIdUri string) (*types.Transaction, error) {
	return _RunbaTest.contract.Transact(opts, "updateTokenUri", _tokenId, _tokenIdUri)
}

// UpdateTokenUri is a paid mutator transaction binding the contract method 0xd31af484.
//
// Solidity: function updateTokenUri(uint256 _tokenId, string _tokenIdUri) returns()
func (_RunbaTest *RunbaTestSession) UpdateTokenUri(_tokenId *big.Int, _tokenIdUri string) (*types.Transaction, error) {
	return _RunbaTest.Contract.UpdateTokenUri(&_RunbaTest.TransactOpts, _tokenId, _tokenIdUri)
}

// UpdateTokenUri is a paid mutator transaction binding the contract method 0xd31af484.
//
// Solidity: function updateTokenUri(uint256 _tokenId, string _tokenIdUri) returns()
func (_RunbaTest *RunbaTestTransactorSession) UpdateTokenUri(_tokenId *big.Int, _tokenIdUri string) (*types.Transaction, error) {
	return _RunbaTest.Contract.UpdateTokenUri(&_RunbaTest.TransactOpts, _tokenId, _tokenIdUri)
}

// RunbaTestApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the RunbaTest contract.
type RunbaTestApprovalIterator struct {
	Event *RunbaTestApproval // Event containing the contract specifics and raw log

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
func (it *RunbaTestApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RunbaTestApproval)
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
		it.Event = new(RunbaTestApproval)
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
func (it *RunbaTestApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RunbaTestApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RunbaTestApproval represents a Approval event raised by the RunbaTest contract.
type RunbaTestApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_RunbaTest *RunbaTestFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*RunbaTestApprovalIterator, error) {

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

	logs, sub, err := _RunbaTest.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &RunbaTestApprovalIterator{contract: _RunbaTest.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_RunbaTest *RunbaTestFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *RunbaTestApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _RunbaTest.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RunbaTestApproval)
				if err := _RunbaTest.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_RunbaTest *RunbaTestFilterer) ParseApproval(log types.Log) (*RunbaTestApproval, error) {
	event := new(RunbaTestApproval)
	if err := _RunbaTest.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RunbaTestApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the RunbaTest contract.
type RunbaTestApprovalForAllIterator struct {
	Event *RunbaTestApprovalForAll // Event containing the contract specifics and raw log

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
func (it *RunbaTestApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RunbaTestApprovalForAll)
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
		it.Event = new(RunbaTestApprovalForAll)
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
func (it *RunbaTestApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RunbaTestApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RunbaTestApprovalForAll represents a ApprovalForAll event raised by the RunbaTest contract.
type RunbaTestApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_RunbaTest *RunbaTestFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*RunbaTestApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _RunbaTest.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &RunbaTestApprovalForAllIterator{contract: _RunbaTest.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_RunbaTest *RunbaTestFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *RunbaTestApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _RunbaTest.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RunbaTestApprovalForAll)
				if err := _RunbaTest.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_RunbaTest *RunbaTestFilterer) ParseApprovalForAll(log types.Log) (*RunbaTestApprovalForAll, error) {
	event := new(RunbaTestApprovalForAll)
	if err := _RunbaTest.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RunbaTestClaimBoxIterator is returned from FilterClaimBox and is used to iterate over the raw logs and unpacked data for ClaimBox events raised by the RunbaTest contract.
type RunbaTestClaimBoxIterator struct {
	Event *RunbaTestClaimBox // Event containing the contract specifics and raw log

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
func (it *RunbaTestClaimBoxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RunbaTestClaimBox)
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
		it.Event = new(RunbaTestClaimBox)
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
func (it *RunbaTestClaimBoxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RunbaTestClaimBoxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RunbaTestClaimBox represents a ClaimBox event raised by the RunbaTest contract.
type RunbaTestClaimBox struct {
	TokenId  *big.Int
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterClaimBox is a free log retrieval operation binding the contract event 0xef3fe58350bbd6cecb2eee87220741c24603a7a5f0ed19286e3e9c1cb7a75df9.
//
// Solidity: event ClaimBox(uint256 _tokenId, address _operator)
func (_RunbaTest *RunbaTestFilterer) FilterClaimBox(opts *bind.FilterOpts) (*RunbaTestClaimBoxIterator, error) {

	logs, sub, err := _RunbaTest.contract.FilterLogs(opts, "ClaimBox")
	if err != nil {
		return nil, err
	}
	return &RunbaTestClaimBoxIterator{contract: _RunbaTest.contract, event: "ClaimBox", logs: logs, sub: sub}, nil
}

// WatchClaimBox is a free log subscription operation binding the contract event 0xef3fe58350bbd6cecb2eee87220741c24603a7a5f0ed19286e3e9c1cb7a75df9.
//
// Solidity: event ClaimBox(uint256 _tokenId, address _operator)
func (_RunbaTest *RunbaTestFilterer) WatchClaimBox(opts *bind.WatchOpts, sink chan<- *RunbaTestClaimBox) (event.Subscription, error) {

	logs, sub, err := _RunbaTest.contract.WatchLogs(opts, "ClaimBox")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RunbaTestClaimBox)
				if err := _RunbaTest.contract.UnpackLog(event, "ClaimBox", log); err != nil {
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
func (_RunbaTest *RunbaTestFilterer) ParseClaimBox(log types.Log) (*RunbaTestClaimBox, error) {
	event := new(RunbaTestClaimBox)
	if err := _RunbaTest.contract.UnpackLog(event, "ClaimBox", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RunbaTestOpenBoxIterator is returned from FilterOpenBox and is used to iterate over the raw logs and unpacked data for OpenBox events raised by the RunbaTest contract.
type RunbaTestOpenBoxIterator struct {
	Event *RunbaTestOpenBox // Event containing the contract specifics and raw log

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
func (it *RunbaTestOpenBoxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RunbaTestOpenBox)
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
		it.Event = new(RunbaTestOpenBox)
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
func (it *RunbaTestOpenBoxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RunbaTestOpenBoxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RunbaTestOpenBox represents a OpenBox event raised by the RunbaTest contract.
type RunbaTestOpenBox struct {
	TokenId  *big.Int
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOpenBox is a free log retrieval operation binding the contract event 0x1363abdc7b7584f1c5f0c1032e066b036050a3f41ea052a527297a92a34a452c.
//
// Solidity: event OpenBox(uint256 _tokenId, address _operator)
func (_RunbaTest *RunbaTestFilterer) FilterOpenBox(opts *bind.FilterOpts) (*RunbaTestOpenBoxIterator, error) {

	logs, sub, err := _RunbaTest.contract.FilterLogs(opts, "OpenBox")
	if err != nil {
		return nil, err
	}
	return &RunbaTestOpenBoxIterator{contract: _RunbaTest.contract, event: "OpenBox", logs: logs, sub: sub}, nil
}

// WatchOpenBox is a free log subscription operation binding the contract event 0x1363abdc7b7584f1c5f0c1032e066b036050a3f41ea052a527297a92a34a452c.
//
// Solidity: event OpenBox(uint256 _tokenId, address _operator)
func (_RunbaTest *RunbaTestFilterer) WatchOpenBox(opts *bind.WatchOpts, sink chan<- *RunbaTestOpenBox) (event.Subscription, error) {

	logs, sub, err := _RunbaTest.contract.WatchLogs(opts, "OpenBox")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RunbaTestOpenBox)
				if err := _RunbaTest.contract.UnpackLog(event, "OpenBox", log); err != nil {
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
func (_RunbaTest *RunbaTestFilterer) ParseOpenBox(log types.Log) (*RunbaTestOpenBox, error) {
	event := new(RunbaTestOpenBox)
	if err := _RunbaTest.contract.UnpackLog(event, "OpenBox", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RunbaTestOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RunbaTest contract.
type RunbaTestOwnershipTransferredIterator struct {
	Event *RunbaTestOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RunbaTestOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RunbaTestOwnershipTransferred)
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
		it.Event = new(RunbaTestOwnershipTransferred)
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
func (it *RunbaTestOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RunbaTestOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RunbaTestOwnershipTransferred represents a OwnershipTransferred event raised by the RunbaTest contract.
type RunbaTestOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RunbaTest *RunbaTestFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RunbaTestOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RunbaTest.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RunbaTestOwnershipTransferredIterator{contract: _RunbaTest.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RunbaTest *RunbaTestFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RunbaTestOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RunbaTest.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RunbaTestOwnershipTransferred)
				if err := _RunbaTest.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_RunbaTest *RunbaTestFilterer) ParseOwnershipTransferred(log types.Log) (*RunbaTestOwnershipTransferred, error) {
	event := new(RunbaTestOwnershipTransferred)
	if err := _RunbaTest.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RunbaTestSetBoxURIIterator is returned from FilterSetBoxURI and is used to iterate over the raw logs and unpacked data for SetBoxURI events raised by the RunbaTest contract.
type RunbaTestSetBoxURIIterator struct {
	Event *RunbaTestSetBoxURI // Event containing the contract specifics and raw log

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
func (it *RunbaTestSetBoxURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RunbaTestSetBoxURI)
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
		it.Event = new(RunbaTestSetBoxURI)
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
func (it *RunbaTestSetBoxURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RunbaTestSetBoxURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RunbaTestSetBoxURI represents a SetBoxURI event raised by the RunbaTest contract.
type RunbaTestSetBoxURI struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetBoxURI is a free log retrieval operation binding the contract event 0xb30eea395822d1634bee3d4815af251d02b453d82d259c3ccf42b4a1e65199a7.
//
// Solidity: event SetBoxURI(uint256 _tokenId)
func (_RunbaTest *RunbaTestFilterer) FilterSetBoxURI(opts *bind.FilterOpts) (*RunbaTestSetBoxURIIterator, error) {

	logs, sub, err := _RunbaTest.contract.FilterLogs(opts, "SetBoxURI")
	if err != nil {
		return nil, err
	}
	return &RunbaTestSetBoxURIIterator{contract: _RunbaTest.contract, event: "SetBoxURI", logs: logs, sub: sub}, nil
}

// WatchSetBoxURI is a free log subscription operation binding the contract event 0xb30eea395822d1634bee3d4815af251d02b453d82d259c3ccf42b4a1e65199a7.
//
// Solidity: event SetBoxURI(uint256 _tokenId)
func (_RunbaTest *RunbaTestFilterer) WatchSetBoxURI(opts *bind.WatchOpts, sink chan<- *RunbaTestSetBoxURI) (event.Subscription, error) {

	logs, sub, err := _RunbaTest.contract.WatchLogs(opts, "SetBoxURI")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RunbaTestSetBoxURI)
				if err := _RunbaTest.contract.UnpackLog(event, "SetBoxURI", log); err != nil {
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
func (_RunbaTest *RunbaTestFilterer) ParseSetBoxURI(log types.Log) (*RunbaTestSetBoxURI, error) {
	event := new(RunbaTestSetBoxURI)
	if err := _RunbaTest.contract.UnpackLog(event, "SetBoxURI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RunbaTestSetSneakerURIIterator is returned from FilterSetSneakerURI and is used to iterate over the raw logs and unpacked data for SetSneakerURI events raised by the RunbaTest contract.
type RunbaTestSetSneakerURIIterator struct {
	Event *RunbaTestSetSneakerURI // Event containing the contract specifics and raw log

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
func (it *RunbaTestSetSneakerURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RunbaTestSetSneakerURI)
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
		it.Event = new(RunbaTestSetSneakerURI)
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
func (it *RunbaTestSetSneakerURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RunbaTestSetSneakerURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RunbaTestSetSneakerURI represents a SetSneakerURI event raised by the RunbaTest contract.
type RunbaTestSetSneakerURI struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetSneakerURI is a free log retrieval operation binding the contract event 0x73c44c51395bcfbc6a4aac8ffb72347e66ca24b2d3f3b57c0c69690a6ecf45fa.
//
// Solidity: event SetSneakerURI(uint256 _tokenId)
func (_RunbaTest *RunbaTestFilterer) FilterSetSneakerURI(opts *bind.FilterOpts) (*RunbaTestSetSneakerURIIterator, error) {

	logs, sub, err := _RunbaTest.contract.FilterLogs(opts, "SetSneakerURI")
	if err != nil {
		return nil, err
	}
	return &RunbaTestSetSneakerURIIterator{contract: _RunbaTest.contract, event: "SetSneakerURI", logs: logs, sub: sub}, nil
}

// WatchSetSneakerURI is a free log subscription operation binding the contract event 0x73c44c51395bcfbc6a4aac8ffb72347e66ca24b2d3f3b57c0c69690a6ecf45fa.
//
// Solidity: event SetSneakerURI(uint256 _tokenId)
func (_RunbaTest *RunbaTestFilterer) WatchSetSneakerURI(opts *bind.WatchOpts, sink chan<- *RunbaTestSetSneakerURI) (event.Subscription, error) {

	logs, sub, err := _RunbaTest.contract.WatchLogs(opts, "SetSneakerURI")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RunbaTestSetSneakerURI)
				if err := _RunbaTest.contract.UnpackLog(event, "SetSneakerURI", log); err != nil {
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
func (_RunbaTest *RunbaTestFilterer) ParseSetSneakerURI(log types.Log) (*RunbaTestSetSneakerURI, error) {
	event := new(RunbaTestSetSneakerURI)
	if err := _RunbaTest.contract.UnpackLog(event, "SetSneakerURI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RunbaTestTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the RunbaTest contract.
type RunbaTestTransferIterator struct {
	Event *RunbaTestTransfer // Event containing the contract specifics and raw log

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
func (it *RunbaTestTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RunbaTestTransfer)
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
		it.Event = new(RunbaTestTransfer)
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
func (it *RunbaTestTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RunbaTestTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RunbaTestTransfer represents a Transfer event raised by the RunbaTest contract.
type RunbaTestTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_RunbaTest *RunbaTestFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*RunbaTestTransferIterator, error) {

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

	logs, sub, err := _RunbaTest.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &RunbaTestTransferIterator{contract: _RunbaTest.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_RunbaTest *RunbaTestFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *RunbaTestTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _RunbaTest.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RunbaTestTransfer)
				if err := _RunbaTest.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_RunbaTest *RunbaTestFilterer) ParseTransfer(log types.Log) (*RunbaTestTransfer, error) {
	event := new(RunbaTestTransfer)
	if err := _RunbaTest.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RunbaTestUpdateTokenIdIterator is returned from FilterUpdateTokenId and is used to iterate over the raw logs and unpacked data for UpdateTokenId events raised by the RunbaTest contract.
type RunbaTestUpdateTokenIdIterator struct {
	Event *RunbaTestUpdateTokenId // Event containing the contract specifics and raw log

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
func (it *RunbaTestUpdateTokenIdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RunbaTestUpdateTokenId)
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
		it.Event = new(RunbaTestUpdateTokenId)
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
func (it *RunbaTestUpdateTokenIdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RunbaTestUpdateTokenIdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RunbaTestUpdateTokenId represents a UpdateTokenId event raised by the RunbaTest contract.
type RunbaTestUpdateTokenId struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateTokenId is a free log retrieval operation binding the contract event 0x4515a5ddb22c0a35420d783b66f71e6c8b3111ae5d5891f269eb07a2a9619a68.
//
// Solidity: event UpdateTokenId(uint256 _tokenId)
func (_RunbaTest *RunbaTestFilterer) FilterUpdateTokenId(opts *bind.FilterOpts) (*RunbaTestUpdateTokenIdIterator, error) {

	logs, sub, err := _RunbaTest.contract.FilterLogs(opts, "UpdateTokenId")
	if err != nil {
		return nil, err
	}
	return &RunbaTestUpdateTokenIdIterator{contract: _RunbaTest.contract, event: "UpdateTokenId", logs: logs, sub: sub}, nil
}

// WatchUpdateTokenId is a free log subscription operation binding the contract event 0x4515a5ddb22c0a35420d783b66f71e6c8b3111ae5d5891f269eb07a2a9619a68.
//
// Solidity: event UpdateTokenId(uint256 _tokenId)
func (_RunbaTest *RunbaTestFilterer) WatchUpdateTokenId(opts *bind.WatchOpts, sink chan<- *RunbaTestUpdateTokenId) (event.Subscription, error) {

	logs, sub, err := _RunbaTest.contract.WatchLogs(opts, "UpdateTokenId")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RunbaTestUpdateTokenId)
				if err := _RunbaTest.contract.UnpackLog(event, "UpdateTokenId", log); err != nil {
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
func (_RunbaTest *RunbaTestFilterer) ParseUpdateTokenId(log types.Log) (*RunbaTestUpdateTokenId, error) {
	event := new(RunbaTestUpdateTokenId)
	if err := _RunbaTest.contract.UnpackLog(event, "UpdateTokenId", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
