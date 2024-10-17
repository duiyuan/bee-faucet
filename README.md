# Faucet

A faucet tool supports distributing test tokens like Sepolia-ETH and Solana (Devnet).

# Supported networks:

Chain:

- Ethereum Sepolia
- Base Sepolia
- Arbitrum Sepolia

Token:
ETH

# Download & Build

Build it by yourself from source code or directly download a pre-built release version from [GitHub](https://github.com/duiyuan/faucet/releases).

```
$ // Build it from source code
$ git clone <github_repo>
$ cd ...
$ go build -o beefaucet
```

# How to use

```
$ beefaucet -h // or go run main.go -h

Usage:
  beefaucet [flags]
  beefaucet [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  list        List all blockchains you can faucet.

Flags:
  -c, --chain string      which chain you want get Sepolia ETH (default "ethereum")
  -e, --endpoint string   Endpoint for blockchain nodes (default "http://139.196.213.90:8595")
  -h, --help              help for beefaucet
  -w, --wallet string     Which wallet/address faucet for (default "0x7F92031F63e01A6ADA475F6E8A637Ce752f8d7D2")

```

### Do Faucet

```
$ beefaucet faucet -w 0x17c3Da8F476a185f707aEA0f1C2A64F04142EbD7 // or go run main.go
Start to deposit Sepolia ETH to 0x17c3Da8F476a185f707aEA0f1C2A64F04142EbD7

✅ ethereum   faucet 0.0001 ETH! (0x935e8291d39f77a468e877c158caee69f8cd77daaaaeee211da710955dd4ce44)
✅ base       faucet 0.0001 ETH! (0x412f8319a402af33927b11469c50a6db17d98c427f33e1243f78e3b9932e05ca)
✅ arbitrum   faucet 0.0001 ETH! (0xd686882b76db9cad1ede9ccf094f2692e851adb94a28c2de887ff0026a3bf31c)

Task completed, succeed: 3, failed: 0
```

### Supported Chain

```
$ beefaucet list // or go run main.go list
Polygon          ID:1442       (supported)
B3               ID:1993       (supported)
Xpla             ID:47         (supported)
Optimism         ID:11155420   (supported)
BNB              ID:97         (supported)
Creator          ID:4654       (supported)
Ethereum Sepolia ID:11155111   (supported)
Base             ID:11155420   (supported)
Arbitrum         ID:421614     (supported)
```
