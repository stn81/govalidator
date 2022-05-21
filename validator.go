package govalidator

const (
	DefaultTag         = "valid"
	DefaultTagValueSep = ";"
)

type SelfValidator interface {
	Validate() error
}

type Validator interface {
	Validate(value interface{}, args ...string) error
}

type ValidateFunc func(value interface{}, args ...string) error

func (f ValidateFunc) Validate(value interface{}, args ...string) error {
	return f(value, args...)
}

func ValidateStruct(ptr interface{}) error {
	validator := structValidators.get(ptr)
	return validator.Validate(ptr)
}
