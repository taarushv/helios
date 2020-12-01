package services

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// Method to stream new transactions as they're discovered by the node
// We create a new channel to plug into the mempool and listen to new txs
// They're then passed through the tx classifier to be parsed and eventually piped into elastic search
func StreamNewTxs(rpcClient *rpc.Client, fullMode bool) {

	// Go channel to pipe data from client subscription
	newTxsChannel := make(chan common.Hash)

	// Subscribe to receive one time events for new txs
	rpcClient.EthSubscribe(
		context.Background(), newTxsChannel, "newPendingTransactions", // no additional args
	)
	client := GetCurrentClient()
	fmt.Println("Subscribed to mempool txs")

	// Configure chain ID and signer to ensure you're configured to mainnet
	chainID, _ := client.NetworkID(context.Background())
	signer := types.NewEIP155Signer(chainID)

	for {
		select {
		// Code block is executed when a new tx hash is piped to the channel
		case transactionHash := <-newTxsChannel:
			// Get transaction object from hash by querying the client
			tx, is_pending, _ := client.TransactionByHash(context.Background(), transactionHash)
			// If tx is valid and still unconfirmed
			if is_pending {
				_, _ = signer.Sender(tx)
				//fmt.Println(fullMode)
				handleTransaction(tx, client, false, fullMode)
			}
		}
	}
}

func handleTransaction(tx *types.Transaction, client *ethclient.Client, isStealth bool, fullMode bool) {
	// Log the tx and pass it through the classifier
	//fmt.Println("New TX, hash: ", tx.Hash().String())
	txClassifier(tx, client, isStealth, fullMode)
}
