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
	"github.com/taarushv/helios/contracts/uniswap"
)

// UNI related functions
// Bytes bc we want to be blazin' fast and avoid conversion hassles
var swapExactETHForTokens = [4]byte{0x7f, 0xf3, 0x6a, 0xb5}
var swapExactTokensForETH = [4]byte{0x18, 0xcb, 0xaf, 0xe5}
var swapExactTokensForTokens = [4]byte{0x38, 0xed, 0x17, 0x39}
var swapETHForExactTokens = [4]byte{0xfb, 0x3b, 0xdb, 0x41}
var addLiquidityETH = [4]byte{0xf3, 0x05, 0xd7, 0x19}
var swapTokensForExactETH = [4]byte{0x4a, 0x25, 0xd9, 0x4a}
var swapTokensForExactTokens = [4]byte{0x88, 0x03, 0xdb, 0xee}
var swapExactTokensForETHSupportingFeeOnTransferTokens = [4]byte{0x79, 0x1a, 0xc9, 0x47} // For fee-on-transfer/deflationary ERC20s
var removeLiquidityETHWithPermit = [4]byte{0xde, 0xd9, 0x38, 0x2a}                       // For remove liquidity calls coming from proxy yield farms
var addLiquidity = [4]byte{0xe8, 0xe3, 0x37, 0x00}
var swapExactTokensForTokensSupportingFeeOnTransferTokens = [4]byte{0x5c, 0x11, 0xd7, 0x95}
var removeLiquidityETH = [4]byte{0x02, 0x75, 0x1c, 0xec}
var removeLiquidityWithPermit = [4]byte{0x21, 0x95, 0x99, 0x5c}
var removeLiquidity = [4]byte{0xba, 0xa2, 0xab, 0xde}
var swapExactETHForTokensSupportingFeeOnTransferTokens = [4]byte{0xb6, 0xf9, 0xde, 0x95}
var removeLiquidityETHWithPermitSupportingFeeOnTransferTokens = [4]byte{0x5b, 0x0d, 0x59, 0x84}
var removeLiquidityETHSupportingFeeOnTransferTokens = [4]byte{0xaf, 0x29, 0x79, 0xeb}

// standard ABI
var routerAbi, _ = abi.JSON(strings.NewReader(uniswap.UniswapABI))

// Relevant types
type UniswapETHToTokenParsedInput struct {
	AmountOutMin *big.Int
	Path         []common.Address
	Deadline     *big.Int
	To           common.Address
}

type UniswapTokenToETHParsedInput struct {
	AmountIn     *big.Int
	AmountOutMin *big.Int
	Path         []common.Address
	Deadline     *big.Int
	To           common.Address
}

type UniswapTokenToTokenParsedInput struct {
	AmountIn     *big.Int
	AmountOutMin *big.Int
	Path         []common.Address
	Deadline     *big.Int
	To           common.Address
}

type UniswapETHToExactTokensInput struct {
	AmountOut *big.Int
	Path      []common.Address
	Deadline  *big.Int
	To        common.Address
}

type UniswapTokensForExactETHInput struct {
	AmountOut   *big.Int
	AmountInMax *big.Int
	Path        []common.Address
	Deadline    *big.Int
	To          common.Address
}

type UniswapTokensForExactTokensInput struct {
	AmountOut   *big.Int
	AmountInMax *big.Int
	Path        []common.Address
	Deadline    *big.Int
	To          common.Address
}

type UniswapExactTokensForETHSupportingFeeOnTransferTokensInput struct {
	AmountIn     *big.Int
	AmountOutMin *big.Int
	Path         []common.Address
	Deadline     *big.Int
	To           common.Address
}

type UniswapExactETHForTokensSupportingFeeOnTransferTokensInput struct {
	AmountOutMin *big.Int
	Path         []common.Address
	Deadline     *big.Int
	To           common.Address
}

type UniswapTradeFinal struct {
	AmountIn          float64  `json:"amountIn"`
	AmountOutMin      float64  `json:"amountOutMin"`
	Path              []string `json:"path"`
	Deadline          int64    `json:"deadline"`
	To                string   `json:"to"`
	OutputTokenSymbol string   `json:"outputTokenSymbol"`
	OutputTokenName   string   `json:""`
}

type UniswapAddLiquidityETHInput struct {
	Token              common.Address
	AmountTokenDesired *big.Int
	AmountETHMin       *big.Int
	AmountTokenMin     *big.Int
	Deadline           *big.Int
	To                 common.Address
}

