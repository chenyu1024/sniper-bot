package runner

import (
	"github.com/ethereum/go-ethereum/common"
	"time"
)

func (e *ethRunner) Start(chain string) {
	//path0 := common.HexToAddress("0xc778417E063141139Fce010982780140Aa0cD5Ab")
	//path1 := common.HexToAddress("0xaD6D458402F60fD3Bd25163575031ACDce07538D")
	//e.sBuy(chain, path0, path1, big.NewInt(100000))
	//e.buy(chain, path0, path1, big.NewInt(100000000000000), big.NewInt(10))

	// [start] 生产数据
	//initDB()

	//tokenIdList, tokenUriList := getFields(101, 200, "sneaker_test")
	//// [end] 生产数据
	//to := common.HexToAddress("")
	//log.Println("start mint")
	//e.batchMint(chain, to, tokenIdList, tokenUriList)
	//log.Println("end mint")

	//log.Println("start approve to airdrop")
	//tokenIdList, _ := getFields(9, 10, "sneaker_test")
	//e.approveFromLeader(chain, tokenIdList)
	//log.Println("end approve to airdrop")

	//log.Println("start add airdrop list")
	//tokenIdList, _ := getFields(1, 5, "sneaker_test")
	//e.addAirdropList(chain, tokenIdList)
	//log.Println("end add airdrop list")

	//log.Println("start add airdrop list")
	//tokenIdList, _ := getFields(9, 10, "sneaker_test")
	//e.batchAirdrop(chain, tokenIdList, []common.Address{
	//	common.HexToAddress("0x178d16d1dc5ed19d39f0b4afbb3dac622ae2293f"),
	//	common.HexToAddress("0x8a380b77c72f7ba9f09b776550a5b5446825d04f"),
	//})
	//log.Println("end add airdrop list")

	//log.Println("start open box")
	//boxList := []int64{11728149, 13533325, 14820016, 12358882, 12507878, 16819690, 19268340, 18585940, 15180473, 13132826, 16525372, 10593182, 17458850, 16707761, 17295576, 13559469, 18945636, 12213957, 13775372, 13726309, 15245243, 17219629, 16295476, 11940653, 13368348}
	//for _, u := range boxList {
	//	e.openBox(chain, big.NewInt(u))
	//	log.Println(u)
	//	time.Sleep(time.Second * 10)
	//}
	//log.Println("end open box")

	var addresses = []common.Address{
		common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"),
		common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92265"),
		common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92264"),
		common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92263"),
	}
	for i := 0; i < len(addresses); i++ {
		e.addLeader(chain, addresses[i])
		time.Sleep(time.Millisecond * 50)
	}
}
