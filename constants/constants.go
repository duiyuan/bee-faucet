package constants

import "githun.com/duiyuan/faucet/utils"

type ChainItem struct {
	Name string
	ID   int32
}

type Short string

const (
	ETH     = 11155111
	ARB     = 421614
	BASE    = 11155420
	OPT     = 11155420
	BNB     = 97
	POLYGON = 1442
	B3      = 1993
	MODE    = 919
	CREATOR = 4654
	XPLA    = 47
)

var ChainList = map[Short]ChainItem{
	"ETH":     {ID: ETH, Name: "Ethereum Sepolia"},
	"Base":    {ID: BASE, Name: "Base"},
	"ARB":     {ID: ARB, Name: "Arbitrum"},
	"OPT":     {ID: OPT, Name: "Optimism"},
	"POLY":    {ID: POLYGON, Name: "Polygon"},
	"BNB":     {ID: BNB, Name: "BNB"},
	"B3":      {ID: B3, Name: "B3"},
	"Creator": {ID: CREATOR, Name: "Creator"},
	"Xpla":    {ID: XPLA, Name: "Xpla"},
}

var ChainShortNames = utils.GetKeys(ChainList)
var ChainItems = utils.GetValues(ChainList)
var ChainLongNames = utils.Map(ChainItems, func(item ChainItem) string {
	return item.Name
})

var MaxChainNameLen = utils.Reduce(ChainLongNames, 0, func(prev int, now string) int {
	if prev > len(now) {
		return prev
	}
	return len(now)
})