type UniswapAddLiquidityETHFinalInput struct {
	TokenAddress             string  `json:"tokenAddress"`
	LiquidityProviderAddress string  `json:"liquidityProviderAddress"`
	Deadline                 int64   `json:"deadline"`
	AmountTokenDesired       float64 `json:"amountTokenDesired"`
	AmountTokenMin           float64 `json:"amountTokenMin"`
	AmountEthMin             float64 `json:"amountEthMin"`
}

type UniswapRemoveLiquidityETHWithPermit struct {
	// (address token,uint256 liquidity,uint256 amountTokenMin,uint256 amountETHMin,address to,uint256 deadline,bool approveMax,uint8 v,bytes32 r,bytes32 s )
	Token          common.Address
	Liquidity      *big.Int
	AmountTokenMin *big.Int
	AmountETHMin   *big.Int
	Deadline       *big.Int
	To             common.Address
	ApproveMax     bool
	V              uint8
	R              [32]byte
	S              [32]byte
}

type UniswapRemoveLiquidityETHInput struct {
	Token          common.Address
	Liquidity      *big.Int
	AmountTokenMin *big.Int
	AmountETHMin   *big.Int
	Deadline       *big.Int
	To             common.Address
}

type UniswapRemoveLiquidityETHSupportingFeeOnTransferTokens struct {
	Token          common.Address
	Liquidity      *big.Int
	AmountTokenMin *big.Int
	AmountETHMin   *big.Int
	Deadline       *big.Int
	To             common.Address
}

type UniswapRemoveLiquidityETHFinalInput struct {
	TokenAddress             string  `json:"tokenAddress"`
	LPTokenAmount            float64 `json:"lPTokenAmount"`
	AmountTokenMin           float64 `json:"amountTokenMin"`
	AmountETHMin             float64 `json:"amountEthMin"`
	Deadline                 int64   `json:"deadline"`
	LiquidityProviderAddress string  `json:"liquidityProviderAddress"`
}

type UniswapRemoveLiquidityETHWithPermitSupportingFeeOnTransferTokensInput struct {
	Token          common.Address
	Liquidity      *big.Int
	AmountTokenMin *big.Int
	AmountETHMin   *big.Int
	Deadline       *big.Int
	To             common.Address
	ApproveMax     bool
	V              uint8
	R              [32]byte
	S              [32]byte
}

type UniswapRemoveLiquidityWithPermitInput struct {
	TokenA     common.Address
	TokenB     common.Address
	Liquidity  *big.Int
	AmountAMin *big.Int
	AmountBMin *big.Int
	To         common.Address
	Deadline   *big.Int
	ApproveMax bool
	V          uint8
	R          [32]byte
	S          [32]byte
}

type UniswapRemoveLiquidityInput struct {
	TokenA     common.Address
	TokenB     common.Address
	Liquidity  *big.Int
	AmountAMin *big.Int
	AmountBMin *big.Int
	Deadline   *big.Int
	To         common.Address
}

type UniswapRemoveLiquidityFinalInput struct {
	TokenA     string  `json:"tokenA"`
	TokenB     string  `json:"tokenB"`
	AmountAMin float64 `json:"amountAMin"`
	AmountBMin float64 `json:"amountBMin"`
	Deadline   int64   `json:"deadline"`
}

type UniswapAddLiquidityInput struct {
	TokenA         common.Address
	TokenB         common.Address
	AmountADesired *big.Int
	AmountBDesired *big.Int
	AmountAMin     *big.Int
	AmountBMin     *big.Int
	Deadline       *big.Int
	To             common.Address
}

type UniswapAddLiquidityFinalInput struct {
	TokenAAddress            string  `json:"tokenAAddress"`
	TokenBAddress            string  `json:"tokenBAddress"`
	AmountADesired           float64 `json:"amountADesired"`
	AmountBDesired           float64 `json:"amountBDesired"`
	AmountAMin               float64 `json:"amountAMin"`
	AmountBMin               float64 `json:"amountBMin"`
	Deadline                 int64   `json:"deadline"`
	LiquidityProviderAddress string  `json:"liquidityProviderAddress"`
}

type UniswapExactTokensForTokensSupportingFeeOnTransferTokensInput struct {
	AmountIn     *big.Int
	AmountOutMin *big.Int
	Path         []common.Address
	Deadline     *big.Int
	To           common.Address
}

