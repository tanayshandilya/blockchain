package main

import (
	"fmt"
	"os"

	"github.com/tanayshandilya/blockchain/core"
	"github.com/tanayshandilya/blockchain/core/encoding"
)

func main() {
	blockChain := new(core.BlockChain)
	blockChain.Initialize()

	json, err := encoding.JsonEncode(blockChain, true)

	if err != nil {
		fmt.Println(err)
	}

	os.WriteFile("./build/ledger.json", json, 0644)

}
