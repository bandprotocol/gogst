package markets

import (
	"fmt"
	"time"

	"github.com/bandprotocol/gogst/holidays"
	"github.com/bandprotocol/gogst/utils"
)

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
