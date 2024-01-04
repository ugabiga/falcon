package service

import "errors"

var (
	ErrExceedLimit     = errors.New("exceed_limit")
	ErrWrongExchange   = errors.New("wrong_exchange")
	ErrWrongCurrency   = errors.New("wrong_currency")
	ErrorNoRows        = errors.New("no_rows")
	ErrDoNotHaveAccess = errors.New("do_not_have_access")
	ErrUnauthorized    = errors.New("unauthorized")
)
