package carbon

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_DiffInYears(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		param    string // 参数值
		expected int64  // 期望值
	}{
		{"2020-08-05 13:14:15", "2020-07-28 13:14:00", 0},
		{"2020-12-31 13:14:15", "2021-01-01 13:14:15", 0},
		{"2020-08-05 13:14:15", "2021-08-28 13:14:59", 1},
		{"2020-08-05 13:14:15", "2018-08-28 13:14:59", -2},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffInYears(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_DiffAbsInYears(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		param    string // 参数值
		expected int64  // 期望值
	}{
		{"2020-08-05 13:14:15", "2020-07-28 13:14:00", 0},
		{"2020-12-31 13:14:15", "2021-01-01 13:14:15", 0},
		{"2020-08-05 13:14:15", "2021-08-28 13:14:59", 1},
		{"2020-08-05 13:14:15", "2018-08-28 13:14:59", 2},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffAbsInYears(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_DiffInMonths(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		param    string // 参数值
		expected int64  // 期望值
	}{
		{"2020-08-05 13:14:15", "2020-07-28 13:14:00", 0},
		{"2020-08-05 13:14:15", "2020-09-06 13:14:59", 1},
		{"2020-12-31 13:14:15", "2021-01-31 13:14:15", 1},
		{"2020-08-05 13:14:15", "2021-08-28 13:14:59", 12},
		{"2020-08-05 13:14:15", "2018-08-28 13:14:59", -24},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffInMonths(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_DiffAbsInMonths(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		param    string // 参数值
		expected int64  // 期望值
	}{
		{"2020-08-05 13:14:15", "2020-07-28 13:14:00", 0},
		{"2020-12-31 13:14:15", "2021-01-01 13:14:15", 0},
		{"2020-08-05 13:14:15", "2021-08-28 13:14:59", 12},
		{"2020-08-05 13:14:15", "2018-08-28 13:14:59", 24},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffAbsInMonths(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_DiffInWeeks(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		param    string // 参数值
		expected int64  // 期望值
	}{
		{"2020-08-05 13:14:15", "2020-07-28 13:14:00", -1},
		{"2020-08-05 13:14:15", "2020-07-28 13:14:15", -1},
		{"2020-08-05 13:14:15", "2020-07-28 13:14:59", -1},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:00", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:59", 0},
		{"2020-08-05 13:14:15", "2020-08-12 13:14:00", 0},
		{"2020-08-05 13:14:15", "2020-08-12 13:14:15", 1},
		{"2020-08-05 13:14:15", "2020-08-12 13:14:59", 1},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffInWeeks(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_DiffAbsInWeeks(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		param    string // 参数值
		expected int64  // 期望值
	}{
		{"2020-08-05 13:14:15", "2020-07-28 13:14:00", 1},
		{"2020-08-05 13:14:15", "2020-07-28 13:14:15", 1},
		{"2020-08-05 13:14:15", "2020-07-28 13:14:59", 1},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:00", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:59", 0},
		{"2020-08-05 13:14:15", "2020-08-12 13:14:00", 0},
		{"2020-08-05 13:14:15", "2020-08-12 13:14:15", 1},
		{"2020-08-05 13:14:15", "2020-08-12 13:14:59", 1},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffAbsInWeeks(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_DiffInDays(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		param    string // 参数值
		expected int64  // 期望值
	}{
		{"2020-08-05 13:14:15", "2020-08-04 13:00:00", -1},
		{"2020-08-05 13:14:15", "2020-08-04 13:14:15", -1},
		{"2020-08-05 13:14:15", "2020-08-04 13:14:59", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:00:00", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:59", 0},
		{"2020-08-05 13:14:15", "2020-08-06 13:00:00", 0},
		{"2020-08-05 13:14:15", "2020-08-06 13:14:15", 1},
		{"2020-08-05 13:14:15", "2020-08-06 13:14:59", 1},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffInDays(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_DiffAbsInDays(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		param    string // 参数值
		expected int64  // 期望值
	}{
		{"2020-08-05 13:14:15", "2020-08-04 13:00:00", 1},
		{"2020-08-05 13:14:15", "2020-08-04 13:14:15", 1},
		{"2020-08-05 13:14:15", "2020-08-04 13:14:59", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:00:00", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:59", 0},
		{"2020-08-05 13:14:15", "2020-08-06 13:00:00", 0},
		{"2020-08-05 13:14:15", "2020-08-06 13:14:15", 1},
		{"2020-08-05 13:14:15", "2020-08-06 13:14:59", 1},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffAbsInDays(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_DiffInHours(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		param    string // 参数值
		expected int64  // 期望值
	}{
		{"2020-08-05 13:14:15", "2020-08-05 12:14:00", -1},
		{"2020-08-05 13:14:15", "2020-08-05 12:14:15", -1},
		{"2020-08-05 13:14:15", "2020-08-05 12:14:59", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:00", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:59", 0},
		{"2020-08-05 13:14:15", "2020-08-05 14:14:00", 0},
		{"2020-08-05 13:14:15", "2020-08-05 14:14:15", 1},
		{"2020-08-05 13:14:15", "2020-08-05 14:14:59", 1},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffInHours(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_DiffAbsInHours(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		param    string // 参数值
		expected int64  // 期望值
	}{
		{"2020-08-05 13:14:15", "2020-08-05 12:14:00", 1},
		{"2020-08-05 13:14:15", "2020-08-05 12:14:15", 1},
		{"2020-08-05 13:14:15", "2020-08-05 12:14:59", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:00", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:59", 0},
		{"2020-08-05 13:14:15", "2020-08-05 14:14:00", 0},
		{"2020-08-05 13:14:15", "2020-08-05 14:14:15", 1},
		{"2020-08-05 13:14:15", "2020-08-05 14:14:59", 1},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffAbsInHours(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_DiffInMinutes(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		param    string // 参数值
		expected int64  // 期望值
	}{
		{"2020-08-05 13:14:15", "2020-08-05 13:13:00", -1},
		{"2020-08-05 13:14:15", "2020-08-05 13:13:15", -1},
		{"2020-08-05 13:14:15", "2020-08-05 13:13:59", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:15:00", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:15:15", 1},
		{"2020-08-05 13:14:15", "2020-08-05 13:15:59", 1},
		{"2020-08-05 13:14:15", "2020-08-05 13:16:00", 1},
		{"2020-08-05 13:14:15", "2020-08-05 13:16:15", 2},
		{"2020-08-05 13:14:15", "2020-08-05 13:16:59", 2},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffInMinutes(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_DiffAbsInMinutes(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		param    string // 参数值
		expected int64  // 期望值
	}{
		{"2020-08-05 13:14:15", "2020-08-05 13:13:00", 1},
		{"2020-08-05 13:14:15", "2020-08-05 13:13:15", 1},
		{"2020-08-05 13:14:15", "2020-08-05 13:13:59", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:15:00", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:15:15", 1},
		{"2020-08-05 13:14:15", "2020-08-05 13:15:59", 1},
		{"2020-08-05 13:14:15", "2020-08-05 13:16:00", 1},
		{"2020-08-05 13:14:15", "2020-08-05 13:16:15", 2},
		{"2020-08-05 13:14:15", "2020-08-05 13:16:59", 2},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffAbsInMinutes(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_DiffInSeconds(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		param    string // 参数值
		expected int64  // 期望值
	}{
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:20", 5},
		{"2020-08-05 13:14:20", "2020-08-05 13:14:15", -5},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffInSeconds(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_DiffAbsInSeconds(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		param    string // 参数值
		expected int64  // 期望值
	}{
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", 0},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:20", 5},
		{"2020-08-05 13:14:20", "2020-08-05 13:14:15", 5},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.DiffAbsInSeconds(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_DiffInString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    Carbon // 输入值
		expected string // 期望值
	}{
		{Now(), "just now"},

		{Now().AddYearsNoOverflow(1), "-1 year"},
		{Now().SubYearsNoOverflow(1), "1 year"},
		{Now().AddYearsNoOverflow(10), "-10 years"},
		{Now().SubYearsNoOverflow(10), "10 years"},

		{Now().AddMonthsNoOverflow(1), "-1 month"},
		{Now().SubMonthsNoOverflow(1), "1 month"},
		{Now().AddMonthsNoOverflow(10), "-10 months"},
		{Now().SubMonthsNoOverflow(10), "10 months"},

		{Now().AddDays(1), "-1 day"},
		{Now().SubDays(1), "1 day"},
		{Now().AddDays(10), "-1 week"},
		{Now().SubDays(10), "1 week"},

		{Now().AddHours(1), "-1 hour"},
		{Now().SubHours(1), "1 hour"},
		{Now().AddHours(10), "-10 hours"},
		{Now().SubHours(10), "10 hours"},

		{Now().AddMinutes(1), "-1 minute"},
		{Now().SubMinutes(1), "1 minute"},
		{Now().AddMinutes(10), "-10 minutes"},
		{Now().SubMinutes(10), "10 minutes"},

		{Now().AddSeconds(1), "-1 second"},
		{Now().SubSeconds(1), "1 second"},
		{Now().AddSeconds(10), "-10 seconds"},
		{Now().SubSeconds(10), "10 seconds"},
	}

	for index, test := range tests {
		c := test.input
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.DiffInString(), "Current test index is "+strconv.Itoa(index))
	}
	for index, test := range tests {
		c := test.input
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.DiffInString(Now()), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLangError_DiffInString(t *testing.T) {
	lang := NewLanguage()
	lang.SetLocale("xxx")
	c := Now().SetLanguage(lang).AddMonths(1)
	assert.NotNil(t, c.Error, "It should catch an exception in DiffInString()")
	assert.Equal(t, "", c.DiffInString())
}

func TestCarbon_DiffAbsInString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    Carbon // 输入值
		expected string // 期望值
	}{
		{Now(), "just now"},

		{Now().AddYearsNoOverflow(1), "1 year"},
		{Now().SubYearsNoOverflow(1), "1 year"},
		{Now().AddYearsNoOverflow(10), "10 years"},
		{Now().SubYearsNoOverflow(10), "10 years"},

		{Now().AddMonthsNoOverflow(1), "1 month"},
		{Now().SubMonthsNoOverflow(1), "1 month"},
		{Now().AddMonthsNoOverflow(10), "10 months"},
		{Now().SubMonthsNoOverflow(10), "10 months"},

		{Now().AddDays(1), "1 day"},
		{Now().SubDays(1), "1 day"},
		{Now().AddDays(10), "1 week"},
		{Now().SubDays(10), "1 week"},

		{Now().AddHours(1), "1 hour"},
		{Now().SubHours(1), "1 hour"},
		{Now().AddHours(10), "10 hours"},
		{Now().SubHours(10), "10 hours"},

		{Now().AddMinutes(1), "1 minute"},
		{Now().SubMinutes(1), "1 minute"},
		{Now().AddMinutes(10), "10 minutes"},
		{Now().SubMinutes(10), "10 minutes"},

		{Now().AddSeconds(1), "1 second"},
		{Now().SubSeconds(1), "1 second"},
		{Now().AddSeconds(10), "10 seconds"},
		{Now().SubSeconds(10), "10 seconds"},
	}

	for index, test := range tests {
		c := test.input
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.DiffAbsInString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := test.input
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.DiffAbsInString(Now()), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLangError_DiffAbsInString(t *testing.T) {
	lang := NewLanguage()
	lang.SetLocale("xxx")
	c := Now().SetLanguage(lang).AddMonths(1)
	assert.NotNil(t, c.Error, "It should catch an exception in DiffInStringWithAbs()")
	assert.Equal(t, "", c.DiffAbsInString())
}

func TestCarbon_DiffForHumans(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    Carbon // 输入值
		expected string // 期望值
	}{
		{Now(), "just now"},

		{Now().AddYearsNoOverflow(1), "1 year from now"},
		{Now().SubYearsNoOverflow(1), "1 year ago"},
		{Now().AddYearsNoOverflow(10), "10 years from now"},
		{Now().SubYearsNoOverflow(10), "10 years ago"},

		{Now().AddMonthsNoOverflow(1), "1 month from now"},
		{Now().SubMonthsNoOverflow(1), "1 month ago"},
		{Now().AddMonthsNoOverflow(10), "10 months from now"},
		{Now().SubMonthsNoOverflow(10), "10 months ago"},

		{Now().AddDays(1), "1 day from now"},
		{Now().SubDays(1), "1 day ago"},
		{Now().AddDays(10), "1 week from now"},
		{Now().SubDays(10), "1 week ago"},

		{Now().AddHours(1), "1 hour from now"},
		{Now().SubHours(1), "1 hour ago"},
		{Now().AddHours(10), "10 hours from now"},
		{Now().SubHours(10), "10 hours ago"},

		{Now().AddMinutes(1), "1 minute from now"},
		{Now().SubMinutes(1), "1 minute ago"},
		{Now().AddMinutes(10), "10 minutes from now"},
		{Now().SubMinutes(10), "10 minutes ago"},

		{Now().AddSeconds(1), "1 second from now"},
		{Now().SubSeconds(1), "1 second ago"},
		{Now().AddSeconds(10), "10 seconds from now"},
		{Now().SubSeconds(10), "10 seconds ago"},
	}

	for index, test := range tests {
		c := test.input
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.DiffForHumans(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestLangError_DiffForHumans(t *testing.T) {
	lang := NewLanguage()
	lang.SetLocale("xxx")
	c := Now().SetLanguage(lang).AddMonths(1)
	assert.NotNil(t, c.Error, "It should catch an exception in DiffForHumans()")
	assert.Equal(t, "", c.DiffForHumans())
}
