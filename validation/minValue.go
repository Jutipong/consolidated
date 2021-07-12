package validation

import "consolidated/base"

type MinValueRule struct {
	FieldName string
	Value     float64
	Min       int
}

//## Validate Rule = 2.2
func MinValue(fields *[]MinValueRule, errs *[]string) float32 {
	for _, f := range *fields {
		if f.Value <= float64(f.Min) {
			*errs = append(*errs, f.FieldName)
		}
	}

	if len(*errs) > 0 {
		return 2.2
	} else {
		return base.Successfully
	}
}
