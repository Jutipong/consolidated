package utils

import (
	"consolidated/base"
	"fmt"

	"github.com/gookit/validate"
)

//## ValidRule
type Rule struct {
	Id    float32
	Base  string
	Value int
}

func ValidField(payload interface{}, field string, rules []Rule) (errs interface{}) {
	var strRules string
	v := validate.New(payload)
	validationRule := base.MasterRule()

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
