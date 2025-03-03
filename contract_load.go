package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"go-learning/base"
	"go-learning/store"
)

func mainLoad() {
	client := base.Client()
	address := common.HexToAddress("0x85eb059056041a999e6dc235af4a2B9636b758c4")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")
	_ = instance
}
