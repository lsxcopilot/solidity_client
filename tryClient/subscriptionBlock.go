package tryClient

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 订阅新区块,实时获取区块信息。订阅区块需要使用WebSocket连接以太坊节点。
/**
区块链节点处理区块头和区块体是异步的，这种设计是故意的，目的是快速广播区块头让网络知道有新区块，然后慢慢验证交易。因此你的代码也需要适应这种异步性
*/
func SubscriptionBlock() {
	//获取客户端
	client, err := ethclient.Dial("wss://sepolia.infura.io/ws/v3/f986659957214e21adfd5cc6b9a86e63")
	if err != nil {
		log.Fatal(err)
	}

	//创建一个区块头的通道
	headers := make(chan *types.Header)

	//订阅新区块头
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}
	//持续监听新区块头
	for {
		//select 监听通道和错误通道
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			//计算头哈希 转成16进制字符串输出
			fmt.Println("十六进制头地址", header.Hash().Hex())
			//根据区块头哈希获取完整区块信息
			/**
			BlockByHash 实现：直接查询本地数据库
			如果节点还没存储完整区块，返回 "not found"

			BlockByNumber 实现：
			1. 先查本地数据库
			2. 如果没有，可以向其他节点请求
			3. 有更多的容错机制

			BlockByNumber 的工作方式：
			1. 节点内部可能已经缓存了该高度的区块
			2. 即使没有，节点可以从对等节点请求
			3. 区块号是确定的，哈希可能因重组变化

			BlockByHash 的工作方式：
			1. 必须精确匹配哈希
			2. 如果节点还没收到完整区块，直接失败
			3. 区块重组时哈希会变化
			*/
			//block, err := client.BlockByHash(context.Background(), header.Hash())
			block, err := client.BlockByNumber(context.Background(), header.Number)
			if err != nil {
				fmt.Println("根据区块头哈希获取完整区块信息失败")
				log.Fatal(err)
			}
			//输出区块哈希并换成16进制字符串
			fmt.Println(block.Hash().Hex())
			//用64位无符号整数输出区块高度
			fmt.Println(block.Number().Uint64())
			fmt.Println(block.Time())
			fmt.Println(block.Nonce())
			//输出区块内交易数量
			fmt.Println(len(block.Transactions()))
		}
	}
}
