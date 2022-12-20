// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package freeMint

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

// FreeMintMetaData contains all meta data concerning the FreeMint contract.
var FreeMintMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIRUNBANFT\",\"name\":\"_run\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"FMPos\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PPos\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"_boxList\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"_shoeList\",\"type\":\"string[]\"}],\"name\":\"addFreeMintData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"_boxList\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"_shoeList\",\"type\":\"string[]\"}],\"name\":\"addPublicMintData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flipSaleState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"freeClaimed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"freeMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"freeMintEnd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"freeMintStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"freeTokenIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"freeTokenIdsBoxUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"freeTokenIdsShoeUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxFreeMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTotalMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"publicClaimed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"}],\"name\":\"publicMint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publicMintEnd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publicMintPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publicMintStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"publicTokenIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"publicTokenIdsBoxUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"publicTokenIdsShoeUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"resetOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"runC\",\"outputs\":[{\"internalType\":\"contractIRUNBANFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"saleIsActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pos\",\"type\":\"uint256\"}],\"name\":\"setFMPos\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_end\",\"type\":\"uint256\"}],\"name\":\"setFreeMintTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxMint\",\"type\":\"uint256\"}],\"name\":\"setMaxFreeMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxTotalMint\",\"type\":\"uint256\"}],\"name\":\"setMaxMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_fm\",\"type\":\"bytes32\"}],\"name\":\"setMerkleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pos\",\"type\":\"uint256\"}],\"name\":\"setPPos\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setPublicMintPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_end\",\"type\":\"uint256\"}],\"name\":\"setPublicMintTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// FreeMintABI is the input ABI used to generate the binding from.
// Deprecated: Use FreeMintMetaData.ABI instead.
var FreeMintABI = FreeMintMetaData.ABI

// FreeMint is an auto generated Go binding around an Ethereum contract.
type FreeMint struct {
	FreeMintCaller     // Read-only binding to the contract
	FreeMintTransactor // Write-only binding to the contract
	FreeMintFilterer   // Log filterer for contract events
}

