package homework

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main1() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/712a252d3b78462587d679771e4e1dc3")
	if err != nil {
		log.Fatal(err)
	}

	// 若传入 nil，它将返回最新的区块头
	// header, err := client.HeaderByNumber(context.Background(), nil)
	// fmt.Println(header.Number.Uint64())     // 9910918
	// fmt.Println(header.Time)                // 1766654688
	// fmt.Println(header.Difficulty.Uint64()) // 0
	// fmt.Println(header.Hash())              // 0xeece3d1b19bbe9a6155230532117e4bed8f8345942e373428a8d562a49ba2bde

	// 查询区块
	blockNumber := big.NewInt(9910918)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Hash())
	fmt.Println(block.Time())
	fmt.Println(len(block.Transactions()))

	// 发送交易
	privateKey, err := crypto.HexToECDSA("958426ca0b0d6a74b91fd90bc1c2db95ed9dd31f5f59284b1ed7f3e70a42503f")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	// .(*ecdsa.PublicKey) 是类型断言语法，意思是“我断言这个接口里实际存放的是 *ecdsa.PublicKey 类型的值”;
	// 如果断言成功，它就是转换后的 *ecdsa.PublicKey 指针；失败则为 nil。
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//  nonce 就是地址的“交易计数器/创建计数器”，用来让交易有序、防重放、让合约地址可确定
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(1000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000)             // in units
	// 只是对本地连接的那一台以太坊节点发一个 JSON-RPC 调用：eth_gasPrice。
	// 节点在收到这个请求后，会按自己内存交易池（mempool）里当前 pending 交易的报价分布快速算出一个“建议值”
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())

}
