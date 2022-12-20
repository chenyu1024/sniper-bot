package cdnft_test

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"testing"
	"wxjkzcy/contract/cdnft"
)

var (
	rpc         = "https://data-seed-prebsc-1-s3.binance.org:8545"
	cdNFTAddr   = "0xB81912298ebe49Aadadafe4E51e204c742F77d9e"
	ownerPK     = "5b51e207c26d1a8b16a59ee6149ae2b51418f3e15da0d29461bc16b908a82ebd"
	testChainId = int64(97)
)

func TestMintNFT(t *testing.T) {
	// init cdnft instance
	client, _ := ethclient.Dial(rpc)
	cdNftInstance, _ := cdnft.NewCdnft(common.HexToAddress(cdNFTAddr), client)
	// init owner
	ownerPrivateKey, _ := crypto.HexToECDSA(ownerPK)
	ownerAuth, _ := bind.NewKeyedTransactorWithChainID(ownerPrivateKey, big.NewInt(testChainId))

	recipient, tokenURI, ifid, log := common.HexToAddress("0x54F1F088a217378D8D9464BDc8Fa190cda83B24a"), "https://docs.openzeppelin.com/", "id123456789", "所有权转移"

	tx, err := cdNftInstance.MintNFT(ownerAuth, recipient, tokenURI, ifid, log)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(tx.Hash())
}
