// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mynft

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// MyNFT721MetaData contains all meta data concerning the MyNFT721 contract.
var MyNFT721MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"mintNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506040518060400160405280600581526020017f4d594e46540000000000000000000000000000000000000000000000000000008152506040518060400160405280600481526020017f4d4e465400000000000000000000000000000000000000000000000000000000815250816000908161008c91906102f4565b50806001908161009c91906102f4565b5050506103c6565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061012557607f821691505b602082108103610138576101376100de565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026101a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610163565b6101aa8683610163565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b60006101f16101ec6101e7846101c2565b6101cc565b6101c2565b9050919050565b6000819050919050565b61020b836101d6565b61021f610217826101f8565b848454610170565b825550505050565b600090565b610234610227565b61023f818484610202565b505050565b5b818110156102635761025860008261022c565b600181019050610245565b5050565b601f8211156102a8576102798161013e565b61028284610153565b81016020851015610291578190505b6102a561029d85610153565b830182610244565b50505b505050565b600082821c905092915050565b60006102cb600019846008026102ad565b1980831691505092915050565b60006102e483836102ba565b9150826002028217905092915050565b6102fd826100a4565b67ffffffffffffffff811115610316576103156100af565b5b610320825461010d565b61032b828285610267565b600060209050601f83116001811461035e576000841561034c578287015190505b61035685826102d8565b8655506103be565b601f19841661036c8661013e565b60005b828110156103945784890151825560018201915060208501945060208101905061036f565b868310156103b157848901516103ad601f8916826102ba565b8355505b6001600288020188555050505b505050505050565b611fd1806103d56000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806370a0823111610097578063b88d4fde11610066578063b88d4fde14610284578063c87b56dd146102a0578063e985e9c5146102d0578063eacabe1414610300576100f5565b806370a08231146101fc57806375794a3c1461022c57806395d89b411461024a578063a22cb46514610268576100f5565b8063095ea7b3116100d3578063095ea7b31461017857806323b872dd1461019457806342842e0e146101b05780636352211e146101cc576100f5565b806301ffc9a7146100fa57806306fdde031461012a578063081812fc14610148575b600080fd5b610114600480360381019061010f91906114b8565b61031c565b6040516101219190611500565b60405180910390f35b6101326103fe565b60405161013f91906115ab565b60405180910390f35b610162600480360381019061015d9190611603565b610490565b60405161016f9190611671565b60405180910390f35b610192600480360381019061018d91906116b8565b6104ac565b005b6101ae60048036038101906101a991906116f8565b6104c2565b005b6101ca60048036038101906101c591906116f8565b6105c4565b005b6101e660048036038101906101e19190611603565b6105e4565b6040516101f39190611671565b60405180910390f35b6102166004803603810190610211919061174b565b6105f6565b6040516102239190611787565b60405180910390f35b6102346106b0565b6040516102419190611787565b60405180910390f35b6102526106b6565b60405161025f91906115ab565b60405180910390f35b610282600480360381019061027d91906117ce565b610748565b005b61029e60048036038101906102999190611943565b61075e565b005b6102ba60048036038101906102b59190611603565b610783565b6040516102c791906115ab565b60405180910390f35b6102ea60048036038101906102e591906119c6565b610828565b6040516102f79190611500565b60405180910390f35b61031a60048036038101906103159190611aa7565b6108bc565b005b60007f80ac58cd000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806103e757507f5b5e139f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b806103f757506103f682610907565b5b9050919050565b60606000805461040d90611b32565b80601f016020809104026020016040519081016040528092919081815260200182805461043990611b32565b80156104865780601f1061045b57610100808354040283529160200191610486565b820191906000526020600020905b81548152906001019060200180831161046957829003601f168201915b5050505050905090565b600061049b82610971565b506104a5826109f9565b9050919050565b6104be82826104b9610a36565b610a3e565b5050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036105345760006040517f64a0ae9200000000000000000000000000000000000000000000000000000000815260040161052b9190611671565b60405180910390fd5b60006105488383610543610a36565b610a50565b90508373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146105be578382826040517f64283d7b0000000000000000000000000000000000000000000000000000000081526004016105b593929190611b63565b60405180910390fd5b50505050565b6105df8383836040518060200160405280600081525061075e565b505050565b60006105ef82610971565b9050919050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036106695760006040517f89c62b640000000000000000000000000000000000000000000000000000000081526004016106609190611671565b60405180910390fd5b600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60065481565b6060600180546106c590611b32565b80601f01602080910402602001604051908101604052809291908181526020018280546106f190611b32565b801561073e5780601f106107135761010080835404028352916020019161073e565b820191906000526020600020905b81548152906001019060200180831161072157829003601f168201915b5050505050905090565b61075a610753610a36565b8383610c6a565b5050565b6107698484846104c2565b61077d610774610a36565b85858585610dd9565b50505050565b60606007600083815260200190815260200160002080546107a390611b32565b80601f01602080910402602001604051908101604052809291908181526020018280546107cf90611b32565b801561081c5780601f106107f15761010080835404028352916020019161081c565b820191906000526020600020905b8154815290600101906020018083116107ff57829003601f168201915b50505050509050919050565b6000600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b6108c882600654610f8a565b8060076000600654815260200190815260200160002090816108ea9190611d46565b50600660008154809291906108fe90611e47565b91905055505050565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b60008061097d83610fa8565b9050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036109f057826040517f7e2732890000000000000000000000000000000000000000000000000000000081526004016109e79190611787565b60405180910390fd5b80915050919050565b60006004600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b600033905090565b610a4b8383836001610fe5565b505050565b600080610a5c84610fa8565b9050600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614610a9e57610a9d8184866111aa565b5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610b2f57610ae0600085600080610fe5565b6001600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825403925050819055505b600073ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614610bb2576001600360008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055505b846002600086815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4809150509392505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610cdb57816040517f5b08ba18000000000000000000000000000000000000000000000000000000008152600401610cd29190611671565b60405180910390fd5b80600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c3183604051610dcc9190611500565b60405180910390a3505050565b60008373ffffffffffffffffffffffffffffffffffffffff163b1115610f83578273ffffffffffffffffffffffffffffffffffffffff1663150b7a02868685856040518563ffffffff1660e01b8152600401610e389493929190611ee4565b6020604051808303816000875af1925050508015610e7457506040513d601f19601f82011682018060405250810190610e719190611f45565b60015b610ef8573d8060008114610ea4576040519150601f19603f3d011682016040523d82523d6000602084013e610ea9565b606091505b506000815103610ef057836040517f64a0ae92000000000000000000000000000000000000000000000000000000008152600401610ee79190611671565b60405180910390fd5b805160208201fd5b63150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614610f8157836040517f64a0ae92000000000000000000000000000000000000000000000000000000008152600401610f789190611671565b60405180910390fd5b505b5050505050565b610fa482826040518060200160405280600081525061126e565b5050565b60006002600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b808061101e5750600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b1561115257600061102e84610971565b9050600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415801561109957508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b80156110ac57506110aa8184610828565b155b156110ee57826040517fa9fbf51f0000000000000000000000000000000000000000000000000000000081526004016110e59190611671565b60405180910390fd5b811561115057838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45b505b836004600085815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050565b6111b5838383611292565b61126957600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361122a57806040517f7e2732890000000000000000000000000000000000000000000000000000000081526004016112219190611787565b60405180910390fd5b81816040517f177e802f000000000000000000000000000000000000000000000000000000008152600401611260929190611f72565b60405180910390fd5b505050565b6112788383611353565b61128d611283610a36565b6000858585610dd9565b505050565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415801561134a57508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16148061130b575061130a8484610828565b5b8061134957508273ffffffffffffffffffffffffffffffffffffffff16611331836109f9565b73ffffffffffffffffffffffffffffffffffffffff16145b5b90509392505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036113c55760006040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016113bc9190611671565b60405180910390fd5b60006113d383836000610a50565b9050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146114475760006040517f73c6ac6e00000000000000000000000000000000000000000000000000000000815260040161143e9190611671565b60405180910390fd5b505050565b6000604051905090565b600080fd5b600080fd5b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b61149581611460565b81146114a057600080fd5b50565b6000813590506114b28161148c565b92915050565b6000602082840312156114ce576114cd611456565b5b60006114dc848285016114a3565b91505092915050565b60008115159050919050565b6114fa816114e5565b82525050565b600060208201905061151560008301846114f1565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561155557808201518184015260208101905061153a565b60008484015250505050565b6000601f19601f8301169050919050565b600061157d8261151b565b6115878185611526565b9350611597818560208601611537565b6115a081611561565b840191505092915050565b600060208201905081810360008301526115c58184611572565b905092915050565b6000819050919050565b6115e0816115cd565b81146115eb57600080fd5b50565b6000813590506115fd816115d7565b92915050565b60006020828403121561161957611618611456565b5b6000611627848285016115ee565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061165b82611630565b9050919050565b61166b81611650565b82525050565b60006020820190506116866000830184611662565b92915050565b61169581611650565b81146116a057600080fd5b50565b6000813590506116b28161168c565b92915050565b600080604083850312156116cf576116ce611456565b5b60006116dd858286016116a3565b92505060206116ee858286016115ee565b9150509250929050565b60008060006060848603121561171157611710611456565b5b600061171f868287016116a3565b9350506020611730868287016116a3565b9250506040611741868287016115ee565b9150509250925092565b60006020828403121561176157611760611456565b5b600061176f848285016116a3565b91505092915050565b611781816115cd565b82525050565b600060208201905061179c6000830184611778565b92915050565b6117ab816114e5565b81146117b657600080fd5b50565b6000813590506117c8816117a2565b92915050565b600080604083850312156117e5576117e4611456565b5b60006117f3858286016116a3565b9250506020611804858286016117b9565b9150509250929050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61185082611561565b810181811067ffffffffffffffff8211171561186f5761186e611818565b5b80604052505050565b600061188261144c565b905061188e8282611847565b919050565b600067ffffffffffffffff8211156118ae576118ad611818565b5b6118b782611561565b9050602081019050919050565b82818337600083830152505050565b60006118e66118e184611893565b611878565b90508281526020810184848401111561190257611901611813565b5b61190d8482856118c4565b509392505050565b600082601f83011261192a5761192961180e565b5b813561193a8482602086016118d3565b91505092915050565b6000806000806080858703121561195d5761195c611456565b5b600061196b878288016116a3565b945050602061197c878288016116a3565b935050604061198d878288016115ee565b925050606085013567ffffffffffffffff8111156119ae576119ad61145b565b5b6119ba87828801611915565b91505092959194509250565b600080604083850312156119dd576119dc611456565b5b60006119eb858286016116a3565b92505060206119fc858286016116a3565b9150509250929050565b600067ffffffffffffffff821115611a2157611a20611818565b5b611a2a82611561565b9050602081019050919050565b6000611a4a611a4584611a06565b611878565b905082815260208101848484011115611a6657611a65611813565b5b611a718482856118c4565b509392505050565b600082601f830112611a8e57611a8d61180e565b5b8135611a9e848260208601611a37565b91505092915050565b60008060408385031215611abe57611abd611456565b5b6000611acc858286016116a3565b925050602083013567ffffffffffffffff811115611aed57611aec61145b565b5b611af985828601611a79565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680611b4a57607f821691505b602082108103611b5d57611b5c611b03565b5b50919050565b6000606082019050611b786000830186611662565b611b856020830185611778565b611b926040830184611662565b949350505050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302611bfc7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611bbf565b611c068683611bbf565b95508019841693508086168417925050509392505050565b6000819050919050565b6000611c43611c3e611c39846115cd565b611c1e565b6115cd565b9050919050565b6000819050919050565b611c5d83611c28565b611c71611c6982611c4a565b848454611bcc565b825550505050565b600090565b611c86611c79565b611c91818484611c54565b505050565b5b81811015611cb557611caa600082611c7e565b600181019050611c97565b5050565b601f821115611cfa57611ccb81611b9a565b611cd484611baf565b81016020851015611ce3578190505b611cf7611cef85611baf565b830182611c96565b50505b505050565b600082821c905092915050565b6000611d1d60001984600802611cff565b1980831691505092915050565b6000611d368383611d0c565b9150826002028217905092915050565b611d4f8261151b565b67ffffffffffffffff811115611d6857611d67611818565b5b611d728254611b32565b611d7d828285611cb9565b600060209050601f831160018114611db05760008415611d9e578287015190505b611da88582611d2a565b865550611e10565b601f198416611dbe86611b9a565b60005b82811015611de657848901518255600182019150602085019450602081019050611dc1565b86831015611e035784890151611dff601f891682611d0c565b8355505b6001600288020188555050505b505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611e52826115cd565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611e8457611e83611e18565b5b600182019050919050565b600081519050919050565b600082825260208201905092915050565b6000611eb682611e8f565b611ec08185611e9a565b9350611ed0818560208601611537565b611ed981611561565b840191505092915050565b6000608082019050611ef96000830187611662565b611f066020830186611662565b611f136040830185611778565b8181036060830152611f258184611eab565b905095945050505050565b600081519050611f3f8161148c565b92915050565b600060208284031215611f5b57611f5a611456565b5b6000611f6984828501611f30565b91505092915050565b6000604082019050611f876000830185611662565b611f946020830184611778565b939250505056fea2646970667358221220ba2027466bb6b6fcca258d093ce930daa580708bc1ef37c4b1ea28004dbb65a164736f6c634300081c0033",
}

