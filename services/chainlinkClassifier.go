package services

import (
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	. "github.com/logrusorgru/aurora"
	"github.com/taarushv/helios/contracts/chainlinkACA"
)

// Chainlink contract where oracles send their price updates
// All contracts access latest weighed price data from this contract (via EACAggregatorProxy)
var accessControlledAggregatorAbi, _ = abi.JSON(strings.NewReader(chainlinkACA.ChainlinkACAABI))

// Function signature to catch the oracle updates
var submitOraclePriceUpdate = [4]byte{0x20, 0x2e, 0xe0, 0xed}

type chainlinkOracleACA struct {
	pair            string // "ETH/USD"
	contractAddress string // "0xf00ba7..."
}

var priceFeedACAList = []chainlinkOracleACA{
	chainlinkOracleACA{
		pair:            "ADA/USD",
		contractAddress: "0xf94800E6e36b0dc860F6f31e7cDf1086099E8c0E",
	},
	chainlinkOracleACA{
		pair:            "AUD/USD",
		contractAddress: "0x3A33c0eFD0EB8fd38a6E1904dF1E32f95F67616b",
	},
	chainlinkOracleACA{
		pair:            "BAT/ETH",
		contractAddress: "0x3146392934Da3AE09447CD7Fe4061d8aa96B50ae",
	},
	chainlinkOracleACA{
		pair:            "BCH/USD",
		contractAddress: "0x744704c31a2E46AD60c7CDf0212933B4c4c2c9eC",
	},
	chainlinkOracleACA{
		pair:            "BNB/USD",
		contractAddress: "0x90888CDDaD598570c6eDC443eee9aaDB63cDA3C4",
	},
	chainlinkOracleACA{
		pair:            "BNT/ETH",
		contractAddress: "0x0A3ec7050884F00B5D9b11131De59DCed0fAFeDB",
	},
	chainlinkOracleACA{
		pair:            "BNT/USD",
		contractAddress: "0x42Dec4a0882756497DB9843a556a55dcd70b1995",
	},
	chainlinkOracleACA{
		pair:            "BTC/ETH",
		contractAddress: "0xbD72DA70007E47AAf1BBD84918675392cf6885F7",
	},
	chainlinkOracleACA{
		pair:            "BTC/USD",
		contractAddress: "0xF570deEffF684D964dc3E15E1F9414283E3f7419",
	},
	chainlinkOracleACA{
		pair:            "BUSD/ETH",
		contractAddress: "0x661BE809784E094eA70F980939Cf3f09337A3178",
	},
	chainlinkOracleACA{
		pair:            "BZRX/ETH",
		contractAddress: "0xb58b218365CD12AD765bA8A9F88E881D2b2a01C2",
	},
	chainlinkOracleACA{
		pair:            "CHF/USD",
		contractAddress: "0xdf005CaD29AAC8b1170960807f99B62aaeD1bb0a",
	},
	chainlinkOracleACA{
		pair:            "COMP/USD",
		contractAddress: "0xdbd020CAeF83eFd542f4De03e3cF0C28A4428bd5",
	},
	chainlinkOracleACA{
		pair:            "DAI/ETH",
		contractAddress: "0xd866A07Dea5Ee3c093e21d33660b5579C21F140b",
	},
	chainlinkOracleACA{
		pair:            "DAI/USD",
		contractAddress: "0xFe16F630Eb0Ca70661B071360701abf980126d3e",
	},
	chainlinkOracleACA{
		pair:            "ENJ/ETH",
		contractAddress: "0x20aff4833e5D261bB34BC3980d88aD17A3FE90Dc",
	},
	chainlinkOracleACA{
		pair:            "ETC/USD",
		contractAddress: "0x41306Eb5fC11A68C284c19Ba3B9510c0252E0a34",
	},
	chainlinkOracleACA{
		pair:            "ETH/USD",
		contractAddress: "0x00c7A37B03690fb9f41b5C5AF8131735C7275446",
	},
	chainlinkOracleACA{
		pair:            "ETH/XDR",
		contractAddress: "0x460DE59c7768e7ff17939F01Cb84F965c4E88266",
	},
	chainlinkOracleACA{
		pair:            "EUR/USD",
		contractAddress: "0x8f71c9c583248A11CAcBbC8FD0D5dFa483D3b109",
	},
	chainlinkOracleACA{
		pair:            "FTM/ETH",
		contractAddress: "0x3aaFb0E5b57bb19D02FcB6656059B34d8E79471f",
	},
	chainlinkOracleACA{
		pair:            "FTSE/GBP",
		contractAddress: "0xc95B41df94F3890122B2bcEf9005AFDe17773dB2",
	},
	chainlinkOracleACA{
		pair:            "FastGas/Gwei",
		contractAddress: "0xca947C9ddF31EE6c2E994EFB794Fdb0819AEEeD0",
	},
	chainlinkOracleACA{
		pair:            "GBP/USD",
		contractAddress: "0x3a6e27b663593E34a7FB80bA9544d9E8BAbdF001",
	},
	chainlinkOracleACA{
		pair:            "JPY/USD",
		contractAddress: "0x87CFEA02C8322653a7335C6f72Be19ce54ECbFb5",
	},
	chainlinkOracleACA{
		pair:            "KNC/ETH",
		contractAddress: "0x075Fe11b3Dd9c605f7fd09FF9310e3E37baaBC9e",
	},
	chainlinkOracleACA{
		pair:            "KNC/USD",
		contractAddress: "0xa811Ff165b082c0507Ce9a5a660Fb3D7eEeCb88A",
	},
	chainlinkOracleACA{
		pair:            "LEND/ETH",
		contractAddress: "0x0a2539a614E28ddDD02ebF396D2611B4f7d552FC",
	},
	chainlinkOracleACA{
		pair:            "LEND/USD",
		contractAddress: "0x0227fb846b48e209d56D79b0A3109FdA561db821",
	},
	chainlinkOracleACA{
		pair:            "LINK/ETH",
		contractAddress: "0x7E6C635d6A53B5033D1B0ceE84ecCeA9096859e4",
	},
	chainlinkOracleACA{
		pair:            "LINK/USD",
		contractAddress: "0x8cDE021F0BfA5f82610e8cE46493cF66AC04Af53",
	},
	chainlinkOracleACA{
		pair:            "LRC/ETH",
		contractAddress: "0x174b0B72ca036b0b64F190Ed83630751195D3362",
	},
	chainlinkOracleACA{
		pair:            "LTC/USD",
		contractAddress: "0x3F2d1Ff4930318B5a7c301E1bf7e703DcF6D83E3",
	},
	chainlinkOracleACA{
		pair:            "MANA/ETH",
		contractAddress: "0x3162C2De0C254B97d869a070929B518b5B9B56B3",
	},
	chainlinkOracleACA{
		pair:            "MKR/ETH",
		contractAddress: "0x204A6FE11De66aa463879f47F3533Dd87d47020D",
	},
	chainlinkOracleACA{
		pair:            "N225/JPY",
		contractAddress: "0x4Fa0655c09E0b5B2F50F1bd861B2d9BC63ccBBCB",
	},
	/*
		chainlinkOracleACA{//older?
			pair:"OXT/USD",
			contractAddress:"0x11eF34572CcaB4c85f0BAf03c36a14e0A9C8C7eA",
		}
	*/
	chainlinkOracleACA{
		pair:            "REN/ETH",
		contractAddress: "0x1A53bF1BFfFB7A2B33e1931D33423C7C94f675ee",
	},
	chainlinkOracleACA{
		pair:            "REN/USD",
		contractAddress: "0xD286AF227B7b0695387E279B9956540818B1dc2a",
	},
	chainlinkOracleACA{
		pair:            "REP/ETH",
		contractAddress: "0x8e1BB728b37832754A260D99B5467fE6d164c068",
	},
	chainlinkOracleACA{
		pair:            "SNX/ETH",
		contractAddress: "0x93d7bBf4CF42bDA5Bb86EaDFad09271040cc10e8",
	},
	chainlinkOracleACA{
		pair:            "SNX/USD",
		contractAddress: "0xC8DB8d5869510Bb1FCd3Bd7C7624c1b49c652ef8",
	},
	chainlinkOracleACA{
		pair:            "SUSD/ETH",
		contractAddress: "0x060f728deB96875F992C97414eFf2B3ef6c58EC7",
	},
	chainlinkOracleACA{
		pair:            "SXP/USD",
		contractAddress: "0x4A75bfa5B740263e88655a3e1e78892bfc7b036a",
	},
	chainlinkOracleACA{
		pair:            "TUSD/ETH",
		contractAddress: "0x0c632eC5982e3A8dC116a02ebA7A419efec170B1",
	},
	chainlinkOracleACA{
		pair:            "USDC/ETH",
		contractAddress: "0x00d02526CA08488342aB634de3B2d0050ecC7f60",
	},
	chainlinkOracleACA{
		pair:            "USDT/ETH",
		contractAddress: "0x1058a82C25F55aB8ab0cE717F3e6e164E80f1A0B",
	},
	chainlinkOracleACA{
		pair:            "WOM/ETH",
		contractAddress: "0xd8F46C7e5f9a0DCFac79a79763634FCE9302985f",
	},
	chainlinkOracleACA{
		pair:            "XAG/USD",
		contractAddress: "0xF320E19B2ED82F1B226b006cD43FE600FEA56615",
	},
	chainlinkOracleACA{
		pair:            "XAU/USD",
		contractAddress: "0x06A7689149cf04DacFDE555d1e1EAD7dD7370316",
	},
	chainlinkOracleACA{
		pair:            "XHV/USD",
		contractAddress: "0x9c6b454D42a69088719fF8B2Ab30C04808c8D061",
	},
	chainlinkOracleACA{
		pair:            "XRP/USD",
		contractAddress: "0x75Ed2f61837c3D9Ef1BF0af4DB84664DC6fe56bC",
	},
	chainlinkOracleACA{
		pair:            "XTZ/USD",
		contractAddress: "0x7391BB54a24719DA7DD81c2E5176cf954D7f7635",
	},
	chainlinkOracleACA{
		pair:            "YFI/ETH",
		contractAddress: "0x4A03707a1bfeFc2836A69B1A6A6bd752270041A9",
	},
	chainlinkOracleACA{
		pair:            "ZRX/ETH",
		contractAddress: "0xE03b49682965A1EB5230D41f96E10896dc563F0D",
	},
	chainlinkOracleACA{
		pair:            "sDEFI/USD",
		contractAddress: "0x25367741a23464b41B4aB978Bd8092d56a3590C0",
	},
}