// Functions to trade tokens

func HandleSwapExactETHForTokens(tx *types.Transaction, client *ethclient.Client, isStealth bool, fullMode bool) {
	var trade UniswapETHToTokenParsedInput
	method, _ := routerAbi.MethodById((swapExactETHForTokens)[:])
	if err := method.Inputs.Unpack(
		&trade, tx.Data()[4:],
	); err != nil {
		log.Fatal(err)
	}
	tradePathString := make([]string, len(trade.Path))
	for i, s := range trade.Path {
		tradePathString[i] = s.Hex()
	}
	final := UniswapTradeFinal{
		AmountIn:          formatEthWeiToEther(tx.Value()),
		AmountOutMin:      formatERC20Decimals(trade.AmountOutMin, trade.Path[len(trade.Path)-1], client),
		Path:              tradePathString,
		Deadline:          trade.Deadline.Int64(),
		To:                trade.To.Hex(),
		OutputTokenSymbol: getTokenSymbol(trade.Path[len(trade.Path)-1], client),
		OutputTokenName:   getTokenName(trade.Path[len(trade.Path)-1], client),
	}
	if fullMode {
		handleUniFinalTrade(tx, client, isStealth, final)
	}

}

func HandleSwapExactTokensForETH(tx *types.Transaction, client *ethclient.Client, isStealth bool, fullMode bool) {
	var trade UniswapTokenToETHParsedInput
	method, _ := routerAbi.MethodById((swapExactTokensForETH)[:])
	if err := method.Inputs.Unpack(
		&trade, tx.Data()[4:],
	); err != nil {
		log.Fatal(err)
	}
	tradePathString := make([]string, len(trade.Path))
	for i, s := range trade.Path {
		tradePathString[i] = s.Hex()
	}
	final := UniswapTradeFinal{
		AmountIn:          formatERC20Decimals(trade.AmountIn, trade.Path[0], client),
		AmountOutMin:      formatEthWeiToEther(trade.AmountOutMin),
		Path:              tradePathString,
		Deadline:          trade.Deadline.Int64(),
		To:                trade.To.Hex(),
		OutputTokenSymbol: "ETH",
		OutputTokenName:   "Ether",
	}
	fmt.Println()
	fmt.Println(Red("New TX: Uniswap Trade"))
	fmt.Println("Hash: ", tx.Hash().Hex(), " Amount: ", final.AmountIn, getTokenSymbol(trade.Path[0], client), " For: ", final.OutputTokenSymbol)
	if fullMode {
		handleUniFinalTrade(tx, client, isStealth, final)
	}
}

func HandleSwapExactTokensForTokens(tx *types.Transaction, client *ethclient.Client, isStealth bool, fullMode bool) {
	var trade UniswapTokenToTokenParsedInput
	method, _ := routerAbi.MethodById((swapExactTokensForTokens)[:])
	if err := method.Inputs.Unpack(
		&trade, tx.Data()[4:],
	); err != nil {
		log.Fatal(err)
	}
	tradePathString := make([]string, len(trade.Path))
	for i, s := range trade.Path {
		tradePathString[i] = s.Hex()
	}
	final := UniswapTradeFinal{
		AmountIn:          formatERC20Decimals(trade.AmountIn, trade.Path[0], client),
		AmountOutMin:      formatERC20Decimals(trade.AmountOutMin, trade.Path[len(trade.Path)-1], client),
		Path:              tradePathString,
		Deadline:          trade.Deadline.Int64(),
		To:                trade.To.Hex(),
		OutputTokenSymbol: getTokenSymbol(trade.Path[len(trade.Path)-1], client),
		OutputTokenName:   getTokenName(trade.Path[len(trade.Path)-1], client),
	}
	fmt.Println()
	fmt.Println(Red("New TX: Uniswap Trade"))
	fmt.Println("Hash: ", tx.Hash().Hex(), " Amount: ", final.AmountIn, getTokenSymbol(trade.Path[0], client), " For: ", final.OutputTokenSymbol)
	if fullMode {
		handleUniFinalTrade(tx, client, isStealth, final)
	}
}

