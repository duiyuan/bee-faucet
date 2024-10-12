package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var wallet string
var chain string

var rootCmd = &cobra.Command{
	Use:   "faucet",
	Short: "Faucet ETH Token for Ethereum or L2 base on Ethereum",
	Run: func(cmd *cobra.Command, args []string) {
		faucet(wallet)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&wallet, "wallet", "w", "0x7F92031F63e01A6ADA475F6E8A637Ce752f8d7D2", "Which wallet/address faucet for")
	rootCmd.PersistentFlags().StringVarP(&chain, "chain", "c", "ethereum", "which chain you want get Sepolia ETH")
}

func Excute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
