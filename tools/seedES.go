package tools

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

// All the functions to seed elastic search with relevant required data
// Index 4bytes

func Seed4Bytes() {
	absPath, _ := filepath.Abs("./data/fn-signatures")

	files, err := ioutil.ReadDir(absPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		fileValue, _ := filepath.Abs("./data/fn-signatures/" + f.Name())
		buf, _ := ioutil.ReadFile(fileValue)
		value := string(buf)
		insert4BytesKV(f.Name(), value)

	}
}

func insert4BytesKV(fnSignature string, fnIdentifier string) {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error connecting to es client: %s", err)
	}
	body := struct {
		Signature  string `json:"fnSignature"`
		Identifier string `json:"fnIdentifier"`
	}{}
	body.Signature = fnSignature
	body.Identifier = fnIdentifier
	jsonBytes, _ := json.Marshal(body)
	// Set up the request object.
	req := esapi.IndexRequest{
		Index:   "4bytes",
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
