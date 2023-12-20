package service

import (
	"context"
	"errors"
	"github.com/ugabiga/falcon/internal/ent"
	"github.com/ugabiga/falcon/internal/ent/tradingaccount"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrWrongExchange = errors.New("wrong_exchange")
	ErrWrongCurrency = errors.New("wrong_currency")
	ErrorNoRows      = errors.New("no_rows")
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

	encryptedCredential, err := s.encrypt(credential)
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
		encryptedPhrase, err := s.encrypt(phrase)
		if err != nil {
			return nil, err
		}
		createQuery.SetPhrase(encryptedPhrase)
	}

	t, err := createQuery.Save(ctx)
	if err != nil {
		return nil, err
	}
	return t, nil

}

func (s TradingAccountService) Get(ctx context.Context, userID uint64) ([]*ent.TradingAccount, error) {
	return s.db.TradingAccount.Query().Where(
		tradingaccount.UserIDEQ(userID),
	).All(ctx)
}

func (s TradingAccountService) GetByID(ctx context.Context, userID, tradingAccountID uint64) (*ent.TradingAccount, error) {
	return s.db.TradingAccount.Query().Where(
		tradingaccount.UserIDEQ(userID),
		tradingaccount.IDEQ(tradingAccountID),
	).First(ctx)
}

func (s TradingAccountService) Update(
	ctx context.Context,
	tradingAccountID uint64,
	userID uint64,
	exchange string,
	currency string,
	Identifier string,
	credential string,
	phrase string,
) error {
	if err := s.validateExchange(exchange); err != nil {
		return err
	}

	if err := s.validateCurrency(currency); err != nil {
		return err
	}

	encryptedCredential, err := s.encrypt(credential)
	if err != nil {
		return err
	}

	updateQuery := s.db.TradingAccount.Update().
		Where(
			tradingaccount.IDEQ(tradingAccountID),
			tradingaccount.UserIDEQ(userID),
		).
		SetExchange(exchange).
		SetCurrency(currency).
		SetIdentifier(Identifier).
		SetCredential(encryptedCredential)

	if phrase != "" {
		encryptedPhrase, err := s.encrypt(phrase)
		if err != nil {
			return err
		}
		updateQuery.SetPhrase(encryptedPhrase)
	}

	updateCount, err := updateQuery.Save(ctx)
	if err != nil {
		return err
	}

	if updateCount <= 0 {
		return ErrorNoRows
	}

	return nil
}

func (s TradingAccountService) Delete(ctx context.Context, userID, tradingAccountID uint64) error {
	deleteCount, err := s.db.TradingAccount.Delete().Where(
		tradingaccount.IDEQ(tradingAccountID),
		tradingaccount.UserIDEQ(userID),
	).Exec(ctx)
	if err != nil {
		return err
	}

	if deleteCount <= 0 {
		return ErrorNoRows
	}

	return nil
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

func (s TradingAccountService) encrypt(credential string) (string, error) {
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