func HandleSwapETHForExactTokens(tx *types.Transaction, client *ethclient.Client, isStealth bool, fullMode bool) {
	var trade UniswapETHToExactTokensInput
	method, _ := routerAbi.MethodById((swapETHForExactTokens)[:])
	if err := method.Inputs.Unpack(
		&trade, tx.Data()[4:],
	); err != nil {
		log.Fatal(err)
	}
	tradePathString := make([]string, len(trade.Path))
	for i, s := range trade.Path {
		tradePathString[i] = s.Hex()
	}
	final := UniswapTradeFinal{
		AmountIn:          formatEthWeiToEther(tx.Value()),
		AmountOutMin:      formatERC20Decimals(trade.AmountOut, trade.Path[len(trade.Path)-1], client),
		Path:              tradePathString,
		Deadline:          trade.Deadline.Int64(),
		To:                trade.To.Hex(),
		OutputTokenSymbol: getTokenSymbol(trade.Path[len(trade.Path)-1], client),
		OutputTokenName:   getTokenName(trade.Path[len(trade.Path)-1], client),
	}
	fmt.Println()
	fmt.Println(Red("New TX: Uniswap Trade"))
	fmt.Println("Hash: ", tx.Hash().Hex(), " Amount: ", final.AmountIn, "ETH", " For: ", final.OutputTokenSymbol)
	if fullMode {
		handleUniFinalTrade(tx, client, isStealth, final)
	}
}

func HandleSwapTokensForExactEth(tx *types.Transaction, client *ethclient.Client, isStealth bool, fullMode bool) {
	var trade UniswapTokensForExactETHInput
	method, _ := routerAbi.MethodById((swapTokensForExactETH)[:])
	if err := method.Inputs.Unpack(
		&trade, tx.Data()[4:],
	); err != nil {
		log.Fatal(err)
	}
	tradePathString := make([]string, len(trade.Path))
	for i, s := range trade.Path {
		tradePathString[i] = s.Hex()
	}
	final := UniswapTradeFinal{
		AmountIn:          formatERC20Decimals(trade.AmountInMax, trade.Path[0], client),
		AmountOutMin:      formatEthWeiToEther(trade.AmountOut),
		Path:              tradePathString,
		Deadline:          trade.Deadline.Int64(),
		To:                trade.To.Hex(),
		OutputTokenSymbol: "ETH",
		OutputTokenName:   "Ether",
	}
	fmt.Println()
	fmt.Println(Red("New TX: Uniswap Trade"))
	fmt.Println("Hash: ", tx.Hash().Hex(), " Amount: ", final.AmountIn, getTokenSymbol(trade.Path[0], client), " For: ", final.OutputTokenSymbol)
	if fullMode {
		handleUniFinalTrade(tx, client, isStealth, final)
	}
}

func HandleSwapTokensForExactTokens(tx *types.Transaction, client *ethclient.Client, isStealth bool, fullMode bool) {
	var trade UniswapTokensForExactTokensInput
	method, _ := routerAbi.MethodById((swapTokensForExactTokens)[:])
	if err := method.Inputs.Unpack(
		&trade, tx.Data()[4:],
	); err != nil {
		log.Fatal(err)
	}
	tradePathString := make([]string, len(trade.Path))
	for i, s := range trade.Path {
		tradePathString[i] = s.Hex()
	}
	final := UniswapTradeFinal{
		AmountIn:          formatERC20Decimals(trade.AmountInMax, trade.Path[0], client),
		AmountOutMin:      formatERC20Decimals(trade.AmountOut, trade.Path[len(trade.Path)-1], client),
		Path:              tradePathString,
		Deadline:          trade.Deadline.Int64(),
		To:                trade.To.Hex(),
		OutputTokenSymbol: getTokenSymbol(trade.Path[len(trade.Path)-1], client),
		OutputTokenName:   getTokenName(trade.Path[len(trade.Path)-1], client),
	}
	fmt.Println()
	fmt.Println(Red("New TX: Uniswap Trade"))
	fmt.Println("Hash: ", tx.Hash().Hex(), " Amount: ", final.AmountIn, getTokenSymbol(trade.Path[0], client), " For: ", final.OutputTokenSymbol)
	if fullMode {
		handleUniFinalTrade(tx, client, isStealth, final)
	}
}

