package carbon

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_IsZero(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", true},
		{"0", true},
		{"0000-00-00", true},
		{"00:00:00", true},
		{"0000-00-00 00:00:00", true},

		{"2020-08-05", false},
	}

	for index, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsZero(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsValid(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-08-05", true},
	}

	for index, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsValid(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsInvalid(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", true},
		{"0", true},
		{"0000-00-00", true},
		{"00:00:00", true},
		{"0000-00-00 00:00:00", true},

		{"2020-08-05", false},
	}

	for index, test := range tests {
		c := Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsInvalid(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsNow(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    Carbon // 输入值
		expected bool   // 期望值
	}{
		{Parse(""), false},
		{Parse("0"), false},
		{Parse("0000-00-00"), false},
		{Parse("00:00:00"), false},
		{Parse("0000-00-00 00:00:00"), false},

		{Tomorrow(), false},
		{Now(), true},
		{Yesterday(), false},
	}

	for index, test := range tests {
		c := test.input
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsNow(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsFuture(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    Carbon // 输入值
		expected bool   // 期望值
	}{
		{Parse(""), false},
		{Parse("0"), false},
		{Parse("0000-00-00"), false},
		{Parse("00:00:00"), false},
		{Parse("0000-00-00 00:00:00"), false},

		{Tomorrow(), true},
		{Now(), false},
		{Yesterday(), false},
	}

	for index, test := range tests {
		c := test.input
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsFuture(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsPast(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    Carbon // 输入值
		expected bool   // 期望值
	}{
		{Parse(""), false},
		{Parse("0"), false},
		{Parse("0000-00-00"), false},
		{Parse("00:00:00"), false},
		{Parse("0000-00-00 00:00:00"), false},

		{Tomorrow(), false},
		{Now(), false},
		{Yesterday(), true},
	}

	for index, test := range tests {
		c := test.input
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsPast(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsLeapYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2015-01-01", false},
		{"2016-01-01", true},
		{"2017-01-01", false},
		{"2018-01-01", false},
		{"2019-01-01", false},
		{"2020-01-01", true},
		{"2021-01-01", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsLeapYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsLongYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2015-01-01", true},
		{"2016-01-01", false},
		{"2017-01-01", false},
		{"2018-01-01", false},
		{"2019-01-01", false},
		{"2020-01-01", true},
		{"2021-01-01", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsLongYear(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsJanuary(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-01-01", true},
		{"2020-02-01", false},
		{"2020-03-01", false},
		{"2020-04-01", false},
		{"2020-05-01", false},
		{"2020-06-01", false},
		{"2020-07-01", false},
		{"2020-08-01", false},
		{"2020-09-01", false},
		{"2020-10-01", false},
		{"2020-11-01", false},
		{"2020-12-01", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsJanuary(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsFebruary(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-01-01", false},
		{"2020-02-01", true},
		{"2020-03-01", false},
		{"2020-04-01", false},
		{"2020-05-01", false},
		{"2020-06-01", false},
		{"2020-07-01", false},
		{"2020-08-01", false},
		{"2020-09-01", false},
		{"2020-10-01", false},
		{"2020-11-01", false},
		{"2020-12-01", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsFebruary(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsMarch(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-01-01", false},
		{"2020-02-01", false},
		{"2020-03-01", true},
		{"2020-04-01", false},
		{"2020-05-01", false},
		{"2020-06-01", false},
		{"2020-07-01", false},
		{"2020-08-01", false},
		{"2020-09-01", false},
		{"2020-10-01", false},
		{"2020-11-01", false},
		{"2020-12-01", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsMarch(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsApril(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-01-01", false},
		{"2020-02-01", false},
		{"2020-03-01", false},
		{"2020-04-01", true},
		{"2020-05-01", false},
		{"2020-06-01", false},
		{"2020-07-01", false},
		{"2020-08-01", false},
		{"2020-09-01", false},
		{"2020-10-01", false},
		{"2020-11-01", false},
		{"2020-12-01", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsApril(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsMay(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-01-01", false},
		{"2020-02-01", false},
		{"2020-03-01", false},
		{"2020-04-01", false},
		{"2020-05-01", true},
		{"2020-06-01", false},
		{"2020-07-01", false},
		{"2020-08-01", false},
		{"2020-09-01", false},
		{"2020-10-01", false},
		{"2020-11-01", false},
		{"2020-12-01", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsMay(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsJune(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-01-01", false},
		{"2020-02-01", false},
		{"2020-03-01", false},
		{"2020-04-01", false},
		{"2020-05-01", false},
		{"2020-06-01", true},
		{"2020-07-01", false},
		{"2020-08-01", false},
		{"2020-09-01", false},
		{"2020-10-01", false},
		{"2020-11-01", false},
		{"2020-12-01", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsJune(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsJuly(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-01-01", false},
		{"2020-02-01", false},
		{"2020-03-01", false},
		{"2020-04-01", false},
		{"2020-05-01", false},
		{"2020-06-01", false},
		{"2020-07-01", true},
		{"2020-08-01", false},
		{"2020-09-01", false},
		{"2020-10-01", false},
		{"2020-11-01", false},
		{"2020-12-01", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsJuly(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsAugust(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-01-01", false},
		{"2020-02-01", false},
		{"2020-03-01", false},
		{"2020-04-01", false},
		{"2020-05-01", false},
		{"2020-06-01", false},
		{"2020-07-01", false},
		{"2020-08-01", true},
		{"2020-09-01", false},
		{"2020-10-01", false},
		{"2020-11-01", false},
		{"2020-12-01", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsAugust(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsSeptember(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-01-01", false},
		{"2020-02-01", false},
		{"2020-03-01", false},
		{"2020-04-01", false},
		{"2020-05-01", false},
		{"2020-06-01", false},
		{"2020-07-01", false},
		{"2020-08-01", false},
		{"2020-09-01", true},
		{"2020-10-01", false},
		{"2020-11-01", false},
		{"2020-12-01", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsSeptember(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsOctober(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-01-01", false},
		{"2020-02-01", false},
		{"2020-03-01", false},
		{"2020-04-01", false},
		{"2020-05-01", false},
		{"2020-06-01", false},
		{"2020-07-01", false},
		{"2020-08-01", false},
		{"2020-09-01", false},
		{"2020-10-01", true},
		{"2020-11-01", false},
		{"2020-12-01", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsOctober(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsNovember(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-01-01", false},
		{"2020-02-01", false},
		{"2020-03-01", false},
		{"2020-04-01", false},
		{"2020-05-01", false},
		{"2020-06-01", false},
		{"2020-07-01", false},
		{"2020-08-01", false},
		{"2020-09-01", false},
		{"2020-10-01", false},
		{"2020-11-01", true},
		{"2020-12-01", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsNovember(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsDecember(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-01-01", false},
		{"2020-02-01", false},
		{"2020-03-01", false},
		{"2020-04-01", false},
		{"2020-05-01", false},
		{"2020-06-01", false},
		{"2020-07-01", false},
		{"2020-08-01", false},
		{"2020-09-01", false},
		{"2020-10-01", false},
		{"2020-11-01", false},
		{"2020-12-01", true},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsDecember(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsMonday(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-10-05", true},
		{"2020-10-06", false},
		{"2020-10-07", false},
		{"2020-10-08", false},
		{"2020-10-09", false},
		{"2020-10-10", false},
		{"2020-10-11", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsMonday(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsTuesday(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-10-05", false},
		{"2020-10-06", true},
		{"2020-10-07", false},
		{"2020-10-08", false},
		{"2020-10-09", false},
		{"2020-10-10", false},
		{"2020-10-11", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsTuesday(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsWednesday(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-10-05", false},
		{"2020-10-06", false},
		{"2020-10-07", true},
		{"2020-10-08", false},
		{"2020-10-09", false},
		{"2020-10-10", false},
		{"2020-10-11", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsWednesday(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsThursday(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-10-05", false},
		{"2020-10-06", false},
		{"2020-10-07", false},
		{"2020-10-08", true},
		{"2020-10-09", false},
		{"2020-10-10", false},
		{"2020-10-11", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsThursday(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsFriday(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-10-05", false},
		{"2020-10-06", false},
		{"2020-10-07", false},
		{"2020-10-08", false},
		{"2020-10-09", true},
		{"2020-10-10", false},
		{"2020-10-11", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsFriday(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsSaturday(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-10-05", false},
		{"2020-10-06", false},
		{"2020-10-07", false},
		{"2020-10-08", false},
		{"2020-10-09", false},
		{"2020-10-10", true},
		{"2020-10-11", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsSaturday(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsSunday(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-10-05", false},
		{"2020-10-06", false},
		{"2020-10-07", false},
		{"2020-10-08", false},
		{"2020-10-09", false},
		{"2020-10-10", false},
		{"2020-10-11", true},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsSunday(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsWeekday(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-10-05", true},
		{"2020-10-06", true},
		{"2020-10-07", true},
		{"2020-10-08", true},
		{"2020-10-09", true},
		{"2020-10-10", false},
		{"2020-10-11", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsWeekday(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsWeekend(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected bool   // 期望值
	}{
		{"", false},
		{"0", false},
		{"0000-00-00", false},
		{"00:00:00", false},
		{"0000-00-00 00:00:00", false},

		{"2020-10-05", false},
		{"2020-10-06", false},
		{"2020-10-07", false},
		{"2020-10-08", false},
		{"2020-10-09", false},
		{"2020-10-10", true},
		{"2020-10-11", true},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsWeekend(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsYesterday(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    Carbon // 输入值
		expected bool   // 期望值
	}{
		{NewCarbon(), false},
		{Now(), false},
		{Yesterday(), true},
		{Tomorrow(), false},
	}

	for index, test := range tests {
		c := test.input
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsYesterday(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsToday(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    Carbon // 输入值
		expected bool   // 期望值
	}{
		{NewCarbon(), false},
		{Now(), true},
		{Yesterday(), false},
		{Tomorrow(), false},
	}

	for index, test := range tests {
		c := test.input
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsToday(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsTomorrow(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    Carbon // 输入值
		expected bool   // 期望值
	}{
		{NewCarbon(), false},
		{Now(), false},
		{Yesterday(), false},
		{Tomorrow(), true},
	}

	for index, test := range tests {
		c := test.input
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsTomorrow(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsSameCentury(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input1   Carbon // 输入值
		input2   Carbon // 输入值
		expected bool   // 期望值
	}{
		{Parse(""), Parse(""), false},
		{Parse("2020-08-05"), Parse("3020-08-05"), false},
		{Parse("2020-08-05"), Parse("2099-08-05"), true},
	}

	for index, test := range tests {
		assert.Nil(test.input1.Error)
		assert.Nil(test.input2.Error)
		assert.Equal(test.expected, test.input1.IsSameCentury(test.input2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsSameDecade(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input1   Carbon // 输入值
		input2   Carbon // 输入值
		expected bool   // 期望值
	}{
		{Parse(""), Parse(""), false},
		{Parse("2020-08-05"), Parse("2030-08-05"), false},
		{Parse("2020-08-05"), Parse("2021-08-05"), true},
		{Parse("2020-01-01"), Parse("2120-01-31"), true},
	}

	for index, test := range tests {
		assert.Nil(test.input1.Error)
		assert.Nil(test.input2.Error)
		assert.Equal(test.expected, test.input1.IsSameDecade(test.input2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsSameYear(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input1   Carbon // 输入值
		input2   Carbon // 输入值
		expected bool   // 期望值
	}{
		{Parse(""), Parse(""), false},
		{Parse("2020-08-05"), Parse("2021-08-05"), false},
		{Parse("2020-01-01"), Parse("2020-12-31"), true},
	}

	for index, test := range tests {
		assert.Nil(test.input1.Error)
		assert.Nil(test.input2.Error)
		assert.Equal(test.expected, test.input1.IsSameYear(test.input2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsSameQuarter(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input1   Carbon // 输入值
		input2   Carbon // 输入值
		expected bool   // 期望值
	}{
		{Parse(""), Parse(""), false},
		{Parse("2020-08-05"), Parse("2020-01-05"), false},
		{Parse("2020-01-01"), Parse("2020-01-31"), true},
	}

	for index, test := range tests {
		assert.Nil(test.input1.Error)
		assert.Nil(test.input2.Error)
		assert.Equal(test.expected, test.input1.IsSameQuarter(test.input2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsSameMonth(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input1   Carbon // 输入值
		input2   Carbon // 输入值
		expected bool   // 期望值
	}{
		{Parse(""), Parse(""), false},
		{Parse("2020-08-05"), Parse("2021-08-05"), false},
		{Parse("2020-01-01"), Parse("2020-01-31"), true},
	}

	for index, test := range tests {
		assert.Nil(test.input1.Error)
		assert.Nil(test.input2.Error)
		assert.Equal(test.expected, test.input1.IsSameMonth(test.input2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsSameDay(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input1   Carbon // 输入值
		input2   Carbon // 输入值
		expected bool   // 期望值
	}{
		{Parse(""), Parse(""), false},
		{Parse("2020-08-05 13:14:15"), Parse("2021-08-05 13:14:15"), false},
		{Parse("2020-08-05 00:00:00"), Parse("2020-08-05 13:14:15"), true},
	}

	for index, test := range tests {
		assert.Nil(test.input1.Error)
		assert.Nil(test.input2.Error)
		assert.Equal(test.expected, test.input1.IsSameDay(test.input2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsSameHour(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input1   Carbon // 输入值
		input2   Carbon // 输入值
		expected bool   // 期望值
	}{
		{Parse(""), Parse(""), false},
		{Parse("2020-08-05 13:14:15"), Parse("2021-08-05 13:14:15"), false},
		{Parse("2020-08-05 13:00:00"), Parse("2020-08-05 13:14:15"), true},
	}

	for index, test := range tests {
		assert.Nil(test.input1.Error)
		assert.Nil(test.input2.Error)
		assert.Equal(test.expected, test.input1.IsSameHour(test.input2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsSameMinute(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input1   Carbon // 输入值
		input2   Carbon // 输入值
		expected bool   // 期望值
	}{
		{Parse(""), Parse(""), false},
		{Parse("2020-08-05 13:14:15"), Parse("2021-08-05 13:14:15"), false},
		{Parse("2020-08-05 13:14:00"), Parse("2020-08-05 13:14:15"), true},
	}

	for index, test := range tests {
		assert.Nil(test.input1.Error)
		assert.Nil(test.input2.Error)
		assert.Equal(test.expected, test.input1.IsSameMinute(test.input2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsSameSecond(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input1   Carbon // 输入值
		input2   Carbon // 输入值
		expected bool   // 期望值
	}{
		{Parse(""), Parse(""), false},
		{Parse("2020-08-05 13:14:15"), Parse("2021-08-05 13:14:15"), false},
		{Parse("2020-08-05 13:14:15"), Parse("2020-08-05 13:14:15"), true},
	}

	for index, test := range tests {
		assert.Nil(test.input1.Error)
		assert.Nil(test.input2.Error)
		assert.Equal(test.expected, test.input1.IsSameSecond(test.input2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Compare(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		param1   string // 输入参数1
		param2   string // 输入参数2
		expected bool   // 期望值
	}{
		{"2020-08-05", ">", "2020-08-04", true},
		{"2020-08-05", "<", "2020-08-04", false},
		{"2020-08-05", "<", "2020-08-06", true},
		{"2020-08-05", ">", "2020-08-06", false},
		{"2020-08-05", "=", "2020-08-05", true},
		{"2020-08-05", ">=", "2020-08-05", true},
		{"2020-08-05", "<=", "2020-08-05", true},
		{"2020-08-05", "!=", "2020-08-05", false},
		{"2020-08-05", "<>", "2020-08-05", false},
		{"2020-08-05", "!=", "2020-08-04", true},
		{"2020-08-05", "<>", "2020-08-04", true},
		{"2020-08-05", "+", "2020-08-04", false},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param2)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.Compare(test.param1, c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Gt(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		param    string // 参数值
		expected bool   // 期望值
	}{
		{"2020-08-05", "2020-08-05", false},
		{"2020-08-05", "2020-08-04", true},
		{"2020-08-05", "2020-08-06", false},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.Gt(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Lt(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		param    string // 参数值
		expected bool   // 期望值
	}{
		{"2020-08-05", "2020-08-05", false},
		{"2020-08-05", "2020-08-04", false},
		{"2020-08-05", "2020-08-06", true},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.Lt(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Eq(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		param    string // 参数值
		expected bool   // 期望值
	}{
		{"2020-08-05", "2020-08-05", true},
		{"2020-08-05", "2020-08-04", false},
		{"2020-08-05", "2020-08-06", false},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.Eq(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Ne(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		param    string // 参数值
		expected bool   // 期望值
	}{
		{"2020-08-05", "2020-08-05", false},
		{"2020-08-05", "2020-08-04", true},
		{"2020-08-05", "2020-08-06", true},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.Ne(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Gte(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		param    string // 参数值
		expected bool   // 期望值
	}{
		{"2020-08-05", "2020-08-05", true},
		{"2020-08-05", "2020-08-04", true},
		{"2020-08-05", "2020-08-06", false},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.Gte(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Lte(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		param    string // 参数值
		expected bool   // 期望值
	}{
		{"2020-08-05", "2020-08-05", true},
		{"2020-08-05", "2020-08-04", false},
		{"2020-08-05", "2020-08-06", true},
	}

	for index, test := range tests {
		c1, c2 := Parse(test.input), Parse(test.param)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Equal(test.expected, c1.Lte(c2), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_Between(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		param1   string // 输入参数1
		param2   string // 输入参数2
		expected bool   // 期望值
	}{
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", "2020-08-05 13:14:15", false},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", "2020-08-06 13:14:15", false},
		{"2020-08-05 13:14:15", "2020-08-04 13:14:15", "2020-08-05 13:14:15", false},
		{"2020-08-05 13:14:15", "2020-08-04 13:14:15", "2020-08-06 13:14:15", true},
	}

	for index, test := range tests {
		c1, c2, c3 := Parse(test.input), Parse(test.param1), Parse(test.param2)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Nil(c3.Error)
		assert.Equal(test.expected, c1.Between(c2, c3), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_BetweenIncludedStart(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		param1   string // 输入参数1
		param2   string // 输入参数2
		expected bool   // 期望值
	}{
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", "2020-08-05 13:14:15", false},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", "2020-08-06 13:14:15", true},
		{"2020-08-05 13:14:15", "2020-08-04 13:14:15", "2020-08-05 13:14:15", false},
		{"2020-08-05 13:14:15", "2020-08-04 13:14:15", "2020-08-06 13:14:15", true},
	}

	for index, test := range tests {
		c1, c2, c3 := Parse(test.input), Parse(test.param1), Parse(test.param2)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Nil(c3.Error)
		assert.Equal(test.expected, c1.BetweenIncludedStart(c2, c3), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_BetweenIncludedEnd(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		param1   string // 输入参数1
		param2   string // 输入参数2
		expected bool   // 期望值
	}{
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", "2020-08-05 13:14:15", false},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", "2020-08-06 13:14:15", false},
		{"2020-08-05 13:14:15", "2020-08-04 13:14:15", "2020-08-05 13:14:15", true},
		{"2020-08-05 13:14:15", "2020-08-04 13:14:15", "2020-08-06 13:14:15", true},
	}

	for index, test := range tests {
		c1, c2, c3 := Parse(test.input), Parse(test.param1), Parse(test.param2)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Nil(c3.Error)
		assert.Equal(test.expected, c1.BetweenIncludedEnd(c2, c3), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_BetweenIncludedBoth(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		param1   string // 输入参数1
		param2   string // 输入参数2
		expected bool   // 期望值
	}{
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", "2020-08-05 13:14:15", true},
		{"2020-08-05 13:14:15", "2020-08-05 13:14:15", "2020-08-06 13:14:15", true},
		{"2020-08-05 13:14:15", "2020-08-04 13:14:15", "2020-08-05 13:14:15", true},
		{"2020-08-05 13:14:15", "2020-08-04 13:14:15", "2020-08-06 13:14:15", true},
		{"2020-08-05 13:14:15", "2020-08-06 13:14:15", "2020-08-06 13:14:15", false},
	}

	for index, test := range tests {
		c1, c2, c3 := Parse(test.input), Parse(test.param1), Parse(test.param2)
		assert.Nil(c1.Error)
		assert.Nil(c2.Error)
		assert.Nil(c3.Error)
		assert.Equal(test.expected, c1.BetweenIncludedBoth(c2, c3), "Current test index is "+strconv.Itoa(index))
	}
}

func TestError_Comparer(t *testing.T) {
	time1, time2, time3, operator := "2020-13-50", "xxx", "xxx", ">"
	assert.True(t, Parse(time1).IsZero(), "It should catch an exception in IsZero()")
	assert.True(t, Parse(time1).IsInvalid(), "It should catch an exception in IsInvalid()")
	assert.True(t, Parse(time1).IsInvalid(), "It should catch an exception in IsInvalid()")
	assert.True(t, Parse(time1).Ne(Parse(time2)), "It should catch an exception in Ne()")
	assert.False(t, Parse(time1).Compare(operator, Parse(time2)), "It should catch an exception in Compare()")
	assert.False(t, Parse(time1).Gt(Parse(time2)), "It should catch an exception in Gt()")
	assert.False(t, Parse(time1).Lt(Parse(time2)), "It should catch an exception in Lt()")
	assert.False(t, Parse(time1).Eq(Parse(time2)), "It should catch an exception in Eq()")
	assert.False(t, Parse(time1).Gte(Parse(time2)), "It should catch an exception in Gte()")
	assert.False(t, Parse(time1).Lte(Parse(time2)), "It should catch an exception in Lte()")
	assert.False(t, Parse(time1).Between(Parse(time2), Parse(time3)), "It should catch an exception in Between()")
	assert.False(t, Parse(time1).BetweenIncludedStart(Parse(time2), Parse(time3)), "It should catch an exception in BetweenIncludedStart()")
	assert.False(t, Parse(time1).BetweenIncludedEnd(Parse(time2), Parse(time3)), "It should catch an exception in BetweenIncludedEnd()")
	assert.False(t, Parse(time1).BetweenIncludedBoth(Parse(time2), Parse(time3)), "It should catch an exception in BetweenIncludedBoth()")
}
