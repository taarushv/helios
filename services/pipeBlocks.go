package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Pipe new blocks into elastic search

func pipeBlock(block *types.Block) {
	fmt.Println("MINED: Block #", block.Number())
	for _, tx := range block.Transactions() {
		fmt.Println(tx.Hash().Hex())
		TxMinedUpdate(tx.Hash().Hex(), GetCurrentClient())
	}
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error connecting to es client: %s", err)
	}

	body := struct {
		No           int64  `json:"blockNo"`
		Hash         string `json:"blockHash"`
		BlockTainted bool   `json:"blockTainted"`
	}{}
	body.No = block.Number().Int64()
	body.Hash = block.Hash().Hex()
	var client = GetCurrentClient()
	body.BlockTainted, _ = CheckBlockReorderTaint(block, client)
	jsonBytes, _ := json.Marshal(body)
	// Set up the request object.
	req := esapi.IndexRequest{
		Index:   "blocks",
		Body:    bytes.NewReader(jsonBytes),
		Refresh: "true",
	}
	// Perform the request with the client.
	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("[%s] Error indexing document", res.Status())
		fmt.Println(res)
		panic("a prob")
	} else {
		// Deserialize the response into a map.
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and indexed document version.
			log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
		}
	}
}

// Check if the block txs order indicates deviation from standard client rules
// A block is said to be tainted if a miner prioritizes a tx over other despite the latter paying higher gas
// If a miner were to be extracting MEV, they would prioritize execution of their own above else (despite gas priority)
// TODO: Lot's of test cases to ensure accuracy, return addition info about types of tx prioritized
func CheckBlockReorderTaint(block *types.Block, client *ethclient.Client) (bool, int) {
	var temp []uint64
	var gasPricesFinal []uint64
	// Iterate through all transactions in the block
	for _, tx := range block.Transactions() {
		// Push to the slice
		temp = append(gasPricesFinal, tx.GasPrice().Uint64())
		gasPricesFinal = temp
	}
	// If any tx (iterated by index in block) is prioritized despite the next tx paying more, we flag the block
	for i, v := range gasPricesFinal {
		// If we reach the end of the block
		if i == len(gasPricesFinal)-1 {
			return false, -1
		}
		// Deviates from standard client rules
		// Also returns index at which deviation occurs
		if v < gasPricesFinal[i+1] {
			return true, i
		}
	}
	// Returning -1 if there is no deviation
	return false, -1
}