func HandleSwapExactTokensForETHSupportingFeeOnTransferTokens(tx *types.Transaction, client *ethclient.Client, isStealth bool, fullMode bool) {
	var trade UniswapExactTokensForETHSupportingFeeOnTransferTokensInput
	method, _ := routerAbi.MethodById((swapExactTokensForETHSupportingFeeOnTransferTokens)[:])
	if err := method.Inputs.Unpack(
		&trade, tx.Data()[4:],
	); err != nil {
		log.Fatal(err)
	}
	tradePathString := make([]string, len(trade.Path))
	for i, s := range trade.Path {
		tradePathString[i] = s.Hex()
	}
	final := UniswapTradeFinal{
		AmountIn:          formatERC20Decimals(trade.AmountIn, trade.Path[0], client),
		AmountOutMin:      formatEthWeiToEther(trade.AmountOutMin),
		Path:              tradePathString,
		Deadline:          trade.Deadline.Int64(),
		To:                trade.To.Hex(),
		OutputTokenSymbol: "ETH",
		OutputTokenName:   "Ether",
	}
	fmt.Println()
	fmt.Println(Red("New TX: Uniswap Trade"))
	fmt.Println("Hash: ", tx.Hash().Hex(), " Amount: ", final.AmountIn, getTokenSymbol(trade.Path[0], client), " For: ", final.OutputTokenSymbol)
	if fullMode {
		handleUniFinalTrade(tx, client, isStealth, final)
	}
}

func HandleSwapExactTokensForTokensSupportingFeeOnTransferTokens(tx *types.Transaction, client *ethclient.Client, isStealth bool, fullMode bool) {
	var trade UniswapExactTokensForTokensSupportingFeeOnTransferTokensInput
	method, _ := routerAbi.MethodById((swapExactTokensForTokensSupportingFeeOnTransferTokens)[:])
	if err := method.Inputs.Unpack(
		&trade, tx.Data()[4:],
	); err != nil {
		log.Fatal(err)
	}
	tradePathString := make([]string, len(trade.Path))
	for i, s := range trade.Path {
		tradePathString[i] = s.Hex()
	}
	final := UniswapTradeFinal{
		AmountIn:          formatERC20Decimals(trade.AmountIn, trade.Path[0], client),
		AmountOutMin:      formatERC20Decimals(trade.AmountOutMin, trade.Path[len(trade.Path)-1], client),
		Path:              tradePathString,
		Deadline:          trade.Deadline.Int64(),
		To:                trade.To.Hex(),
		OutputTokenSymbol: getTokenSymbol(trade.Path[len(trade.Path)-1], client),
		OutputTokenName:   getTokenName(trade.Path[len(trade.Path)-1], client),
	}
	fmt.Println()
	fmt.Println(Red("New TX: Uniswap Trade"))
	fmt.Println("Hash: ", tx.Hash().Hex(), " Amount: ", final.AmountIn, getTokenSymbol(trade.Path[0], client), " For: ", final.OutputTokenSymbol)
	if fullMode {
		handleUniFinalTrade(tx, client, isStealth, final)
	}
}

func HandleSwapExactETHForTokensSupportingFeeOnTransferTokens(tx *types.Transaction, client *ethclient.Client, isStealth bool, fullMode bool) {
	var trade UniswapExactETHForTokensSupportingFeeOnTransferTokensInput
	method, _ := routerAbi.MethodById((swapExactETHForTokensSupportingFeeOnTransferTokens)[:])
	if err := method.Inputs.Unpack(
		&trade, tx.Data()[4:],
	); err != nil {
		log.Fatal(err)
	}
	tradePathString := make([]string, len(trade.Path))
	for i, s := range trade.Path {
		tradePathString[i] = s.Hex()
	}
	final := UniswapTradeFinal{
		AmountIn:          formatEthWeiToEther(tx.Value()),
		AmountOutMin:      formatERC20Decimals(trade.AmountOutMin, trade.Path[len(trade.Path)-1], client),
		Path:              tradePathString,
		Deadline:          trade.Deadline.Int64(),
		To:                trade.To.Hex(),
		OutputTokenSymbol: getTokenSymbol(trade.Path[len(trade.Path)-1], client),
		OutputTokenName:   getTokenName(trade.Path[len(trade.Path)-1], client),
	}
	fmt.Println()
	fmt.Println(Red("New TX: Uniswap Trade"))
	fmt.Println("Hash: ", tx.Hash().Hex(), " Amount: ", final.AmountIn, getTokenSymbol(trade.Path[0], client), " For: ", final.OutputTokenSymbol)
	if fullMode {
		handleUniFinalTrade(tx, client, isStealth, final)
	}
}

