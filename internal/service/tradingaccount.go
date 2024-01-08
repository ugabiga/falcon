package service

import (
	"context"
	"github.com/AlekSi/pointer"
	"github.com/ugabiga/falcon/internal/common/encryption"
	"github.com/ugabiga/falcon/internal/model"
	"github.com/ugabiga/falcon/internal/repository"
	"log"
)

const (
	TradingAccountCreationLimit = 2
)

type TradingAccountService struct {
	encryption *encryption.Encryption
	repo       *repository.DynamoRepository
}

func NewTradingAccountService(
	encryption *encryption.Encryption,
	repo *repository.DynamoRepository,
) *TradingAccountService {
	return &TradingAccountService{
		encryption: encryption,
		repo:       repo,
	}
}

func (s TradingAccountService) Create(ctx context.Context, userID string, name string, exchange string, key string, secret string, phrase string) (*model.TradingAccount, error) {
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

	tradingAccount := model.TradingAccount{
		UserID:   userID,
		Name:     name,
		Exchange: exchange,
		IP:       ip,
		Key:      key,
		Secret:   encryptedSecret,
	}

	if phrase != "" {
		encryptedPhrase, err := s.encrypt(phrase)
		if err != nil {
			return nil, err
		}
		tradingAccount.Phrase = encryptedPhrase
	}

	newTradingAccount, err := s.repo.CreateTradingAccount(
		ctx,
		tradingAccount,
	)
	if err != nil {
		return nil, err
	}

	return newTradingAccount, nil

}

func (s TradingAccountService) validateExceedLimit(ctx context.Context, userID string) error {
	count, err := s.repo.CountTradingAccountsByUserID(ctx, userID)
	if err != nil {
		return err
	}

	if count >= TradingAccountCreationLimit {
		return ErrExceedLimit
	}

	return nil
}

func (s TradingAccountService) GetByUserID(ctx context.Context, userID string) ([]model.TradingAccount, error) {
	tradingAccounts, err := s.repo.GetTradingAccountsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return tradingAccounts, nil
}

func (s TradingAccountService) Update(
	ctx context.Context,
	tradingAccountID string,
	userID string,
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

	tradingAccount, err := s.repo.GetTradingAccount(ctx, userID, tradingAccountID)
	if err != nil {
		return err
	}

	if tradingAccount.UserID != userID {
		return ErrUnAuthorizedAction
	}

	inputTradingAccount := model.TradingAccount{
		ID:        tradingAccountID,
		UserID:    userID,
		IP:        tradingAccount.IP,
		Secret:    tradingAccount.Secret,
		CreatedAt: tradingAccount.CreatedAt,
	}

	if name != nil {
		inputTradingAccount.Name = *name
	}
	if exchange != nil {
		inputTradingAccount.Exchange = *exchange
	}
	if key != nil {
		inputTradingAccount.Key = *key
	}
	if secret != nil {
		encryptedSecret, err := s.encrypt(pointer.GetString(secret))
		if err != nil {
			return err
		}
		inputTradingAccount.Secret = encryptedSecret
	}
	if phrase != nil {
		encryptedPhrase, err := s.encrypt(pointer.GetString(phrase))
		if err != nil {
			return err
		}
		inputTradingAccount.Phrase = encryptedPhrase
	}

	_, err = s.repo.UpdateTradingAccount(ctx, inputTradingAccount)
	if err != nil {
		return err
	}

	return nil
}

func (s TradingAccountService) Delete(ctx context.Context, userID string, tradingAccountID string) error {
	log.Printf("delete trading account: %s, %s", userID, tradingAccountID)
	tradingAccount, err := s.repo.GetTradingAccount(ctx, userID, tradingAccountID)
	if err != nil {
		log.Printf("failed to get trading account: %v", err)
		return err
	}

	if tradingAccount.UserID != userID {
		return ErrUnAuthorizedAction
	}

	tasks, err := s.repo.GetTasksByTradingAccountID(ctx, tradingAccountID)
	if err != nil {
		log.Printf("failed to get tasks: %v", err)
		return err
	}

	for _, task := range tasks {
		if err := s.repo.DeleteTask(ctx, tradingAccountID, task.ID); err != nil {
			log.Printf("failed to delete task: %v", err)
			return err
		}
	}

	if err := s.repo.DeleteTradingAccount(ctx, userID, tradingAccountID); err != nil {
		log.Printf("failed to delete trading account: %v", err)
		return err
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
	return s.encryption.Encrypt(secret)
}

func (s TradingAccountService) availableIP() (string, error) {
	// TODO : implement
	return "192.168.0.1", nil
}
