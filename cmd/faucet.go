package cmd

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/spf13/cobra"
	"githun.com/duiyuan/faucet/constants"
	"githun.com/duiyuan/faucet/core"
	"githun.com/duiyuan/faucet/utils"
)

var faucetCmd = &cobra.Command{
	Use:    "faucet",
	Short:  "Faucet Sepolia ETH",
	Hidden: true,
	Run: func(cmd *cobra.Command, arg []string) {
		Dofaucet(chain, wallet)
	},
}

type TransferResp struct {
	Hash   string `json:"hash"`
	Amount string `json:"amount"`
	Token  string `json:"token"`
	Chain  string `json:"chain"`
	Err    string `json:"error"`
}

type TxResp struct {
	Err     int            `json:"err"`
	Message string         `json:"message"`
	Data    []TransferResp `json:"data"`
}

func GetMaxLen(data []TransferResp) int {
	return utils.Reduce(data, 10, func(p int, n TransferResp) int {
		l := len(n.Chain)
		if l > p {
			return l
		}
		return p
	})
}

func claim(chain string) (item *TransferResp, err error) {
	resp, err := core.ClaimToken[TxResp](
		fmt.Sprintf("%s/%s", endpoint, "transfer"),
		strings.ToLower(chain),
		wallet,
	)

	if err != nil {
		return
	}

	e, message := resp.Err, resp.Message
	if e != 0 {
		err = errors.New(message)
		return
	}

	data := resp.Data
	item = &data[0]
	return
}

func Dofaucet(chain string, wallet string) {
	var succeed, failed int32
	var wg sync.WaitGroup

	chains := utils.Filter(constants.ChainItems, func(item constants.ChainItem) bool {
		return item.Supported
	})

	fmt.Printf("\nStart to deposit Sepolia ETH to %s \n", wallet)
	fmt.Println("")

	for _, c := range chains {
		wg.Add(1)
		go func() {
			var format string

			defer wg.Done()

			item, err := claim(c.Name)
			if err != nil {
				format = fmt.Sprintf("❌ %%-%ds error %%s\n", 10)
				fmt.Printf(format, c.Name, err)
				atomic.AddInt32(&failed, 1)
				return
			}

			chain, amount, token, hash, e := item.Chain, item.Amount, item.Token, item.Hash, item.Err
			if e == "" {
				format = fmt.Sprintf("✅ %%-%ds faucet %%s %%s! (%%s)\n", 10)
				fmt.Printf(format, chain, amount, token, hash)
				atomic.AddInt32(&succeed, 1)
			} else {
				format = fmt.Sprintf("❌ %%-%ds error %%s %%s! (%%s)\n", 10)
				fmt.Printf(format, chain, amount, token, err)
				atomic.AddInt32(&failed, 1)
			}
		}()
	}
	wg.Wait()

	fmt.Printf("\nTask completed, succeed: %d, failed: %d \n", succeed, failed)
	fmt.Println("")
}

func init() {
	rootCmd.AddCommand(faucetCmd)
}
