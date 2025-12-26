package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func main4() {
	// 成随机私钥的 GenerateKey
	// privateKey, err := crypto.GenerateKey()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("privateKey...", privateKey)

	// 如果已经有了私钥的 Hex 字符串，也可以使用 HexToECDSA 方法恢复私钥：
	privateKey, err := crypto.HexToECDSA("ccec5314acec3d18eae81b6bd988b844fc4f7f7d3c828b351de6d0fede02d3f2")
	if err != nil {
		log.Fatal(err)
	}

	// 将私钥转换为字节
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("私钥转换为字节...", hexutil.Encode(privateKeyBytes)[2:]) // 去掉'0x'
	// publicKey
	publicKey := privateKey.Public()
	// fmt.Println("publicKey...", publicKey)

	// 将其转换为十六进制的过程与我们使用转化私钥的过程类似。
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("from pubKey:", hexutil.Encode(publicKeyBytes)[4:]) // 去掉'0x04'

	// go-ethereum 加密包有一个 PubkeyToAddress 方法，它接受一个 ECDSA 公钥，并返回公共地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address)
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println("full:", hexutil.Encode(hash.Sum(nil)[:]))
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:])) // 原长32位，截去12位，保留后20位
}
