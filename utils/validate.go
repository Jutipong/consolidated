package utils

import (
	"consolidated/base"
	"fmt"
	"strconv"
	"strings"

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
	FieldName  string
	Value      string
	ValueFloat float64

	Length    int
	Condition []string
}

//## Validate Request By library
func ValidateByRule(obj interface{}, ruleId float32, validateRules []ValidateRules) (errMsg []string) {
	_rule := base.GetRule(ruleId)

	switch ruleId {
	case 2.1:
		validateFloat64(_rule, validateRules, &errMsg)
	case 4:
		validateByCondition(_rule, validateRules, &errMsg)
	default:
		validateLibrary(_rule, validateRules, &errMsg)
	}

	return errMsg
}

func validateLibrary(_rule map[string]string, validateRules []ValidateRules, errMsg *[]string) {
	for _, rule := range validateRules {
		for _, fields := range rule.RuleFields {
			var strRules string
			v := validate.New(rule.Obj)

			switch _rule["Type"] {
			case "":
				strRules = fmt.Sprintf("%v|", _rule["Rule"])
			case "number":
				strRules = fmt.Sprintf("%v:%v|", _rule["Rule"], fields.Length)
			}

			v.StringRule(fields.FieldName, strRules)
			if !v.Validate() {
				fmt.Println(v.Errors)
				for fieldName, value := range v.Errors {
					fmt.Println(value)
					*errMsg = append(*errMsg, fieldName)
				}
			}
		}
	}
}

func validateByCondition(_rule map[string]string, validateRules []ValidateRules, errMsg *[]string) {
	for _, rule := range validateRules {
		for _, fields := range rule.RuleFields {
			if !Contains(fields.Condition, fields.Value) {
				*errMsg = append(*errMsg, fields.FieldName)
			}
		}
	}
}

func validateFloat64(_rule map[string]string, validateRules []ValidateRules, errMsg *[]string) {
	validateFloat64Length(_rule, validateRules, errMsg)
	if len(*errMsg) == 0 {
		validateDigit(_rule, validateRules, errMsg)
	}
}

func validateFloat64Length(_rule map[string]string, validateRules []ValidateRules, errMsg *[]string) {
	for _, rule := range validateRules {
		for _, fields := range rule.RuleFields {
			str := strconv.FormatFloat(fields.ValueFloat, 'f', -1, 64)
			if len(str) > fields.Length {
				*errMsg = append(*errMsg, fields.FieldName)
			}
		}
	}
}

func validateDigit(_rule map[string]string, validateRules []ValidateRules, errMsg *[]string) {
	for _, rule := range validateRules {
		for _, fields := range rule.RuleFields {
			arr := strings.SplitN(strconv.FormatFloat(fields.ValueFloat, 'f', -1, 64), ".", 2)
			fmt.Println(arr)
			if len(arr) != 2 {
				*errMsg = append(*errMsg, fields.FieldName)
			} else {
				if len(arr[1]) > 2 {
					*errMsg = append(*errMsg, fields.FieldName)
				}
			}
		}
	}
}

// func validateFloat64MaxValue(_rule map[string]string, validateRules []ValidateRules, errMsg *[]string) {
// 	for _, rule := range validateRules {
// 		for _, fields := range rule.RuleFields {
// 			str := strconv.FormatFloat(fields.ValueFloat, 'f', -1, 64)
// 			if len(str) > fields.Length {
// 				*errMsg = append(*errMsg, fields.FieldName)
// 			}
// 		}
// 	}
// }

//## GetValidateError
type ErrsMsg struct {
	Field       string
	Description []string
}

func GetValidateError(v *validate.Validation) interface{} {
	errs := []ErrsMsg{}

	for fieldName, validations := range v.Errors {
		err := ErrsMsg{
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
