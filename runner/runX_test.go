package runner

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
	airdropAbi "wxjkzcy/contract/airdrop"
	nftAbi "wxjkzcy/contract/other"
)

var (
	//leaderAddress     = "0xf1d7ad7b3ee5929d2b44dbfd452b7791bcb0c213" // 项目方钱包
	leaderAddress     = "0x03eac140f4c4cf2228e6be37e5d1cb8e914b23a6" // 项目方钱包
	testLeaderAddress = "0x54F1F088a217378D8D9464BDc8Fa190cda83B24a" // 项目方钱包
	leaderPrivateKey  = "a428326eb2ea0c16edbbe552f4e4350abf9dbaa5bbf2c2d5ff4b3d909c6a0ee7"
	ownerAddress      = "0x59321acA9557E3d35Ab12F8B249e4f7A43DB5Eb3" // owner钱包地址
	//ownerPrivateKey   = "79fd71e5b6cb287494dc36c06f1cfdd91d4a1e776e0f3ea1827d68224be01576"
	ownerPrivateKey   = "4a0fb9b76a1155caa0b9dd0653f7ecb724421c8d251fde3d6dc7a6af67c8b649" //test owner
	newOwnerAddress   = "0x89b2a39e930E49f587F8F428177411E2241Fa4C9"
	newOwnerAddressPk = "4a0fb9b76a1155caa0b9dd0653f7ecb724421c8d251fde3d6dc7a6af67c8b649"
	airdropAddress    = "0x79e3Da0fb21A50947d5AB203Ab7E5F522eDcac43" // 空投合约地址
	nftAddress        = "0xf7e3e2196b759eF67BB04F5c893A1ab46CD70F5c" // test RunX
	//nftAddress = "0xfa6998b51F9f91580C92359aF2D7644319584717" // RunX
	//rpcAddress = "https://bsc-dataseed1.defibit.io/"
	rpcAddress    = "https://data-seed-prebsc-1-s1.binance.org:8545/"
	headerUri     = "https://runba-global.s3.ap-southeast-1.amazonaws.com/on-chain-asset-json/"
	headerUriTest = "https://test.czymln.com/data/"
	footUri       = ".json"
	reserveJson   = "https://www.runba.xyz/"
	testChainId   = int64(97)
	mainChainId   = int64(56)
	playerAddress = "0xD22752Cb93751d6d732eE751448C45c72C37DeD2" // 玩家地址
	tokenId1      = big.NewInt(14155288)
	jsonUri1      = "8a24c024f"
	boxUri1       = "479738c2a"
	tokenId2      = big.NewInt(18888888)
	jsonUri2      = "00fcfe9ee"
	boxUri2       = "00fcfe9ec"
	//tokenId3          = big.NewInt(19581636)
	//jsonUri3          = "7824729f4"
	//boxUri3           = "be85754b2"
	airdropTokenId1 = big.NewInt(17792347)
	//airdropTokenId2 = big.NewInt(18697063)
	//airdropTokenId3 = big.NewInt(15100058)
	hyperPayAddress = "0x99Ee8E4e3061Cc767a9d4965De14a0e194958F82"
)

type AbiInstances struct {
	nftInstance      *nftAbi.Abi
	airdropInstance  *airdropAbi.Airdrop
	ownerPrivateKey  *ecdsa.PrivateKey
	leaderPrivateKey *ecdsa.PrivateKey
	ethClient        *ethclient.Client
}

var AbiIns AbiInstances

// 初始化
func init() {
	InitDB()
	client, err := ethclient.Dial(rpcAddress)
	if err != nil {
		log.Fatalln(err.Error())
	}
	nftInstance, err := nftAbi.NewAbi(common.HexToAddress(nftAddress), client)
	if err != nil {
		log.Fatalln(err.Error())
	}
	airdropInstance, err := airdropAbi.NewAirdrop(common.HexToAddress(airdropAddress), client)
	if err != nil {
		log.Fatalln(err.Error())
	}
	ownerPrivateKey, err := crypto.HexToECDSA(ownerPrivateKey)
	if err != nil {
		log.Fatalln(err.Error())
	}
	leaderPrivateKey, err := crypto.HexToECDSA(leaderPrivateKey)
	if err != nil {
		log.Fatalln(err.Error())
	}
	AbiIns = AbiInstances{
		nftInstance:      nftInstance,
		airdropInstance:  airdropInstance,
		ownerPrivateKey:  ownerPrivateKey,
		leaderPrivateKey: leaderPrivateKey,
		ethClient:        client,
	}
}

