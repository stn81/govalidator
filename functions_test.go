package govalidator

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestIsAlpha(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"\n", false},
		{"\r", false},
		{"Ⅸ", false},
		{"", true},
		{"   fooo   ", false},
		{"abc!!!", false},
		{"abc1", false},
		{"abc〩", false},
		{"abc", true},
		{"소주", false},
		{"ABC", true},
		{"FoObAr", true},
		{"소aBC", false},
		{"소", false},
		{"달기&Co.", false},
		{"〩Hours", false},
		{"\ufff0", false},
		{"\u0070", true},  //UTF-8(ASCII): p
		{"\u0026", false}, //UTF-8(ASCII): &
		{"\u0030", false}, //UTF-8(ASCII): 0
		{"123", false},
		{"0123", false},
		{"-00123", false},
		{"0", false},
		{"-0", false},
		{"123.123", false},
		{" ", false},
		{".", false},
		{"-1¾", false},
		{"1¾", false},
		{"〥〩", false},
		{"모자", false},
		{"ix", true},
		{"۳۵۶۰", false},
		{"1--", false},
		{"1-1", false},
		{"-", false},
		{"--", false},
		{"1++", false},
		{"1+1", false},
		{"+", false},
		{"++", false},
		{"+1", false},
	}
	for _, test := range tests {
		err := IsAlpha(test.param)
		if test.expected {
			require.NoError(t, err, "check IsAlpha(%s)", test.param)
		} else {
			require.Error(t, err, "check IsAlpha(%s)", test.param)
		}
	}
}

func TestIsAlphanumeric(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"\n", false},
		{"\r", false},
		{"Ⅸ", false},
		{"", true},
		{"   fooo   ", false},
		{"abc!!!", false},
		{"abc123", true},
		{"ABC111", true},
		{"abc1", true},
		{"abc〩", false},
		{"abc", true},
		{"소주", false},
		{"ABC", true},
		{"FoObAr", true},
		{"소aBC", false},
		{"소", false},
		{"달기&Co.", false},
		{"〩Hours", false},
		{"\ufff0", false},
		{"\u0070", true},  //UTF-8(ASCII): p
		{"\u0026", false}, //UTF-8(ASCII): &
		{"\u0030", true},  //UTF-8(ASCII): 0
		{"123", true},
		{"0123", true},
		{"-00123", false},
		{"0", true},
		{"-0", false},
		{"123.123", false},
		{" ", false},
		{".", false},
		{"-1¾", false},
		{"1¾", false},
		{"〥〩", false},
		{"모자", false},
		{"ix", true},
		{"۳۵۶۰", false},
		{"1--", false},
		{"1-1", false},
		{"-", false},
		{"--", false},
		{"1++", false},
		{"1+1", false},
		{"+", false},
		{"++", false},
		{"+1", false},
	}
	for _, test := range tests {
		err := IsAlphanumeric(test.param)
		if test.expected {
			require.NoError(t, err, "check IsAlphanumeric(%s)", test.param)
		} else {
			require.Error(t, err, "check IsAlphanumeric(%s)", test.param)
		}
	}
}

func TestIsNumeric(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"\n", false},
		{"\r", false},
		{"Ⅸ", false},
		{"", true},
		{"   fooo   ", false},
		{"abc!!!", false},
		{"abc1", false},
		{"abc〩", false},
		{"abc", false},
		{"소주", false},
		{"ABC", false},
		{"FoObAr", false},
		{"소aBC", false},
		{"소", false},
		{"달기&Co.", false},
		{"〩Hours", false},
		{"\ufff0", false},
		{"\u0070", false}, //UTF-8(ASCII): p
		{"\u0026", false}, //UTF-8(ASCII): &
		{"\u0030", true},  //UTF-8(ASCII): 0
		{"123", true},
		{"0123", true},
		{"-00123", false},
		{"+00123", false},
		{"0", true},
		{"-0", false},
		{"123.123", false},
		{" ", false},
		{".", false},
		{"12𐅪3", false},
		{"-1¾", false},
		{"1¾", false},
		{"〥〩", false},
		{"모자", false},
		{"ix", false},
		{"۳۵۶۰", false},
		{"1--", false},
		{"1-1", false},
		{"-", false},
		{"--", false},
		{"1++", false},
		{"1+1", false},
		{"+", false},
		{"++", false},
		{"+1", false},
	}
	for _, test := range tests {
		err := IsNumeric(test.param)
		if test.expected {
			require.NoError(t, err, "check IsNumeric(%s)", test.param)
		} else {
			require.Error(t, err, "check IsNumeric(%s)", test.param)
		}
	}
}

