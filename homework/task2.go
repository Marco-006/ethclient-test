package homework

// "ethclient-test/store"

// "ethclient-test/homework"

// const (
// 	contractAddr2 = "0xB24f37224f33dC99e9069fc3C66D6BD90306C77a"
// )

// func main() {

// 	client, err := ethclient.Dial("https://sepolia.infura.io/v3/712a252d3b78462587d679771e4e1dc3")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	//  实例化合约 —— 直接用包名 + 结构体名
// 	storeContract, err := homework.NewHomework(common.HexToAddress(contractAddr2), client)
// 	// storeContract, err := homework.NewCounter(common.HexToAddress(contractAddr2), client)
// 	// storeContract, err := store.NewStore(common.HexToAddress(contractAddr2), client)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// 私钥去掉0x
// 	privateKey, err := crypto.HexToECDSA("958426ca0b0d6a74b91fd90bc1c2db95ed9dd31f5f59284b1ed7f3e70a42503f")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// 构造一个 交易授权 对象
// 	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
// 	if err != nil {
// 		log.Printf("---", err)
// 	}

// 	num_0, err := storeContract.GetNum(&bind.CallOpts{Context: context.Background()})
// 	if err != nil {
// 		log.Fatalf("...", err)
// 	}
// 	fmt.Println("current num =", num_0)

// 	// storeContract.AddNum()
// 	// auth, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
// 	tx, err := storeContract.AddNum(opt)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("tx hash:", tx.Hash().Hex())

// 	// 如果需要重新查询的话，需要重新连接

// 	// num_1, err := storeContract.GetNum(&bind.CallOpts{Context: context.Background()})
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// fmt.Println("current num =", num_1)

// 	// 执行结果
// 	// current num = 3
// 	// tx hash: 0xc107f61fbd3807e1e2384bd98f8ba872a90ec6965b53638476c05dddf7cb4db9
// 	// current num = 3
// }
