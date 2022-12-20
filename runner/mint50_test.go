package runner

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"testing"
	nftAbi "wxjkzcy/contract/other"
	"wxjkzcy/contract/runbaTest"
)

var (
	//runbaTestChainId         = 97
	//runbaOwner               = "0x54F1F088a217378D8D9464BDc8Fa190cda83B24a"
	//runbaPOwnerPK            = "5b51e207c26d1a8b16a59ee6149ae2b51418f3e15da0d29461bc16b908a82ebd"
	//testRunbaLeaderAddress   = "0x54F1F088a217378D8D9464BDc8Fa190cda83B24a"
	//testRunbaLeaderPK        = "5b51e207c26d1a8b16a59ee6149ae2b51418f3e15da0d29461bc16b908a82ebd"
	//rpcTestRunbaAddress      = "https://data-seed-prebsc-2-s1.binance.org:8545/"
	//runbaTestContractAddress = "0x167ABc73f1AD06994799801c03C2F105477D15a8"

	runbaTestChainId         = 56
	runbaOwner               = "0x59321acA9557E3d35Ab12F8B249e4f7A43DB5Eb3"
	runbaPOwnerPK            = "79fd71e5b6cb287494dc36c06f1cfdd91d4a1e776e0f3ea1827d68224be01576"
	testRunbaLeaderAddress   = "0x03eac140f4c4cf2228e6be37e5d1cb8e914b23a6"
	testRunbaLeaderPK        = ""
	rpcTestRunbaAddress      = "https://bsc.nodereal.io"
	runbaTestContractAddress = "0xfa6998b51F9f91580C92359aF2D7644319584717"
)

func TestMint50Test(t *testing.T) {
	InitDB()
	tokenIdListAll, jsonUriListAll, boxUriListAll := getBoxListByQuality("", "", 0)
	// init owner
	ownerRunbaPrivateKey, _ := crypto.HexToECDSA(runbaPOwnerPK)
	ownerRunbaAuth, _ := bind.NewKeyedTransactorWithChainID(ownerRunbaPrivateKey, big.NewInt(int64(runbaTestChainId)))
	// init runba instance
	client, err := ethclient.Dial(rpcTestRunbaAddress)
	//runbaInstance, err := nftAbi.NewAbi(common.HexToAddress(runbaTestContractAddress), client) // 主网
	runbaInstance, err := runbaTest.NewRunbaTest(common.HexToAddress(runbaTestContractAddress), client) // 测试网
	if err != nil {
		log.Fatalln(err.Error())
	}
	to := common.HexToAddress(testRunbaLeaderAddress)
	tx, err := runbaInstance.MintBatch200(ownerRunbaAuth, to, tokenIdListAll, jsonUriListAll, boxUriListAll) // 主网
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(tx.Hash())
	//txnIsOk(tx.Hash())
}
func TestGetOwnerByTokenId(t *testing.T) {
	// init runba instance
	client, _ := ethclient.Dial(rpcTestRunbaAddress)
	runbaInstance, _ := runbaTest.NewRunbaTest(common.HexToAddress(runbaTestContractAddress), client)
	owner, _ := runbaInstance.OwnerOf(nil, big.NewInt(14528294))
	log.Println(owner)
}
func TestBatchApproveTest(t *testing.T) {
	InitDB()
	// init runba instance
	client, err := ethclient.Dial(rpcTestRunbaAddress)
	runbaInstance, err := runbaTest.NewRunbaTest(common.HexToAddress(runbaTestContractAddress), client)
	if err != nil {
		log.Fatalln(err.Error())
	}
	tokenIdListAll, _, _ := getBoxListByQuality("S", "sneaker_info_main", 2)
	// init leader
	leaderRunbaPrivateKey, _ := crypto.HexToECDSA(testRunbaLeaderPK)
	leaderRunbaAuth, _ := bind.NewKeyedTransactorWithChainID(leaderRunbaPrivateKey, big.NewInt(int64(runbaTestChainId)))
	// start approve to runbaNftOwner
	tx, err := runbaInstance.BatchApprove(leaderRunbaAuth, tokenIdListAll, common.HexToAddress(runbaOwner))
	if err != nil {
		log.Fatalln(err.Error())
	}
	txnIsOk(tx.Hash())
}
func TestAirdropTest(t *testing.T) {
	InitDB()
	// init runba instance
	client, _ := ethclient.Dial(rpcTestRunbaAddress)
	runbaInstance, _ := runbaTest.NewRunbaTest(common.HexToAddress(runbaTestContractAddress), client)
	// init owner
	ownerRunbaPrivateKey, _ := crypto.HexToECDSA(runbaPOwnerPK)
	ownerRunbaAuth, _ := bind.NewKeyedTransactorWithChainID(ownerRunbaPrivateKey, big.NewInt(int64(runbaTestChainId)))
	// make tokenId list
	tokenIdListAll, _, _ := getBoxListByQuality("S", "sneaker_info_main", 2)
	// make airdrop list
	toCorrectList := make([]common.Address, 0)
	for i := 0; i < len(tokenIdListAll); i++ {
		toCorrectList = append(toCorrectList, common.HexToAddress("hyperPayAddressTest"))
	}
	for _, v := range toCorrectList {
		// 验证list
		if v != common.HexToAddress("hyperPayAddressTest") {
			log.Fatalln("error")
		}
	}
	// start airdrop
	airdropTx, err := runbaInstance.Airdrop(ownerRunbaAuth, tokenIdListAll, toCorrectList)
	if err != nil {
		log.Fatalln(err.Error())
	}
	txnIsOk(airdropTx.Hash())
}

func TestNewOwnerToYong(t *testing.T) {
	yongAddress := "0x5594447DeE112b4BE34c97DCa1d1e439d5b67B6C"
	// init runba instance
	client, _ := ethclient.Dial(rpcTestRunbaAddress)
	runbaInstance, _ := nftAbi.NewAbi(common.HexToAddress(runbaTestContractAddress), client)
	// init owner
	ownerRunbaPrivateKey, _ := crypto.HexToECDSA(runbaPOwnerPK)
	ownerRunbaAuth, _ := bind.NewKeyedTransactorWithChainID(ownerRunbaPrivateKey, big.NewInt(int64(runbaTestChainId)))
	transferOwn, err := runbaInstance.TransferOwnership(ownerRunbaAuth, common.HexToAddress(yongAddress))
	if err != nil {
		log.Fatalln(err.Error())
	}

	txnIsOk(transferOwn.Hash())
	newOwner, err := runbaInstance.Owner(nil)
	if newOwner != common.HexToAddress(yongAddress) {
		t.Fatal("new owner set failed")
	}
}
func TestUpdateMintStatusTwice(t *testing.T) {
	InitDB()
	tokenIdList, _, _ := getBoxListByQuality("", "", 1)
	for _, v := range tokenIdList {
		id := updateIsMint("sneaker_info_main_copy1", 1, v.Int64())
		t.Log(id)
	}
}
