package converter

import (
	"github.com/antlabs/deepcopy"
	"github.com/ugabiga/falcon/internal/ent"
	"github.com/ugabiga/falcon/internal/graph/generated"
)

func ToTask(inputData *ent.Task) (*generated.Task, error) {
	var result generated.Task
	result.ID = IntToString(inputData.ID)
	if err := deepcopy.CopyEx(&result, inputData); err != nil {
		return nil, err
	}

	return &result, nil
}

func ToTasks(inputData []*ent.Task) ([]*generated.Task, error) {
	resultList := make([]*generated.Task, 0, len(inputData))
	for _, v := range inputData {
		result, err := ToTask(v)
		if err != nil {
			return nil, err
		}
		resultList = append(resultList, result)
	}
	return resultList, nil
}
