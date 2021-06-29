package utils

import "encoding/json"

func SerializeObject(objs interface{}) string {
	b, err := json.Marshal(&objs)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func DeserializeObject(str string, st interface{}) interface{} {
	st = json.Unmarshal([]byte(str), &st)
	return st
}
