package eventstore

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 获取合约事件
const (
	abiStr = `[{"inputs":[{"internalType":"string","name":"_version","type":"string"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"bytes32","name":"key","type":"bytes32"},{"indexed":false,"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"ItemSet","type":"event"},{"inputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"name":"items","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"key","type":"bytes32"},{"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"setItem","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"version","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`
)

func Eventstore() {
	//获取客户端
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/rGBqI4PDZG8o5pjuB0U98")
	if err != nil {
		panic(err)
	}

	//获取合约地址
	//将16进制字符串转换成以太坊地址类型
	contractAddress := common.HexToAddress("0xC4eE04fCb3D307e34F42d8167182E95cc3e3A34d")

	//构造过滤合约事件的查询条件(过滤出这个合约地址下的操作事件日志)
	//需要指定区块查询范围否则范围太大
	ethQuery := ethereum.FilterQuery{
		//从这个编号的区块开始查询
		FromBlock: big.NewInt(9838738),
		//到这个区块编号结束
		ToBlock:   big.NewInt(9838745),
		Addresses: []common.Address{contractAddress},
	}

	fmt.Println("开始获取日志")
	//根据过滤条件，查询合约事件
	logs, err := client.FilterLogs(context.Background(), ethQuery)
	if err != nil {
		log.Fatal(err)
	}

	if len(logs) == 0 {
		fmt.Println("获取日志信息 为空 ", logs)
	}

	//读取合约生成的abi文件
	contractAbi, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		log.Fatal(err)
	}

	//使用abi文件解析日志
	for _, VLog := range logs {
		fmt.Println("区块的hash值", VLog.BlockHash.Hex())
		fmt.Println("区块的编号", VLog.BlockNumber)
		fmt.Println("交易的hash值", VLog.TxHash.Hex())
		//匿名结构体
		//使用小写字符无法进行导出操作，字段必须导出：首字母大写，如 Key 而不是 key
		event := struct {
			Key   [32]byte
			Value [32]byte
		}{}
		//abi根据传入的事件名称，和数据VLog.Data,将数据封装成event
		err := contractAbi.UnpackIntoInterface(&event, "ItemSet", VLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("事件的key", common.Bytes2Hex(event.Key[:]))
		fmt.Println("事件的value", common.Bytes2Hex(event.Value[:]))

		var topics []string
		for i := range VLog.Topics {
			topics = append(topics, VLog.Topics[i].Hex())
		}

		//topic[0]: 事件签名哈希
		fmt.Println("代表事件签名哈希 topics[0]:", topics[0])
		if len(topics) > 1 {
			fmt.Print("indexed topics:", topics[1:])
		}
	}
	/*
		//每个以太坊事件日志包含
			{
				"address": "0xcontractAddress",
				"topics": [
					"0x91e4c654...",  // topic0: 事件签名哈希
					"0xkeyhash...",   // topic1: 第一个indexed参数
					"0xvaluehash..."  // topic2: 第二个indexed参数
				],
				"data": "0x..."       // 非indexed参数的数据
			}
			//标明是indexed的参数，会进入topic1，topic2
			event ItemSet(
				bytes32 indexed key,    // 进入 topics[1]
				bytes32 indexed value,  // 进入 topics[2]
				uint256 timestamp,      // 进入 data 部分
				address sender          // 进入 data 部分
			);
	*/
	//事件签名哈希
	eventSignature := []byte("ItemSet(bytes32,bytes32)")
	hash := crypto.Keccak256Hash(eventSignature)
	fmt.Println("事件签名哈希 signature topics:", hash.Hex())
}

// 订阅事件(需要 websocket RPC URL)
func SubscribeEvent() {
	//获取客户端
	client, err := ethclient.Dial("wss://sepolia.infura.io/ws/v3/f986659957214e21adfd5cc6b9a86e63")
	if err != nil {
		log.Fatal(err)
	}
	//设置过滤条件，只查询这个合约地址的事件
	contractAddress := common.HexToAddress("0xC4eE04fCb3D307e34F42d8167182E95cc3e3A34d")
	query := ethereum.FilterQuery{
		/**
		[]common.Address：类型声明，表示这是一个 common.Address 类型的切片
		{contractAddress}：初始化内容，包含一个元素 contractAddress
		例如：
			Addresses: []common.Address{contract1, contract2},  // 监听两个合约
		*/
		Addresses: []common.Address{contractAddress},
	}

	/**
	make：Go 的内置函数，用于创建 slice、map 和 channel
	chan：关键字，表示创建的是通道
	types.Log：通道传递的数据类型（以太坊事件日志类型）
	*/
	//创建日志渠道
	logs := make(chan types.Log)

	//SubscribeFilterLogs订阅日志信息
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	//构建abi对象
	contractAbi, err := abi.JSON(strings.NewReader(string(abiStr)))
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog.BlockHash.Hex())
			fmt.Println(vLog.BlockNumber)
			fmt.Println(vLog.TxHash.Hex())
			event := struct {
				Key   [32]byte
				Value [32]byte
			}{}
			err := contractAbi.UnpackIntoInterface(&event, "ItemSet", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(common.Bytes2Hex(event.Key[:]))
			fmt.Println(common.Bytes2Hex(event.Value[:]))

			var topics []string
			for i := range vLog.Topics {
				topics = append(topics, vLog.Topics[i].Hex())
			}
			fmt.Println("topic[0] 代表方法签名:", topics[0])
			if len(topics) > 1 {
				fmt.Println("index topic:", topics[1:])
			}
		}
	}

}
