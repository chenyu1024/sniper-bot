// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mint

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

// MintMetaData contains all meta data concerning the Mint contract.
var MintMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractRUNBANFT\",\"name\":\"_run\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"FMMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FMPos\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OGMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OGPos\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PPos\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WLMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WLPos\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"_boxList\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"_shoeList\",\"type\":\"string[]\"}],\"name\":\"addFreeMintData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"_boxList\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"_shoeList\",\"type\":\"string[]\"}],\"name\":\"addOgMintData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"_boxList\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"_shoeList\",\"type\":\"string[]\"}],\"name\":\"addPublicMintData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"_boxList\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"_shoeList\",\"type\":\"string[]\"}],\"name\":\"addWlMintData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"freeClaimed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"freeMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"freeMintEnd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"freeMintStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"freeTokenIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"freeTokenIdsBoxUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"freeTokenIdsShoeUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxFreeMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTotalMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"ogClaimed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"ogMint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ogMintEnd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ogMintPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ogMintStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ogTokenIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ogTokenIdsBoxUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ogTokenIdsShoeUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"publicClaimed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publicMint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publicMintEnd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publicMintPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publicMintStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"publicTokenIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"publicTokenIdsBoxUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"publicTokenIdsShoeUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"resetOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"runC\",\"outputs\":[{\"internalType\":\"contractRUNBANFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_fm\",\"type\":\"bytes32\"}],\"name\":\"setFMMerkleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pos\",\"type\":\"uint256\"}],\"name\":\"setFMPos\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_end\",\"type\":\"uint256\"}],\"name\":\"setFreeMintTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxMint\",\"type\":\"uint256\"}],\"name\":\"setMaxFreeMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxTotalMint\",\"type\":\"uint256\"}],\"name\":\"setMaxMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_og\",\"type\":\"bytes32\"}],\"name\":\"setOGMerkleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setOGMintPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_end\",\"type\":\"uint256\"}],\"name\":\"setOGMintTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pos\",\"type\":\"uint256\"}],\"name\":\"setOGPos\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pos\",\"type\":\"uint256\"}],\"name\":\"setPPos\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setPublicMintPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_end\",\"type\":\"uint256\"}],\"name\":\"setPublicMintTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_wl\",\"type\":\"bytes32\"}],\"name\":\"setWLMerkleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setWLMintPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pos\",\"type\":\"uint256\"}],\"name\":\"setWLPos\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_end\",\"type\":\"uint256\"}],\"name\":\"setWhiteListMintTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_new\",\"type\":\"address\"}],\"name\":\"transferToNewOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"whiteListMint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wlClaimed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wlMintEnd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wlMintPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wlMintStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"wlTokenIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"wlTokenIdsBoxUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"wlTokenIdsShoeUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// MintABI is the input ABI used to generate the binding from.
// Deprecated: Use MintMetaData.ABI instead.
var MintABI = MintMetaData.ABI

// Mint is an auto generated Go binding around an Ethereum contract.
type Mint struct {
	MintCaller     // Read-only binding to the contract
	MintTransactor // Write-only binding to the contract
	MintFilterer   // Log filterer for contract events
}

