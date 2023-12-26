package converter

import (
	"github.com/antlabs/deepcopy"
	"github.com/ugabiga/falcon/internal/ent"
	"github.com/ugabiga/falcon/internal/graph/generated"
)

func ToTradingAccount(inputUser *ent.TradingAccount) (*generated.TradingAccount, error) {
	var result generated.TradingAccount
	result.ID = IntToString(inputUser.ID)
	if err := deepcopy.CopyEx(&result, inputUser); err != nil {
		return nil, err
	}

	return &result, nil
}

func ToTradingAccounts(inputData []*ent.TradingAccount) ([]*generated.TradingAccount, error) {
	resultList := make([]*generated.TradingAccount, 0, len(inputData))
	for _, v := range inputData {
		result, err := ToTradingAccount(v)
		if err != nil {
			return nil, err
		}
		resultList = append(resultList, result)
	}
	return resultList, nil
}
