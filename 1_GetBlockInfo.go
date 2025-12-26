package main

import (
	"context"
	"fmt"
	"log"

	// "math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main1() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/712a252d3b78462587d679771e4e1dc3")
	if err != nil {
		log.Fatal(err)
	}

	// blockNumber := big.NewInt(5671744)

	// 若传入 nil，它将返回最新的区块头
	header, err := client.HeaderByNumber(context.Background(), nil)
	fmt.Println(header.Number.Uint64())     // 5671744
	fmt.Println(header.Time)                // 1712798400
	fmt.Println(header.Difficulty.Uint64()) // 0
	fmt.Println(header.Hash())              // 0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5

	fmt.Println(header.Hash().Hex()) // 0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5

	if err != nil {
		log.Fatal(err)
	}
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Number().Uint64())     // 5671744
	fmt.Println(block.Time())                // 1712798400
	fmt.Println(block.Difficulty().Uint64()) // 0
	fmt.Println(block.Hash().Hex())          // 0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5
	fmt.Println(len(block.Transactions()))   // 70
	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count) // 70

	// 9885267
	// 1766307588
	// 0
	// 0x5412b45e59e4766517f51a631568eeb2e73b8790394973583a0d10d7cb2310ae
	// 9885267
	// 1766307588
	// 0
	// 0x5412b45e59e4766517f51a631568eeb2e73b8790394973583a0d10d7cb2310ae
	// 82
	// 82

}
