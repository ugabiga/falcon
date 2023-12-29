package converter

import (
	"github.com/antlabs/deepcopy"
	"github.com/ugabiga/falcon/internal/ent"
	"github.com/ugabiga/falcon/internal/graph/generated"
)

func ToTradingAccount(inputData *ent.TradingAccount) (*generated.TradingAccount, error) {
	var result generated.TradingAccount
	if err := deepcopy.CopyEx(&result, inputData); err != nil {
		return nil, err
	}

	result.Tasks = make([]*generated.Task, 0, len(inputData.Edges.Tasks))
	for _, v := range inputData.Edges.Tasks {
		newVal, err := ToTask(v)
		if err != nil {
			return nil, err
		}
		result.Tasks = append(result.Tasks, newVal)
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
