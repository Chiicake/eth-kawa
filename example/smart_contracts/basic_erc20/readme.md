solc
```bash
solcjs --bin --abi ./basic_erc20.sol -o ./build
```

abigen
```bash
abigen --bin=./build/basic_erc20_sol_BasicERC20.bin --abi=./build/basic_erc20_sol_BasicERC20.abi --pkg=basic_erc20 --out=basic_erc20.go
```