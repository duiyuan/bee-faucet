package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"githun.com/duiyuan/faucet/core"
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
}

type TxResp struct {
	Err     int            `json:"err"`
	Message string         `json:"message"`
	Data    []TransferResp `json:"data"`
}

func Dofaucet(chain string, wallet string) {
	resp, err := core.ClaimToken[TxResp](
		chain,
		wallet,
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	// var resp TxResp
	// if err := mapstructure.Decode(result, &resp); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	if err, message := resp.Err, resp.Message; err != 0 {
		fmt.Println(message)
		return
	}

	data := resp.Data

	fmt.Printf("\nStart to Deposit Sepolia ETH to %s \n", wallet)
	fmt.Println("")
	for _, item := range data {
		format := fmt.Sprintf("🚀 %%-%ds success, faucet %%s %%s! Hash: %%s\n", 10)
		fmt.Printf(format, item.Chain, item.Amount, item.Token, item.Hash)
	}
	fmt.Printf("\nEnd faucet, succeed: %d, failed: %d \n", succeed, failed)
	fmt.Println("")
}

func init() {
	rootCmd.AddCommand(faucetCmd)
}
