package holidays

import (
	"time"

	"github.com/bandprotocol/gogst/utils"
)

// Mapping between the holiday date to the early closing hour
// If the early closing hour is -1, then the market closes that day
// Assumption: early closing hour means no post market
var calendar = map[string]time.Duration{
	"2022-01-17": -1,
	"2022-02-21": -1,
	"2022-04-15": -1,
	"2022-05-30": -1,
	"2022-06-20": -1,
	"2022-07-04": -1,
	"2022-09-05": -1,
	"2022-11-24": utils.Time(13, 0, 0),
	"2022-12-26": -1,
}
