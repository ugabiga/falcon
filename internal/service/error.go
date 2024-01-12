package service

import "errors"

var (
	ErrTickerNotFound          = errors.New("ticker_not_found")
	ErrExceedLimit             = errors.New("exceed_limit")
	ErrWrongExchange           = errors.New("wrong_exchange")
	ErrWrongCurrency           = errors.New("wrong_currency")
	ErrUnAuthorizedAction      = errors.New("unauthorized_action")
	ErrAlreadyExists           = errors.New("already_exists")
	ErrSizeNotSatisfiedMinimum = errors.New("size_not_satisfied_minimum")
)
