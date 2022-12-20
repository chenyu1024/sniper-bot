package runner

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/spf13/viper"
	"log"
	"math/big"
	"wxjkzcy/consts"
	airdropAbi "wxjkzcy/contract/airdrop"
)

func (e *ethRunner) getAirdropAbiInstance(chain string) *airdropAbi.Airdrop {
	targetAddress := common.HexToAddress(viper.GetString("airdropContract"))
	instance, err := airdropAbi.NewAirdrop(targetAddress, e.getClient(chain))
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Printf("[STATUS] Load abi success.")
	return instance
}

func (e *ethRunner) addAirdropList(chain string, tokenIds []*big.Int) {
	auth, err := bind.NewKeyedTransactorWithChainID(e.privateKey, big.NewInt(consts.ChainIdMap[chain]))
	if err != nil {
		log.Fatal(err)
	}
	nonce, err := e.getClient(chain).PendingNonceAt(context.Background(), e.fromAddress)
	log.Println("Nonce:", nonce)
	if err != nil {
		log.Fatalln(err.Error())
	}
	auth.Nonce = nil
	auth.GasLimit = viper.GetUint64("gasLimit")
	gas, _ := big.NewFloat(viper.GetFloat64("gasPrice") * params.GWei).Int(nil)
	auth.GasPrice = gas

	instance := e.getAirdropAbiInstance(chain)
	tx, err := instance.AddAirdropTokenIds(auth, tokenIds)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println(tx.Hash().String())
}
