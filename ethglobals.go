package main

import (
	log "logutil"
	"github.com/ethereum/go-ethereum/ethclient"
)

var ethercl *ethclient.Client;

func InitializeEthClient(){
	cl, err := ethclient.Dial("ws://127.0.0.1:8545")
	if(err != nil){
		//panic(err)
	}

	ethercl = cl;

	log.Info("Initialized ethclient!")
}