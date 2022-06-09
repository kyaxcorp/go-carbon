package carbon

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestCarbon_SetTimezone(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		timezone string // 输入参数
		expected string // 期望值
	}{
		{"0000-00-00 00:00:00", PRC, ""},
		{"2020-08-05 13:14:15", PRC, "2020-08-05 13:14:15"},
		{"2020-08-05 13:14:15", Tokyo, "2020-08-05 12:14:15"},
		{"2020-08-05 13:14:15", London, "2020-08-05 20:14:15"},
	}

	for index, test := range tests {
		c := SetTimezone(test.timezone).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SetLocale(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		locale   string // 输入参数
		expected string // 期望值
	}{
		{"0000-00-00", "en", ""},
		{"2020-08-05", "en", "August"},
		{"2020-08-05", "jp", "はちがつ"},
		{"2020-08-05", "zh-CN", "八月"},
	}

	for index, test := range tests {
		c := SetLocale(test.locale).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToMonthString(PRC), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input).SetLocale(test.locale)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToMonthString(PRC), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SetLanguage(t *testing.T) {
	lang := NewLanguage()
	resources := map[string]string{
		"seasons": "spring|summer|autumn|winter",
	}
	lang.SetLocale("en")
	if lang.Error == nil {
		lang.SetResources(resources)
	}

	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"2020-08-05", "summer"},
	}

	for index, test := range tests {
		c := Parse(test.input).SetLanguage(lang)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Season(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := SetLanguage(lang).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Season(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SetDateTime(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input                                  string // 输入值
		year, month, day, hour, minute, second int    // 输入参数
		expected                               string // 期望值
	}{
		{"2020-01-01", 2019, 02, 02, 13, 14, 15, "2019-02-02 13:14:15"},
		{"2020-01-01", 2019, 02, 31, 13, 14, 15, "2019-03-03 13:14:15"},
	}

	for index, test := range tests {
		c := Parse(test.input).SetDateTime(test.year, test.month, test.day, test.hour, test.minute, test.second)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SetDateTimeMilli(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input                                               string // 输入值
		year, month, day, hour, minute, second, millisecond int    // 输入参数
		expected                                            string // 期望值
	}{
		{"2020-01-01", 2019, 02, 02, 13, 14, 15, 999, "2019-02-02 13:14:15.999 +0800 CST"},
		{"2020-01-01", 2019, 02, 31, 13, 14, 15, 999, "2019-03-03 13:14:15.999 +0800 CST"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC).SetDateTimeMilli(test.year, test.month, test.day, test.hour, test.minute, test.second, test.millisecond)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SetDateTimeMicro(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input                                               string // 输入值
		year, month, day, hour, minute, second, microsecond int    // 输入参数
		expected                                            string // 期望值
	}{
		{"2020-01-01", 2019, 02, 02, 13, 14, 15, 999999, "2019-02-02 13:14:15.999999 +0800 CST"},
		{"2020-01-01", 2019, 02, 31, 13, 14, 15, 999999, "2019-03-03 13:14:15.999999 +0800 CST"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC).SetDateTimeMicro(test.year, test.month, test.day, test.hour, test.minute, test.second, test.microsecond)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SetDateTimeNano(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input                                              string // 输入值
		year, month, day, hour, minute, second, nanosecond int    // 输入参数
		expected                                           string // 期望值
	}{
		{"2020-01-01", 2019, 02, 02, 13, 14, 15, 999999999, "2019-02-02 13:14:15.999999999 +0800 CST"},
		{"2020-01-01", 2019, 02, 31, 13, 14, 15, 999999999, "2019-03-03 13:14:15.999999999 +0800 CST"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC).SetDateTimeNano(test.year, test.month, test.day, test.hour, test.minute, test.second, test.nanosecond)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SetDate(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input            string // 输入值
		year, month, day int    // 输入参数
		expected         string // 期望值
	}{
		{"2020-01-01", 2019, 02, 02, "2019-02-02 00:00:00"},
		{"2020-01-01", 2019, 02, 31, "2019-03-03 00:00:00"},
	}

	for index, test := range tests {
		c := Parse(test.input).SetDate(test.year, test.month, test.day)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SetDateMilli(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input                         string // 输入值
		year, month, day, millisecond int    // 输入参数
		expected                      string // 期望值
	}{
		{"2020-01-01", 2019, 02, 02, 999, "2019-02-02 00:00:00.999"},
		{"2020-01-01", 2019, 02, 31, 999, "2019-03-03 00:00:00.999"},
	}

	for index, test := range tests {
		c := Parse(test.input).SetDateMilli(test.year, test.month, test.day, test.millisecond)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeMilliString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SetDateMicro(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input                         string // 输入值
		year, month, day, microsecond int    // 输入参数
		expected                      string // 期望值
	}{
		{"2020-01-01", 2019, 02, 02, 999999, "2019-02-02 00:00:00.999999"},
		{"2020-01-01", 2019, 02, 31, 999999, "2019-03-03 00:00:00.999999"},
	}

	for index, test := range tests {
		c := Parse(test.input).SetDateMicro(test.year, test.month, test.day, test.microsecond)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeMicroString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SetDateNano(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input                        string // 输入值
		year, month, day, nanosecond int    // 输入参数
		expected                     string // 期望值
	}{
		{"2020-01-01", 2019, 02, 02, 999999999, "2019-02-02 00:00:00.999999999"},
		{"2020-01-01", 2019, 02, 31, 999999999, "2019-03-03 00:00:00.999999999"},
	}

	for index, test := range tests {
		c := Parse(test.input).SetDateNano(test.year, test.month, test.day, test.nanosecond)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeNanoString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SetTime(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input                string // 输入值
		hour, minute, second int    // 输入参数
		expected             string // 期望值
	}{
		{"2020-08-05", 13, 14, 15, "2020-08-05 13:14:15"},
		{"2020-08-05", 13, 14, 90, "2020-08-05 13:15:30"},
	}

	for index, test := range tests {
		c := Parse(test.input).SetTime(test.hour, test.minute, test.second)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SetTimeMilli(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input                             string // 输入值
		hour, minute, second, millisecond int    // 输入参数
		expected                          string // 期望值
	}{
		{"2020-08-05", 13, 14, 15, 999, "2020-08-05 13:14:15.999"},
		{"2020-08-05", 13, 14, 90, 999, "2020-08-05 13:15:30.999"},
	}

	for index, test := range tests {
		c := Parse(test.input).SetTimeMilli(test.hour, test.minute, test.second, test.millisecond)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeMilliString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SetTimeMicro(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input                             string // 输入值
		hour, minute, second, microsecond int    // 输入参数
		expected                          string // 期望值
	}{
		{"2020-08-05", 13, 14, 15, 999999, "2020-08-05 13:14:15.999999"},
		{"2020-08-05", 13, 14, 90, 999999, "2020-08-05 13:15:30.999999"},
	}

	for index, test := range tests {
		c := Parse(test.input).SetTimeMicro(test.hour, test.minute, test.second, test.microsecond)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeMicroString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SetTimeNano(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input                            string // 输入值
		hour, minute, second, nanosecond int    // 输入参数
		expected                         string // 期望值
	}{
		{"2020-08-05", 13, 14, 15, 999999999, "2020-08-05 13:14:15.999999999"},
		{"2020-08-05", 13, 14, 90, 999999999, "2020-08-05 13:15:30.999999999"},
	}

	for index, test := range tests {
		c := Parse(test.input).SetTimeNano(test.hour, test.minute, test.second, test.nanosecond)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeNanoString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SetYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		year     int    // 输入参数
		expected string // 期望值
	}{
		{"2020-01-01", 2019, "2019-01-01"},
		{"2020-01-31", 2019, "2019-01-31"},
		{"2020-02-01", 2019, "2019-02-01"},
		{"2020-02-28", 2019, "2019-02-28"},
		{"2020-02-29", 2019, "2019-03-01"},
	}

	for index, test := range tests {
		c := Parse(test.input).SetYear(test.year)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SetYearNoOverflow(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		year     int    // 输入参数
		expected string // 期望值
	}{
		{"2020-01-01", 2019, "2019-01-01"},
		{"2020-01-31", 2019, "2019-01-31"},
		{"2020-02-01", 2019, "2019-02-01"},
		{"2020-02-28", 2019, "2019-02-28"},
		{"2020-02-29", 2019, "2019-02-28"},
	}

	for index, test := range tests {
		c := Parse(test.input).SetYearNoOverflow(test.year)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SetMonth(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		month    int    // 输入参数
		expected string // 期望值
	}{
		{"2020-01-01", 2, "2020-02-01"},
		{"2020-01-30", 2, "2020-03-01"},
		{"2020-01-31", 2, "2020-03-02"},
		{"2020-08-05", 2, "2020-02-05"},
	}

	for index, test := range tests {
		c := Parse(test.input).SetMonth(test.month)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SetMonthNoOverflow(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		month    int    // 输入参数
		expected string // 期望值
	}{
		{"2020-01-01", 2, "2020-02-01"},
		{"2020-01-30", 2, "2020-02-29"},
		{"2020-01-31", 2, "2020-02-29"},
		{"2020-08-05", 2, "2020-02-05"},
	}

	for index, test := range tests {
		c := Parse(test.input).SetMonthNoOverflow(test.month)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SetWeekStartsAt(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		week     string // 输入参数
		expected string // 期望值
	}{
		{"", Sunday, ""},
		{"0000-00-00 00:00:00", Sunday, ""},
		{"", Monday, ""},
		{"0000-00-00 00:00:00", Monday, ""},

		{"2021-06-13", Sunday, "2021-06-13 00:00:00"},
		{"2021-06-14", Sunday, "2021-06-13 00:00:00"},
		{"2021-06-18", Sunday, "2021-06-13 00:00:00"},

		{"2021-06-13", Monday, "2021-06-07 00:00:00"},
		{"2021-06-14", Monday, "2021-06-14 00:00:00"},
		{"2021-06-18", Monday, "2021-06-14 00:00:00"},

		{"2021-06-13", Tuesday, "2021-06-08 00:00:00"},
		{"2021-06-14", Tuesday, "2021-06-08 00:00:00"},
		{"2021-06-18", Tuesday, "2021-06-15 00:00:00"},

		{"2021-06-13", Wednesday, "2021-06-09 00:00:00"},
		{"2021-06-14", Wednesday, "2021-06-09 00:00:00"},
		{"2021-06-18", Wednesday, "2021-06-16 00:00:00"},

		{"2021-06-13", Thursday, "2021-06-10 00:00:00"},
		{"2021-06-14", Thursday, "2021-06-10 00:00:00"},
		{"2021-06-18", Thursday, "2021-06-17 00:00:00"},

		{"2021-06-13", Friday, "2021-06-11 00:00:00"},
		{"2021-06-14", Friday, "2021-06-11 00:00:00"},
		{"2021-06-18", Friday, "2021-06-18 00:00:00"},

		{"2021-06-13", Saturday, "2021-06-12 00:00:00"},
		{"2021-06-14", Saturday, "2021-06-12 00:00:00"},
		{"2021-06-18", Saturday, "2021-06-12 00:00:00"},
	}

	for index, test := range tests {
		c := Parse(test.input).SetWeekStartsAt(test.week).StartOfWeek()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SetDay(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		day      int    // 输入参数
		expected string // 期望值
	}{
		{"2020-01-01", 31, "2020-01-31"},
		{"2020-02-01", 31, "2020-03-02"},
		{"2020-02-28", 31, "2020-03-02"},
		{"2020-02-29", 31, "2020-03-02"},
	}

	for index, test := range tests {
		c := Parse(test.input).SetDay(test.day)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SetHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		hour     int    // 输入参数
		expected string // 期望值
	}{
		{"2020-08-05 13:14:15", 10, "2020-08-05 10:14:15"},
		{"2020-08-05 13:14:15", 24, "2020-08-06 00:14:15"},
	}

	for index, test := range tests {
		c := Parse(test.input).SetHour(test.hour)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SetMinute(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		minute   int    // 输入参数
		expected string // 期望值
	}{
		{"2020-08-05 13:14:15", 10, "2020-08-05 13:10:15"},
		{"2020-08-05 13:14:15", 60, "2020-08-05 14:00:15"},
	}

	for index, test := range tests {
		c := Parse(test.input).SetMinute(test.minute)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SetSecond(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		second   int    // 输入参数
		expected int    // 期望值
	}{
		{"2020-08-05 13:14:15", 10, 10},
		{"2020-08-05 13:14:15", 59, 59},
	}

	for index, test := range tests {
		c := Parse(test.input).SetSecond(test.second)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Second(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SetMillisecond(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input       string // 输入值
		millisecond int    // 输入参数
		expected    int    // 期望值
	}{
		{"2020-08-05 13:14:15", 100, 100},
		{"2020-08-05 13:14:15", 999, 999},
	}

	for index, test := range tests {
		c := Parse(test.input).SetMillisecond(test.millisecond)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Millisecond(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SetMicrosecond(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input       string // 输入值
		microsecond int    // 输入参数
		expected    int    // 期望值
	}{
		{"2020-08-05 13:14:15", 100000, 100000},
		{"2020-08-05 13:14:15", 999999, 999999},
	}

	for index, test := range tests {
		c := Parse(test.input).SetMicrosecond(test.microsecond)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Microsecond(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SetNanosecond(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input      string // 输入值
		nanosecond int    // 输入参数
		expected   int    // 期望值
	}{
		{"2020-08-05 13:14:15", 100000000, 100000000},
		{"2020-08-05 13:14:15", 999999999, 999999999},
	}

	for index, test := range tests {
		c := Parse(test.input).SetNanosecond(test.nanosecond)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Nanosecond(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestError_Setter(t *testing.T) {
	input, timezone, locale, year, month, day, hour, minute, second, millisecond, microsecond, nanosecond := "2020-08-50 13:14:15", "xxx", "xxx", 2020, 8, 50, 13, 14, 15, 999, 999999, 999999999
	c := Parse(input)

	tz1 := SetTimezone(timezone)
	assert.NotNil(t, tz1.Error, "It should catch an exception in SetTimezone()")
	tz2 := Now(timezone).SetTimezone(timezone)
	assert.NotNil(t, tz2.Error, "It should catch an exception in SetTimezone()")

	loc1 := SetLocale(locale)
	assert.NotNil(t, loc1.Error, "It should catch an exception in SetLocale()")
	loc2 := SetLocale(locale).SetLocale(PRC)
	assert.NotNil(t, loc2.Error, "It should catch an exception in SetLocale()")

	lang := NewLanguage()
	lang.SetLocale(locale)
	lang1 := SetLanguage(lang)
	lang2 := SetLocale(locale).SetLanguage(lang)
	assert.NotNil(t, lang1.Error, "It should catch an exception in SetLanguage()")
	assert.NotNil(t, lang2.SetLanguage(lang).Error, "It should catch an exception in SetLanguage()")

	assert.NotNil(t, c.SetDateTime(year, month, day, hour, minute, second).Error, "It should catch an exception in SetDateTime()")
	assert.NotNil(t, c.SetDateTimeMilli(year, month, day, hour, minute, second, millisecond).Error, "It should catch an exception in SetDateTimeMilli()")
	assert.NotNil(t, c.SetDateTimeMicro(year, month, day, hour, minute, second, microsecond).Error, "It should catch an exception in SetDateTimeMicro()")
	assert.NotNil(t, c.SetDateTimeNano(year, month, day, hour, minute, second, nanosecond).Error, "It should catch an exception in SetDateTimeNano()")

	assert.NotNil(t, c.SetDate(year, month, day).Error, "It should catch an exception in SeDate()")
	assert.NotNil(t, c.SetDateMilli(year, month, day, millisecond).Error, "It should catch an exception in SetDateMilli()")
	assert.NotNil(t, c.SetDateMicro(year, month, day, microsecond).Error, "It should catch an exception in SetDateMicro()")
	assert.NotNil(t, c.SetDateNano(year, month, day, nanosecond).Error, "It should catch an exception in SetDateNano()")

	assert.NotNil(t, c.SetTime(hour, minute, second).Error, "It should catch an exception in SetTime()")
	assert.NotNil(t, c.SetTimeMilli(hour, minute, second, millisecond).Error, "It should catch an exception in SetTimeMilli()")
	assert.NotNil(t, c.SetTimeMicro(hour, minute, second, microsecond).Error, "It should catch an exception in SetTimeMicro()")
	assert.NotNil(t, c.SetTimeNano(hour, minute, second, nanosecond).Error, "It should catch an exception in SetTimeNano()")

	assert.NotNil(t, c.SetYear(year).Error, "It should catch an exception in SetYear()")
	assert.NotNil(t, c.SetYearNoOverflow(year).Error, "It should catch an exception in SetYearNoOverflow()")
	assert.NotNil(t, c.SetMonth(month).Error, "It should catch an exception in SetMonth()")
	assert.NotNil(t, c.SetMonthNoOverflow(month).Error, "It should catch an exception in SetMonthNoOverflow()")
	assert.NotNil(t, c.SetDay(day).Error, "It should catch an exception in SetDay()")
	assert.NotNil(t, c.SetHour(hour).Error, "It should catch an exception in SetHour()")
	assert.NotNil(t, c.SetMinute(minute).Error, "It should catch an exception in SetMinute()")
	assert.NotNil(t, c.SetSecond(second).Error, "It should catch an exception in SetSecond()")
	assert.NotNil(t, c.SetMillisecond(millisecond).Error, "It should catch an exception in SetMillisecond()")
	assert.NotNil(t, c.SetMicrosecond(microsecond).Error, "It should catch an exception in SetMicrosecond()")
	assert.NotNil(t, c.SetNanosecond(nanosecond).Error, "It should catch an exception in SetNanosecond()")
}
