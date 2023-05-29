package utils

import (
	"time"

	"github.com/ibra-bybuy/go-wsports-events/pkg/utils/datetime"
)

func EndDate() string {
	return datetime.YearMonthDayHoursMinsSecs(time.Now().Add(-(time.Hour * 2)).UTC()) + ".000Z"
}
