package tryClient

//创建新钱包
import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func CreateWallet() {
	//使用crypto包下的方法GenerateKey生成随机私钥的
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	//如果有了私钥的Hex字符串，可以使用HexToECDSA方法恢复私钥
	//privateKey,err = crypto.HexToECDSA("ccec5314acec3d18eae81b6bd988b844fc4f7f7d3c828b351de6d0fede02d3f2")

	//把 ECDSA 私钥对象转换成原始字节数组，结果：[]byte 类型的私钥原始数据
	privateKeyBytes := crypto.FromECDSA(privateKey)
	// 1. hexutil.Encode() 把字节数组转换成带"0x"前缀的16进制字符串
	// 2. [2:] 去掉字符串开头的"0x"
	//这就是用于签署交易的私钥，将被视为密码，永远不应该被共享给别人，因为谁拥有它可以访问你的所有资产
	fmt.Println("私钥:", hexutil.Encode(privateKeyBytes)[2:])
	//由于公钥是从私钥派生的，因此 go-ethereum 的加密私钥具有一个返回公钥的 Public 方法。
	publicKey := privateKey.Public()
	// 1. publicKey 是一个接口类型（可能是多种公钥格式）
	// 2. .(*ecdsa.PublicKey) 尝试转换成 ECDSA 公钥类型
	// 3. ok 返回 true 表示转换成功，false 表示失败
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	//把 ECDSA 公钥对象转换成原始字节，然后转换成16进制字符串并去掉前4个字符。
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	/**
	0x 是16进制前缀
	04 是 ECDSA 公钥的标准标识符
	去掉这4个字符后，得到的就是公钥的X和Y坐标数据
	原始公钥字节: [04] + [X坐标] + [Y坐标]
	十六进制: "0x04" + "X坐标(64字符)" + "Y坐标(64字符)"
	*/
	fmt.Println("from pubKey:", hexutil.Encode(publicKeyBytes)[4:])
	//获取公钥地址(十六进制)
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("公钥地址：", address)

	//以下生成以太坊钱包地址的核心算法，和上面的方式最终生成的公钥地址一样。

	//公共地址其实就是公钥的 Keccak-256 哈希
	//创建 Keccak256 哈希计算器（以太坊使用的哈希算法）
	hash := sha3.NewLegacyKeccak256()
	//将公钥字节（去掉第一个标识字节）写入哈希计算器
	hash.Write(publicKeyBytes[1:])
	/**
	1.hash.Sum(nil) 计算最终的哈希值
	2.[12:] 取哈希值的最后20个字节（因为以太坊地址是20字节）,Keccak256 哈希结果是32字节
	3.hexutil.Encode() 转换成16进制字符串
	*/
	fmt.Println("full:", hexutil.Encode(hash.Sum(nil)[:]))
	fmt.Println("分割后的地址:", hexutil.Encode(hash.Sum(nil)[12:]))
}
