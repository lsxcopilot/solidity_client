package tryClient

//获取区块查询信息
import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

func BlockQuary() {
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/rGBqI4PDZG8o5pjuB0U98")
	if err != nil {
		fmt.Println(err)
	}

	//指定区块号
	blockNumber := big.NewInt(5671744)
	fmt.Println("-------------区块头部-----------------")

	//通过区块号查询区块头,返回区块头信息,blockNumber如果是nil则返回的是最新区块
	header, err := client.HeaderByNumber(context.Background(), blockNumber)
	if err != nil {
		fmt.Println(err)
		return
	}
	//区块头号
	fmt.Println(header.Number.Uint64()) // 5671744
	//区块的产生时间
	fmt.Println(header.Time) // 1712798400
	//当前区块难度
	fmt.Println(header.Difficulty.Uint64()) // 0
	//获取头部哈希
	fmt.Println(header.Hash().Hex()) // 0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5

	fmt.Println("-------------区块部份-----------------")
	//根据区块号获取区块信息
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Number().Uint64())     // 5671744
	fmt.Println(block.Time())                // 1712798400
	fmt.Println(block.Difficulty().Uint64()) // 0
	fmt.Println(block.Hash().Hex())          // 0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5
	//获取区块的交易的数量
	fmt.Println(len(block.Transactions())) // 70

	fmt.Println("-------------通过客户端获取区块信息-----------------")
	// 使用 context.Background() 创建一个基础上下文
	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count) // 70
}
