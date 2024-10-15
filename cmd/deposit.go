package cmd

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/spf13/cobra"
	"githun.com/duiyuan/faucet/constants"
	"githun.com/duiyuan/faucet/thirdweb"
)

var (
	wg      sync.WaitGroup
	succeed int32
	failed  int32
)

var depositCmd = &cobra.Command{
	Use:   "deposit",
	Short: "Deposit sepolia to build-in accounts",
	Run: func(cmd *cobra.Command, arg []string) {
		Deposit(thirdweb.Account)
	},
}

func Deposit(toAddr string) {
	blockchains := constants.ChainItems

	fmt.Printf("\nStart to Deposit Sepolia ETH to %s \n", toAddr)
	fmt.Println("")

	for _, chain := range blockchains {
		wg.Add(1)
		go func() {
			if err := thirdweb.ClaimToken(chain.ID, chain.Name, toAddr); err != nil {
				atomic.AddInt32(&failed, 1)
			} else {
				atomic.AddInt32(&succeed, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Printf("\nEnd faucet, succeed: %d, failed: %d \n", succeed, failed)
}

func init() {
	rootCmd.AddCommand(depositCmd)
}
