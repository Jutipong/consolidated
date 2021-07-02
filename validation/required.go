package validation

type RuleField struct {
	FidelName string
	Value     interface{}
}

//## Validate Rule = 1
func Required(fields []RuleField, errs *[]string) float32 {
	for _, f := range fields {
		if f.Value == "" || f.Value == nil {
			*errs = append(*errs, f.FidelName)
		}
	}

	if len(*errs) > 0 {
		return 1
	} else {
		return 0000
	}
}
