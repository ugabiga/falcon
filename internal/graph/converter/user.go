package converter

import (
	"github.com/antlabs/deepcopy"
	"github.com/ugabiga/falcon/internal/ent"
	"github.com/ugabiga/falcon/internal/graph/generated"
	"strconv"
)

func ToUser(inputUser *ent.User) (*generated.User, error) {
	var newUser generated.User
	newUser.ID = IntToString(inputUser.ID)
	if err := deepcopy.CopyEx(&newUser, inputUser); err != nil {
		return nil, err
	}

	newUser.Authentications = make([]*generated.Authentication, 0, len(inputUser.Edges.Authentications))
	for _, v := range inputUser.Edges.Authentications {
		var a generated.Authentication
		if err := deepcopy.CopyEx(&a, v); err != nil {
			return nil, err
		}
		a.ID = IntToString(v.ID)
		newUser.Authentications = append(newUser.Authentications, &a)
	}

	newUser.TradingAccounts = make([]*generated.TradingAccount, 0, len(inputUser.Edges.TradingAccounts))
	for _, v := range inputUser.Edges.TradingAccounts {
		var a generated.TradingAccount
		if err := deepcopy.CopyEx(&a, v); err != nil {
			return nil, err
		}
		a.ID = IntToString(v.ID)
		newUser.TradingAccounts = append(newUser.TradingAccounts, &a)
	}

	return &newUser, nil
}

func IntToString(iid int) string {
	return strconv.Itoa(iid)
}

func StringToInt(sid string) int {
	atoi, err := strconv.Atoi(sid)
	if err != nil {
		return 0
	}
	return atoi
}
