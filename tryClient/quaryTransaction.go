package tryClient

// 获取区块的交易信息
import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func QuaryTransaction() {
	//获取client对象
	// 连接到 Sepolia 测试网
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/rGBqI4PDZG8o5pjuB0U98")
	if err != nil {
		log.Fatal(err)
	}

	//获取链id，是那条链上的交易
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//获取区块信息
	blockNumber := big.NewInt(5671744)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	//循环遍历区块的交易信息
	for _, tx := range block.Transactions() {
		fmt.Println(tx.Hash().Hex())        // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
		fmt.Println(tx.Value().String())    // 100000000000000000
		fmt.Println(tx.Gas())               // 21000
		fmt.Println(tx.GasPrice().Uint64()) // 100000000000
		fmt.Println(tx.Nonce())             // 245132
		fmt.Println(tx.Data())              // []
		fmt.Println(tx.To().Hex())          // 0x8F9aFd209339088Ced7Bc0f57Fe08566ADda3587

		// 创建 EIP-155 签名器：types.NewEIP155Signer(chainID)，创建符合 EIP-155 标准的交易签名器的函数
		// EIP-155 签名器简单说：它确保你的交易签名只对特定区块链有效，防止交易被恶意重放到其他链上。

		// 相当于：
		// 1. 创建一个知道当前区块链ID的签名验证器
		// 2. 用这个验证器从交易签名中解析出"谁发送了这笔交易"
		// 3. 返回发送者的钱包地址
		//这行代码就是通过交易链ID的签名来找出"谁发了这笔交易"。
		if sender, err := types.Sender(types.NewEIP155Signer(chainID), tx); err == nil {
			fmt.Println("sender", sender.Hex()) // 0x2CdA41645F2dBffB852a605E92B185501801FC28
		} else {
			log.Fatal(err)
		}

		//获取交易的请求结果，是成功还是失败
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(receipt.Status) // 1
		fmt.Println(receipt.Logs)   // []
		break
	}

	//将区块转成哈希对象
	blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5")
	//查看这个区块有多少交易
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}

	for idx := uint(0); idx < count; idx++ {
		// 根据索引获取区块中的第 idx 笔交易
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}

		// 打印这笔交易的哈希
		fmt.Println(tx.Hash().Hex()) // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
		break
	}

	//将交易哈希字符串转换为哈希对象
	txHash := common.HexToHash("0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5")
	//通过哈希查询交易详情，并返回三个值：
	//tx: 交易对象（包含发送者、接收者、金额等信息）
	//isPending: 交易是否还在等待被打包（true=等待中，false=已打包）
	//err: 错误信息
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(isPending)       // false
	fmt.Println(tx.Hash().Hex()) // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5.Println(isPending)

}
