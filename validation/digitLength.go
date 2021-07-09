package validation

import (
	"fmt"
	"strconv"
	"strings"
)

type DigitLengthRule struct {
	FieldName string
	Value     float64
}

//## Validate Rule = 2.1
func DigitLength(fields []DigitLengthRule, errs *[]string) float32 {
	for _, f := range fields {
		arr := strings.SplitN(strconv.FormatFloat(f.Value, 'f', -1, 64), ".", 2)
		fmt.Println(arr)
		if len(arr) != 2 {
			*errs = append(*errs, f.FieldName)
		} else {
			if len(arr[1]) > 2 {
				*errs = append(*errs, f.FieldName)
			}
		}
	}

	if len(*errs) > 0 {
		return 2.1
	} else {
		return 0000
	}
}
