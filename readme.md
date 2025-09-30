This is a simple project to learn the eth development.

As the pre-requisite, install ganache and run it, the ganache will listen on http://localhost:8545
```bash
npm install -g ganache-cli
ganache-cli
```
Install a Solidity compiler
```bash
npm install --global solc
```

Install abigen tool
```bash
go install github.com/ethereum/go-ethereum/cmd/abigen@latest
```
abigen --bin=./build/store_sol_Store.bin --abi=./build/store_sol_Store.abi --pkg=store --out=store.go