func TestIsLowerCase(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", true},
		{"abc123", true},
		{"abc", true},
		{"a b c", true},
		{"abcß", true},
		{"abcẞ", false},
		{"ABCẞ", false},
		{"tr竪s 端ber", true},
		{"fooBar", false},
		{"123ABC", false},
		{"ABC123", false},
		{"ABC", false},
		{"S T R", false},
		{"fooBar", false},
		{"abacaba123", true},
	}
	for _, test := range tests {
		err := IsLowerCase(test.param)
		if test.expected {
			require.NoError(t, err, "check IsLowerCase(%s)", test.param)
		} else {
			require.Error(t, err, "check IsLowerCase(%s)", test.param)
		}
	}
}

func TestIsUpperCase(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", true},
		{"abc123", false},
		{"abc", false},
		{"a b c", false},
		{"abcß", false},
		{"abcẞ", false},
		{"ABCẞ", true},
		{"tr竪s 端ber", false},
		{"fooBar", false},
		{"123ABC", true},
		{"ABC123", true},
		{"ABC", true},
		{"S T R", true},
		{"fooBar", false},
		{"abacaba123", false},
	}
	for _, test := range tests {
		err := IsUpperCase(test.param)
		if test.expected {
			require.NoError(t, err, "check IsUpperCase(%s)", test.param)
		} else {
			require.Error(t, err, "check IsUpperCase(%s)", test.param)
		}
	}
}

func TestIsInt(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"-2147483648", true},          //Signed 32 Bit Min Int
		{"2147483647", true},           //Signed 32 Bit Max Int
		{"-2147483649", true},          //Signed 32 Bit Min Int - 1
		{"2147483648", true},           //Signed 32 Bit Max Int + 1
		{"4294967295", true},           //Unsigned 32 Bit Max Int
		{"4294967296", true},           //Unsigned 32 Bit Max Int + 1
		{"-9223372036854775808", true}, //Signed 64 Bit Min Int
		{"9223372036854775807", true},  //Signed 64 Bit Max Int
		{"-9223372036854775809", true}, //Signed 64 Bit Min Int - 1
		{"9223372036854775808", true},  //Signed 64 Bit Max Int + 1
		{"18446744073709551615", true}, //Unsigned 64 Bit Max Int
		{"18446744073709551616", true}, //Unsigned 64 Bit Max Int + 1
		{"", true},
		{"123", true},
		{"0", true},
		{"-0", true},
		{"+0", true},
		{"01", false},
		{"123.123", false},
		{" ", false},
		{"000", false},
	}
	for _, test := range tests {
		err := IsInt(test.param)
		if test.expected {
			require.NoError(t, err, "check IsInt(%s)", test.param)
		} else {
			require.Error(t, err, "check IsInt(%s)", test.param)
		}
	}
}

func TestIsEmail(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", true},
		{"foo@bar.com", true},
		{"x@x.x", true},
		{"foo@bar.com.au", true},
		{"foo+bar@bar.com", true},
		{"foo@bar.coffee", true},
		{"foo@bar.coffee..coffee", false},
		{"foo@bar.bar.coffee", true},
		{"foo@bar.中文网", true},
		{"invalidemail@", false},
		{"invalid.com", false},
		{"@invalid.com", false},
		{"test|123@m端ller.com", true},
		{"hans@m端ller.com", true},
		{"hans.m端ller@test.com", true},
		{"NathAn.daVIeS@DomaIn.cOM", true},
		{"NATHAN.DAVIES@DOMAIN.CO.UK", true},
	}
	for _, test := range tests {
		err := IsEmail(test.param)
		if test.expected {
			require.NoError(t, err, "isEmail(%s)", test.param)
		} else {
			require.Error(t, err, "isEmail(%s)", test.param)
		}
	}
}

