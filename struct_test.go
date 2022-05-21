package govalidator

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stn81/dynamic"
	"github.com/stretchr/testify/require"
)

func TestTagValidator(t *testing.T) {
	emptySt := &struct {
		Name string `valid:"required"`
	}{}
	err := ValidateStruct(emptySt)
	require.NotNil(t, err)

	errs := err.(Errors)
	nameErr := errs.FindByName("Name")
	require.NotNil(t, nameErr)
	require.Equal(t, ErrIsRequired, nameErr.Err)

	alphaSt := &struct {
		Name string `valid:"alpha"`
	}{
		"234",
	}
	err = ValidateStruct(alphaSt)
	require.NotNil(t, err)

	errs = err.(Errors)
	nameErr = errs.FindByName("Name")
	require.NotNil(t, nameErr)
	require.Equal(t, ErrInvalidAlpha, nameErr.Err)
}

type aType struct {
	Name  string `valid:"required"`
	Value int    `valid:"range(1,3)"`
}

func (a *aType) Validate() error {
	if a.Value == 3 {
		return errors.New("aType validate failed")
	}
	return nil
}

func TestTypeValidator(t *testing.T) {
	a := &aType{Name: "zhangsan", Value: 3}
	err := ValidateStruct(a)
	require.Error(t, err)
	require.Contains(t, err.Error(), "aType validate failed")

	b := &aType{Name: "zhangsan", Value: 4}
	err = ValidateStruct(b)
	require.Error(t, err)
}

func TestSkipValidator(t *testing.T) {
	st := &struct {
		Name string `valid:"skipempty;alpha"`
	}{
		"234",
	}
	err := ValidateStruct(st)
	require.Error(t, err)

	st2 := &struct {
		Name string `valid:"skipempty;alpha"`
	}{}
	err2 := ValidateStruct(st2)
	require.NoError(t, err2)
}

type stEmbeded struct {
	ID int `json:"id" valid:"required"`
}

type stOutter struct {
	Objs []*stEmbeded `valid:"required"`
}

func TestStructSliceFields(t *testing.T) {
	st := &stOutter{
		Objs: []*stEmbeded{
			{0},
		},
	}

	err := ValidateStruct(st)
	require.Error(t, err)
}

type goodDiveStruct struct {
	Name *string `valid:"dive;alpha"`
}

type badDiveStruct struct {
	Name *string `valid:"alpha"`
}

func TestStructDiveField(t *testing.T) {
	strData := "abc"
	st := &goodDiveStruct{Name: &strData}
	err := ValidateStruct(st)
	require.NoError(t, err)

	st2 := &badDiveStruct{Name: &strData}
	defer func() {
		r := recover()
		require.Contains(t, fmt.Sprint(r), "not string")
	}()
	ValidateStruct(st2)
}

type stNilDiveStruct struct {
	Values []*stEmbeded `valid:"required"`
}

func TestNilDive(t *testing.T) {
	st := &stNilDiveStruct{
		Values: []*stEmbeded{nil, {1}},
	}
	err := ValidateStruct(st)
	require.NoError(t, err)

	st2 := &stNilDiveStruct{
		Values: []*stEmbeded{nil, {0}},
	}
	err2 := ValidateStruct(st2)
	require.Error(t, err2)
	require.Contains(t, err2.Error(), "is required")
}

type stDynamicField struct {
	Content *dynamic.Type `json:"data"`
}

func TestDynamicField(t *testing.T) {
	st := &stDynamicField{
		Content: dynamic.New(stEmbeded{0}),
	}
	err := ValidateStruct(st)
	require.Error(t, err)
	require.Contains(t, err.Error(), "data.id")

	st2 := &stDynamicField{
		Content: dynamic.New(nil),
	}
	err2 := ValidateStruct(st2)
	require.NoError(t, err2)
}