// Functions to add liquidity

func HandleAddLiquidity(tx *types.Transaction, client *ethclient.Client, isStealth bool, fullMode bool) {
	var unpacked UniswapAddLiquidityInput
	method, _ := routerAbi.MethodById((addLiquidity)[:])
	if err := method.Inputs.Unpack(
		&unpacked, tx.Data()[4:],
	); err != nil {
		log.Fatal(err)
	}

	final := UniswapAddLiquidityFinalInput{
		TokenAAddress:            unpacked.TokenA.Hex(),
		TokenBAddress:            unpacked.TokenB.Hex(),
		LiquidityProviderAddress: unpacked.To.Hex(),
		Deadline:                 unpacked.Deadline.Int64(),
		AmountAMin:               formatERC20Decimals(unpacked.AmountAMin, unpacked.TokenA, client),
		AmountBMin:               formatERC20Decimals(unpacked.AmountBMin, unpacked.TokenB, client),
		AmountADesired:           formatERC20Decimals(unpacked.AmountADesired, unpacked.TokenA, client),
		AmountBDesired:           formatERC20Decimals(unpacked.AmountBDesired, unpacked.TokenB, client),
	}
	fmt.Println()
	fmt.Println(Red("New TX: Uniswap Add Liquidity"))
	fmt.Println("Hash: ", tx.Hash().Hex())
	if fullMode {
		handleUniAddLiq(tx, client, isStealth, final)
	}
}

func HandleAddLiquidityETH(tx *types.Transaction, client *ethclient.Client, isStealth bool, fullMode bool) {
	var addLiquidity UniswapAddLiquidityETHInput
	method, _ := routerAbi.MethodById((addLiquidityETH)[:])
	if err := method.Inputs.Unpack(
		&addLiquidity, tx.Data()[4:],
	); err != nil {
		log.Fatal(err)
	}

	final := UniswapAddLiquidityETHFinalInput{
		TokenAddress:             addLiquidity.Token.Hex(),
		LiquidityProviderAddress: addLiquidity.To.Hex(),
		Deadline:                 addLiquidity.Deadline.Int64(),
		AmountTokenDesired:       formatERC20Decimals(addLiquidity.AmountTokenDesired, addLiquidity.Token, client),
		AmountTokenMin:           formatERC20Decimals(addLiquidity.AmountTokenMin, addLiquidity.Token, client),
		AmountEthMin:             formatEthWeiToEther(addLiquidity.AmountETHMin),
	}
	fmt.Println(Red("New TX: Uniswap Add Liquidity"))
	fmt.Println("Hash: ", tx.Hash().Hex())
	if fullMode {
		handleUniAddETHLiq(tx, client, isStealth, final)
	}
}

// Functions to remove liquidity

func HandleRemoveLiquidityETHWithPermit(tx *types.Transaction, client *ethclient.Client, isStealth bool, fullMode bool) {
	var removeLiquidity UniswapRemoveLiquidityETHWithPermit
	method, _ := routerAbi.MethodById((removeLiquidityETH)[:])
	if err := method.Inputs.Unpack(
		&removeLiquidity, tx.Data()[4:],
	); err != nil {
		log.Fatal(err)
	}

	final := UniswapRemoveLiquidityETHFinalInput{
		TokenAddress:             removeLiquidity.Token.Hex(),
		LiquidityProviderAddress: removeLiquidity.To.Hex(),
		Deadline:                 removeLiquidity.Deadline.Int64(),
		LPTokenAmount:            formatEthWeiToEther(removeLiquidity.Liquidity), // UNI LP tokens have 18 decimals too
		AmountTokenMin:           formatERC20Decimals(removeLiquidity.AmountTokenMin, removeLiquidity.Token, client),
		AmountETHMin:             formatEthWeiToEther(removeLiquidity.AmountETHMin),
	}
	fmt.Println()
	fmt.Println(Red("New TX: Uniswap Remove Liquidity"))
	fmt.Println("Hash: ", tx.Hash().Hex())
	if fullMode {
		handleUniRemoveETHLiq(tx, client, isStealth, final)
	}

}

