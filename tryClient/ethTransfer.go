package tryClient

//转账操作
import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// EthTransfer 以太币转账
/**
整体流程：
1.连接以太坊客户端
2.加载发送方私钥并获取对应的公钥和地址
3.获取发送方账户的nonce值
4.设置转账金额和gas价格
5.构建交易对象
6.获取网络Chain ID
7.使用私钥对交易进行签名
8.发送签名后的交易到以太坊网络
*/
// 1000 0000 0000 0000 000
func EthTransfer() {

	//获取客户端
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/rGBqI4PDZG8o5pjuB0U98")
	if err != nil {
		log.Fatal(err)
	}
	//ECDSA 是以太坊用来生成钱包、签名交易的核心密码学算法。
	//十六进制转换成ECDSA私钥对象，把保存在字符串中的私钥还原成程序能用的私钥对象。
	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}

	//提取公钥,公钥是从私钥派生的
	publicKey := privateKey.Public()
	//将公钥转换成ECDSA类型的公钥
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	//获取公钥地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	//获取应该发送的nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	//设置发送的wei
	value := big.NewInt(100000000000000000)
	//设置gas的转账上限为21000个单位
	gasLimit := uint64(21000)
	//SuggestGasPrice 函数，用于根据'x'个先前块来获得平均gas价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//将十六进制字符串转成地址
	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	var data []byte
	//进调用 NewTransaction 来生成我们的未签名以太坊事务参数包括：nonce,目标地址，转账金额，gas费用上限，gas金额，附带的数据
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	//使用发件人的私钥对事物进行签名
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//调用 SignTx 方法，该方法接受一个未签名的事务和我们之前构造的私钥,返回签名好的事务
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	//发送事务
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent:%s", signedTx.Hash().Hex())
}
