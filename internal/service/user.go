package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/ugabiga/falcon/internal/ent"
	"github.com/ugabiga/falcon/internal/ent/authentication"
	"github.com/ugabiga/falcon/internal/ent/user"
)

type UserService struct {
	db *ent.Client
}

func NewUserService(
	db *ent.Client,
) *UserService {
	return &UserService{
		db: db,
	}
}

type UserWithOptions struct {
	WithAuthentications bool
	WithTradingAccounts bool
}

func (s UserService) Query(ctx context.Context, where ent.UserWhereInput) ([]*ent.User, error) {
	_, err := s.db.User.Create().
		SetName("dummy").
		SetTimezone("UTC").
		Save(ctx)
	if err != nil {
		return nil, err
	}

	query, err := where.Filter(s.db.User.Query())
	if err != nil {
		return nil, err
	}
	users, err := query.All(ctx)

	return users, nil
}

func (s UserService) CreateDummy(
	ctx context.Context, userID int, withOptions UserWithOptions,
) (
	*ent.User, error,
) {
	u, err := s.db.User.Create().
		SetName("dummy").
		SetTimezone("UTC").
		Save(ctx)
	if err != nil {
		return nil, err
	}

	if withOptions.WithAuthentications {
		a, err := s.db.Authentication.Create().
			SetUserID(u.ID).
			SetIdentifier(uuid.New().String()).
			SetCredential(uuid.New().String()).
			SetProvider(authentication.ProviderGoogle).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		u.Edges.Authentications = append(u.Edges.Authentications, a)
	}

	if withOptions.WithTradingAccounts {
		ta, err := s.db.TradingAccount.Create().
			SetUserID(u.ID).
			SetExchange("binance").
			SetIdentifier(uuid.New().String()).
			SetCredential(uuid.New().String()).
			SetIP("127.0.0.1").
			Save(ctx)
		if err != nil {
			return nil, err
		}
		u.Edges.TradingAccounts = append(u.Edges.TradingAccounts, ta)
	}

	return u, nil
}

func (s UserService) GetByID(ctx context.Context, userID int) (*ent.User, error) {
	u, err := s.db.User.Query().
		Where(
			user.IDEQ(userID),
		).
		First(ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}
func (s UserService) Update(ctx context.Context, userID int, user *ent.User) (*ent.User, error) {
	u, err := s.db.User.UpdateOneID(userID).
		SetName(user.Name).
		SetTimezone(user.Timezone).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}
