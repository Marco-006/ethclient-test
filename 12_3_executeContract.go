package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	contractAddr = "0x5D54A3962136297a7c8D800A20487d59e54D487a"
)

func main12_3() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/712a252d3b78462587d679771e4e1dc3")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("958426ca0b0d6a74b91fd90bc1c2db95ed9dd31f5f59284b1ed7f3e70a42503f")
	if err != nil {
		log.Fatal(err)
	}

	// 获取公钥地址
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 获取 nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// 估算 gas 价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 准备交易数据
	methodSignature := []byte("setItem(bytes32,bytes32)")
	methodSelector := crypto.Keccak256(methodSignature)[:4]

	var key [32]byte
	var value [32]byte
	copy(key[:], []byte("demo_save_key_no_use_abi"))
	copy(value[:], []byte("demo_save_value_no_use_abi_11111"))

	// 组合调用数据
	var input []byte
	input = append(input, methodSelector...)
	input = append(input, key[:]...)
	input = append(input, value[:]...)

	// 创建交易并签名交易
	chainID := big.NewInt(int64(11155111))
	tx := types.NewTransaction(nonce, common.HexToAddress(contractAddr), big.NewInt(0), 300000, gasPrice, input)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 降低GAS消耗的一种方式
	// GasTipCap： 给矿工/验证者的小费上限（单位 Wei），即旧语境里的 “maxPriorityFeePerGas”。真正的小费 = min(gasTipCap, gasFeeCap −
	// baseFee)。

	// gasFeeCap : 用户愿意为单位 Gas 支付的最高总价（单位 Wei），即旧语境里的 “maxFeePerGas”。必须 ≥ 基础费（baseFee）且 ≥ gasTipCap，
	// 否则会被直接丢弃

	// types.NewTx(&types.DynamicFeeTx{
	// 	ChainID : chainID,
	// 	Nonce :   nonce,
	// 	GasTipCap  *big.Int // a.k.a. maxPriorityFeePerGas
	// 	GasFeeCap  *big.Int // a.k.a. maxFeePerGas
	// 	Gas        uint64
	// 	To         *common.Address `rlp:"nil"` // nil means contract creation
	// 	Value      *big.Int
	// 	Data       []byte
	// })

	// 发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	// 0x9150e5b28f3c0ded2f981ccb7c2c26ea3a1bc85338cce876b8352c11099b750b
	fmt.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())
	_, err = waitForReceipt(client, signedTx.Hash())
	if err != nil {
		log.Fatal(err)
	}

	// 查询刚刚设置的值
	itemsSignature := []byte("items(bytes32)")
	itemsSelector := crypto.Keccak256(itemsSignature)[:4]

	var callInput []byte
	callInput = append(callInput, itemsSelector...)
	callInput = append(callInput, key[:]...)

	to := common.HexToAddress(contractAddr)
	callMsg := ethereum.CallMsg{
		To:   &to,
		Data: callInput,
	}

	// 解析返回值
	result, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		log.Fatal(err)
	}

	var unpacked [32]byte
	copy(unpacked[:], result)
	fmt.Println("is value saving in contract equals to origin value:", unpacked == value)
}

func waitForReceipt(client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	for {
		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		if err == nil {
			return receipt, nil
		}
		if err != ethereum.NotFound {
			return nil, err
		}
		// 等待一段时间后再次查询
		time.Sleep(1 * time.Second)
	}
}
