package govalidator

import (
	"fmt"
	"reflect"
)

// StrIDValidator the strid tag validator
type StrIDValidator struct {
	RegExValidator Validator
	MinLen         int
	MaxLen         int
}

// NewStrIDValidator create a new strid validator
func NewStrIDValidator(minLen, maxLen int, regex string) Validator {
	v := &StrIDValidator{
		MinLen:         minLen,
		MaxLen:         maxLen,
		RegExValidator: NewRegExValidator(regex),
	}
	return v
}

// Validate implements the Validator interface
func (v *StrIDValidator) Validate(value interface{}, args ...string) error {
	str, ok := value.(string)
	if !ok {
		panic(fmt.Errorf("expected string type, not %v", reflect.TypeOf(value)))
	}

	if len(str) < v.MinLen || len(str) > v.MaxLen {
		return fmt.Errorf("string length not in [%v,%v]", v.MinLen, v.MaxLen)
	}
	return v.RegExValidator.Validate(value)
}
