package main

/*
RPC Functions:
- DeployNewTxtStorage() - Deploys a new TxtStorage and returns to the client the given address
- GetEventsForData() - Get the logged events regarding RequestData() for a specific address
- GetEventsForReputation() - Gets the logged events regarding ReputationGiven() for a specific address
- GetReputation() - Gets the reputation for a specific address
- GetPackedData() - Gets a JSON containing the information regarding the given TxtStorage
*/

import(
	"context"
	"fmt"
	"math/big"
	gcrypt "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"contracts"
	"crypto/ecdsa"
	log "logutil"

	"encoding/json"

	"strconv"

	golog "log"
)

type TxtStoragePack struct{
	Address string
	Reputation int64
	HashFile string
	FileName string
}

func DeployNewTxtStorage(fileName string, fileHash string, price int64, privateKey string, clientAddr string) string{  // We MUST automate the clients registration process
	log.Call("DeployNewTxtStorage() ...")
	privateK, err := gcrypt.HexToECDSA(privateKey)
	if(err != nil){
		//panic(err)
		log.Info("We've got an error while converting private key to ECDSA.")
	}

	pubK := privateK.Public()
	pubECDSA := pubK.(*ecdsa.PublicKey)
	fromAddr := gcrypt.PubkeyToAddress(*pubECDSA)


	gasPrice, err := ethercl.SuggestGasPrice(context.Background())
	if(err != nil){
		//panic(err)
		log.Info("Error while calling SuggestGasPrice()")
	}

	nonce, err := ethercl.PendingNonceAt(context.Background(), fromAddr)
	if(err != nil){
		//panic(err)
		log.Info("Error while calling PendingNonceAt()")
	}

	auth := bind.NewKeyedTransactor(privateK)
	auth.GasLimit = uint64(300000)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasPrice = gasPrice
	auth.Value = big.NewInt(int64(0))

	address, tx, instance, err := contracts.DeployTxtStorage(auth, ethercl, fileName, big.NewInt(price), fileHash)
	_ = tx
	_ = instance
	if(err != nil){
		//panic(err)
		log.Info("Error while deploying TxtStorage")
		return ""
	}

	// Add this to the client's list of contracts..
	//clients[clientAddr].contracts[address.Hash().Hex()] = instance

	fmt.Println("Created instance of TxtStorage on address:", address.Hash().Hex())


	// Time to set  up the address thingy..
	defer func() {

	go RegisterLogEventsForTxtUploader(clients[clientAddr], "ReputationGiven", instance, address.Hash().Hex())
	go RegisterLogEventsForTxtUploader(clients[clientAddr], "RequestFiles", instance, address.Hash().Hex())

	}()
	
	return address.Hash().Hex()
}

// We don't store events at all in the server,
// we instead allow clients to submit requests over the server in order to list all the event logs in a contract.
// cAddr: The address in which the contract has been deployed.
// clientId: The client which is making the call.
func GetEventsForReputation(cAddr string, clientId int) string{
	log.Call("GetEventsForReputation")
	instance, err := contracts.NewTxtStorage(common.HexToAddress(cAddr), ethercl)
	if(err != nil){
		log.Info(err.Error())
	}

	iter, _ := instance.FilterReputationGiven(&bind.FilterOpts{
		Context: context.Background()})

	client := GetClientById(clientId)
	
	if(client.client_id != 0){
		for iter.Next(){
			// pack each data and send it to a client
			client.stub.AddLogEvent(iter.Event.Sender.Hash().Hex(), "ReputationGiven", cAddr)
		}
		return "Ok"
	}
	
	log.Info("Invalid CID")
	return "Client ID has not been found."
}

func GetEventsForData(cAddr string, clientId int) string{
	log.Call("GetEventsForData")
	instance, err := contracts.NewTxtStorage(common.HexToAddress(cAddr), ethercl)
	if(err != nil){
		log.Info(err.Error())
	}

	iter, _ := instance.FilterRequestData(&bind.FilterOpts{
		Context: context.Background()})

	client := GetClientById(clientId)
	
	if(client.client_id != 0){
		for iter.Next(){
			// pack each data and send it to a client
			client.stub.AddLogEvent(iter.Event.Sender.Hash().Hex(), "RequestFiles", cAddr)
		}
		return "Ok"
	}
	
	log.Info("Invalid CID")
	return "Client ID has not been found."
}

func GetReputation(cId int, cAddr string) string{
	log.Call("GetReputation from " + strconv.Itoa(cId) + " @" + cAddr)
	client := GetClientById(cId)
	if(client.client_id == 0){
		log.Info("Invalid CID")
		return "Invalid CID:0"
	}

	rep,  _ := client.m_contracts[cAddr].GetReputation(nil)
	return "Ok:" + rep.String()
}

func GetPackedData(cId int, cAddr string) string{
	log.Call("GetPackedData from " + string(cId) + " @" + cAddr)
	client := GetClientById(cId)
	if(client.client_id == 0){
		return "Invalid CID:0"
	}

	instance, _ := contracts.NewTxtStorage(common.HexToAddress(cAddr), ethercl)  // Maybe?
	fName, err := instance.GetName(nil)
	if(err != nil){
		golog.Fatal(err)
	}

	hFile, err := instance.GetHashData(nil)
	if(err != nil){
		golog.Fatal(err)
	}

	rep, err := instance.GetReputation(nil)
	if(err != nil){
		golog.Fatal(err)
	}

	pStruct := TxtStoragePack{Address: cAddr, FileName: fName, HashFile: hFile, Reputation: rep.Int64()}
	
	jString, _ := json.Marshal(pStruct)
	
	return string(jString[:])
}