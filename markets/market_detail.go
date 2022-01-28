package markets

import "time"

type MarketDetail struct {
	TimeZone        string
	OpenPreMarket   time.Duration
	OpenRegMarket   time.Duration
	CloseRegMarket  time.Duration
	ClosePostMarket time.Duration
}

func Time(hour, min, sec int) time.Duration {
	return time.Duration(hour)*time.Hour + time.Duration(min)*time.Minute + time.Duration(sec)*time.Second
}

var MARKETS = map[Market]MarketDetail{
	US_STOCK: {"America/New_York", Time(4, 0, 0), Time(9, 30, 0), Time(16, 0, 0), Time(20, 0, 0)},
	CA_STOCK: {"America/New_York", -1, Time(9, 30, 0), Time(16, 0, 0), Time(17, 0, 0)},
}
