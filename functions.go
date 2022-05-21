package govalidator

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

var (
	notNumberRegexp     = regexp.MustCompile("[^0-9]+")
	whiteSpacesAndMinus = regexp.MustCompile(`[\s-]+`)
)

const (
	maxURLRuneCount   = 2083
	minURLRuneCount   = 3
	RF3339WithoutZone = "2006-01-02T15:04:05"
)

var (
	// ErrSkip skip the after tag validators
	ErrSkip                       = errors.New("skip")
	ErrNotString                  = errors.New("not string")
	ErrInvalidEmail               = errors.New("invalid email")
	ErrInvalidURL                 = errors.New("invalid url")
	ErrInvalidAlpha               = errors.New("invalid alpha")
	ErrInvalidUTFLetter           = errors.New("invalid UTF Letter")
	ErrInvalidAlphanumeric        = errors.New("invalid alpha or numeric")
	ErrInvalidNumeric             = errors.New("invalid numeric")
	ErrNotLowerCase               = errors.New("not lower case")
	ErrNotUpperCase               = errors.New("not upper case")
	ErrNotHasLowerCase            = errors.New("not has lower case")
	ErrNotHasUpperCase            = errors.New("not has upper case")
	ErrNotInteger                 = errors.New("not integer")
	ErrNotFloat                   = errors.New("not float")
	ErrNotEmpty                   = errors.New("not empty")
	ErrIsRequired                 = errors.New("is required")
	ErrInvalidJSON                = errors.New("invalid json")
	ErrInvalidASCII               = errors.New("invalid ascii")
	ErrNotPrintableASCII          = errors.New("not printable ascii")
	ErrInvalidBase64              = errors.New("invalid base64")
	ErrInvalidIP                  = errors.New("invalid IP")
	ErrInvalidPort                = errors.New("invalid port")
	ErrInvalidIPv4                = errors.New("invalid IPv4")
	ErrInvalidIPv6                = errors.New("invalid IPv6")
	ErrInvalidCIDR                = errors.New("invalid CIDR")
	ErrInvalidMAC                 = errors.New("invalid MAC")
	ErrInvalidLatitude            = errors.New("invalid latitude")
	ErrInvalidLongtitude          = errors.New("invalid longtitude")
	ErrInvalidISO4217CurrencyCode = errors.New("invalid ISO4217 currency code")
)

func ErrNotInList(value interface{}, args ...string) error {
	return fmt.Errorf("%v not in list: [%v]", value, strings.Join(args, ","))
}

func ErrRegexpNotMatch(value, pattern string) error {
	return fmt.Errorf("%v not match pattern: %v", value, pattern)
}

func ErrNumArgsInvalid(funcName string, expected int) error {
	return fmt.Errorf("function %v need %v arguments", funcName, expected)
}

func ErrInvalidLength(got, min, max int) error {
	return fmt.Errorf("length should in range [%v, %v], but got %v", min, max, got)
}

func ErrNotInRange(value interface{}, min, max string) error {
	return fmt.Errorf("should in range [%v, %v], but got %v", min, max, value)
}

func ErrLessThanMin(value interface{}, min interface{}) error {
	return fmt.Errorf("should be great than %v, but got %v", min, value)
}

func ErrGreatThanMax(value interface{}, max interface{}) error {
	return fmt.Errorf("should be less than %v, but got %v", max, value)
}

func ErrInvalidTime(str, format string) error {
	return fmt.Errorf("invalid time: %v, format should be: %v", str, format)
}

func ErrInvalidHash(method string, value string) error {
	return fmt.Errorf("invalid %v hash: %v", method, value)
}

func assertString(value interface{}) string {
	val := reflect.ValueOf(value)
	if val.Kind() != reflect.String {
		panic(ErrNotString)
	}
	return val.String()
}

// IsEmail check if the string is an email.
func IsEmail(value interface{}, args ...string) error {
	str := assertString(value)
	if str == "" {
		return nil
	}

	if rxEmail.MatchString(str) {
		return nil
	}
	return ErrInvalidEmail
}

// IsURL check if the string is an URL.
func IsURL(value interface{}, args ...string) error {
	str := assertString(value)
	if str == "" {
		return nil
	}

	if utf8.RuneCountInString(str) >= maxURLRuneCount || len(str) <= minURLRuneCount || strings.HasPrefix(str, ".") {
		return ErrInvalidURL
	}
	strTemp := str
	if strings.Contains(str, ":") && !strings.Contains(str, "://") {
		// support no indicated urlscheme but with colon for port number
		// http:// is appended so url.Parse will succeed, strTemp used so it does not impact rxURL.MatchString
		strTemp = "http://" + str
	}
	u, err := url.Parse(strTemp)
	if err != nil {
		return ErrInvalidURL
	}
	if strings.HasPrefix(u.Host, ".") {
		return ErrInvalidURL
	}
	if u.Host == "" && (u.Path != "" && !strings.Contains(u.Path, ".")) {
		return ErrInvalidURL
	}

	if rxURL.MatchString(str) {
		return nil
	}
	return ErrInvalidURL
}

