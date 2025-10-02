solc
```bash
solcjs --bin --abi ./store.sol -o ./build
```

abigen
```bash
abigen --bin=./build/store_sol_Store.bin --abi=./build/store_sol_Store.abi --pkg=store --out=store.go
```