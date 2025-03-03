package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func mainW() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHEX := hexutil.Encode(privateKeyBytes)
	//我们现在可以使用 go-ethereum hexutil 包将它转换为十六进制字符串，该包提供了一个带有字节切片的 Encode 方法。 然后我们在十六进制编码之后删除“0x”。
	fmt.Println("privateKey:", privateKeyHEX, "=====>result:", privateKeyHEX[2:]) // 0xfad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	publicKeyHEX := hexutil.Encode(publicKeyBytes)
	//将其转换为十六进制的过程与我们使用转化私钥的过程类似。 我们剥离了 0x 和前 2 个字符 04，它始终是 EC 前缀，不是必需的。
	fmt.Println("publicKey:", publicKeyHEX, "=====>result:", publicKeyHEX[4:]) // 0x049a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address) // 0x96216849c49358B10257cb55b28eA603c874b05E

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:])) // 0x96216849c49358b10257cb55b28ea603c874b05e
}
