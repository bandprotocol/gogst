package markets

import (
	"time"

	"github.com/bandprotocol/gogst/utils"
)

type MarketDetail struct {
	TimeZone        string
	OpenPreMarket   time.Duration
	OpenRegMarket   time.Duration
	CloseRegMarket  time.Duration
	ClosePostMarket time.Duration
}

var MARKETS = map[Market]MarketDetail{
	US_STOCK: {"America/New_York", utils.Time(4, 0, 0), utils.Time(9, 30, 0), utils.Time(16, 0, 0), utils.Time(20, 0, 0)},
	CA_STOCK: {"America/New_York", -1, utils.Time(9, 30, 0), utils.Time(16, 0, 0), utils.Time(17, 0, 0)},
}
