package validation

import (
	"consolidated/base"
	"fmt"
	"strconv"
	"unicode/utf8"
)

type MaxLengthRule struct {
	FieldName string
	Value     interface{}
	Length    int
}

//## Validate Rule = 2
func MaxLength(fields *[]MaxLengthRule, errs *[]string) float32 {
	for _, f := range *fields {
		switch f.Value.(type) {
		case float64:
			str := strconv.FormatFloat(f.Value.(float64), 'f', -1, 64)
			if len(str) > f.Length {
				*errs = append(*errs, f.FieldName)
			}
		default:
			len := utf8.RuneCountInString(fmt.Sprintf("%v", f.Value))
			if len > f.Length {
				*errs = append(*errs, f.FieldName)
			}
		}
	}

	if len(*errs) > 0 {
		return 2
	} else {
		return base.Successfully
	}
}
