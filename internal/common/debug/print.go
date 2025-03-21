package debug

import (
	"encoding/json"
	"log"
	"strings"
)

func ToJSONStr(obj interface{}) string {
	bytes, err := json.MarshalIndent(obj, "", "\t")
	if err != nil {
		log.Println("Error marshalling JSON:", err)
		return ""
	}

	return "\n" + string(bytes)
}
func ToJSONInlineStr(obj interface{}) string {
	bytes, err := json.Marshal(obj)
	if err != nil {
		log.Println("Error marshalling JSON:", err)
		return ""
	}
	formattedStr := string(bytes)
	formattedStr = strings.ReplaceAll(formattedStr, `\n`, "")

	return " " + formattedStr
}
func FromByteToJSONInLineStr(obj []byte) string {
	var target map[string]interface{}
	if err := json.Unmarshal(obj, &target); err != nil {
		log.Println(err)
		return ""
	}

	bytes, err := json.Marshal(target)
	if err != nil {
		log.Println("Error marshalling JSON:", err)
		return ""
	}
	formattedStr := string(bytes)
	formattedStr = strings.ReplaceAll(formattedStr, `\n`, "")

	return " " + formattedStr
}
