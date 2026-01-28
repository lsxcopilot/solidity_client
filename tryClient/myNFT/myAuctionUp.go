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

// MynftMetaData contains all meta data concerning the Mynft contract.
var MynftMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"auctionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"auctionMapping\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startingPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"highestBidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duringTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"ended\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duringTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftContractAddress\",\"type\":\"address\"}],\"name\":\"createAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionId\",\"type\":\"uint256\"}],\"name\":\"endAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"getLatestPrice\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"getTokenDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_priceFeed\",\"type\":\"address\"}],\"name\":\"setTokenPriceFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenPriceFeedMapping\",\"outputs\":[{\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50611f908061001f6000396000f3fe6080604052600436106100915760003560e01c8063742f0a9011610059578063742f0a90146101aa578063785c7cf6146101c65780638129fc1c14610203578063819c9e4c1461021a578063b9a2de3a1461024357610091565b80630f265dc01461009657806310782f8f146100dc57806316345f181461010757806327927b3e14610144578063674417ae14610181575b600080fd5b3480156100a257600080fd5b506100bd60048036038101906100b8919061140d565b61026c565b6040516100d39a999897969594939291906114a5565b60405180910390f35b3480156100e857600080fd5b506100f161034d565b6040516100fe9190611541565b60405180910390f35b34801561011357600080fd5b5061012e60048036038101906101299190611588565b610353565b60405161013b91906115ce565b60405180910390f35b34801561015057600080fd5b5061016b60048036038101906101669190611588565b610497565b6040516101789190611648565b60405180910390f35b34801561018d57600080fd5b506101a860048036038101906101a39190611663565b6104ca565b005b6101c460048036038101906101bf91906116a3565b6105dc565b005b3480156101d257600080fd5b506101ed60048036038101906101e89190611588565b610bcc565b6040516101fa9190611541565b60405180910390f35b34801561020f57600080fd5b50610218610c5f565b005b34801561022657600080fd5b50610241600480360381019061023c91906116f6565b610dd6565b005b34801561024f57600080fd5b5061026a6004803603810190610265919061140d565b6110d2565b005b60016020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010154908060020154908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154908060050154908060060160009054906101000a900460ff16908060060160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060070154908060080160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508a565b60025481565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663feaf968c6040518163ffffffff1660e01b815260040160a060405180830381865afa158015610464573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061048891906117e0565b50505091505080915050919050565b60046020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600060029054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461055a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610551906118b8565b60405180910390fd5b80600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b6000600160008581526020019081526020016000209050806004015481600501546106079190611907565b4210801561062457508060060160009054906101000a900460ff16155b610663576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161065a90611987565b60405180910390fd5b60008060129050600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614610753578373ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa1580156106e9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061070d91906119e0565b60ff1690506305f5e10081600a6107249190611b40565b61072d86610353565b876107389190611b8b565b6107429190611bfc565b61074c9190611bfc565b9150610790565b3494506305f5e100670de0b6b3a764000061076e6000610353565b876107799190611b8b565b6107839190611bfc565b61078d9190611bfc565b91505b60006305f5e100670de0b6b3a76400006107aa6000610353565b86600101546107b99190611b8b565b6107c39190611bfc565b6107cd9190611bfc565b905060008073ffffffffffffffffffffffffffffffffffffffff168560060160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610891576305f5e10083600a61083a9190611b40565b6108678760060160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16610353565b87600201546108769190611b8b565b6108809190611bfc565b61088a9190611bfc565b90506108f2565b6305f5e100670de0b6b3a76400006108cc8760060160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16610353565b87600201546108db9190611b8b565b6108e59190611bfc565b6108ef9190611bfc565b90505b81841015801561090157508084115b610940576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161093790611c79565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168560060160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603610a0c578460030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc86600201549081150290604051600060405180830381858888f19350505050158015610a06573d6000803e3d6000fd5b50610b33565b8573ffffffffffffffffffffffffffffffffffffffff166323b872dd308760030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168a6040518463ffffffff1660e01b8152600401610a6d93929190611c99565b6020604051808303816000875af1158015610a8c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ab09190611cfc565b508573ffffffffffffffffffffffffffffffffffffffff166323b872dd33308a6040518463ffffffff1660e01b8152600401610aee93929190611c99565b6020604051808303816000875af1158015610b0d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b319190611cfc565b505b348560020181905550338560030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550858560060160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050505050505050565b6000808273ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c1a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c3e91906119e0565b60ff16905080600003610c55576012915050610c5a565b809150505b919050565b60008060019054906101000a900460ff16159050808015610c905750600160008054906101000a900460ff1660ff16105b80610cbd5750610c9f306113af565b158015610cbc5750600160008054906101000a900460ff1660ff16145b5b610cfc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cf390611d9b565b60405180910390fd5b60016000806101000a81548160ff021916908360ff1602179055508015610d39576001600060016101000a81548160ff0219169083151502179055505b33600060026101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508015610dd35760008060016101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024986001604051610dca9190611df6565b60405180910390a15b50565b60008411610e19576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e1090611e5d565b60405180910390fd5b6040518061014001604052803373ffffffffffffffffffffffffffffffffffffffff16815260200185815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001848152602001428152602001600015158152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020018381526020018273ffffffffffffffffffffffffffffffffffffffff1681525060016000600254815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101556040820151816002015560608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506080820151816004015560a0820151816005015560c08201518160060160006101000a81548160ff02191690831515021790555060e08201518160060160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061010082015181600701556101208201518160080160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509050506002600081548092919061105890611e7d565b91905055508073ffffffffffffffffffffffffffffffffffffffff166323b872dd3330856040518463ffffffff1660e01b815260040161109a93929190611c99565b600060405180830381600087803b1580156110b457600080fd5b505af11580156110c8573d6000803e3d6000fd5b5050505050505050565b6000600160008381526020019081526020016000209050806004015481600501546110fd9190611907565b42101561113f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161113690611f11565b60405180910390fd5b8060080160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd308360030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684600701546040518463ffffffff1660e01b81526004016111c893929190611c99565b600060405180830381600087803b1580156111e257600080fd5b505af11580156111f6573d6000803e3d6000fd5b50505050600073ffffffffffffffffffffffffffffffffffffffff168160060160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036112c2578060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc479081150290604051600060405180830381858888f193505050501580156112bc573d6000803e3d6000fd5b5061138e565b8060060160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8260000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1683600201546040518363ffffffff1660e01b8152600401611349929190611f31565b6020604051808303816000875af1158015611368573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061138c9190611cfc565b505b60018160060160006101000a81548160ff0219169083151502179055505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600080fd5b6000819050919050565b6113ea816113d7565b81146113f557600080fd5b50565b600081359050611407816113e1565b92915050565b600060208284031215611423576114226113d2565b5b6000611431848285016113f8565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006114658261143a565b9050919050565b6114758161145a565b82525050565b611484816113d7565b82525050565b60008115159050919050565b61149f8161148a565b82525050565b6000610140820190506114bb600083018d61146c565b6114c8602083018c61147b565b6114d5604083018b61147b565b6114e2606083018a61146c565b6114ef608083018961147b565b6114fc60a083018861147b565b61150960c0830187611496565b61151660e083018661146c565b61152461010083018561147b565b61153261012083018461146c565b9b9a5050505050505050505050565b6000602082019050611556600083018461147b565b92915050565b6115658161145a565b811461157057600080fd5b50565b6000813590506115828161155c565b92915050565b60006020828403121561159e5761159d6113d2565b5b60006115ac84828501611573565b91505092915050565b6000819050919050565b6115c8816115b5565b82525050565b60006020820190506115e360008301846115bf565b92915050565b6000819050919050565b600061160e6116096116048461143a565b6115e9565b61143a565b9050919050565b6000611620826115f3565b9050919050565b600061163282611615565b9050919050565b61164281611627565b82525050565b600060208201905061165d6000830184611639565b92915050565b6000806040838503121561167a576116796113d2565b5b600061168885828601611573565b925050602061169985828601611573565b9150509250929050565b6000806000606084860312156116bc576116bb6113d2565b5b60006116ca868287016113f8565b93505060206116db868287016113f8565b92505060406116ec86828701611573565b9150509250925092565b600080600080608085870312156117105761170f6113d2565b5b600061171e878288016113f8565b945050602061172f878288016113f8565b9350506040611740878288016113f8565b925050606061175187828801611573565b91505092959194509250565b600069ffffffffffffffffffff82169050919050565b61177c8161175d565b811461178757600080fd5b50565b60008151905061179981611773565b92915050565b6117a8816115b5565b81146117b357600080fd5b50565b6000815190506117c58161179f565b92915050565b6000815190506117da816113e1565b92915050565b600080600080600060a086880312156117fc576117fb6113d2565b5b600061180a8882890161178a565b955050602061181b888289016117b6565b945050604061182c888289016117cb565b935050606061183d888289016117cb565b925050608061184e8882890161178a565b9150509295509295909350565b600082825260208201905092915050565b7f4f6e6c792061646d696e2063616e207365742070726963652066656564000000600082015250565b60006118a2601d8361185b565b91506118ad8261186c565b602082019050919050565b600060208201905081810360008301526118d181611895565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611912826113d7565b915061191d836113d7565b9250828201905080821115611935576119346118d8565b5b92915050565b7f41756374696f6e2068617320656e646564000000000000000000000000000000600082015250565b600061197160118361185b565b915061197c8261193b565b602082019050919050565b600060208201905081810360008301526119a081611964565b9050919050565b600060ff82169050919050565b6119bd816119a7565b81146119c857600080fd5b50565b6000815190506119da816119b4565b92915050565b6000602082840312156119f6576119f56113d2565b5b6000611a04848285016119cb565b91505092915050565b60008160011c9050919050565b6000808291508390505b6001851115611a6457808604811115611a4057611a3f6118d8565b5b6001851615611a4f5780820291505b8081029050611a5d85611a0d565b9450611a24565b94509492505050565b600082611a7d5760019050611b39565b81611a8b5760009050611b39565b8160018114611aa15760028114611aab57611ada565b6001915050611b39565b60ff841115611abd57611abc6118d8565b5b8360020a915084821115611ad457611ad36118d8565b5b50611b39565b5060208310610133831016604e8410600b8410161715611b0f5782820a905083811115611b0a57611b096118d8565b5b611b39565b611b1c8484846001611a1a565b92509050818404811115611b3357611b326118d8565b5b81810290505b9392505050565b6000611b4b826113d7565b9150611b56836113d7565b9250611b837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484611a6d565b905092915050565b6000611b96826113d7565b9150611ba1836113d7565b9250828202611baf816113d7565b91508282048414831517611bc657611bc56118d8565b5b5092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000611c07826113d7565b9150611c12836113d7565b925082611c2257611c21611bcd565b5b828204905092915050565b7f42696420616d6f756e7420697320746f6f206c6f770000000000000000000000600082015250565b6000611c6360158361185b565b9150611c6e82611c2d565b602082019050919050565b60006020820190508181036000830152611c9281611c56565b9050919050565b6000606082019050611cae600083018661146c565b611cbb602083018561146c565b611cc8604083018461147b565b949350505050565b611cd98161148a565b8114611ce457600080fd5b50565b600081519050611cf681611cd0565b92915050565b600060208284031215611d1257611d116113d2565b5b6000611d2084828501611ce7565b91505092915050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b6000611d85602e8361185b565b9150611d9082611d29565b604082019050919050565b60006020820190508181036000830152611db481611d78565b9050919050565b6000819050919050565b6000611de0611ddb611dd684611dbb565b6115e9565b6119a7565b9050919050565b611df081611dc5565b82525050565b6000602082019050611e0b6000830184611de7565b92915050565b7f6e6f20686176652070726963652063616e206c657373207468616e207a65726f600082015250565b6000611e4760208361185b565b9150611e5282611e11565b602082019050919050565b60006020820190508181036000830152611e7681611e3a565b9050919050565b6000611e88826113d7565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611eba57611eb96118d8565b5b600182019050919050565b7f41756374696f6e206973207374696c6c206f6e676f696e670000000000000000600082015250565b6000611efb60188361185b565b9150611f0682611ec5565b602082019050919050565b60006020820190508181036000830152611f2a81611eee565b9050919050565b6000604082019050611f46600083018561146c565b611f53602083018461147b565b939250505056fea264697066735822122039463f0a53ae759563156c1e1e1cac3d9ab66f359f5f61b0afda7020d546099864736f6c634300081c0033",
}

