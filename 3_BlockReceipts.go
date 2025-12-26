package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func main3() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/712a252d3b78462587d679771e4e1dc3")
	if err != nil {
		log.Fatal(err)
	}

	// 方法1：使用区块的高度或哈希以外
	// 区块的高度
	blockNumber := big.NewInt(9885267)
	blockHash := common.HexToHash("0x5412b45e59e4766517f51a631568eeb2e73b8790394973583a0d10d7cb2310ae")

	receiptByHash, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithHash(blockHash, false))
	if err != nil {
		log.Fatal(err)
	}

	receiptsByNum, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Int64())))
	if err != nil {
		log.Fatal(err)
	}
	// BlockReceipts 返回的 []*types.Receipt 切片里的每一个 *types.Receipt 都是新分配的指针
	fmt.Println(receiptByHash[0] == receiptsByNum[0]) // true
	// fmt.Println("-----", receiptByHash[0])
	// fmt.Println("-----", receiptsByNum[0])

	for _, receipt := range receiptByHash {
		fmt.Println(receipt.Status)                // 1
		fmt.Println(receipt.Logs)                  // []
		fmt.Println(receipt.TxHash.Hex())          // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
		fmt.Println(receipt.TransactionIndex)      // 0
		fmt.Println(receipt.ContractAddress.Hex()) // 0x0000000000000000000000000000000000000000
		break
	}

	// 方法2：使用交易哈希查询
	txHash := common.HexToHash("0x7fee370d0e449a33d90c0a60cafcd1b77c30f610e74a161930c6bf6c2777d524")
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(receipt.Status)                // 1
	fmt.Println(receipt.Logs)                  // []
	fmt.Println(receipt.TxHash.Hex())          // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
	fmt.Println(receipt.TransactionIndex)      // 0
	fmt.Println(receipt.ContractAddress.Hex()) // 0x0000000000000000000000000000000000000000

	// false
	// 1
	// []
	// 0x7fee370d0e449a33d90c0a60cafcd1b77c30f610e74a161930c6bf6c2777d524
	// 0
	// 0x0000000000000000000000000000000000000000
	// 1
	// []
	// 0x7fee370d0e449a33d90c0a60cafcd1b77c30f610e74a161930c6bf6c2777d524
	// 0
	// 0x0000000000000000000000000000000000000000
}
