package validation

import "time"

type DateTimeRule struct {
	FieldName string
	Value     string
}

//## Validate Rule = 5.1
func YYYYMMDD(fields []DateTimeRule, errs *[]string) float32 {
	for _, f := range fields {
		_, err := time.Parse("20060102", f.Value)
		if err != nil {
			*errs = append(*errs, f.FieldName)
		}
	}

	if len(*errs) > 0 {
		return 5.1
	} else {
		return 0000
	}
}

//## Validate Rule = 5.2
func HHMMSS(fields []DateTimeRule, errs *[]string) float32 {
	for _, f := range fields {
		_, err := time.Parse("15:04:05", f.Value)
		if err != nil {
			*errs = append(*errs, f.FieldName)
		}
	}

	if len(*errs) > 0 {
		return 5.2
	} else {
		return 0000
	}
}
