package token_vmt_test

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"log"
	"math/big"
	"testing"
	"time"
	"wxjkzcy/contract/token/vmt"
	"wxjkzcy/contract/token/vstStorage"
)

var (
	//owner_address       = "0x4F69FC7bb16e4869FE51BF2E60d860b040a63282"
	rpcAddress  = "https://data-seed-prebsc-1-s1.binance.org:8545/"
	VST_ADDRESS = common.HexToAddress("0x32460D300C6e02c9c63c19e1EfB8A944829ccE5c")

	storageContract     = "0x525f5D722D1d2e4BEf68E0ce068310D0409d6dF1"
	VST_STORAGE_ADDRESS = common.HexToAddress("0x525f5D722D1d2e4BEf68E0ce068310D0409d6dF1")

	owner_pri      = "31ec230279cfb9274bc27b68416d83b1aada47cda887f09546e1d567e6bcc1d2"
	account1       = "0x41d7C8879F832584672744e245a253cEAf0ec012"
	account2       = "0x7c02f3D0c26b4B1CFf9c543D13f90f4e9988E05c"
	testChainId    = int64(97)
	mintAmount     = big.NewInt(1000000000) //
	transferAmount = big.NewInt(500)
)

type Setting struct {
	VstInstance        *vmt.Vmt
	VstStorageInstance *vstStorage.VstS
	ownerPk            *ecdsa.PrivateKey
	client             *ethclient.Client
}

var VmtSetting Setting

func init() {
	client, err := ethclient.Dial(rpcAddress)
	if err != nil {
		log.Fatalln(err.Error())
	}
	vstInstance, err := vmt.NewVmt(VST_ADDRESS, client)
	if err != nil {
		log.Fatalln(err.Error())
	}
	vstStorageInstance, err := vstStorage.NewVstS(VST_STORAGE_ADDRESS, client)
	if err != nil {
		log.Fatalln(err.Error())
	}
	accountPk, err := crypto.HexToECDSA(owner_pri)
	if err != nil {
		log.Fatalln(err.Error())
	}
	VmtSetting = Setting{
		VstInstance:        vstInstance,
		VstStorageInstance: vstStorageInstance,
		ownerPk:            accountPk,
		client:             client,
	}
}

func txnIsOk(hash common.Hash) bool {
	for i := 1; ; i++ {
		receipt, _ := VmtSetting.client.TransactionReceipt(context.Background(), hash)
		if receipt == nil {
			continue
		} else if big.NewInt(int64(receipt.Status)).Cmp(big.NewInt(1)) == 0 {
			//log.Printf("交易成功！")
			return true
		} else if big.NewInt(int64(receipt.Status)).Cmp(big.NewInt(1)) != 0 {
			log.Fatalln("交易失败！")
			return false
		}
		time.Sleep(time.Second)
	}
}

func TestVstMint(t *testing.T) {
	var m = VmtSetting
	vAuth, err := bind.NewKeyedTransactorWithChainID(m.ownerPk, big.NewInt(testChainId))
	if err != nil {
		log.Fatalln(err.Error())
	}
	ethers, _ := big.NewFloat(params.Ether).Int(nil)
	mintAmountWithoutDecimals := big.NewInt(0).Mul(ethers, mintAmount)
	// start mint
	tx, err := m.VstInstance.Mint(vAuth, common.HexToAddress(storageContract), mintAmountWithoutDecimals)
	if err != nil {
		t.Fatal(err.Error())
	}
	log.Println(tx.Hash().String())
	txnIsOk(tx.Hash())

	accountBalance, _ := m.VstInstance.BalanceOf(nil, common.HexToAddress(storageContract))
	if accountBalance.Cmp(mintAmountWithoutDecimals) != 0 {
		t.Log(accountBalance)
		t.Log(mintAmountWithoutDecimals)
		t.Fatal("mint数量不符")
	}
}
