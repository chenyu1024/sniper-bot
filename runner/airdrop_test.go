package runner

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
	"testing"
)

func TestSetMerkleRoot(t *testing.T) {
	var a = AbiIns
	ownerAuth, err := bind.NewKeyedTransactorWithChainID(a.ownerPrivateKey, big.NewInt(testChainId))
	if err != nil {
		log.Fatal(err.Error())
	}
	var root = common.Hex2Bytes("0x470eecd4764aaeb53d8f558679c759ddf569389bd832215f7b48e439fb531d2c")
	root32 := byteTo32(root)
	m, err := a.airdropInstance.SetMerkleRoot(ownerAuth, root32)
	if err != nil {
		t.Fatal(err.Error())
	}
	txnIsOk(m.Hash())
	afterSet, err := a.airdropInstance.MerkleRoot(nil)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Logf("%s", afterSet)
}
func TestSetAirdropList(t *testing.T) {
	var a = AbiIns
	ownerAuth, err := bind.NewKeyedTransactorWithChainID(a.ownerPrivateKey, big.NewInt(testChainId))
	if err != nil {
		log.Fatal(err.Error())
	}
	var tokenIdList = []*big.Int{airdropTokenId1}
	addTx, err := a.airdropInstance.AddAirdropTokenIds(ownerAuth, tokenIdList)
	if err != nil {
		t.Fatal(err.Error())
	}
	txnIsOk(addTx.Hash())

}
func TestSetNftContractAddress(t *testing.T) {
	var a = AbiIns
	ownerAuth, err := bind.NewKeyedTransactorWithChainID(a.ownerPrivateKey, big.NewInt(testChainId))
	if err != nil {
		log.Fatal(err.Error())
	}
	tx, err := a.airdropInstance.SetNftAddress(ownerAuth, common.HexToAddress(nftAddress))
	if err != nil {
		t.Fatal(err.Error())
	}
	txnIsOk(tx.Hash())
	address, err := a.airdropInstance.RunXAddress(nil)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Logf("runContract: %s", address.String())
	if address.String() != nftAddress {
		t.Fatal("fail")
	}
}

func byteTo32(s []byte) [32]byte {
	b := [32]byte{}
	for index, v := range s {
		if index == 32 {
			break
		}
		b[index] = v
	}
	return b
}
