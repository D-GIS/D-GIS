package main

/*
import (
	_"github.com/ethereum/go-ethereum/common"
	"context"
	"math/big"
	"crypto/ecdsa"
	"fmt"
	"log"
	_"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"

	_"contracts"
)

func main() {
	client, err := ethclient.Dial("wss://ropsten.infura.io/ws")
	

	if err != nil{
		log.Fatal(err)
	}
	

	fmt.Println("We have a connection!")

	privateKey, err := crypto.HexToECDSA("391d83052b7db85305127c4726010dea806613b9e23aa27a98087fecd859e834")
	if(err != nil){
		log.Fatal(err)
	}

	publicKey := privateKey.Public()


	publicECDSA := publicKey.(*ecdsa.PublicKey)

	fromAddress := crypto.PubkeyToAddress(*publicECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	gasPrice, err := client.SuggestGasPrice(context.Background())

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice
	auth.Value = big.NewInt(int64(0))
	
}*/