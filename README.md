# Faucet

A faucet tool supports distributing test tokens like Sepolia-ETH and Solana (Devnet).

# Supported Network

Chain:

- Ethereum Sepolia
- Base Sepolia
- Arbitrum Sepolia
- Optimism Sepolia

Token:

- ETH
- 0xbee USDT

# Download

[Here](https://github.com/duiyuan/faucet/releases).

# Build from source code

Or you can build it from source code, if you update code locally.

```
$ git clone <github_repo>
$ cd ...
$ go build -o beetools
```

# How to use

```
$ beetools -h // or go run main.go -h

Usage:
  beetools [flags]
  beetools [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  list        List all blockchains you can faucet.

Flags:
  -c, --chain string      which chain you want get Sepolia ETH (default "ethereum")
  -e, --endpoint string   Endpoint for blockchain nodes (default "http://139.196.213.90:8595")
  -h, --help              help for beetools
  -w, --wallet string     Which wallet/address faucet for (default "0x7F92031F63e01A6ADA475F6E8A637Ce752f8d7D2")

```

### Faucet

```
$ beetools faucet -w 0x17c3Da8F476a185f707aEA0f1C2A64F04142EbD7 --chain optimism --token usdt // Retrive one token

```

----------|-------|--------|--------|------------|-------------------------------------------------------------------------------------------------------------|
| CHAIN | TOKEN | AMOUNT | STATUS | STATUSTEXT | LINK |
|----------|-------|--------|--------|------------|-------------------------------------------------------------------------------------------------------------|
| optimism | usdt | 10 | OK | OK | https://sepolia-optimism.etherscan.io/tx/0x22ec7511050c4d5c1b218ac1fe8724ee187cb611a2ce271059e18f7f2a3e8878 |
|----------|-------|--------|--------|------------|-------------------------------------------------------------------------------------------------------------|

```

### Rate Limit

You can request the faucet once every 6 hours; violating this will result in rate limiting for your requests.

```

$ beetools faucet -w 0x17c3Da8F476a185f707aEA0f1C2A64F04142EbD7 --chain base --token usdt
$
|----------|-------|--------|--------|--------------------------------|------|
| CHAIN | TOKEN | AMOUNT | STATUS | STATUSTEXT | LINK |
|----------|-------|--------|--------|--------------------------------|------|
| ethereum | usdt | | X | rate limit, please retry after | |
| | | | | 05:59:18 | |
|----------|-------|--------|--------|--------------------------------|------|

```

### Supported Chain

```

$ beetools list // or go run main.go list
Arbitrum ID:421614 (Supported)
B3 ID:1993 (Not supported)
BNB ID:97 (Not supported)
Base ID:11155420 (Supported)
Creator ID:4654 (Not supported)
Ethereum ID:11155111 (Supported)
Optimism ID:11155420 (Supported)
Polygon ID:1442 (Not supported)
Xpla ID:47 (Not supported)

```

```