func TestIsURL(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", true},
		{"http://foo.bar#com", true},
		{"http://foobar.com", true},
		{"https://foobar.com", true},
		{"foobar.com", true},
		{"http://foobar.coffee/", true},
		{"http://foobar.中文网/", true},
		{"http://foobar.org/", true},
		{"http://foobar.ORG", true},
		{"http://foobar.org:8080/", true},
		{"ftp://foobar.ru/", true},
		{"ftp.foo.bar", true},
		{"http://user:pass@www.foobar.com/", true},
		{"http://user:pass@www.foobar.com/path/file", true},
		{"http://127.0.0.1/", true},
		{"http://duckduckgo.com/?q=%2F", true},
		{"http://localhost:3000/", true},
		{"http://foobar.com/?foo=bar#baz=qux", true},
		{"http://foobar.com?foo=bar", true},
		{"http://www.xn--froschgrn-x9a.net/", true},
		{"http://foobar.com/a-", true},
		{"http://foobar.پاکستان/", true},
		{"http://foobar.c_o_m", false},
		{"http://_foobar.com", false},
		{"http://foo_bar.com", true},
		{"http://user:pass@foo_bar_bar.bar_foo.com", true},
		{"xyz://foobar.com", false},
		// {"invalid.", false}, is it false like "localhost."?
		{".com", false},
		{"rtmp://foobar.com", false},
		{"http://localhost:3000/", true},
		{"http://foobar.com#baz=qux", true},
		{"http://foobar.com/t$-_.+!*\\'(),", true},
		{"http://www.foobar.com/~foobar", true},
		{"http://www.-foobar.com/", false},
		{"http://www.foo---bar.com/", false},
		{"http://r6---snnvoxuioq6.googlevideo.com", true},
		{"mailto:someone@example.com", true},
		{"irc://irc.server.org/channel", false},
		{"irc://#channel@network", true},
		{"/abs/test/dir", false},
		{"./rel/test/dir", false},
		{"http://foo^bar.org", false},
		{"http://foo&*bar.org", false},
		{"http://foo&bar.org", false},
		{"http://foo bar.org", false},
		{"http://foo.bar.org", true},
		{"http://www.foo.bar.org", true},
		{"http://www.foo.co.uk", true},
		{"foo", false},
		{"http://.foo.com", false},
		{"http://,foo.com", false},
		{",foo.com", false},
		{"http://myservice.:9093/", true},
		// according to issues #62 #66
		{"https://pbs.twimg.com/profile_images/560826135676588032/j8fWrmYY_normal.jpeg", true},
		// according to #125
		{"http://prometheus-alertmanager.service.q:9093", true},
		{"aio1_alertmanager_container-63376c45:9093", true},
		{"https://www.logn-123-123.url.with.sigle.letter.d:12345/url/path/foo?bar=zzz#user", true},
		{"http://me.example.com", true},
		{"http://www.me.example.com", true},
		{"https://farm6.static.flickr.com", true},
		{"https://zh.wikipedia.org/wiki/Wikipedia:%E9%A6%96%E9%A1%B5", true},
		{"google", false},
		// According to #87
		{"http://hyphenated-host-name.example.co.in", true},
		{"http://cant-end-with-hyphen-.example.com", false},
		{"http://-cant-start-with-hyphen.example.com", false},
		{"http://www.domain-can-have-dashes.com", true},
		{"http://m.abcd.com/test.html", true},
		{"http://m.abcd.com/a/b/c/d/test.html?args=a&b=c", true},
		{"http://[::1]:9093", true},
		{"http://[::1]:909388", false},
		{"1200::AB00:1234::2552:7777:1313", false},
		{"http://[2001:db8:a0b:12f0::1]/index.html", true},
		{"http://[1200:0000:AB00:1234:0000:2552:7777:1313]", true},
		{"http://user:pass@[::1]:9093/a/b/c/?a=v#abc", true},
		{"https://127.0.0.1/a/b/c?a=v&c=11d", true},
		{"https://foo_bar.example.com", true},
		{"http://foo_bar.example.com", true},
		{"http://foo_bar_fizz_buzz.example.com", true},
		{"http://_cant_start_with_underescore", false},
		{"http://cant_end_with_underescore_", false},
		{"foo_bar.example.com", true},
		{"foo_bar_fizz_buzz.example.com", true},
		{"http://hello_world.example.com", true},
		// According to #212
		{"foo_bar-fizz-buzz:1313", true},
		{"foo_bar-fizz-buzz:13:13", false},
		{"foo_bar-fizz-buzz://1313", false},
	}
	for _, test := range tests {
		err := IsURL(test.param)
		if test.expected {
			require.NoError(t, err, "check IsURL(%s)", test.param)
		} else {
			require.Error(t, err, "check IsURL(%s)", test.param)
		}
	}
}

