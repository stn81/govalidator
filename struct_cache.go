package govalidator

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/stn81/dynamic"
)

var structValidators = structValidatorCache{}

type structValidatorCache struct {
	store sync.Map
}

func (c *structValidatorCache) get(ptr interface{}) Validator {
	val := reflect.ValueOf(ptr)
	if val.Kind() != reflect.Ptr {
		panic(fmt.Errorf("not struct ptr: %v", getTypeName(reflect.TypeOf(ptr))))
	}

	typ := val.Type()
	kind := typ.Kind()
	if kind != reflect.Ptr && kind != reflect.Struct {
		panic(fmt.Errorf("not struct or struct ptr: %v", getTypeName(reflect.TypeOf(ptr))))
	}

	var validator Validator
	switch typ.Kind() {
	case reflect.Struct:
		validator = c.register(typ)
	case reflect.Ptr:
		elemType := typ.Elem()
		validator = c.register(elemType)
		validator = &DiveValidator{validator}
	}

	return validator
}

func (c *structValidatorCache) register(typ reflect.Type) Validator {
	if typ.Kind() != reflect.Struct {
		panic(fmt.Errorf("not struct type: %v", typ))
	}

	// check duplicated registration
	value, ok := c.store.Load(typ)
	if ok {
		return value.(Validator)
	}

	// parse struct
	numFields := typ.NumField()
	fields := make([]*field, 0, numFields)
	for i := 0; i < numFields; i++ {
		structField := typ.Field(i)

		// skip private field
		if structField.PkgPath != "" {
			continue
		}

		validTag := structField.Tag.Get(DefaultTag)
		if validTag == "-" {
			continue
		}

		validators := []Validator{}
		diveCount := 0

		// collect Tag Validator
		if validTag != "" {
			tags := split(validTag, DefaultTagValueSep)
			for _, tag := range tags {
				if tag == "dive" {
					diveCount += 1
					continue
				}
				tagValidator := c.parseTagValidator(tag)
				for i := 0; i < diveCount; i++ {
					tagValidator = &DiveValidator{tagValidator}
				}
				validators = append(validators, tagValidator)
			}
		}

		// collect struct SelfValidator
		selfValidator := c.parseSelfValidator(structField.Type)
		if selfValidator != nil {
			validators = append(validators, selfValidator)
		}

		// check dynamic field
		if structField.Type.Kind() == reflect.Ptr {
			if dynamic.IsDynamic(structField.Type) {
				validator := &DynamicFieldValidator{}
				validators = append(validators, validator)
			}
		}

		fi := &field{
			index:      i,
			name:       getFieldName(structField),
			validators: validators,
		}

		fields = append(fields, fi)
	}

	stValidator := &structValidator{
		typ:    typ,
		fields: fields,
	}
	c.store.Store(typ, stValidator)
	return stValidator
}

func (c *structValidatorCache) parseTagValidator(tag string) Validator {
	var name string
	var args []string
	var customErr error

	pCustomErr := strings.Index(tag, "~")
	if pCustomErr != -1 {
		customErr = errors.New(tag[pCustomErr+1:])
		tag = tag[:pCustomErr]
	}

	pStart := strings.Index(tag, "(")
	if pStart == -1 {
		name = tag
	} else {
		pEnd := strings.Index(tag, ")")
		if pEnd == -1 {
			panic(ErrUnmatchedParenthesis)
		}
		name = tag[:pStart]
		args = split(tag[pStart+1:pEnd], ",")
	}

	validator := TagValidatorMap.Get(name)
	if validator == nil {
		panic(ErrUnknownTagValidator(name))
	}

	tagValidator := &TagValidator{
		Name:      name,
		CustomErr: customErr,
		Validator: validator,
		Args:      args,
	}
	return tagValidator
}

func (c *structValidatorCache) parseSelfValidator(typ reflect.Type) Validator {
	switch typ.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Array, reflect.Map:
		elemType := typ.Elem()
		validator := c.parseSelfValidator(elemType)
		if validator != nil {
			return &DiveValidator{validator}
		}
	case reflect.Struct:
		return c.register(typ)
	}

	return nil
}