// MynftABI is the input ABI used to generate the binding from.
// Deprecated: Use MynftMetaData.ABI instead.
var MynftABI = MynftMetaData.ABI

// MynftBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MynftMetaData.Bin instead.
var MynftBin = MynftMetaData.Bin

// DeployMynft deploys a new Ethereum contract, binding an instance of Mynft to it.
func DeployMynft(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Mynft, error) {
	parsed, err := MynftMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MynftBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Mynft{MynftCaller: MynftCaller{contract: contract}, MynftTransactor: MynftTransactor{contract: contract}, MynftFilterer: MynftFilterer{contract: contract}}, nil
}

// Mynft is an auto generated Go binding around an Ethereum contract.
type Mynft struct {
	MynftCaller     // Read-only binding to the contract
	MynftTransactor // Write-only binding to the contract
	MynftFilterer   // Log filterer for contract events
}

// MynftCaller is an auto generated read-only Go binding around an Ethereum contract.
type MynftCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MynftTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MynftTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MynftFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MynftFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MynftSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MynftSession struct {
	Contract     *Mynft            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MynftCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MynftCallerSession struct {
	Contract *MynftCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MynftTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MynftTransactorSession struct {
	Contract     *MynftTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MynftRaw is an auto generated low-level Go binding around an Ethereum contract.
type MynftRaw struct {
	Contract *Mynft // Generic contract binding to access the raw methods on
}

// MynftCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MynftCallerRaw struct {
	Contract *MynftCaller // Generic read-only contract binding to access the raw methods on
}

// MynftTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MynftTransactorRaw struct {
	Contract *MynftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMynft creates a new instance of Mynft, bound to a specific deployed contract.
func NewMynft(address common.Address, backend bind.ContractBackend) (*Mynft, error) {
	contract, err := bindMynft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mynft{MynftCaller: MynftCaller{contract: contract}, MynftTransactor: MynftTransactor{contract: contract}, MynftFilterer: MynftFilterer{contract: contract}}, nil
}

// NewMynftCaller creates a new read-only instance of Mynft, bound to a specific deployed contract.
func NewMynftCaller(address common.Address, caller bind.ContractCaller) (*MynftCaller, error) {
	contract, err := bindMynft(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MynftCaller{contract: contract}, nil
}

// NewMynftTransactor creates a new write-only instance of Mynft, bound to a specific deployed contract.
func NewMynftTransactor(address common.Address, transactor bind.ContractTransactor) (*MynftTransactor, error) {
	contract, err := bindMynft(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MynftTransactor{contract: contract}, nil
}

// NewMynftFilterer creates a new log filterer instance of Mynft, bound to a specific deployed contract.
func NewMynftFilterer(address common.Address, filterer bind.ContractFilterer) (*MynftFilterer, error) {
	contract, err := bindMynft(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MynftFilterer{contract: contract}, nil
}

// bindMynft binds a generic wrapper to an already deployed contract.
func bindMynft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MynftMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mynft *MynftRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mynft.Contract.MynftCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mynft *MynftRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mynft.Contract.MynftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mynft *MynftRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mynft.Contract.MynftTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mynft *MynftCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mynft.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mynft *MynftTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mynft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mynft *MynftTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mynft.Contract.contract.Transact(opts, method, params...)
}

// AuctionId is a free data retrieval call binding the contract method 0x10782f8f.
//
// Solidity: function auctionId() view returns(uint256)
func (_Mynft *MynftCaller) AuctionId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mynft.contract.Call(opts, &out, "auctionId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AuctionId is a free data retrieval call binding the contract method 0x10782f8f.
//
// Solidity: function auctionId() view returns(uint256)
func (_Mynft *MynftSession) AuctionId() (*big.Int, error) {
	return _Mynft.Contract.AuctionId(&_Mynft.CallOpts)
}

// AuctionId is a free data retrieval call binding the contract method 0x10782f8f.
//
// Solidity: function auctionId() view returns(uint256)
func (_Mynft *MynftCallerSession) AuctionId() (*big.Int, error) {
	return _Mynft.Contract.AuctionId(&_Mynft.CallOpts)
}

// AuctionMapping is a free data retrieval call binding the contract method 0x0f265dc0.
//
// Solidity: function auctionMapping(uint256 ) view returns(address seller, uint256 startingPrice, uint256 highestBid, address highestBidder, uint256 duringTime, uint256 startTime, bool ended, address tokenAddress, uint256 tokenId, address nftAddress)
func (_Mynft *MynftCaller) AuctionMapping(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Seller        common.Address
	StartingPrice *big.Int
	HighestBid    *big.Int
	HighestBidder common.Address
	DuringTime    *big.Int
	StartTime     *big.Int
	Ended         bool
	TokenAddress  common.Address
	TokenId       *big.Int
	NftAddress    common.Address
}, error) {
	var out []interface{}
	err := _Mynft.contract.Call(opts, &out, "auctionMapping", arg0)

	outstruct := new(struct {
		Seller        common.Address
		StartingPrice *big.Int
		HighestBid    *big.Int
		HighestBidder common.Address
		DuringTime    *big.Int
		StartTime     *big.Int
		Ended         bool
		TokenAddress  common.Address
		TokenId       *big.Int
		NftAddress    common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Seller = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.StartingPrice = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.HighestBid = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.HighestBidder = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.DuringTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Ended = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.TokenAddress = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.NftAddress = *abi.ConvertType(out[9], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// AuctionMapping is a free data retrieval call binding the contract method 0x0f265dc0.
//
// Solidity: function auctionMapping(uint256 ) view returns(address seller, uint256 startingPrice, uint256 highestBid, address highestBidder, uint256 duringTime, uint256 startTime, bool ended, address tokenAddress, uint256 tokenId, address nftAddress)
func (_Mynft *MynftSession) AuctionMapping(arg0 *big.Int) (struct {
	Seller        common.Address
	StartingPrice *big.Int
	HighestBid    *big.Int
	HighestBidder common.Address
	DuringTime    *big.Int
	StartTime     *big.Int
	Ended         bool
	TokenAddress  common.Address
	TokenId       *big.Int
	NftAddress    common.Address
}, error) {
	return _Mynft.Contract.AuctionMapping(&_Mynft.CallOpts, arg0)
}

// AuctionMapping is a free data retrieval call binding the contract method 0x0f265dc0.
//
// Solidity: function auctionMapping(uint256 ) view returns(address seller, uint256 startingPrice, uint256 highestBid, address highestBidder, uint256 duringTime, uint256 startTime, bool ended, address tokenAddress, uint256 tokenId, address nftAddress)
func (_Mynft *MynftCallerSession) AuctionMapping(arg0 *big.Int) (struct {
	Seller        common.Address
	StartingPrice *big.Int
	HighestBid    *big.Int
	HighestBidder common.Address
	DuringTime    *big.Int
	StartTime     *big.Int
	Ended         bool
	TokenAddress  common.Address
	TokenId       *big.Int
	NftAddress    common.Address
}, error) {
	return _Mynft.Contract.AuctionMapping(&_Mynft.CallOpts, arg0)
}

// GetTokenDecimals is a free data retrieval call binding the contract method 0x785c7cf6.
//
// Solidity: function getTokenDecimals(address tokenAddress) view returns(uint256 decimals)
func (_Mynft *MynftCaller) GetTokenDecimals(opts *bind.CallOpts, tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mynft.contract.Call(opts, &out, "getTokenDecimals", tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenDecimals is a free data retrieval call binding the contract method 0x785c7cf6.
//
// Solidity: function getTokenDecimals(address tokenAddress) view returns(uint256 decimals)
func (_Mynft *MynftSession) GetTokenDecimals(tokenAddress common.Address) (*big.Int, error) {
	return _Mynft.Contract.GetTokenDecimals(&_Mynft.CallOpts, tokenAddress)
}

// GetTokenDecimals is a free data retrieval call binding the contract method 0x785c7cf6.
//
// Solidity: function getTokenDecimals(address tokenAddress) view returns(uint256 decimals)
func (_Mynft *MynftCallerSession) GetTokenDecimals(tokenAddress common.Address) (*big.Int, error) {
	return _Mynft.Contract.GetTokenDecimals(&_Mynft.CallOpts, tokenAddress)
}

// TokenPriceFeedMapping is a free data retrieval call binding the contract method 0x27927b3e.
//
// Solidity: function tokenPriceFeedMapping(address ) view returns(address)
func (_Mynft *MynftCaller) TokenPriceFeedMapping(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Mynft.contract.Call(opts, &out, "tokenPriceFeedMapping", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenPriceFeedMapping is a free data retrieval call binding the contract method 0x27927b3e.
//
// Solidity: function tokenPriceFeedMapping(address ) view returns(address)
func (_Mynft *MynftSession) TokenPriceFeedMapping(arg0 common.Address) (common.Address, error) {
	return _Mynft.Contract.TokenPriceFeedMapping(&_Mynft.CallOpts, arg0)
}

// TokenPriceFeedMapping is a free data retrieval call binding the contract method 0x27927b3e.
//
// Solidity: function tokenPriceFeedMapping(address ) view returns(address)
func (_Mynft *MynftCallerSession) TokenPriceFeedMapping(arg0 common.Address) (common.Address, error) {
	return _Mynft.Contract.TokenPriceFeedMapping(&_Mynft.CallOpts, arg0)
}

// Bid is a paid mutator transaction binding the contract method 0x742f0a90.
//
// Solidity: function bid(uint256 _auctionId, uint256 _bidAmount, address _tokenAddress) payable returns()
func (_Mynft *MynftTransactor) Bid(opts *bind.TransactOpts, _auctionId *big.Int, _bidAmount *big.Int, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Mynft.contract.Transact(opts, "bid", _auctionId, _bidAmount, _tokenAddress)
}

// Bid is a paid mutator transaction binding the contract method 0x742f0a90.
//
// Solidity: function bid(uint256 _auctionId, uint256 _bidAmount, address _tokenAddress) payable returns()
func (_Mynft *MynftSession) Bid(_auctionId *big.Int, _bidAmount *big.Int, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Mynft.Contract.Bid(&_Mynft.TransactOpts, _auctionId, _bidAmount, _tokenAddress)
}

// Bid is a paid mutator transaction binding the contract method 0x742f0a90.
//
// Solidity: function bid(uint256 _auctionId, uint256 _bidAmount, address _tokenAddress) payable returns()
func (_Mynft *MynftTransactorSession) Bid(_auctionId *big.Int, _bidAmount *big.Int, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Mynft.Contract.Bid(&_Mynft.TransactOpts, _auctionId, _bidAmount, _tokenAddress)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x819c9e4c.
//
// Solidity: function createAuction(uint256 startPrice, uint256 duringTime, uint256 tokenId, address nftContractAddress) returns()
func (_Mynft *MynftTransactor) CreateAuction(opts *bind.TransactOpts, startPrice *big.Int, duringTime *big.Int, tokenId *big.Int, nftContractAddress common.Address) (*types.Transaction, error) {
	return _Mynft.contract.Transact(opts, "createAuction", startPrice, duringTime, tokenId, nftContractAddress)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x819c9e4c.
//
// Solidity: function createAuction(uint256 startPrice, uint256 duringTime, uint256 tokenId, address nftContractAddress) returns()
func (_Mynft *MynftSession) CreateAuction(startPrice *big.Int, duringTime *big.Int, tokenId *big.Int, nftContractAddress common.Address) (*types.Transaction, error) {
	return _Mynft.Contract.CreateAuction(&_Mynft.TransactOpts, startPrice, duringTime, tokenId, nftContractAddress)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x819c9e4c.
//
// Solidity: function createAuction(uint256 startPrice, uint256 duringTime, uint256 tokenId, address nftContractAddress) returns()
func (_Mynft *MynftTransactorSession) CreateAuction(startPrice *big.Int, duringTime *big.Int, tokenId *big.Int, nftContractAddress common.Address) (*types.Transaction, error) {
	return _Mynft.Contract.CreateAuction(&_Mynft.TransactOpts, startPrice, duringTime, tokenId, nftContractAddress)
}

// EndAuction is a paid mutator transaction binding the contract method 0xb9a2de3a.
//
// Solidity: function endAuction(uint256 _auctionId) returns()
func (_Mynft *MynftTransactor) EndAuction(opts *bind.TransactOpts, _auctionId *big.Int) (*types.Transaction, error) {
	return _Mynft.contract.Transact(opts, "endAuction", _auctionId)
}

// EndAuction is a paid mutator transaction binding the contract method 0xb9a2de3a.
//
// Solidity: function endAuction(uint256 _auctionId) returns()
func (_Mynft *MynftSession) EndAuction(_auctionId *big.Int) (*types.Transaction, error) {
	return _Mynft.Contract.EndAuction(&_Mynft.TransactOpts, _auctionId)
}

// EndAuction is a paid mutator transaction binding the contract method 0xb9a2de3a.
//
// Solidity: function endAuction(uint256 _auctionId) returns()
func (_Mynft *MynftTransactorSession) EndAuction(_auctionId *big.Int) (*types.Transaction, error) {
	return _Mynft.Contract.EndAuction(&_Mynft.TransactOpts, _auctionId)
}

// GetLatestPrice is a paid mutator transaction binding the contract method 0x16345f18.
//
// Solidity: function getLatestPrice(address _tokenAddress) returns(int256)
func (_Mynft *MynftTransactor) GetLatestPrice(opts *bind.TransactOpts, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Mynft.contract.Transact(opts, "getLatestPrice", _tokenAddress)
}

// GetLatestPrice is a paid mutator transaction binding the contract method 0x16345f18.
//
// Solidity: function getLatestPrice(address _tokenAddress) returns(int256)
func (_Mynft *MynftSession) GetLatestPrice(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Mynft.Contract.GetLatestPrice(&_Mynft.TransactOpts, _tokenAddress)
}

// GetLatestPrice is a paid mutator transaction binding the contract method 0x16345f18.
//
// Solidity: function getLatestPrice(address _tokenAddress) returns(int256)
func (_Mynft *MynftTransactorSession) GetLatestPrice(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Mynft.Contract.GetLatestPrice(&_Mynft.TransactOpts, _tokenAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Mynft *MynftTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mynft.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Mynft *MynftSession) Initialize() (*types.Transaction, error) {
	return _Mynft.Contract.Initialize(&_Mynft.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Mynft *MynftTransactorSession) Initialize() (*types.Transaction, error) {
	return _Mynft.Contract.Initialize(&_Mynft.TransactOpts)
}

// SetTokenPriceFeed is a paid mutator transaction binding the contract method 0x674417ae.
//
// Solidity: function setTokenPriceFeed(address _tokenAddress, address _priceFeed) returns()
func (_Mynft *MynftTransactor) SetTokenPriceFeed(opts *bind.TransactOpts, _tokenAddress common.Address, _priceFeed common.Address) (*types.Transaction, error) {
	return _Mynft.contract.Transact(opts, "setTokenPriceFeed", _tokenAddress, _priceFeed)
}

// SetTokenPriceFeed is a paid mutator transaction binding the contract method 0x674417ae.
//
// Solidity: function setTokenPriceFeed(address _tokenAddress, address _priceFeed) returns()
func (_Mynft *MynftSession) SetTokenPriceFeed(_tokenAddress common.Address, _priceFeed common.Address) (*types.Transaction, error) {
	return _Mynft.Contract.SetTokenPriceFeed(&_Mynft.TransactOpts, _tokenAddress, _priceFeed)
}

// SetTokenPriceFeed is a paid mutator transaction binding the contract method 0x674417ae.
//
// Solidity: function setTokenPriceFeed(address _tokenAddress, address _priceFeed) returns()
func (_Mynft *MynftTransactorSession) SetTokenPriceFeed(_tokenAddress common.Address, _priceFeed common.Address) (*types.Transaction, error) {
	return _Mynft.Contract.SetTokenPriceFeed(&_Mynft.TransactOpts, _tokenAddress, _priceFeed)
}

// MynftInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Mynft contract.
type MynftInitializedIterator struct {
	Event *MynftInitialized // Event containing the contract specifics and raw log

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
func (it *MynftInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MynftInitialized)
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
		it.Event = new(MynftInitialized)
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
func (it *MynftInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MynftInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MynftInitialized represents a Initialized event raised by the Mynft contract.
type MynftInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Mynft *MynftFilterer) FilterInitialized(opts *bind.FilterOpts) (*MynftInitializedIterator, error) {

	logs, sub, err := _Mynft.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MynftInitializedIterator{contract: _Mynft.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Mynft *MynftFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MynftInitialized) (event.Subscription, error) {

	logs, sub, err := _Mynft.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MynftInitialized)
				if err := _Mynft.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Mynft *MynftFilterer) ParseInitialized(log types.Log) (*MynftInitialized, error) {
	event := new(MynftInitialized)
	if err := _Mynft.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
