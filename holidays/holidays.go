package holidays

import (
	"time"
)

func IsHoliday(date time.Time) (bool, time.Duration) {
	earlyCloseTime, ok := calendar[date.Format("2006-01-02")]
	return ok, earlyCloseTime
}
