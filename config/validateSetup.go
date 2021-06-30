package config

import (
	"time"

	"github.com/gookit/validate"
)

func SetupCustomValidate() {
	_validationRule := MasterRule()

	validate.Config(func(opt *validate.GlobalOption) {
		opt.StopOnError = false
		// opt.SkipOnEmpty = false
	})

	//## Tag => YYYYMMDD
	validate.AddValidator("YYYYMMDD", func(strDate string) bool {
		_, err := time.Parse("20060102", strDate)
		if err == nil {
			return true
		} else {
			return false
		}
	})
	validate.AddGlobalMessages(map[string]string{
		"YYYYMMDD": _validationRule[5.1]["Message"],
	})

	//
	validate.AddGlobalMessages(map[string]string{
		"required": _validationRule[1]["Message"],
	})
}
