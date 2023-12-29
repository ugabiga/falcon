package converter

import (
	"github.com/antlabs/deepcopy"
	"github.com/ugabiga/falcon/internal/ent"
	"github.com/ugabiga/falcon/internal/graph/generated"
)

func ToUser(inputUser *ent.User) (*generated.User, error) {
	var newUser generated.User
	if err := deepcopy.CopyEx(&newUser, inputUser); err != nil {
		return nil, err
	}

	newUser.Authentications = make([]*generated.Authentication, 0, len(inputUser.Edges.Authentications))
	for _, v := range inputUser.Edges.Authentications {
		var a generated.Authentication
		if err := deepcopy.CopyEx(&a, v); err != nil {
			return nil, err
		}
		newUser.Authentications = append(newUser.Authentications, &a)
	}

	newUser.TradingAccounts = make([]*generated.TradingAccount, 0, len(inputUser.Edges.TradingAccounts))
	for _, v := range inputUser.Edges.TradingAccounts {
		var a generated.TradingAccount
		if err := deepcopy.CopyEx(&a, v); err != nil {
			return nil, err
		}
		newUser.TradingAccounts = append(newUser.TradingAccounts, &a)
	}

	return &newUser, nil
}

func ToUsers(inputUsers []*ent.User) ([]*generated.User, error) {
	users := make([]*generated.User, 0, len(inputUsers))
	for _, v := range inputUsers {
		user, err := ToUser(v)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
