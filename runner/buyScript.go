package runner

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/spf13/viper"
	"log"
	"math/big"
	"wxjkzcy/consts"
	"wxjkzcy/contract/buy"
)

func (e *ethRunner) getBuyAbiInstance(chain string) *buy.Buy {
	targetAddress := common.HexToAddress(viper.GetString("0x6faC9eFc3C0414A94E7156A95C4EcB99a6F7a121"))
	instance, err := buy.NewBuy(targetAddress, e.getClient(chain))
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Printf("[STATUS] Load abi success.")
	return instance
}
func (e *ethRunner) sBuy(chain string, path0, path1 common.Address, amount *big.Int) {
	instance := e.getBuyAbiInstance(chain)
	auth, err := bind.NewKeyedTransactorWithChainID(e.privateKey, big.NewInt(consts.ChainIdMap[chain]))
	if err != nil {
		log.Fatal(err.Error())
	}
	auth.GasLimit = viper.GetUint64("gasLimit")
	gas, _ := big.NewFloat(viper.GetFloat64("gasPrice") * params.GWei).Int(nil)
	auth.GasPrice = gas
	buy1, err := instance.SimpleBuyv1(auth, path0, path1, amount)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println(buy1.Hash().String())
}

func (e *ethRunner) buy(chain string, path0, path1 common.Address, amount, slippage *big.Int) {
	instance := e.getBuyAbiInstance(chain)
	auth, err := bind.NewKeyedTransactorWithChainID(e.privateKey, big.NewInt(consts.ChainIdMap[chain]))
	if err != nil {
		log.Fatal(err.Error())
	}
	auth.GasLimit = viper.GetUint64("gasLimit")
	gas, _ := big.NewFloat(viper.GetFloat64("gasPrice") * params.GWei).Int(nil)
	auth.GasPrice = gas
	buy1, err := instance.SimpleBuy(auth, path0, path1, amount, slippage)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println(buy1.Hash().String())
}
