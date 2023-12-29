package converter

import (
	"github.com/antlabs/deepcopy"
	"github.com/ugabiga/falcon/internal/ent"
	"github.com/ugabiga/falcon/internal/graph/generated"
)

func ToTask(inputData *ent.Task) (*generated.Task, error) {
	var result generated.Task

	if err := deepcopy.CopyEx(&result, inputData); err != nil {
		return nil, err
	}

	result.TaskHistories = make([]*generated.TaskHistory, 0, len(inputData.Edges.TaskHistories))
	for _, v := range inputData.Edges.TaskHistories {
		newVal, err := ToTaskHistory(v)
		if err != nil {
			return nil, err
		}
		result.TaskHistories = append(result.TaskHistories, newVal)
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
