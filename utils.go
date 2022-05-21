package govalidator

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

func contains(str, substring string) bool {
	return strings.Contains(str, substring)
}

func matches(str, pattern string) bool {
	match, _ := regexp.MatchString(pattern, str)
	return match
}

func getTypeName(typ reflect.Type) string {
	pkgPath := typ.PkgPath()
	if pkgPath != "" {
		pkgPath += "."
	}
	return pkgPath + typ.Name()
}

func IsEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String, reflect.Array:
		return v.Len() == 0
	case reflect.Map, reflect.Slice:
		return v.Len() == 0 || v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}

	return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}

func split(str string, sep string) (r []string) {
	p := strings.Split(str, sep)
	if len(p) == 1 && p[0] == "" {
		r = []string{}
		return
	}

	r = make([]string, 0, len(p))
	for _, v := range p {
		v = strings.TrimSpace(v)
		if len(v) > 0 {
			r = append(r, v)
		}
	}
	return
}

// IsZeroValue return true if the value is zero value
func IsZeroValue(value reflect.Value) bool {
	if !value.IsValid() {
		return true
	}

	typ := value.Type()
	if !typ.Comparable() {
		panic(fmt.Errorf("type is not comparable: %v", typ))
	}
	return reflect.Zero(typ).Interface() == value.Interface()
}