func HandleRemoveLiquidityETH(tx *types.Transaction, client *ethclient.Client, isStealth bool, fullMode bool) {
	var unpack UniswapRemoveLiquidityETHInput
	method, _ := routerAbi.MethodById((removeLiquidityETH)[:])
	if err := method.Inputs.Unpack(
		&unpack, tx.Data()[4:],
	); err != nil {
		log.Fatal(err)
	}

	final := UniswapRemoveLiquidityETHFinalInput{
		TokenAddress:             unpack.Token.Hex(),
		LiquidityProviderAddress: unpack.To.Hex(),
		Deadline:                 unpack.Deadline.Int64(),
		AmountTokenMin:           formatERC20Decimals(unpack.AmountTokenMin, unpack.Token, client),
		AmountETHMin:             formatEthWeiToEther(unpack.AmountETHMin),
		LPTokenAmount:            formatEthWeiToEther(unpack.Liquidity),
	}
	fmt.Println()
	fmt.Println(Red("New TX: Uniswap Remove Liquidity"))
	fmt.Println("Hash: ", tx.Hash().Hex())
	if fullMode {
		handleUniRemoveETHLiq(tx, client, isStealth, final)
	}
}

func HandleRemoveLiquidityWithPermit(tx *types.Transaction, client *ethclient.Client, isStealth bool, fullMode bool) {
	var unpack UniswapRemoveLiquidityWithPermitInput
	method, _ := routerAbi.MethodById((removeLiquidityWithPermit)[:])
	if err := method.Inputs.Unpack(
		&unpack, tx.Data()[4:],
	); err != nil {
		log.Fatal(err)
	}

	final := UniswapRemoveLiquidityFinalInput{
		TokenA:     unpack.TokenA.Hex(),
		TokenB:     unpack.TokenB.Hex(),
		AmountAMin: formatERC20Decimals(unpack.AmountAMin, unpack.TokenA, client),
		AmountBMin: formatERC20Decimals(unpack.AmountBMin, unpack.TokenB, client),
		Deadline:   unpack.Deadline.Int64(),
	}
	fmt.Println()
	fmt.Println(Red("New TX: Uniswap Remove Liquidity"))
	fmt.Println("Hash: ", tx.Hash().Hex())
	if fullMode {
		handleUniRemoveLiq(tx, client, isStealth, final)
	}
}

func HandleRemoveLiquidity(tx *types.Transaction, client *ethclient.Client, isStealth bool, fullMode bool) {
	var unpack UniswapRemoveLiquidityInput
	method, _ := routerAbi.MethodById((removeLiquidity)[:])
	if err := method.Inputs.Unpack(
		&unpack, tx.Data()[4:],
	); err != nil {
		log.Fatal(err)
	}

	final := UniswapRemoveLiquidityFinalInput{
		TokenA:     unpack.TokenA.Hex(),
		TokenB:     unpack.TokenB.Hex(),
		AmountAMin: formatERC20Decimals(unpack.AmountAMin, unpack.TokenA, client),
		AmountBMin: formatERC20Decimals(unpack.AmountBMin, unpack.TokenB, client),
		Deadline:   unpack.Deadline.Int64(),
	}
	fmt.Println()
	fmt.Println(Red("New TX: Uniswap Remove Liquidity"))
	fmt.Println("Hash: ", tx.Hash().Hex())
	if fullMode {
		handleUniRemoveLiq(tx, client, isStealth, final)
	}

}

func HandleRemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(tx *types.Transaction, client *ethclient.Client, isStealth bool, fullMode bool) {
	var unpack UniswapRemoveLiquidityETHWithPermitSupportingFeeOnTransferTokensInput
	method, _ := routerAbi.MethodById((removeLiquidityETH)[:])
	if err := method.Inputs.Unpack(
		&unpack, tx.Data()[4:],
	); err != nil {
		log.Fatal(err)
	}

	final := UniswapRemoveLiquidityETHFinalInput{
		TokenAddress:             unpack.Token.Hex(),
		LiquidityProviderAddress: unpack.To.Hex(),
		Deadline:                 unpack.Deadline.Int64(),
		AmountTokenMin:           formatERC20Decimals(unpack.AmountTokenMin, unpack.Token, client),
		AmountETHMin:             formatEthWeiToEther(unpack.AmountETHMin),
		LPTokenAmount:            formatEthWeiToEther(unpack.Liquidity), // UNI LP tokens have 18 decimals too
	}
	fmt.Println()
	fmt.Println(Red("New TX: Uniswap Remove Liquidity"))
	fmt.Println("Hash: ", tx.Hash().Hex())
	if fullMode {
		handleUniRemoveETHLiq(tx, client, isStealth, final)
	}
}

