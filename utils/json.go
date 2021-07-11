package utils

import "encoding/json"

func JsonSerialize(payload interface{}) string {
	b, err := json.Marshal(&payload)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func JsonDeserialize(str string, st interface{}) {
	json.Unmarshal([]byte(str), &st)
}
