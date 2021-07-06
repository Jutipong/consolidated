package utils

import "encoding/json"

func JsonSerialize(objs interface{}) string {
	b, err := json.Marshal(&objs)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func JsonDeserialize(str string, st interface{}) {
	json.Unmarshal([]byte(str), &st)
}
