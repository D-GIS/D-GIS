package main


import (
	"net/http"
    "github.com/hprose/hprose-golang/rpc"
	log "logutil"
)

func main() {
	log.InitializeLogger()
	InitializeEthClient()
	InitializeClients()
	server := rpc.NewHTTPService()

	
	// TxtStorage functions
	server.AddFunction("DeployTxtStorage", DeployNewTxtStorage)
	server.AddFunction("GetPackedData", GetPackedData)
	server.AddFunction("GetReputation", GetReputation)
	server.AddFunction("GetEventsForReputation", GetEventsForReputation)
	server.AddFunction("GetEventsForData", GetEventsForData)

	// Clients functions
	server.AddFunction("RegisterClient", RegisterClient)

	log.Info("Registered server functions!")
	http.ListenAndServe(":8080", server)
}