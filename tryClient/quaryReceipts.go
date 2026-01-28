package tryClient

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func QuaryReceipts() {
	//获取测试网客户端
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/rGBqI4PDZG8o5pjuB0U98")
	//如果获取失败打印日志
	if err != nil {
		log.Fatal(err)
	}
	//获取指定编号的区块
	/**
	创建一个值为 5671744 的大整数对象。

	普通的 int 类型有大小限制
	big.Int 可以处理任意大的整数（比如区块链中的巨大数值）
	big.NewInt(5671744) 就是把普通的数字 5671744 转换成能处理大数的格式
	*/
	blockNumber := big.NewInt(5671744)
	//就是把16进制字符串转换成区块链哈希对象(Hex : 表示16进制)
	blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5")

	//rpc.BlockNumberOrHashWithHash(blockHash, false)创建一个区块查询参数，指定用哈希来查询区块，并且不要求返回完整交易详情。可以优化性能，只获取交易哈希而不是完整交易
	//blockHash：要查询的区块哈希 false：不返回完整的交易对象（只返回交易哈希）
	receiptByHash, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithHash(blockHash, false))
	if err != nil {
		log.Fatal(err)
	}

	//rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Int64())) 创建一个区块查询参数，指定用区块号来查询区块。
	//rpc.BlockNumber()：把普通数字转换成区块链能理解的区块号格式
	receiptByNum, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Int64())))
	if err != nil {
		log.Fatal(err)
	}
	//两种获取方式拿到的是同一个收据
	fmt.Println(receiptByHash[0] == receiptByNum[0])

	//打印获取的详细信息
	for _, receipt := range receiptByHash {

		fmt.Println("收据状态：", receipt.Status)
		fmt.Println("收据日志：", receipt.Logs)
		fmt.Println("收据交易的hash的16进制：", receipt.TxHash.Hex())
		fmt.Println("收据所处的交易编号", receipt.TransactionIndex)
		fmt.Println("收据所处的合约地址", receipt.ContractAddress.Hex())
		break
	}
	//将16进制转成hash
	txHash := common.HexToHash("0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5")
	//通过交易哈希查询这笔交易的执行结果收据。
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("收据状态：", receipt.Status)
	fmt.Println("收据日志：", receipt.Logs)
	fmt.Println("收据交易的hash的16进制：", receipt.TxHash.Hex())
	fmt.Println("收据所处的交易编号", receipt.TransactionIndex)
	/**
	只在部署合约的交易中有效
	返回新创建的智能合约地址
	如果是普通转账交易，返回空地址（0x0）
	*/
	fmt.Println("合约地址：", receipt.ContractAddress.Hex())
}
