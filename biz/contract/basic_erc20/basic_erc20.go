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
	Bin: "0x60806040526012600260006101000a81548160ff021916908360ff1602179055503480156200002d57600080fd5b50604051620017e5380380620017e58339818101604052810190620000539190620003fa565b82600090805190602001906200006b92919062000172565b5081600190805190602001906200008492919062000172565b50600260009054906101000a900460ff1660ff16600a620000a6919062000617565b81620000b3919062000668565b600381905550600354600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055503373ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef600354604051620001619190620006da565b60405180910390a35050506200075c565b828054620001809062000726565b90600052602060002090601f016020900481019282620001a45760008555620001f0565b82601f10620001bf57805160ff1916838001178555620001f0565b82800160010185558215620001f0579182015b82811115620001ef578251825591602001919060010190620001d2565b5b509050620001ff919062000203565b5090565b5b808211156200021e57600081600090555060010162000204565b5090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6200028b8262000240565b810181811067ffffffffffffffff82111715620002ad57620002ac62000251565b5b80604052505050565b6000620002c262000222565b9050620002d0828262000280565b919050565b600067ffffffffffffffff821115620002f357620002f262000251565b5b620002fe8262000240565b9050602081019050919050565b60005b838110156200032b5780820151818401526020810190506200030e565b838111156200033b576000848401525b50505050565b6000620003586200035284620002d5565b620002b6565b9050828152602081018484840111156200037757620003766200023b565b5b620003848482856200030b565b509392505050565b600082601f830112620003a457620003a362000236565b5b8151620003b684826020860162000341565b91505092915050565b6000819050919050565b620003d481620003bf565b8114620003e057600080fd5b50565b600081519050620003f481620003c9565b92915050565b6000806000606084860312156200041657620004156200022c565b5b600084015167ffffffffffffffff81111562000437576200043662000231565b5b62000445868287016200038c565b935050602084015167ffffffffffffffff81111562000469576200046862000231565b5b62000477868287016200038c565b92505060406200048a86828701620003e3565b9150509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008160011c9050919050565b6000808291508390505b60018511156200052257808604811115620004fa57620004f962000494565b5b60018516156200050a5780820291505b80810290506200051a85620004c3565b9450620004da565b94509492505050565b6000826200053d576001905062000610565b816200054d576000905062000610565b81600181146200056657600281146200057157620005a7565b600191505062000610565b60ff84111562000586576200058562000494565b5b8360020a915084821115620005a0576200059f62000494565b5b5062000610565b5060208310610133831016604e8410600b8410161715620005e15782820a905083811115620005db57620005da62000494565b5b62000610565b620005f08484846001620004d0565b925090508184048111156200060a576200060962000494565b5b81810290505b9392505050565b60006200062482620003bf565b91506200063183620003bf565b9250620006607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84846200052b565b905092915050565b60006200067582620003bf565b91506200068283620003bf565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615620006be57620006bd62000494565b5b828202905092915050565b620006d481620003bf565b82525050565b6000602082019050620006f16000830184620006c9565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200073f57607f821691505b60208210811415620007565762000755620006f7565b5b50919050565b611079806200076c6000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c8063313ce56711610066578063313ce5671461013457806370a082311461015257806395d89b4114610182578063a9059cbb146101a0578063dd62ed3e146101d057610093565b806306fdde0314610098578063095ea7b3146100b657806318160ddd146100e657806323b872dd14610104575b600080fd5b6100a0610200565b6040516100ad9190610adb565b60405180910390f35b6100d060048036038101906100cb9190610b96565b61028e565b6040516100dd9190610bf1565b60405180910390f35b6100ee6103ef565b6040516100fb9190610c1b565b60405180910390f35b61011e60048036038101906101199190610c36565b6103f5565b60405161012b9190610bf1565b60405180910390f35b61013c610756565b6040516101499190610ca5565b60405180910390f35b61016c60048036038101906101679190610cc0565b610769565b6040516101799190610c1b565b60405180910390f35b61018a610781565b6040516101979190610adb565b60405180910390f35b6101ba60048036038101906101b59190610b96565b61080f565b6040516101c79190610bf1565b60405180910390f35b6101ea60048036038101906101e59190610ced565b610a1d565b6040516101f79190610c1b565b60405180910390f35b6000805461020d90610d5c565b80601f016020809104026020016040519081016040528092919081815260200182805461023990610d5c565b80156102865780601f1061025b57610100808354040283529160200191610286565b820191906000526020600020905b81548152906001019060200180831161026957829003601f168201915b505050505081565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614156102ff576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102f690610e00565b60405180910390fd5b81600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516103dd9190610c1b565b60405180910390a36001905092915050565b60035481565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610466576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161045d90610e92565b60405180910390fd5b81600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410156104e8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104df90610efe565b60405180910390fd5b81600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410156105a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161059e90610f6a565b60405180910390fd5b81600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546105f69190610fb9565b9250508190555081600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461064c9190610fed565b9250508190555081600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546106df9190610fb9565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516107439190610c1b565b60405180910390a3600190509392505050565b600260009054906101000a900460ff1681565b60046020528060005260406000206000915090505481565b6001805461078e90610d5c565b80601f01602080910402602001604051908101604052809291908181526020018280546107ba90610d5c565b80156108075780601f106107dc57610100808354040283529160200191610807565b820191906000526020600020905b8154815290600101906020018083116107ea57829003601f168201915b505050505081565b60008073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610880576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161087790610e92565b60405180910390fd5b81600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015610902576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108f990610efe565b60405180910390fd5b81600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546109519190610fb9565b9250508190555081600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546109a79190610fed565b925050819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610a0b9190610c1b565b60405180910390a36001905092915050565b6005602052816000526040600020602052806000526040600020600091509150505481565b600081519050919050565b600082825260208201905092915050565b60005b83811015610a7c578082015181840152602081019050610a61565b83811115610a8b576000848401525b50505050565b6000601f19601f8301169050919050565b6000610aad82610a42565b610ab78185610a4d565b9350610ac7818560208601610a5e565b610ad081610a91565b840191505092915050565b60006020820190508181036000830152610af58184610aa2565b905092915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610b2d82610b02565b9050919050565b610b3d81610b22565b8114610b4857600080fd5b50565b600081359050610b5a81610b34565b92915050565b6000819050919050565b610b7381610b60565b8114610b7e57600080fd5b50565b600081359050610b9081610b6a565b92915050565b60008060408385031215610bad57610bac610afd565b5b6000610bbb85828601610b4b565b9250506020610bcc85828601610b81565b9150509250929050565b60008115159050919050565b610beb81610bd6565b82525050565b6000602082019050610c066000830184610be2565b92915050565b610c1581610b60565b82525050565b6000602082019050610c306000830184610c0c565b92915050565b600080600060608486031215610c4f57610c4e610afd565b5b6000610c5d86828701610b4b565b9350506020610c6e86828701610b4b565b9250506040610c7f86828701610b81565b9150509250925092565b600060ff82169050919050565b610c9f81610c89565b82525050565b6000602082019050610cba6000830184610c96565b92915050565b600060208284031215610cd657610cd5610afd565b5b6000610ce484828501610b4b565b91505092915050565b60008060408385031215610d0457610d03610afd565b5b6000610d1285828601610b4b565b9250506020610d2385828601610b4b565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610d7457607f821691505b60208210811415610d8857610d87610d2d565b5b50919050565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b6000610dea602283610a4d565b9150610df582610d8e565b604082019050919050565b60006020820190508181036000830152610e1981610ddd565b9050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b6000610e7c602383610a4d565b9150610e8782610e20565b604082019050919050565b60006020820190508181036000830152610eab81610e6f565b9050919050565b7f45524332303a20696e73756666696369656e742062616c616e63650000000000600082015250565b6000610ee8601b83610a4d565b9150610ef382610eb2565b602082019050919050565b60006020820190508181036000830152610f1781610edb565b9050919050565b7f45524332303a20616c6c6f77616e636520657863656564656400000000000000600082015250565b6000610f54601983610a4d565b9150610f5f82610f1e565b602082019050919050565b60006020820190508181036000830152610f8381610f47565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610fc482610b60565b9150610fcf83610b60565b925082821015610fe257610fe1610f8a565b5b828203905092915050565b6000610ff882610b60565b915061100383610b60565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561103857611037610f8a565b5b82820190509291505056fea26469706673582212206c6da1ee3a5ac98d2c815fc38f17291aefe5aa56c6b476b330c7ea732e3b098b64736f6c63430008080033",
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
