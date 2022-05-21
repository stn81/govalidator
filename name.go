package govalidator

import (
	"reflect"
	"strings"
)

func toJSONName(tag string) string {
	if tag == "" {
		return ""
	}

	// JSON name always comes first. If there's no options then split[0] is
	// JSON name, if JSON name is not set, then split[0] is an empty string.
	split := strings.SplitN(tag, ",", 2)

	name := split[0]

	// However it is possible that the field is skipped when
	// (de-)serializing from/to JSON, in which case assume that there is no
	// tag name to use
	if name == "-" {
		return ""
	}
	return name
}

func getFieldName(field reflect.StructField) string {
	if name := field.Tag.Get("query"); name != "" {
		return name
	}

	if name := field.Tag.Get("rest"); name != "" {
		return name
	}

	if name := field.Tag.Get("json"); name != "" {
		return toJSONName(name)
	}
	return field.Name
}
