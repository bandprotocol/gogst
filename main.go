package main

import (
	"fmt"

	"github.com/bandprotocol/gogst/markets"
)

func main() {
	// Returns 0, Invalid stock market
	fmt.Println(markets.GetMarketStatusByMarket(markets.UNDEFINED))

	fmt.Println(markets.GetMarketStatus("GLXY"))
}
