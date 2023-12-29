package converter

import (
	"github.com/antlabs/deepcopy"
	"github.com/ugabiga/falcon/internal/ent"
	"github.com/ugabiga/falcon/internal/graph/generated"
)

func ToTaskHistory(inputData *ent.TaskHistory) (*generated.TaskHistory, error) {
	var result generated.TaskHistory
	if err := deepcopy.CopyEx(&result, inputData); err != nil {
		return nil, err
	}

	return &result, nil
}

func ToTaskHistories(inputData []*ent.TaskHistory) ([]*generated.TaskHistory, error) {
	resultList := make([]*generated.TaskHistory, 0, len(inputData))
	for _, v := range inputData {
		result, err := ToTaskHistory(v)
		if err != nil {
			return nil, err
		}
		resultList = append(resultList, result)
	}
	return resultList, nil
}
