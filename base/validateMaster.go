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

	//## 1 => required
	key = 1
	validationRule[key] = make(map[string]string)
	description := map[string]string{}
	description["Message"] = "Mandatory field"
	description["Code"] = "V001"
	validationRule[key] = description

	//## 2 => Field max length
	key = 2
	validationRule[key] = make(map[string]string)
	description = map[string]string{}
	description["Message"] = "Invalid fields length"
	description["Code"] = "V002"
	validationRule[key] = description

	//## 2.1 => Field min length
	key = 2.1
	validationRule[key] = make(map[string]string)
	description = map[string]string{}
	description["Message"] = "Invalid fields length"
	description["Code"] = "V002"
	validationRule[key] = description

	//## 2.2 => Digit length
	key = 2.2
	validationRule[key] = make(map[string]string)
	description = map[string]string{}
	description["Message"] = "Invalid fields length (.2 digit)"
	description["Code"] = "V002"
	validationRule[key] = description

	//## Character set
	//## 3.1 => Field length
	// key = 3.1
	// validationRule[key] = make(map[string]string)
	// description = map[string]string{}
	// description["Message"] = "Character set: a b c d e f g h i j k l m n o p q r s t u v w x y z"
	// description["Code"] = "V003"
	// description["Rule"] = "lower" //custom
	// validationRule[key] = description

	//## 3.2 => Field length
	// key = 3.2
	// validationRule[key] = make(map[string]string)
	// description = map[string]string{}
	// description["Message"] = "Character set: A B C D E F G H I J K L M N O P Q R S T U V W X Y Z"
	// description["Code"] = "V003"
	// description["Rule"] = "Upper" //custom
	// validationRule[key] = description

	//## 3.3 => Field length
	key = 3
	validationRule[key] = make(map[string]string)
	description = map[string]string{}
	description["Message"] = "Character set: 0 1 2 3 4 5 6 7 8 9 ."
	description["Code"] = "V003"
	validationRule[key] = description

	//## 4 => Field length
	key = 4
	validationRule[key] = make(map[string]string)
	description = map[string]string{}
	description["Message"] = "Fix value"
	description["Code"] = "V004"
	validationRule[key] = description

	//##Date pattern
	//## 5.1 => Field length
	key = 5.1
	validationRule[key] = make(map[string]string)
	description = map[string]string{}
	description["Message"] = "Invalid fields type"
	description["Code"] = "V005"
	validationRule[key] = description

	//## 5.2 => Field length
	key = 5.2
	validationRule[key] = make(map[string]string)
	description = map[string]string{}
	description["Message"] = "format: HH:MM:SS"
	description["Code"] = "V005"
	validationRule[key] = description

	//## 5.3 => Field length
	key = 5.3
	validationRule[key] = make(map[string]string)
	description = map[string]string{}
	description["Message"] = "format: YYYYMMDD HH:MM:SS"
	description["Code"] = "V005"
	validationRule[key] = description

	key = 400
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
