package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const maxRetry = 5
const retryInterval = 200 * time.Millisecond

func main9_1() {
	client, err := ethclient.Dial("wss://sepolia.infura.io/ws/v3/712a252d3b78462587d679771e4e1dc3")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	log.Println("waiting for new blocks...")
	for {
		select {
		case err := <-sub.Err():
			log.Printf("subscription error: %v", err)
			time.Sleep(2 * time.Second) // 简单重连可在这里实现
		case header := <-headers:
			// 异步+重试，避免订阅循环被卡住
			go func(h *types.Header) {
				var block *types.Block
				var err error
				for i := 0; i < maxRetry; i++ {
					block, err = client.BlockByHash(context.Background(), h.Hash())
					if err == nil {
						break
					}
					if i < maxRetry-1 {
						time.Sleep(retryInterval)
					}
				}
				if err != nil {
					log.Printf("⚠️  block %s still not found after %d retries", h.Hash(), maxRetry)
					return
				}
				fmt.Println("✅ new block",
					"num=", block.Number(),
					"hash=", block.Hash().Hex(),
					"time=", time.Unix(int64(block.Time()), 0).Format("15:04:05"),
					"txs=", len(block.Transactions()))
			}(header)
		}
	}
}
