package carbon

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_Season(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-01-05", "Winter"},
		{"2020-02-05", "Winter"},
		{"2020-03-05", "Spring"},
		{"2020-04-05", "Spring"},
		{"2020-05-05", "Spring"},
		{"2020-06-05", "Summer"},
		{"2020-07-05", "Summer"},
		{"2020-08-05", "Summer"},
		{"2020-09-05", "Autumn"},
		{"2020-10-05", "Autumn"},
		{"2020-11-05", "Autumn"},
		{"2020-12-05", "Winter"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.Season(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_StartOfSeason(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-01-15 00:00:00", "2019-12-01 00:00:00"},
		{"2020-02-15 00:00:00", "2019-12-01 00:00:00"},
		{"2020-03-15 00:00:00", "2020-03-01 00:00:00"},
		{"2020-04-15 23:59:59", "2020-03-01 00:00:00"},
		{"2020-05-15 23:59:59", "2020-03-01 00:00:00"},
		{"2020-06-15 23:59:59", "2020-06-01 00:00:00"},
		{"2020-07-15 23:59:59", "2020-06-01 00:00:00"},
		{"2020-08-15 13:14:15", "2020-06-01 00:00:00"},
		{"2020-09-15 13:14:15", "2020-09-01 00:00:00"},
		{"2020-10-15", "2020-09-01 00:00:00"},
		{"2020-11-15", "2020-09-01 00:00:00"},
		{"2020-12-15", "2020-12-01 00:00:00"},
		{"2021-01-15", "2020-12-01 00:00:00"},
		{"2021-01-15", "2020-12-01 00:00:00"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).StartOfSeason()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_EndOfSeason(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-01-15 00:00:00", "2020-02-29 23:59:59"},
		{"2020-02-15 00:00:00", "2020-02-29 23:59:59"},
		{"2020-03-15 00:00:00", "2020-05-31 23:59:59"},
		{"2020-04-15 23:59:59", "2020-05-31 23:59:59"},
		{"2020-05-15 23:59:59", "2020-05-31 23:59:59"},
		{"2020-06-15 23:59:59", "2020-08-31 23:59:59"},
		{"2020-07-15 23:59:59", "2020-08-31 23:59:59"},
		{"2020-08-15 13:14:15", "2020-08-31 23:59:59"},
		{"2020-09-15 13:14:15", "2020-11-30 23:59:59"},
		{"2020-10-15", "2020-11-30 23:59:59"},
		{"2020-11-15", "2020-11-30 23:59:59"},
		{"2020-12-15", "2021-02-28 23:59:59"},
		{"2021-01-15", "2021-02-28 23:59:59"},
		{"2021-02-15", "2021-02-28 23:59:59"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input).EndOfSeason()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsSpring(t *testing.T) {
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
		{"2020-04-01", true},
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
		assert.Equal(test.expected, c.IsSpring(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsSummer(t *testing.T) {
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
		{"2020-07-01", true},
		{"2020-08-01", true},
		{"2020-09-01", false},
		{"2020-10-01", false},
		{"2020-11-01", false},
		{"2020-12-01", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsSummer(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsAutumn(t *testing.T) {
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
		{"2020-10-01", true},
		{"2020-11-01", true},
		{"2020-12-01", false},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsAutumn(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_IsWinter(t *testing.T) {
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
		{"2020-12-01", true},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.IsWinter(), "Current test index is "+strconv.Itoa(index))
	}
}
