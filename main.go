package main

import mynft "client/tryClient/myNFT"

//"client/tryClient/store"
//"client/tryClient"

func main() {

	//tryClient.BlockQuary()
	//tryClient.QuaryTransaction()
	//tryClient.QuaryReceipts()
	//tryClient.CreateWallet()
	//tryClient.EthTransfer()  没有测试有问题
	//tryClient.ERC20Transfer() //没有测试有问题
	//tryClient.AccountBalanceQuary()
	//tryClient.ERC20BalanceQuary()
	//tryClient.SubscriptionBlock()
	//store.DeployStoreTestByABIgen()
	//store.DeployStoreTestByEthclient() //使用ethclient部署合约,有问题通过合约地址查询合约没有合约的编译文件
	//loadingstore.LoadingContract()
	//excusestore.ExcuseStoreFirst()
	//excusestore.ExcuseStoreSecned()
	//eventstore.Eventstore()
	//eventstore.SubscribeEvent()
	//mynft.ExcuseStoreNFT()

	// one : https://ipfs.io/ipfs/bafkreidjtxhcmm3mnynv76qhu5fqdqnohroylliwzsr43xhs67o6xhzg6a
	// two : https://ipfs.io/ipfs/bafkreicdrcv2wuzeyyx2gxn2widjmzp3b7x4kgm64k6346bm6mfnc3axdi
	// three :https://ipfs.io/ipfs/bafkreifns3yxp3cxywpx2frvu2274drk6knt427bikdaf4uobyp26j54lu
	//mynft.MintNFT("https://ipfs.io/ipfs/bafkreicdrcv2wuzeyyx2gxn2widjmzp3b7x4kgm64k6346bm6mfnc3axdi")

	mynft.Apprave(11)

}