func TestIsFloat(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", true},
		{"  ", false},
		{"-.123", false},
		{"abacaba", false},
		{"1f", false},
		{"-1f", false},
		{"+1f", false},
		{"123", true},
		{"123.", true},
		{"123.123", true},
		{"-123.123", true},
		{"+123.123", true},
		{"0.123", true},
		{"-0.123", true},
		{"+0.123", true},
		{".0", true},
		{"01.123", true},
		{"-0.22250738585072011e-307", true},
		{"+0.22250738585072011e-307", true},
	}
	for _, test := range tests {
		err := IsFloat(test.param)
		if test.expected {
			require.NoError(t, err, "check IsFloat(%s)", test.param)
		} else {
			require.Error(t, err, "check IsFloat(%s)", test.param)
		}
	}
}

func TestIsNull(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"abacaba", false},
		{"", true},
	}
	for _, test := range tests {
		err := IsEmpty(test.param)
		if test.expected {
			require.NoError(t, err, "check IsNull(%s)", test.param)
		} else {
			require.Error(t, err, "check IsNull(%s)", test.param)
		}
	}
}

func TestIsJSON(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", true},
		{"145", true},
		{"asdf", false},
		{"123:f00", false},
		{"{\"Name\":\"Alice\",\"Body\":\"Hello\",\"Time\":1294706395881547000}", true},
		{"{}", true},
		{"{\"Key\":{\"Key\":{\"Key\":123}}}", true},
		{"[]", true},
		{"null", true},
	}
	for _, test := range tests {
		err := IsJSON(test.param)
		if test.expected {
			require.NoError(t, err, "check IsJSON(%s)", test.param)
		} else {
			require.Error(t, err, "check IsJSON(%s)", test.param)
		}
	}
}

func TestIsASCII(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", true},
		{"ｆｏｏbar", false},
		{"ｘｙｚ０９８", false},
		{"１２３456", false},
		{"ｶﾀｶﾅ", false},
		{"foobar", true},
		{"0987654321", true},
		{"test@example.com", true},
		{"1234abcDEF", true},
		{"", true},
	}
	for _, test := range tests {
		err := IsASCII(test.param)
		if test.expected {
			require.NoError(t, err, "check IsASCII(%s)", test.param)
		} else {
			require.Error(t, err, "check IsASCII(%s)", test.param)
		}
	}
}

func TestIsPrintableASCII(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", true},
		{"ｆｏｏbar", false},
		{"ｘｙｚ０９８", false},
		{"１２３456", false},
		{"ｶﾀｶﾅ", false},
		{"foobar", true},
		{"0987654321", true},
		{"test@example.com", true},
		{"1234abcDEF", true},
		{"newline\n", false},
		{"\x19test\x7F", false},
	}
	for _, test := range tests {
		err := IsPrintableASCII(test.param)
		if test.expected {
			require.NoError(t, err, "check IsPrintableASCII(%s)", test.param)
		} else {
			require.Error(t, err, "check IsPrintableASCII(%s)", test.param)
		}
	}
}

func TestIsBase64(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"TG9yZW0gaXBzdW0gZG9sb3Igc2l0IGFtZXQsIGNvbnNlY3RldHVyIGFkaXBpc2NpbmcgZWxpdC4=", true},
		{"Vml2YW11cyBmZXJtZW50dW0gc2VtcGVyIHBvcnRhLg==", true},
		{"U3VzcGVuZGlzc2UgbGVjdHVzIGxlbw==", true},
		{"MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuMPNS1Ufof9EW/M98FNw" +
			"UAKrwflsqVxaxQjBQnHQmiI7Vac40t8x7pIb8gLGV6wL7sBTJiPovJ0V7y7oc0Ye" +
			"rhKh0Rm4skP2z/jHwwZICgGzBvA0rH8xlhUiTvcwDCJ0kc+fh35hNt8srZQM4619" +
			"FTgB66Xmp4EtVyhpQV+t02g6NzK72oZI0vnAvqhpkxLeLiMCyrI416wHm5Tkukhx" +
			"QmcL2a6hNOyu0ixX/x2kSFXApEnVrJ+/IxGyfyw8kf4N2IZpW5nEP847lpfj0SZZ" +
			"Fwrd1mnfnDbYohX2zRptLy2ZUn06Qo9pkG5ntvFEPo9bfZeULtjYzIl6K8gJ2uGZ" + "HQIDAQAB", true},
		{"12345", false},
		{"", true},
		{"Vml2YW11cyBmZXJtZtesting123", false},
	}
	for _, test := range tests {
		err := IsBase64(test.param)
		if test.expected {
			require.NoError(t, err, "check IsBase64(%s)", test.param)
		} else {
			require.Error(t, err, "check IsBase64(%s)", test.param)
		}
	}
}

