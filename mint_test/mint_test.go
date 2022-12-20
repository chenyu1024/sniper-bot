package mint_test

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"log"
	"math/big"
	"testing"
	"wxjkzcy/contract/mint"
	nftAbi "wxjkzcy/contract/other"
)
import "wxjkzcy/runner"

var (
	rpcTestRunbaAddress = "https://data-seed-prebsc-2-s1.binance.org:8545/"

	mintContract             = "0x6E7a06fb6a9371b8157730A422D5c43aa67C9bC1"
	runbaTestContractAddress = "0x0F2d60576cddB843EFfb1b2859AAAdEC36078237"

	mintTestChainId         = int64(97)
	mintOwnerPK             = "31ec230279cfb9274bc27b68416d83b1aada47cda887f09546e1d567e6bcc1d2"
	mintContractOriginOwner = "0x4F69FC7bb16e4869FE51BF2E60d860b040a63282"

	ded2PK = "f3057d4dd422222a71ccf209101e4913e53b6be6f9107745221caa0c0b50bd63"
)

func TestAddSetting(t *testing.T) {
	runner.InitDB()
	// init runba instance
	client, err := ethclient.Dial(rpcTestRunbaAddress)
	mintInstance, err := mint.NewMint(common.HexToAddress(mintContract), client)
	if err != nil {
		log.Fatalln(err.Error())
	}
	// init owner
	ownerRunbaPrivateKey, _ := crypto.HexToECDSA(mintOwnerPK)
	// 【free mint】
	{
		// set free mint airdrop data
		ownerMintAuth, _ := bind.NewKeyedTransactorWithChainID(ownerRunbaPrivateKey, big.NewInt(mintTestChainId))
		tokenIds, box, shoe := runner.GetTestBoxAndShoe("C")
		log.Println(tokenIds)
		tx, err := mintInstance.AddFreeMintData(ownerMintAuth, tokenIds, box, shoe)
		if err != nil {
			log.Fatalln(err.Error())
		}
		log.Println(tx.Hash())
	}
	{
		// set free mint timestamp
		ownerMintAuth, _ := bind.NewKeyedTransactorWithChainID(ownerRunbaPrivateKey, big.NewInt(mintTestChainId))
		tx, err := mintInstance.SetFreeMintTime(ownerMintAuth, big.NewInt(1668596009), big.NewInt(1668682417))
		if err != nil {
			log.Fatalln(err.Error())
		}
		log.Println(tx.Hash())
	}
	{
		// set free mint merkle root
		freeMintMerkleTree := common.HexToHash("0xbfdf5a87cff87241b72cb5b4a5f6f0f5f0ffb1d2c12471fa2d844c51fd377ca2")
		ownerMintAuth, _ := bind.NewKeyedTransactorWithChainID(ownerRunbaPrivateKey, big.NewInt(mintTestChainId))
		tx, err := mintInstance.SetFMMerkleRoot(ownerMintAuth, freeMintMerkleTree)
		if err != nil {
			log.Fatalln(err.Error())
		}
		log.Println(tx.Hash())
	}
	// 【og mint】
	/*{
		// set og mint airdrop data
		ownerMintAuth, _ := bind.NewKeyedTransactorWithChainID(ownerRunbaPrivateKey, big.NewInt(mintTestChainId))
		tokenIds, box, shoe := runner.GetTestBoxAndShoe("B")
		log.Println(tokenIds)
		tx, err := mintInstance.AddOgMintData(ownerMintAuth, tokenIds, box, shoe)
		if err != nil {
			log.Fatalln(err.Error())
		}
		log.Println(tx.Hash())
	}
	{
		// set og mint timestamp
		ownerMintAuth, _ := bind.NewKeyedTransactorWithChainID(ownerRunbaPrivateKey, big.NewInt(mintTestChainId))
		tx, err := mintInstance.SetOGMintTime(ownerMintAuth, big.NewInt(1668596009), big.NewInt(1668682417))
		if err != nil {
			log.Fatalln(err.Error())
		}
		log.Println(tx.Hash())
	}
	{
		// set og mint merkle root
		ogMintMerkleTree := common.HexToHash("0xbfdf5a87cff87241b72cb5b4a5f6f0f5f0ffb1d2c12471fa2d844c51fd377ca2")
		ownerMintAuth, _ := bind.NewKeyedTransactorWithChainID(ownerRunbaPrivateKey, big.NewInt(mintTestChainId))
		tx, err := mintInstance.SetOGMerkleRoot(ownerMintAuth, ogMintMerkleTree)
		if err != nil {
			log.Fatalln(err.Error())
		}
		log.Println(tx.Hash())
	}*/
}
func TestFreeMintNft(t *testing.T) {
	// init runba instance
	client, err := ethclient.Dial(rpcTestRunbaAddress)
	mintInstance, err := mint.NewMint(common.HexToAddress(mintContract), client)
	if err != nil {
		log.Fatalln(err.Error())
	}
	// init owner
	ownerRunbaPrivateKey, _ := crypto.HexToECDSA(mintOwnerPK)
	ownerMintAuth, _ := bind.NewKeyedTransactorWithChainID(ownerRunbaPrivateKey, big.NewInt(mintTestChainId))
	//ownerMintAuth.Value = big.NewInt()
	merkleProof := [][32]byte{
		common.HexToHash("0x4aff374b14c44d97ac6da76e0da7dceff1c990bf3732833476b89b1957c10475"),
		common.HexToHash("0x6bf2c17399f57e4c811c027d093725416585a970655bcf17c5899154e33dc05b"),
		common.HexToHash("0xf3054ba188efdd1ffadc87256faa52d3e615b0430cfb4b48f72b323f826896c7"),
	}
	tx, err := mintInstance.FreeMint(ownerMintAuth, merkleProof)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println(tx.Hash())
}

