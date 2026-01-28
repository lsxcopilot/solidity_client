package mynft

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	//100,86400,5,0xd3d3cfa5755623486f2Fc904122DE9685Cf29db2
	//bafkreicwlwk6jfv22fhpzullinygbwm7bivuc3fuiufa4425jdraywhogm
	//拍卖会合约地址
	contractAddr = "0xA915DCa50f66cA420951D2f1c834D1b536Ca0A20"
	nftAddress   = "0xd3d3cfa5755623486f2Fc904122DE9685Cf29db2"
	account2     = "0xfE1d3b475B9C0072681deC2De6b36EA83A5edfc6"
	nftUrl       = "https://ipfs.io/ipfs/bafkreicwlwk6jfv22fhpzullinygbwm7bivuc3fuiufa4425jdraywhogm"
	privateKey   = "085d06eb0e8f9733067a4734a7fec9a9c88ccce1bc4bf375b5c208fc523cd70c"
)

func ExcuseStoreNFT() {
	//获取客户端
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/rGBqI4PDZG8o5pjuB0U98")

	if err != nil {
		log.Fatal("11111", err)
	}
	//加载拍卖会合约(调用生成的go合约代码中的New方法),获取到加载后的合约对象
	store, err := NewMynft(common.HexToAddress(contractAddr), client)
	if err != nil {
		log.Fatal("22222", err)
	}

	//加载nft合约(调用生成的go合约代码中的New方法),获取到加载后的合约对象
	nft, err := NewMyNFT721(common.HexToAddress(nftAddress), client)
	if err != nil {
		log.Fatal("22222", err)
	}

	//执行合约方法
	//需要先获取私钥
	privateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal("33333", err)
	}
	//10,86400,5,0xd3d3cfa5755623486f2Fc904122DE9685Cf29db2
	//获取交易发送器
	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatal("44444", err)
	}

	//创建nft
	commonAddress2 := common.HexToAddress(account2)
	mint, minErr := nft.MintNFT(opt, commonAddress2, nftUrl)
	if minErr != nil {
		log.Fatal("mintErr", minErr)
	}
	log.Printf("创建nft编号 tx sent: %s", mint.Hash().Hex())

	//等待交易被挖矿确认
	receipt, mintErr := bind.WaitMined(context.Background(), client, mint)
	if mintErr != nil {
		log.Fatal(mintErr)
	}
	fmt.Printf("创建nft交易确认在区块: %d\n", receipt.BlockNumber)

	//获取下一个tokenId
	callOpt := &bind.CallOpts{Context: context.Background()}
	nextTokenId, err := nft.NextTokenId(callOpt)
	if err != nil {
		log.Fatal("nextTokenId", err)
	}
	log.Printf("下一个tokenId: %s", nextTokenId.String())

	//进行授权操作
	txApprove, approveErr := nft.Approve(opt, common.HexToAddress(contractAddr), nextTokenId.Sub(nextTokenId, big.NewInt(1)))
	if approveErr != nil {
		log.Fatal("approveErr", approveErr)
	}

	log.Printf("授权操作编号tx sent: %s", txApprove.Hash().Hex())
	receiptApprove, approveErr := bind.WaitMined(context.Background(), client, txApprove)
	if approveErr != nil {
		log.Fatal(approveErr)
	}
	fmt.Printf("授权交易确认在区块: %d\n", receiptApprove.BlockNumber)

	//授权成功之后,创建拍卖
	//输出方法执行结果,添加key-value
	/*
		uint256 startPrice,
		uint256 duringTime,
		uint256 tokenId,
		address nftContractAddress
	*/
	value := big.NewInt(10)                               //起拍价格0.001eth
	duringTime := big.NewInt(86400)                       //持续时间1天
	key := nextTokenId.Sub(nextTokenId, big.NewInt(1))    //tokenId
	nftContractAddress := common.HexToAddress(nftAddress) //NFT合约地址

	//调用者是我自己的私钥用户(创建拍卖之前需要先拥有对应的NFT,就是授权给拍卖合约地址)
	tx, err := store.CreateAuction(opt, value, duringTime, key, nftContractAddress)
	if err != nil {
		log.Fatal("55555: ", err.Error())
	}

	log.Printf("创建拍卖的编号tx sent: %s", tx.Hash().Hex())

	receiptCreate, createAuctionErr := bind.WaitMined(context.Background(), client, tx)
	if createAuctionErr != nil {
		log.Fatal(createAuctionErr)
	}
	fmt.Printf("创建拍卖交易确认在区块: %d\n", receiptCreate.BlockNumber)

	//获取数据,根据生成的go代码的合约，生成调用函数所需要的对象
	valueContract, err := store.AuctionMapping(callOpt, key)
	if err != nil {
		log.Fatal("66666", err)
	}
	log.Printf("value from contract: %s", string(valueContract.TokenId.Bytes()))

}

