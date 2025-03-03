package base

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Client() *ethclient.Client {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/xx")
	if err != nil {
		log.Fatal(err)
		return client
	}
	return client
}

func WsClient() *ethclient.Client {
	client, err := ethclient.Dial("wss://eth-sepolia.g.alchemy.com/v2/xx")
	if err != nil {
		log.Fatal(err)
		return client
	}
	return client
}

func Address() common.Address {
	address := common.HexToAddress("0x3A2287539c04b28868B65215129E60eF8Ad2a720")
	return address
}

func ContractAddress() common.Address {
	address := common.HexToAddress("0x76967e1B238bb048E4Da500B39ACB426cf199032")
	return address
}

func PrivateKey() *ecdsa.PrivateKey {
	privateKey, err := crypto.HexToECDSA("xx")
	if err != nil {
		log.Fatal(err)
	}
	return privateKey
}

func TargetAddress() common.Address {
	address := common.HexToAddress("0x062437AfA9a90Ae56655F65D61486A812d06Dfbf")
	return address
}