func TestOGMintNft(t *testing.T) {
	// init runba instance
	client, err := ethclient.Dial(rpcTestRunbaAddress)
	mintInstance, err := mint.NewMint(common.HexToAddress(mintContract), client)
	if err != nil {
		log.Fatalln(err.Error())
	}
	// init owner
	ownerRunbaPrivateKey, _ := crypto.HexToECDSA(ded2PK)
	for i := 0; i < 4; i++ {
		//{
		ownerMintAuth, _ := bind.NewKeyedTransactorWithChainID(ownerRunbaPrivateKey, big.NewInt(mintTestChainId))

		amountIn, _ := big.NewFloat(0.1 * params.Ether).Int(nil)
		ownerMintAuth.Value = amountIn

		merkleProof := [][32]byte{
			common.HexToHash("0x773608c543c78bc15be1195542d01e8f072d416a48fc0c24167d3868c8a1699a"),
			common.HexToHash("0x6bf2c17399f57e4c811c027d093725416585a970655bcf17c5899154e33dc05b"),
			common.HexToHash("0xf3054ba188efdd1ffadc87256faa52d3e615b0430cfb4b48f72b323f826896c7"),
		}
		tx, err := mintInstance.OgMint(ownerMintAuth, merkleProof)
		if err != nil {
			log.Fatalln(err.Error())
		}
		log.Println(tx.Hash())
		//}
	}
}
func TestMintNft1(t *testing.T) {
	// init runba instance
	client, err := ethclient.Dial(rpcTestRunbaAddress)
	mintInstance, err := mint.NewMint(common.HexToAddress(mintContract), client)
	if err != nil {
		log.Fatalln(err.Error())
	}
	// init owner
	ownerRunbaPrivateKey, _ := crypto.HexToECDSA(ded2PK)
	ownerMintAuth, _ := bind.NewKeyedTransactorWithChainID(ownerRunbaPrivateKey, big.NewInt(mintTestChainId))
	//ownerMintAuth.Value = big.NewInt()
	merkleProof := [][32]byte{
		common.HexToHash("0x773608c543c78bc15be1195542d01e8f072d416a48fc0c24167d3868c8a1699a"),
		common.HexToHash("0x6bf2c17399f57e4c811c027d093725416585a970655bcf17c5899154e33dc05b"),
		common.HexToHash("0xf3054ba188efdd1ffadc87256faa52d3e615b0430cfb4b48f72b323f826896c7"),
	}
	tx, err := mintInstance.FreeMint(ownerMintAuth, merkleProof)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println(tx.Hash())
}

// nft 转移给 mint合约
func TestNewOwnerToMintContract(t *testing.T) {
	// init runba instance
	client, _ := ethclient.Dial(rpcTestRunbaAddress)
	runbaInstance, _ := nftAbi.NewAbi(common.HexToAddress(runbaTestContractAddress), client)
	// init owner
	ownerRunbaPrivateKey, _ := crypto.HexToECDSA(mintOwnerPK)
	ownerRunbaAuth, _ := bind.NewKeyedTransactorWithChainID(ownerRunbaPrivateKey, big.NewInt(mintTestChainId))
	transferOwn, err := runbaInstance.TransferOwnership(ownerRunbaAuth, common.HexToAddress(mintContract))
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println(transferOwn.Hash())
}
func TestTransferReset(t *testing.T) {
	// init runba instance
	client, err := ethclient.Dial(rpcTestRunbaAddress)
	mintInstance, err := mint.NewMint(common.HexToAddress(mintContract), client)
	if err != nil {
		log.Fatalln(err.Error())
	}
	// init owner
	ownerRunbaPrivateKey, _ := crypto.HexToECDSA(mintOwnerPK)
	ownerMintAuth, _ := bind.NewKeyedTransactorWithChainID(ownerRunbaPrivateKey, big.NewInt(mintTestChainId))
	// transfer
	tx, err := mintInstance.TransferToNewOwner(ownerMintAuth, common.HexToAddress(mintContractOriginOwner))
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println(tx.Hash())
	// 1.owner还没转移成功
	// 2. pos有错误
}
func TestWithdraw(t *testing.T) {
	// init runba instance
	client, err := ethclient.Dial(rpcTestRunbaAddress)
	mintInstance, err := mint.NewMint(common.HexToAddress(mintContract), client)
	if err != nil {
		log.Fatalln(err.Error())
	}
	// init owner
	ownerRunbaPrivateKey, _ := crypto.HexToECDSA(mintOwnerPK)
	ownerMintAuth, _ := bind.NewKeyedTransactorWithChainID(ownerRunbaPrivateKey, big.NewInt(mintTestChainId))
	tx, err := mintInstance.Withdraw(ownerMintAuth)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(tx.Hash())
}