// 查询订单状态
func txnIsOk(hash common.Hash) bool {
	for i := 1; ; i++ {
		receipt, _ := AbiIns.ethClient.TransactionReceipt(context.Background(), hash)
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

func TestAddLeader(t *testing.T) {
	var a = AbiIns
	ownerAuth, err := bind.NewKeyedTransactorWithChainID(a.ownerPrivateKey, big.NewInt(mainChainId))
	if err != nil {
		log.Fatal(err.Error())
	}
	beforeAddStatus, _ := a.nftInstance.IsLeader(nil, common.HexToAddress(leaderAddress))
	var leaderList = []common.Address{common.HexToAddress(leaderAddress)}
	addTx, err := a.nftInstance.AddLeader(ownerAuth, leaderList)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log("tx hash:", addTx.Hash())
	txnIsOk(addTx.Hash())
	afterAddStatus, err := a.nftInstance.IsLeader(nil, common.HexToAddress(leaderAddress))
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Logf("before add leader: %v", beforeAddStatus)
	t.Logf("after add leader: %v", afterAddStatus)
	if !beforeAddStatus != afterAddStatus {
		t.Fatal("设置项目方失败")
	}
}

func TestTransferOwnership(t *testing.T) {
	var a = AbiIns
	ownerAuth, err := bind.NewKeyedTransactorWithChainID(a.ownerPrivateKey, big.NewInt(mainChainId))
	if err != nil {
		log.Fatal(err.Error())
	}
	yongAddress := "0x5594447DeE112b4BE34c97DCa1d1e439d5b67B6C"
	transferOwn, err := a.nftInstance.TransferOwnership(ownerAuth, common.HexToAddress(yongAddress))
	if err != nil {
		log.Fatalln(err.Error())
	}
	txnIsOk(transferOwn.Hash())
	newOwner, err := a.nftInstance.Owner(nil)
	if newOwner != common.HexToAddress(yongAddress) {
		t.Fatal("new owner set failed")
	}
}

//铸造单个鞋盒
//func TestMint(t *testing.T) {
//	var a = AbiIns
//	ownerAuth, err := bind.NewKeyedTransactorWithChainID(a.ownerPrivateKey, big.NewInt(testChainId))
//	if err != nil {
//		log.Fatal(err.Error())
//	}
//	// owner发起mint操作给项目方
//	to := common.HexToAddress(playerAddress)
//	tx, err := a.nftInstance.Mint(ownerAuth, to, tokenId2, jsonUri2, boxUri2)
//	if err != nil {
//		log.Fatal(err.Error())
//	}
//	t.Log("mint hash ", tx.Hash().String())
//	// wait for tx finished
//	txnIsOk(tx.Hash())
//	// 查询tokenId所有者地址
//	tokenOwnerAddress, err := a.nftInstance.OwnerOf(nil, tokenId1)
//	if err != nil {
//		log.Fatalln(err.Error())
//	}
//	t.Logf("to address:%s", to.String())
//	t.Logf("tokenId:%s,owner:%s", tokenId1.String(), tokenOwnerAddress.String())
//	if tokenOwnerAddress == to {
//		t.Log("铸造完成")
//	}
//}
func TestBatchMint(t *testing.T) {
	// [start] 生产数据
	InitDB()
	tokenIdList, jsonUriList, boxUriList := getFields(701, 900, "sneaker_info_main_copy1")
	// [end] 生产数据

	var a = AbiIns
	ownerAuth, err := bind.NewKeyedTransactorWithChainID(a.ownerPrivateKey, big.NewInt(testChainId))
	if err != nil {
		log.Fatal(err.Error())
	}
	ownerAuth.GasPrice = big.NewInt(10 * params.GWei)
	// owner发起mint操作给项目方
	to := common.HexToAddress(leaderAddress)

	tx, err := a.nftInstance.MintBatch200(ownerAuth, to, tokenIdList, jsonUriList, boxUriList)
	if err != nil {
		log.Fatal(err.Error())
	}
	t.Log("mint hash:", tx.Hash())
	txnIsOk(tx.Hash())
	//tokenOwner2, _ := a.nftInstance.OwnerOf(nil, tokenId2)
	//tokenOwner3, _ := a.nftInstance.OwnerOf(nil, tokenId3)
	//t.Logf("mint to address :%s", to.String())
	//t.Logf("tokenId:%s Owner:%s", tokenId2.String(), tokenOwner2.String())
	//t.Logf("tokenId:%s Owner:%s", tokenId3.String(), tokenOwner3.String())
	//if tokenOwner2 != tokenOwner3 {
	//	t.Log("fail")
	//}
}

// 测试空投
func TestAirdrop(t *testing.T) {
	wrongPK := "79fd71e5b6cb287494dc36c06f1cfdd91d4a1e776e0f3ea1827d68224be01576"
	ownerPrivateKey, err := crypto.HexToECDSA(wrongPK)
	if err != nil {
		log.Fatalln(err.Error())
	}
	var a = AbiIns
	//ownerAuth, err := bind.NewKeyedTransactorWithChainID(ownerPrivateKey, big.NewInt(mainChainId))
	ownerAuth, err := bind.NewKeyedTransactorWithChainID(ownerPrivateKey, big.NewInt(testChainId))
	if err != nil {
		log.Fatal(err.Error())
	}
	// 授权
	//leaderAuth, err := bind.NewKeyedTransactorWithChainID(a.leaderPrivateKey, big.NewInt(testChainId))
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//var tokenIdList = []*big.Int{tokenId1}
	//var toList = []common.Address{common.HexToAddress(playerAddress)}

	// 项目方授权给owner
	//txApprove, err := a.nftInstance.BatchApprove(leaderAuth, tokenIdList, common.HexToAddress(newOwnerAddress))
	//if err != nil {
	//	log.Fatalln(err.Error())
	//}
	//txnIsOk(txApprove.Hash())
	//t.Log("授权完成")
	// 合约 空投给玩家
	tokenIdList, _, _ := getBoxListByQuality("", "", 1)
	toCorrectList := make([]common.Address, 0)
	for i := 0; i < len(tokenIdList); i++ {
		toCorrectList = append(toCorrectList, common.HexToAddress(hyperPayAddress))
	}
	// 验证list
	for _, v := range toCorrectList {
		if v != common.HexToAddress(hyperPayAddress) {
			log.Fatalln("sd")
		}
	}
	airdropTx, err := a.nftInstance.Airdrop(ownerAuth, tokenIdList, toCorrectList)
	if err != nil {
		log.Fatalln(err.Error())
	}
	txnIsOk(airdropTx.Hash())
	//tokenOwnerAddress1, _ := a.nftInstance.OwnerOf(nil, tokenId1)
	//tokenOwnerAddress2, _ := a.nftInstance.OwnerOf(nil, tokenId2)
	//tokenOwnerAddress3, _ := a.nftInstance.OwnerOf(nil, tokenId3)
	//log.Printf("tokenId:%s,owner:%s", tokenId1.String(), tokenOwnerAddress1.String())
	//log.Printf("tokenId:%s,owner:%s", tokenId2.String(), tokenOwnerAddress2.String())
	//log.Printf("tokenId:%s,owner:%s", tokenId3.String(), tokenOwnerAddress3.String())
	//if toList[0] != tokenOwnerAddress1 ||
	//	toList[1] != tokenOwnerAddress2 ||
	//	toList[2] != tokenOwnerAddress3 {
	//	t.Fatal("空投失败")
	//}
}

// 开盒
func TestOpenBox(t *testing.T) {
	var a = AbiIns
	leaderAuth, err := bind.NewKeyedTransactorWithChainID(a.leaderPrivateKey, big.NewInt(testChainId))
	if err != nil {
		log.Fatal(err.Error())
	}
	beforeOpen, err := a.nftInstance.IsOpen(nil, tokenId1)
	if err != nil {
		log.Fatalln(err.Error())
	}
	openBoxTx, err := a.nftInstance.OpenBox(leaderAuth, tokenId1)
	if err != nil {
		log.Fatalln(err.Error())
	}
	txnIsOk(openBoxTx.Hash())
	afterOpen, err := a.nftInstance.IsOpen(nil, tokenId1)
	if err != nil {
		log.Fatalln(err.Error())
	}
	t.Logf("before open status: %v", beforeOpen)
	t.Logf("after open status: %v", afterOpen)
	if !beforeOpen && !afterOpen {
		t.Fatal("failed")
	}
}

func TestSetHeaderUri(t *testing.T) {
	var a = AbiIns
	//ownerAuth, err := bind.NewKeyedTransactorWithChainID(a.ownerPrivateKey, big.NewInt(testChainId))
	ownerAuth, err := bind.NewKeyedTransactorWithChainID(a.ownerPrivateKey, big.NewInt(mainChainId))
	if err != nil {
		log.Fatal(err.Error())
	}
	txn, err := a.nftInstance.SetHeaderUri(ownerAuth, headerUri)
	if err != nil {
		log.Fatal(err.Error())
	}
	txnIsOk(txn.Hash())
	//newHeader, err := a.nftInstance.HeaderUri(nil)
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//if newHeader != headerUri {
	//	t.Fatal("header设置失败")
	//}

}

func TestSetFootUri(t *testing.T) {
	var a = AbiIns
	//ownerAuth, err := bind.NewKeyedTransactorWithChainID(a.ownerPrivateKey, big.NewInt(testChainId))
	ownerAuth, err := bind.NewKeyedTransactorWithChainID(a.ownerPrivateKey, big.NewInt(mainChainId))
	if err != nil {
		log.Fatal(err.Error())
	}
	txn, err := a.nftInstance.SetFootUri(ownerAuth, footUri)
	if err != nil {
		log.Fatal(err.Error())
	}
	txnIsOk(txn.Hash())
	//newFoot, err := a.nftInstance.FootUri(nil)
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//if newFoot != footUri {
	//	t.Fatal("foot设置失败")
	//}
}

func TestSetReserveJson(t *testing.T) {
	var a = AbiIns
	//ownerAuth, err := bind.NewKeyedTransactorWithChainID(a.ownerPrivateKey, big.NewInt(testChainId))
	ownerAuth, err := bind.NewKeyedTransactorWithChainID(a.ownerPrivateKey, big.NewInt(mainChainId))
	if err != nil {
		log.Fatal(err.Error())
	}
	txn, err := a.nftInstance.SetDefaultUri(ownerAuth, reserveJson)
	if err != nil {
		log.Fatal(err.Error())
	}
	txnIsOk(txn.Hash())
	//newReserveJson, err := a.nftInstance.SetDefaultUri(nil)
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//if newReserveJson != reserveJson {
	//	t.Fatal("reserveJson设置失败")
	//}
}

//func TestBatchMint(t *testing.T) {
//	// [start] 生产数据
//	initDB()
//	tokenIdList, tokenUriList := getFields(101, 200, "sneaker_test_10000")
//	// [end] 生产数据
//	to := common.HexToAddress("")
//	log.Println("start mint")
//	e.batchMint(chain, to, tokenIdList, tokenUriList)
//	log.Println("end mint")
//}
func TestGetBoxListByQualityAndMint200(t *testing.T) {
	// 4a0fb9b76a1155caa0b9dd0653f7ecb724421c8d251fde3d6dc7a6af67c8b649 (4c9)
	// 5b51e207c26d1a8b16a59ee6149ae2b51418f3e15da0d29461bc16b908a82ebd (24a)
	//tokenIdListC, jsonUriListC, boxUriListC := getBoxListByQuality("C", "sneaker_info_main", 23)
	//log.Println(tokenIdListC)
	//log.Println(jsonUriListC)
	//log.Println(boxUriListC)

	//tokenIdListB, jsonUriListB, boxUriListB := getBoxListByQuality("B", "sneaker_info_main", 20)
	//log.Println(tokenIdListB)
	//log.Println(jsonUriListB)
	//log.Println(boxUriListB)
	//tokenIdListA, jsonUriListA, boxUriListA := getBoxListByQuality("A", "sneaker_info_main", 5)
	//log.Println(tokenIdListA)
	//log.Println(jsonUriListA)
	//log.Println(boxUriListA)
	//tokenIdListS, jsonUriListS, boxUriListS := getBoxListByQuality("S", "sneaker_info_main", 2)
	tokenIdListAll, jsonUriListAll, boxUriListAll := getBoxListByQuality("S", "sneaker_info_main", 2)
	//log.Println(tokenIdListS)
	//log.Println(jsonUriListS)
	//log.Println(boxUriListS)
	var a = AbiIns
	//ownerAuth, err := bind.NewKeyedTransactorWithChainID(a.ownerPrivateKey, big.NewInt(mainChainId))
	ownerAuth, err := bind.NewKeyedTransactorWithChainID(a.ownerPrivateKey, big.NewInt(testChainId))
	if err != nil {
		log.Fatal(err.Error())
	}
	// owner发起mint操作给项目方
	//to := common.HexToAddress(leaderAddress)
	to := common.HexToAddress(testLeaderAddress)

	tx, err := a.nftInstance.MintBatch200(ownerAuth, to, tokenIdListAll, jsonUriListAll, boxUriListAll)
	if err != nil {
		log.Fatal(err.Error())
	}
	t.Log("mint hash:", tx.Hash())
	txnIsOk(tx.Hash())
}

func TestBatchApprove(t *testing.T) {
	wrongPK := "a428326eb2ea0c16edbbe552f4e4350abf9dbaa5bbf2c2d5ff4b3d909c6a0ee7"
	ownerPrivateKey, err := crypto.HexToECDSA(wrongPK)
	if err != nil {
		log.Fatalln(err.Error())
	}

	var a = AbiIns
	wrongAuth, err := bind.NewKeyedTransactorWithChainID(ownerPrivateKey, big.NewInt(mainChainId))
	tokenIdList, _, _ := getBoxListByQuality("", "", 1)
	tx, err := a.nftInstance.BatchApprove(wrongAuth, tokenIdList, common.HexToAddress(ownerAddress))
	if err != nil {
		log.Fatalln(err.Error())
	}
	txnIsOk(tx.Hash())
	log.Println(tx.Hash())
}

//func TestVerifyIds(t *testing.T) {
//	initDB()
//	res := getJsonData("d:/duiying.json")
//	for _, v := range res {
//		dbIdByZcJson := getFieldValue("sneaker_info_main_copy1", v.JsonUri)
//		zcJsonId := v.Id
//		log.Println(dbIdByZcJson, zcJsonId)
//		if !strings.EqualFold(dbIdByZcJson, zcJsonId) {
//			log.Fatalln("数据不一致")
//		}
//	}
//}
func TestAfterTransfer(t *testing.T) {
	var a = AbiIns
	tokenIdList, _, _ := getBoxListByQuality("", "", 1)
	for _, v := range tokenIdList {
		owner, err := a.nftInstance.OwnerOf(nil, v)
		if err != nil {
			log.Fatalln(err.Error())
		}
		if owner != common.HexToAddress("0x03eac140f4c4cf2228e6be37e5d1cb8e914b23a6") {
			log.Fatalln("error owner")
		}
	}
}

// TODO ZCY
func TestUpdateMintStatus(t *testing.T) {
	InitDB()
	tokenIdList, _, _ := getBoxListByQuality("", "", 1)
	for _, v := range tokenIdList {
		id := updateIsMint("sneaker_info_main_copy1", 1, v.Int64())
		t.Log(id)
	}
}
