package main

import (
	"fmt"
	"log"
	"math"
	"math/big"

	token "ethclient-test/erc20" // for demo

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main8() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/712a252d3b78462587d679771e4e1dc3")
	if err != nil {
		log.Fatal(err)
	}
	// Golem (GNT) Address
	// 导入 代币
	tokenAddress := common.HexToAddress("0xa0c513a3c44ADA27C476178ee1AeAc1B1831c877")
	instance, err := token.NewErc20(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// 导入 账户
	address := common.HexToAddress("0x25836239F7b632635F815689389C537133248edb")
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}
	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("name: %s\n", name)         // "name: Golem Network"
	fmt.Printf("symbol: %s\n", symbol)     // "symbol: GNT"
	fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"
	fmt.Printf("wei: %s\n", bal)           // "wei: 74605500647408739782407023"
	fbal := new(big.Float)
	fbal.SetString(bal.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))
	fmt.Printf("balance: %f", value) // "balance: 74605500.647409"

	// solcjs --abi IERC20Metadata.sol --- “用 solcjs 把 IERC20Metadata.sol 编译，只输出合约的 ABI 文件。”
	//

}
