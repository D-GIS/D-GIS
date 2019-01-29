package main


/////////////////////////////////////////
////////////// Regjistron klientet.
////////////////////////////////////////


/*
TODO:
	- Bejm mire qe i ruajme logsat ne memorie? A do na duhen ti perdorim me vone?

*/


import (
	"contracts"
	"context"
	log "logutil"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hprose/hprose-golang/rpc"
)

type LogEvent struct{
	deployedContract	interface{}
	name 		string
	result		interface{}
	client 		Client
	logId		uint64
}

type ClientStub struct{
	AddLogEvent		func(address string, eventType string, cAddr string)
}

type Client struct{
	address string
	client *rpc.HTTPClient
	client_id int
	stub	*ClientStub
	m_contracts	map[string]*contracts.TxtStorage
}

var clients map[string]Client;
var events []LogEvent;

var numclients *int;
var numLogs *uint64;

func InitializeClients(){
	clients = make(map[string]Client)
	events = make([]LogEvent, 0)
	numclients = new(int)
	numLogs = new(uint64)

	*numclients = 1
}

func RegisterClient(addr string) int{
	log.Call("RegisterClient with address: " + addr)

	clientT := rpc.NewHTTPClient("http://" + addr)
	cstub := new(ClientStub)
	clientT.UseService(cstub)

	cl := Client{address:  addr,
				client_id: *numclients + 1,
				client: clientT,
				stub: cstub,
				m_contracts: make(map[string]*contracts.TxtStorage)};

	clients[addr] = cl
	//clients[addr].stub.AddLogEvent("AAA", "AA", "A")

	*numclients++
	return *numclients
}

func GetClientById(id int) Client{
	for _, v := range clients{
		if(v.client_id == id){
			return v
		}
	}
	return Client{client_id: 0}
}


func RegisterLogEventsForTxtUploader(cl Client, nm string, contract interface{}, address string) {
	instance := contract.(*contracts.TxtStorage)
	filesHash, _ := instance.GetHashData(nil)

	//logEventRep := LogEvent{}

	if(nm == "ReputationGiven"){
		repGiven := make(chan *contracts.TxtStorageReputationGiven)
		startNum1 := uint64(0)

		instance.WatchReputationGiven(&bind.WatchOpts{
			Context: context.Background(),
			Start: &startNum1 }, repGiven)
		
		resRep := <- repGiven
		
		cl.stub.AddLogEvent(resRep.Sender.Hash().Hex() + ":" + filesHash, nm, address)
	}

	if(nm == "RequestFiles"){
		getFiles := make(chan *contracts.TxtStorageRequestData)

		startNum2 := uint64(0)
	
		instance.WatchRequestData(&bind.WatchOpts{
			Context: context.Background(),
			Start: &startNum2 }, getFiles)
	
		resFiles := <- getFiles
		cl.stub.AddLogEvent(resFiles.Sender.Hash().Hex() + ":" + filesHash, nm, address)
	}

	//_ = append(events, logEventRep)
	RegisterLogEventsForTxtUploader(cl, nm, contract, address)
}
