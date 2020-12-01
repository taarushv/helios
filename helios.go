package main

import (
	"flag"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/taarushv/helios/services"
	"github.com/taarushv/helios/tools"
)

func main() {
	// Load the env variable (contains ipc path/infura ws api url)
	godotenv.Load(".env")
	// 'Quick' mode displays mempool data on the console
	// 'Full' initiates a elastic search client and stores data
	var modeType = flag.String("mode", "quick", "Quick mode vs with 'full' inserts to elastic search")
	// Flush helper flag to delete all
	var flush = flag.String("flush", "", "Index you want to delete")
	flag.Parse()
	// `go run helios.go -mode=full -flush=transactions` will delete all txs stored in ES
	if *flush != "" {
		fmt.Println("Flushing the index:", *flush)
		// This is irreversable, be careful!
		tools.FlushIndexData(*flush)
	} else {
		// Initiate client, flags: -client=infura or -client=local
		rpcClient := services.InitRPCClient()
		if *modeType == "full" {
			// Stream news txs, store them depending on mode
			services.StreamNewTxs(rpcClient, true)
			//TODO: validate queries before updating a block after being mined
			//services.StreamNewBlocks(rpcClient)
		} else {
			services.StreamNewTxs(rpcClient, false)
		}
	}

}
