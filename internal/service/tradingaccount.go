package service

import (
	"context"
	"github.com/AlekSi/pointer"
	"github.com/ugabiga/falcon/internal/ent"
	"github.com/ugabiga/falcon/internal/ent/tradingaccount"
	"golang.org/x/crypto/bcrypt"
)

const (
	TradingAccountCreationLimit = 2
)

type TradingAccountService struct {
	db *ent.Client
}

func NewTradingAccountService(db *ent.Client) *TradingAccountService {
	return &TradingAccountService{db: db}
}

func (s TradingAccountService) Create(
	ctx context.Context,
	userID int,
	name string,
	exchange string,
	key string,
	secret string,
	phrase string,
) (
	*ent.TradingAccount, error,
) {
	if err := s.validateExchange(exchange); err != nil {
		return nil, err
	}

	encryptedSecret, err := s.encrypt(secret)
	if err != nil {
		return nil, err
	}

	ip, err := s.availableIP()
	if err != nil {
		return nil, err
	}

	if err = s.validateExceedLimit(ctx, userID); err != nil {
		return nil, err
	}

	createQuery := s.db.TradingAccount.Create().
		SetUserID(userID).
		SetName(name).
		SetExchange(exchange).
		SetIP(ip).
		SetKey(key).
		SetSecret(encryptedSecret)

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

func (s TradingAccountService) validateExceedLimit(ctx context.Context, userID int) error {
	count, err := s.db.TradingAccount.Query().Where(
		tradingaccount.UserIDEQ(userID),
	).Count(ctx)
	if err != nil {
		return err
	}

	if count >= TradingAccountCreationLimit {
		return ErrExceedLimit
	}

	return nil
}

func (s TradingAccountService) Get(ctx context.Context, userID int) ([]*ent.TradingAccount, error) {
	return s.db.TradingAccount.Query().Where(
		tradingaccount.UserIDEQ(userID),
	).
		Order(ent.Desc(tradingaccount.FieldID)).
		All(ctx)
}

func (s TradingAccountService) GetWithTask(ctx context.Context, userID int) ([]*ent.TradingAccount, error) {
	query := s.db.TradingAccount.Query().Where(
		tradingaccount.UserIDEQ(userID),
	)

	return query.
		Order(ent.Desc(tradingaccount.FieldID)).
		WithTasks().
		All(ctx)
}

func (s TradingAccountService) First(ctx context.Context, userID int) (*ent.TradingAccount, error) {
	return s.db.TradingAccount.Query().Where(
		tradingaccount.UserIDEQ(userID),
	).
		Order(ent.Desc(tradingaccount.FieldID)).
		First(ctx)
}

func (s TradingAccountService) GetByID(ctx context.Context, userID, tradingAccountID int) (*ent.TradingAccount, error) {
	return s.db.TradingAccount.Query().Where(
		tradingaccount.UserIDEQ(userID),
		tradingaccount.IDEQ(tradingAccountID),
	).First(ctx)
}

func (s TradingAccountService) Update(
	ctx context.Context,
	tradingAccountID int,
	userID int,
	name *string,
	exchange *string,
	key *string,
	secret *string,
	phrase *string,
) error {
	if exchange == nil && key == nil && secret == nil && phrase == nil {
		return nil
	}

	if exchange != nil {
		if err := s.validateExchange(pointer.GetString(exchange)); err != nil {
			return err
		}
	}

	updateQuery := s.db.TradingAccount.Update().
		Where(
			tradingaccount.IDEQ(tradingAccountID),
			tradingaccount.UserIDEQ(userID),
		)
	if name != nil {
		updateQuery.SetName(pointer.GetString(name))
	}
	if exchange != nil {
		updateQuery.SetExchange(pointer.GetString(exchange))
	}
	if key != nil {
		updateQuery.SetKey(pointer.GetString(key))
	}
	if secret != nil {
		encryptedSecret, err := s.encrypt(pointer.GetString(secret))
		if err != nil {
			return err
		}
		updateQuery.SetSecret(encryptedSecret)
	}

	if phrase != nil {
		encryptedPhrase, err := s.encrypt(pointer.GetString(phrase))
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

func (s TradingAccountService) Delete(ctx context.Context, userID, tradingAccountID int) error {
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

func (s TradingAccountService) encrypt(secret string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(secret), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(password), nil
}

func (s TradingAccountService) availableIP() (string, error) {
	// TODO : implement
	return "192.168.0.1", nil
}
