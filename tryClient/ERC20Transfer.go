package tryClient

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
)

// ERC20Transfer  ERC20代币转账
/**
整体流程：
1.连接以太坊客户端
2.加载发送方私钥并获取对应的公钥和地址
3.获取发送方账户的nonce值
4.设置转账金额和gas价格
5.准备调用ERC20合约的transfer函数所需的数据字段
6.估算执行合约所需的gasLimit
7.构建交易对象
8.获取网络Chain ID
9.使用私钥对交易进行签名
10.发送签名后的交易到以太坊网络
*/
func ERC20Transfer() {
	//获取测试网客户端
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/rGBqI4PDZG8o5pjuB0U98")
	if err != nil {
		log.Fatal(err)
	}

	//获取账户私钥
	privateKey, err := crypto.HexToECDSA("4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	if err != nil {
		log.Fatal(err)
	}

	//获取公钥
	publicKey := privateKey.Public()
	//将公钥转换成ECDSA类型的公钥
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	//把公钥转换成对应的以太坊钱包地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	//从链上获取当前账户的nonce值
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	//设置转账金额
	value := big.NewInt(0)
	//获取平均消费gas
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//设置接收地址
	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	//设置合约地址
	// 如果是普通ETH转账：目标账户地址
	// 如果是合约调用：合约地址
	tokenAddress := common.HexToAddress("0x28b149020d2152179873ec60bed6bf7cd705775d")
	//创建包含函数签名字符串的字节切片
	//"transfer(address,uint256)" 是ERC20代币标准的转账函数签名
	//准备调用智能合约的transfer函数所需的标识符和参数编码。
	/**
	常见的函数签名
	// ERC20
	[]byte("transfer(address,uint256)")
	[]byte("approve(address,uint256)")
	[]byte("balanceOf(address)")

	// ERC721
	[]byte("safeTransferFrom(address,address,uint256)")
	*/
	transferFnSignature := []byte("transfer(address,uint256)")
	/**
	sha3：引用了以太坊的加密库（如golang.org/x/crypto/sha3）
	NewLegacyKeccak256()：创建传统Keccak-256哈希器
	hash：返回的哈希计算对象
	*/
	hash := sha3.NewLegacyKeccak256()
	//将函数签名写入哈希对象
	hash.Write(transferFnSignature)
	//获取前四个字节的哈希值作为方法ID,此时方法id是字节数组
	methodID := hash.Sum(nil)[:4]
	//对字节数组进行编码为16进制字符串并打印
	fmt.Println(hexutil.Encode(methodID))
	//对接收地址进行填充，使其符合以太坊ABI规范
	//common.LeftPadBytes作用是将地址左填充到32字节（256位），这是以太坊ABI编码的规范要求。
	//以太坊虚拟机（EVM）要求所有参数必须是32字节的倍数
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	//对字节进行编码为16进制字符串并打印
	fmt.Println(hexutil.Encode(paddedAddress))
	//创建一个大整数对象
	amount := new(big.Int)
	//对象填充值,10进制的1000000000000000000数值
	//kwei (10^3) , mwei(10^6) , gwei(10^9) , szabo(10^12) , finney(10^15) , ether(10^18)
	amount.SetString("1000000000000000000", 10)
	//将金额填充为32字节
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	//对字节进行编码为16进制字符串并打印
	fmt.Println(hexutil.Encode(paddedAmount))
	//将方法ID、填充后的地址和填充后的金额组合成数据字段
	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	//估算执行合约所需的gasLimit
	//EstimateGas：估算Gas消耗的方法
	//CallMsg：模拟调用的消息结构
	//目的：在不实际发送交易的情况下，预计算需要多少Gas
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &tokenAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gasLimit)
	//这是在构建一个以太坊交易对象。这个函数创建了一个完整的待发送交易
	//types.NewTransaction 是在构建一个待签名的原始交易对象，包含了执行所需的所有参数。这是发送交易前的关键步骤，构建后的交易需要签名才能广播到网络
	tx := types.NewTransaction(
		nonce,        // 序号（防止重放攻击）
		tokenAddress, // 目标地址（合约地址）
		value,        // 转账的ETH金额（单位：wei）
		gasLimit,     // Gas上限
		gasPrice,     // Gas价格（单位：wei）
		data,         // 调用合约的函数和数据
	)
	//获取当前连接的网络ID（Chain ID），这是交易签名和网络验证的关键参数。
	//如果没有Chain ID，交易可以在不同链上重放，导致安全问题。
	//获取Chain ID主要是为了EIP-155交易签名，防止跨链重放攻击，并确保交易在正确的网络上执行。这是发送交易前的必要步骤。
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//使用私钥对交易进行签名，确保交易的真实性和完整性。
	// 签名证明：你是私钥的持有者，授权这笔交易
	// 防篡改：签名后的交易内容不能被修改，否则签名无效
	// EIP-155：使用链ID防止跨链重放攻击
	//签名是用私钥证明交易合法性的必要步骤，NewEIP155Signer 创建包含Chain ID的签名器以防止跨链重放攻击，确保交易安全可靠
	signedTx, err := types.SignTx(
		tx,                             // 未签名的原始交易
		types.NewEIP155Signer(chainID), // EIP-155签名器
		privateKey,                     // 发送者的私钥
	)
	if err != nil {
		log.Fatal(err)
	}
	//发送签名后的交易到以太坊网络
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
