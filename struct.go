package govalidator

import (
	"fmt"
	"reflect"

	"github.com/stn81/dynamic"
)

type TagValidator struct {
	Name      string
	CustomErr error
	Validator Validator
	Args      []string
}

func (v *TagValidator) Validate(value interface{}, args ...string) error {
	err := v.Validator.Validate(value, v.Args...)
	if err != nil {
		if v.CustomErr != nil {
			err = v.CustomErr
		}
		return err
	}
	return nil
}

type DiveValidator struct {
	Validator Validator
}

func (v *DiveValidator) Validate(value interface{}, args ...string) error {
	val := reflect.ValueOf(value)

	switch val.Kind() {
	case reflect.Ptr:
		if val.IsNil() {
			return nil
		}
		ind := reflect.Indirect(val)
		if err := v.Validator.Validate(ind.Interface()); err != nil {
			return err
		}
	case reflect.Slice, reflect.Array:
		size := val.Len()
		for i := 0; i < size; i++ {
			ind := val.Index(i)
			if err := v.Validator.Validate(ind.Interface()); err != nil {
				return err
			}
		}
	case reflect.Map:
		keys := val.MapKeys()
		for _, key := range keys {
			ind := val.MapIndex(key)
			if err := v.Validator.Validate(ind.Interface()); err != nil {
				return err
			}
		}
	default:
		panic(ErrNotIndirectType(val.Type()))
	}

	// check struct
	if validator, ok := value.(SelfValidator); ok {
		if err := validator.Validate(); err != nil {
			return err
		}
	}

	return nil
}

type DynamicFieldValidator struct{}

func (v *DynamicFieldValidator) Validate(value interface{}, args ...string) error {
	val := reflect.ValueOf(dynamic.GetValue(value.(*dynamic.Type)))
	if !val.IsValid() {
		return nil
	}
	typ := val.Type()
	validator := structValidators.parseSelfValidator(typ)
	if validator == nil {
		return nil
	}
	return validator.Validate(val.Interface())
}

type field struct {
	index      int
	name       string
	validators []Validator
}

type structValidator struct {
	typ    reflect.Type
	fields []*field
}

func (v *structValidator) Validate(value interface{}, args ...string) error {
	val := reflect.ValueOf(value)

	if val.Kind() != reflect.Struct {
		panic(fmt.Errorf("not struct type: %v", val.Type()))
	}

	// check each fields
	var errs Errors
	for _, field := range v.fields {
		fieldVal := val.Field(field.index)
	validatorsLoop:
		for _, validator := range field.validators {
			err := validator.Validate(fieldVal.Interface())
			switch {
			case err == ErrSkip:
				break validatorsLoop
			case err != nil:
				errs.Append(err, field.name)
			}
		}
	}
	if !errs.Empty() {
		return errs
	}

	return nil
}
