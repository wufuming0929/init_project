package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"go-learning/base"
	"log"

	"github.com/ethereum/go-ethereum/common"
)

func mainByte() {
	client := base.Client()

	contractAddress := common.HexToAddress("0x85eb059056041a999e6dc235af4a2B9636b758c4")
	bytecode, err := client.CodeAt(context.Background(), contractAddress, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hex.EncodeToString(bytecode)) // 60806...10029
}
