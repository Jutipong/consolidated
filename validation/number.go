package validation

import (
	"consolidated/base"
	"regexp"
)

type NumberRule struct {
	FieldName string
	Value     string
}

//## Validate Rule = 3
func Number(fields *[]NumberRule, errs *[]string) float32 {
	regexp := regexp.MustCompile(`^[0-9\.\/]+$$`)
	for _, f := range *fields {
		if match := regexp.MatchString(f.Value); !match {
			*errs = append(*errs, f.FieldName)
		}
	}

	if len(*errs) > 0 {
		return 3
	} else {
		return base.Successfully
	}
}
