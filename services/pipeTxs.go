package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"time"

	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	. "github.com/logrusorgru/aurora"

	"github.com/taarushv/helios/contracts/erc20"
)

// Mempool => ES index document
func handleDirectTransfer(tx *types.Transaction, client *ethclient.Client, isStealth bool) {
	// Connect to our es client
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error connecting to es client: %s", err)
	}
	// Start building a document
	body := struct {
		// Time the tx was discovered and other details (for data analysis + Kibana)
		TimeSeen int64  `json:"timeFirstDiscovered"`
		Hash     string `json:"txHash"`
		// Classifcation and other info
		Type     string  `json:"txType"`
		From     string  `json:"from"`
		To       string  `json:"to"`
		Value    float64 `json:"txValue"`
		Nonce    uint64  `json:"nonce"`
		GasPrice float64 `json:"gasPrice"`
		Gas      float64 `json:"gas"`
		// Custom tags that are updated after a block including the tx is mined
		Mined         bool  `json:"txMined"`
		BlockIncluded int64 `json:"blockIncluded"`
		Failed        bool  `json:"txFailed"`
		Stealth       bool  `json:"txStealth"`
		// WIP
		// LocalLogs are the logs of the txs when executed against the local EVM (ganache fork that's updated every block)
		// Tags are to classify post logs, to track
		LocalLogs []string `json:"localLogs"`
		Tags      []string `json:"tags"`
	}{}
	body.TimeSeen = time.Now().Unix()
	body.Hash = tx.Hash().Hex()
	body.Type = "directTransfer"
	body.From = getTxSenderAddress(tx, client)
	body.To = tx.To().Hex()
	body.Value = formatEthWeiToEther(tx.Value())
	body.Nonce = tx.Nonce()
	formattedGas := new(big.Int).SetUint64(tx.Gas())
	body.Gas = formatEthWeiToEther(formattedGas)
	body.GasPrice = formatEthWeiToEther(tx.GasPrice())
	body.Mined = isTxMined(tx.Hash().Hex(), client)
	body.Failed = hasTxFailed(tx.Hash().Hex(), client)
	if isStealth {
		body.BlockIncluded = getBlockNoByTxHash(tx.Hash().Hex(), client)
	}
	body.Stealth = isStealth
	jsonBytes, _ := json.Marshal(body)
	// Set up the request object.
	req := esapi.IndexRequest{
		Index:   "transactions",
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

func handleContractDeployment(tx *types.Transaction, client *ethclient.Client, isStealth bool) {
	fmt.Println("Contract deployment, TX:", tx.Hash().String())
	// Connect to our es client
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error connecting to es client: %s", err)
	}
	// Start building a document
	body := struct {
		// Time the tx was discovered and other details (for data analysis + Kibana)
		TimeSeen int64  `json:"timeFirstDiscovered"`
		Hash     string `json:"txHash"`
		// Classifcation and other info
		Type     string  `json:"txType"`
		From     string  `json:"from"`
		To       string  `json:"to"`
		Value    float64 `json:"txValue"`
		Nonce    uint64  `json:"nonce"`
		GasPrice float64 `json:"gasPrice"`
		Gas      float64 `json:"gas"`
		// Custom tags that are updated after a block including the tx is mined
		Mined         bool  `json:"txMined"`
		BlockIncluded int64 `json:"blockIncluded"`
		Failed        bool  `json:"txFailed"`
		Stealth       bool  `json:"txStealth"`
		// WIP
		// LocalLogs are the logs of the txs when executed against the local EVM (ganache fork that's updated every block)
		// Tags are to classify post logs, to track
		LocalLogs []string `json:"localLogs"`
		Tags      []string `json:"tags"`
	}{}
	body.TimeSeen = time.Now().Unix()
	body.Hash = tx.Hash().Hex()
	body.Type = "contractDeployment"
	body.From = getTxSenderAddress(tx, client)
	body.Value = formatEthWeiToEther(tx.Value())
	body.Nonce = tx.Nonce()
	formattedGas := new(big.Int).SetUint64(tx.Gas())
	body.Gas = formatEthWeiToEther(formattedGas)
	body.GasPrice = formatEthWeiToEther(tx.GasPrice())
	body.Mined = isTxMined(tx.Hash().Hex(), client)
	body.Failed = hasTxFailed(tx.Hash().Hex(), client)
	if isStealth {
		body.BlockIncluded = getBlockNoByTxHash(tx.Hash().Hex(), client)
	}
	body.Stealth = isStealth
	jsonBytes, _ := json.Marshal(body)
	// Set up the request object.
	req := esapi.IndexRequest{
		Index:   "transactions",
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

type ERC20TransferParsedInput struct {
	from         string
	to           string
	value        float64
	tokenAddress string
	tokenSymbol  string
	txHash       string
}

var erc20Abi, _ = abi.JSON(strings.NewReader(erc20.Erc20ABI))

func handleERC20Approve(tx *types.Transaction, client *ethclient.Client, isStealth bool, fullMode bool) {
	type finalParsedData struct {
		TokensSpender string `json:"tokensSpender"`
		TokenAddress  string `json:"tokenAddress"`
		TokenSymbol   string `json:"tokenSymbol"`
	}
	tokenInstance, _ := erc20.NewErc20(*tx.To(), client)
	tokenSymbol, _ := tokenInstance.Symbol(nil)
	//fmt.Println(tokenSymbol)
	fmt.Println()
	fmt.Println(Blue("New TX: ERC20 Approval"))
	fmt.Println("Hash: ", tx.Hash().Hex(), " Token: ", tokenSymbol)
	if fullMode {
		// Connect to our es client
		es, err := elasticsearch.NewDefaultClient()
		if err != nil {
			log.Fatalf("Error connecting to es client: %s", err)
		}
		// Start building a document
		body := struct {
			// Time the tx was discovered and other details (for data analysis + Kibana)
			TimeSeen int64  `json:"timeFirstDiscovered"`
			Hash     string `json:"txHash"`
			// Classifcation and other info
			Type            string          `json:"txType"`
			FinalParsedData finalParsedData `json:"finalParsedData"`
			From            string          `json:"from"`
			To              string          `json:"to"`
			Value           float64         `json:"txValue"`
			Nonce           uint64          `json:"nonce"`
			GasPrice        float64         `json:"gasPrice"`
			Gas             float64         `json:"gas"`
			// Custom tags that are updated after a block including the tx is mined
			Mined         bool  `json:"txMined"`
			BlockIncluded int64 `json:"blockIncluded"`
			Failed        bool  `json:"txFailed"`
			Stealth       bool  `json:"txStealth"`
			// WIP
			// LocalLogs are the logs of the txs when executed against the local EVM (ganache fork that's updated every block)
			// Tags are to classify post logs, to track
			LocalLogs []string `json:"localLogs"`
			Tags      []string `json:"tags"`
		}{}
		body.TimeSeen = time.Now().Unix()
		body.Hash = tx.Hash().Hex()
		body.Type = "erc20Approve"
		body.FinalParsedData = finalParsedData{TokenAddress: tx.To().String(), TokenSymbol: tokenSymbol, TokensSpender: common.BytesToAddress(tx.Data()[4:36]).Hex()}
		body.From = getTxSenderAddress(tx, client)
		body.To = tx.To().Hex()
		body.Value = formatEthWeiToEther(tx.Value())
		body.Nonce = tx.Nonce()
		formattedGas := new(big.Int).SetUint64(tx.Gas())
		body.Gas = formatEthWeiToEther(formattedGas)
		body.GasPrice = formatEthWeiToEther(tx.GasPrice())
		body.Mined = isTxMined(tx.Hash().Hex(), client)
		body.Failed = hasTxFailed(tx.Hash().Hex(), client)
		if isStealth {
			body.BlockIncluded = getBlockNoByTxHash(tx.Hash().Hex(), client)
		}
		body.Stealth = isStealth
		jsonBytes, _ := json.Marshal(body)
		// Set up the request object.
		req := esapi.IndexRequest{
			Index:   "transactions",
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

}

func handleERC20Transfer(tx *types.Transaction, client *ethclient.Client, isStealth bool, fullMode bool) {

	type finalParsedData struct {
		TokenTo      string  `json:"tokenTo"`
		TokenFrom    string  `json:"tokenFrom"`
		TokenAmount  float64 `json:"tokenAmount"`
		TokenAddress string  `json:"tokenAddress"`
		TokenSymbol  string  `json:"tokenSymbol"`
	}

	tokensSent := new(big.Int).SetBytes((tx.Data()[36:68]))
	tokenInstance, _ := erc20.NewErc20(*tx.To(), client)
	tokenSymbol, _ := tokenInstance.Symbol(nil)
	fmt.Println()
	fmt.Println(Blue("New TX: ERC20 Transfer"))
	fmt.Println("Hash: ", tx.Hash().Hex(), " Amount: ", formatERC20Decimals(tokensSent, *tx.To(), client), tokenSymbol)
	if fullMode {
		// Connect to our es client
		es, err := elasticsearch.NewDefaultClient()
		if err != nil {
			log.Fatalf("Error connecting to es client: %s", err)
		}
		// Start building a document
		body := struct {
			// Time the tx was discovered and other details (for data analysis + Kibana)
			TimeSeen int64  `json:"timeFirstDiscovered"`
			Hash     string `json:"txHash"`
			// Classifcation and other info
			Type            string          `json:"txType"`
			FinalParsedData finalParsedData `json:"finalParsedData"`
			From            string          `json:"from"`
			To              string          `json:"to"`
			Value           float64         `json:"txValue"`
			Nonce           uint64          `json:"nonce"`
			GasPrice        float64         `json:"gasPrice"`
			Gas             float64         `json:"gas"`
			// Custom tags that are updated after a block including the tx is mined
			Mined         bool  `json:"txMined"`
			BlockIncluded int64 `json:"blockIncluded"`
			Failed        bool  `json:"txFailed"`
			Stealth       bool  `json:"txStealth"`
			// WIP
			// LocalLogs are the logs of the txs when executed against the local EVM (ganache fork that's updated every block)
			// Tags are to classify post logs, to track
			LocalLogs []string `json:"localLogs"`
			Tags      []string `json:"tags"`
		}{}
		body.TimeSeen = time.Now().Unix()
		body.Hash = tx.Hash().Hex()
		body.Type = "erc20Transfer"
		body.FinalParsedData = finalParsedData{TokenAddress: tx.To().String(), TokenSymbol: tokenSymbol, TokenFrom: getTxSenderAddress(tx, client), TokenTo: common.BytesToAddress(tx.Data()[4:36]).Hex(), TokenAmount: formatERC20Decimals(tokensSent, *tx.To(), client)}
		body.From = getTxSenderAddress(tx, client)
		body.To = tx.To().Hex()
		body.Value = formatEthWeiToEther(tx.Value())
		body.Nonce = tx.Nonce()
		formattedGas := new(big.Int).SetUint64(tx.Gas())
		body.Gas = formatEthWeiToEther(formattedGas)
		body.GasPrice = formatEthWeiToEther(tx.GasPrice())
		body.Mined = isTxMined(tx.Hash().Hex(), client)
		body.Failed = hasTxFailed(tx.Hash().Hex(), client)
		if isStealth {
			body.BlockIncluded = getBlockNoByTxHash(tx.Hash().Hex(), client)
		}
		body.Stealth = isStealth
		jsonBytes, _ := json.Marshal(body)
		// Set up the request object.
		req := esapi.IndexRequest{
			Index:   "transactions",
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

}

func handleLinkOracleUpdate(tx *types.Transaction, client *ethclient.Client, isStealth bool, final chainlinkOraclePriceUpdate) {

	// Connect to our es client
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error connecting to es client: %s", err)
	}
	// Start building a document
	body := struct {
		// Time the tx was discovered and other details (for data analysis + Kibana)
		TimeSeen int64  `json:"timeFirstDiscovered"`
		Hash     string `json:"txHash"`
		// Classifcation and other info
		Type            string                     `json:"txType"`
		FinalParsedData chainlinkOraclePriceUpdate `json:"finalParsedData"`
		From            string                     `json:"from"`
		To              string                     `json:"to"`
		Value           float64                    `json:"txValue"`
		Nonce           uint64                     `json:"nonce"`
		GasPrice        float64                    `json:"gasPrice"`
		Gas             float64                    `json:"gas"`
		// Custom tags that are updated after a block including the tx is mined
		Mined         bool  `json:"txMined"`
		BlockIncluded int64 `json:"blockIncluded"`
		Failed        bool  `json:"txFailed"`
		Stealth       bool  `json:"txStealth"`
		// WIP
		// LocalLogs are the logs of the txs when executed against the local EVM (ganache fork that's updated every block)
		// Tags are to classify post logs, to track
		LocalLogs []string `json:"localLogs"`
		Tags      []string `json:"tags"`
	}{}
	body.TimeSeen = time.Now().Unix()
	body.Hash = tx.Hash().Hex()
	body.Type = "linkOracleUpdate"
	body.FinalParsedData = final
	body.From = getTxSenderAddress(tx, client)
	body.To = tx.To().Hex()
	body.Value = formatEthWeiToEther(tx.Value())
	body.Nonce = tx.Nonce()
	formattedGas := new(big.Int).SetUint64(tx.Gas())
	body.Gas = formatEthWeiToEther(formattedGas)
	body.GasPrice = formatEthWeiToEther(tx.GasPrice())
	body.Mined = isTxMined(tx.Hash().Hex(), client)
	body.Failed = hasTxFailed(tx.Hash().Hex(), client)
	if isStealth {
		body.BlockIncluded = getBlockNoByTxHash(tx.Hash().Hex(), client)
	}
	body.Stealth = isStealth
	jsonBytes, _ := json.Marshal(body)
	// Set up the request object.
	req := esapi.IndexRequest{
		Index:   "transactions",
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

func handleUniAddETHLiq(tx *types.Transaction, client *ethclient.Client, isStealth bool, final UniswapAddLiquidityETHFinalInput) {

	// Connect to our es client
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error connecting to es client: %s", err)
	}
	// Start building a document
	body := struct {
		// Time the tx was discovered and other details (for data analysis + Kibana)
		TimeSeen int64  `json:"timeFirstDiscovered"`
		Hash     string `json:"txHash"`
		// Classifcation and other info
		Type            string                           `json:"txType"`
		FinalParsedData UniswapAddLiquidityETHFinalInput `json:"finalParsedData"`
		From            string                           `json:"from"`
		To              string                           `json:"to"`
		Value           float64                          `json:"txValue"`
		Nonce           uint64                           `json:"nonce"`
		GasPrice        float64                          `json:"gasPrice"`
		Gas             float64                          `json:"gas"`
		// Custom tags that are updated after a block including the tx is mined
		Mined         bool  `json:"txMined"`
		BlockIncluded int64 `json:"blockIncluded"`
		Failed        bool  `json:"txFailed"`
		Stealth       bool  `json:"txStealth"`
		// WIP
		// LocalLogs are the logs of the txs when executed against the local EVM (ganache fork that's updated every block)
		// Tags are to classify post logs, to track
		LocalLogs []string `json:"localLogs"`
		Tags      []string `json:"tags"`
	}{}
	body.TimeSeen = time.Now().Unix()
	body.Hash = tx.Hash().Hex()
	body.Type = "uniswapAddLiqETH"
	body.FinalParsedData = final
	body.From = getTxSenderAddress(tx, client)
	body.To = tx.To().Hex()
	body.Value = formatEthWeiToEther(tx.Value())
	body.Nonce = tx.Nonce()
	formattedGas := new(big.Int).SetUint64(tx.Gas())
	body.Gas = formatEthWeiToEther(formattedGas)
	body.GasPrice = formatEthWeiToEther(tx.GasPrice())
	body.Mined = isTxMined(tx.Hash().Hex(), client)
	body.Failed = hasTxFailed(tx.Hash().Hex(), client)
	if isStealth {
		body.BlockIncluded = getBlockNoByTxHash(tx.Hash().Hex(), client)
	}
	body.Stealth = isStealth
	jsonBytes, _ := json.Marshal(body)
	// Set up the request object.
	req := esapi.IndexRequest{
		Index:   "transactions",
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

func handleUniAddLiq(tx *types.Transaction, client *ethclient.Client, isStealth bool, final UniswapAddLiquidityFinalInput) {

	// Connect to our es client
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error connecting to es client: %s", err)
	}
	// Start building a document
	body := struct {
		// Time the tx was discovered and other details (for data analysis + Kibana)
		TimeSeen int64  `json:"timeFirstDiscovered"`
		Hash     string `json:"txHash"`
		// Classifcation and other info
		Type            string                        `json:"txType"`
		FinalParsedData UniswapAddLiquidityFinalInput `json:"finalParsedData"`
		From            string                        `json:"from"`
		To              string                        `json:"to"`
		Value           float64                       `json:"txValue"`
		Nonce           uint64                        `json:"nonce"`
		GasPrice        float64                       `json:"gasPrice"`
		Gas             float64                       `json:"gas"`
		// Custom tags that are updated after a block including the tx is mined
		Mined         bool  `json:"txMined"`
		BlockIncluded int64 `json:"blockIncluded"`
		Failed        bool  `json:"txFailed"`
		Stealth       bool  `json:"txStealth"`
		// WIP
		// LocalLogs are the logs of the txs when executed against the local EVM (ganache fork that's updated every block)
		// Tags are to classify post logs, to track
		LocalLogs []string `json:"localLogs"`
		Tags      []string `json:"tags"`
	}{}
	body.TimeSeen = time.Now().Unix()
	body.Hash = tx.Hash().Hex()
	body.Type = "uniswapAddLiq"
	body.FinalParsedData = final
	body.From = getTxSenderAddress(tx, client)
	body.To = tx.To().Hex()
	body.Value = formatEthWeiToEther(tx.Value())
	body.Nonce = tx.Nonce()
	formattedGas := new(big.Int).SetUint64(tx.Gas())
	body.Gas = formatEthWeiToEther(formattedGas)
	body.GasPrice = formatEthWeiToEther(tx.GasPrice())
	body.Mined = isTxMined(tx.Hash().Hex(), client)
	body.Failed = hasTxFailed(tx.Hash().Hex(), client)
	if isStealth {
		body.BlockIncluded = getBlockNoByTxHash(tx.Hash().Hex(), client)
	}
	body.Stealth = isStealth
	jsonBytes, _ := json.Marshal(body)
	// Set up the request object.
	req := esapi.IndexRequest{
		Index:   "transactions",
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

func handleUniRemoveETHLiq(tx *types.Transaction, client *ethclient.Client, isStealth bool, final UniswapRemoveLiquidityETHFinalInput) {
	// Connect to our es client
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error connecting to es client: %s", err)
	}
	// Start building a document
	body := struct {
		// Time the tx was discovered and other details (for data analysis + Kibana)
		TimeSeen int64  `json:"timeFirstDiscovered"`
		Hash     string `json:"txHash"`
		// Classifcation and other info
		Type            string                              `json:"txType"`
		FinalParsedData UniswapRemoveLiquidityETHFinalInput `json:"finalParsedData"`
		From            string                              `json:"from"`
		To              string                              `json:"to"`
		Value           float64                             `json:"txValue"`
		Nonce           uint64                              `json:"nonce"`
		GasPrice        float64                             `json:"gasPrice"`
		Gas             float64                             `json:"gas"`
		// Custom tags that are updated after a block including the tx is mined
		Mined         bool  `json:"txMined"`
		BlockIncluded int64 `json:"blockIncluded"`
		Failed        bool  `json:"txFailed"`
		Stealth       bool  `json:"txStealth"`
		// WIP
		// LocalLogs are the logs of the txs when executed against the local EVM (ganache fork that's updated every block)
		// Tags are to classify post logs, to track
		LocalLogs []string `json:"localLogs"`
		Tags      []string `json:"tags"`
	}{}
	body.TimeSeen = time.Now().Unix()
	body.Hash = tx.Hash().Hex()
	body.Type = "uniswapRemoveLiqETH"
	body.FinalParsedData = final
	body.From = getTxSenderAddress(tx, client)
	body.To = tx.To().Hex()
	body.Value = formatEthWeiToEther(tx.Value())
	body.Nonce = tx.Nonce()
	formattedGas := new(big.Int).SetUint64(tx.Gas())
	body.Gas = formatEthWeiToEther(formattedGas)
	body.GasPrice = formatEthWeiToEther(tx.GasPrice())
	body.Mined = isTxMined(tx.Hash().Hex(), client)
	body.Failed = hasTxFailed(tx.Hash().Hex(), client)
	if isStealth {
		body.BlockIncluded = getBlockNoByTxHash(tx.Hash().Hex(), client)
	}
	body.Stealth = isStealth
	jsonBytes, _ := json.Marshal(body)
	// Set up the request object.
	req := esapi.IndexRequest{
		Index:   "transactions",
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

func handleUniRemoveLiq(tx *types.Transaction, client *ethclient.Client, isStealth bool, final UniswapRemoveLiquidityFinalInput) {

	// Connect to our es client
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error connecting to es client: %s", err)
	}
	// Start building a document
	body := struct {
		// Time the tx was discovered and other details (for data analysis + Kibana)
		TimeSeen int64  `json:"timeFirstDiscovered"`
		Hash     string `json:"txHash"`
		// Classifcation and other info
		Type            string                           `json:"txType"`
		FinalParsedData UniswapRemoveLiquidityFinalInput `json:"finalParsedData"`
		From            string                           `json:"from"`
		To              string                           `json:"to"`
		Value           float64                          `json:"txValue"`
		Nonce           uint64                           `json:"nonce"`
		GasPrice        float64                          `json:"gasPrice"`
		Gas             float64                          `json:"gas"`
		// Custom tags that are updated after a block including the tx is mined
		Mined         bool  `json:"txMined"`
		BlockIncluded int64 `json:"blockIncluded"`
		Failed        bool  `json:"txFailed"`
		Stealth       bool  `json:"txStealth"`
		// WIP
		// LocalLogs are the logs of the txs when executed against the local EVM (ganache fork that's updated every block)
		// Tags are to classify post logs, to track
		LocalLogs []string `json:"localLogs"`
		Tags      []string `json:"tags"`
	}{}
	body.TimeSeen = time.Now().Unix()
	body.Hash = tx.Hash().Hex()
	body.Type = "uniswapRemoveLiq"
	body.FinalParsedData = final
	body.From = getTxSenderAddress(tx, client)
	body.To = tx.To().Hex()
	body.Value = formatEthWeiToEther(tx.Value())
	body.Nonce = tx.Nonce()
	formattedGas := new(big.Int).SetUint64(tx.Gas())
	body.Gas = formatEthWeiToEther(formattedGas)
	body.GasPrice = formatEthWeiToEther(tx.GasPrice())
	body.Mined = isTxMined(tx.Hash().Hex(), client)
	body.Failed = hasTxFailed(tx.Hash().Hex(), client)
	if isStealth {
		body.BlockIncluded = getBlockNoByTxHash(tx.Hash().Hex(), client)
	}
	body.Stealth = isStealth
	jsonBytes, _ := json.Marshal(body)
	// Set up the request object.
	req := esapi.IndexRequest{
		Index:   "transactions",
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

func handleUniFinalTrade(tx *types.Transaction, client *ethclient.Client, isStealth bool, final UniswapTradeFinal) {

	// Connect to our es client
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error connecting to es client: %s", err)
	}
	// Start building a document
	body := struct {
		// Time the tx was discovered and other details (for data analysis + Kibana)
		TimeSeen int64  `json:"timeFirstDiscovered"`
		Hash     string `json:"txHash"`
		// Classifcation and other info
		Type            string            `json:"txType"`
		FinalParsedData UniswapTradeFinal `json:"finalParsedData"`
		From            string            `json:"from"`
		To              string            `json:"to"`
		Value           float64           `json:"txValue"`
		Nonce           uint64            `json:"nonce"`
		GasPrice        float64           `json:"gasPrice"`
		Gas             float64           `json:"gas"`
		// Custom tags that are updated after a block including the tx is mined
		Mined         bool  `json:"txMined"`
		BlockIncluded int64 `json:"blockIncluded"`
		Failed        bool  `json:"txFailed"`
		Stealth       bool  `json:"txStealth"`
		// WIP
		// LocalLogs are the logs of the txs when executed against the local EVM (ganache fork that's updated every block)
		// Tags are to classify post logs, to track
		LocalLogs []string `json:"localLogs"`
		Tags      []string `json:"tags"`
	}{}
	body.TimeSeen = time.Now().Unix()
	body.Hash = tx.Hash().Hex()
	body.Type = "uniswapTrade"
	body.FinalParsedData = final
	body.From = getTxSenderAddress(tx, client)
	body.To = tx.To().Hex()
	body.Value = formatEthWeiToEther(tx.Value())
	body.Nonce = tx.Nonce()
	formattedGas := new(big.Int).SetUint64(tx.Gas())
	body.Gas = formatEthWeiToEther(formattedGas)
	body.GasPrice = formatEthWeiToEther(tx.GasPrice())
	body.Mined = isTxMined(tx.Hash().Hex(), client)
	body.Failed = hasTxFailed(tx.Hash().Hex(), client)
	if isStealth {
		body.BlockIncluded = getBlockNoByTxHash(tx.Hash().Hex(), client)
	}
	body.Stealth = isStealth
	jsonBytes, _ := json.Marshal(body)
	// Set up the request object.
	req := esapi.IndexRequest{
		Index:   "transactions",
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

func handleMiscTx(tx *types.Transaction, client *ethclient.Client, isStealth bool) {
	// Connect to our es client
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error connecting to es client: %s", err)
	}
	// Start building a document
	body := struct {
		// Time the tx was discovered and other details (for data analysis + Kibana)
		TimeSeen int64  `json:"timeFirstDiscovered"`
		Hash     string `json:"txHash"`
		// Classifcation and other info
		Type     string  `json:"txType"`
		From     string  `json:"from"`
		To       string  `json:"to"`
		Value    float64 `json:"txValue"`
		Nonce    uint64  `json:"nonce"`
		GasPrice float64 `json:"gasPrice"`
		Gas      float64 `json:"gas"`
		// Custom tags that are updated after a block including the tx is mined
		Mined         bool  `json:"txMined"`
		BlockIncluded int64 `json:"blockIncluded"`
		Failed        bool  `json:"txFailed"`
		Stealth       bool  `json:"txStealth"`
		// WIP
		// LocalLogs are the logs of the txs when executed against the local EVM (ganache fork that's updated every block)
		// Tags are to classify post logs, to track
		LocalLogs []string `json:"localLogs"`
		Tags      []string `json:"tags"`
	}{}
	body.TimeSeen = time.Now().Unix()
	body.Hash = tx.Hash().Hex()
	body.Type = "miscTx"
	body.From = getTxSenderAddress(tx, client)
	body.To = tx.To().Hex()
	body.Value = formatEthWeiToEther(tx.Value())
	body.Nonce = tx.Nonce()
	formattedGas := new(big.Int).SetUint64(tx.Gas())
	body.Gas = formatEthWeiToEther(formattedGas)
	body.GasPrice = formatEthWeiToEther(tx.GasPrice())
	body.Mined = isTxMined(tx.Hash().Hex(), client)
	body.Failed = hasTxFailed(tx.Hash().Hex(), client)
	if isStealth {
		body.BlockIncluded = getBlockNoByTxHash(tx.Hash().Hex(), client)
	}
	body.Stealth = isStealth
	jsonBytes, _ := json.Marshal(body)
	// Set up the request object.
	req := esapi.IndexRequest{
		Index:   "transactions",
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

func handleEdgeTx(tx *types.Transaction, client *ethclient.Client, isStealth bool) {
	// Connect to our es client
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error connecting to es client: %s", err)
	}
	// Start building a document
	body := struct {
		// Time the tx was discovered and other details (for data analysis + Kibana)
		TimeSeen int64  `json:"timeFirstDiscovered"`
		Hash     string `json:"txHash"`
		// Classifcation and other info
		Type     string  `json:"txType"`
		From     string  `json:"from"`
		To       string  `json:"to"`
		Value    float64 `json:"txValue"`
		Nonce    uint64  `json:"nonce"`
		GasPrice float64 `json:"gasPrice"`
		Gas      float64 `json:"gas"`
		// Custom tags that are updated after a block including the tx is mined
		Mined         bool  `json:"txMined"`
		BlockIncluded int64 `json:"blockIncluded"`
		Failed        bool  `json:"txFailed"`
		Stealth       bool  `json:"txStealth"`
		// WIP
		// LocalLogs are the logs of the txs when executed against the local EVM (ganache fork that's updated every block)
		// Tags are to classify post logs, to track
		LocalLogs []string `json:"localLogs"`
		Tags      []string `json:"tags"`
	}{}
	body.TimeSeen = time.Now().Unix()
	body.Hash = tx.Hash().Hex()
	body.Type = "edgeTx"
	body.From = getTxSenderAddress(tx, client)
	body.To = tx.To().Hex()
	body.Value = formatEthWeiToEther(tx.Value())
	body.Nonce = tx.Nonce()
	formattedGas := new(big.Int).SetUint64(tx.Gas())
	body.Gas = formatEthWeiToEther(formattedGas)
	body.GasPrice = formatEthWeiToEther(tx.GasPrice())
	body.Mined = isTxMined(tx.Hash().Hex(), client)
	body.Failed = hasTxFailed(tx.Hash().Hex(), client)
	if isStealth {
		body.BlockIncluded = getBlockNoByTxHash(tx.Hash().Hex(), client)
	}
	body.Stealth = isStealth
	jsonBytes, _ := json.Marshal(body)
	// Set up the request object.
	req := esapi.IndexRequest{
		Index:   "transactions",
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