// MintCaller is an auto generated read-only Go binding around an Ethereum contract.
type MintCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MintTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MintFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MintSession struct {
	Contract     *Mint             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MintCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MintCallerSession struct {
	Contract *MintCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MintTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MintTransactorSession struct {
	Contract     *MintTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MintRaw is an auto generated low-level Go binding around an Ethereum contract.
type MintRaw struct {
	Contract *Mint // Generic contract binding to access the raw methods on
}

// MintCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MintCallerRaw struct {
	Contract *MintCaller // Generic read-only contract binding to access the raw methods on
}

// MintTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MintTransactorRaw struct {
	Contract *MintTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMint creates a new instance of Mint, bound to a specific deployed contract.
func NewMint(address common.Address, backend bind.ContractBackend) (*Mint, error) {
	contract, err := bindMint(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mint{MintCaller: MintCaller{contract: contract}, MintTransactor: MintTransactor{contract: contract}, MintFilterer: MintFilterer{contract: contract}}, nil
}

// NewMintCaller creates a new read-only instance of Mint, bound to a specific deployed contract.
func NewMintCaller(address common.Address, caller bind.ContractCaller) (*MintCaller, error) {
	contract, err := bindMint(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MintCaller{contract: contract}, nil
}

// NewMintTransactor creates a new write-only instance of Mint, bound to a specific deployed contract.
func NewMintTransactor(address common.Address, transactor bind.ContractTransactor) (*MintTransactor, error) {
	contract, err := bindMint(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MintTransactor{contract: contract}, nil
}

// NewMintFilterer creates a new log filterer instance of Mint, bound to a specific deployed contract.
func NewMintFilterer(address common.Address, filterer bind.ContractFilterer) (*MintFilterer, error) {
	contract, err := bindMint(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MintFilterer{contract: contract}, nil
}

// bindMint binds a generic wrapper to an already deployed contract.
func bindMint(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MintABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mint *MintRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mint.Contract.MintCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mint *MintRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mint.Contract.MintTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mint *MintRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mint.Contract.MintTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mint *MintCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mint.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mint *MintTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mint.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mint *MintTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mint.Contract.contract.Transact(opts, method, params...)
}

// FMMerkleRoot is a free data retrieval call binding the contract method 0x309999de.
//
// Solidity: function FMMerkleRoot() view returns(bytes32)
func (_Mint *MintCaller) FMMerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "FMMerkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FMMerkleRoot is a free data retrieval call binding the contract method 0x309999de.
//
// Solidity: function FMMerkleRoot() view returns(bytes32)
func (_Mint *MintSession) FMMerkleRoot() ([32]byte, error) {
	return _Mint.Contract.FMMerkleRoot(&_Mint.CallOpts)
}

// FMMerkleRoot is a free data retrieval call binding the contract method 0x309999de.
//
// Solidity: function FMMerkleRoot() view returns(bytes32)
func (_Mint *MintCallerSession) FMMerkleRoot() ([32]byte, error) {
	return _Mint.Contract.FMMerkleRoot(&_Mint.CallOpts)
}

// FMPos is a free data retrieval call binding the contract method 0x3f19f9c6.
//
// Solidity: function FMPos() view returns(uint256)
func (_Mint *MintCaller) FMPos(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "FMPos")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FMPos is a free data retrieval call binding the contract method 0x3f19f9c6.
//
// Solidity: function FMPos() view returns(uint256)
func (_Mint *MintSession) FMPos() (*big.Int, error) {
	return _Mint.Contract.FMPos(&_Mint.CallOpts)
}

// FMPos is a free data retrieval call binding the contract method 0x3f19f9c6.
//
// Solidity: function FMPos() view returns(uint256)
func (_Mint *MintCallerSession) FMPos() (*big.Int, error) {
	return _Mint.Contract.FMPos(&_Mint.CallOpts)
}

// OGMerkleRoot is a free data retrieval call binding the contract method 0x83df8d8d.
//
// Solidity: function OGMerkleRoot() view returns(bytes32)
func (_Mint *MintCaller) OGMerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "OGMerkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OGMerkleRoot is a free data retrieval call binding the contract method 0x83df8d8d.
//
// Solidity: function OGMerkleRoot() view returns(bytes32)
func (_Mint *MintSession) OGMerkleRoot() ([32]byte, error) {
	return _Mint.Contract.OGMerkleRoot(&_Mint.CallOpts)
}

// OGMerkleRoot is a free data retrieval call binding the contract method 0x83df8d8d.
//
// Solidity: function OGMerkleRoot() view returns(bytes32)
func (_Mint *MintCallerSession) OGMerkleRoot() ([32]byte, error) {
	return _Mint.Contract.OGMerkleRoot(&_Mint.CallOpts)
}

// OGPos is a free data retrieval call binding the contract method 0xe9f7018b.
//
// Solidity: function OGPos() view returns(uint256)
func (_Mint *MintCaller) OGPos(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "OGPos")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OGPos is a free data retrieval call binding the contract method 0xe9f7018b.
//
// Solidity: function OGPos() view returns(uint256)
func (_Mint *MintSession) OGPos() (*big.Int, error) {
	return _Mint.Contract.OGPos(&_Mint.CallOpts)
}

// OGPos is a free data retrieval call binding the contract method 0xe9f7018b.
//
// Solidity: function OGPos() view returns(uint256)
func (_Mint *MintCallerSession) OGPos() (*big.Int, error) {
	return _Mint.Contract.OGPos(&_Mint.CallOpts)
}

// PPos is a free data retrieval call binding the contract method 0x87d593c7.
//
// Solidity: function PPos() view returns(uint256)
func (_Mint *MintCaller) PPos(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "PPos")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PPos is a free data retrieval call binding the contract method 0x87d593c7.
//
// Solidity: function PPos() view returns(uint256)
func (_Mint *MintSession) PPos() (*big.Int, error) {
	return _Mint.Contract.PPos(&_Mint.CallOpts)
}

// PPos is a free data retrieval call binding the contract method 0x87d593c7.
//
// Solidity: function PPos() view returns(uint256)
func (_Mint *MintCallerSession) PPos() (*big.Int, error) {
	return _Mint.Contract.PPos(&_Mint.CallOpts)
}

// WLMerkleRoot is a free data retrieval call binding the contract method 0xaea48328.
//
// Solidity: function WLMerkleRoot() view returns(bytes32)
func (_Mint *MintCaller) WLMerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "WLMerkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WLMerkleRoot is a free data retrieval call binding the contract method 0xaea48328.
//
// Solidity: function WLMerkleRoot() view returns(bytes32)
func (_Mint *MintSession) WLMerkleRoot() ([32]byte, error) {
	return _Mint.Contract.WLMerkleRoot(&_Mint.CallOpts)
}

// WLMerkleRoot is a free data retrieval call binding the contract method 0xaea48328.
//
// Solidity: function WLMerkleRoot() view returns(bytes32)
func (_Mint *MintCallerSession) WLMerkleRoot() ([32]byte, error) {
	return _Mint.Contract.WLMerkleRoot(&_Mint.CallOpts)
}

// WLPos is a free data retrieval call binding the contract method 0x260c9a29.
//
// Solidity: function WLPos() view returns(uint256)
func (_Mint *MintCaller) WLPos(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "WLPos")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WLPos is a free data retrieval call binding the contract method 0x260c9a29.
//
// Solidity: function WLPos() view returns(uint256)
func (_Mint *MintSession) WLPos() (*big.Int, error) {
	return _Mint.Contract.WLPos(&_Mint.CallOpts)
}

// WLPos is a free data retrieval call binding the contract method 0x260c9a29.
//
// Solidity: function WLPos() view returns(uint256)
func (_Mint *MintCallerSession) WLPos() (*big.Int, error) {
	return _Mint.Contract.WLPos(&_Mint.CallOpts)
}

// FreeClaimed is a free data retrieval call binding the contract method 0x61c0b6a0.
//
// Solidity: function freeClaimed(address ) view returns(uint256)
func (_Mint *MintCaller) FreeClaimed(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "freeClaimed", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FreeClaimed is a free data retrieval call binding the contract method 0x61c0b6a0.
//
// Solidity: function freeClaimed(address ) view returns(uint256)
func (_Mint *MintSession) FreeClaimed(arg0 common.Address) (*big.Int, error) {
	return _Mint.Contract.FreeClaimed(&_Mint.CallOpts, arg0)
}

// FreeClaimed is a free data retrieval call binding the contract method 0x61c0b6a0.
//
// Solidity: function freeClaimed(address ) view returns(uint256)
func (_Mint *MintCallerSession) FreeClaimed(arg0 common.Address) (*big.Int, error) {
	return _Mint.Contract.FreeClaimed(&_Mint.CallOpts, arg0)
}

// FreeMintEnd is a free data retrieval call binding the contract method 0x44fdd445.
//
// Solidity: function freeMintEnd() view returns(uint256)
func (_Mint *MintCaller) FreeMintEnd(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "freeMintEnd")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FreeMintEnd is a free data retrieval call binding the contract method 0x44fdd445.
//
// Solidity: function freeMintEnd() view returns(uint256)
func (_Mint *MintSession) FreeMintEnd() (*big.Int, error) {
	return _Mint.Contract.FreeMintEnd(&_Mint.CallOpts)
}

// FreeMintEnd is a free data retrieval call binding the contract method 0x44fdd445.
//
// Solidity: function freeMintEnd() view returns(uint256)
func (_Mint *MintCallerSession) FreeMintEnd() (*big.Int, error) {
	return _Mint.Contract.FreeMintEnd(&_Mint.CallOpts)
}

// FreeMintStart is a free data retrieval call binding the contract method 0x70706ab3.
//
// Solidity: function freeMintStart() view returns(uint256)
func (_Mint *MintCaller) FreeMintStart(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "freeMintStart")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FreeMintStart is a free data retrieval call binding the contract method 0x70706ab3.
//
// Solidity: function freeMintStart() view returns(uint256)
func (_Mint *MintSession) FreeMintStart() (*big.Int, error) {
	return _Mint.Contract.FreeMintStart(&_Mint.CallOpts)
}

// FreeMintStart is a free data retrieval call binding the contract method 0x70706ab3.
//
// Solidity: function freeMintStart() view returns(uint256)
func (_Mint *MintCallerSession) FreeMintStart() (*big.Int, error) {
	return _Mint.Contract.FreeMintStart(&_Mint.CallOpts)
}

// FreeTokenIds is a free data retrieval call binding the contract method 0x2f827403.
//
// Solidity: function freeTokenIds(uint256 ) view returns(uint256)
func (_Mint *MintCaller) FreeTokenIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "freeTokenIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FreeTokenIds is a free data retrieval call binding the contract method 0x2f827403.
//
// Solidity: function freeTokenIds(uint256 ) view returns(uint256)
func (_Mint *MintSession) FreeTokenIds(arg0 *big.Int) (*big.Int, error) {
	return _Mint.Contract.FreeTokenIds(&_Mint.CallOpts, arg0)
}

// FreeTokenIds is a free data retrieval call binding the contract method 0x2f827403.
//
// Solidity: function freeTokenIds(uint256 ) view returns(uint256)
func (_Mint *MintCallerSession) FreeTokenIds(arg0 *big.Int) (*big.Int, error) {
	return _Mint.Contract.FreeTokenIds(&_Mint.CallOpts, arg0)
}

// FreeTokenIdsBoxUri is a free data retrieval call binding the contract method 0x5793a713.
//
// Solidity: function freeTokenIdsBoxUri(uint256 ) view returns(string)
func (_Mint *MintCaller) FreeTokenIdsBoxUri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "freeTokenIdsBoxUri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// FreeTokenIdsBoxUri is a free data retrieval call binding the contract method 0x5793a713.
//
// Solidity: function freeTokenIdsBoxUri(uint256 ) view returns(string)
func (_Mint *MintSession) FreeTokenIdsBoxUri(arg0 *big.Int) (string, error) {
	return _Mint.Contract.FreeTokenIdsBoxUri(&_Mint.CallOpts, arg0)
}

// FreeTokenIdsBoxUri is a free data retrieval call binding the contract method 0x5793a713.
//
// Solidity: function freeTokenIdsBoxUri(uint256 ) view returns(string)
func (_Mint *MintCallerSession) FreeTokenIdsBoxUri(arg0 *big.Int) (string, error) {
	return _Mint.Contract.FreeTokenIdsBoxUri(&_Mint.CallOpts, arg0)
}

// FreeTokenIdsShoeUri is a free data retrieval call binding the contract method 0x0ead5b2c.
//
// Solidity: function freeTokenIdsShoeUri(uint256 ) view returns(string)
func (_Mint *MintCaller) FreeTokenIdsShoeUri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "freeTokenIdsShoeUri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// FreeTokenIdsShoeUri is a free data retrieval call binding the contract method 0x0ead5b2c.
//
// Solidity: function freeTokenIdsShoeUri(uint256 ) view returns(string)
func (_Mint *MintSession) FreeTokenIdsShoeUri(arg0 *big.Int) (string, error) {
	return _Mint.Contract.FreeTokenIdsShoeUri(&_Mint.CallOpts, arg0)
}

// FreeTokenIdsShoeUri is a free data retrieval call binding the contract method 0x0ead5b2c.
//
// Solidity: function freeTokenIdsShoeUri(uint256 ) view returns(string)
func (_Mint *MintCallerSession) FreeTokenIdsShoeUri(arg0 *big.Int) (string, error) {
	return _Mint.Contract.FreeTokenIdsShoeUri(&_Mint.CallOpts, arg0)
}

// MaxFreeMint is a free data retrieval call binding the contract method 0xa591252d.
//
// Solidity: function maxFreeMint() view returns(uint256)
func (_Mint *MintCaller) MaxFreeMint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "maxFreeMint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxFreeMint is a free data retrieval call binding the contract method 0xa591252d.
//
// Solidity: function maxFreeMint() view returns(uint256)
func (_Mint *MintSession) MaxFreeMint() (*big.Int, error) {
	return _Mint.Contract.MaxFreeMint(&_Mint.CallOpts)
}

// MaxFreeMint is a free data retrieval call binding the contract method 0xa591252d.
//
// Solidity: function maxFreeMint() view returns(uint256)
func (_Mint *MintCallerSession) MaxFreeMint() (*big.Int, error) {
	return _Mint.Contract.MaxFreeMint(&_Mint.CallOpts)
}

// MaxTotalMint is a free data retrieval call binding the contract method 0xf4170e98.
//
// Solidity: function maxTotalMint() view returns(uint256)
func (_Mint *MintCaller) MaxTotalMint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "maxTotalMint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTotalMint is a free data retrieval call binding the contract method 0xf4170e98.
//
// Solidity: function maxTotalMint() view returns(uint256)
func (_Mint *MintSession) MaxTotalMint() (*big.Int, error) {
	return _Mint.Contract.MaxTotalMint(&_Mint.CallOpts)
}

// MaxTotalMint is a free data retrieval call binding the contract method 0xf4170e98.
//
// Solidity: function maxTotalMint() view returns(uint256)
func (_Mint *MintCallerSession) MaxTotalMint() (*big.Int, error) {
	return _Mint.Contract.MaxTotalMint(&_Mint.CallOpts)
}

// OgClaimed is a free data retrieval call binding the contract method 0x73fc16ad.
//
// Solidity: function ogClaimed(address ) view returns(uint256)
func (_Mint *MintCaller) OgClaimed(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "ogClaimed", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OgClaimed is a free data retrieval call binding the contract method 0x73fc16ad.
//
// Solidity: function ogClaimed(address ) view returns(uint256)
func (_Mint *MintSession) OgClaimed(arg0 common.Address) (*big.Int, error) {
	return _Mint.Contract.OgClaimed(&_Mint.CallOpts, arg0)
}

// OgClaimed is a free data retrieval call binding the contract method 0x73fc16ad.
//
// Solidity: function ogClaimed(address ) view returns(uint256)
func (_Mint *MintCallerSession) OgClaimed(arg0 common.Address) (*big.Int, error) {
	return _Mint.Contract.OgClaimed(&_Mint.CallOpts, arg0)
}

// OgMintEnd is a free data retrieval call binding the contract method 0xf5facba3.
//
// Solidity: function ogMintEnd() view returns(uint256)
func (_Mint *MintCaller) OgMintEnd(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "ogMintEnd")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OgMintEnd is a free data retrieval call binding the contract method 0xf5facba3.
//
// Solidity: function ogMintEnd() view returns(uint256)
func (_Mint *MintSession) OgMintEnd() (*big.Int, error) {
	return _Mint.Contract.OgMintEnd(&_Mint.CallOpts)
}

// OgMintEnd is a free data retrieval call binding the contract method 0xf5facba3.
//
// Solidity: function ogMintEnd() view returns(uint256)
func (_Mint *MintCallerSession) OgMintEnd() (*big.Int, error) {
	return _Mint.Contract.OgMintEnd(&_Mint.CallOpts)
}

// OgMintPrice is a free data retrieval call binding the contract method 0x9d1f1c47.
//
// Solidity: function ogMintPrice() view returns(uint256)
func (_Mint *MintCaller) OgMintPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "ogMintPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OgMintPrice is a free data retrieval call binding the contract method 0x9d1f1c47.
//
// Solidity: function ogMintPrice() view returns(uint256)
func (_Mint *MintSession) OgMintPrice() (*big.Int, error) {
	return _Mint.Contract.OgMintPrice(&_Mint.CallOpts)
}

// OgMintPrice is a free data retrieval call binding the contract method 0x9d1f1c47.
//
// Solidity: function ogMintPrice() view returns(uint256)
func (_Mint *MintCallerSession) OgMintPrice() (*big.Int, error) {
	return _Mint.Contract.OgMintPrice(&_Mint.CallOpts)
}

// OgMintStart is a free data retrieval call binding the contract method 0x627711c3.
//
// Solidity: function ogMintStart() view returns(uint256)
func (_Mint *MintCaller) OgMintStart(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "ogMintStart")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OgMintStart is a free data retrieval call binding the contract method 0x627711c3.
//
// Solidity: function ogMintStart() view returns(uint256)
func (_Mint *MintSession) OgMintStart() (*big.Int, error) {
	return _Mint.Contract.OgMintStart(&_Mint.CallOpts)
}

// OgMintStart is a free data retrieval call binding the contract method 0x627711c3.
//
// Solidity: function ogMintStart() view returns(uint256)
func (_Mint *MintCallerSession) OgMintStart() (*big.Int, error) {
	return _Mint.Contract.OgMintStart(&_Mint.CallOpts)
}

// OgTokenIds is a free data retrieval call binding the contract method 0x72bcdaf1.
//
// Solidity: function ogTokenIds(uint256 ) view returns(uint256)
func (_Mint *MintCaller) OgTokenIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "ogTokenIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OgTokenIds is a free data retrieval call binding the contract method 0x72bcdaf1.
//
// Solidity: function ogTokenIds(uint256 ) view returns(uint256)
func (_Mint *MintSession) OgTokenIds(arg0 *big.Int) (*big.Int, error) {
	return _Mint.Contract.OgTokenIds(&_Mint.CallOpts, arg0)
}

// OgTokenIds is a free data retrieval call binding the contract method 0x72bcdaf1.
//
// Solidity: function ogTokenIds(uint256 ) view returns(uint256)
func (_Mint *MintCallerSession) OgTokenIds(arg0 *big.Int) (*big.Int, error) {
	return _Mint.Contract.OgTokenIds(&_Mint.CallOpts, arg0)
}

// OgTokenIdsBoxUri is a free data retrieval call binding the contract method 0x35ccdb01.
//
// Solidity: function ogTokenIdsBoxUri(uint256 ) view returns(string)
func (_Mint *MintCaller) OgTokenIdsBoxUri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "ogTokenIdsBoxUri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// OgTokenIdsBoxUri is a free data retrieval call binding the contract method 0x35ccdb01.
//
// Solidity: function ogTokenIdsBoxUri(uint256 ) view returns(string)
func (_Mint *MintSession) OgTokenIdsBoxUri(arg0 *big.Int) (string, error) {
	return _Mint.Contract.OgTokenIdsBoxUri(&_Mint.CallOpts, arg0)
}

// OgTokenIdsBoxUri is a free data retrieval call binding the contract method 0x35ccdb01.
//
// Solidity: function ogTokenIdsBoxUri(uint256 ) view returns(string)
func (_Mint *MintCallerSession) OgTokenIdsBoxUri(arg0 *big.Int) (string, error) {
	return _Mint.Contract.OgTokenIdsBoxUri(&_Mint.CallOpts, arg0)
}

// OgTokenIdsShoeUri is a free data retrieval call binding the contract method 0xd0964033.
//
// Solidity: function ogTokenIdsShoeUri(uint256 ) view returns(string)
func (_Mint *MintCaller) OgTokenIdsShoeUri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "ogTokenIdsShoeUri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// OgTokenIdsShoeUri is a free data retrieval call binding the contract method 0xd0964033.
//
// Solidity: function ogTokenIdsShoeUri(uint256 ) view returns(string)
func (_Mint *MintSession) OgTokenIdsShoeUri(arg0 *big.Int) (string, error) {
	return _Mint.Contract.OgTokenIdsShoeUri(&_Mint.CallOpts, arg0)
}

// OgTokenIdsShoeUri is a free data retrieval call binding the contract method 0xd0964033.
//
// Solidity: function ogTokenIdsShoeUri(uint256 ) view returns(string)
func (_Mint *MintCallerSession) OgTokenIdsShoeUri(arg0 *big.Int) (string, error) {
	return _Mint.Contract.OgTokenIdsShoeUri(&_Mint.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mint *MintCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mint *MintSession) Owner() (common.Address, error) {
	return _Mint.Contract.Owner(&_Mint.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mint *MintCallerSession) Owner() (common.Address, error) {
	return _Mint.Contract.Owner(&_Mint.CallOpts)
}

// PublicClaimed is a free data retrieval call binding the contract method 0xb5b1cd7c.
//
// Solidity: function publicClaimed(address ) view returns(uint256)
func (_Mint *MintCaller) PublicClaimed(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "publicClaimed", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PublicClaimed is a free data retrieval call binding the contract method 0xb5b1cd7c.
//
// Solidity: function publicClaimed(address ) view returns(uint256)
func (_Mint *MintSession) PublicClaimed(arg0 common.Address) (*big.Int, error) {
	return _Mint.Contract.PublicClaimed(&_Mint.CallOpts, arg0)
}

// PublicClaimed is a free data retrieval call binding the contract method 0xb5b1cd7c.
//
// Solidity: function publicClaimed(address ) view returns(uint256)
func (_Mint *MintCallerSession) PublicClaimed(arg0 common.Address) (*big.Int, error) {
	return _Mint.Contract.PublicClaimed(&_Mint.CallOpts, arg0)
}

// PublicMintEnd is a free data retrieval call binding the contract method 0x2a196e1b.
//
// Solidity: function publicMintEnd() view returns(uint256)
func (_Mint *MintCaller) PublicMintEnd(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "publicMintEnd")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PublicMintEnd is a free data retrieval call binding the contract method 0x2a196e1b.
//
// Solidity: function publicMintEnd() view returns(uint256)
func (_Mint *MintSession) PublicMintEnd() (*big.Int, error) {
	return _Mint.Contract.PublicMintEnd(&_Mint.CallOpts)
}

// PublicMintEnd is a free data retrieval call binding the contract method 0x2a196e1b.
//
// Solidity: function publicMintEnd() view returns(uint256)
func (_Mint *MintCallerSession) PublicMintEnd() (*big.Int, error) {
	return _Mint.Contract.PublicMintEnd(&_Mint.CallOpts)
}

// PublicMintPrice is a free data retrieval call binding the contract method 0xdc53fd92.
//
// Solidity: function publicMintPrice() view returns(uint256)
func (_Mint *MintCaller) PublicMintPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "publicMintPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PublicMintPrice is a free data retrieval call binding the contract method 0xdc53fd92.
//
// Solidity: function publicMintPrice() view returns(uint256)
func (_Mint *MintSession) PublicMintPrice() (*big.Int, error) {
	return _Mint.Contract.PublicMintPrice(&_Mint.CallOpts)
}

// PublicMintPrice is a free data retrieval call binding the contract method 0xdc53fd92.
//
// Solidity: function publicMintPrice() view returns(uint256)
func (_Mint *MintCallerSession) PublicMintPrice() (*big.Int, error) {
	return _Mint.Contract.PublicMintPrice(&_Mint.CallOpts)
}

// PublicMintStart is a free data retrieval call binding the contract method 0x8cfec4c0.
//
// Solidity: function publicMintStart() view returns(uint256)
func (_Mint *MintCaller) PublicMintStart(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "publicMintStart")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PublicMintStart is a free data retrieval call binding the contract method 0x8cfec4c0.
//
// Solidity: function publicMintStart() view returns(uint256)
func (_Mint *MintSession) PublicMintStart() (*big.Int, error) {
	return _Mint.Contract.PublicMintStart(&_Mint.CallOpts)
}

// PublicMintStart is a free data retrieval call binding the contract method 0x8cfec4c0.
//
// Solidity: function publicMintStart() view returns(uint256)
func (_Mint *MintCallerSession) PublicMintStart() (*big.Int, error) {
	return _Mint.Contract.PublicMintStart(&_Mint.CallOpts)
}

// PublicTokenIds is a free data retrieval call binding the contract method 0xf5653f4e.
//
// Solidity: function publicTokenIds(uint256 ) view returns(uint256)
func (_Mint *MintCaller) PublicTokenIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "publicTokenIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PublicTokenIds is a free data retrieval call binding the contract method 0xf5653f4e.
//
// Solidity: function publicTokenIds(uint256 ) view returns(uint256)
func (_Mint *MintSession) PublicTokenIds(arg0 *big.Int) (*big.Int, error) {
	return _Mint.Contract.PublicTokenIds(&_Mint.CallOpts, arg0)
}

// PublicTokenIds is a free data retrieval call binding the contract method 0xf5653f4e.
//
// Solidity: function publicTokenIds(uint256 ) view returns(uint256)
func (_Mint *MintCallerSession) PublicTokenIds(arg0 *big.Int) (*big.Int, error) {
	return _Mint.Contract.PublicTokenIds(&_Mint.CallOpts, arg0)
}

// PublicTokenIdsBoxUri is a free data retrieval call binding the contract method 0xc2bcab10.
//
// Solidity: function publicTokenIdsBoxUri(uint256 ) view returns(string)
func (_Mint *MintCaller) PublicTokenIdsBoxUri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "publicTokenIdsBoxUri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PublicTokenIdsBoxUri is a free data retrieval call binding the contract method 0xc2bcab10.
//
// Solidity: function publicTokenIdsBoxUri(uint256 ) view returns(string)
func (_Mint *MintSession) PublicTokenIdsBoxUri(arg0 *big.Int) (string, error) {
	return _Mint.Contract.PublicTokenIdsBoxUri(&_Mint.CallOpts, arg0)
}

// PublicTokenIdsBoxUri is a free data retrieval call binding the contract method 0xc2bcab10.
//
// Solidity: function publicTokenIdsBoxUri(uint256 ) view returns(string)
func (_Mint *MintCallerSession) PublicTokenIdsBoxUri(arg0 *big.Int) (string, error) {
	return _Mint.Contract.PublicTokenIdsBoxUri(&_Mint.CallOpts, arg0)
}

// PublicTokenIdsShoeUri is a free data retrieval call binding the contract method 0xb60f78eb.
//
// Solidity: function publicTokenIdsShoeUri(uint256 ) view returns(string)
func (_Mint *MintCaller) PublicTokenIdsShoeUri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "publicTokenIdsShoeUri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PublicTokenIdsShoeUri is a free data retrieval call binding the contract method 0xb60f78eb.
//
// Solidity: function publicTokenIdsShoeUri(uint256 ) view returns(string)
func (_Mint *MintSession) PublicTokenIdsShoeUri(arg0 *big.Int) (string, error) {
	return _Mint.Contract.PublicTokenIdsShoeUri(&_Mint.CallOpts, arg0)
}

// PublicTokenIdsShoeUri is a free data retrieval call binding the contract method 0xb60f78eb.
//
// Solidity: function publicTokenIdsShoeUri(uint256 ) view returns(string)
func (_Mint *MintCallerSession) PublicTokenIdsShoeUri(arg0 *big.Int) (string, error) {
	return _Mint.Contract.PublicTokenIdsShoeUri(&_Mint.CallOpts, arg0)
}

// RunC is a free data retrieval call binding the contract method 0xa664d160.
//
// Solidity: function runC() view returns(address)
func (_Mint *MintCaller) RunC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "runC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RunC is a free data retrieval call binding the contract method 0xa664d160.
//
// Solidity: function runC() view returns(address)
func (_Mint *MintSession) RunC() (common.Address, error) {
	return _Mint.Contract.RunC(&_Mint.CallOpts)
}

// RunC is a free data retrieval call binding the contract method 0xa664d160.
//
// Solidity: function runC() view returns(address)
func (_Mint *MintCallerSession) RunC() (common.Address, error) {
	return _Mint.Contract.RunC(&_Mint.CallOpts)
}

// WlClaimed is a free data retrieval call binding the contract method 0x5a089b30.
//
// Solidity: function wlClaimed(address ) view returns(uint256)
func (_Mint *MintCaller) WlClaimed(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "wlClaimed", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WlClaimed is a free data retrieval call binding the contract method 0x5a089b30.
//
// Solidity: function wlClaimed(address ) view returns(uint256)
func (_Mint *MintSession) WlClaimed(arg0 common.Address) (*big.Int, error) {
	return _Mint.Contract.WlClaimed(&_Mint.CallOpts, arg0)
}

// WlClaimed is a free data retrieval call binding the contract method 0x5a089b30.
//
// Solidity: function wlClaimed(address ) view returns(uint256)
func (_Mint *MintCallerSession) WlClaimed(arg0 common.Address) (*big.Int, error) {
	return _Mint.Contract.WlClaimed(&_Mint.CallOpts, arg0)
}

// WlMintEnd is a free data retrieval call binding the contract method 0x62a66ea9.
//
// Solidity: function wlMintEnd() view returns(uint256)
func (_Mint *MintCaller) WlMintEnd(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "wlMintEnd")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WlMintEnd is a free data retrieval call binding the contract method 0x62a66ea9.
//
// Solidity: function wlMintEnd() view returns(uint256)
func (_Mint *MintSession) WlMintEnd() (*big.Int, error) {
	return _Mint.Contract.WlMintEnd(&_Mint.CallOpts)
}

// WlMintEnd is a free data retrieval call binding the contract method 0x62a66ea9.
//
// Solidity: function wlMintEnd() view returns(uint256)
func (_Mint *MintCallerSession) WlMintEnd() (*big.Int, error) {
	return _Mint.Contract.WlMintEnd(&_Mint.CallOpts)
}

// WlMintPrice is a free data retrieval call binding the contract method 0x2c4e9fc6.
//
// Solidity: function wlMintPrice() view returns(uint256)
func (_Mint *MintCaller) WlMintPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "wlMintPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WlMintPrice is a free data retrieval call binding the contract method 0x2c4e9fc6.
//
// Solidity: function wlMintPrice() view returns(uint256)
func (_Mint *MintSession) WlMintPrice() (*big.Int, error) {
	return _Mint.Contract.WlMintPrice(&_Mint.CallOpts)
}

// WlMintPrice is a free data retrieval call binding the contract method 0x2c4e9fc6.
//
// Solidity: function wlMintPrice() view returns(uint256)
func (_Mint *MintCallerSession) WlMintPrice() (*big.Int, error) {
	return _Mint.Contract.WlMintPrice(&_Mint.CallOpts)
}

// WlMintStart is a free data retrieval call binding the contract method 0x9f8532c1.
//
// Solidity: function wlMintStart() view returns(uint256)
func (_Mint *MintCaller) WlMintStart(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "wlMintStart")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WlMintStart is a free data retrieval call binding the contract method 0x9f8532c1.
//
// Solidity: function wlMintStart() view returns(uint256)
func (_Mint *MintSession) WlMintStart() (*big.Int, error) {
	return _Mint.Contract.WlMintStart(&_Mint.CallOpts)
}

// WlMintStart is a free data retrieval call binding the contract method 0x9f8532c1.
//
// Solidity: function wlMintStart() view returns(uint256)
func (_Mint *MintCallerSession) WlMintStart() (*big.Int, error) {
	return _Mint.Contract.WlMintStart(&_Mint.CallOpts)
}

// WlTokenIds is a free data retrieval call binding the contract method 0x76e16d28.
//
// Solidity: function wlTokenIds(uint256 ) view returns(uint256)
func (_Mint *MintCaller) WlTokenIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "wlTokenIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WlTokenIds is a free data retrieval call binding the contract method 0x76e16d28.
//
// Solidity: function wlTokenIds(uint256 ) view returns(uint256)
func (_Mint *MintSession) WlTokenIds(arg0 *big.Int) (*big.Int, error) {
	return _Mint.Contract.WlTokenIds(&_Mint.CallOpts, arg0)
}

// WlTokenIds is a free data retrieval call binding the contract method 0x76e16d28.
//
// Solidity: function wlTokenIds(uint256 ) view returns(uint256)
func (_Mint *MintCallerSession) WlTokenIds(arg0 *big.Int) (*big.Int, error) {
	return _Mint.Contract.WlTokenIds(&_Mint.CallOpts, arg0)
}

// WlTokenIdsBoxUri is a free data retrieval call binding the contract method 0x84eaa704.
//
// Solidity: function wlTokenIdsBoxUri(uint256 ) view returns(string)
func (_Mint *MintCaller) WlTokenIdsBoxUri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "wlTokenIdsBoxUri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// WlTokenIdsBoxUri is a free data retrieval call binding the contract method 0x84eaa704.
//
// Solidity: function wlTokenIdsBoxUri(uint256 ) view returns(string)
func (_Mint *MintSession) WlTokenIdsBoxUri(arg0 *big.Int) (string, error) {
	return _Mint.Contract.WlTokenIdsBoxUri(&_Mint.CallOpts, arg0)
}

// WlTokenIdsBoxUri is a free data retrieval call binding the contract method 0x84eaa704.
//
// Solidity: function wlTokenIdsBoxUri(uint256 ) view returns(string)
func (_Mint *MintCallerSession) WlTokenIdsBoxUri(arg0 *big.Int) (string, error) {
	return _Mint.Contract.WlTokenIdsBoxUri(&_Mint.CallOpts, arg0)
}

// WlTokenIdsShoeUri is a free data retrieval call binding the contract method 0x58bbd44f.
//
// Solidity: function wlTokenIdsShoeUri(uint256 ) view returns(string)
func (_Mint *MintCaller) WlTokenIdsShoeUri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "wlTokenIdsShoeUri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// WlTokenIdsShoeUri is a free data retrieval call binding the contract method 0x58bbd44f.
//
// Solidity: function wlTokenIdsShoeUri(uint256 ) view returns(string)
func (_Mint *MintSession) WlTokenIdsShoeUri(arg0 *big.Int) (string, error) {
	return _Mint.Contract.WlTokenIdsShoeUri(&_Mint.CallOpts, arg0)
}

// WlTokenIdsShoeUri is a free data retrieval call binding the contract method 0x58bbd44f.
//
// Solidity: function wlTokenIdsShoeUri(uint256 ) view returns(string)
func (_Mint *MintCallerSession) WlTokenIdsShoeUri(arg0 *big.Int) (string, error) {
	return _Mint.Contract.WlTokenIdsShoeUri(&_Mint.CallOpts, arg0)
}

// AddFreeMintData is a paid mutator transaction binding the contract method 0x3ce375e0.
//
// Solidity: function addFreeMintData(uint256[] _tokenIds, string[] _boxList, string[] _shoeList) returns()
func (_Mint *MintTransactor) AddFreeMintData(opts *bind.TransactOpts, _tokenIds []*big.Int, _boxList []string, _shoeList []string) (*types.Transaction, error) {
	return _Mint.contract.Transact(opts, "addFreeMintData", _tokenIds, _boxList, _shoeList)
}

// AddFreeMintData is a paid mutator transaction binding the contract method 0x3ce375e0.
//
// Solidity: function addFreeMintData(uint256[] _tokenIds, string[] _boxList, string[] _shoeList) returns()
func (_Mint *MintSession) AddFreeMintData(_tokenIds []*big.Int, _boxList []string, _shoeList []string) (*types.Transaction, error) {
	return _Mint.Contract.AddFreeMintData(&_Mint.TransactOpts, _tokenIds, _boxList, _shoeList)
}

// AddFreeMintData is a paid mutator transaction binding the contract method 0x3ce375e0.
//
// Solidity: function addFreeMintData(uint256[] _tokenIds, string[] _boxList, string[] _shoeList) returns()
func (_Mint *MintTransactorSession) AddFreeMintData(_tokenIds []*big.Int, _boxList []string, _shoeList []string) (*types.Transaction, error) {
	return _Mint.Contract.AddFreeMintData(&_Mint.TransactOpts, _tokenIds, _boxList, _shoeList)
}

// AddOgMintData is a paid mutator transaction binding the contract method 0xdcfcf4b6.
//
// Solidity: function addOgMintData(uint256[] _tokenIds, string[] _boxList, string[] _shoeList) returns()
func (_Mint *MintTransactor) AddOgMintData(opts *bind.TransactOpts, _tokenIds []*big.Int, _boxList []string, _shoeList []string) (*types.Transaction, error) {
	return _Mint.contract.Transact(opts, "addOgMintData", _tokenIds, _boxList, _shoeList)
}

// AddOgMintData is a paid mutator transaction binding the contract method 0xdcfcf4b6.
//
// Solidity: function addOgMintData(uint256[] _tokenIds, string[] _boxList, string[] _shoeList) returns()
func (_Mint *MintSession) AddOgMintData(_tokenIds []*big.Int, _boxList []string, _shoeList []string) (*types.Transaction, error) {
	return _Mint.Contract.AddOgMintData(&_Mint.TransactOpts, _tokenIds, _boxList, _shoeList)
}

// AddOgMintData is a paid mutator transaction binding the contract method 0xdcfcf4b6.
//
// Solidity: function addOgMintData(uint256[] _tokenIds, string[] _boxList, string[] _shoeList) returns()
func (_Mint *MintTransactorSession) AddOgMintData(_tokenIds []*big.Int, _boxList []string, _shoeList []string) (*types.Transaction, error) {
	return _Mint.Contract.AddOgMintData(&_Mint.TransactOpts, _tokenIds, _boxList, _shoeList)
}

// AddPublicMintData is a paid mutator transaction binding the contract method 0xebfcc606.
//
// Solidity: function addPublicMintData(uint256[] _tokenIds, string[] _boxList, string[] _shoeList) returns()
func (_Mint *MintTransactor) AddPublicMintData(opts *bind.TransactOpts, _tokenIds []*big.Int, _boxList []string, _shoeList []string) (*types.Transaction, error) {
	return _Mint.contract.Transact(opts, "addPublicMintData", _tokenIds, _boxList, _shoeList)
}

// AddPublicMintData is a paid mutator transaction binding the contract method 0xebfcc606.
//
// Solidity: function addPublicMintData(uint256[] _tokenIds, string[] _boxList, string[] _shoeList) returns()
func (_Mint *MintSession) AddPublicMintData(_tokenIds []*big.Int, _boxList []string, _shoeList []string) (*types.Transaction, error) {
	return _Mint.Contract.AddPublicMintData(&_Mint.TransactOpts, _tokenIds, _boxList, _shoeList)
}

// AddPublicMintData is a paid mutator transaction binding the contract method 0xebfcc606.
//
// Solidity: function addPublicMintData(uint256[] _tokenIds, string[] _boxList, string[] _shoeList) returns()
func (_Mint *MintTransactorSession) AddPublicMintData(_tokenIds []*big.Int, _boxList []string, _shoeList []string) (*types.Transaction, error) {
	return _Mint.Contract.AddPublicMintData(&_Mint.TransactOpts, _tokenIds, _boxList, _shoeList)
}

// AddWlMintData is a paid mutator transaction binding the contract method 0xb745f42d.
//
// Solidity: function addWlMintData(uint256[] _tokenIds, string[] _boxList, string[] _shoeList) returns()
func (_Mint *MintTransactor) AddWlMintData(opts *bind.TransactOpts, _tokenIds []*big.Int, _boxList []string, _shoeList []string) (*types.Transaction, error) {
	return _Mint.contract.Transact(opts, "addWlMintData", _tokenIds, _boxList, _shoeList)
}

// AddWlMintData is a paid mutator transaction binding the contract method 0xb745f42d.
//
// Solidity: function addWlMintData(uint256[] _tokenIds, string[] _boxList, string[] _shoeList) returns()
func (_Mint *MintSession) AddWlMintData(_tokenIds []*big.Int, _boxList []string, _shoeList []string) (*types.Transaction, error) {
	return _Mint.Contract.AddWlMintData(&_Mint.TransactOpts, _tokenIds, _boxList, _shoeList)
}

// AddWlMintData is a paid mutator transaction binding the contract method 0xb745f42d.
//
// Solidity: function addWlMintData(uint256[] _tokenIds, string[] _boxList, string[] _shoeList) returns()
func (_Mint *MintTransactorSession) AddWlMintData(_tokenIds []*big.Int, _boxList []string, _shoeList []string) (*types.Transaction, error) {
	return _Mint.Contract.AddWlMintData(&_Mint.TransactOpts, _tokenIds, _boxList, _shoeList)
}

// FreeMint is a paid mutator transaction binding the contract method 0x88d15d50.
//
// Solidity: function freeMint(bytes32[] _merkleProof) returns()
func (_Mint *MintTransactor) FreeMint(opts *bind.TransactOpts, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _Mint.contract.Transact(opts, "freeMint", _merkleProof)
}

// FreeMint is a paid mutator transaction binding the contract method 0x88d15d50.
//
// Solidity: function freeMint(bytes32[] _merkleProof) returns()
func (_Mint *MintSession) FreeMint(_merkleProof [][32]byte) (*types.Transaction, error) {
	return _Mint.Contract.FreeMint(&_Mint.TransactOpts, _merkleProof)
}

// FreeMint is a paid mutator transaction binding the contract method 0x88d15d50.
//
// Solidity: function freeMint(bytes32[] _merkleProof) returns()
func (_Mint *MintTransactorSession) FreeMint(_merkleProof [][32]byte) (*types.Transaction, error) {
	return _Mint.Contract.FreeMint(&_Mint.TransactOpts, _merkleProof)
}

// OgMint is a paid mutator transaction binding the contract method 0x918ed5d5.
//
// Solidity: function ogMint(bytes32[] _merkleProof) payable returns()
func (_Mint *MintTransactor) OgMint(opts *bind.TransactOpts, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _Mint.contract.Transact(opts, "ogMint", _merkleProof)
}

// OgMint is a paid mutator transaction binding the contract method 0x918ed5d5.
//
// Solidity: function ogMint(bytes32[] _merkleProof) payable returns()
func (_Mint *MintSession) OgMint(_merkleProof [][32]byte) (*types.Transaction, error) {
	return _Mint.Contract.OgMint(&_Mint.TransactOpts, _merkleProof)
}

// OgMint is a paid mutator transaction binding the contract method 0x918ed5d5.
//
// Solidity: function ogMint(bytes32[] _merkleProof) payable returns()
func (_Mint *MintTransactorSession) OgMint(_merkleProof [][32]byte) (*types.Transaction, error) {
	return _Mint.Contract.OgMint(&_Mint.TransactOpts, _merkleProof)
}

// PublicMint is a paid mutator transaction binding the contract method 0x26092b83.
//
// Solidity: function publicMint() payable returns()
func (_Mint *MintTransactor) PublicMint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mint.contract.Transact(opts, "publicMint")
}

// PublicMint is a paid mutator transaction binding the contract method 0x26092b83.
//
// Solidity: function publicMint() payable returns()
func (_Mint *MintSession) PublicMint() (*types.Transaction, error) {
	return _Mint.Contract.PublicMint(&_Mint.TransactOpts)
}

// PublicMint is a paid mutator transaction binding the contract method 0x26092b83.
//
// Solidity: function publicMint() payable returns()
func (_Mint *MintTransactorSession) PublicMint() (*types.Transaction, error) {
	return _Mint.Contract.PublicMint(&_Mint.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mint *MintTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mint.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mint *MintSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mint.Contract.RenounceOwnership(&_Mint.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mint *MintTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mint.Contract.RenounceOwnership(&_Mint.TransactOpts)
}

// ResetOwnership is a paid mutator transaction binding the contract method 0xbe3af789.
//
// Solidity: function resetOwnership(address _newOwner) returns()
func (_Mint *MintTransactor) ResetOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Mint.contract.Transact(opts, "resetOwnership", _newOwner)
}

// ResetOwnership is a paid mutator transaction binding the contract method 0xbe3af789.
//
// Solidity: function resetOwnership(address _newOwner) returns()
func (_Mint *MintSession) ResetOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Mint.Contract.ResetOwnership(&_Mint.TransactOpts, _newOwner)
}

// ResetOwnership is a paid mutator transaction binding the contract method 0xbe3af789.
//
// Solidity: function resetOwnership(address _newOwner) returns()
func (_Mint *MintTransactorSession) ResetOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Mint.Contract.ResetOwnership(&_Mint.TransactOpts, _newOwner)
}

// SetFMMerkleRoot is a paid mutator transaction binding the contract method 0x5e6b248a.
//
// Solidity: function setFMMerkleRoot(bytes32 _fm) returns()
func (_Mint *MintTransactor) SetFMMerkleRoot(opts *bind.TransactOpts, _fm [32]byte) (*types.Transaction, error) {
	return _Mint.contract.Transact(opts, "setFMMerkleRoot", _fm)
}

// SetFMMerkleRoot is a paid mutator transaction binding the contract method 0x5e6b248a.
//
// Solidity: function setFMMerkleRoot(bytes32 _fm) returns()
func (_Mint *MintSession) SetFMMerkleRoot(_fm [32]byte) (*types.Transaction, error) {
	return _Mint.Contract.SetFMMerkleRoot(&_Mint.TransactOpts, _fm)
}

// SetFMMerkleRoot is a paid mutator transaction binding the contract method 0x5e6b248a.
//
// Solidity: function setFMMerkleRoot(bytes32 _fm) returns()
func (_Mint *MintTransactorSession) SetFMMerkleRoot(_fm [32]byte) (*types.Transaction, error) {
	return _Mint.Contract.SetFMMerkleRoot(&_Mint.TransactOpts, _fm)
}

// SetFMPos is a paid mutator transaction binding the contract method 0xac9421ca.
//
// Solidity: function setFMPos(uint256 _pos) returns()
func (_Mint *MintTransactor) SetFMPos(opts *bind.TransactOpts, _pos *big.Int) (*types.Transaction, error) {
	return _Mint.contract.Transact(opts, "setFMPos", _pos)
}

// SetFMPos is a paid mutator transaction binding the contract method 0xac9421ca.
//
// Solidity: function setFMPos(uint256 _pos) returns()
func (_Mint *MintSession) SetFMPos(_pos *big.Int) (*types.Transaction, error) {
	return _Mint.Contract.SetFMPos(&_Mint.TransactOpts, _pos)
}

// SetFMPos is a paid mutator transaction binding the contract method 0xac9421ca.
//
// Solidity: function setFMPos(uint256 _pos) returns()
func (_Mint *MintTransactorSession) SetFMPos(_pos *big.Int) (*types.Transaction, error) {
	return _Mint.Contract.SetFMPos(&_Mint.TransactOpts, _pos)
}

// SetFreeMintTime is a paid mutator transaction binding the contract method 0xf26b8119.
//
// Solidity: function setFreeMintTime(uint256 _start, uint256 _end) returns()
func (_Mint *MintTransactor) SetFreeMintTime(opts *bind.TransactOpts, _start *big.Int, _end *big.Int) (*types.Transaction, error) {
	return _Mint.contract.Transact(opts, "setFreeMintTime", _start, _end)
}

// SetFreeMintTime is a paid mutator transaction binding the contract method 0xf26b8119.
//
// Solidity: function setFreeMintTime(uint256 _start, uint256 _end) returns()
func (_Mint *MintSession) SetFreeMintTime(_start *big.Int, _end *big.Int) (*types.Transaction, error) {
	return _Mint.Contract.SetFreeMintTime(&_Mint.TransactOpts, _start, _end)
}

// SetFreeMintTime is a paid mutator transaction binding the contract method 0xf26b8119.
//
// Solidity: function setFreeMintTime(uint256 _start, uint256 _end) returns()
func (_Mint *MintTransactorSession) SetFreeMintTime(_start *big.Int, _end *big.Int) (*types.Transaction, error) {
	return _Mint.Contract.SetFreeMintTime(&_Mint.TransactOpts, _start, _end)
}

// SetMaxFreeMint is a paid mutator transaction binding the contract method 0x742a4c9b.
//
// Solidity: function setMaxFreeMint(uint256 _maxMint) returns()
func (_Mint *MintTransactor) SetMaxFreeMint(opts *bind.TransactOpts, _maxMint *big.Int) (*types.Transaction, error) {
	return _Mint.contract.Transact(opts, "setMaxFreeMint", _maxMint)
}

// SetMaxFreeMint is a paid mutator transaction binding the contract method 0x742a4c9b.
//
// Solidity: function setMaxFreeMint(uint256 _maxMint) returns()
func (_Mint *MintSession) SetMaxFreeMint(_maxMint *big.Int) (*types.Transaction, error) {
	return _Mint.Contract.SetMaxFreeMint(&_Mint.TransactOpts, _maxMint)
}

// SetMaxFreeMint is a paid mutator transaction binding the contract method 0x742a4c9b.
//
// Solidity: function setMaxFreeMint(uint256 _maxMint) returns()
func (_Mint *MintTransactorSession) SetMaxFreeMint(_maxMint *big.Int) (*types.Transaction, error) {
	return _Mint.Contract.SetMaxFreeMint(&_Mint.TransactOpts, _maxMint)
}

// SetMaxMint is a paid mutator transaction binding the contract method 0x547520fe.
//
// Solidity: function setMaxMint(uint256 _maxTotalMint) returns()
func (_Mint *MintTransactor) SetMaxMint(opts *bind.TransactOpts, _maxTotalMint *big.Int) (*types.Transaction, error) {
	return _Mint.contract.Transact(opts, "setMaxMint", _maxTotalMint)
}

// SetMaxMint is a paid mutator transaction binding the contract method 0x547520fe.
//
// Solidity: function setMaxMint(uint256 _maxTotalMint) returns()
func (_Mint *MintSession) SetMaxMint(_maxTotalMint *big.Int) (*types.Transaction, error) {
	return _Mint.Contract.SetMaxMint(&_Mint.TransactOpts, _maxTotalMint)
}

// SetMaxMint is a paid mutator transaction binding the contract method 0x547520fe.
//
// Solidity: function setMaxMint(uint256 _maxTotalMint) returns()
func (_Mint *MintTransactorSession) SetMaxMint(_maxTotalMint *big.Int) (*types.Transaction, error) {
	return _Mint.Contract.SetMaxMint(&_Mint.TransactOpts, _maxTotalMint)
}

// SetOGMerkleRoot is a paid mutator transaction binding the contract method 0x25c2c020.
//
// Solidity: function setOGMerkleRoot(bytes32 _og) returns()
func (_Mint *MintTransactor) SetOGMerkleRoot(opts *bind.TransactOpts, _og [32]byte) (*types.Transaction, error) {
	return _Mint.contract.Transact(opts, "setOGMerkleRoot", _og)
}

// SetOGMerkleRoot is a paid mutator transaction binding the contract method 0x25c2c020.
//
// Solidity: function setOGMerkleRoot(bytes32 _og) returns()
func (_Mint *MintSession) SetOGMerkleRoot(_og [32]byte) (*types.Transaction, error) {
	return _Mint.Contract.SetOGMerkleRoot(&_Mint.TransactOpts, _og)
}

// SetOGMerkleRoot is a paid mutator transaction binding the contract method 0x25c2c020.
//
// Solidity: function setOGMerkleRoot(bytes32 _og) returns()
func (_Mint *MintTransactorSession) SetOGMerkleRoot(_og [32]byte) (*types.Transaction, error) {
	return _Mint.Contract.SetOGMerkleRoot(&_Mint.TransactOpts, _og)
}

// SetOGMintPrice is a paid mutator transaction binding the contract method 0x74e6ddf7.
//
// Solidity: function setOGMintPrice(uint256 _price) returns()
func (_Mint *MintTransactor) SetOGMintPrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _Mint.contract.Transact(opts, "setOGMintPrice", _price)
}

// SetOGMintPrice is a paid mutator transaction binding the contract method 0x74e6ddf7.
//
// Solidity: function setOGMintPrice(uint256 _price) returns()
func (_Mint *MintSession) SetOGMintPrice(_price *big.Int) (*types.Transaction, error) {
	return _Mint.Contract.SetOGMintPrice(&_Mint.TransactOpts, _price)
}

// SetOGMintPrice is a paid mutator transaction binding the contract method 0x74e6ddf7.
//
// Solidity: function setOGMintPrice(uint256 _price) returns()
func (_Mint *MintTransactorSession) SetOGMintPrice(_price *big.Int) (*types.Transaction, error) {
	return _Mint.Contract.SetOGMintPrice(&_Mint.TransactOpts, _price)
}

// SetOGMintTime is a paid mutator transaction binding the contract method 0x373be6a2.
//
// Solidity: function setOGMintTime(uint256 _start, uint256 _end) returns()
func (_Mint *MintTransactor) SetOGMintTime(opts *bind.TransactOpts, _start *big.Int, _end *big.Int) (*types.Transaction, error) {
	return _Mint.contract.Transact(opts, "setOGMintTime", _start, _end)
}

// SetOGMintTime is a paid mutator transaction binding the contract method 0x373be6a2.
//
// Solidity: function setOGMintTime(uint256 _start, uint256 _end) returns()
func (_Mint *MintSession) SetOGMintTime(_start *big.Int, _end *big.Int) (*types.Transaction, error) {
	return _Mint.Contract.SetOGMintTime(&_Mint.TransactOpts, _start, _end)
}

// SetOGMintTime is a paid mutator transaction binding the contract method 0x373be6a2.
//
// Solidity: function setOGMintTime(uint256 _start, uint256 _end) returns()
func (_Mint *MintTransactorSession) SetOGMintTime(_start *big.Int, _end *big.Int) (*types.Transaction, error) {
	return _Mint.Contract.SetOGMintTime(&_Mint.TransactOpts, _start, _end)
}

// SetOGPos is a paid mutator transaction binding the contract method 0x8bc746f0.
//
// Solidity: function setOGPos(uint256 _pos) returns()
func (_Mint *MintTransactor) SetOGPos(opts *bind.TransactOpts, _pos *big.Int) (*types.Transaction, error) {
	return _Mint.contract.Transact(opts, "setOGPos", _pos)
}

// SetOGPos is a paid mutator transaction binding the contract method 0x8bc746f0.
//
// Solidity: function setOGPos(uint256 _pos) returns()
func (_Mint *MintSession) SetOGPos(_pos *big.Int) (*types.Transaction, error) {
	return _Mint.Contract.SetOGPos(&_Mint.TransactOpts, _pos)
}

// SetOGPos is a paid mutator transaction binding the contract method 0x8bc746f0.
//
// Solidity: function setOGPos(uint256 _pos) returns()
func (_Mint *MintTransactorSession) SetOGPos(_pos *big.Int) (*types.Transaction, error) {
	return _Mint.Contract.SetOGPos(&_Mint.TransactOpts, _pos)
}

// SetPPos is a paid mutator transaction binding the contract method 0xbc5bca27.
//
// Solidity: function setPPos(uint256 _pos) returns()
func (_Mint *MintTransactor) SetPPos(opts *bind.TransactOpts, _pos *big.Int) (*types.Transaction, error) {
	return _Mint.contract.Transact(opts, "setPPos", _pos)
}

// SetPPos is a paid mutator transaction binding the contract method 0xbc5bca27.
//
// Solidity: function setPPos(uint256 _pos) returns()
func (_Mint *MintSession) SetPPos(_pos *big.Int) (*types.Transaction, error) {
	return _Mint.Contract.SetPPos(&_Mint.TransactOpts, _pos)
}

// SetPPos is a paid mutator transaction binding the contract method 0xbc5bca27.
//
// Solidity: function setPPos(uint256 _pos) returns()
func (_Mint *MintTransactorSession) SetPPos(_pos *big.Int) (*types.Transaction, error) {
	return _Mint.Contract.SetPPos(&_Mint.TransactOpts, _pos)
}

// SetPublicMintPrice is a paid mutator transaction binding the contract method 0x5d82cf6e.
//
// Solidity: function setPublicMintPrice(uint256 _price) returns()
func (_Mint *MintTransactor) SetPublicMintPrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _Mint.contract.Transact(opts, "setPublicMintPrice", _price)
}

// SetPublicMintPrice is a paid mutator transaction binding the contract method 0x5d82cf6e.
//
// Solidity: function setPublicMintPrice(uint256 _price) returns()
func (_Mint *MintSession) SetPublicMintPrice(_price *big.Int) (*types.Transaction, error) {
	return _Mint.Contract.SetPublicMintPrice(&_Mint.TransactOpts, _price)
}

// SetPublicMintPrice is a paid mutator transaction binding the contract method 0x5d82cf6e.
//
// Solidity: function setPublicMintPrice(uint256 _price) returns()
func (_Mint *MintTransactorSession) SetPublicMintPrice(_price *big.Int) (*types.Transaction, error) {
	return _Mint.Contract.SetPublicMintPrice(&_Mint.TransactOpts, _price)
}

// SetPublicMintTime is a paid mutator transaction binding the contract method 0x911690e2.
//
// Solidity: function setPublicMintTime(uint256 _start, uint256 _end) returns()
func (_Mint *MintTransactor) SetPublicMintTime(opts *bind.TransactOpts, _start *big.Int, _end *big.Int) (*types.Transaction, error) {
	return _Mint.contract.Transact(opts, "setPublicMintTime", _start, _end)
}

// SetPublicMintTime is a paid mutator transaction binding the contract method 0x911690e2.
//
// Solidity: function setPublicMintTime(uint256 _start, uint256 _end) returns()
func (_Mint *MintSession) SetPublicMintTime(_start *big.Int, _end *big.Int) (*types.Transaction, error) {
	return _Mint.Contract.SetPublicMintTime(&_Mint.TransactOpts, _start, _end)
}

// SetPublicMintTime is a paid mutator transaction binding the contract method 0x911690e2.
//
// Solidity: function setPublicMintTime(uint256 _start, uint256 _end) returns()
func (_Mint *MintTransactorSession) SetPublicMintTime(_start *big.Int, _end *big.Int) (*types.Transaction, error) {
	return _Mint.Contract.SetPublicMintTime(&_Mint.TransactOpts, _start, _end)
}

// SetWLMerkleRoot is a paid mutator transaction binding the contract method 0x7d44fd11.
//
// Solidity: function setWLMerkleRoot(bytes32 _wl) returns()
func (_Mint *MintTransactor) SetWLMerkleRoot(opts *bind.TransactOpts, _wl [32]byte) (*types.Transaction, error) {
	return _Mint.contract.Transact(opts, "setWLMerkleRoot", _wl)
}

// SetWLMerkleRoot is a paid mutator transaction binding the contract method 0x7d44fd11.
//
// Solidity: function setWLMerkleRoot(bytes32 _wl) returns()
func (_Mint *MintSession) SetWLMerkleRoot(_wl [32]byte) (*types.Transaction, error) {
	return _Mint.Contract.SetWLMerkleRoot(&_Mint.TransactOpts, _wl)
}

// SetWLMerkleRoot is a paid mutator transaction binding the contract method 0x7d44fd11.
//
// Solidity: function setWLMerkleRoot(bytes32 _wl) returns()
func (_Mint *MintTransactorSession) SetWLMerkleRoot(_wl [32]byte) (*types.Transaction, error) {
	return _Mint.Contract.SetWLMerkleRoot(&_Mint.TransactOpts, _wl)
}

// SetWLMintPrice is a paid mutator transaction binding the contract method 0xe1e62b02.
//
// Solidity: function setWLMintPrice(uint256 _price) returns()
func (_Mint *MintTransactor) SetWLMintPrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _Mint.contract.Transact(opts, "setWLMintPrice", _price)
}

// SetWLMintPrice is a paid mutator transaction binding the contract method 0xe1e62b02.
//
// Solidity: function setWLMintPrice(uint256 _price) returns()
func (_Mint *MintSession) SetWLMintPrice(_price *big.Int) (*types.Transaction, error) {
	return _Mint.Contract.SetWLMintPrice(&_Mint.TransactOpts, _price)
}

// SetWLMintPrice is a paid mutator transaction binding the contract method 0xe1e62b02.
//
// Solidity: function setWLMintPrice(uint256 _price) returns()
func (_Mint *MintTransactorSession) SetWLMintPrice(_price *big.Int) (*types.Transaction, error) {
	return _Mint.Contract.SetWLMintPrice(&_Mint.TransactOpts, _price)
}

// SetWLPos is a paid mutator transaction binding the contract method 0x6b527d24.
//
// Solidity: function setWLPos(uint256 _pos) returns()
func (_Mint *MintTransactor) SetWLPos(opts *bind.TransactOpts, _pos *big.Int) (*types.Transaction, error) {
	return _Mint.contract.Transact(opts, "setWLPos", _pos)
}

// SetWLPos is a paid mutator transaction binding the contract method 0x6b527d24.
//
// Solidity: function setWLPos(uint256 _pos) returns()
func (_Mint *MintSession) SetWLPos(_pos *big.Int) (*types.Transaction, error) {
	return _Mint.Contract.SetWLPos(&_Mint.TransactOpts, _pos)
}

// SetWLPos is a paid mutator transaction binding the contract method 0x6b527d24.
//
// Solidity: function setWLPos(uint256 _pos) returns()
func (_Mint *MintTransactorSession) SetWLPos(_pos *big.Int) (*types.Transaction, error) {
	return _Mint.Contract.SetWLPos(&_Mint.TransactOpts, _pos)
}

// SetWhiteListMintTime is a paid mutator transaction binding the contract method 0x3c01a68c.
//
// Solidity: function setWhiteListMintTime(uint256 _start, uint256 _end) returns()
func (_Mint *MintTransactor) SetWhiteListMintTime(opts *bind.TransactOpts, _start *big.Int, _end *big.Int) (*types.Transaction, error) {
	return _Mint.contract.Transact(opts, "setWhiteListMintTime", _start, _end)
}

// SetWhiteListMintTime is a paid mutator transaction binding the contract method 0x3c01a68c.
//
// Solidity: function setWhiteListMintTime(uint256 _start, uint256 _end) returns()
func (_Mint *MintSession) SetWhiteListMintTime(_start *big.Int, _end *big.Int) (*types.Transaction, error) {
	return _Mint.Contract.SetWhiteListMintTime(&_Mint.TransactOpts, _start, _end)
}

// SetWhiteListMintTime is a paid mutator transaction binding the contract method 0x3c01a68c.
//
// Solidity: function setWhiteListMintTime(uint256 _start, uint256 _end) returns()
func (_Mint *MintTransactorSession) SetWhiteListMintTime(_start *big.Int, _end *big.Int) (*types.Transaction, error) {
	return _Mint.Contract.SetWhiteListMintTime(&_Mint.TransactOpts, _start, _end)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mint *MintTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Mint.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mint *MintSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mint.Contract.TransferOwnership(&_Mint.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mint *MintTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mint.Contract.TransferOwnership(&_Mint.TransactOpts, newOwner)
}

// TransferToNewOwner is a paid mutator transaction binding the contract method 0x7372e9be.
//
// Solidity: function transferToNewOwner(address _new) returns()
func (_Mint *MintTransactor) TransferToNewOwner(opts *bind.TransactOpts, _new common.Address) (*types.Transaction, error) {
	return _Mint.contract.Transact(opts, "transferToNewOwner", _new)
}

// TransferToNewOwner is a paid mutator transaction binding the contract method 0x7372e9be.
//
// Solidity: function transferToNewOwner(address _new) returns()
func (_Mint *MintSession) TransferToNewOwner(_new common.Address) (*types.Transaction, error) {
	return _Mint.Contract.TransferToNewOwner(&_Mint.TransactOpts, _new)
}

// TransferToNewOwner is a paid mutator transaction binding the contract method 0x7372e9be.
//
// Solidity: function transferToNewOwner(address _new) returns()
func (_Mint *MintTransactorSession) TransferToNewOwner(_new common.Address) (*types.Transaction, error) {
	return _Mint.Contract.TransferToNewOwner(&_Mint.TransactOpts, _new)
}

// WhiteListMint is a paid mutator transaction binding the contract method 0x97254e55.
//
// Solidity: function whiteListMint(bytes32[] _merkleProof) payable returns()
func (_Mint *MintTransactor) WhiteListMint(opts *bind.TransactOpts, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _Mint.contract.Transact(opts, "whiteListMint", _merkleProof)
}

// WhiteListMint is a paid mutator transaction binding the contract method 0x97254e55.
//
// Solidity: function whiteListMint(bytes32[] _merkleProof) payable returns()
func (_Mint *MintSession) WhiteListMint(_merkleProof [][32]byte) (*types.Transaction, error) {
	return _Mint.Contract.WhiteListMint(&_Mint.TransactOpts, _merkleProof)
}

// WhiteListMint is a paid mutator transaction binding the contract method 0x97254e55.
//
// Solidity: function whiteListMint(bytes32[] _merkleProof) payable returns()
func (_Mint *MintTransactorSession) WhiteListMint(_merkleProof [][32]byte) (*types.Transaction, error) {
	return _Mint.Contract.WhiteListMint(&_Mint.TransactOpts, _merkleProof)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Mint *MintTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mint.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Mint *MintSession) Withdraw() (*types.Transaction, error) {
	return _Mint.Contract.Withdraw(&_Mint.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Mint *MintTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Mint.Contract.Withdraw(&_Mint.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Mint *MintTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Mint.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Mint *MintSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Mint.Contract.Fallback(&_Mint.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Mint *MintTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Mint.Contract.Fallback(&_Mint.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Mint *MintTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mint.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Mint *MintSession) Receive() (*types.Transaction, error) {
	return _Mint.Contract.Receive(&_Mint.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Mint *MintTransactorSession) Receive() (*types.Transaction, error) {
	return _Mint.Contract.Receive(&_Mint.TransactOpts)
}

// MintOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Mint contract.
type MintOwnershipTransferredIterator struct {
	Event *MintOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MintOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintOwnershipTransferred)
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
		it.Event = new(MintOwnershipTransferred)
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
func (it *MintOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintOwnershipTransferred represents a OwnershipTransferred event raised by the Mint contract.
type MintOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mint *MintFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MintOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mint.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MintOwnershipTransferredIterator{contract: _Mint.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mint *MintFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MintOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mint.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintOwnershipTransferred)
				if err := _Mint.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Mint *MintFilterer) ParseOwnershipTransferred(log types.Log) (*MintOwnershipTransferred, error) {
	event := new(MintOwnershipTransferred)
	if err := _Mint.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
