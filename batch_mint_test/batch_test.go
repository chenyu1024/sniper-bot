package batch_mint_test

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strings"
	"testing"
	"wxjkzcy/contract/freeMint"
)

var (
	amount          = 10
	mintTestChainId = int64(97)
	rpcTest         = "https://quiet-distinguished-tree.bsc-testnet.discover.quiknode.pro/a2bf3760742e089ff5e7168a4e0e7cbcb5b96c36/"
	MintContract    = common.HexToAddress("0xef4e44A358C47EdCC33AE76C8e48911005944EBc")
	NFTContract     = common.HexToAddress("0xab8197250666C8af2B09cd6347e15E6c9ad8cAe2")
	mintOwnerPK     = "83f41cacb91895a4f0a57d18cd3e3f1513cdec49f8f6188ab6e626815ecfdf01"
)

/*
1. js处理---0~200.txt
2. 批量转账
3. deploy nft 、 deploy mint，transfer to mint

*/

// 读取私钥
type Wallet struct {
	Address string `json:"address"`
	Pk      string `json:"pk"`
}
type Wallets struct {
	WalletSlice []Wallet `json:"wallets"`
}

func readJson(path string) []byte {
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer jsonFile.Close()
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return jsonData
}

// 生成私钥slice
func genPkSlice() (pkSlice []string) {
	var wsRead Wallets
	res := readJson("d:/wallet.json")
	_ = json.Unmarshal(res, &wsRead)
	pkSlice = make([]string, 0)
	for _, v := range wsRead.WalletSlice {
		//log.Println(v.Address)
		pkSlice = append(pkSlice, v.Pk)
	}
	return
}

// 读取Merkel tree并生成 Merkel slice
func genProofSlice() [][][32]byte {
	// n个地址所有的Merkel proof
	proofAll := make([][][32]byte, 0)

	for i := 0; i < amount; i++ {
		path := fmt.Sprintf("d:/proof/%d.txt", i)
		proofs := genSingleProof(path)
		proofAll = append(proofAll, proofs)
	}
	//log.Println(proofAll)
	return proofAll
}
func genSingleProof(path string) [][32]byte {
	var pkSlice = make([][32]byte, 0)
	file, err := os.Open(path)
	if err != nil {
		log.Println("open error")
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		str = strings.TrimSuffix(str, "\r\n")
		if err == io.EOF {
			log.Println("已经读到文件末尾")
			break
		}
		res := common.HexToHash(str)
		pkSlice = append(pkSlice, res)
	}
	return pkSlice
}

func TestBatchMint(t *testing.T) {
	pks := genPkSlice()
	pfs := genProofSlice()

	// init runba instance
	client, err := ethclient.Dial(rpcTest)
	if err != nil {
		t.Fatal(err.Error())
	}
	mintInstance, err := freeMint.NewFreeMint(MintContract, client)
	if err != nil {
		t.Fatal(err.Error())
	}
	for i := 0; i < 10; i++ {
		ownerPrivateKey, _ := crypto.HexToECDSA(pks[i])
		ownerRunbaAuth, _ := bind.NewKeyedTransactorWithChainID(ownerPrivateKey, big.NewInt(mintTestChainId))
		tx, err := mintInstance.FreeMint(ownerRunbaAuth, pfs[i])
		if err != nil {
			t.Log(err.Error())
		}
		t.Log(tx.Hash())
	}

}
