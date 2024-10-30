package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/189/golodash"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"githun.com/duiyuan/beetools/constants"
	"githun.com/duiyuan/beetools/core"
)

var faucetCmd = &cobra.Command{
	Use:    "faucet",
	Short:  "Faucet Sepolia ETH",
	Hidden: true,
	Run: func(cmd *cobra.Command, arg []string) {
		f := &Faucet{
			Chain:  chain,
			Token:  token,
			Wallet: wallet,
		}

		if err := f.Initial(); err != nil {
			fmt.Println(err)
			return
		}
		f.DripOne(token)
		f.Render()

	},
	Example: `
		beefaucet -w 0x17c3Da8F476a185f707aEA0f1C2A64F04142EbD7 --chain base --token usdt // Get one token
	`,
}

type TransferResp struct {
	Hash     string `json:"hash"`
	Amount   string `json:"amount"`
	Token    string `json:"token"`
	Chain    string `json:"chain"`
	Explorer string `json:"explorer"`
	Err      string `json:"error"`
}

type TxResp struct {
	Err     int            `json:"err"`
	Message string         `json:"message"`
	Data    []TransferResp `json:"data"`
}

func GetMaxLen(data []TransferResp) int {
	return golodash.Reduce(data, 10, func(p int, n TransferResp) int {
		l := len(n.Chain)
		if l > p {
			return l
		}
		return p
	})
}

func claim(chain string, token string, wallet string) (item *TransferResp, err error) {
	resp, err := core.ClaimToken[TxResp](
		fmt.Sprintf("%s/%s", endpoint, "transfer"),
		strings.ToLower(chain),
		strings.ToLower(token),
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

type Faucet struct {
	Chain       string
	Token       string
	Wallet      string
	TableWriter *tablewriter.Table
}

func (f *Faucet) Initial() error {
	f.TableWriter = tablewriter.NewWriter(os.Stdout)
	f.TableWriter.SetHeader([]string{"Chain", "Token", "Amount", "Status", "StatusText", "Link"})
	return f.ShouldParams()
}

func (f *Faucet) ShouldParams() error {
	chain := f.Chain
	token := f.Token

	lowerChains := golodash.Map(constants.ChainLongNames, func(name string) string {
		return strings.ToLower(name)
	})

	chainExpected := strings.Join(lowerChains, ",")
	tokenExpected := strings.Join(constants.Tokens, ",")

	has := golodash.Some(constants.ChainLongNames, func(ch string) bool {
		return strings.ToLower(ch) == chain
	})

	if !has {
		return fmt.Errorf("unsupported chain:%s, chain must be one of %s", chain, chainExpected)
	}

	has = golodash.Some(constants.Tokens, func(t string) bool {
		return t == token
	})

	if !has {
		return fmt.Errorf("unsupported token:%s, token must be one of %s", token, tokenExpected)
	}

	return nil
}

func (f *Faucet) Render() {
	f.TableWriter.SetBorder(true)
	f.TableWriter.SetHeaderLine(true)
	f.TableWriter.SetCenterSeparator("|")
	f.TableWriter.SetColumnSeparator("|")
	f.TableWriter.SetRowSeparator("-")
	f.TableWriter.Render()
}

func (f *Faucet) DripOne(token string) error {
	chain := f.Chain
	wallet := f.Wallet

	item, err := claim(chain, token, wallet)
	if err != nil {
		row := []string{
			chain, token, "0", "X", err.Error(), "",
		}
		f.TableWriter.Append(row)
		return err
	}

	chain, amount, token, explorer, e := item.Chain, item.Amount, item.Token, item.Explorer, item.Err
	if e == "" {
		row := []string{
			chain, token, amount, "OK", "OK", explorer,
		}
		f.TableWriter.Append(row)
		return nil
	} else {
		row := []string{
			chain, token, "0", "X", err.Error(), "",
		}
		f.TableWriter.Append(row)
		return err
	}
}

func (f *Faucet) Batch() {
	var succeed, failed int32
	var wg sync.WaitGroup

	wallet := f.Wallet

	fmt.Printf("\nStart to deposit Sepolia ETH to %s \n", wallet)
	fmt.Println("")

	for _, t := range []string{"ETH", "USDT"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := f.DripOne(t); err != nil {
				atomic.AddInt32(&failed, 1)
				return
			}
			atomic.AddInt32(&succeed, 1)
		}()
	}
	wg.Wait()

	fmt.Printf("\nTasks completed, succeed: %d, failed: %d \n", succeed, failed)
	fmt.Println("")
}

func init() {
	rootCmd.AddCommand(faucetCmd)
}
