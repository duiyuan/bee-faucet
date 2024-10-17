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

✅ optimism   faucet 0.0001 ETH! (0x7a0d2bed16c9bce297d009ae6e6b812718bd034274e35a582532b2e4196e831b)
✅ ethereum   faucet 0.0001 ETH! (0x0892599ef49b4fc5071f1c0a6bd286df2026f5ecf70a059af3071d71b912046f)
✅ arbitrum   faucet 0.0001 ETH! (0xb37afd1a14edfe6f7c67e9b2e0a5524413c4eeb57f57d810491bdd1b1a818445)
✅ base       faucet 0.0001 ETH! (0x9b5547273ac751f31439f000732a91aa1b3310dfcdb17b83346fef77d88acacb)

Task completed, succeed: 4, failed: 0
```

### Supported Chain

```
$ beefaucet list // or go run main.go list
Arbitrum ID:421614     (Supported)
B3       ID:1993       (Not supported)
BNB      ID:97         (Not supported)
Base     ID:11155420   (Supported)
Creator  ID:4654       (Not supported)
Ethereum ID:11155111   (Supported)
Optimism ID:11155420   (Supported)
Polygon  ID:1442       (Not supported)
Xpla     ID:47         (Not supported)
```
