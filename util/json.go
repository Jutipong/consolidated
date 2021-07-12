package util

import "encoding/json"

//## ตัวอย่างการใช้งาน => util.JsonSerialize(payload)
func JsonSerialize(payload interface{}) string {
	b, err := json.Marshal(&payload)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

//## ตัวอย่างการใช้งาน => util.JsonDeserialize(str, &result)
func JsonDeserialize(str string, st interface{}) {
	json.Unmarshal([]byte(str), &st)
}