func TestIsIP(t *testing.T) {
	t.Parallel()

	// Without version
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", true},
		{"127.0.0.1", true},
		{"0.0.0.0", true},
		{"255.255.255.255", true},
		{"1.2.3.4", true},
		{"::1", true},
		{"2001:db8:0000:1:1:1:1:1", true},
		{"300.0.0.0", false},
	}
	for _, test := range tests {
		err := IsIP(test.param)
		if test.expected {
			require.NoError(t, err, "check IsIP(%s)", test.param)
		} else {
			require.Error(t, err, "check IsIP(%s)", test.param)
		}
	}

	// IPv4
	tests = []struct {
		param    string
		expected bool
	}{
		{"", true},
		{"127.0.0.1", true},
		{"0.0.0.0", true},
		{"255.255.255.255", true},
		{"1.2.3.4", true},
		{"::1", false},
		{"2001:db8:0000:1:1:1:1:1", false},
		{"300.0.0.0", false},
	}
	for _, test := range tests {
		err := IsIPv4(test.param)
		if test.expected {
			require.NoError(t, err, "check IsIPv4(%s)", test.param)
		} else {
			require.Error(t, err, "check IsIPv4(%s)", test.param)
		}

	}

	// IPv6
	tests = []struct {
		param    string
		expected bool
	}{
		{"", true},
		{"127.0.0.1", false},
		{"0.0.0.0", false},
		{"255.255.255.255", false},
		{"1.2.3.4", false},
		{"::1", true},
		{"2001:db8:0000:1:1:1:1:1", true},
		{"300.0.0.0", false},
	}
	for _, test := range tests {
		err := IsIPv6(test.param)
		if test.expected {
			require.NoError(t, err, "check IsIPv6(%s)", test.param)
		} else {
			require.Error(t, err, "check IsIPv6(%s)", test.param)
		}

	}
}

func TestIsPort(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"1", true},
		{"65535", true},
		{"0", false},
		{"65536", false},
		{"65538", false},
	}

	for _, test := range tests {
		err := IsPort(test.param)
		if test.expected {
			require.NoError(t, err, "check IsPort(%s)", test.param)
		} else {
			require.Error(t, err, "check IsPort(%s)", test.param)
		}

	}
}

func TestIsMAC(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"3D:F2:C9:A6:B3:4F", true},
		{"3D-F2-C9-A6-B3:4F", false},
		{"123", false},
		{"", true},
		{"abacaba", false},
	}
	for _, test := range tests {
		err := IsMAC(test.param)
		if test.expected {
			require.NoError(t, err, "check IsMAC(%s)", test.param)
		} else {
			require.Error(t, err, "check IsMAC(%s)", test.param)
		}

	}
}

func TestIsLatitude(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"-90.000", true},
		{"+90", true},
		{"47.1231231", true},
		{"+99.9", false},
		{"108", false},
	}
	for _, test := range tests {
		err := IsLatitude(test.param)
		if test.expected {
			require.NoError(t, err, "check IsLatitude(%s)", test.param)
		} else {
			require.Error(t, err, "check IsLatitude(%s)", test.param)
		}

	}
}

func TestIsLongitude(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", true},
		{"-180.000", true},
		{"180.1", false},
		{"+73.234", true},
		{"+382.3811", false},
		{"23.11111111", true},
	}
	for _, test := range tests {
		err := IsLongitude(test.param)
		if test.expected {
			require.NoError(t, err, "check IsLongitude(%s)", test.param)
		} else {
			require.Error(t, err, "check IsLongitude(%s)", test.param)
		}

	}
}