// one : https://ipfs.io/ipfs/bafkreidjtxhcmm3mnynv76qhu5fqdqnohroylliwzsr43xhs67o6xhzg6a
// two : https://ipfs.io/ipfs/bafkreicdrcv2wuzeyyx2gxn2widjmzp3b7x4kgm64k6346bm6mfnc3axdi
// three :https://ipfs.io/ipfs/bafkreifns3yxp3cxywpx2frvu2274drk6knt427bikdaf4uobyp26j54lu
func MintNFT(NFTUrl string) {
	//获取客户端
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/rGBqI4PDZG8o5pjuB0U98")

	if err != nil {
		log.Fatal("11111", err)
	}

	//加载nft合约(调用生成的go合约代码中的New方法),获取到加载后的合约对象
	nft, err := NewMyNFT721(common.HexToAddress(nftAddress), client)
	if err != nil {
		log.Fatal("22222", err)
	}

	//执行合约方法
	//需要先获取私钥
	privateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal("33333", err)
	}
	//10,86400,5,0xd3d3cfa5755623486f2Fc904122DE9685Cf29db2
	//获取交易发送器
	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatal("44444", err)
	}

	//创建nft
	commonAddress2 := common.HexToAddress(account2)
	mint, minErr := nft.MintNFT(opt, commonAddress2, NFTUrl)
	if minErr != nil {
		log.Fatal("mintErr", minErr)
	}
	log.Printf("创建nft编号 tx sent: %s", mint.Hash().Hex())

	//等待交易被挖矿确认
	receipt, mintErr := bind.WaitMined(context.Background(), client, mint)
	if mintErr != nil {
		log.Fatal(mintErr)
	}
	fmt.Printf("创建nft交易确认在区块: %d\n", receipt.BlockNumber)

	//获取下一个tokenId
	callOpt := &bind.CallOpts{Context: context.Background()}
	nextTokenId, err := nft.NextTokenId(callOpt)
	if err != nil {
		log.Fatal("nextTokenId", err)
	}
	log.Printf("下一个tokenId: %s", nextTokenId.String())
}

func Apprave(tokenId int64) {
	//获取客户端
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/rGBqI4PDZG8o5pjuB0U98")

	if err != nil {
		log.Fatal("11111", err)
	}

	//加载nft合约(调用生成的go合约代码中的New方法),获取到加载后的合约对象
	nft, err := NewMyNFT721(common.HexToAddress(nftAddress), client)
	if err != nil {
		log.Fatal("22222", err)
	}

	//执行合约方法
	//需要先获取私钥
	privateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal("33333", err)
	}

	//获取交易发送器
	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatal("44444", err)
	}
	//先判断是否已经有授权有则直接返回，没有则进行授权
	callOpt := &bind.CallOpts{Context: context.Background()}
	approvedAddress, err := nft.GetApproved(callOpt, big.NewInt(tokenId))
	log.Printf("获取到的授权地址:", approvedAddress.Hex())
	//获取零地址
	var zeroAddress = common.Address{}
	//零地址则没有进行授权
	if approvedAddress == zeroAddress {
		//进行授权操作
		txApprove, approveErr := nft.Approve(opt, common.HexToAddress(contractAddr), big.NewInt(tokenId))
		if approveErr != nil {
			log.Fatal("approveErr", approveErr)
		}

		log.Printf("授权操作编号tx sent: %s", txApprove.Hash().Hex())
		receiptApprove, approveErr := bind.WaitMined(context.Background(), client, txApprove)
		if approveErr != nil {
			log.Fatal(approveErr)
		}
		fmt.Printf("授权交易确认在区块: %d\n", receiptApprove.BlockNumber)
	} else {
		log.Fatal("已进行过授权,无需重复授权")
	}

}
