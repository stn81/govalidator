package govalidator

import (
	"fmt"
	"reflect"
	"regexp"
)

// RegExValidator is the regexp validator, performs better than regex tag
type RegExValidator struct {
	Pattern string
	RegEx   *regexp.Regexp
}

// NewRegExValidator create a new regexp validator
func NewRegExValidator(pattern string) *RegExValidator {
	regex := regexp.MustCompile(pattern)
	return &RegExValidator{
		Pattern: pattern,
		RegEx:   regex,
	}
}

// Validate implements the Validator interface
func (v *RegExValidator) Validate(value interface{}, args ...string) error {
	str, ok := value.(string)
	if !ok {
		panic(fmt.Errorf("not string type: %v", reflect.TypeOf(value)))
	}

	if str == "" {
		return nil
	}

	if v.RegEx.MatchString(str) {
		return nil
	}
	return ErrRegexpNotMatch(str, v.Pattern)
}
