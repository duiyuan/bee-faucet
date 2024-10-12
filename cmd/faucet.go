package cmd

import (
	"fmt"
	"sync"
	"sync/atomic"

	"githun.com/duiyuan/faucet/constants"
	"githun.com/duiyuan/faucet/thirdweb"
)

var (
	wg      sync.WaitGroup
	succeed int32
	failed  int32
)

func faucet(toAddr string) {
	blockchains := constants.ChainItems

	fmt.Printf("\nStart to faucet Sepolia ETH to %s \n", toAddr)
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
