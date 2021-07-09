package base

var validationRule = map[float32]map[string]string{}

func MasterRule() map[float32]map[string]string {
	return validationRule
}

func GetRule(id float32) map[string]string {
	return validationRule[id]
}

func InitMasterRule() {
	var key float32

	//## 1 => Mandatory field
	key = 1
	validationRule[key] = make(map[string]string)
	description := map[string]string{}
	description["Message"] = "Mandatory field"
	description["Code"] = "V001"
	validationRule[key] = description

	//## 2 => Invalid fields length
	key = 2
	validationRule[key] = make(map[string]string)
	description = map[string]string{}
	description["Message"] = "Invalid fields length"
	description["Code"] = "V002"
	validationRule[key] = description

	//## 2.1 => Invalid fields length (.2 digit)
	key = 2.1
	validationRule[key] = make(map[string]string)
	description = map[string]string{}
	description["Message"] = "Invalid fields length (.2 digit)"
	description["Code"] = "V002"
	validationRule[key] = description

	//## 3.3 => Character set: 0 1 2 3 4 5 6 7 8 9 .
	key = 3
	validationRule[key] = make(map[string]string)
	description = map[string]string{}
	description["Message"] = "Character set: 0 1 2 3 4 5 6 7 8 9 ."
	description["Code"] = "V003"
	validationRule[key] = description

	//## 4 => Fix value
	key = 4
	validationRule[key] = make(map[string]string)
	description = map[string]string{}
	description["Message"] = "Fix value"
	description["Code"] = "V004"
	validationRule[key] = description

	//##Date pattern
	//## 5.1 => Invalid fields type
	key = 5.1
	validationRule[key] = make(map[string]string)
	description = map[string]string{}
	description["Message"] = "Invalid fields type"
	description["Code"] = "V005"
	validationRule[key] = description

	//## 5.2 => format: HH:MM:SS
	key = 5.2
	validationRule[key] = make(map[string]string)
	description = map[string]string{}
	description["Message"] = "format: HH:MM:SS"
	description["Code"] = "V005"
	validationRule[key] = description

	//## 5.3 => format: YYYYMMDD HH:MM:SS
	key = 5.3
	validationRule[key] = make(map[string]string)
	description = map[string]string{}
	description["Message"] = "format: YYYYMMDD HH:MM:SS"
	description["Code"] = "V005"
	validationRule[key] = description

	//=================================//
	//		Stutus		   //
	//=================================//
	key = 400
	validationRule[key] = make(map[string]string)
	description = map[string]string{}
	description["Message"] = "Bad Request"
	description["Code"] = "400"
	validationRule[key] = description

	key = 401
	validationRule[key] = make(map[string]string)
	description = map[string]string{}
	description["Message"] = "Unauthorized"
	description["Code"] = "400"
	validationRule[key] = description

	key = 0000
	validationRule[key] = make(map[string]string)
	description = map[string]string{}
	description["Message"] = "Success"
	description["Code"] = "V001"
	validationRule[key] = description
}
