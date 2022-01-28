package main

import (
	"fmt"
	"log"

	"github.com/bandprotocol/gogst/markets"
)

func main() {
	// Returns invalid Invalid stock market
	fmt.Println(markets.GetMarketStatusByMarket(markets.UNDEFINED))

	status, err := markets.GetMarketStatus("AAPL")
	if err != nil {
		log.Fatal(err)
	}
	if status == markets.PRE || status == markets.POST {
		fmt.Println("Extended time")
	} else {
		fmt.Printf("%s time\n", status)
	}
}
