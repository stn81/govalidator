package govalidator

import (
	"sync"
)

var TagValidatorMap = &tagValidatorMap{}

type tagValidatorMap struct {
	store sync.Map
}

func (m *tagValidatorMap) Get(name string) Validator {
	value, ok := m.store.Load(name)
	if !ok {
		return nil
	}
	return value.(Validator)
}

func (m *tagValidatorMap) RegisterValidator(name string, validator Validator) {
	m.store.Store(name, validator)
}

func (m *tagValidatorMap) RegisterValidateFunc(name string, validateFunc ValidateFunc) {
	m.RegisterValidator(name, validateFunc)
}