type chainlinkOraclePriceUpdate struct {
	Oracle                                         string  `json:"oracleNodeAddress"`     // EOA address of the individual oracle
	PairDescription                                string  `json:"oraclePairDescription"` // "ETH/USD", "BTC/USD" etc
	RoundId                                        int64   `json:"oracleRoundId"`
	Submission                                     float64 `json:"oraclePriceSubmission"` // New price added to the aggregator
	CurrentPrice                                   float64 `json:"oracleCurrentPrice"`
	NextPriceIfExecutedInIsolation                 float64 `json:"oracleNextPriceIfExecutedInIsolation"`
	NextPriceIfExecutedWithOtherOracleTxsInMempool float64 `json:"oracleNextPriceIfExecutedWithOtherOracleTxsInMempool"`
}

// This function serves two purposes
// 1) Check if the list of chainlink oracle pricefeeds (that I bootstrapped in `priceFeedACAList`) are valid
// 2) To run on init to check if all the pricefeeds are valid (useful to check sanity of ACA's when we add new oracles)
func ChainlinkACAListSanityCheck(client *ethclient.Client) bool {
	for _, priceFeed := range priceFeedACAList {
		// Create a contract instance to check if it's valid
		aggregatorInstance, _ := chainlinkACA.NewChainlinkACA(common.HexToAddress(priceFeed.contractAddress), client)
		aggregatorDesc, _ := aggregatorInstance.Description(nil)
		trimmedAggregatorDesc := strings.ReplaceAll(aggregatorDesc, " ", "")
		if priceFeed.pair != trimmedAggregatorDesc {
			log.Fatalln("Chainlink ACA oracle list sanity checked failed. Pair does not exist on mainnet:", priceFeed.pair)
			return false
		}
	}
	return true
}

