This is a simple project to learn the eth development.

As the pre-requisite, install ganache and run it, the ganache will listen on http://localhost:8545
```bash
npm install -g ganache
ganache
```
Install a Solidity compiler
```bash
npm install -g solc@0.8.8
```

Install abigen tool
```bash
go install github.com/ethereum/go-ethereum/cmd/abigen@latest
```

部署合约目前只能使用 Remix 编译后部署
但是需要使用 abigen 生成go合约代码
具体流程为
1. 启动ganache
2. 使用Remix编译合约
3. 替换infra中的合约地址来加载合约