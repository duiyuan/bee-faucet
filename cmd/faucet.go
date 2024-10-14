package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"githun.com/duiyuan/faucet/core"
)

var faucetCmd = &cobra.Command{
	Use:   "faucet",
	Short: "Faucet Sepolia ETH",
	Run: func(cmd *cobra.Command, arg []string) {
		Dofaucet(chain, wallet)
	},
}

func Dofaucet(chain string, wallet string) {
	faucet, err := core.ClaimToken(
		chain,
		wallet,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	hash := faucet["hash"]
	fmt.Printf("Hash %s \n", hash)
}

func init() {
	rootCmd.AddCommand(faucetCmd)
}
