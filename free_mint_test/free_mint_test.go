package free_mint_test

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"testing"
	"wxjkzcy/contract/freeMint"
	"wxjkzcy/contract/runbaTest"
	"wxjkzcy/runner"
)

var (
	rpcTestRunbaAddress                = "https://quiet-distinguished-tree.bsc-testnet.discover.quiknode.pro/a2bf3760742e089ff5e7168a4e0e7cbcb5b96c36/"
	mintContract                       = "0xef4e44A358C47EdCC33AE76C8e48911005944EBc"
	mintOwnerPK                        = "83f41cacb91895a4f0a57d18cd3e3f1513cdec49f8f6188ab6e626815ecfdf01"
	mintTestChainId                    = int64(97)
	freeMintStartTime, freeMintEndTime = big.NewInt(1671501166), big.NewInt(1671546646)
	freeMintMerkleRoot                 = common.HexToHash("0xcc7fbac43ed9422fcc1c9832f54f21c6c9b5f0f81dc8b308305b986bffef2b61")

	runbaNftAddress = "0xab8197250666C8af2B09cd6347e15E6c9ad8cAe2"
	runbaOwnerPK    = "83f41cacb91895a4f0a57d18cd3e3f1513cdec49f8f6188ab6e626815ecfdf01"
)

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}
func TestInitSetting(t *testing.T) {
	runner.InitDB()
	client, err := ethclient.Dial(rpcTestRunbaAddress)
	mintInstance, err := freeMint.NewFreeMint(common.HexToAddress(mintContract), client)
	checkErr(err)
	ownerMintPrivateKey, _ := crypto.HexToECDSA(mintOwnerPK)
	/*{
		// set free mint start&end time
		ownerMintAuth, _ := bind.NewKeyedTransactorWithChainID(ownerMintPrivateKey, big.NewInt(mintTestChainId))
		tx, err := mintInstance.SetFreeMintTime(ownerMintAuth, freeMintStartTime, freeMintEndTime)
		checkErr(err)
		log.Println("set freeMint Time", tx.Hash())
	}*/
	{
		// set merkel root
		ownerMintAuth, _ := bind.NewKeyedTransactorWithChainID(ownerMintPrivateKey, big.NewInt(mintTestChainId))
		tx, err := mintInstance.SetMerkleRoot(ownerMintAuth, freeMintMerkleRoot)
		checkErr(err)
		log.Println("set merkel root", tx.Hash())
	}
	// ******** warning ********* //
	//{
	//	// add free mint data
	//	tokenIds, boxes, shoes := runner.GetFreeMintData()
	//	//log.Fatalln(boxes, shoes)
	//	ownerMintAuth, _ := bind.NewKeyedTransactorWithChainID(ownerMintPrivateKey, big.NewInt(mintTestChainId))
	//	tx, err := mintInstance.AddFreeMintData(ownerMintAuth, tokenIds, boxes, shoes)
	//	checkErr(err)
	//	log.Println("add free mint data", tx.Hash())
	//}
}
func TestOwnerRunbaToMint(t *testing.T) {
	ownerRunbaPrivateKey, _ := crypto.HexToECDSA(runbaOwnerPK)
	ownerRunbaAuth, _ := bind.NewKeyedTransactorWithChainID(ownerRunbaPrivateKey, big.NewInt(mintTestChainId))
	// init runba instance
	client, err := ethclient.Dial(rpcTestRunbaAddress)
	checkErr(err)
	runbaInstance, err := runbaTest.NewRunbaTest(common.HexToAddress(runbaNftAddress), client)
	checkErr(err)
	to := common.HexToAddress("0xef4e44A358C47EdCC33AE76C8e48911005944EBc")
	//owner1, _ := runbaInstance.Owner(nil)
	//log.Fatalln(owner1)

	tx, err := runbaInstance.TransferOwnership(ownerRunbaAuth, to)
	checkErr(err)
	log.Println("transfer ownership", tx.Hash())
}

func TestFreeMint(t *testing.T) {
	client, err := ethclient.Dial(rpcTestRunbaAddress)
	mintInstance, err := freeMint.NewFreeMint(common.HexToAddress(mintContract), client)
	checkErr(err)
	ownerMintPrivateKey, _ := crypto.HexToECDSA(mintOwnerPK)
	ownerMintAuth, _ := bind.NewKeyedTransactorWithChainID(ownerMintPrivateKey, big.NewInt(mintTestChainId))
	merkleProof := [][32]byte{
		common.HexToHash("0x0a79d27e4168da6925d23297fa79e2fcf97abbe477d9db7e7c4dddb90ccfd83f"),
		common.HexToHash("0x9985727bab1fccc97e5401f0517ed720f761b2e872dcb216014f02e5295ccb20"),
	}
	tx, err := mintInstance.FreeMint(ownerMintAuth, merkleProof)
	checkErr(err)
	log.Println("mint success", tx.Hash())
}
func TestFilp(t *testing.T) {
	client, err := ethclient.Dial(rpcTestRunbaAddress)
	mintInstance, err := freeMint.NewFreeMint(common.HexToAddress(mintContract), client)
	checkErr(err)
	ownerMintPrivateKey, _ := crypto.HexToECDSA(mintOwnerPK)
	ownerMintAuth, _ := bind.NewKeyedTransactorWithChainID(ownerMintPrivateKey, big.NewInt(mintTestChainId))
	tx, err := mintInstance.FlipSaleState(ownerMintAuth)
	checkErr(err)
	log.Println("mint success", tx.Hash())
}
func TestFPOS(t *testing.T) {
	client, err := ethclient.Dial(rpcTestRunbaAddress)
	mintInstance, err := freeMint.NewFreeMint(common.HexToAddress(mintContract), client)
	checkErr(err)
	//ownerMintPrivateKey, _ := crypto.HexToECDSA(mintOwnerPK)
	//ownerMintAuth, _ := bind.NewKeyedTransactorWithChainID(ownerMintPrivateKey, big.NewInt(mintTestChainId))
	tx, err := mintInstance.FMPos(nil)
	checkErr(err)
	log.Println("mint success", tx)
}
func TestResetOwnership(t *testing.T) {
	client, err := ethclient.Dial(rpcTestRunbaAddress)
	mintInstance, err := freeMint.NewFreeMint(common.HexToAddress(mintContract), client)
	checkErr(err)
	ownerMintPrivateKey, _ := crypto.HexToECDSA(mintOwnerPK)
	ownerMintAuth, _ := bind.NewKeyedTransactorWithChainID(ownerMintPrivateKey, big.NewInt(mintTestChainId))
	tx, err := mintInstance.ResetOwnership(ownerMintAuth, common.HexToAddress("0x59321acA9557E3d35Ab12F8B249e4f7A43DB5Eb3"))
	checkErr(err)
	log.Println(tx.Hash())
}
func TestCheckMetadata(t *testing.T) {
	client, err := ethclient.Dial("https://bsc.nodereal.io")
	checkErr(err)
	runbaInstance, err := runbaTest.NewRunbaTest(common.HexToAddress("0xfa6998b51f9f91580c92359af2d7644319584717"), client)
	checkErr(err)
	//str, err := runbaInstance.TokenURI(nil, big.NewInt(17997724))
	str, err := runbaInstance.IsOpen(nil, big.NewInt(17997724))
	checkErr(err)
	log.Println(str)
}
