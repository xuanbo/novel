package util

import "encoding/json"

func ToJsonByte(data interface{}) []byte {
	result, err := json.Marshal(data)
	if err != nil {
		return []byte("")
	}
	return result
}
