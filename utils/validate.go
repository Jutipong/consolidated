package utils

import (
	"consolidated/config"
	"fmt"

	"github.com/gookit/validate"
)

// type RuleValidate struct {
// 	RuleId float32
// 	Rules  []Rules
// }

type ValidateRules struct {
	Obj        interface{}
	RuleFields []RuleField
}

type RuleField struct {
	FieldName string
	Len       int
	Condition []string
}

//## ใช้ในการ Validate Request ตามลำดับ Rule => 1, 2, 3, .......
func ValidateByRule(obj interface{}, ruleId float32, validateRules []ValidateRules) []string {
	_rule := config.GetRule(ruleId)
	var errMsg []string

	for _, rule := range validateRules {
		for _, fields := range rule.RuleFields {
			var strRules string
			v := validate.New(rule.Obj)

			switch _rule["Type"] {
			case "":
				if len(fields.Condition) == 0 {
					strRules = fmt.Sprintf("%v|", _rule["Rule"])
				}
				// else {
				// 	//condition
				// }
			case "number":
				strRules = fmt.Sprintf("%v:%v|", _rule["Rule"], fields.Len)
			}

			v.StringRule(fields.FieldName, strRules)
			if !v.Validate() {
				for fieldName := range v.Errors {
					errMsg = append(errMsg, fieldName)
				}
			}
		}
	}

	return errMsg
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

// //## ValidRule
// type Rule struct {
// 	Id    float32
// 	Base  string
// 	Value int
// }

// //## ใช้สำหรับ Validate model ทั่วไป
// func ValidField(obj interface{}, field string, rules []Rule) (errs interface{}) {
// 	validationRule := config.MasterRule()

// 	for _, rule := range rules {
// 		v := validate.New(obj)
// 		var strRules string
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

// 		v.StringRule(field, strRules)
// 		if !v.Validate() {
// 			errs = GetValidateError(v)
// 			return errs
// 		}
// 	}
// 	return errs
// }

//Backup
// //## ใช้ในการ Validate Request ตามลำดับ Rule => 1, 2, 3, .......
// func ValidateByRule(obj interface{}, ruleId float32, rules []ValidateRule) []string {
// 	_rule := config.GetRule(ruleId)
// 	var errMsg []string

// 	for _, rule := range rules {
// 		var strRules string
// 		v := validate.New(obj)

// 		switch _rule["Type"] {
// 		case "":
// 			if len(rule.Condition) == 0 {
// 				strRules = fmt.Sprintf("%v|", _rule["Rule"])
// 			}
// 			// else {
// 			// 	//condition
// 			// }
// 		case "number":
// 			strRules = fmt.Sprintf("%v:%v|", _rule["Rule"], rule.Value)
// 		}

// 		v.StringRule(rule.FieldName, strRules)
// 		if !v.Validate() {
// 			for _, er := range GetFieldNameError(v) {
// 				errMsg = append(errMsg, er)
// 			}
// 		}
// 	}

// 	return errMsg
// }

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
