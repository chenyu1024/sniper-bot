package token_vmt_test

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"log"
	"math/big"
	"testing"
)

func TestTransfer(t *testing.T) {
	var m = VmtSetting
	vAuth, err := bind.NewKeyedTransactorWithChainID(m.ownerPk, big.NewInt(testChainId))
	if err != nil {
		log.Fatalln(err.Error())
	}
	ethers, _ := big.NewFloat(params.Ether).Int(nil)
	transferAmountWithoutDecimals := big.NewInt(0).Mul(ethers, transferAmount)
	// start transfer
	tx1, err := m.VstStorageInstance.SendTORunBa(vAuth, VST_ADDRESS, common.HexToAddress(account1), transferAmountWithoutDecimals)
	if err != nil {
		log.Println(123)
		log.Fatalln(err.Error())
	}
	txnIsOk(tx1.Hash())
	tx2, err := m.VstStorageInstance.SendTORunBa(vAuth, VST_ADDRESS, common.HexToAddress(account2), transferAmountWithoutDecimals)
	if err != nil {
		log.Println(123)
		log.Fatalln(err.Error())
	}
	txnIsOk(tx2.Hash())
	balance1, _ := m.VstInstance.BalanceOf(nil, common.HexToAddress(account1))
	if balance1.Cmp(transferAmountWithoutDecimals) != 0 {
		t.Log(balance1)
		t.Log(transferAmountWithoutDecimals)
		t.Fatal("测试失败，转账余额错误1")
	}
	balance2, _ := m.VstInstance.BalanceOf(nil, common.HexToAddress(account2))
	if balance2.Cmp(transferAmountWithoutDecimals) != 0 {
		t.Fatal("测试失败，转账余额错误2")
	}
}