// MyNFT721ABI is the input ABI used to generate the binding from.
// Deprecated: Use MyNFT721MetaData.ABI instead.
var MyNFT721ABI = MyNFT721MetaData.ABI

// MyNFT721Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MyNFT721MetaData.Bin instead.
var MyNFT721Bin = MyNFT721MetaData.Bin

// DeployMyNFT721 deploys a new Ethereum contract, binding an instance of MyNFT721 to it.
func DeployMyNFT721(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MyNFT721, error) {
	parsed, err := MyNFT721MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MyNFT721Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MyNFT721{MyNFT721Caller: MyNFT721Caller{contract: contract}, MyNFT721Transactor: MyNFT721Transactor{contract: contract}, MyNFT721Filterer: MyNFT721Filterer{contract: contract}}, nil
}

// MyNFT721 is an auto generated Go binding around an Ethereum contract.
type MyNFT721 struct {
	MyNFT721Caller     // Read-only binding to the contract
	MyNFT721Transactor // Write-only binding to the contract
	MyNFT721Filterer   // Log filterer for contract events
}

// MyNFT721Caller is an auto generated read-only Go binding around an Ethereum contract.
type MyNFT721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyNFT721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MyNFT721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyNFT721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MyNFT721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyNFT721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MyNFT721Session struct {
	Contract     *MyNFT721         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MyNFT721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MyNFT721CallerSession struct {
	Contract *MyNFT721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MyNFT721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MyNFT721TransactorSession struct {
	Contract     *MyNFT721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MyNFT721Raw is an auto generated low-level Go binding around an Ethereum contract.
type MyNFT721Raw struct {
	Contract *MyNFT721 // Generic contract binding to access the raw methods on
}

// MyNFT721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MyNFT721CallerRaw struct {
	Contract *MyNFT721Caller // Generic read-only contract binding to access the raw methods on
}

// MyNFT721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MyNFT721TransactorRaw struct {
	Contract *MyNFT721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMyNFT721 creates a new instance of MyNFT721, bound to a specific deployed contract.
func NewMyNFT721(address common.Address, backend bind.ContractBackend) (*MyNFT721, error) {
	contract, err := bindMyNFT721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MyNFT721{MyNFT721Caller: MyNFT721Caller{contract: contract}, MyNFT721Transactor: MyNFT721Transactor{contract: contract}, MyNFT721Filterer: MyNFT721Filterer{contract: contract}}, nil
}

// NewMyNFT721Caller creates a new read-only instance of MyNFT721, bound to a specific deployed contract.
func NewMyNFT721Caller(address common.Address, caller bind.ContractCaller) (*MyNFT721Caller, error) {
	contract, err := bindMyNFT721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MyNFT721Caller{contract: contract}, nil
}

// NewMyNFT721Transactor creates a new write-only instance of MyNFT721, bound to a specific deployed contract.
func NewMyNFT721Transactor(address common.Address, transactor bind.ContractTransactor) (*MyNFT721Transactor, error) {
	contract, err := bindMyNFT721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MyNFT721Transactor{contract: contract}, nil
}

// NewMyNFT721Filterer creates a new log filterer instance of MyNFT721, bound to a specific deployed contract.
func NewMyNFT721Filterer(address common.Address, filterer bind.ContractFilterer) (*MyNFT721Filterer, error) {
	contract, err := bindMyNFT721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MyNFT721Filterer{contract: contract}, nil
}

// bindMyNFT721 binds a generic wrapper to an already deployed contract.
func bindMyNFT721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MyNFT721MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyNFT721 *MyNFT721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyNFT721.Contract.MyNFT721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyNFT721 *MyNFT721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyNFT721.Contract.MyNFT721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyNFT721 *MyNFT721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyNFT721.Contract.MyNFT721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyNFT721 *MyNFT721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyNFT721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyNFT721 *MyNFT721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyNFT721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyNFT721 *MyNFT721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyNFT721.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MyNFT721 *MyNFT721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MyNFT721.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MyNFT721 *MyNFT721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MyNFT721.Contract.BalanceOf(&_MyNFT721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_MyNFT721 *MyNFT721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MyNFT721.Contract.BalanceOf(&_MyNFT721.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_MyNFT721 *MyNFT721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MyNFT721.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_MyNFT721 *MyNFT721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _MyNFT721.Contract.GetApproved(&_MyNFT721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_MyNFT721 *MyNFT721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _MyNFT721.Contract.GetApproved(&_MyNFT721.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MyNFT721 *MyNFT721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _MyNFT721.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MyNFT721 *MyNFT721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _MyNFT721.Contract.IsApprovedForAll(&_MyNFT721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_MyNFT721 *MyNFT721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _MyNFT721.Contract.IsApprovedForAll(&_MyNFT721.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MyNFT721 *MyNFT721Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MyNFT721.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MyNFT721 *MyNFT721Session) Name() (string, error) {
	return _MyNFT721.Contract.Name(&_MyNFT721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MyNFT721 *MyNFT721CallerSession) Name() (string, error) {
	return _MyNFT721.Contract.Name(&_MyNFT721.CallOpts)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_MyNFT721 *MyNFT721Caller) NextTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MyNFT721.contract.Call(opts, &out, "nextTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_MyNFT721 *MyNFT721Session) NextTokenId() (*big.Int, error) {
	return _MyNFT721.Contract.NextTokenId(&_MyNFT721.CallOpts)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_MyNFT721 *MyNFT721CallerSession) NextTokenId() (*big.Int, error) {
	return _MyNFT721.Contract.NextTokenId(&_MyNFT721.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_MyNFT721 *MyNFT721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MyNFT721.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_MyNFT721 *MyNFT721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _MyNFT721.Contract.OwnerOf(&_MyNFT721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_MyNFT721 *MyNFT721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _MyNFT721.Contract.OwnerOf(&_MyNFT721.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MyNFT721 *MyNFT721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MyNFT721.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MyNFT721 *MyNFT721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MyNFT721.Contract.SupportsInterface(&_MyNFT721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MyNFT721 *MyNFT721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MyNFT721.Contract.SupportsInterface(&_MyNFT721.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MyNFT721 *MyNFT721Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MyNFT721.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MyNFT721 *MyNFT721Session) Symbol() (string, error) {
	return _MyNFT721.Contract.Symbol(&_MyNFT721.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MyNFT721 *MyNFT721CallerSession) Symbol() (string, error) {
	return _MyNFT721.Contract.Symbol(&_MyNFT721.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_MyNFT721 *MyNFT721Caller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _MyNFT721.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_MyNFT721 *MyNFT721Session) TokenURI(tokenId *big.Int) (string, error) {
	return _MyNFT721.Contract.TokenURI(&_MyNFT721.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_MyNFT721 *MyNFT721CallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _MyNFT721.Contract.TokenURI(&_MyNFT721.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_MyNFT721 *MyNFT721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MyNFT721.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_MyNFT721 *MyNFT721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MyNFT721.Contract.Approve(&_MyNFT721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_MyNFT721 *MyNFT721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MyNFT721.Contract.Approve(&_MyNFT721.TransactOpts, to, tokenId)
}

// MintNFT is a paid mutator transaction binding the contract method 0xeacabe14.
//
// Solidity: function mintNFT(address to, string uri) returns()
func (_MyNFT721 *MyNFT721Transactor) MintNFT(opts *bind.TransactOpts, to common.Address, uri string) (*types.Transaction, error) {
	return _MyNFT721.contract.Transact(opts, "mintNFT", to, uri)
}

// MintNFT is a paid mutator transaction binding the contract method 0xeacabe14.
//
// Solidity: function mintNFT(address to, string uri) returns()
func (_MyNFT721 *MyNFT721Session) MintNFT(to common.Address, uri string) (*types.Transaction, error) {
	return _MyNFT721.Contract.MintNFT(&_MyNFT721.TransactOpts, to, uri)
}

// MintNFT is a paid mutator transaction binding the contract method 0xeacabe14.
//
// Solidity: function mintNFT(address to, string uri) returns()
func (_MyNFT721 *MyNFT721TransactorSession) MintNFT(to common.Address, uri string) (*types.Transaction, error) {
	return _MyNFT721.Contract.MintNFT(&_MyNFT721.TransactOpts, to, uri)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_MyNFT721 *MyNFT721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MyNFT721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_MyNFT721 *MyNFT721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MyNFT721.Contract.SafeTransferFrom(&_MyNFT721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_MyNFT721 *MyNFT721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MyNFT721.Contract.SafeTransferFrom(&_MyNFT721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_MyNFT721 *MyNFT721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _MyNFT721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_MyNFT721 *MyNFT721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _MyNFT721.Contract.SafeTransferFrom0(&_MyNFT721.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_MyNFT721 *MyNFT721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _MyNFT721.Contract.SafeTransferFrom0(&_MyNFT721.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MyNFT721 *MyNFT721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _MyNFT721.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MyNFT721 *MyNFT721Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _MyNFT721.Contract.SetApprovalForAll(&_MyNFT721.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_MyNFT721 *MyNFT721TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _MyNFT721.Contract.SetApprovalForAll(&_MyNFT721.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_MyNFT721 *MyNFT721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MyNFT721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_MyNFT721 *MyNFT721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MyNFT721.Contract.TransferFrom(&_MyNFT721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_MyNFT721 *MyNFT721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _MyNFT721.Contract.TransferFrom(&_MyNFT721.TransactOpts, from, to, tokenId)
}

// MyNFT721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MyNFT721 contract.
type MyNFT721ApprovalIterator struct {
	Event *MyNFT721Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyNFT721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyNFT721Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyNFT721Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyNFT721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyNFT721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyNFT721Approval represents a Approval event raised by the MyNFT721 contract.
type MyNFT721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_MyNFT721 *MyNFT721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*MyNFT721ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MyNFT721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MyNFT721ApprovalIterator{contract: _MyNFT721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_MyNFT721 *MyNFT721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MyNFT721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MyNFT721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyNFT721Approval)
				if err := _MyNFT721.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_MyNFT721 *MyNFT721Filterer) ParseApproval(log types.Log) (*MyNFT721Approval, error) {
	event := new(MyNFT721Approval)
	if err := _MyNFT721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyNFT721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the MyNFT721 contract.
type MyNFT721ApprovalForAllIterator struct {
	Event *MyNFT721ApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyNFT721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyNFT721ApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyNFT721ApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyNFT721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyNFT721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyNFT721ApprovalForAll represents a ApprovalForAll event raised by the MyNFT721 contract.
type MyNFT721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_MyNFT721 *MyNFT721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*MyNFT721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _MyNFT721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &MyNFT721ApprovalForAllIterator{contract: _MyNFT721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_MyNFT721 *MyNFT721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *MyNFT721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _MyNFT721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyNFT721ApprovalForAll)
				if err := _MyNFT721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_MyNFT721 *MyNFT721Filterer) ParseApprovalForAll(log types.Log) (*MyNFT721ApprovalForAll, error) {
	event := new(MyNFT721ApprovalForAll)
	if err := _MyNFT721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MyNFT721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MyNFT721 contract.
type MyNFT721TransferIterator struct {
	Event *MyNFT721Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyNFT721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyNFT721Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyNFT721Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyNFT721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyNFT721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyNFT721Transfer represents a Transfer event raised by the MyNFT721 contract.
type MyNFT721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_MyNFT721 *MyNFT721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*MyNFT721TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MyNFT721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MyNFT721TransferIterator{contract: _MyNFT721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_MyNFT721 *MyNFT721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MyNFT721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _MyNFT721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyNFT721Transfer)
				if err := _MyNFT721.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_MyNFT721 *MyNFT721Filterer) ParseTransfer(log types.Log) (*MyNFT721Transfer, error) {
	event := new(MyNFT721Transfer)
	if err := _MyNFT721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
