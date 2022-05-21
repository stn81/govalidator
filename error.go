package govalidator

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// Errors is an array of multiple errors and conforms to the error interface.
type Errors []*Error

func (es *Errors) Append(err error, path string) {
	switch v := err.(type) {
	case Errors:
		for _, e := range v {
			e.Name = path + "." + e.Name
		}
		*es = append(*es, v...)
	case *Error:
		v.Name = path + "." + v.Name
		*es = append(*es, v)
	default:
		e := &Error{
			Name: path,
			Err:  v,
		}
		*es = append(*es, e)
	}
}

func (es Errors) Error() string {
	var errs []string
	for _, e := range es {
		errs = append(errs, e.Error())
	}
	return strings.Join(errs, ";")
}

func (es *Errors) Empty() bool {
	return len(*es) == 0
}

func (es *Errors) FindByName(name string) *Error {
	for _, err := range *es {
		if err.Name == name {
			return err
		}
	}
	return nil
}

// Error encapsulates a name, an error and whether there's a custom error message or not.
type Error struct {
	Name string
	Err  error
}

func (e Error) Error() string {
	if e.Name == "" {
		return e.Err.Error()
	}
	return e.Name + ": " + e.Err.Error()
}

func ErrNotExpectedType(got, expected string) error {
	return fmt.Errorf("expected type `%v`, got `%v`", expected, got)
}

func ErrUnsupportedType(typ reflect.Type) error {
	return fmt.Errorf("unsupported type: `%v`", getTypeName(typ))
}

func ErrInvalidTag(tag string, err error) error {
	return fmt.Errorf("invalid tag: %v, %v", tag, err)
}

func ErrFuncInvalid(name string) error {
	return fmt.Errorf("func invalid: %v", name)
}

func ErrUnknownTagValidator(name string) error {
	return fmt.Errorf("unknown tag validator: %v", name)
}

func ErrNotIndirectType(typ reflect.Type) error {
	return fmt.Errorf("expected type(ptr,slice,array,map), but got %v", typ)
}

var (
	ErrUnmatchedParenthesis = errors.New("unmatched parenthesis")
)
