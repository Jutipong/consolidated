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
	// validationRule := map[float32]map[string]string{}

	//## 1 => required
	key = 1
	validationRule[key] = make(map[string]string)
	descript := map[string]string{}
	descript["Message"] = "Mandatory field"
	descript["ErrorCode"] = "V001"
	descript["Rule"] = "required"
	validationRule[key] = descript

	//## 2 => Field length
	key = 2
	validationRule[key] = make(map[string]string)
	descript = map[string]string{}
	descript["Message"] = "Field length"
	descript["ErrorCode"] = "V002"
	descript["Rule"] = "maxLen"
	descript["Type"] = "number"
	validationRule[key] = descript

	//## Character set
	//## 3.1 => Field length
	key = 3.1
	validationRule[key] = make(map[string]string)
	descript = map[string]string{}
	descript["Message"] = "Character set: a b c d e f g h i j k l m n o p q r s t u v w x y z"
	descript["ErrorCode"] = "V003"
	descript["Rule"] = "lower" //custom
	validationRule[key] = descript

	//## 3.2 => Field length
	key = 3.2
	validationRule[key] = make(map[string]string)
	descript = map[string]string{}
	descript["Message"] = "Character set: A B C D E F G H I J K L M N O P Q R S T U V W X Y Z"
	descript["ErrorCode"] = "V003"
	descript["Rule"] = "Upper" //custom
	validationRule[key] = descript

	//## 3.3 => Field length
	key = 3.3
	validationRule[key] = make(map[string]string)
	descript = map[string]string{}
	descript["Message"] = "Character set: 0 1 2 3 4 5 6 7 8 9 ."
	descript["ErrorCode"] = "V003"
	descript["Rule"] = "number" //custom
	validationRule[key] = descript

	//## 4 => Field length
	key = 4
	validationRule[key] = make(map[string]string)
	descript = map[string]string{}
	descript["Message"] = "Fix value"
	descript["ErrorCode"] = "V004"
	descript["Rule"] = "fixValue" //custom
	validationRule[key] = descript

	//##Date pattern
	//## 5.1 => Field length
	key = 5.1
	validationRule[key] = make(map[string]string)
	descript = map[string]string{}
	descript["Message"] = "format: YYYYMMDD"
	descript["ErrorCode"] = "V005"
	descript["Rule"] = "YYYYMMDD" //custom
	validationRule[key] = descript

	//## 5.2 => Field length
	key = 5.2
	validationRule[key] = make(map[string]string)
	descript = map[string]string{}
	descript["Message"] = "format: HH:MM:SS"
	descript["ErrorCode"] = "V005"
	descript["Rule"] = "hhmmss" //custom
	validationRule[key] = descript

	//## 5.3 => Field length
	key = 5.3
	validationRule[key] = make(map[string]string)
	descript = map[string]string{}
	descript["Message"] = "format: YYYYMMDD HH:MM:SS"
	descript["ErrorCode"] = "V005"
	descript["Rule"] = "yyyymmmddhhmmss" //custom
	validationRule[key] = descript

	// b, _ := json.Marshal(validationRule[key])
	// fmt.Println(string(b))
}
