package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"githun.com/duiyuan/faucet/constants"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all blockchains you can faucet.",
	Run:   listAllChain,
}

func listAllChain(cmd *cobra.Command, args []string) {

	for _, chain := range constants.ChainItems {
		format := fmt.Sprintf("%%-%ds ID:%%-10d (supported)\n", constants.MaxChainNameLen)
		fmt.Printf(format, chain.Name, chain.ID)
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
}