// IsAlpha check if the string contains only letters (a-zA-Z). Empty string is valid.
func IsAlpha(value interface{}, args ...string) error {
	str := assertString(value)
	if str == "" {
		return nil
	}
	if rxAlpha.MatchString(str) {
		return nil
	}
	return ErrInvalidAlpha
}

// IsAlphanumeric check if the string contains only letters and numbers. Empty string is valid.
func IsAlphanumeric(value interface{}, args ...string) error {
	str := assertString(value)
	if str == "" {
		return nil
	}
	if rxAlphanumeric.MatchString(str) {
		return nil
	}
	return ErrInvalidAlphanumeric
}

// IsNumeric check if the string contains only numbers. Empty string is valid.
func IsNumeric(value interface{}, args ...string) error {
	str := assertString(value)
	if str == "" {
		return nil
	}
	if rxNumeric.MatchString(str) {
		return nil
	}
	return ErrInvalidNumeric
}

// IsLowerCase check if the string is lowercase. Empty string is valid.
func IsLowerCase(value interface{}, args ...string) error {
	str := assertString(value)
	if str == "" {
		return nil
	}
	if str == strings.ToLower(str) {
		return nil
	}
	return ErrNotLowerCase
}

// IsUpperCase check if the string is uppercase. Empty string is valid.
func IsUpperCase(value interface{}, args ...string) error {
	str := assertString(value)
	if str == "" {
		return nil
	}
	if str == strings.ToUpper(str) {
		return nil
	}
	return ErrNotUpperCase
}

// IsInt check if the string is an integer. Empty string is valid.
func IsInt(value interface{}, args ...string) error {
	str := assertString(value)
	if str == "" {
		return nil
	}
	if rxInt.MatchString(str) {
		return nil
	}
	return ErrNotInteger
}

// IsFloat check if the string is a float.
func IsFloat(value interface{}, args ...string) error {
	str := assertString(value)
	if str == "" {
		return nil
	}

	if rxFloat.MatchString(str) {
		return nil
	}
	return ErrNotFloat
}

// IsEmpty check if the string is null.
func IsEmpty(value interface{}, args ...string) error {
	val := reflect.ValueOf(value)
	if IsEmptyValue(val) {
		return nil
	}
	return ErrNotEmpty
}

// Length check if the string's length (in bytes) falls in a range.
func Length(value interface{}, args ...string) error {
	if len(args) != 2 {
		panic(ErrNumArgsInvalid("length", 2))
	}

	var length int
	switch v := value.(type) {
	case string:
		length = utf8.RuneCountInString(v)
	case []byte:
		length = len(v)
	}

	min := GetInt(args[0])
	max := GetInt(args[1])
	if length >= min && length <= max {
		return nil
	}
	return ErrInvalidLength(length, min, max)
}

// IsJSON check if the string is valid JSON (note: uses json.Unmarshal).
func IsJSON(value interface{}, args ...string) error {
	str := assertString(value)
	if str == "" {
		return nil
	}

	if json.Valid([]byte(str)) {
		return nil
	}
	return ErrInvalidJSON
}

// IsASCII check if the string contains ASCII chars only. Empty string is valid.
func IsASCII(value interface{}, args ...string) error {
	str := assertString(value)
	if str == "" {
		return nil
	}
	if rxASCII.MatchString(str) {
		return nil
	}
	return ErrInvalidASCII
}

// IsPrintableASCII check if the string contains printable ASCII chars only. Empty string is valid.
func IsPrintableASCII(value interface{}, args ...string) error {
	str := assertString(value)
	if str == "" {
		return nil
	}
	if rxPrintableASCII.MatchString(str) {
		return nil
	}
	return ErrNotPrintableASCII
}

// IsBase64 check if a string is base64 encoded.
func IsBase64(value interface{}, args ...string) error {
	str := assertString(value)
	if str == "" {
		return nil
	}

	if rxBase64.MatchString(str) {
		return nil
	}
	return ErrInvalidBase64
}

