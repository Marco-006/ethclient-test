package main

// const (
// 	contractAddr = "0x5D54A3962136297a7c8D800A20487d59e54D487a"
// )

// func main() {
// 	client, err := ethclient.Dial("https://sepolia.infura.io/v3/712a252d3b78462587d679771e4e1dc3")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	storeContract, err := store.NewStore(common.HexToAddress(contractAddr), client)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	privateKey, err := crypto.HexToECDSA("958426ca0b0d6a74b91fd90bc1c2db95ed9dd31f5f59284b1ed7f3e70a42503f")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var key [32]byte
// 	var value [32]byte

// 	copy(key[:], []byte("demo_save_key"))
// 	copy(value[:], []byte("demo_save_value11111"))
//     // sepolia网络中  chinId是固定的
// 	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	tx, err := storeContract.SetItem(opt, key, value)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("tx hash:", tx.Hash().Hex())

// 	callOpt := &bind.CallOpts{Context: context.Background()}
// 	valueInContract, err := storeContract.Items(callOpt, key)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("is value saving in contract equals to origin value:", valueInContract == value)
// }
