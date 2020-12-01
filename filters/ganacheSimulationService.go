package services

// TODO: Migrate panther from JS to Golang
/*

import (
	"context"
	"fmt"
	"log"
	"os/exec"

	"github.com/ethereum/go-ethereum/ethclient"
)

func GetGanacheBlockNo() {
	ganacheForkClient, err := ethclient.Dial("http://127.0.0.1:1337")
	if err != nil {
		log.Fatal(err)
	}
	bN, err := ganacheForkClient.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bN) // 25893180161173005034
}

func KillLocalGanacheInstance() {
	killCommand := "lsof -n -i4TCP:1337 | grep LISTEN | awk '{ print $2 }' | xargs kill"
	err := exec.Command("bash", "-c", killCommand).Run()
	if err != nil {
		fmt.Println("Error refreshing the local ganache instance")
	}
}

func StartLocalGanaceInstance() {
	var cmd = exec.Command("bash", "-c", "ganache-cli --port 1337 --fork http://0.0.0.0:8545 --deterministic")
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Starting Ganache fork @ block#")
	GetGanacheBlockNo()
}
*/
