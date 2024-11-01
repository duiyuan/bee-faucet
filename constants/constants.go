package constants

import (
	"github.com/189/golodash"
)

type ChainItem struct {
	Name      string
	ID        int32
	Supported bool
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

var Tokens = []string{"eth", "usdt"}

var ChainList = map[Short]ChainItem{
	"ETH":  {ID: ETH, Name: "Ethereum", Supported: true},
	"Base": {ID: BASE, Name: "Base", Supported: true},
	"ARB":  {ID: ARB, Name: "Arbitrum", Supported: true},
	"OPT":  {ID: OPT, Name: "Optimism", Supported: true},
	// "POLY":    {ID: POLYGON, Name: "Polygon", Supported: false},
	// "BNB":     {ID: BNB, Name: "BNB", Supported: false},
	// "B3":      {ID: B3, Name: "B3", Supported: false},
	// "Creator": {ID: CREATOR, Name: "Creator", Supported: false},
	// "Xpla":    {ID: XPLA, Name: "Xpla", Supported: false},
}

var ChainShortNames = golodash.GetKeys(ChainList)
var ChainItems = golodash.GetValues(ChainList)
var ChainLongNames = golodash.Map(ChainItems, func(item ChainItem) string {
	return item.Name
})

var MaxChainNameLen = golodash.Reduce(ChainLongNames, 0, func(prev int, now string) int {
	if prev > len(now) {
		return prev
	}
	return len(now)
})
