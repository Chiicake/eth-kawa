// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package basic_erc20

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

// BasicErc20MetaData contains all meta data concerning the BasicErc20 contract.
var BasicErc20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_initialSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052601260025f6101000a81548160ff021916908360ff16021790555034801561002a575f5ffd5b5060405161180e38038061180e833981810160405281019061004c91906102cd565b825f908161005a919061055c565b50816001908161006a919061055c565b5060025f9054906101000a900460ff1660ff16600a6100899190610787565b8161009491906107d1565b60038190555060035460045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055503373ffffffffffffffffffffffffffffffffffffffff165f73ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60035460405161013d9190610821565b60405180910390a350505061083a565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6101ac82610166565b810181811067ffffffffffffffff821117156101cb576101ca610176565b5b80604052505050565b5f6101dd61014d565b90506101e982826101a3565b919050565b5f67ffffffffffffffff82111561020857610207610176565b5b61021182610166565b9050602081019050919050565b8281835e5f83830152505050565b5f61023e610239846101ee565b6101d4565b90508281526020810184848401111561025a57610259610162565b5b61026584828561021e565b509392505050565b5f82601f8301126102815761028061015e565b5b815161029184826020860161022c565b91505092915050565b5f819050919050565b6102ac8161029a565b81146102b6575f5ffd5b50565b5f815190506102c7816102a3565b92915050565b5f5f5f606084860312156102e4576102e3610156565b5b5f84015167ffffffffffffffff8111156103015761030061015a565b5b61030d8682870161026d565b935050602084015167ffffffffffffffff81111561032e5761032d61015a565b5b61033a8682870161026d565b925050604061034b868287016102b9565b9150509250925092565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806103a357607f821691505b6020821081036103b6576103b561035f565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026104187fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826103dd565b61042286836103dd565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61045d6104586104538461029a565b61043a565b61029a565b9050919050565b5f819050919050565b61047683610443565b61048a61048282610464565b8484546103e9565b825550505050565b5f5f905090565b6104a1610492565b6104ac81848461046d565b505050565b5b818110156104cf576104c45f82610499565b6001810190506104b2565b5050565b601f821115610514576104e5816103bc565b6104ee846103ce565b810160208510156104fd578190505b610511610509856103ce565b8301826104b1565b50505b505050565b5f82821c905092915050565b5f6105345f1984600802610519565b1980831691505092915050565b5f61054c8383610525565b9150826002028217905092915050565b61056582610355565b67ffffffffffffffff81111561057e5761057d610176565b5b610588825461038c565b6105938282856104d3565b5f60209050601f8311600181146105c4575f84156105b2578287015190505b6105bc8582610541565b865550610623565b601f1984166105d2866103bc565b5f5b828110156105f9578489015182556001820191506020850194506020810190506105d4565b868310156106165784890151610612601f891682610525565b8355505b6001600288020188555050505b505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f8160011c9050919050565b5f5f8291508390505b60018511156106ad578086048111156106895761068861062b565b5b60018516156106985780820291505b80810290506106a685610658565b945061066d565b94509492505050565b5f826106c55760019050610780565b816106d2575f9050610780565b81600181146106e857600281146106f257610721565b6001915050610780565b60ff8411156107045761070361062b565b5b8360020a91508482111561071b5761071a61062b565b5b50610780565b5060208310610133831016604e8410600b84101617156107565782820a9050838111156107515761075061062b565b5b610780565b6107638484846001610664565b9250905081840481111561077a5761077961062b565b5b81810290505b9392505050565b5f6107918261029a565b915061079c8361029a565b92506107c97fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84846106b6565b905092915050565b5f6107db8261029a565b91506107e68361029a565b92508282026107f48161029a565b9150828204841483151761080b5761080a61062b565b5b5092915050565b61081b8161029a565b82525050565b5f6020820190506108345f830184610812565b92915050565b610fc7806108475f395ff3fe608060405234801561000f575f5ffd5b5060043610610091575f3560e01c8063313ce56711610064578063313ce5671461013157806370a082311461014f57806395d89b411461017f578063a9059cbb1461019d578063dd62ed3e146101cd57610091565b806306fdde0314610095578063095ea7b3146100b357806318160ddd146100e357806323b872dd14610101575b5f5ffd5b61009d6101fd565b6040516100aa9190610a7e565b60405180910390f35b6100cd60048036038101906100c89190610b2f565b610288565b6040516100da9190610b87565b60405180910390f35b6100eb6103e3565b6040516100f89190610baf565b60405180910390f35b61011b60048036038101906101169190610bc8565b6103e9565b6040516101289190610b87565b60405180910390f35b610139610737565b6040516101469190610c33565b60405180910390f35b61016960048036038101906101649190610c4c565b610749565b6040516101769190610baf565b60405180910390f35b61018761075e565b6040516101949190610a7e565b60405180910390f35b6101b760048036038101906101b29190610b2f565b6107ea565b6040516101c49190610b87565b60405180910390f35b6101e760048036038101906101e29190610c77565b6109ee565b6040516101f49190610baf565b60405180910390f35b5f805461020990610ce2565b80601f016020809104026020016040519081016040528092919081815260200182805461023590610ce2565b80156102805780601f1061025757610100808354040283529160200191610280565b820191905f5260205f20905b81548152906001019060200180831161026357829003601f168201915b505050505081565b5f5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036102f7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102ee90610d82565b60405180910390fd5b8160055f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516103d19190610baf565b60405180910390a36001905092915050565b60035481565b5f5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610458576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161044f90610e10565b60405180910390fd5b8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410156104d8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104cf90610e78565b60405180910390fd5b8160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015610593576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161058a90610ee0565b60405180910390fd5b8160045f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546105df9190610f2b565b925050819055508160045f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546106329190610f5e565b925050819055508160055f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546106c09190610f2b565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516107249190610baf565b60405180910390a3600190509392505050565b60025f9054906101000a900460ff1681565b6004602052805f5260405f205f915090505481565b6001805461076b90610ce2565b80601f016020809104026020016040519081016040528092919081815260200182805461079790610ce2565b80156107e25780601f106107b9576101008083540402835291602001916107e2565b820191905f5260205f20905b8154815290600101906020018083116107c557829003601f168201915b505050505081565b5f5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610859576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161085090610e10565b60405180910390fd5b8160045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205410156108d9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108d090610e78565b60405180910390fd5b8160045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546109259190610f2b565b925050819055508160045f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546109789190610f5e565b925050819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516109dc9190610baf565b60405180910390a36001905092915050565b6005602052815f5260405f20602052805f5260405f205f91509150505481565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610a5082610a0e565b610a5a8185610a18565b9350610a6a818560208601610a28565b610a7381610a36565b840191505092915050565b5f6020820190508181035f830152610a968184610a46565b905092915050565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610acb82610aa2565b9050919050565b610adb81610ac1565b8114610ae5575f5ffd5b50565b5f81359050610af681610ad2565b92915050565b5f819050919050565b610b0e81610afc565b8114610b18575f5ffd5b50565b5f81359050610b2981610b05565b92915050565b5f5f60408385031215610b4557610b44610a9e565b5b5f610b5285828601610ae8565b9250506020610b6385828601610b1b565b9150509250929050565b5f8115159050919050565b610b8181610b6d565b82525050565b5f602082019050610b9a5f830184610b78565b92915050565b610ba981610afc565b82525050565b5f602082019050610bc25f830184610ba0565b92915050565b5f5f5f60608486031215610bdf57610bde610a9e565b5b5f610bec86828701610ae8565b9350506020610bfd86828701610ae8565b9250506040610c0e86828701610b1b565b9150509250925092565b5f60ff82169050919050565b610c2d81610c18565b82525050565b5f602082019050610c465f830184610c24565b92915050565b5f60208284031215610c6157610c60610a9e565b5b5f610c6e84828501610ae8565b91505092915050565b5f5f60408385031215610c8d57610c8c610a9e565b5b5f610c9a85828601610ae8565b9250506020610cab85828601610ae8565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610cf957607f821691505b602082108103610d0c57610d0b610cb5565b5b50919050565b7f45524332303a20617070726f766520746f20746865207a65726f2061646472655f8201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b5f610d6c602283610a18565b9150610d7782610d12565b604082019050919050565b5f6020820190508181035f830152610d9981610d60565b9050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f20616464725f8201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b5f610dfa602383610a18565b9150610e0582610da0565b604082019050919050565b5f6020820190508181035f830152610e2781610dee565b9050919050565b7f45524332303a20696e73756666696369656e742062616c616e636500000000005f82015250565b5f610e62601b83610a18565b9150610e6d82610e2e565b602082019050919050565b5f6020820190508181035f830152610e8f81610e56565b9050919050565b7f45524332303a20616c6c6f77616e6365206578636565646564000000000000005f82015250565b5f610eca601983610a18565b9150610ed582610e96565b602082019050919050565b5f6020820190508181035f830152610ef781610ebe565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610f3582610afc565b9150610f4083610afc565b9250828203905081811115610f5857610f57610efe565b5b92915050565b5f610f6882610afc565b9150610f7383610afc565b9250828201905080821115610f8b57610f8a610efe565b5b9291505056fea26469706673582212207de8c40267c0bd45ddc337fe0b5286ba0ed79f70e07a6ed004b00425a9fdb3bf64736f6c634300081e0033",
}