// Core method to identify and classify oracle updates
func handleChainlinkOracleUpdate(tx *types.Transaction, client *ethclient.Client, isStealth bool, fullMode bool) {

	// Check if it's a legit chainlink oracle update (i.e pre-approved link contracts + check sender against their respective oracles)
	// TODO: Execute against ganache fork

	ACAInstance, _ := chainlinkACA.NewChainlinkACA(*tx.To(), client)
	pairDescription, _ := ACAInstance.Description(nil)
	pairDecimals, _ := ACAInstance.Decimals(nil)
	pairCurrentPrice, _ := ACAInstance.LatestAnswer(nil)

	final := chainlinkOraclePriceUpdate{
		Oracle:                         getTxSenderAddress(tx, client),
		PairDescription:                pairDescription,
		RoundId:                        new(big.Int).SetBytes((tx.Data()[4:36])).Int64(),
		Submission:                     formatChainlinkOraclePrice(new(big.Int).SetBytes((tx.Data()[36:68])), pairDecimals),
		CurrentPrice:                   formatChainlinkOraclePrice(pairCurrentPrice, pairDecimals),
		NextPriceIfExecutedInIsolation: 0, // TODO: What will the next price be if this oracle update tx is executed
		NextPriceIfExecutedWithOtherOracleTxsInMempool: 0, // Same as above but if *all* the pending oracle updates in the mempool are executed
	}
	fmt.Println()
	fmt.Println(Green("New TX: Chainlink Oracle Update"))
	fmt.Println("Hash: ", tx.Hash().Hex(), " Pair: ", final.PairDescription)
	fmt.Println("Current Price: ", final.CurrentPrice, " Submission: ", final.Submission)
	if fullMode {
		handleLinkOracleUpdate(tx, client, isStealth, final)
	}
}
