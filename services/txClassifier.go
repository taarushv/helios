package services

import (
	"bytes"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	. "github.com/logrusorgru/aurora"
)

// Store function signature bytes for quick classification
var erc20Transfer = []byte{0xa9, 0x05, 0x9c, 0xbb}
var erc20Approve = []byte{0x09, 0x5e, 0xa7, 0xb3}
var linkOracleUpdate = []byte{0x20, 0x2e, 0xe0, 0xed}

// Store contract addresses for quick classification
var uniV2routerAddress = "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"

// Core classifier to tag txs in the mempool before they're executed
// We classify a tx and then pipe it into elastic search as a document entry
// Ex: Oracle updates (to backrun + liquidate underwater positions)
// Trades, to either frontrun and arb the trade or backrun a large order to take advantage of slippage
// Basic ERC20 approvals/transfers
func txClassifier(tx *types.Transaction, client *ethclient.Client, isStealth bool, fullMode bool) {
	// If the tx has no code to be executed, it's a direct ETH transfer (or contract fallback function?)
	if len(tx.Data()) == 0 {
		fmt.Println()
		fmt.Println(Yellow("New TX: ETH Direct Transfer"))
		fmt.Println("Hash: ", tx.Hash().Hex(), " Value: ", formatEthWeiToEther(tx.Value()))
		if fullMode {
			handleDirectTransfer(tx, client, isStealth)
		}
	} else {
		// If the tx has no recepient it's a contract deployment
		if tx.To() == nil {
			fmt.Println(White("New TX: Contract Deployment"))
			fmt.Println("Hash: ", tx.Hash().Hex())
			if fullMode {
				handleContractDeployment(tx, client, isStealth)
			}
		} else {
			// Now that we've ruled out the base cases, we classify contract interactions via function signature
			// We also have a KV pair of fn signatures + identifiers in the "4bytes" index (for use in nested methods)
			// "filters" and their assosiated insert methods go here
			if len(tx.Data()) >= 4 {
				// Check first 4 bytes against our filters
				if bytes.Equal(tx.Data()[:4], erc20Approve) { // Standard ERC20 Approve and Transfer
					handleERC20Approve(tx, client, isStealth, fullMode)
				} else if bytes.Equal(tx.Data()[:4], erc20Transfer) {
					handleERC20Transfer(tx, client, isStealth, fullMode)
				} else if tx.To().Hex() == uniV2routerAddress { // Uniswap related trades
					handleUniswapTrade(tx, client, isStealth, fullMode)
				} else if bytes.Equal(tx.Data()[:4], linkOracleUpdate) { // Chainlink oracle updates
					handleChainlinkOracleUpdate(tx, client, isStealth, fullMode)
				} else {
					// "Everything else" for now, until I add more filters
					// TODO: Identify method by querying "4bytes"
					if fullMode {
						handleMiscTx(tx, client, isStealth)
					}

				}
			} else {
				if fullMode {
					// Weird txs (<4 bytes indicates no function identifier, likely a wallet error or using data field to stamp bytes)
					handleEdgeTx(tx, client, isStealth)
				}

			}

		}
	}

}
