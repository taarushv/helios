package services

import (
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"
	"unsafe"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

// Initiate a client to handle all communication with the Ethereum protocol
// You can either connect to a local geth ipc client or use infura
var (
	infuraEndpoint = goDotEnvVariable("INFURA_WS_URL")
	ipcPath        = goDotEnvVariable("GETH_IPC_PATH")
)

var clientType = flag.String("client", ipcPath, "Gateway to the ethereum protocol")

func DialInfuraClient() *ethclient.Client {
	client, err := ethclient.Dial(infuraEndpoint)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func DialLocalClient() *ethclient.Client {
	// Connect via geth's ethclient
	client, err := ethclient.Dial(ipcPath)
	if err != nil {
		fmt.Println("Error connecting to geth ipc socket")
		log.Fatalln(err)
	}
	return client
}

// ethclient NOT rpcclient
func GetCurrentClient() *ethclient.Client {
	flag.Parse()
	if *clientType == "infura" {
		return DialInfuraClient()
	} else {
		return DialLocalClient()
	}
}
func InitRPCClient() *rpc.Client {
	var clientValue reflect.Value
	clientValue = reflect.ValueOf(GetCurrentClient()).Elem()
	fieldStruct := clientValue.FieldByName("c")
	clientPointer := reflect.NewAt(fieldStruct.Type(), unsafe.Pointer(fieldStruct.UnsafeAddr())).Elem()
	finalClient, _ := clientPointer.Interface().(*rpc.Client)
	return finalClient
}
