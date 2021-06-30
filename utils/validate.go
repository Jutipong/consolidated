package utils

import (
	"fmt"

	"github.com/gookit/validate"
)

var validationRule = map[float32]map[string]string{}

//## Helper
type ErrosMsg struct {
	Field       string
	Description []string
}

func GetValidateError(v *validate.Validation) interface{} {
	errs := []ErrosMsg{}

	for fieldName, validations := range v.Errors {
		err := ErrosMsg{
			Field:       fieldName,
			Description: make([]string, 0),
		}

		for _, errMessage := range validations {
			err.Description = append(err.Description, errMessage)
		}
		errs = append(errs, err)
	}
	return errs
}

func ValidRule(payload interface{}, field string, rules []Rule) (errs interface{}) {
	var strRules string
	v := validate.New(payload)

	for _, rule := range rules {
		if len(rule.Base) > 0 {
			strRules = fmt.Sprintf("%v", rule.Base)
		}

		_v := validationRule[rule.Id]
		switch _v["Type"] {
		case "":
			strRules += fmt.Sprintf("%v|", validationRule[rule.Id]["Rule"])
		case "number":
			strRules += fmt.Sprintf("%v:%v|", validationRule[rule.Id]["Rule"], rule.Value)
		}
	}

	v.StringRule(field, strRules)
	if !v.Validate() {
		errs = GetValidateError(v)
	}
	return errs
}

//## Master Config
type Rule struct {
	Id   float32
	Base string
	// Rule  string = "Rule"
	// Type  string
	Value int
}

func InitValidationRule() {
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
	descript["Rule"] = "numberx" //custom
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
	descript["Rule"] = "yyyymmmdd" //custom
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
