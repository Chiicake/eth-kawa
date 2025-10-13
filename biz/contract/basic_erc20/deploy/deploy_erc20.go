package main

import (
	"context"
	"crypto/ecdsa"
	erc20 "eth-kawa/example/smart_contracts/basic_erc20" // 替换为你的ERC20 ABI包路径
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// 连接到以太坊节点
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatalf("无法连接到节点: %v", err)
	}

	// 加载部署者私钥
	privateKey, err := crypto.HexToECDSA("869a589a1b188e9d1f84c74ad0dd6eaa6eb2f59460dd6c03bff0648736cd265e")
	if err != nil {
		log.Fatalf("无效的私钥: %v", err)
	}

	// 从私钥获取公钥和地址
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("公钥类型转换失败")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 获取当前地址的nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("获取nonce失败: %v", err)
	}

	// 获取建议的gas价格
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("获取gas价格失败: %v", err)
	}

	// 创建交易签名器
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // 部署合约不需要发送ETH
	auth.GasLimit = uint64(600000) // ERC20部署需要更多gas，适当调高
	auth.GasPrice = gasPrice

	// ERC20代币参数 - 根据需要修改
	tokenName := "MyBasicToken"          // 代币名称
	tokenSymbol := "MBT"                 // 代币符号
	initialSupply := big.NewInt(1000000) // 初始供应量（会自动乘以10^18）

	// 部署合约
	address, tx, instance, err := erc20.DeployBasicErc20(
		auth,
		client,
		tokenName,
		tokenSymbol,
		initialSupply,
	)
	if err != nil {
		log.Fatalf("部署合约失败: %v", err)
	}

	// 输出部署结果
	fmt.Printf("合约部署地址: %s\n", address.Hex())
	fmt.Printf("交易哈希: %s\n", tx.Hash().Hex())

	// 保留实例引用，可用于后续合约交互
	_ = instance
}
