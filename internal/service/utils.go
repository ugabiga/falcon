package service

import (
	"errors"
	"github.com/adhocore/gronx"
	"time"
)

var (
	ErrTickerNotFound     = errors.New("ticker not found")
	ErrExceedLimit        = errors.New("exceed_limit")
	ErrWrongExchange      = errors.New("wrong_exchange")
	ErrWrongCurrency      = errors.New("wrong_currency")
	ErrorNoRows           = errors.New("no_rows")
	ErrDoNotHaveAccess    = errors.New("do_not_have_access")
	ErrUnAuthorizedAction = errors.New("unauthorized_action")
	ErrAlreadyExist       = errors.New("trading_account_already_exist")
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