// FreeMintCaller is an auto generated read-only Go binding around an Ethereum contract.
type FreeMintCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FreeMintTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FreeMintTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FreeMintFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FreeMintFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FreeMintSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FreeMintSession struct {
	Contract     *FreeMint         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FreeMintCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FreeMintCallerSession struct {
	Contract *FreeMintCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// FreeMintTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FreeMintTransactorSession struct {
	Contract     *FreeMintTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// FreeMintRaw is an auto generated low-level Go binding around an Ethereum contract.
type FreeMintRaw struct {
	Contract *FreeMint // Generic contract binding to access the raw methods on
}

// FreeMintCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FreeMintCallerRaw struct {
	Contract *FreeMintCaller // Generic read-only contract binding to access the raw methods on
}

// FreeMintTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FreeMintTransactorRaw struct {
	Contract *FreeMintTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFreeMint creates a new instance of FreeMint, bound to a specific deployed contract.
func NewFreeMint(address common.Address, backend bind.ContractBackend) (*FreeMint, error) {
	contract, err := bindFreeMint(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FreeMint{FreeMintCaller: FreeMintCaller{contract: contract}, FreeMintTransactor: FreeMintTransactor{contract: contract}, FreeMintFilterer: FreeMintFilterer{contract: contract}}, nil
}

// NewFreeMintCaller creates a new read-only instance of FreeMint, bound to a specific deployed contract.
func NewFreeMintCaller(address common.Address, caller bind.ContractCaller) (*FreeMintCaller, error) {
	contract, err := bindFreeMint(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FreeMintCaller{contract: contract}, nil
}

// NewFreeMintTransactor creates a new write-only instance of FreeMint, bound to a specific deployed contract.
func NewFreeMintTransactor(address common.Address, transactor bind.ContractTransactor) (*FreeMintTransactor, error) {
	contract, err := bindFreeMint(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FreeMintTransactor{contract: contract}, nil
}

// NewFreeMintFilterer creates a new log filterer instance of FreeMint, bound to a specific deployed contract.
func NewFreeMintFilterer(address common.Address, filterer bind.ContractFilterer) (*FreeMintFilterer, error) {
	contract, err := bindFreeMint(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FreeMintFilterer{contract: contract}, nil
}

// bindFreeMint binds a generic wrapper to an already deployed contract.
func bindFreeMint(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FreeMintABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FreeMint *FreeMintRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FreeMint.Contract.FreeMintCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FreeMint *FreeMintRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreeMint.Contract.FreeMintTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FreeMint *FreeMintRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FreeMint.Contract.FreeMintTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FreeMint *FreeMintCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FreeMint.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FreeMint *FreeMintTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreeMint.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FreeMint *FreeMintTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FreeMint.Contract.contract.Transact(opts, method, params...)
}

// FMPos is a free data retrieval call binding the contract method 0x3f19f9c6.
//
// Solidity: function FMPos() view returns(uint256)
func (_FreeMint *FreeMintCaller) FMPos(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FreeMint.contract.Call(opts, &out, "FMPos")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FMPos is a free data retrieval call binding the contract method 0x3f19f9c6.
//
// Solidity: function FMPos() view returns(uint256)
func (_FreeMint *FreeMintSession) FMPos() (*big.Int, error) {
	return _FreeMint.Contract.FMPos(&_FreeMint.CallOpts)
}

// FMPos is a free data retrieval call binding the contract method 0x3f19f9c6.
//
// Solidity: function FMPos() view returns(uint256)
func (_FreeMint *FreeMintCallerSession) FMPos() (*big.Int, error) {
	return _FreeMint.Contract.FMPos(&_FreeMint.CallOpts)
}

// PPos is a free data retrieval call binding the contract method 0x87d593c7.
//
// Solidity: function PPos() view returns(uint256)
func (_FreeMint *FreeMintCaller) PPos(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FreeMint.contract.Call(opts, &out, "PPos")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PPos is a free data retrieval call binding the contract method 0x87d593c7.
//
// Solidity: function PPos() view returns(uint256)
func (_FreeMint *FreeMintSession) PPos() (*big.Int, error) {
	return _FreeMint.Contract.PPos(&_FreeMint.CallOpts)
}

// PPos is a free data retrieval call binding the contract method 0x87d593c7.
//
// Solidity: function PPos() view returns(uint256)
func (_FreeMint *FreeMintCallerSession) PPos() (*big.Int, error) {
	return _FreeMint.Contract.PPos(&_FreeMint.CallOpts)
}

// FreeClaimed is a free data retrieval call binding the contract method 0x61c0b6a0.
//
// Solidity: function freeClaimed(address ) view returns(uint256)
func (_FreeMint *FreeMintCaller) FreeClaimed(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FreeMint.contract.Call(opts, &out, "freeClaimed", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FreeClaimed is a free data retrieval call binding the contract method 0x61c0b6a0.
//
// Solidity: function freeClaimed(address ) view returns(uint256)
func (_FreeMint *FreeMintSession) FreeClaimed(arg0 common.Address) (*big.Int, error) {
	return _FreeMint.Contract.FreeClaimed(&_FreeMint.CallOpts, arg0)
}

// FreeClaimed is a free data retrieval call binding the contract method 0x61c0b6a0.
//
// Solidity: function freeClaimed(address ) view returns(uint256)
func (_FreeMint *FreeMintCallerSession) FreeClaimed(arg0 common.Address) (*big.Int, error) {
	return _FreeMint.Contract.FreeClaimed(&_FreeMint.CallOpts, arg0)
}

// FreeMintEnd is a free data retrieval call binding the contract method 0x44fdd445.
//
// Solidity: function freeMintEnd() view returns(uint256)
func (_FreeMint *FreeMintCaller) FreeMintEnd(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FreeMint.contract.Call(opts, &out, "freeMintEnd")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FreeMintEnd is a free data retrieval call binding the contract method 0x44fdd445.
//
// Solidity: function freeMintEnd() view returns(uint256)
func (_FreeMint *FreeMintSession) FreeMintEnd() (*big.Int, error) {
	return _FreeMint.Contract.FreeMintEnd(&_FreeMint.CallOpts)
}

// FreeMintEnd is a free data retrieval call binding the contract method 0x44fdd445.
//
// Solidity: function freeMintEnd() view returns(uint256)
func (_FreeMint *FreeMintCallerSession) FreeMintEnd() (*big.Int, error) {
	return _FreeMint.Contract.FreeMintEnd(&_FreeMint.CallOpts)
}

// FreeMintStart is a free data retrieval call binding the contract method 0x70706ab3.
//
// Solidity: function freeMintStart() view returns(uint256)
func (_FreeMint *FreeMintCaller) FreeMintStart(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FreeMint.contract.Call(opts, &out, "freeMintStart")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FreeMintStart is a free data retrieval call binding the contract method 0x70706ab3.
//
// Solidity: function freeMintStart() view returns(uint256)
func (_FreeMint *FreeMintSession) FreeMintStart() (*big.Int, error) {
	return _FreeMint.Contract.FreeMintStart(&_FreeMint.CallOpts)
}

// FreeMintStart is a free data retrieval call binding the contract method 0x70706ab3.
//
// Solidity: function freeMintStart() view returns(uint256)
func (_FreeMint *FreeMintCallerSession) FreeMintStart() (*big.Int, error) {
	return _FreeMint.Contract.FreeMintStart(&_FreeMint.CallOpts)
}

// FreeTokenIds is a free data retrieval call binding the contract method 0x2f827403.
//
// Solidity: function freeTokenIds(uint256 ) view returns(uint256)
func (_FreeMint *FreeMintCaller) FreeTokenIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FreeMint.contract.Call(opts, &out, "freeTokenIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FreeTokenIds is a free data retrieval call binding the contract method 0x2f827403.
//
// Solidity: function freeTokenIds(uint256 ) view returns(uint256)
func (_FreeMint *FreeMintSession) FreeTokenIds(arg0 *big.Int) (*big.Int, error) {
	return _FreeMint.Contract.FreeTokenIds(&_FreeMint.CallOpts, arg0)
}

// FreeTokenIds is a free data retrieval call binding the contract method 0x2f827403.
//
// Solidity: function freeTokenIds(uint256 ) view returns(uint256)
func (_FreeMint *FreeMintCallerSession) FreeTokenIds(arg0 *big.Int) (*big.Int, error) {
	return _FreeMint.Contract.FreeTokenIds(&_FreeMint.CallOpts, arg0)
}

// FreeTokenIdsBoxUri is a free data retrieval call binding the contract method 0x5793a713.
//
// Solidity: function freeTokenIdsBoxUri(uint256 ) view returns(string)
func (_FreeMint *FreeMintCaller) FreeTokenIdsBoxUri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _FreeMint.contract.Call(opts, &out, "freeTokenIdsBoxUri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// FreeTokenIdsBoxUri is a free data retrieval call binding the contract method 0x5793a713.
//
// Solidity: function freeTokenIdsBoxUri(uint256 ) view returns(string)
func (_FreeMint *FreeMintSession) FreeTokenIdsBoxUri(arg0 *big.Int) (string, error) {
	return _FreeMint.Contract.FreeTokenIdsBoxUri(&_FreeMint.CallOpts, arg0)
}

// FreeTokenIdsBoxUri is a free data retrieval call binding the contract method 0x5793a713.
//
// Solidity: function freeTokenIdsBoxUri(uint256 ) view returns(string)
func (_FreeMint *FreeMintCallerSession) FreeTokenIdsBoxUri(arg0 *big.Int) (string, error) {
	return _FreeMint.Contract.FreeTokenIdsBoxUri(&_FreeMint.CallOpts, arg0)
}

// FreeTokenIdsShoeUri is a free data retrieval call binding the contract method 0x0ead5b2c.
//
// Solidity: function freeTokenIdsShoeUri(uint256 ) view returns(string)
func (_FreeMint *FreeMintCaller) FreeTokenIdsShoeUri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _FreeMint.contract.Call(opts, &out, "freeTokenIdsShoeUri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// FreeTokenIdsShoeUri is a free data retrieval call binding the contract method 0x0ead5b2c.
//
// Solidity: function freeTokenIdsShoeUri(uint256 ) view returns(string)
func (_FreeMint *FreeMintSession) FreeTokenIdsShoeUri(arg0 *big.Int) (string, error) {
	return _FreeMint.Contract.FreeTokenIdsShoeUri(&_FreeMint.CallOpts, arg0)
}

// FreeTokenIdsShoeUri is a free data retrieval call binding the contract method 0x0ead5b2c.
//
// Solidity: function freeTokenIdsShoeUri(uint256 ) view returns(string)
func (_FreeMint *FreeMintCallerSession) FreeTokenIdsShoeUri(arg0 *big.Int) (string, error) {
	return _FreeMint.Contract.FreeTokenIdsShoeUri(&_FreeMint.CallOpts, arg0)
}

// MaxFreeMint is a free data retrieval call binding the contract method 0xa591252d.
//
// Solidity: function maxFreeMint() view returns(uint256)
func (_FreeMint *FreeMintCaller) MaxFreeMint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FreeMint.contract.Call(opts, &out, "maxFreeMint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxFreeMint is a free data retrieval call binding the contract method 0xa591252d.
//
// Solidity: function maxFreeMint() view returns(uint256)
func (_FreeMint *FreeMintSession) MaxFreeMint() (*big.Int, error) {
	return _FreeMint.Contract.MaxFreeMint(&_FreeMint.CallOpts)
}

// MaxFreeMint is a free data retrieval call binding the contract method 0xa591252d.
//
// Solidity: function maxFreeMint() view returns(uint256)
func (_FreeMint *FreeMintCallerSession) MaxFreeMint() (*big.Int, error) {
	return _FreeMint.Contract.MaxFreeMint(&_FreeMint.CallOpts)
}

// MaxTotalMint is a free data retrieval call binding the contract method 0xf4170e98.
//
// Solidity: function maxTotalMint() view returns(uint256)
func (_FreeMint *FreeMintCaller) MaxTotalMint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FreeMint.contract.Call(opts, &out, "maxTotalMint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTotalMint is a free data retrieval call binding the contract method 0xf4170e98.
//
// Solidity: function maxTotalMint() view returns(uint256)
func (_FreeMint *FreeMintSession) MaxTotalMint() (*big.Int, error) {
	return _FreeMint.Contract.MaxTotalMint(&_FreeMint.CallOpts)
}

// MaxTotalMint is a free data retrieval call binding the contract method 0xf4170e98.
//
// Solidity: function maxTotalMint() view returns(uint256)
func (_FreeMint *FreeMintCallerSession) MaxTotalMint() (*big.Int, error) {
	return _FreeMint.Contract.MaxTotalMint(&_FreeMint.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_FreeMint *FreeMintCaller) MerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FreeMint.contract.Call(opts, &out, "merkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_FreeMint *FreeMintSession) MerkleRoot() ([32]byte, error) {
	return _FreeMint.Contract.MerkleRoot(&_FreeMint.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_FreeMint *FreeMintCallerSession) MerkleRoot() ([32]byte, error) {
	return _FreeMint.Contract.MerkleRoot(&_FreeMint.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FreeMint *FreeMintCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FreeMint.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FreeMint *FreeMintSession) Owner() (common.Address, error) {
	return _FreeMint.Contract.Owner(&_FreeMint.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FreeMint *FreeMintCallerSession) Owner() (common.Address, error) {
	return _FreeMint.Contract.Owner(&_FreeMint.CallOpts)
}

// PublicClaimed is a free data retrieval call binding the contract method 0xb5b1cd7c.
//
// Solidity: function publicClaimed(address ) view returns(uint256)
func (_FreeMint *FreeMintCaller) PublicClaimed(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FreeMint.contract.Call(opts, &out, "publicClaimed", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PublicClaimed is a free data retrieval call binding the contract method 0xb5b1cd7c.
//
// Solidity: function publicClaimed(address ) view returns(uint256)
func (_FreeMint *FreeMintSession) PublicClaimed(arg0 common.Address) (*big.Int, error) {
	return _FreeMint.Contract.PublicClaimed(&_FreeMint.CallOpts, arg0)
}

// PublicClaimed is a free data retrieval call binding the contract method 0xb5b1cd7c.
//
// Solidity: function publicClaimed(address ) view returns(uint256)
func (_FreeMint *FreeMintCallerSession) PublicClaimed(arg0 common.Address) (*big.Int, error) {
	return _FreeMint.Contract.PublicClaimed(&_FreeMint.CallOpts, arg0)
}

// PublicMintEnd is a free data retrieval call binding the contract method 0x2a196e1b.
//
// Solidity: function publicMintEnd() view returns(uint256)
func (_FreeMint *FreeMintCaller) PublicMintEnd(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FreeMint.contract.Call(opts, &out, "publicMintEnd")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PublicMintEnd is a free data retrieval call binding the contract method 0x2a196e1b.
//
// Solidity: function publicMintEnd() view returns(uint256)
func (_FreeMint *FreeMintSession) PublicMintEnd() (*big.Int, error) {
	return _FreeMint.Contract.PublicMintEnd(&_FreeMint.CallOpts)
}

// PublicMintEnd is a free data retrieval call binding the contract method 0x2a196e1b.
//
// Solidity: function publicMintEnd() view returns(uint256)
func (_FreeMint *FreeMintCallerSession) PublicMintEnd() (*big.Int, error) {
	return _FreeMint.Contract.PublicMintEnd(&_FreeMint.CallOpts)
}

// PublicMintPrice is a free data retrieval call binding the contract method 0xdc53fd92.
//
// Solidity: function publicMintPrice() view returns(uint256)
func (_FreeMint *FreeMintCaller) PublicMintPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FreeMint.contract.Call(opts, &out, "publicMintPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PublicMintPrice is a free data retrieval call binding the contract method 0xdc53fd92.
//
// Solidity: function publicMintPrice() view returns(uint256)
func (_FreeMint *FreeMintSession) PublicMintPrice() (*big.Int, error) {
	return _FreeMint.Contract.PublicMintPrice(&_FreeMint.CallOpts)
}

// PublicMintPrice is a free data retrieval call binding the contract method 0xdc53fd92.
//
// Solidity: function publicMintPrice() view returns(uint256)
func (_FreeMint *FreeMintCallerSession) PublicMintPrice() (*big.Int, error) {
	return _FreeMint.Contract.PublicMintPrice(&_FreeMint.CallOpts)
}

// PublicMintStart is a free data retrieval call binding the contract method 0x8cfec4c0.
//
// Solidity: function publicMintStart() view returns(uint256)
func (_FreeMint *FreeMintCaller) PublicMintStart(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FreeMint.contract.Call(opts, &out, "publicMintStart")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PublicMintStart is a free data retrieval call binding the contract method 0x8cfec4c0.
//
// Solidity: function publicMintStart() view returns(uint256)
func (_FreeMint *FreeMintSession) PublicMintStart() (*big.Int, error) {
	return _FreeMint.Contract.PublicMintStart(&_FreeMint.CallOpts)
}

// PublicMintStart is a free data retrieval call binding the contract method 0x8cfec4c0.
//
// Solidity: function publicMintStart() view returns(uint256)
func (_FreeMint *FreeMintCallerSession) PublicMintStart() (*big.Int, error) {
	return _FreeMint.Contract.PublicMintStart(&_FreeMint.CallOpts)
}

// PublicTokenIds is a free data retrieval call binding the contract method 0xf5653f4e.
//
// Solidity: function publicTokenIds(uint256 ) view returns(uint256)
func (_FreeMint *FreeMintCaller) PublicTokenIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FreeMint.contract.Call(opts, &out, "publicTokenIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PublicTokenIds is a free data retrieval call binding the contract method 0xf5653f4e.
//
// Solidity: function publicTokenIds(uint256 ) view returns(uint256)
func (_FreeMint *FreeMintSession) PublicTokenIds(arg0 *big.Int) (*big.Int, error) {
	return _FreeMint.Contract.PublicTokenIds(&_FreeMint.CallOpts, arg0)
}

// PublicTokenIds is a free data retrieval call binding the contract method 0xf5653f4e.
//
// Solidity: function publicTokenIds(uint256 ) view returns(uint256)
func (_FreeMint *FreeMintCallerSession) PublicTokenIds(arg0 *big.Int) (*big.Int, error) {
	return _FreeMint.Contract.PublicTokenIds(&_FreeMint.CallOpts, arg0)
}

// PublicTokenIdsBoxUri is a free data retrieval call binding the contract method 0xc2bcab10.
//
// Solidity: function publicTokenIdsBoxUri(uint256 ) view returns(string)
func (_FreeMint *FreeMintCaller) PublicTokenIdsBoxUri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _FreeMint.contract.Call(opts, &out, "publicTokenIdsBoxUri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PublicTokenIdsBoxUri is a free data retrieval call binding the contract method 0xc2bcab10.
//
// Solidity: function publicTokenIdsBoxUri(uint256 ) view returns(string)
func (_FreeMint *FreeMintSession) PublicTokenIdsBoxUri(arg0 *big.Int) (string, error) {
	return _FreeMint.Contract.PublicTokenIdsBoxUri(&_FreeMint.CallOpts, arg0)
}

// PublicTokenIdsBoxUri is a free data retrieval call binding the contract method 0xc2bcab10.
//
// Solidity: function publicTokenIdsBoxUri(uint256 ) view returns(string)
func (_FreeMint *FreeMintCallerSession) PublicTokenIdsBoxUri(arg0 *big.Int) (string, error) {
	return _FreeMint.Contract.PublicTokenIdsBoxUri(&_FreeMint.CallOpts, arg0)
}

// PublicTokenIdsShoeUri is a free data retrieval call binding the contract method 0xb60f78eb.
//
// Solidity: function publicTokenIdsShoeUri(uint256 ) view returns(string)
func (_FreeMint *FreeMintCaller) PublicTokenIdsShoeUri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _FreeMint.contract.Call(opts, &out, "publicTokenIdsShoeUri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PublicTokenIdsShoeUri is a free data retrieval call binding the contract method 0xb60f78eb.
//
// Solidity: function publicTokenIdsShoeUri(uint256 ) view returns(string)
func (_FreeMint *FreeMintSession) PublicTokenIdsShoeUri(arg0 *big.Int) (string, error) {
	return _FreeMint.Contract.PublicTokenIdsShoeUri(&_FreeMint.CallOpts, arg0)
}

// PublicTokenIdsShoeUri is a free data retrieval call binding the contract method 0xb60f78eb.
//
// Solidity: function publicTokenIdsShoeUri(uint256 ) view returns(string)
func (_FreeMint *FreeMintCallerSession) PublicTokenIdsShoeUri(arg0 *big.Int) (string, error) {
	return _FreeMint.Contract.PublicTokenIdsShoeUri(&_FreeMint.CallOpts, arg0)
}

// RunC is a free data retrieval call binding the contract method 0xa664d160.
//
// Solidity: function runC() view returns(address)
func (_FreeMint *FreeMintCaller) RunC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FreeMint.contract.Call(opts, &out, "runC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RunC is a free data retrieval call binding the contract method 0xa664d160.
//
// Solidity: function runC() view returns(address)
func (_FreeMint *FreeMintSession) RunC() (common.Address, error) {
	return _FreeMint.Contract.RunC(&_FreeMint.CallOpts)
}

// RunC is a free data retrieval call binding the contract method 0xa664d160.
//
// Solidity: function runC() view returns(address)
func (_FreeMint *FreeMintCallerSession) RunC() (common.Address, error) {
	return _FreeMint.Contract.RunC(&_FreeMint.CallOpts)
}

// SaleIsActive is a free data retrieval call binding the contract method 0xeb8d2444.
//
// Solidity: function saleIsActive() view returns(bool)
func (_FreeMint *FreeMintCaller) SaleIsActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FreeMint.contract.Call(opts, &out, "saleIsActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SaleIsActive is a free data retrieval call binding the contract method 0xeb8d2444.
//
// Solidity: function saleIsActive() view returns(bool)
func (_FreeMint *FreeMintSession) SaleIsActive() (bool, error) {
	return _FreeMint.Contract.SaleIsActive(&_FreeMint.CallOpts)
}

// SaleIsActive is a free data retrieval call binding the contract method 0xeb8d2444.
//
// Solidity: function saleIsActive() view returns(bool)
func (_FreeMint *FreeMintCallerSession) SaleIsActive() (bool, error) {
	return _FreeMint.Contract.SaleIsActive(&_FreeMint.CallOpts)
}

// AddFreeMintData is a paid mutator transaction binding the contract method 0x3ce375e0.
//
// Solidity: function addFreeMintData(uint256[] _tokenIds, string[] _boxList, string[] _shoeList) returns()
func (_FreeMint *FreeMintTransactor) AddFreeMintData(opts *bind.TransactOpts, _tokenIds []*big.Int, _boxList []string, _shoeList []string) (*types.Transaction, error) {
	return _FreeMint.contract.Transact(opts, "addFreeMintData", _tokenIds, _boxList, _shoeList)
}

// AddFreeMintData is a paid mutator transaction binding the contract method 0x3ce375e0.
//
// Solidity: function addFreeMintData(uint256[] _tokenIds, string[] _boxList, string[] _shoeList) returns()
func (_FreeMint *FreeMintSession) AddFreeMintData(_tokenIds []*big.Int, _boxList []string, _shoeList []string) (*types.Transaction, error) {
	return _FreeMint.Contract.AddFreeMintData(&_FreeMint.TransactOpts, _tokenIds, _boxList, _shoeList)
}

// AddFreeMintData is a paid mutator transaction binding the contract method 0x3ce375e0.
//
// Solidity: function addFreeMintData(uint256[] _tokenIds, string[] _boxList, string[] _shoeList) returns()
func (_FreeMint *FreeMintTransactorSession) AddFreeMintData(_tokenIds []*big.Int, _boxList []string, _shoeList []string) (*types.Transaction, error) {
	return _FreeMint.Contract.AddFreeMintData(&_FreeMint.TransactOpts, _tokenIds, _boxList, _shoeList)
}

// AddPublicMintData is a paid mutator transaction binding the contract method 0xebfcc606.
//
// Solidity: function addPublicMintData(uint256[] _tokenIds, string[] _boxList, string[] _shoeList) returns()
func (_FreeMint *FreeMintTransactor) AddPublicMintData(opts *bind.TransactOpts, _tokenIds []*big.Int, _boxList []string, _shoeList []string) (*types.Transaction, error) {
	return _FreeMint.contract.Transact(opts, "addPublicMintData", _tokenIds, _boxList, _shoeList)
}

// AddPublicMintData is a paid mutator transaction binding the contract method 0xebfcc606.
//
// Solidity: function addPublicMintData(uint256[] _tokenIds, string[] _boxList, string[] _shoeList) returns()
func (_FreeMint *FreeMintSession) AddPublicMintData(_tokenIds []*big.Int, _boxList []string, _shoeList []string) (*types.Transaction, error) {
	return _FreeMint.Contract.AddPublicMintData(&_FreeMint.TransactOpts, _tokenIds, _boxList, _shoeList)
}

// AddPublicMintData is a paid mutator transaction binding the contract method 0xebfcc606.
//
// Solidity: function addPublicMintData(uint256[] _tokenIds, string[] _boxList, string[] _shoeList) returns()
func (_FreeMint *FreeMintTransactorSession) AddPublicMintData(_tokenIds []*big.Int, _boxList []string, _shoeList []string) (*types.Transaction, error) {
	return _FreeMint.Contract.AddPublicMintData(&_FreeMint.TransactOpts, _tokenIds, _boxList, _shoeList)
}

// FlipSaleState is a paid mutator transaction binding the contract method 0x34918dfd.
//
// Solidity: function flipSaleState() returns()
func (_FreeMint *FreeMintTransactor) FlipSaleState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreeMint.contract.Transact(opts, "flipSaleState")
}

// FlipSaleState is a paid mutator transaction binding the contract method 0x34918dfd.
//
// Solidity: function flipSaleState() returns()
func (_FreeMint *FreeMintSession) FlipSaleState() (*types.Transaction, error) {
	return _FreeMint.Contract.FlipSaleState(&_FreeMint.TransactOpts)
}

// FlipSaleState is a paid mutator transaction binding the contract method 0x34918dfd.
//
// Solidity: function flipSaleState() returns()
func (_FreeMint *FreeMintTransactorSession) FlipSaleState() (*types.Transaction, error) {
	return _FreeMint.Contract.FlipSaleState(&_FreeMint.TransactOpts)
}

// FreeMint is a paid mutator transaction binding the contract method 0x88d15d50.
//
// Solidity: function freeMint(bytes32[] _merkleProof) returns()
func (_FreeMint *FreeMintTransactor) FreeMint(opts *bind.TransactOpts, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _FreeMint.contract.Transact(opts, "freeMint", _merkleProof)
}

// FreeMint is a paid mutator transaction binding the contract method 0x88d15d50.
//
// Solidity: function freeMint(bytes32[] _merkleProof) returns()
func (_FreeMint *FreeMintSession) FreeMint(_merkleProof [][32]byte) (*types.Transaction, error) {
	return _FreeMint.Contract.FreeMint(&_FreeMint.TransactOpts, _merkleProof)
}

// FreeMint is a paid mutator transaction binding the contract method 0x88d15d50.
//
// Solidity: function freeMint(bytes32[] _merkleProof) returns()
func (_FreeMint *FreeMintTransactorSession) FreeMint(_merkleProof [][32]byte) (*types.Transaction, error) {
	return _FreeMint.Contract.FreeMint(&_FreeMint.TransactOpts, _merkleProof)
}

// PublicMint is a paid mutator transaction binding the contract method 0x2db11544.
//
// Solidity: function publicMint(uint256 _quantity) payable returns()
func (_FreeMint *FreeMintTransactor) PublicMint(opts *bind.TransactOpts, _quantity *big.Int) (*types.Transaction, error) {
	return _FreeMint.contract.Transact(opts, "publicMint", _quantity)
}

// PublicMint is a paid mutator transaction binding the contract method 0x2db11544.
//
// Solidity: function publicMint(uint256 _quantity) payable returns()
func (_FreeMint *FreeMintSession) PublicMint(_quantity *big.Int) (*types.Transaction, error) {
	return _FreeMint.Contract.PublicMint(&_FreeMint.TransactOpts, _quantity)
}

// PublicMint is a paid mutator transaction binding the contract method 0x2db11544.
//
// Solidity: function publicMint(uint256 _quantity) payable returns()
func (_FreeMint *FreeMintTransactorSession) PublicMint(_quantity *big.Int) (*types.Transaction, error) {
	return _FreeMint.Contract.PublicMint(&_FreeMint.TransactOpts, _quantity)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FreeMint *FreeMintTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreeMint.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FreeMint *FreeMintSession) RenounceOwnership() (*types.Transaction, error) {
	return _FreeMint.Contract.RenounceOwnership(&_FreeMint.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FreeMint *FreeMintTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FreeMint.Contract.RenounceOwnership(&_FreeMint.TransactOpts)
}

// ResetOwnership is a paid mutator transaction binding the contract method 0xbe3af789.
//
// Solidity: function resetOwnership(address _newOwner) returns()
func (_FreeMint *FreeMintTransactor) ResetOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _FreeMint.contract.Transact(opts, "resetOwnership", _newOwner)
}

// ResetOwnership is a paid mutator transaction binding the contract method 0xbe3af789.
//
// Solidity: function resetOwnership(address _newOwner) returns()
func (_FreeMint *FreeMintSession) ResetOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _FreeMint.Contract.ResetOwnership(&_FreeMint.TransactOpts, _newOwner)
}

// ResetOwnership is a paid mutator transaction binding the contract method 0xbe3af789.
//
// Solidity: function resetOwnership(address _newOwner) returns()
func (_FreeMint *FreeMintTransactorSession) ResetOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _FreeMint.Contract.ResetOwnership(&_FreeMint.TransactOpts, _newOwner)
}

// SetFMPos is a paid mutator transaction binding the contract method 0xac9421ca.
//
// Solidity: function setFMPos(uint256 _pos) returns()
func (_FreeMint *FreeMintTransactor) SetFMPos(opts *bind.TransactOpts, _pos *big.Int) (*types.Transaction, error) {
	return _FreeMint.contract.Transact(opts, "setFMPos", _pos)
}

// SetFMPos is a paid mutator transaction binding the contract method 0xac9421ca.
//
// Solidity: function setFMPos(uint256 _pos) returns()
func (_FreeMint *FreeMintSession) SetFMPos(_pos *big.Int) (*types.Transaction, error) {
	return _FreeMint.Contract.SetFMPos(&_FreeMint.TransactOpts, _pos)
}

// SetFMPos is a paid mutator transaction binding the contract method 0xac9421ca.
//
// Solidity: function setFMPos(uint256 _pos) returns()
func (_FreeMint *FreeMintTransactorSession) SetFMPos(_pos *big.Int) (*types.Transaction, error) {
	return _FreeMint.Contract.SetFMPos(&_FreeMint.TransactOpts, _pos)
}

// SetFreeMintTime is a paid mutator transaction binding the contract method 0xf26b8119.
//
// Solidity: function setFreeMintTime(uint256 _start, uint256 _end) returns()
func (_FreeMint *FreeMintTransactor) SetFreeMintTime(opts *bind.TransactOpts, _start *big.Int, _end *big.Int) (*types.Transaction, error) {
	return _FreeMint.contract.Transact(opts, "setFreeMintTime", _start, _end)
}

// SetFreeMintTime is a paid mutator transaction binding the contract method 0xf26b8119.
//
// Solidity: function setFreeMintTime(uint256 _start, uint256 _end) returns()
func (_FreeMint *FreeMintSession) SetFreeMintTime(_start *big.Int, _end *big.Int) (*types.Transaction, error) {
	return _FreeMint.Contract.SetFreeMintTime(&_FreeMint.TransactOpts, _start, _end)
}

// SetFreeMintTime is a paid mutator transaction binding the contract method 0xf26b8119.
//
// Solidity: function setFreeMintTime(uint256 _start, uint256 _end) returns()
func (_FreeMint *FreeMintTransactorSession) SetFreeMintTime(_start *big.Int, _end *big.Int) (*types.Transaction, error) {
	return _FreeMint.Contract.SetFreeMintTime(&_FreeMint.TransactOpts, _start, _end)
}

// SetMaxFreeMint is a paid mutator transaction binding the contract method 0x742a4c9b.
//
// Solidity: function setMaxFreeMint(uint256 _maxMint) returns()
func (_FreeMint *FreeMintTransactor) SetMaxFreeMint(opts *bind.TransactOpts, _maxMint *big.Int) (*types.Transaction, error) {
	return _FreeMint.contract.Transact(opts, "setMaxFreeMint", _maxMint)
}

// SetMaxFreeMint is a paid mutator transaction binding the contract method 0x742a4c9b.
//
// Solidity: function setMaxFreeMint(uint256 _maxMint) returns()
func (_FreeMint *FreeMintSession) SetMaxFreeMint(_maxMint *big.Int) (*types.Transaction, error) {
	return _FreeMint.Contract.SetMaxFreeMint(&_FreeMint.TransactOpts, _maxMint)
}

// SetMaxFreeMint is a paid mutator transaction binding the contract method 0x742a4c9b.
//
// Solidity: function setMaxFreeMint(uint256 _maxMint) returns()
func (_FreeMint *FreeMintTransactorSession) SetMaxFreeMint(_maxMint *big.Int) (*types.Transaction, error) {
	return _FreeMint.Contract.SetMaxFreeMint(&_FreeMint.TransactOpts, _maxMint)
}

// SetMaxMint is a paid mutator transaction binding the contract method 0x547520fe.
//
// Solidity: function setMaxMint(uint256 _maxTotalMint) returns()
func (_FreeMint *FreeMintTransactor) SetMaxMint(opts *bind.TransactOpts, _maxTotalMint *big.Int) (*types.Transaction, error) {
	return _FreeMint.contract.Transact(opts, "setMaxMint", _maxTotalMint)
}

// SetMaxMint is a paid mutator transaction binding the contract method 0x547520fe.
//
// Solidity: function setMaxMint(uint256 _maxTotalMint) returns()
func (_FreeMint *FreeMintSession) SetMaxMint(_maxTotalMint *big.Int) (*types.Transaction, error) {
	return _FreeMint.Contract.SetMaxMint(&_FreeMint.TransactOpts, _maxTotalMint)
}

// SetMaxMint is a paid mutator transaction binding the contract method 0x547520fe.
//
// Solidity: function setMaxMint(uint256 _maxTotalMint) returns()
func (_FreeMint *FreeMintTransactorSession) SetMaxMint(_maxTotalMint *big.Int) (*types.Transaction, error) {
	return _FreeMint.Contract.SetMaxMint(&_FreeMint.TransactOpts, _maxTotalMint)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x7cb64759.
//
// Solidity: function setMerkleRoot(bytes32 _fm) returns()
func (_FreeMint *FreeMintTransactor) SetMerkleRoot(opts *bind.TransactOpts, _fm [32]byte) (*types.Transaction, error) {
	return _FreeMint.contract.Transact(opts, "setMerkleRoot", _fm)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x7cb64759.
//
// Solidity: function setMerkleRoot(bytes32 _fm) returns()
func (_FreeMint *FreeMintSession) SetMerkleRoot(_fm [32]byte) (*types.Transaction, error) {
	return _FreeMint.Contract.SetMerkleRoot(&_FreeMint.TransactOpts, _fm)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x7cb64759.
//
// Solidity: function setMerkleRoot(bytes32 _fm) returns()
func (_FreeMint *FreeMintTransactorSession) SetMerkleRoot(_fm [32]byte) (*types.Transaction, error) {
	return _FreeMint.Contract.SetMerkleRoot(&_FreeMint.TransactOpts, _fm)
}

// SetPPos is a paid mutator transaction binding the contract method 0xbc5bca27.
//
// Solidity: function setPPos(uint256 _pos) returns()
func (_FreeMint *FreeMintTransactor) SetPPos(opts *bind.TransactOpts, _pos *big.Int) (*types.Transaction, error) {
	return _FreeMint.contract.Transact(opts, "setPPos", _pos)
}

// SetPPos is a paid mutator transaction binding the contract method 0xbc5bca27.
//
// Solidity: function setPPos(uint256 _pos) returns()
func (_FreeMint *FreeMintSession) SetPPos(_pos *big.Int) (*types.Transaction, error) {
	return _FreeMint.Contract.SetPPos(&_FreeMint.TransactOpts, _pos)
}

// SetPPos is a paid mutator transaction binding the contract method 0xbc5bca27.
//
// Solidity: function setPPos(uint256 _pos) returns()
func (_FreeMint *FreeMintTransactorSession) SetPPos(_pos *big.Int) (*types.Transaction, error) {
	return _FreeMint.Contract.SetPPos(&_FreeMint.TransactOpts, _pos)
}

// SetPublicMintPrice is a paid mutator transaction binding the contract method 0x5d82cf6e.
//
// Solidity: function setPublicMintPrice(uint256 _price) returns()
func (_FreeMint *FreeMintTransactor) SetPublicMintPrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _FreeMint.contract.Transact(opts, "setPublicMintPrice", _price)
}

// SetPublicMintPrice is a paid mutator transaction binding the contract method 0x5d82cf6e.
//
// Solidity: function setPublicMintPrice(uint256 _price) returns()
func (_FreeMint *FreeMintSession) SetPublicMintPrice(_price *big.Int) (*types.Transaction, error) {
	return _FreeMint.Contract.SetPublicMintPrice(&_FreeMint.TransactOpts, _price)
}

// SetPublicMintPrice is a paid mutator transaction binding the contract method 0x5d82cf6e.
//
// Solidity: function setPublicMintPrice(uint256 _price) returns()
func (_FreeMint *FreeMintTransactorSession) SetPublicMintPrice(_price *big.Int) (*types.Transaction, error) {
	return _FreeMint.Contract.SetPublicMintPrice(&_FreeMint.TransactOpts, _price)
}

// SetPublicMintTime is a paid mutator transaction binding the contract method 0x911690e2.
//
// Solidity: function setPublicMintTime(uint256 _start, uint256 _end) returns()
func (_FreeMint *FreeMintTransactor) SetPublicMintTime(opts *bind.TransactOpts, _start *big.Int, _end *big.Int) (*types.Transaction, error) {
	return _FreeMint.contract.Transact(opts, "setPublicMintTime", _start, _end)
}

// SetPublicMintTime is a paid mutator transaction binding the contract method 0x911690e2.
//
// Solidity: function setPublicMintTime(uint256 _start, uint256 _end) returns()
func (_FreeMint *FreeMintSession) SetPublicMintTime(_start *big.Int, _end *big.Int) (*types.Transaction, error) {
	return _FreeMint.Contract.SetPublicMintTime(&_FreeMint.TransactOpts, _start, _end)
}

// SetPublicMintTime is a paid mutator transaction binding the contract method 0x911690e2.
//
// Solidity: function setPublicMintTime(uint256 _start, uint256 _end) returns()
func (_FreeMint *FreeMintTransactorSession) SetPublicMintTime(_start *big.Int, _end *big.Int) (*types.Transaction, error) {
	return _FreeMint.Contract.SetPublicMintTime(&_FreeMint.TransactOpts, _start, _end)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FreeMint *FreeMintTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FreeMint.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FreeMint *FreeMintSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FreeMint.Contract.TransferOwnership(&_FreeMint.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FreeMint *FreeMintTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FreeMint.Contract.TransferOwnership(&_FreeMint.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_FreeMint *FreeMintTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreeMint.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_FreeMint *FreeMintSession) Withdraw() (*types.Transaction, error) {
	return _FreeMint.Contract.Withdraw(&_FreeMint.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_FreeMint *FreeMintTransactorSession) Withdraw() (*types.Transaction, error) {
	return _FreeMint.Contract.Withdraw(&_FreeMint.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_FreeMint *FreeMintTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _FreeMint.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_FreeMint *FreeMintSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _FreeMint.Contract.Fallback(&_FreeMint.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_FreeMint *FreeMintTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _FreeMint.Contract.Fallback(&_FreeMint.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FreeMint *FreeMintTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FreeMint.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FreeMint *FreeMintSession) Receive() (*types.Transaction, error) {
	return _FreeMint.Contract.Receive(&_FreeMint.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FreeMint *FreeMintTransactorSession) Receive() (*types.Transaction, error) {
	return _FreeMint.Contract.Receive(&_FreeMint.TransactOpts)
}

// FreeMintOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FreeMint contract.
type FreeMintOwnershipTransferredIterator struct {
	Event *FreeMintOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FreeMintOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FreeMintOwnershipTransferred)
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
		it.Event = new(FreeMintOwnershipTransferred)
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
func (it *FreeMintOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FreeMintOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FreeMintOwnershipTransferred represents a OwnershipTransferred event raised by the FreeMint contract.
type FreeMintOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FreeMint *FreeMintFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FreeMintOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FreeMint.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FreeMintOwnershipTransferredIterator{contract: _FreeMint.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FreeMint *FreeMintFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FreeMintOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FreeMint.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FreeMintOwnershipTransferred)
				if err := _FreeMint.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FreeMint *FreeMintFilterer) ParseOwnershipTransferred(log types.Log) (*FreeMintOwnershipTransferred, error) {
	event := new(FreeMintOwnershipTransferred)
	if err := _FreeMint.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
