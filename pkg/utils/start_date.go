package utils

import (
	"time"

	"github.com/ibra-bybuy/go-wsports-events/pkg/constants"
	"github.com/ibra-bybuy/go-wsports-events/pkg/utils/datetime"
)

func StartDate(sport string) string {
	t := time.Now().Truncate(time.Hour - 10)
	if sport == constants.FOOTBALL_TYPE {
		t = time.Now().Truncate(time.Minute - 150)
	} else if sport == constants.MMA_TYPE {
		t = time.Now().Truncate(time.Hour - 10)
	}

	return datetime.Full(t)
}
