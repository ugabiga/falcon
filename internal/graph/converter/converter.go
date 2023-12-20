package converter

import (
	"github.com/AlekSi/pointer"
	"github.com/antlabs/deepcopy"
	"reflect"
	"strconv"
)

func BindWhereInput(input, target interface{}) interface{} {
	if err := deepcopy.CopyEx(target, input); err != nil {
		return target
	}

	convertWhereIDs(input, target)

	return target
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

func convertWhereIDs(input interface{}, target interface{}) {
	val := reflect.ValueOf(input).Elem()
	targetVal := reflect.ValueOf(target).Elem()

	idPair := map[string]string{
		"ID":      "ID",
		"IDNeq":   "IDNEQ",
		"IDGt":    "IDGT",
		"IDGte":   "IDGTE",
		"IDLt":    "IDLT",
		"IDLte":   "IDLTE",
		"IDIn":    "IDIn",
		"IDNotIn": "IDNotIn",
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := val.Type().Field(i).Name

		if targetFieldName, ok := idPair[fieldName]; ok {
			convertWhereID(field, targetVal, targetFieldName)
			continue
		}

		if fieldName == "IDIn" || fieldName == "IDNotIn" {
			if !field.IsNil() {
				targetField := targetVal.FieldByName(fieldName)
				for idx := 0; idx < field.Len(); idx++ {
					targetField.Index(idx).Set(reflect.ValueOf(
						StringToInt(field.Index(idx).String())),
					)
				}
			}
		}

	}
}

func convertWhereID(field, targetVal reflect.Value, targetFieldName string) {
	if !field.IsNil() {
		targetField := targetVal.FieldByName(targetFieldName)
		targetField.Set(reflect.ValueOf(pointer.ToInt(StringToInt(field.Elem().String()))))
	}
}
