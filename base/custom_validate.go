package base

import (
	"time"

	"github.com/gookit/validate"
)

func SetupCustomValidate() {
	validate.Config(func(opt *validate.GlobalOption) {
		opt.StopOnError = false
		// opt.SkipOnEmpty = false
	})

	//## Tag => YYYYMMDD
	validate.AddValidator("YYYYMMDD", func(strDate string) bool {
		_, err := time.Parse("2006-01-02", strDate)
		if err == nil {
			return true
		} else {
			return false
		}
	})
	validate.AddGlobalMessages(map[string]string{
		"YYYYMMDD": "Date pattern นะจ๊ะ",
	})
}
