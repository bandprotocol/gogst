package markets

import (
	"fmt"
	"time"

	"github.com/bandprotocol/gogst/holidays"
	"github.com/bandprotocol/gogst/utils"
)

var STOCK_BASE = map[string]Market{
	"AAPL":  US_STOCK,
	"ARKK":  US_STOCK,
	"AMC":   US_STOCK,
	"COIN":  US_STOCK,
	"MSFT":  US_STOCK,
	"AMD":   US_STOCK,
	"HOOD":  US_STOCK,
	"GS":    US_STOCK,
	"AMZN":  US_STOCK,
	"TSLA":  US_STOCK,
	"GOOGL": US_STOCK,
	"SLV":   US_STOCK,
	"ABNB":  US_STOCK,
	"SQ":    US_STOCK,
	"FB":    US_STOCK,
	"IAU":   US_STOCK,
	"NFLX":  US_STOCK,
	"USO":   US_STOCK,
	"GME":   US_STOCK,
	"QQQ":   US_STOCK,
	"BABA":  US_STOCK,
	"VIXY":  US_STOCK,
	"SPY":   US_STOCK,
	"TWTR":  US_STOCK,
	"GLXY":  CA_STOCK,
}

func GetMarketStatusByMarket(stock Market) (MarketStatus, error) {
	marketDetail, ok := MARKETS[stock]
	if !ok {
		return INVALID, fmt.Errorf("Invalid stock market")
	}

	tz, err := time.LoadLocation(marketDetail.TimeZone)
	if err != nil {
		return INVALID, err
	}
	now := time.Now().In(tz)

	holiday, earlyCloseTime := holidays.IsHoliday(now)
	if holiday {
		marketDetail.CloseRegMarket = earlyCloseTime
		marketDetail.ClosePostMarket = -1
	}

	if now.Weekday() == time.Saturday || now.Weekday() == time.Sunday || earlyCloseTime == -1 {
		return CLOSE, nil
	}

	d := utils.Time(now.Clock())
	if marketDetail.OpenPreMarket != -1 {
		if marketDetail.OpenPreMarket <= d && d < marketDetail.OpenRegMarket {
			return PRE, nil
		}
	}

	if marketDetail.OpenRegMarket <= d && d < marketDetail.CloseRegMarket {
		return REG, nil
	}

	if marketDetail.ClosePostMarket != -1 {
		if marketDetail.CloseRegMarket <= d && d < marketDetail.ClosePostMarket {
			return POST, nil
		}
	}

	return CLOSE, nil
}

func GetMarketStatus(symbol string) (MarketStatus, error) {
	stock, ok := STOCK_BASE[symbol]
	if !ok {
		return INVALID, fmt.Errorf("%s is not supported", symbol)
	}
	return GetMarketStatusByMarket(stock)
}
