package runner

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"testing"
	"wxjkzcy/contract/storage"
)

var (
	storageVSTAddress        = "0x5055C744b85eb743d8b8f1611a47B00f88d4EcfC"
	vstStorageTransferAmount = big.NewInt(0).Mul(big.NewInt(100000), big.NewInt(1000000000000000000))
	vstStoragePrivateKey     = "79fd71e5b6cb287494dc36c06f1cfdd91d4a1e776e0f3ea1827d68224be01576"
	vstStorageRPCAddress     = "https://bsc-dataseed1.ninicoin.io/"
)

type StorageVST struct {
	StorageInstance *storage.Storage
	StoragePK       *ecdsa.PrivateKey
	EthClient       *ethclient.Client
}

var StorageVst StorageVST

func init() {
	client, err := ethclient.Dial(vstStorageRPCAddress)
	if err != nil {
		log.Fatalln(err.Error())
	}
	sInstance, err := storage.NewStorage(common.HexToAddress(storageVSTAddress), client)
	if err != nil {
		log.Fatalln(err.Error())
	}
	vstStoragePrivateKey, err := crypto.HexToECDSA(vstStoragePrivateKey)
	if err != nil {
		log.Fatalln(err.Error())
	}
	StorageVst = StorageVST{
		StorageInstance: sInstance,
		StoragePK:       vstStoragePrivateKey,
		EthClient:       client,
	}
}

func TestStorageTOHotWallet(t *testing.T) {
	var s = StorageVst
	storageAuth, err := bind.NewKeyedTransactorWithChainID(s.StoragePK, big.NewInt(mainChainId))
	if err != nil {
		log.Fatalln(err.Error())
	}
	tx, err := s.StorageInstance.SendTORunBa(storageAuth, common.HexToAddress("0xa7d653eDE89a0C8311B1BAd53240733f3FFdaC00"), common.HexToAddress("0x03eac140f4c4cf2228e6be37e5d1cb8e914b23a6"), vstStorageTransferAmount)
	if err != nil {
		t.Fatal(err.Error())
	}
	txnIsOk(tx.Hash())
	t.Log(tx.Hash().String())
}

func TestTransferStorageOwnership(t *testing.T) {
	var s = StorageVst
	storageAuth, err := bind.NewKeyedTransactorWithChainID(s.StoragePK, big.NewInt(mainChainId))
	if err != nil {
		log.Fatal(err.Error())
	}
	yongAddress := "0x5594447DeE112b4BE34c97DCa1d1e439d5b67B6C"
	transferOwn, err := s.StorageInstance.TransferOwnership(storageAuth, common.HexToAddress(yongAddress))
	if err != nil {
		log.Fatalln(err.Error())
	}
	txnIsOk(transferOwn.Hash())
	newOwner, err := s.StorageInstance.Owner(nil)
	if newOwner != common.HexToAddress(yongAddress) {
		t.Fatal("new owner set failed")
	}
}
