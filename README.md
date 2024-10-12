# Faucet

A faucet tool that supports sending free Sepolia ETH to specified address.

# How To Download

Build it by yourself from the source code or directly download a pre-built release version from [GitHub](https://github.com/duiyuan/faucet/releases).

```
$ // Build it from source code
$ git clone <github_repo>
$ cd ...
$ go build -o beefaucet
```

# How to use

```
$ ./beefaucet -h // or go run main.go -h

Faucet ETH Token for Ethereum or L2 base on Ethereum

Usage:
  faucet [flags]
  faucet [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  list        List all blockchains you can faucet.

Flags:
  -c, --chain string    which chain you want get Sepolia ETH (default "ethereum")
  -h, --help            help for faucet
  -w, --wallet string   Which wallet/address faucet for (default "0x7F92031F63e01A6ADA475F6E8A637Ce752f8d7D2")

```

### Do Faucet

```
$ ./beefaucet // or go run main.go
Start to faucet ETH for wallet: 0x7F92031F63e01A6ADA475F6E8A637Ce752f8d7D2

  Xpla             done, facuet 0.1 ETH
  Base             done, facuet 0.1 ETH
  Creator          done, facuet 0.1 ETH
  Ethereum Sepolia done, facuet 0.1 ETH
  B3               done, facuet 0.1 ETH
  Optimism         done, facuet 0.1 ETH
  Polygon          done, facuet 0.1 ETH
  Arbitrum         done, facuet 0.1 ETH
  BNB              done, facuet 0.1 ETH

End faucet, succeed: 9, failed: 0
```

### Supported Chain

```
$ ./beefaucet list // or go run main.go list
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
