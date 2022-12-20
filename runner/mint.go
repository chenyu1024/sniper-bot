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
	userAbi "wxjkzcy/contract/other"
)

// abigen -abi .\tomato.abi -pkg abi -out abi.go
func (e *ethRunner) getAbiInstance(chain string) *userAbi.Abi {
	targetAddress := common.HexToAddress(viper.GetString("targetContract"))
	instance, err := userAbi.NewAbi(targetAddress, e.getClient(chain))
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Printf("[STATUS] Load abi success.")
	return instance
}
func (e *ethRunner) getOwner(chain string) common.Address {
	instance := e.getAbiInstance(chain)
	owner, err := instance.Owner(nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("owner", owner.String())
	return owner
}

func (e *ethRunner) batchMint(chain string, to common.Address, tokenIds []*big.Int, tokenUris, boxUris []string) {
	auth, err := bind.NewKeyedTransactorWithChainID(e.privateKey, big.NewInt(consts.ChainIdMap[chain]))
	if err != nil {
		log.Fatal(err)
	}
	nonce, err := e.getClient(chain).PendingNonceAt(context.Background(), e.fromAddress)
	//noncePending, err := e.getClient(chain).PendingTransactionCount(context.Background())
	//suggestGas, err := e.getClient(chain).SuggestGasPrice(context.Background())
	log.Println("Nonce:", nonce)
	//log.Printf("suggest Gas %v", suggestGas.String())
	//log.Println(noncePending)
	if err != nil {
		log.Fatalln(err.Error())
	}
	//auth.Nonce = big.NewInt(int64(nonce))
	auth.Nonce = nil
	//auth.Value = nil
	auth.GasLimit = viper.GetUint64("gasLimit")
	gas, _ := big.NewFloat(viper.GetFloat64("gasPrice") * params.GWei).Int(nil)
	auth.GasPrice = gas

	instance := e.getAbiInstance(chain)
	tx, err := instance.MintBatch200(auth, to, tokenIds, tokenUris, boxUris)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println(tx.Hash().String())

	//instance.CheckCertification(nil, )
}

func (e *ethRunner) approveFromLeader(chain string, tokenIds []*big.Int) {
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

	instance := e.getAbiInstance(chain)
	tx, err := instance.BatchApprove(auth, tokenIds, common.HexToAddress("todesadasd")) // TODO
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println(tx.Hash().String())
}

func (e *ethRunner) isLeader(chain string) bool {
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

	instance := e.getAbiInstance(chain)
	isLeader, err := instance.IsLeader(nil, common.HexToAddress("0x89b2a39e930E49f587F8F428177411E2241Fa4C9"))
	if err != nil {
		log.Fatalln(err.Error())
	}
	return isLeader
}

func (e *ethRunner) batchAirdrop(chain string, tokenIdList []*big.Int, toList []common.Address) {
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

	instance := e.getAbiInstance(chain)
	tx, err := instance.Airdrop(auth, tokenIdList, toList)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println(tx.Hash().String())
}

func (e *ethRunner) openBox(chain string, tokenId *big.Int) {
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

	instance := e.getAbiInstance(chain)
	tx, err := instance.OpenBox(auth, tokenId)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println(tx.Hash().String())
}
func (e *ethRunner) addLeader(chain string, address common.Address) {
	instance := e.getAbiInstance(chain)

	auth, err := bind.NewKeyedTransactorWithChainID(e.privateKey, big.NewInt(consts.ChainIdMap[chain]))
	if err != nil {
		log.Fatal(err)
	}
	leader, err := instance.AddLeader(auth, []common.Address{address})
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println(leader.Nonce(), "|", leader.Hash())

}
