package filters

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taarushv/helios/contracts/erc20"
)

// WIP
type LogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
}

func TransfersInBlock(block *types.Block, client *ethclient.Client) {
	fmt.Println(block.Number())
	for _, tx := range block.Transactions() {
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			fmt.Println("Failed to fetch tx logs")
			log.Fatal(err)
		}
		contractAbi, err := abi.JSON(strings.NewReader(string(erc20.Erc20ABI)))

		if err != nil {
			log.Fatal(err)
		}
		logTransferSig := []byte("Transfer(address,address,uint256)")
		logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
		for _, vLog := range receipt.Logs {
			if vLog.Topics[0].Hex() == logTransferSigHash.Hex() {
				fmt.Println("ERC20 Transfer")
				var event = struct {
					Value *big.Int
				}{}
				err := contractAbi.Unpack(&event, "Transfer", vLog.Data)
				if err != nil {
					fmt.Println("ooops")
					fmt.Println(err)
				}
				fmt.Println("From", common.BytesToAddress(vLog.Topics[1].Bytes()).Hex())
				fmt.Println("To", common.BytesToAddress(vLog.Topics[2].Bytes()).Hex())
				fmt.Println("Tokens", event.Value.String())
			}
		}
	}

}
