package helper

import (
	"errors"
	"fmt"
	"reflect"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// FieldError contains all functions to get error details

type FieldError interface {

	// returns the validation tag that failed. if the
	// validation was an alias, this will return the
	// alias name and not the underlying tag that failed.
	//
	// eg. alias "iscolor": "hexcolor|rgb|rgba|hsl|hsla"
	// will return "iscolor"
	Tag() string

	// returns the validation tag that failed, even if an
	// alias the actual tag within the alias will be returned.
	// If an 'or' validation fails the entire or will be returned.
	//
	// eg. alias "iscolor": "hexcolor|rgb|rgba|hsl|hsla"
	// will return "hexcolor|rgb|rgba|hsl|hsla"
	ActualTag() string

	// returns the namespace for the field error, with the tag
	// name taking precedence over the fields actual name.
	//
	// eg. JSON name "User.fname"
	//
	// See StructNamespace() for a version that returns actual names.
	//
	// NOTE: this field can be blank when validating a single primitive field
	// using validate.Field(...) as there is no way to extract it's name
	Namespace() string

	// returns the namespace for the field error, with the fields
	// actual name.
	//
	// eq. "User.FirstName" see Namespace for comparison
	//
	// NOTE: this field can be blank when validating a single primitive field
	// using validate.Field(...) as there is no way to extract it's name
	StructNamespace() string

	// returns the fields name with the tag name taking precedence over the
	// fields actual name.
	//
	// eq. JSON name "fname"
	// see StructField for comparison
	Field() string

	// returns the fields actual name from the struct, when able to determine.
	//
	// eq.  "FirstName"
	// see Field for comparison
	StructField() string

	// returns the actual fields value in case needed for creating the error
	// message
	Value() interface{}

	// returns the param value, in string form for comparison; this will also
	// help with generating an error message
	Param() string

	// Kind returns the Field's reflect Kind
	//
	// eg. time.Time's kind is a struct
	Kind() reflect.Kind

	// Type returns the Field's reflect Type
	//
	// // eg. time.Time's type is time.Time
	Type() reflect.Type

	// returns the FieldError's translated error
	// from the provided 'ut.Translator' and registered 'TranslationFunc'
	//
	// NOTE: if no registered translator can be found it returns the same as
	// calling fe.Error()
	//Error()
	Translate(ut ut.Translator) string
}

type ValidationErrors []FieldError

type ValidationError struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}

//var Validate = validator.New()

func Descriptive(verr validator.ValidationErrors) []ValidationError {
	errs := []ValidationError{}
	for _, f := range verr {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}
		if err == "required" {
			errs = append(errs, ValidationError{Field: f.Field(), Reason: "field is required"})
		} else {
			errs = append(errs, ValidationError{Field: f.Field(), Reason: err})
		}
	}
	return errs
}

func GetErrShouldBind(err error) []ValidationError {
	errs := []ValidationError{}
	if err != nil {
		var valid validator.ValidationErrors
		if errors.As(err, &valid) {
			errs = append(errs, Descriptive(valid)...)
		}
	}
	return errs
}
