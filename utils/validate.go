package utils

import "github.com/gookit/validate"

type ErrosMsg struct {
	Field       string
	Description []map[string]string
}

func GetValidateError(v *validate.Validation) interface{} {
	errs := []ErrosMsg{}

	for fieldName, validations := range v.Errors {
		err := ErrosMsg{
			Field:       fieldName,
			Description: make([]map[string]string, 0),
		}

		for validateType, errMessage := range validations {
			des := map[string]string{}
			des[validateType] = errMessage
			err.Description = append(err.Description, des)
		}
		errs = append(errs, err)
	}
	return errs
}