// IsHash checks if a string is a hash of type algorithm.
// Algorithm is one of ['md4', 'md5', 'sha1', 'sha256', 'sha384', 'sha512', 'ripemd128', 'ripemd160', 'tiger128', 'tiger160', 'tiger192', 'crc32', 'crc32b']
func IsHash(value interface{}, args ...string) error {
	if len(args) != 1 {
		panic(ErrNumArgsInvalid("hash", 1))
	}

	str := assertString(value)
	if str == "" {
		return nil
	}

	algorithm := GetString(args[0])
	length := "0"
	algo := strings.ToLower(algorithm)

	if algo == "crc32" || algo == "crc32b" {
		length = "8"
	} else if algo == "md5" || algo == "md4" || algo == "ripemd128" || algo == "tiger128" {
		length = "32"
	} else if algo == "sha1" || algo == "ripemd160" || algo == "tiger160" {
		length = "40"
	} else if algo == "tiger192" {
		length = "48"
	} else if algo == "sha256" {
		length = "64"
	} else if algo == "sha384" {
		length = "96"
	} else if algo == "sha512" {
		length = "128"
	} else {
		return ErrInvalidHash(algo, str)
	}

	if matches(str, "^[a-f0-9]{"+length+"}$") {
		return nil
	}
	return ErrInvalidHash(algo, str)
}

// IsIP checks if a string is either IP version 4 or 6.
func IsIP(value interface{}, args ...string) error {
	str := assertString(value)
	if str == "" {
		return nil
	}

	if net.ParseIP(str) != nil {
		return nil
	}
	return ErrInvalidIP
}

// IsPort checks if a string represents a valid port
func IsPort(value interface{}, args ...string) error {
	str := assertString(value)
	if str == "" {
		return nil
	}

	if i, err := strconv.Atoi(str); err == nil && i > 0 && i < 65536 {
		return nil
	}
	return ErrInvalidPort
}

// IsIPv4 check if the string is an IP version 4.
func IsIPv4(value interface{}, args ...string) error {
	str := assertString(value)
	if str == "" {
		return nil
	}

	ip := net.ParseIP(str)
	if ip != nil && strings.Contains(str, ".") {
		return nil
	}
	return ErrInvalidIPv4
}

// IsIPv6 check if the string is an IP version 6.
func IsIPv6(value interface{}, args ...string) error {
	str := assertString(value)
	if str == "" {
		return nil
	}

	ip := net.ParseIP(str)
	if ip != nil && strings.Contains(str, ":") {
		return nil
	}
	return ErrInvalidIPv6
}

// IsCIDR check if the string is an valid CIDR notiation (IPV4 & IPV6)
func IsCIDR(value interface{}, args ...string) error {
	str := assertString(value)
	if str == "" {
		return nil
	}

	if _, _, err := net.ParseCIDR(str); err == nil {
		return nil
	}
	return ErrInvalidCIDR
}

// IsMAC check if a string is valid MAC address.
// Possible MAC formats:
// 01:23:45:67:89:ab
// 01:23:45:67:89:ab:cd:ef
// 01-23-45-67-89-ab
// 01-23-45-67-89-ab-cd-ef
// 0123.4567.89ab
// 0123.4567.89ab.cdef
func IsMAC(value interface{}, args ...string) error {
	str := assertString(value)
	if str == "" {
		return nil
	}

	if _, err := net.ParseMAC(str); err == nil {
		return nil
	}
	return ErrInvalidMAC
}

// IsLatitude check if a string is valid latitude.
func IsLatitude(value interface{}, args ...string) error {
	str := assertString(value)
	if rxLatitude.MatchString(str) {
		return nil
	}
	return ErrInvalidLatitude
}

// IsLongitude check if a string is valid longitude.
func IsLongitude(value interface{}, args ...string) error {
	str := assertString(value)
	if str == "" {
		return nil
	}

	if rxLongitude.MatchString(str) {
		return nil
	}
	return ErrInvalidLongtitude
}

// IsTime check if string is valid according to given format
func IsTime(value interface{}, args ...string) error {
	if len(args) != 1 {
		panic(ErrNumArgsInvalid("time", 1))
	}

	str := assertString(value)
	if str == "" {
		return nil
	}

	format := GetString(args[0])
	if _, err := time.Parse(format, str); err == nil {
		return nil
	}
	return ErrInvalidTime(str, format)
}

// IsRFC3339 check if string is valid timestamp value according to RFC3339
func IsRFC3339(value interface{}, args ...string) error {
	return IsTime(value, time.RFC3339)
}

// IsRFC3339WithoutZone check if string is valid timestamp value according to RFC3339 which excludes the timezone.
func IsRFC3339WithoutZone(value interface{}, args ...string) error {
	return IsTime(value, RF3339WithoutZone)
}

// IsISO4217 check if string is valid ISO currency code
func IsISO4217(value interface{}, args ...string) error {
	str := assertString(value)
	if str == "" {
		return nil
	}

	for _, currency := range ISO4217List {
		if str == currency {
			return nil
		}
	}

	return ErrInvalidISO4217CurrencyCode
}

