package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taarushv/helios/contracts/erc20"
)

func getTxSenderAddress(tx *types.Transaction, client *ethclient.Client) string {
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	msg, _ := tx.AsMessage(types.NewEIP155Signer(chainID))
	return msg.From().Hex()
}

func formatEthWeiToEther(etherAmount *big.Int) float64 {
	var base, exponent = big.NewInt(10), big.NewInt(18)
	denominator := base.Exp(base, exponent, nil)
	// Convert to float for precision
	tokensSentFloat := new(big.Float).SetInt(etherAmount)
	denominatorFloat := new(big.Float).SetInt(denominator)
	// Divide and return the final result
	final, _ := new(big.Float).Quo(tokensSentFloat, denominatorFloat).Float64()
	return final
}

func isTxMined(txHash string, client *ethclient.Client) bool {
	finalTxHash := common.HexToHash(txHash)
	_, isPending, err := client.TransactionByHash(context.Background(), finalTxHash)
	if err != nil {
		log.Fatal(err)
	}
	return !isPending
}

func hasTxFailed(txHash string, client *ethclient.Client) bool {
	if isTxMined(txHash, client) {
		receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash(txHash))
		if err != nil {
			log.Fatal(err)
		}
		if receipt.Status == 1 {
			return false
		} else {
			return true
		}
	} else {
		return false
	}
}

func getTxIDByHash(txHash string) string {
	var (
		r map[string]interface{}
	)
	es, _ := elasticsearch.NewDefaultClient()
	// 3. Search for the indexed documents
	//
	// Build the request body.
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"txHash": txHash,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	// Perform the search request.
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("transactions"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	// Print the ID and document source for each hit.
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		return fmt.Sprintf("%v", hit.(map[string]interface{})["_id"])
	}
	return ""
}

func getBlockNoByTxHash(txHash string, client *ethclient.Client) int64 {
	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash(txHash))
	if err != nil {
		log.Fatal(err)
	}
	return receipt.BlockNumber.Int64()
}

// After a block is mined, we iterate through the txs and mark the ones that've been mined
func TxMinedUpdate(txHash string, client *ethclient.Client) {
	txID := getTxIDByHash(txHash)
	if txID != "" {
		type final struct {
			TxMined       bool  `json:"txMined"`
			TxFailed      bool  `json:"txFailed"`
			BlockIncluded int64 `json:"blockIncluded"`
		}
		body := struct {
			Doc final `json:"doc"`
		}{}

		body.Doc = final{TxMined: isTxMined(txHash, client), TxFailed: hasTxFailed(txHash, client), BlockIncluded: getBlockNoByTxHash(txHash, client)}
		es, _ := elasticsearch.NewDefaultClient()
		jsonBytes, _ := json.Marshal(body)
		// tag:a0f4e902d18460337684d74ea932fbe9[]
		res, err := es.Update(
			"transactions",
			txID,
			bytes.NewReader(jsonBytes),
			es.Update.WithPretty(),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			fmt.Println("Error getting the response:", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	} else {
		tx, _, err := client.TransactionByHash(context.Background(), common.HexToHash(txHash))
		if err != nil {
			log.Fatal(err)
		}
		// We re-examine "stealth txs" (only full mode)
		txClassifier(tx, client, true, true)
	}
}

// Format # of tokens transferred into required float
func formatERC20Decimals(tokensSent *big.Int, tokenAddress common.Address, client *ethclient.Client) float64 {
	// Create a ERC20 instance and connect to geth to get decimals
	tokenInstance, _ := erc20.NewErc20(tokenAddress, client)
	decimals, _ := tokenInstance.Decimals(nil)
	// Construct a denominator based on the decimals
	// 18 decimals would result in denominator = 10^18
	var base, exponent = big.NewInt(10), big.NewInt(int64(decimals))
	denominator := base.Exp(base, exponent, nil)
	// Convert to float for precision
	tokensSentFloat := new(big.Float).SetInt(tokensSent)
	denominatorFloat := new(big.Float).SetInt(denominator)
	// Divide and return the final result
	final, _ := new(big.Float).Quo(tokensSentFloat, denominatorFloat).Float64()
	// TODO Take big.Accuracy into account
	return final
}

func formatChainlinkOraclePrice(submission *big.Int, decimals uint8) float64 {
	var decimalsInt64 int64
	decimalsInt64 = int64(decimals)
	var base, exponent = big.NewInt(10), big.NewInt(decimalsInt64)
	denominator := base.Exp(base, exponent, nil)
	// Convert to float for precision
	price := new(big.Float).SetInt(submission)
	priceFloat := new(big.Float).SetInt(denominator)
	// Divide and return the final result
	final, _ := new(big.Float).Quo(price, priceFloat).Float64()
	return final
}

func getTokenSymbol(tokenAddress common.Address, client *ethclient.Client) string {
	tokenIntance, _ := erc20.NewErc20(tokenAddress, client)
	sym, _ := tokenIntance.Symbol(nil)
	return sym
}
func getTokenName(tokenAddress common.Address, client *ethclient.Client) string {
	tokenIntance, _ := erc20.NewErc20(tokenAddress, client)
	name, _ := tokenIntance.Name(nil)
	return name
}
