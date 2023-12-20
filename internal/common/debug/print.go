package debug

import (
	"encoding/json"
	"log"
)

func ToJSONStr(obj interface{}) string {
	bytes, err := json.MarshalIndent(obj, "", "\t")
	if err != nil {
		log.Println("Error marshalling JSON:", err)
		return ""
	}

	return "\n" + string(bytes)
}
