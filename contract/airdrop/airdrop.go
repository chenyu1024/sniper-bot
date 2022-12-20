// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package airdrop

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

// AirdropMetaData contains all meta data concerning the Airdrop contract.
var AirdropMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"ClaimBox\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIdList\",\"type\":\"uint256[]\"}],\"name\":\"addAirdropTokenIds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"airdropList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"getAirdrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pos\",\"type\":\"uint256\"}],\"name\":\"getAirdropList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurAirdropPos\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNFTContractAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mintEnable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"runXAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"setMerkleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_enable\",\"type\":\"bool\"}],\"name\":\"setMintEnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setNftAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pos\",\"type\":\"uint256\"}],\"name\":\"setPos\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AirdropABI is the input ABI used to generate the binding from.
// Deprecated: Use AirdropMetaData.ABI instead.
var AirdropABI = AirdropMetaData.ABI

// Airdrop is an auto generated Go binding around an Ethereum contract.
type Airdrop struct {
	AirdropCaller     // Read-only binding to the contract
	AirdropTransactor // Write-only binding to the contract
	AirdropFilterer   // Log filterer for contract events
}

// AirdropCaller is an auto generated read-only Go binding around an Ethereum contract.
type AirdropCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AirdropTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AirdropTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AirdropFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AirdropFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AirdropSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AirdropSession struct {
	Contract     *Airdrop          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AirdropCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AirdropCallerSession struct {
	Contract *AirdropCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AirdropTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AirdropTransactorSession struct {
	Contract     *AirdropTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AirdropRaw is an auto generated low-level Go binding around an Ethereum contract.
type AirdropRaw struct {
	Contract *Airdrop // Generic contract binding to access the raw methods on
}

// AirdropCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AirdropCallerRaw struct {
	Contract *AirdropCaller // Generic read-only contract binding to access the raw methods on
}

// AirdropTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AirdropTransactorRaw struct {
	Contract *AirdropTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAirdrop creates a new instance of Airdrop, bound to a specific deployed contract.
func NewAirdrop(address common.Address, backend bind.ContractBackend) (*Airdrop, error) {
	contract, err := bindAirdrop(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Airdrop{AirdropCaller: AirdropCaller{contract: contract}, AirdropTransactor: AirdropTransactor{contract: contract}, AirdropFilterer: AirdropFilterer{contract: contract}}, nil
}

// NewAirdropCaller creates a new read-only instance of Airdrop, bound to a specific deployed contract.
func NewAirdropCaller(address common.Address, caller bind.ContractCaller) (*AirdropCaller, error) {
	contract, err := bindAirdrop(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AirdropCaller{contract: contract}, nil
}

// NewAirdropTransactor creates a new write-only instance of Airdrop, bound to a specific deployed contract.
func NewAirdropTransactor(address common.Address, transactor bind.ContractTransactor) (*AirdropTransactor, error) {
	contract, err := bindAirdrop(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AirdropTransactor{contract: contract}, nil
}

// NewAirdropFilterer creates a new log filterer instance of Airdrop, bound to a specific deployed contract.
func NewAirdropFilterer(address common.Address, filterer bind.ContractFilterer) (*AirdropFilterer, error) {
	contract, err := bindAirdrop(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AirdropFilterer{contract: contract}, nil
}

// bindAirdrop binds a generic wrapper to an already deployed contract.
func bindAirdrop(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AirdropABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Airdrop *AirdropRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Airdrop.Contract.AirdropCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Airdrop *AirdropRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Airdrop.Contract.AirdropTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Airdrop *AirdropRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Airdrop.Contract.AirdropTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Airdrop *AirdropCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Airdrop.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Airdrop *AirdropTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Airdrop.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Airdrop *AirdropTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Airdrop.Contract.contract.Transact(opts, method, params...)
}

// AirdropList is a free data retrieval call binding the contract method 0xcb81b203.
//
// Solidity: function airdropList(uint256 ) view returns(uint256)
func (_Airdrop *AirdropCaller) AirdropList(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "airdropList", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AirdropList is a free data retrieval call binding the contract method 0xcb81b203.
//
// Solidity: function airdropList(uint256 ) view returns(uint256)
func (_Airdrop *AirdropSession) AirdropList(arg0 *big.Int) (*big.Int, error) {
	return _Airdrop.Contract.AirdropList(&_Airdrop.CallOpts, arg0)
}

// AirdropList is a free data retrieval call binding the contract method 0xcb81b203.
//
// Solidity: function airdropList(uint256 ) view returns(uint256)
func (_Airdrop *AirdropCallerSession) AirdropList(arg0 *big.Int) (*big.Int, error) {
	return _Airdrop.Contract.AirdropList(&_Airdrop.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Airdrop *AirdropCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Airdrop *AirdropSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Airdrop.Contract.BalanceOf(&_Airdrop.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Airdrop *AirdropCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Airdrop.Contract.BalanceOf(&_Airdrop.CallOpts, owner)
}

// GetAirdropList is a free data retrieval call binding the contract method 0x7b5d1221.
//
// Solidity: function getAirdropList(uint256 _pos) view returns(uint256)
func (_Airdrop *AirdropCaller) GetAirdropList(opts *bind.CallOpts, _pos *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "getAirdropList", _pos)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAirdropList is a free data retrieval call binding the contract method 0x7b5d1221.
//
// Solidity: function getAirdropList(uint256 _pos) view returns(uint256)
func (_Airdrop *AirdropSession) GetAirdropList(_pos *big.Int) (*big.Int, error) {
	return _Airdrop.Contract.GetAirdropList(&_Airdrop.CallOpts, _pos)
}

// GetAirdropList is a free data retrieval call binding the contract method 0x7b5d1221.
//
// Solidity: function getAirdropList(uint256 _pos) view returns(uint256)
func (_Airdrop *AirdropCallerSession) GetAirdropList(_pos *big.Int) (*big.Int, error) {
	return _Airdrop.Contract.GetAirdropList(&_Airdrop.CallOpts, _pos)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Airdrop *AirdropCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Airdrop *AirdropSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Airdrop.Contract.GetApproved(&_Airdrop.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Airdrop *AirdropCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Airdrop.Contract.GetApproved(&_Airdrop.CallOpts, tokenId)
}

// GetCurAirdropPos is a free data retrieval call binding the contract method 0x1aaf0261.
//
// Solidity: function getCurAirdropPos() view returns(uint256)
func (_Airdrop *AirdropCaller) GetCurAirdropPos(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "getCurAirdropPos")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurAirdropPos is a free data retrieval call binding the contract method 0x1aaf0261.
//
// Solidity: function getCurAirdropPos() view returns(uint256)
func (_Airdrop *AirdropSession) GetCurAirdropPos() (*big.Int, error) {
	return _Airdrop.Contract.GetCurAirdropPos(&_Airdrop.CallOpts)
}

// GetCurAirdropPos is a free data retrieval call binding the contract method 0x1aaf0261.
//
// Solidity: function getCurAirdropPos() view returns(uint256)
func (_Airdrop *AirdropCallerSession) GetCurAirdropPos() (*big.Int, error) {
	return _Airdrop.Contract.GetCurAirdropPos(&_Airdrop.CallOpts)
}

// GetNFTContractAddr is a free data retrieval call binding the contract method 0x7ad345d3.
//
// Solidity: function getNFTContractAddr() view returns(address)
func (_Airdrop *AirdropCaller) GetNFTContractAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "getNFTContractAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNFTContractAddr is a free data retrieval call binding the contract method 0x7ad345d3.
//
// Solidity: function getNFTContractAddr() view returns(address)
func (_Airdrop *AirdropSession) GetNFTContractAddr() (common.Address, error) {
	return _Airdrop.Contract.GetNFTContractAddr(&_Airdrop.CallOpts)
}

// GetNFTContractAddr is a free data retrieval call binding the contract method 0x7ad345d3.
//
// Solidity: function getNFTContractAddr() view returns(address)
func (_Airdrop *AirdropCallerSession) GetNFTContractAddr() (common.Address, error) {
	return _Airdrop.Contract.GetNFTContractAddr(&_Airdrop.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Airdrop *AirdropCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Airdrop *AirdropSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Airdrop.Contract.IsApprovedForAll(&_Airdrop.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Airdrop *AirdropCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Airdrop.Contract.IsApprovedForAll(&_Airdrop.CallOpts, owner, operator)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_Airdrop *AirdropCaller) MerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "merkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_Airdrop *AirdropSession) MerkleRoot() ([32]byte, error) {
	return _Airdrop.Contract.MerkleRoot(&_Airdrop.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_Airdrop *AirdropCallerSession) MerkleRoot() ([32]byte, error) {
	return _Airdrop.Contract.MerkleRoot(&_Airdrop.CallOpts)
}

// MintEnable is a free data retrieval call binding the contract method 0xc1e1ce29.
//
// Solidity: function mintEnable() view returns(bool)
func (_Airdrop *AirdropCaller) MintEnable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "mintEnable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MintEnable is a free data retrieval call binding the contract method 0xc1e1ce29.
//
// Solidity: function mintEnable() view returns(bool)
func (_Airdrop *AirdropSession) MintEnable() (bool, error) {
	return _Airdrop.Contract.MintEnable(&_Airdrop.CallOpts)
}

// MintEnable is a free data retrieval call binding the contract method 0xc1e1ce29.
//
// Solidity: function mintEnable() view returns(bool)
func (_Airdrop *AirdropCallerSession) MintEnable() (bool, error) {
	return _Airdrop.Contract.MintEnable(&_Airdrop.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Airdrop *AirdropCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Airdrop *AirdropSession) Name() (string, error) {
	return _Airdrop.Contract.Name(&_Airdrop.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Airdrop *AirdropCallerSession) Name() (string, error) {
	return _Airdrop.Contract.Name(&_Airdrop.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Airdrop *AirdropCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Airdrop *AirdropSession) Owner() (common.Address, error) {
	return _Airdrop.Contract.Owner(&_Airdrop.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Airdrop *AirdropCallerSession) Owner() (common.Address, error) {
	return _Airdrop.Contract.Owner(&_Airdrop.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address)
func (_Airdrop *AirdropCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "ownerOf", _tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address)
func (_Airdrop *AirdropSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _Airdrop.Contract.OwnerOf(&_Airdrop.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address)
func (_Airdrop *AirdropCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _Airdrop.Contract.OwnerOf(&_Airdrop.CallOpts, _tokenId)
}

// RunXAddress is a free data retrieval call binding the contract method 0xd9bb6dcd.
//
// Solidity: function runXAddress() view returns(address)
func (_Airdrop *AirdropCaller) RunXAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "runXAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RunXAddress is a free data retrieval call binding the contract method 0xd9bb6dcd.
//
// Solidity: function runXAddress() view returns(address)
func (_Airdrop *AirdropSession) RunXAddress() (common.Address, error) {
	return _Airdrop.Contract.RunXAddress(&_Airdrop.CallOpts)
}

// RunXAddress is a free data retrieval call binding the contract method 0xd9bb6dcd.
//
// Solidity: function runXAddress() view returns(address)
func (_Airdrop *AirdropCallerSession) RunXAddress() (common.Address, error) {
	return _Airdrop.Contract.RunXAddress(&_Airdrop.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Airdrop *AirdropCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Airdrop *AirdropSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Airdrop.Contract.SupportsInterface(&_Airdrop.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Airdrop *AirdropCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Airdrop.Contract.SupportsInterface(&_Airdrop.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Airdrop *AirdropCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Airdrop *AirdropSession) Symbol() (string, error) {
	return _Airdrop.Contract.Symbol(&_Airdrop.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Airdrop *AirdropCallerSession) Symbol() (string, error) {
	return _Airdrop.Contract.Symbol(&_Airdrop.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Airdrop *AirdropCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Airdrop *AirdropSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Airdrop.Contract.TokenURI(&_Airdrop.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Airdrop *AirdropCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Airdrop.Contract.TokenURI(&_Airdrop.CallOpts, tokenId)
}

// AddAirdropTokenIds is a paid mutator transaction binding the contract method 0xfdb3b158.
//
// Solidity: function addAirdropTokenIds(uint256[] _tokenIdList) returns()
func (_Airdrop *AirdropTransactor) AddAirdropTokenIds(opts *bind.TransactOpts, _tokenIdList []*big.Int) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "addAirdropTokenIds", _tokenIdList)
}

// AddAirdropTokenIds is a paid mutator transaction binding the contract method 0xfdb3b158.
//
// Solidity: function addAirdropTokenIds(uint256[] _tokenIdList) returns()
func (_Airdrop *AirdropSession) AddAirdropTokenIds(_tokenIdList []*big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.AddAirdropTokenIds(&_Airdrop.TransactOpts, _tokenIdList)
}

// AddAirdropTokenIds is a paid mutator transaction binding the contract method 0xfdb3b158.
//
// Solidity: function addAirdropTokenIds(uint256[] _tokenIdList) returns()
func (_Airdrop *AirdropTransactorSession) AddAirdropTokenIds(_tokenIdList []*big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.AddAirdropTokenIds(&_Airdrop.TransactOpts, _tokenIdList)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Airdrop *AirdropTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Airdrop *AirdropSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.Approve(&_Airdrop.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Airdrop *AirdropTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.Approve(&_Airdrop.TransactOpts, to, tokenId)
}

// GetAirdrop is a paid mutator transaction binding the contract method 0x537a505a.
//
// Solidity: function getAirdrop(bytes32[] _merkleProof) returns()
func (_Airdrop *AirdropTransactor) GetAirdrop(opts *bind.TransactOpts, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "getAirdrop", _merkleProof)
}

// GetAirdrop is a paid mutator transaction binding the contract method 0x537a505a.
//
// Solidity: function getAirdrop(bytes32[] _merkleProof) returns()
func (_Airdrop *AirdropSession) GetAirdrop(_merkleProof [][32]byte) (*types.Transaction, error) {
	return _Airdrop.Contract.GetAirdrop(&_Airdrop.TransactOpts, _merkleProof)
}

// GetAirdrop is a paid mutator transaction binding the contract method 0x537a505a.
//
// Solidity: function getAirdrop(bytes32[] _merkleProof) returns()
func (_Airdrop *AirdropTransactorSession) GetAirdrop(_merkleProof [][32]byte) (*types.Transaction, error) {
	return _Airdrop.Contract.GetAirdrop(&_Airdrop.TransactOpts, _merkleProof)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Airdrop *AirdropTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Airdrop *AirdropSession) RenounceOwnership() (*types.Transaction, error) {
	return _Airdrop.Contract.RenounceOwnership(&_Airdrop.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Airdrop *AirdropTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Airdrop.Contract.RenounceOwnership(&_Airdrop.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Airdrop *AirdropTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Airdrop *AirdropSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.SafeTransferFrom(&_Airdrop.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Airdrop *AirdropTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.SafeTransferFrom(&_Airdrop.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Airdrop *AirdropTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Airdrop *AirdropSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Airdrop.Contract.SafeTransferFrom0(&_Airdrop.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Airdrop *AirdropTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Airdrop.Contract.SafeTransferFrom0(&_Airdrop.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Airdrop *AirdropTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Airdrop *AirdropSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Airdrop.Contract.SetApprovalForAll(&_Airdrop.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Airdrop *AirdropTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Airdrop.Contract.SetApprovalForAll(&_Airdrop.TransactOpts, operator, approved)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x7cb64759.
//
// Solidity: function setMerkleRoot(bytes32 _root) returns()
func (_Airdrop *AirdropTransactor) SetMerkleRoot(opts *bind.TransactOpts, _root [32]byte) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "setMerkleRoot", _root)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x7cb64759.
//
// Solidity: function setMerkleRoot(bytes32 _root) returns()
func (_Airdrop *AirdropSession) SetMerkleRoot(_root [32]byte) (*types.Transaction, error) {
	return _Airdrop.Contract.SetMerkleRoot(&_Airdrop.TransactOpts, _root)
}

// SetMerkleRoot is a paid mutator transaction binding the contract method 0x7cb64759.
//
// Solidity: function setMerkleRoot(bytes32 _root) returns()
func (_Airdrop *AirdropTransactorSession) SetMerkleRoot(_root [32]byte) (*types.Transaction, error) {
	return _Airdrop.Contract.SetMerkleRoot(&_Airdrop.TransactOpts, _root)
}

// SetMintEnable is a paid mutator transaction binding the contract method 0x981fb047.
//
// Solidity: function setMintEnable(bool _enable) returns()
func (_Airdrop *AirdropTransactor) SetMintEnable(opts *bind.TransactOpts, _enable bool) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "setMintEnable", _enable)
}

// SetMintEnable is a paid mutator transaction binding the contract method 0x981fb047.
//
// Solidity: function setMintEnable(bool _enable) returns()
func (_Airdrop *AirdropSession) SetMintEnable(_enable bool) (*types.Transaction, error) {
	return _Airdrop.Contract.SetMintEnable(&_Airdrop.TransactOpts, _enable)
}

// SetMintEnable is a paid mutator transaction binding the contract method 0x981fb047.
//
// Solidity: function setMintEnable(bool _enable) returns()
func (_Airdrop *AirdropTransactorSession) SetMintEnable(_enable bool) (*types.Transaction, error) {
	return _Airdrop.Contract.SetMintEnable(&_Airdrop.TransactOpts, _enable)
}

// SetNftAddress is a paid mutator transaction binding the contract method 0x0b102d1a.
//
// Solidity: function setNftAddress(address _address) returns()
func (_Airdrop *AirdropTransactor) SetNftAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "setNftAddress", _address)
}

// SetNftAddress is a paid mutator transaction binding the contract method 0x0b102d1a.
//
// Solidity: function setNftAddress(address _address) returns()
func (_Airdrop *AirdropSession) SetNftAddress(_address common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.SetNftAddress(&_Airdrop.TransactOpts, _address)
}

// SetNftAddress is a paid mutator transaction binding the contract method 0x0b102d1a.
//
// Solidity: function setNftAddress(address _address) returns()
func (_Airdrop *AirdropTransactorSession) SetNftAddress(_address common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.SetNftAddress(&_Airdrop.TransactOpts, _address)
}

// SetPos is a paid mutator transaction binding the contract method 0x6c37cf7a.
//
// Solidity: function setPos(uint256 _pos) returns()
func (_Airdrop *AirdropTransactor) SetPos(opts *bind.TransactOpts, _pos *big.Int) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "setPos", _pos)
}

// SetPos is a paid mutator transaction binding the contract method 0x6c37cf7a.
//
// Solidity: function setPos(uint256 _pos) returns()
func (_Airdrop *AirdropSession) SetPos(_pos *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.SetPos(&_Airdrop.TransactOpts, _pos)
}

// SetPos is a paid mutator transaction binding the contract method 0x6c37cf7a.
//
// Solidity: function setPos(uint256 _pos) returns()
func (_Airdrop *AirdropTransactorSession) SetPos(_pos *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.SetPos(&_Airdrop.TransactOpts, _pos)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Airdrop *AirdropTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Airdrop *AirdropSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.TransferFrom(&_Airdrop.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Airdrop *AirdropTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.TransferFrom(&_Airdrop.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Airdrop *AirdropTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Airdrop *AirdropSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.TransferOwnership(&_Airdrop.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Airdrop *AirdropTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.TransferOwnership(&_Airdrop.TransactOpts, newOwner)
}

// AirdropApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Airdrop contract.
type AirdropApprovalIterator struct {
	Event *AirdropApproval // Event containing the contract specifics and raw log

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
func (it *AirdropApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropApproval)
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
		it.Event = new(AirdropApproval)
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
func (it *AirdropApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropApproval represents a Approval event raised by the Airdrop contract.
type AirdropApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Airdrop *AirdropFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*AirdropApprovalIterator, error) {

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

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &AirdropApprovalIterator{contract: _Airdrop.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Airdrop *AirdropFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AirdropApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropApproval)
				if err := _Airdrop.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Airdrop *AirdropFilterer) ParseApproval(log types.Log) (*AirdropApproval, error) {
	event := new(AirdropApproval)
	if err := _Airdrop.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Airdrop contract.
type AirdropApprovalForAllIterator struct {
	Event *AirdropApprovalForAll // Event containing the contract specifics and raw log

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
func (it *AirdropApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropApprovalForAll)
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
		it.Event = new(AirdropApprovalForAll)
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
func (it *AirdropApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropApprovalForAll represents a ApprovalForAll event raised by the Airdrop contract.
type AirdropApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Airdrop *AirdropFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*AirdropApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &AirdropApprovalForAllIterator{contract: _Airdrop.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Airdrop *AirdropFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *AirdropApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropApprovalForAll)
				if err := _Airdrop.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Airdrop *AirdropFilterer) ParseApprovalForAll(log types.Log) (*AirdropApprovalForAll, error) {
	event := new(AirdropApprovalForAll)
	if err := _Airdrop.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropClaimBoxIterator is returned from FilterClaimBox and is used to iterate over the raw logs and unpacked data for ClaimBox events raised by the Airdrop contract.
type AirdropClaimBoxIterator struct {
	Event *AirdropClaimBox // Event containing the contract specifics and raw log

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
func (it *AirdropClaimBoxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropClaimBox)
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
		it.Event = new(AirdropClaimBox)
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
func (it *AirdropClaimBoxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropClaimBoxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropClaimBox represents a ClaimBox event raised by the Airdrop contract.
type AirdropClaimBox struct {
	TokenId  *big.Int
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterClaimBox is a free log retrieval operation binding the contract event 0xef3fe58350bbd6cecb2eee87220741c24603a7a5f0ed19286e3e9c1cb7a75df9.
//
// Solidity: event ClaimBox(uint256 _tokenId, address _operator)
func (_Airdrop *AirdropFilterer) FilterClaimBox(opts *bind.FilterOpts) (*AirdropClaimBoxIterator, error) {

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "ClaimBox")
	if err != nil {
		return nil, err
	}
	return &AirdropClaimBoxIterator{contract: _Airdrop.contract, event: "ClaimBox", logs: logs, sub: sub}, nil
}

// WatchClaimBox is a free log subscription operation binding the contract event 0xef3fe58350bbd6cecb2eee87220741c24603a7a5f0ed19286e3e9c1cb7a75df9.
//
// Solidity: event ClaimBox(uint256 _tokenId, address _operator)
func (_Airdrop *AirdropFilterer) WatchClaimBox(opts *bind.WatchOpts, sink chan<- *AirdropClaimBox) (event.Subscription, error) {

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "ClaimBox")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropClaimBox)
				if err := _Airdrop.contract.UnpackLog(event, "ClaimBox", log); err != nil {
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
func (_Airdrop *AirdropFilterer) ParseClaimBox(log types.Log) (*AirdropClaimBox, error) {
	event := new(AirdropClaimBox)
	if err := _Airdrop.contract.UnpackLog(event, "ClaimBox", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Airdrop contract.
type AirdropOwnershipTransferredIterator struct {
	Event *AirdropOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AirdropOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropOwnershipTransferred)
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
		it.Event = new(AirdropOwnershipTransferred)
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
func (it *AirdropOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropOwnershipTransferred represents a OwnershipTransferred event raised by the Airdrop contract.
type AirdropOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Airdrop *AirdropFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AirdropOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AirdropOwnershipTransferredIterator{contract: _Airdrop.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Airdrop *AirdropFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AirdropOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropOwnershipTransferred)
				if err := _Airdrop.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Airdrop *AirdropFilterer) ParseOwnershipTransferred(log types.Log) (*AirdropOwnershipTransferred, error) {
	event := new(AirdropOwnershipTransferred)
	if err := _Airdrop.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AirdropTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Airdrop contract.
type AirdropTransferIterator struct {
	Event *AirdropTransfer // Event containing the contract specifics and raw log

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
func (it *AirdropTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropTransfer)
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
		it.Event = new(AirdropTransfer)
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
func (it *AirdropTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropTransfer represents a Transfer event raised by the Airdrop contract.
type AirdropTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Airdrop *AirdropFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*AirdropTransferIterator, error) {

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

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &AirdropTransferIterator{contract: _Airdrop.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Airdrop *AirdropFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AirdropTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropTransfer)
				if err := _Airdrop.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Airdrop *AirdropFilterer) ParseTransfer(log types.Log) (*AirdropTransfer, error) {
	event := new(AirdropTransfer)
	if err := _Airdrop.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
