package validation

import (
	"consolidated/utils"
)

type FixValueRule struct {
	FieldName  string
	Value      string
	Conditions []string
}

//## Validate Rule = 4
func FixValue(fields []FixValueRule, errs *[]string) float32 {
	for _, f := range fields {
		if !utils.Contains(f.Conditions, f.Value) {
			*errs = append(*errs, f.FieldName)
		}
	}

	if len(*errs) > 0 {
		return 4
	} else {
		return 0000
	}
}
