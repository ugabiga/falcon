package service

import (
	"fmt"
	"github.com/adhocore/gronx"
	"github.com/ugabiga/falcon/internal/ent"
	"time"
)

func dbRollback(tx *ent.Tx, err error) error {
	if rollbackErr := tx.Rollback(); rollbackErr != nil {
		err = fmt.Errorf("%w: %v", err, rollbackErr)
	}
	return err
}
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

	return nextTime.UTC(), nil
}
