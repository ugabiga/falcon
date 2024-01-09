package service

import (
	"github.com/adhocore/gronx"
	"time"
)

func nextCronExecutionTime(cron string, timezone string) (time.Time, error) {
	location, err := time.LoadLocation(timezone)
	if err != nil {
		return time.Time{}, err
	}

	localTime := time.Now().In(location)
	nextTime, err := gronx.NextTickAfter(cron, localTime, true)
	if err != nil {
		return time.Time{}, err
	}

	return nextTime.UTC().Truncate(time.Minute), nil
}
