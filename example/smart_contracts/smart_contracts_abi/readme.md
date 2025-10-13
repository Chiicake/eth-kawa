solc
```bash
solcjs -optimize ./store.sol --bin --abi -o ./build
```

abigen
```bash
abigen --bin=./build/store_sol_Store.bin --abi=./build/store_sol_Store.abi --pkg=store --out=store.go
```