func HandleRemoveLiquidityETHSupportingFeeOnTransferTokens(tx *types.Transaction, client *ethclient.Client, isStealth bool, fullMode bool) {
	var unpack UniswapRemoveLiquidityETHSupportingFeeOnTransferTokens
	method, _ := routerAbi.MethodById((removeLiquidityETHSupportingFeeOnTransferTokens)[:])
	if err := method.Inputs.Unpack(
		&unpack, tx.Data()[4:],
	); err != nil {
		log.Fatal(err)
	}

	final := UniswapRemoveLiquidityETHFinalInput{
		TokenAddress:             unpack.Token.Hex(),
		LiquidityProviderAddress: unpack.To.Hex(),
		Deadline:                 unpack.Deadline.Int64(),
		AmountTokenMin:           formatERC20Decimals(unpack.AmountTokenMin, unpack.Token, client),
		AmountETHMin:             formatEthWeiToEther(unpack.AmountETHMin),
		LPTokenAmount:            formatEthWeiToEther(unpack.Liquidity), // UNI LP tokens have 18 decimals too
	}
	fmt.Println()
	fmt.Println(Red("New TX: Uniswap Remove Liquidity"))
	fmt.Println("Hash: ", tx.Hash().Hex())
	if fullMode {
		handleUniRemoveETHLiq(tx, client, isStealth, final)
	}
}

// Core method that determines the kind of uniswap trade the tx is
func handleUniswapTrade(tx *types.Transaction, client *ethclient.Client, isStealth bool, fullMode bool) {
	// Iterate through each function (ranked by popularity, https://bloxy.info/address/0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D)
	// Store data in the format we need
	txFunctionHash := [4]byte{}
	copy(txFunctionHash[:], tx.Data()[:4])
	switch txFunctionHash {
	case swapExactETHForTokens:
		HandleSwapExactETHForTokens(tx, client, isStealth, fullMode)
	case swapExactTokensForETH:
		HandleSwapExactTokensForETH(tx, client, isStealth, fullMode)
	case swapExactTokensForTokens:
		HandleSwapExactTokensForTokens(tx, client, isStealth, fullMode)
	case swapETHForExactTokens:
		HandleSwapETHForExactTokens(tx, client, isStealth, fullMode)
	case addLiquidityETH: // ADD LIQ
		HandleAddLiquidityETH(tx, client, isStealth, fullMode)
	case swapTokensForExactETH:
		HandleSwapTokensForExactEth(tx, client, isStealth, fullMode)
	case swapTokensForExactTokens:
		HandleSwapTokensForExactTokens(tx, client, isStealth, fullMode)
	case swapExactTokensForETHSupportingFeeOnTransferTokens:
		HandleSwapExactTokensForETHSupportingFeeOnTransferTokens(tx, client, isStealth, fullMode)
	case removeLiquidityETHWithPermit: // REMOVE LIQ
		HandleRemoveLiquidityETHWithPermit(tx, client, isStealth, fullMode)
	case addLiquidity: // ADD LIQ
		HandleAddLiquidity(tx, client, isStealth, fullMode)
	case swapExactTokensForTokensSupportingFeeOnTransferTokens:
		HandleSwapExactTokensForTokensSupportingFeeOnTransferTokens(tx, client, isStealth, fullMode)
	case removeLiquidityETH: // REMOVE LIQ
		HandleRemoveLiquidityETH(tx, client, isStealth, fullMode)
	case removeLiquidityWithPermit: // REMOVE LIQ
		HandleRemoveLiquidityWithPermit(tx, client, isStealth, fullMode)
	case removeLiquidity: // REMOVE LIQ
		HandleRemoveLiquidity(tx, client, isStealth, fullMode)
	case swapExactETHForTokensSupportingFeeOnTransferTokens:
		HandleSwapExactETHForTokensSupportingFeeOnTransferTokens(tx, client, isStealth, fullMode)
	case removeLiquidityETHWithPermitSupportingFeeOnTransferTokens: // REMOVE LIQ
		HandleRemoveLiquidityETHWithPermitSupportingFeeOnTransferTokens(tx, client, isStealth, fullMode)
	case removeLiquidityETHSupportingFeeOnTransferTokens: // REMOVE LIQ
		HandleRemoveLiquidityETHSupportingFeeOnTransferTokens(tx, client, isStealth, fullMode)
	}
}
