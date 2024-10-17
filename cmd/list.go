package cmd

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"
	"githun.com/duiyuan/faucet/constants"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all blockchains you can faucet.",
	Run:   listAllChain,
}

func listAllChain(cmd *cobra.Command, args []string) {
	items := constants.ChainItems

	sort.Slice(items, func(i, j int) bool {
		return items[i].Name < items[j].Name
	})

	for _, chain := range constants.ChainItems {
		supported := "Not supported"
		if chain.Supported {
			supported = "Supported"
		}
		format := fmt.Sprintf("%%-%ds ID:%%-10d (%%s)\n", constants.MaxChainNameLen)
		fmt.Printf(format, chain.Name, chain.ID, supported)
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
}
