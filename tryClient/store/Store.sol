// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;
//编译Store合约 将编译后的二进制字节码文件输出到tryClient/store目录下
// solcjs --bin tryClient/store/Store.sol -o tryClient/store
//生成Store合约的ABI文件 将abi文件输出到tryClient/store目录下
// solcjs --abi tryClient/store/Store.sol -o tryClient/store

//使用abigen工具根据Store合约的ABI和二进制文件生成Go语言绑定代码
// 先安装abigen工具
//go install github.com/ethereum/go-ethereum/cmd/abigen@latest
// 运行abigen命令生成Go绑定代码
// abigen --bin tryClient/store/Store_sol.bin --abi tryClient/store/Store_sol.abi --pkg=tryClient --out=tryClient/store/store_abigen.go
// abigen --bin tryClient/myNFT/ABIAndBIN/myAuctionUpBIN.bin --abi tryClient/myNFT/ABIAndBIN/myAuctionUpABI.abi --pkg=mynft --out=tryClient/myNFT/myAuctionUp.go
// abigen -bin tryClient/myNFT/ABIAndBIN/erc721.bin --abi tryClient/myNFT/ABIAndBIN/erc721.abi --pkg=myNFT721 --out=tryClient/myNFT/myNFT.go
contract Store {
  event ItemSet(bytes32 key, bytes32 value);

  string public version;
  mapping (bytes32 => bytes32) public items;

  constructor(string memory _version) {
    version = _version;
  }

  function setItem(bytes32 key, bytes32 value) external {
    items[key] = value;
    emit ItemSet(key, value);
  }
}
//echo $PATH
/**
  添加 Go 的 bin 目录到 PATH
  export PATH=$PATH:$HOME/go/bin
  export PATH=$PATH:/usr/local/go/bin

  使配置生效
  source ~/.zshrc
 */