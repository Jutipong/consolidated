package utils

import (
	"consolidated/config"
	"fmt"

	"github.com/gookit/validate"
)

//## ValidRule
type Rule struct {
	Id    float32
	Base  string
	Value int
}

//## ใช้สำหรับ Validate model ทั่วไป
func ValidField(obj interface{}, field string, rules []Rule) (errs interface{}) {
	validationRule := config.MasterRule()

	for _, rule := range rules {
		v := validate.New(obj)
		var strRules string
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

		v.StringRule(field, strRules)
		if !v.Validate() {
			errs = GetValidateError(v)
			return errs
		}
	}
	return errs
}

type ValidateRule struct {
	FieldName string
	Value     int
	Condition []string
}

//## ใช้ในการ Validate Request ตามลำดับ Rule => 1, 2, 3, .......
func ValidateByRule(obj interface{}, ruleId float32, value int, rules []ValidateRule) (string, string, []string) {
	_rule := config.GetRule(ruleId)
	var errMsg []string

	for _, rule := range rules {
		var strRules string
		v := validate.New(obj)

		switch _rule["Type"] {
		case "":
			if len(rule.Condition) == 0 {
				strRules = fmt.Sprintf("%v|", _rule["Rule"])
			}
			// else {
			// 	//condition
			// }
		case "number":
			strRules = fmt.Sprintf("%v:%v|", _rule["Rule"], value)
		}

		v.StringRule(rule.FieldName, strRules)
		if !v.Validate() {
			for _, er := range GetFieldNameError(v) {
				errMsg = append(errMsg, er)
			}
		}
	}

	if len(errMsg) == 0 {
		return "", "", errMsg
	} else {
		return _rule["ErrorCode"], _rule["Message"], errMsg
	}

}

//## GetValidateError
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

func GetFieldNameError(v *validate.Validation) []string {
	errs := []string{}

	for fieldName := range v.Errors {
		errs = append(errs, fieldName)
	}
	return errs
}

// func ValidFields(payload interface{}, field string, rules []Rule) (errs interface{}) {
// 	var strRules string
// 	v := validate.New(payload)
// 	validationRule := base.MasterRule()

// 	for _, rule := range rules {
// 		if len(rule.Base) > 0 {
// 			strRules = fmt.Sprintf("%v", rule.Base)
// 		}

// 		_v := validationRule[rule.Id]
// 		switch _v["Type"] {
// 		case "":
// 			strRules += fmt.Sprintf("%v|", validationRule[rule.Id]["Rule"])
// 		case "number":
// 			strRules += fmt.Sprintf("%v:%v|", validationRule[rule.Id]["Rule"], rule.Value)
// 		}
// 	}

// 	v.StringRule(field, strRules)
// 	if !v.Validate() {
// 		errs = GetValidateError(v)
// 	}
// 	return errs
// }
