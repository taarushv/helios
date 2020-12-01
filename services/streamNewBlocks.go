package services

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func handleBlock(blockHash common.Hash, client *ethclient.Client) {
	block, err := client.BlockByHash(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}
	//filters.TransfersInBlock(block, client)
	// Test output from the channel by logging it

	pipeBlock(block)
	// Find out all the transactions that emit ERC20 transfer event
	//fmt.Println(".....")

}

func StreamNewBlocks(client *rpc.Client) {
	// Go channel to pipe data from client subscriptions
	newBlocksChannel := make(chan *types.Header, 10)

	// Subscribe to receive one time events for new txs
	// i.e Pipe new data to the channel every time a block is mined
	client.EthSubscribe(
		context.Background(), newBlocksChannel, "newHeads", // no additional args
	)

	fmt.Println("Subscribed to new blocks")

	// Configure chain ID and signer to ensure you're configured to mainnet
	//chainID, _ := client.NetworkID(context.Background())
	//signer := types.NewEIP155Signer(chainID)

	for {
		select {
		// Code block is executed when a new block is piped to the channel
		case lastBlockHeader := <-newBlocksChannel:
			fmt.Println("New block in channel")
			func() {
				go handleBlock(lastBlockHeader.Hash(), GetCurrentClient())
			}()
		}
	}
}