// RegEx checks if a string matches a given pattern.
func RegEx(value interface{}, args ...string) error {
	if len(args) != 1 {
		panic(ErrNumArgsInvalid("regex", 1))
	}

	str := assertString(value)
	if str == "" {
		return nil
	}

	pattern := args[0]
	if matches(str, pattern) {
		return nil
	}
	return ErrRegexpNotMatch(str, pattern)
}

// Min check the min value
func Min(value interface{}, args ...string) error {
	if len(args) != 1 {
		panic(ErrNumArgsInvalid("min", 1))
	}

	switch value.(type) {
	case int, int8, int16, int32, int64:
		val := GetInt64(value)
		min := GetInt64(args[0])
		if val >= min {
			return nil
		}
	case uint, uint8, uint16, uint32, uint64:
		val := GetUint64(value)
		min := GetUint64(args[0])
		if val >= min {
			return nil
		}
	case float32, float64:
		val := GetFloat64(value)
		min := GetFloat64(args[0])
		if val >= min {
			return nil
		}
	}

	return ErrLessThanMin(value, args[0])
}

// Max check the max value
func Max(value interface{}, args ...string) error {
	if len(args) != 1 {
		panic(ErrNumArgsInvalid("max", 1))
	}

	switch value.(type) {
	case int, int8, int16, int32, int64:
		val := GetInt64(value)
		max := GetInt64(args[0])
		if val <= max {
			return nil
		}
	case uint, uint8, uint16, uint32, uint64:
		val := GetUint64(value)
		max := GetUint64(args[0])
		if val <= max {
			return nil
		}
	case float32, float64:
		val := GetFloat64(value)
		max := GetFloat64(args[0])
		if val <= max {
			return nil
		}
	}

	return ErrGreatThanMax(value, args[0])
}

// Range check value range
func Range(value interface{}, args ...string) error {
	if len(args) != 2 {
		panic(ErrNumArgsInvalid("range", 2))
	}

	switch value.(type) {
	case int, int8, int16, int32, int64:
		val := GetInt64(value)
		min := GetInt64(args[0])
		max := GetInt64(args[1])
		if val >= min && val <= max {
			return nil
		}
	case uint, uint8, uint16, uint32, uint64:
		val := GetUint64(value)
		min := GetUint64(args[0])
		max := GetUint64(args[1])
		if val >= min && val <= max {
			return nil
		}
	case float32, float64:
		val := GetFloat64(value)
		min := GetFloat64(args[0])
		max := GetFloat64(args[1])
		if val >= min && val <= max {
			return nil
		}
	}

	return ErrNotInRange(value, args[0], args[1])
}

// IsIn check if string str is a member of the set of strings params
func IsIn(value interface{}, args ...string) error {
	str := GetString(value)
	if str == "" {
		return nil
	}
	for _, arg := range args {
		if str == arg {
			return nil
		}
	}

	return ErrNotInList(value, args...)
}

// Required check where value is not empty value
func Required(value interface{}, args ...string) error {
	val := reflect.ValueOf(value)
	if !IsEmptyValue(val) {
		return nil
	}
	return ErrIsRequired
}

// SkipEmpty validator
func SkipEmpty(value interface{}, args ...string) error {
	val := reflect.ValueOf(value)
	if !IsEmptyValue(val) {
		return nil
	}
	return ErrSkip
}

// TagMap is a map of functions, that can be used as tags for ValidateStruct function.
var TagMap = map[string]ValidateFunc{
	"email":              IsEmail,
	"url":                IsURL,
	"alpha":              IsAlpha,
	"alphanum":           IsAlphanumeric,
	"numeric":            IsNumeric,
	"lowercase":          IsLowerCase,
	"uppercase":          IsUpperCase,
	"int":                IsInt,
	"float":              IsFloat,
	"empty":              IsEmpty,
	"json":               IsJSON,
	"ascii":              IsASCII,
	"hash":               IsHash,
	"printableascii":     IsPrintableASCII,
	"base64":             IsBase64,
	"ip":                 IsIP,
	"port":               IsPort,
	"ipv4":               IsIPv4,
	"ipv6":               IsIPv6,
	"mac":                IsMAC,
	"latitude":           IsLatitude,
	"longitude":          IsLongitude,
	"rfc3339":            IsRFC3339,
	"rfc3339WithoutZone": IsRFC3339WithoutZone,
	"ISO4217":            IsISO4217,
	"required":           Required,
	"in":                 IsIn,
	"min":                Min,
	"max":                Max,
	"range":              Range,
	"length":             Length,
	"skipempty":          SkipEmpty,
	"regex":              RegEx,
}

func init() {
	for tag, validator := range TagMap {
		TagValidatorMap.RegisterValidateFunc(tag, validator)
	}
}
