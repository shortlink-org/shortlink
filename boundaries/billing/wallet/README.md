# Wallet

This service is responsible for managing the user's wallet.

### Requirements

- **Solidity** for smart contracts
- **Go** for the backend

### Installation

- Install [Ethereum tools](https://geth.ethereum.org/docs/install-and-build/installing-geth)

```shell
$> docker run --rm -v $(pwd):/root ethereum/solc:0.8.17-alpine --abi /root/contracts/Wallet.sol -o /root/build
$> docker run --rm -v $(pwd):/root ethereum/solc:0.8.17-alpine --abi /root/contracts/Store.sol -o /root/build

$> abigen --abi=build/Store.abi --pkg=store --out=store/Store.go
$> abigen --abi=build/Wallet.abi --pkg=wallet --out=wallet/Wallet.go
```
