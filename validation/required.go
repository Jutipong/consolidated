package validation

import (
	"consolidated/base"
	"fmt"
	"strings"
)

type RequiredRule struct {
	FieldName string
	Value     interface{}
}

//## Validate Rule = 1
func Required(fields *[]RequiredRule, errs *[]string) float32 {
	for _, f := range *fields {
		if strings.TrimSpace(fmt.Sprintf("%v", f.Value)) == "" || f.Value == nil {
			*errs = append(*errs, f.FieldName)
		}
	}

	if len(*errs) > 0 {
		return 1
	} else {
		return base.Successfully
	}
}
