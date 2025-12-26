package main

// var StoreABI = `[{"inputs":[{"internalType":"string","name":"_version","type":"string"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"bytes32","name":"key","type":"bytes32"},{"indexed":false,"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"ItemSet","type":"event"},{"inputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"name":"items","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"key","type":"bytes32"},{"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"setItem","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"version","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`

// func main13_1() {
// 	client, err := ethclient.Dial("https://sepolia.infura.io/v3/712a252d3b78462587d679771e4e1dc3")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// 查询最新的区块高度
// 	block, err := client.BlockByNumber(context.Background(), nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(block.Number().Uint64()) // 5671744

// 	contractAddress := common.HexToAddress("0x6725defAEb0f73Fa5C46F55B45cc3Bc6D3E350aA")
// 	query := ethereum.FilterQuery{
// 		FromBlock: big.NewInt(6920583),
// 		// ToBlock:   big.NewInt(2394201),
// 		Addresses: []common.Address{
// 			contractAddress,
// 		},
// 		// Topics: [][]common.Hash{
// 		//  {},
// 		//  {},
// 		// },
// 	}

// 	logs, err := client.FilterLogs(context.Background(), query)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	contractAbi, err := abi.JSON(strings.NewReader(StoreABI))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for _, vLog := range logs {
// 		fmt.Println(vLog.BlockHash.Hex())
// 		fmt.Println(vLog.BlockNumber)
// 		fmt.Println(vLog.TxHash.Hex())
// 		event := struct {
// 			Key   [32]byte
// 			Value [32]byte
// 		}{}
// 		// 以解码原始的日志数据
// 		err := contractAbi.UnpackIntoInterface(&event, "ItemSet", vLog.Data)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		fmt.Println(common.Bytes2Hex(event.Key[:]))
// 		fmt.Println(common.Bytes2Hex(event.Value[:]))
// 		var topics []string
// 		for i := range vLog.Topics {
// 			topics = append(topics, vLog.Topics[i].Hex())
// 		}

// 		fmt.Println("topics[0]=", topics[0])
// 		if len(topics) > 1 {
// 			fmt.Println("indexed topics:", topics[1:])
// 		}
// 	}

// 	eventSignature := []byte("ItemSet(bytes32,bytes32)")
// 	hash := crypto.Keccak256Hash(eventSignature)
// 	fmt.Println("signature topics=", hash.Hex())
// }
