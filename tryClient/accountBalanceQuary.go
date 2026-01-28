package tryClient

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// accountBalanceQuary.go 账户余额查询
/**
整体流程：
1.连接以太坊客户端
2.定义要查询余额的账户地址
3.查询该账户在最新区块的余额
4.查询该账户在指定区块高度的余额
5.将余额从最小单位wei转换为以太（ETH）单位进行显示
6.查询该账户的待处理交易余额
*/
/**
	总结
简单来说：
	PendingBalanceAt = "我现在应该有多少钱？"（包含未确认的交易）
	BalanceAt(nil) = "我账上实际有多少钱？"（只包含已确认的交易）
选择建议：
	用户界面显示：优先使用 PendingBalanceAt 提供更好的用户体验
	财务结算、合约交互：使用 BalanceAt(nil) 确保状态确定
	风险评估：两者结合使用，既要考虑实时状态，也要尊重链上最终确认
*/
func AccountBalanceQuary() {
	//获取客户端
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/rGBqI4PDZG8o5pjuB0U98")
	if err != nil {
		log.Fatal(err)
	}
	//common.HexToAddress() 是将十六进制字符串转换为以太坊地址类型的函数。
	//0x25836239F7b632635F815689389C537133248edb
	account := common.HexToAddress("0xfe1d3b475b9c0072681dec2de6b36ea83a5edfc6")
	//查询账户余额,nil表示查询最新区块的余额
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("初始：account balance: %s", balance.String())

	blockNumber := big.NewInt(5532993)
	balance2, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("account balance2 at block %d: %s", blockNumber, balance2.String())
	//设置一个大数来转换余额单位
	fbalance := new(big.Float)
	//将余额转换成以太为单位
	fbalance.SetString(balance2.String())
	//1e18是以太坊的最小单位wei和以太之间的换算关系
	//Quo 表示除法，将余额除以1e18得到以太为单位的余额
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(1e18))
	log.Printf("转换ETH：account balance at block %d: %f ETH", blockNumber, ethValue)
	//查询待处理交易的账户余额
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("pending account balance: %s", pendingBalance.String())
	//转换成以太为单位
	fpendingBalance := new(big.Float)
	fpendingBalance.SetString(pendingBalance.String())
	pendingEthValue := new(big.Float).Quo(fpendingBalance, big.NewFloat(1e18))
	log.Printf("转换ETH：pending account balance: %f ETH", pendingEthValue)
}
