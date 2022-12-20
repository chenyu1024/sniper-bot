package batchTransfer_test

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"testing"
	"wxjkzcy/contract/batchTransfer"
)

var (
	rpc          = "https://data-seed-prebsc-1-s3.binance.org:8545"
	transferAddr = "0x28274619d441036dBDB5e1aB0520ad8a9049fa28"
	ownerPK      = "83f41cacb91895a4f0a57d18cd3e3f1513cdec49f8f6188ab6e626815ecfdf01"
	testChainId  = int64(97)
	amount       = big.NewInt(3358220000000000)

	/*rpc          = "https://goerli.infura.io/v3/cd0a83ff95ce4f78b73ad93d64be7bff"
	transferAddr = "0x63994BE0Ee17EBD189C7E72745155D10Ad859154"
	ownerPK      = "83f41cacb91895a4f0a57d18cd3e3f1513cdec49f8f6188ab6e626815ecfdf01"
	testChainId  = int64(5)
	amount       = big.NewInt(196000000000000000)*/
)

func TestTransfer(t *testing.T) {
	client, _ := ethclient.Dial(rpc)
	transferInstance, _ := batchTransfer.NewBatchTransfer(common.HexToAddress(transferAddr), client)
	// init owner
	ownerPrivateKey, _ := crypto.HexToECDSA(ownerPK)
	ownerAuth, _ := bind.NewKeyedTransactorWithChainID(ownerPrivateKey, big.NewInt(testChainId))
	to := []common.Address{
		common.HexToAddress("0xfF7e9Ef8DCaf380a40913F0984F5ecbA1e32390f"),
		common.HexToAddress("0xe64EA8cB1C8963782Fde17D7cdB63dDE89c6CfeC"),
		common.HexToAddress("0xFf3dc6932A4fC3930a3EC31130186EDBa45227FA"),
		common.HexToAddress("0xc2aE5D51b5E2999Dfe3aA9077DA0C6AD556AA1c2"),
		common.HexToAddress("0x702b298215784b873e03ba1Ad81f94E5eF845A27"),
		common.HexToAddress("0x8d9a824F098259d8d379515c049473E1D4DA145e"),
		common.HexToAddress("0xC3f32C2814c5Eb1F1B04117e9f7D9F31ecaE056a"),
		common.HexToAddress("0x969dbA8DdE32a59a992Cb2FBF3b287bd3BfBa619"),
		common.HexToAddress("0x85691243711d42147695D148c71235FfaC8eb2Db"),
		common.HexToAddress("0x2fCe338f0da77b9cAD041db109b81Ff07D59574e"),
	}
	tx, err := transferInstance.BatchTransferETH(ownerAuth, to, amount)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(tx.Hash())
}

func TestWithdraw(t *testing.T) {
	client, _ := ethclient.Dial(rpc)
	transferInstance, _ := batchTransfer.NewBatchTransfer(common.HexToAddress(transferAddr), client)
	// init owner
	ownerPrivateKey, _ := crypto.HexToECDSA(ownerPK)
	ownerAuth, _ := bind.NewKeyedTransactorWithChainID(ownerPrivateKey, big.NewInt(testChainId))
	tx, err := transferInstance.WithdrawETH(ownerAuth)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(tx.Hash())
}
