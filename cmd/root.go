package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var wallet string
var chain string
var token string
var endpoint string

var rootCmd = &cobra.Command{
	Use:   "beefaucet",
	Short: "Faucet ETH Token for Ethereum or L2 base on Ethereum",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&wallet, "wallet", "w", "0x7F92031F63e01A6ADA475F6E8A637Ce752f8d7D2", "Which wallet/address faucet for")
	rootCmd.PersistentFlags().StringVarP(&chain, "chain", "c", "ethereum", "which chain you want get Sepolia ETH")
	rootCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "which token you want get it's will be ETH, USDT")
	rootCmd.PersistentFlags().StringVarP(&endpoint, "endpoint", "e", "http://139.196.213.90:8595", "Endpoint for blockchain nodes")
}

func Excute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
