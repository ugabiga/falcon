package converter

import (
	"github.com/AlekSi/pointer"
	"reflect"
	"strconv"
)

func convertWhereIds(input interface{}, target interface{}) {
	val := reflect.ValueOf(input).Elem()
	targetVal := reflect.ValueOf(target).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := val.Type().Field(i).Name

		if fieldName == "ID" {
			if !field.IsNil() {
				targetField := targetVal.FieldByName("ID")
				targetField.Set(reflect.ValueOf(pointer.ToInt(StringToInt(field.Elem().String()))))
			}
		}

		if fieldName == "IDNeq" {
			if !field.IsNil() {
				targetField := targetVal.FieldByName("IDNEQ")
				targetField.Set(reflect.ValueOf(pointer.ToInt(StringToInt(field.Elem().String()))))
			}
		}

		if fieldName == "IDGt" {
			if !field.IsNil() {
				targetField := targetVal.FieldByName("IDGT")
				targetField.Set(reflect.ValueOf(pointer.ToInt(StringToInt(field.Elem().String()))))
			}
		}

		if fieldName == "IDGte" {
			if !field.IsNil() {
				targetField := targetVal.FieldByName("IDGTE")
				targetField.Set(reflect.ValueOf(pointer.ToInt(StringToInt(field.Elem().String()))))
			}
		}

		if fieldName == "IDLt" {
			if !field.IsNil() {
				targetField := targetVal.FieldByName("IDLT")
				targetField.Set(reflect.ValueOf(pointer.ToInt(StringToInt(field.Elem().String()))))
			}
		}

		if fieldName == "IDLte" {
			if !field.IsNil() {
				targetField := targetVal.FieldByName("IDLTE")
				targetField.Set(reflect.ValueOf(pointer.ToInt(StringToInt(field.Elem().String()))))
			}
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