func TestIsTime(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		param    string
		format   string
		expected bool
	}{
		{"2016-12-31 11:00", time.RFC3339, false},
		{"2016-12-31 11:00:00", time.RFC3339, false},
		{"2016-12-31T11:00", time.RFC3339, false},
		{"2016-12-31T11:00:00", time.RFC3339, false},
		{"2016-12-31T11:00:00Z", time.RFC3339, true},
		{"2016-12-31T11:00:00+01:00", time.RFC3339, true},
		{"2016-12-31T11:00:00-01:00", time.RFC3339, true},
		{"2016-12-31T11:00:00.05Z", time.RFC3339, true},
		{"2016-12-31T11:00:00.05-01:00", time.RFC3339, true},
		{"2016-12-31T11:00:00.05+01:00", time.RFC3339, true},
		{"2016-12-31T11:00:00", RF3339WithoutZone, true},
		{"2016-12-31T11:00:00Z", RF3339WithoutZone, false},
		{"2016-12-31T11:00:00+01:00", RF3339WithoutZone, false},
		{"2016-12-31T11:00:00-01:00", RF3339WithoutZone, false},
		{"2016-12-31T11:00:00.05Z", RF3339WithoutZone, false},
		{"2016-12-31T11:00:00.05-01:00", RF3339WithoutZone, false},
		{"2016-12-31T11:00:00.05+01:00", RF3339WithoutZone, false},
	}
	for _, test := range tests {
		err := IsTime(test.param, test.format)
		if test.expected {
			require.NoError(t, err, "check IsTime(%s, %s)", test.param, test.format)
		} else {
			require.Error(t, err, "check IsTime(%s, %s)", test.param, test.format)
		}

	}
}

func TestIsRFC3339(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		param    string
		expected bool
	}{
		{"2016-12-31 11:00", false},
		{"2016-12-31 11:00:00", false},
		{"2016-12-31T11:00", false},
		{"2016-12-31T11:00:00", false},
		{"2016-12-31T11:00:00Z", true},
		{"2016-12-31T11:00:00+01:00", true},
		{"2016-12-31T11:00:00-01:00", true},
		{"2016-12-31T11:00:00.05Z", true},
		{"2016-12-31T11:00:00.05-01:00", true},
		{"2016-12-31T11:00:00.05+01:00", true},
	}
	for _, test := range tests {
		err := IsRFC3339(test.param)
		if test.expected {
			require.NoError(t, err, "check IsRFC3339(%s)", test.param)
		} else {
			require.Error(t, err, "check IsRFC3339(%s)", test.param)
		}

	}
}

func TestIsISO4217(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", true},
		{"ABCD", false},
		{"A", false},
		{"ZZZ", false},
		{"usd", false},
		{"USD", true},
	}
	for _, test := range tests {
		err := IsISO4217(test.param)
		if test.expected {
			require.NoError(t, err, "check IsISO4217(%s)", test.param)
		} else {
			require.Error(t, err, "check IsISO4217(%s)", test.param)
		}

	}
}

func TestLength(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		value    string
		min      string
		max      string
		expected bool
	}{
		{"123456", "0", "100", true},
		{"1239999", "0", "0", false},
		{"1239asdfasf99", "100", "200", false},
		{"1239999asdff29", "10", "30", true},
		{"世界真大", "4", "4", true},
	}
	for _, test := range tests {
		err := Length(test.value, test.min, test.max)
		if test.expected {
			if test.expected {
				require.NoError(t, err, "check Length(%s, %s, %s)", test.value, test.min, test.max)
			} else {
				require.Error(t, err, "check Length(%s, %s, %s)", test.value, test.min, test.max)
			}

		} else {
			require.Error(t, err, "check Length(%s, %s, %s)", test.value, test.min, test.max)
		}
	}
}

func TestIsIn(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		value    string
		params   []string
		expected bool
	}{
		{"PRESENT", []string{"PRESENT"}, true},
		{"PRESENT", []string{"PRESENT", "PRÉSENTE", "NOTABSENT"}, true},
		{"PRÉSENTE", []string{"PRESENT", "PRÉSENTE", "NOTABSENT"}, true},
		{"PRESENT", []string{}, false},
		{"PRESENT", nil, false},
		{"ABSENT", []string{"PRESENT", "PRÉSENTE", "NOTABSENT"}, false},
		{"", []string{"PRESENT", "PRÉSENTE", "NOTABSENT"}, true},
	}
	for _, test := range tests {
		err := IsIn(test.value, test.params...)
		if test.expected {
			require.NoError(t, err, "check IsIn(%s, %v)", test.value, test.params)
		} else {
			require.Error(t, err, "check IsIn(%s, %v)", test.value, test.params)
		}

	}
}