// BasicErc20ABI is the input ABI used to generate the binding from.
// Deprecated: Use BasicErc20MetaData.ABI instead.
var BasicErc20ABI = BasicErc20MetaData.ABI

// BasicErc20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BasicErc20MetaData.Bin instead.
var BasicErc20Bin = BasicErc20MetaData.Bin

// DeployBasicErc20 deploys a new Ethereum contract, binding an instance of BasicErc20 to it.
func DeployBasicErc20(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string, _initialSupply *big.Int) (common.Address, *types.Transaction, *BasicErc20, error) {
	parsed, err := BasicErc20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BasicErc20Bin), backend, _name, _symbol, _initialSupply)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BasicErc20{BasicErc20Caller: BasicErc20Caller{contract: contract}, BasicErc20Transactor: BasicErc20Transactor{contract: contract}, BasicErc20Filterer: BasicErc20Filterer{contract: contract}}, nil
}

// BasicErc20 is an auto generated Go binding around an Ethereum contract.
type BasicErc20 struct {
	BasicErc20Caller     // Read-only binding to the contract
	BasicErc20Transactor // Write-only binding to the contract
	BasicErc20Filterer   // Log filterer for contract events
}

// BasicErc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type BasicErc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicErc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type BasicErc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicErc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasicErc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicErc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasicErc20Session struct {
	Contract     *BasicErc20       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasicErc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasicErc20CallerSession struct {
	Contract *BasicErc20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BasicErc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasicErc20TransactorSession struct {
	Contract     *BasicErc20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BasicErc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type BasicErc20Raw struct {
	Contract *BasicErc20 // Generic contract binding to access the raw methods on
}

// BasicErc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasicErc20CallerRaw struct {
	Contract *BasicErc20Caller // Generic read-only contract binding to access the raw methods on
}

// BasicErc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasicErc20TransactorRaw struct {
	Contract *BasicErc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBasicErc20 creates a new instance of BasicErc20, bound to a specific deployed contract.
func NewBasicErc20(address common.Address, backend bind.ContractBackend) (*BasicErc20, error) {
	contract, err := bindBasicErc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasicErc20{BasicErc20Caller: BasicErc20Caller{contract: contract}, BasicErc20Transactor: BasicErc20Transactor{contract: contract}, BasicErc20Filterer: BasicErc20Filterer{contract: contract}}, nil
}

// NewBasicErc20Caller creates a new read-only instance of BasicErc20, bound to a specific deployed contract.
func NewBasicErc20Caller(address common.Address, caller bind.ContractCaller) (*BasicErc20Caller, error) {
	contract, err := bindBasicErc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasicErc20Caller{contract: contract}, nil
}

// NewBasicErc20Transactor creates a new write-only instance of BasicErc20, bound to a specific deployed contract.
func NewBasicErc20Transactor(address common.Address, transactor bind.ContractTransactor) (*BasicErc20Transactor, error) {
	contract, err := bindBasicErc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasicErc20Transactor{contract: contract}, nil
}

// NewBasicErc20Filterer creates a new log filterer instance of BasicErc20, bound to a specific deployed contract.
func NewBasicErc20Filterer(address common.Address, filterer bind.ContractFilterer) (*BasicErc20Filterer, error) {
	contract, err := bindBasicErc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasicErc20Filterer{contract: contract}, nil
}

// bindBasicErc20 binds a generic wrapper to an already deployed contract.
func bindBasicErc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BasicErc20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicErc20 *BasicErc20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasicErc20.Contract.BasicErc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicErc20 *BasicErc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicErc20.Contract.BasicErc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicErc20 *BasicErc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicErc20.Contract.BasicErc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicErc20 *BasicErc20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BasicErc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicErc20 *BasicErc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicErc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicErc20 *BasicErc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicErc20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_BasicErc20 *BasicErc20Caller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BasicErc20.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_BasicErc20 *BasicErc20Session) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _BasicErc20.Contract.Allowance(&_BasicErc20.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_BasicErc20 *BasicErc20CallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _BasicErc20.Contract.Allowance(&_BasicErc20.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_BasicErc20 *BasicErc20Caller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BasicErc20.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_BasicErc20 *BasicErc20Session) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _BasicErc20.Contract.BalanceOf(&_BasicErc20.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_BasicErc20 *BasicErc20CallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _BasicErc20.Contract.BalanceOf(&_BasicErc20.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BasicErc20 *BasicErc20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BasicErc20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BasicErc20 *BasicErc20Session) Decimals() (uint8, error) {
	return _BasicErc20.Contract.Decimals(&_BasicErc20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BasicErc20 *BasicErc20CallerSession) Decimals() (uint8, error) {
	return _BasicErc20.Contract.Decimals(&_BasicErc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BasicErc20 *BasicErc20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BasicErc20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BasicErc20 *BasicErc20Session) Name() (string, error) {
	return _BasicErc20.Contract.Name(&_BasicErc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BasicErc20 *BasicErc20CallerSession) Name() (string, error) {
	return _BasicErc20.Contract.Name(&_BasicErc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BasicErc20 *BasicErc20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BasicErc20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BasicErc20 *BasicErc20Session) Symbol() (string, error) {
	return _BasicErc20.Contract.Symbol(&_BasicErc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BasicErc20 *BasicErc20CallerSession) Symbol() (string, error) {
	return _BasicErc20.Contract.Symbol(&_BasicErc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BasicErc20 *BasicErc20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BasicErc20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BasicErc20 *BasicErc20Session) TotalSupply() (*big.Int, error) {
	return _BasicErc20.Contract.TotalSupply(&_BasicErc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BasicErc20 *BasicErc20CallerSession) TotalSupply() (*big.Int, error) {
	return _BasicErc20.Contract.TotalSupply(&_BasicErc20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool success)
func (_BasicErc20 *BasicErc20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BasicErc20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool success)
func (_BasicErc20 *BasicErc20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BasicErc20.Contract.Approve(&_BasicErc20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool success)
func (_BasicErc20 *BasicErc20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _BasicErc20.Contract.Approve(&_BasicErc20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool success)
func (_BasicErc20 *BasicErc20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BasicErc20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool success)
func (_BasicErc20 *BasicErc20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BasicErc20.Contract.Transfer(&_BasicErc20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool success)
func (_BasicErc20 *BasicErc20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _BasicErc20.Contract.Transfer(&_BasicErc20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 value) returns(bool success)
func (_BasicErc20 *BasicErc20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, value *big.Int) (*types.Transaction, error) {
	return _BasicErc20.contract.Transact(opts, "transferFrom", sender, recipient, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 value) returns(bool success)
func (_BasicErc20 *BasicErc20Session) TransferFrom(sender common.Address, recipient common.Address, value *big.Int) (*types.Transaction, error) {
	return _BasicErc20.Contract.TransferFrom(&_BasicErc20.TransactOpts, sender, recipient, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 value) returns(bool success)
func (_BasicErc20 *BasicErc20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, value *big.Int) (*types.Transaction, error) {
	return _BasicErc20.Contract.TransferFrom(&_BasicErc20.TransactOpts, sender, recipient, value)
}

// BasicErc20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BasicErc20 contract.
type BasicErc20ApprovalIterator struct {
	Event *BasicErc20Approval // Event containing the contract specifics and raw log

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
func (it *BasicErc20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicErc20Approval)
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
		it.Event = new(BasicErc20Approval)
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
func (it *BasicErc20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicErc20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicErc20Approval represents a Approval event raised by the BasicErc20 contract.
type BasicErc20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BasicErc20 *BasicErc20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BasicErc20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BasicErc20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BasicErc20ApprovalIterator{contract: _BasicErc20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BasicErc20 *BasicErc20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BasicErc20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BasicErc20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicErc20Approval)
				if err := _BasicErc20.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BasicErc20 *BasicErc20Filterer) ParseApproval(log types.Log) (*BasicErc20Approval, error) {
	event := new(BasicErc20Approval)
	if err := _BasicErc20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BasicErc20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BasicErc20 contract.
type BasicErc20TransferIterator struct {
	Event *BasicErc20Transfer // Event containing the contract specifics and raw log

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
func (it *BasicErc20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicErc20Transfer)
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
		it.Event = new(BasicErc20Transfer)
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
func (it *BasicErc20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicErc20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicErc20Transfer represents a Transfer event raised by the BasicErc20 contract.
type BasicErc20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BasicErc20 *BasicErc20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BasicErc20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BasicErc20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BasicErc20TransferIterator{contract: _BasicErc20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BasicErc20 *BasicErc20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BasicErc20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BasicErc20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicErc20Transfer)
				if err := _BasicErc20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BasicErc20 *BasicErc20Filterer) ParseTransfer(log types.Log) (*BasicErc20Transfer, error) {
	event := new(BasicErc20Transfer)
	if err := _BasicErc20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
