package service

import (
	"context"
	"errors"
	"github.com/ugabiga/falcon/internal/ent"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrWrongExchange = errors.New("wrong_exchange")
	ErrWrongCurrency = errors.New("wrong_currency")
)

type TradingAccountService struct {
	db *ent.Client
}

func NewTradingAccountService(db *ent.Client) *TradingAccountService {
	return &TradingAccountService{db: db}
}

func (s TradingAccountService) Create(
	ctx context.Context,
	userID uint64,
	exchange string,
	currency string,
	Identifier string,
	credential string,
	phrase string,
) (
	*ent.TradingAccount, error,
) {
	if err := s.validateExchange(exchange); err != nil {
		return nil, err
	}

	if err := s.validateCurrency(currency); err != nil {
		return nil, err
	}

	encryptedCredential, err := s.encryptCredential(credential)
	if err != nil {
		return nil, err
	}

	ip, err := s.availableIP()
	if err != nil {
		return nil, err
	}

	createQuery := s.db.TradingAccount.Create().
		SetUserID(userID).
		SetExchange(exchange).
		SetCurrency(currency).
		SetIP(ip).
		SetIdentifier(Identifier).
		SetCredential(encryptedCredential)

	if phrase != "" {
		createQuery.SetPhrase(phrase)
	}

	t, err := createQuery.Save(ctx)
	if err != nil {
		return nil, err
	}
	return t, nil

}

func (s TradingAccountService) validateExchange(exchange string) error {
	switch exchange {
	case "binance":
		return nil
	case "upbit":
		return nil
	default:
		return ErrWrongExchange
	}
}

func (s TradingAccountService) validateCurrency(currency string) error {
	// currency code ISO 4217
	switch currency {
	case "KRW":
		return nil
	case "USD":
		return nil
	default:
		return ErrWrongCurrency
	}
}

func (s TradingAccountService) encryptCredential(credential string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(credential), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(password), nil
}

func (s TradingAccountService) availableIP() (string, error) {
	// TODO : implement
	return "", nil
}
