package store

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// DeployStoreTestByABIgen 使用 abigen 生成的合约绑定代码部署 Store 合约的示例
func DeployStoreTestByABIgen() {
	//获取ETH客户端
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/rGBqI4PDZG8o5pjuB0U98")
	if err != nil {
		log.Fatal(err)
	}
	//加载部署合约的私钥
	/**
	// 转换后的 ECDSA 私钥可以：
	type ecdsa.PrivateKey struct {
		PublicKey ecdsa.PublicKey  // 包含对应的公钥
		D         *big.Int         // 私钥本身（大整数）
	}

	// 可以进行的操作：
	// 1. 生成对应的以太坊地址
	// 2. 对交易进行签名
	// 3. 对消息进行签名
	// 4. 加密/解密操作

	crypto.HexToECDSA 的作用：

		格式转换：将人类可读的十六进制字符串转换为计算机可操作的 ECDSA 对象
		启用密码学操作：只有 ECDSA 格式才能进行签名、验证等操作
		提取公钥和地址：从私钥推导出公钥和以太坊地址
		标准化处理：统一私钥的处理方式

		私钥的十六进制字符串只用于存储和传输
		ECDSA 对象用于实际的密码学操作
		转换后应立即验证有效性
		使用后应及时清除内存中的私钥
	*/
	privateKey, err := crypto.HexToECDSA("085d06eb0e8f9733067a4734a7fec9a9c88ccce1bc4bf375b5c208fc523cd70c")
	if err != nil {
		log.Fatal(err)
	}
	// 安全地清除私钥
	//defer crypto.ZeroKey(privateKey)
	//从私钥中获取公钥
	publicKey := privateKey.Public()
	//使用类型断言转换为 *ecdsa.PublicKey 类型
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	/**
	从 ECDSA 公钥生成以太坊地址
	crypto.PubkeyToAddress(*publicKeyECDSA) 的作用：
		生成唯一标识：为每个公钥生成一个唯一的20字节地址
		提供用户友好的格式：十六进制字符串，以 0x 开头
		确保安全性：单向函数，无法从地址反推公钥
	私钥 → ECDSA公钥 → 去掉04前缀 → Keccak-256哈希 → 取后20字节 → 以太坊地址
	记住：在以太坊中，地址是用户、合约的标识符，而 crypto.PubkeyToAddress 是生成这个标识符的标准方式。这是以太坊身份系统的基石。
	*/
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	//获取账户的nonce值
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	//建议的Gas价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//获取网络di
	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	/**
	bind 包的主要功能：
	1. 合约绑定：将 Solidity ABI 转换为 Go 结构体
	2. 交易构建：创建和发送交易到智能合约
	3. 调用封装：提供类型安全的合约方法调用
	4. 事件监听：订阅和处理合约事件

	创建交易发送器
	*/
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	//部署合约
	input := "1.0"
	//调用生成的go文件中(store_abigen)的DeployTryClient方法
	address, tx, instance, err := DeployTryClient(auth, client, input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("使用abigen生成的文件部署后的地址", address.Hex())
	fmt.Println("交易哈希", tx.Hash().Hex())
	//_ = instance,避免编译错误
	_ = instance

}
