package loadingstore

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	//已经部署的合约地址
	contractAddr = "0x239A5baFC5b769B21Ea229024aEAD4fEF96ADdd5"
)

func LoadingContract() {
	//获取客户端
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/rGBqI4PDZG8o5pjuB0U98")
	if err != nil {
		log.Fatal("连接客户端失败:", err)
	}
	//加载合约
	storeContract, err := NewLoadingstore(common.HexToAddress(contractAddr), client)
	if err != nil {
		log.Fatal("加载合约失败:", err)
	}
	_ = storeContract
}
