package filters

import (
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// WIP

// Oracle for the BTC/USDC perp funding rate on dydx
var PBTCUSDCFundingRateOracleAddress = "0x4525D2B71f7f018c9EBddFcD336852A85404e75B"
var PLINKUSDCFundingRateOracleAddress = "0x8B90515C7a99b7Edd97702c04d1E3666281De1B0"

func HandleDydxFundingRateUpdate(tx *types.Transaction, client *ethclient.Client) {
	fmt.Println(PBTCUSDCFundingRateOracleAddress)
	//ef460e36
	//00000000000000000000000000000000000000000000000000000006228bd5b20000000000000000000000000000000000000000000000000000000000000001
	// FR Log 00000000000000000000005F8A503601000000000000000000000006228BD5B2
}
