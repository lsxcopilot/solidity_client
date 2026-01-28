package excusestore

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

/*
*
方式一：使用abigen生成的go合约代码，进行加载执行
*/

const (
	contractAddr = "0xC4eE04fCb3D307e34F42d8167182E95cc3e3A34d"
)

func ExcuseStoreFirst() {
	//获取客户端
	clinet, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/rGBqI4PDZG8o5pjuB0U98")
	if err != nil {
		log.Fatal(err)
	}
	//加载合约(调用生成的go合约代码中的New方法),获取到加载后的合约对象
	store, err := NewLoadingstore(common.HexToAddress(contractAddr), clinet)
	if err != nil {
		log.Fatal(err)
	}
	//执行合约方法
	//需要先获取私钥
	privateKey, err := crypto.HexToECDSA("085d06eb0e8f9733067a4734a7fec9a9c88ccce1bc4bf375b5c208fc523cd70c")
	if err != nil {
		log.Fatal(err)
	}
	var key [32]byte
	var value [32]byte

	copy(key[:], []byte("demo_key"))
	copy(value[:], []byte("demo_value_111"))

	//获取交易发送器
	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatal(err)
	}
	//输出方法执行结果,添加key-value
	tx, err := store.SetItem(opt, key, value)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("tx sent: %s", tx.Hash().Hex())

	//获取数据,根据生成的go代码的合约，生成调用函数所需要的对象
	callOpt := &bind.CallOpts{Context: context.Background()}
	valueContract, err := store.Items(callOpt, key)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("value from contract: %s", string(valueContract[:]))
}

// 方式二：使用abi文件和ethclient加载合约，执行合约方法
func ExcuseStoreSecned() {

	//创建客户端
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/rGBqI4PDZG8o5pjuB0U98")
	if err != nil {
		log.Fatal(err)
	}

	//获取私钥
	//crypto.HexToECDSA()从私钥字符串恢复私钥对象
	privateKey, err := crypto.HexToECDSA("085d06eb0e8f9733067a4734a7fec9a9c88ccce1bc4bf375b5c208fc523cd70c")
	if err != nil {
		log.Fatal(err)
	}

	//获取公钥
	publicKey := privateKey.Public()
	//将公钥通过类型断言的方式转化为ecdsa公钥
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	//通过公钥获取地址
	//crypto.PubkeyToAddress(*publicKeyECDSA) 是将 ECDSA 公钥转换为以太坊地址的函数。
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//获取nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	//评估gas
	gasprice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//准备交易信息(将合约生成的abi文件进行传入)
	contractABI, err := abi.JSON(strings.NewReader(`[{"inputs":[{"internalType":"string","name":"_version","type":"string"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"bytes32","name":"key","type":"bytes32"},{"indexed":false,"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"ItemSet","type":"event"},{"inputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"name":"items","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"key","type":"bytes32"},{"internalType":"bytes32","name":"value","type":"bytes32"}],"name":"setItem","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"version","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"}]`))
	if err != nil {
		log.Fatal(err)
	}

	//准备函数名称
	methodName := "setItem"

	//准备函数参数
	var key [32]byte
	var value [32]byte

	copy(key[:], []byte("demo_key"))
	copy(value[:], []byte("demo_value_222"))

	//打包函数和参数
	input, err := contractABI.Pack(methodName, key, value)
	if err != nil {
		log.Fatal(err)
	}

	//创建交易发送器,并签名
	chainID := big.NewInt(int64(11155111))
	tx := types.NewTransaction(nonce, common.HexToAddress(contractAddr), big.NewInt(0), 100000, gasprice, input)
	//交易签名
	tx, err = types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	//发送交易
	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("tx sent: %s", tx.Hash().Hex())

	//等待交易完成
	waitForReceipt(client, tx.Hash())

	//查询刚才设置的值
	methodName = "items"
	//打包函数和数据
	callInput, err := contractABI.Pack(methodName, key)
	if err != nil {
		log.Fatal(err)
	}
	//调用合约
	to := common.HexToAddress(contractAddr)
	callMsg := ethereum.CallMsg{
		To:   &to,
		Data: callInput,
	}

	//解析返回值
	result, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		log.Fatal(err)
	}

	var unpacked [32]byte //存储解析后的值
	//将原始字节数据解析为 Go 数据结构
	contractABI.UnpackIntoInterface(&unpacked, "items", result)
	log.Printf("value from contract: %s", string(unpacked[:]))

}

func waitForReceipt(client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	for {
		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		if err == nil {
			return receipt, nil
		}
		if err != ethereum.NotFound {
			return nil, err
		}
		// 等待一段时间后再次查询
		time.Sleep(1 * time.Second)
	}
}
