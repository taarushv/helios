package filters

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"reflect"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taarushv/helios/contracts/erc20"
)

func GetTokenSymbol(tokenAddress common.Address, client *ethclient.Client) string {
	tokenIntance, _ := erc20.NewErc20(tokenAddress, client)
	sym, _ := tokenIntance.Symbol(nil)
	return sym
}
func GetTokenName(tokenAddress common.Address, client *ethclient.Client) string {
	tokenIntance, _ := erc20.NewErc20(tokenAddress, client)
	name, _ := tokenIntance.Name(nil)
	return name
}
func getTxSenderAddress(tx *types.Transaction, client *ethclient.Client) string {
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	msg, _ := tx.AsMessage(types.NewEIP155Signer(chainID))
	return msg.From().Hex()
}

// Format # of tokens transferred into required float
func FormatERC20Decimals(tokensSent *big.Int, tokenAddress common.Address, client *ethclient.Client) float64 {
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

func FormatEthWeiToEther(etherAmount *big.Int) float64 {
	var base, exponent = big.NewInt(10), big.NewInt(18)
	denominator := base.Exp(base, exponent, nil)
	// Convert to float for precision
	tokensSentFloat := new(big.Float).SetInt(etherAmount)
	denominatorFloat := new(big.Float).SetInt(denominator)
	// Divide and return the final result
	final, _ := new(big.Float).Quo(tokensSentFloat, denominatorFloat).Float64()
	return final
}

func itemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)

	if arr.Kind() != reflect.Array {
		panic("Invalid data-type")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

func GetFnIdentifierBySignature(fnSignature string) string {
	var r map[string]interface{}
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error connecting to es client: %s", err)
	}
	// 3. Search for the indexed documents
	//
	// Build the request body.
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"fnSignature": fnSignature,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	// Perform the search request.
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("4bytes"),
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
		var value = (hit.(map[string]interface{})["_source"])
		var fnIdentifier = (value.(map[string]interface{})["fnIdentifier"])
		final := fmt.Sprintf("%v", fnIdentifier)
		return final
	}
	return ""
}

func IsTxMined(txHash string, client *ethclient.Client) bool {
	finalTxHash := common.HexToHash(txHash)
	_, isPending, err := client.TransactionByHash(context.Background(), finalTxHash)
	if err != nil {
		log.Fatal(err)
	}
	return !isPending
}

func HasTxFailed(txHash string, client *ethclient.Client) bool {
	if IsTxMined(txHash, client) {
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
