package carbon

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_AddDuration(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		duration string // 输入参数
		expected string // 期望值
	}{
		{"", "10h", ""},
		{"0", "10h", ""},
		{"0000-00-00", "10h", ""},
		{"00:00:00", "10h", ""},
		{"0000-00-00 00:00:00", "10h", ""},

		{"2020-01-01 13:14:15", "10h", "2020-01-01 23:14:15"},
		{"2020-01-01 13:14:15", "10.5h", "2020-01-01 23:44:15"},

		{"2020-01-01 13:14:15", "10m", "2020-01-01 13:24:15"},
		{"2020-01-01 13:14:15", "10.5m", "2020-01-01 13:24:45"},

		{"2020-01-01 13:14:15", "10s", "2020-01-01 13:14:25"},
		{"2020-01-01 13:14:15", "10.5s", "2020-01-01 13:14:25"},
	}

	for index, test := range tests {
		c := Parse(test.input).AddDuration(test.duration)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SubDuration(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		duration string // 输入参数
		expected string // 期望值
	}{
		{"", "10h", ""},
		{"0", "10h", ""},
		{"0000-00-00", "10h", ""},
		{"00:00:00", "10h", ""},
		{"0000-00-00 00:00:00", "10h", ""},

		{"2020-01-01 13:14:15", "10h", "2020-01-01 03:14:15"},
		{"2020-01-01 13:14:15", "10.5h", "2020-01-01 02:44:15"},

		{"2020-01-01 13:14:15", "10m", "2020-01-01 13:04:15"},
		{"2020-01-01 13:14:15", "10.5m", "2020-01-01 13:03:45"},

		{"2020-01-01 13:14:15", "10s", "2020-01-01 13:14:05"},
		{"2020-01-01 13:14:15", "10.5s", "2020-01-01 13:14:04"},
	}

	for index, test := range tests {
		c := Parse(test.input).SubDuration(test.duration)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_AddCenturies(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input     string // 输入值
		centuries int    // 输入参数
		expected  string // 期望值
	}{
		{"", 3, ""},
		{"0", 3, ""},
		{"0000-00-00", 3, ""},
		{"00:00:00", 3, ""},
		{"0000-00-00 00:00:00", 3, ""},

		{"2020-01-01", 3, "2320-01-01"},
		{"2020-01-31", 3, "2320-01-31"},
		{"2020-02-01", 3, "2320-02-01"},
		{"2020-02-28", 3, "2320-02-28"},
		{"2020-02-29", 3, "2320-02-29"},
	}

	for index, test := range tests {
		c := Parse(test.input).AddCenturies(test.centuries)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_AddCenturiesNoOverflow(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input     string // 输入值
		centuries int    // 输入参数
		expected  string // 期望值
	}{
		{"", 3, ""},
		{"0", 3, ""},
		{"0000-00-00", 3, ""},
		{"00:00:00", 3, ""},
		{"0000-00-00 00:00:00", 3, ""},

		{"2020-01-01", 3, "2320-01-01"},
		{"2020-01-31", 3, "2320-01-31"},
		{"2020-02-01", 3, "2320-02-01"},
		{"2020-02-28", 3, "2320-02-28"},
		{"2020-02-29", 3, "2320-02-29"},
	}

	for index, test := range tests {
		c := Parse(test.input).AddCenturiesNoOverflow(test.centuries)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SubCenturies(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input     string // 输入值
		centuries int
		expected  string // 期望值
	}{
		{"", 3, ""},
		{"0", 3, ""},
		{"0000-00-00", 3, ""},
		{"00:00:00", 3, ""},
		{"0000-00-00 00:00:00", 3, ""},

		{"2020-01-01", 3, "1720-01-01"},
		{"2020-01-31", 3, "1720-01-31"},
		{"2020-02-01", 3, "1720-02-01"},
		{"2020-02-28", 3, "1720-02-28"},
		{"2020-02-29", 3, "1720-02-29"},
	}

	for index, test := range tests {
		c := Parse(test.input).SubCenturies(test.centuries)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SubCenturiesNoOverflow(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input     string // 输入值
		centuries int
		expected  string // 期望值
	}{
		{"", 3, ""},
		{"0", 3, ""},
		{"0000-00-00", 3, ""},
		{"00:00:00", 3, ""},
		{"0000-00-00 00:00:00", 3, ""},

		{"2020-01-01", 3, "1720-01-01"},
		{"2020-01-31", 3, "1720-01-31"},
		{"2020-02-01", 3, "1720-02-01"},
		{"2020-02-28", 3, "1720-02-28"},
		{"2020-02-29", 3, "1720-02-29"},
	}

	for index, test := range tests {
		c := Parse(test.input).SubCenturiesNoOverflow(test.centuries)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_AddCentury(t *testing.T) {
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

		{"2020-01-01", "2120-01-01"},
		{"2020-01-31", "2120-01-31"},
		{"2020-02-01", "2120-02-01"},
		{"2020-02-28", "2120-02-28"},
		{"2020-02-29", "2120-02-29"},
	}

	for index, test := range tests {
		c := Parse(test.input).AddCentury()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_AddCenturyNoOverflow(t *testing.T) {
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

		{"2020-01-01", "2120-01-01"},
		{"2020-01-31", "2120-01-31"},
		{"2020-02-01", "2120-02-01"},
		{"2020-02-28", "2120-02-28"},
		{"2020-02-29", "2120-02-29"},
	}

	for index, test := range tests {
		c := Parse(test.input).AddCenturyNoOverflow()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SubCentury(t *testing.T) {
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

		{"2020-01-01", "1920-01-01"},
		{"2020-01-31", "1920-01-31"},
		{"2020-02-01", "1920-02-01"},
		{"2020-02-28", "1920-02-28"},
		{"2020-02-29", "1920-02-29"},
	}

	for index, test := range tests {
		c := Parse(test.input).SubCentury()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SubCenturyNoOverflow(t *testing.T) {
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

		{"2020-01-01", "1920-01-01"},
		{"2020-01-31", "1920-01-31"},
		{"2020-02-01", "1920-02-01"},
		{"2020-02-28", "1920-02-28"},
		{"2020-02-29", "1920-02-29"},
	}

	for index, test := range tests {
		c := Parse(test.input).SubCenturyNoOverflow()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_AddDecades(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		decades  int    // 输入参数
		expected string // 期望值
	}{
		{"", 3, ""},
		{"0", 3, ""},
		{"0000-00-00", 3, ""},
		{"00:00:00", 3, ""},
		{"0000-00-00 00:00:00", 3, ""},

		{"2020-01-01", 3, "2050-01-01"},
		{"2020-01-31", 3, "2050-01-31"},
		{"2020-02-01", 3, "2050-02-01"},
		{"2020-02-28", 3, "2050-02-28"},
		{"2020-02-29", 3, "2050-03-01"},
	}

	for index, test := range tests {
		c := Parse(test.input).AddDecades(test.decades)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_AddDecadesNoOverflow(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input     string // 输入值
		centuries int    // 输入参数
		expected  string // 期望值
	}{
		{"", 3, ""},
		{"0", 3, ""},
		{"0000-00-00", 3, ""},
		{"00:00:00", 3, ""},
		{"0000-00-00 00:00:00", 3, ""},

		{"2020-01-01", 3, "2050-01-01"},
		{"2020-01-31", 3, "2050-01-31"},
		{"2020-02-01", 3, "2050-02-01"},
		{"2020-02-28", 3, "2050-02-28"},
		{"2020-02-29", 3, "2050-02-28"},
	}

	for index, test := range tests {
		c := Parse(test.input).AddDecadesNoOverflow(test.centuries)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_AddDecade(t *testing.T) {
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

		{"2020-01-01", "2030-01-01"},
		{"2020-01-31", "2030-01-31"},
		{"2020-02-01", "2030-02-01"},
		{"2020-02-28", "2030-02-28"},
		{"2020-02-29", "2030-03-01"},
	}

	for index, test := range tests {
		c := Parse(test.input).AddDecade()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_AddDecadeNoOverflow(t *testing.T) {
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

		{"2020-01-01", "2030-01-01"},
		{"2020-01-31", "2030-01-31"},
		{"2020-02-01", "2030-02-01"},
		{"2020-02-28", "2030-02-28"},
		{"2020-02-29", "2030-02-28"},
	}

	for index, test := range tests {
		c := Parse(test.input).AddDecadeNoOverflow()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SubDecades(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		decades  int    // 输入参数
		expected string // 期望值
	}{
		{"", 3, ""},
		{"0", 3, ""},
		{"0000-00-00", 3, ""},
		{"00:00:00", 3, ""},
		{"0000-00-00 00:00:00", 3, ""},

		{"2020-01-01", 3, "1990-01-01"},
		{"2020-01-31", 3, "1990-01-31"},
		{"2020-02-01", 3, "1990-02-01"},
		{"2020-02-28", 3, "1990-02-28"},
		{"2020-02-29", 3, "1990-03-01"},
	}

	for index, test := range tests {
		c := Parse(test.input).SubDecades(test.decades)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SubDecadesNoOverflow(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input     string // 输入值
		centuries int    // 输入参数
		expected  string // 期望值
	}{
		{"", 3, ""},
		{"0", 3, ""},
		{"0000-00-00", 3, ""},
		{"00:00:00", 3, ""},
		{"0000-00-00 00:00:00", 3, ""},

		{"2020-01-01", 3, "1990-01-01"},
		{"2020-01-31", 3, "1990-01-31"},
		{"2020-02-01", 3, "1990-02-01"},
		{"2020-02-28", 3, "1990-02-28"},
		{"2020-02-29", 3, "1990-02-28"},
	}

	for index, test := range tests {
		c := Parse(test.input).SubDecadesNoOverflow(test.centuries)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SubDecade(t *testing.T) {
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

		{"2020-01-01", "2010-01-01"},
		{"2020-01-31", "2010-01-31"},
		{"2020-02-01", "2010-02-01"},
		{"2020-02-28", "2010-02-28"},
		{"2020-02-29", "2010-03-01"},
	}

	for index, test := range tests {
		c := Parse(test.input).SubDecade()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SubDecadeNoOverflow(t *testing.T) {
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

		{"2020-01-01", "2010-01-01"},
		{"2020-01-31", "2010-01-31"},
		{"2020-02-01", "2010-02-01"},
		{"2020-02-28", "2010-02-28"},
		{"2020-02-29", "2010-02-28"},
	}

	for index, test := range tests {
		c := Parse(test.input).SubDecadeNoOverflow()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_AddYears(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		years    int    // 输入参数
		expected string // 期望值
	}{
		{"", 3, ""},
		{"0", 3, ""},
		{"0000-00-00", 3, ""},
		{"00:00:00", 3, ""},
		{"0000-00-00 00:00:00", 3, ""},

		{"2020-01-01", 3, "2023-01-01"},
		{"2020-01-31", 3, "2023-01-31"},
		{"2020-02-01", 3, "2023-02-01"},
		{"2020-02-28", 3, "2023-02-28"},
		{"2020-02-29", 3, "2023-03-01"},
	}

	for index, test := range tests {
		c := Parse(test.input).AddYears(test.years)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_AddYearsNoOverflow(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		years    int    // 输入参数
		expected string // 期望值
	}{
		{"", 3, ""},
		{"0", 3, ""},
		{"0000-00-00", 3, ""},
		{"00:00:00", 3, ""},
		{"0000-00-00 00:00:00", 3, ""},

		{"2020-01-01", 3, "2023-01-01"},
		{"2020-01-31", 3, "2023-01-31"},
		{"2020-02-01", 3, "2023-02-01"},
		{"2020-02-28", 3, "2023-02-28"},
		{"2020-02-29", 3, "2023-02-28"},
	}

	for index, test := range tests {
		c := Parse(test.input).AddYearsNoOverflow(test.years)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SubYears(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		years    int
		expected string // 期望值
	}{
		{"", 3, ""},
		{"0", 3, ""},
		{"0000-00-00", 3, ""},
		{"00:00:00", 3, ""},
		{"0000-00-00 00:00:00", 3, ""},

		{"2020-01-01", 3, "2017-01-01"},
		{"2020-01-31", 3, "2017-01-31"},
		{"2020-02-01", 3, "2017-02-01"},
		{"2020-02-28", 3, "2017-02-28"},
		{"2020-02-29", 3, "2017-03-01"},
	}

	for index, test := range tests {
		c := Parse(test.input).SubYears(test.years)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SubYearsNoOverflow(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		years    int
		expected string // 期望值
	}{
		{"", 3, ""},
		{"0", 3, ""},
		{"0000-00-00", 3, ""},
		{"00:00:00", 3, ""},
		{"0000-00-00 00:00:00", 3, ""},

		{"2020-01-01", 3, "2017-01-01"},
		{"2020-01-31", 3, "2017-01-31"},
		{"2020-02-01", 3, "2017-02-01"},
		{"2020-02-28", 3, "2017-02-28"},
		{"2020-02-29", 3, "2017-02-28"},
	}

	for index, test := range tests {
		c := Parse(test.input).SubYearsNoOverflow(test.years)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_AddYear(t *testing.T) {
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

		{"2020-01-01", "2021-01-01"},
		{"2020-01-31", "2021-01-31"},
		{"2020-02-01", "2021-02-01"},
		{"2020-02-28", "2021-02-28"},
		{"2020-02-29", "2021-03-01"},
	}

	for index, test := range tests {
		c := Parse(test.input).AddYear()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_AddYearNoOverflow(t *testing.T) {
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

		{"2020-01-01", "2021-01-01"},
		{"2020-01-31", "2021-01-31"},
		{"2020-02-01", "2021-02-01"},
		{"2020-02-28", "2021-02-28"},
		{"2020-02-29", "2021-02-28"},
	}

	for index, test := range tests {
		c := Parse(test.input).AddYearNoOverflow()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SubYear(t *testing.T) {
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

		{"2020-01-01", "2019-01-01"},
		{"2020-01-31", "2019-01-31"},
		{"2020-02-01", "2019-02-01"},
		{"2020-02-28", "2019-02-28"},
		{"2020-02-29", "2019-03-01"},
	}

	for index, test := range tests {
		c := Parse(test.input).SubYear()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SubYearNoOverflow(t *testing.T) {
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

		{"2020-01-01", "2019-01-01"},
		{"2020-01-31", "2019-01-31"},
		{"2020-02-01", "2019-02-01"},
		{"2020-02-28", "2019-02-28"},
		{"2020-02-29", "2019-02-28"},
	}

	for index, test := range tests {
		c := Parse(test.input).SubYearNoOverflow()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_AddQuarters(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		quarters int
		expected string // 期望值
	}{
		{"", 2, ""},
		{"0", 2, ""},
		{"0000-00-00", 2, ""},
		{"00:00:00", 2, ""},
		{"0000-00-00 00:00:00", 2, ""},

		{"2019-08-01", 2, "2020-02-01"},
		{"2019-08-31", 2, "2020-03-02"},
		{"2020-01-01", 2, "2020-07-01"},
		{"2020-02-28", 2, "2020-08-28"},
		{"2020-02-29", 2, "2020-08-29"},
		{"2020-08-05", 2, "2021-02-05"},
		{"2020-08-31", 2, "2021-03-03"},
	}

	for index, test := range tests {
		c := Parse(test.input).AddQuarters(test.quarters)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_AddQuartersNoOverflow(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		quarters int
		expected string // 期望值
	}{
		{"", 2, ""},
		{"0", 2, ""},
		{"0000-00-00", 2, ""},
		{"00:00:00", 2, ""},
		{"0000-00-00 00:00:00", 2, ""},

		{"2019-08-01", 2, "2020-02-01"},
		{"2019-08-31", 2, "2020-02-29"},
		{"2020-01-01", 2, "2020-07-01"},
		{"2020-02-28", 2, "2020-08-28"},
		{"2020-02-29", 2, "2020-08-29"},
		{"2020-08-05", 2, "2021-02-05"},
		{"2020-08-31", 2, "2021-02-28"},
	}

	for index, test := range tests {
		c := Parse(test.input).AddQuartersNoOverflow(test.quarters)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SubQuarters(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		quarters int
		expected string // 期望值
	}{
		{"", 2, ""},
		{"0", 2, ""},
		{"0000-00-00", 2, ""},
		{"00:00:00", 2, ""},
		{"0000-00-00 00:00:00", 2, ""},

		{"2019-08-01", 2, "2019-02-01"},
		{"2019-08-31", 2, "2019-03-03"},
		{"2020-01-01", 2, "2019-07-01"},
		{"2020-02-28", 2, "2019-08-28"},
		{"2020-02-29", 2, "2019-08-29"},
		{"2020-08-05", 2, "2020-02-05"},
		{"2020-08-31", 2, "2020-03-02"},
	}

	for index, test := range tests {
		c := Parse(test.input).SubQuarters(test.quarters)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SubQuartersNoOverflow(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		quarters int
		expected string // 期望值
	}{
		{"", 2, ""},
		{"0", 2, ""},
		{"0000-00-00", 2, ""},
		{"00:00:00", 2, ""},
		{"0000-00-00 00:00:00", 2, ""},

		{"2019-08-01", 2, "2019-02-01"},
		{"2019-08-31", 2, "2019-02-28"},
		{"2020-01-01", 2, "2019-07-01"},
		{"2020-02-28", 2, "2019-08-28"},
		{"2020-02-29", 2, "2019-08-29"},
		{"2020-08-05", 2, "2020-02-05"},
		{"2020-08-31", 2, "2020-02-29"},
	}

	for index, test := range tests {
		c := Parse(test.input).SubQuartersNoOverflow(test.quarters)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_AddQuarter(t *testing.T) {
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

		{"2019-11-01", "2020-02-01"},
		{"2019-11-30", "2020-03-01"},
		{"2020-02-28", "2020-05-28"},
		{"2020-02-29", "2020-05-29"},
		{"2020-08-31", "2020-12-01"},
		{"2020-11-01", "2021-02-01"},
		{"2020-11-30", "2021-03-02"},
	}

	for index, test := range tests {
		c := Parse(test.input).AddQuarter()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_AddQuarterNoOverflow(t *testing.T) {
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

		{"2019-11-01", "2020-02-01"},
		{"2019-11-30", "2020-02-29"},
		{"2020-02-28", "2020-05-28"},
		{"2020-02-29", "2020-05-29"},
		{"2020-08-31", "2020-11-30"},
		{"2020-11-01", "2021-02-01"},
		{"2020-11-30", "2021-02-28"},
	}

	for index, test := range tests {
		c := Parse(test.input).AddQuarterNoOverflow()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SubQuarter(t *testing.T) {
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

		{"2019-04-01", "2019-01-01"},
		{"2019-04-30", "2019-01-30"},
		{"2020-04-01", "2020-01-01"},
		{"2020-04-30", "2020-01-30"},
		{"2020-05-01", "2020-02-01"},
		{"2020-05-31", "2020-03-02"},
		{"2020-08-05", "2020-05-05"},
	}

	for index, test := range tests {
		c := Parse(test.input).SubQuarter()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SubQuarterNoOverflow(t *testing.T) {
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

		{"2019-04-01", "2019-01-01"},
		{"2019-04-30", "2019-01-30"},
		{"2020-04-01", "2020-01-01"},
		{"2020-04-30", "2020-01-30"},
		{"2020-05-01", "2020-02-01"},
		{"2020-05-31", "2020-02-29"},
		{"2020-08-05", "2020-05-05"},
	}

	for index, test := range tests {
		c := Parse(test.input).SubQuarterNoOverflow()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_AddMonths(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		months   int
		expected string // 期望值
	}{
		{"", 3, ""},
		{"0", 3, ""},
		{"0000-00-00", 3, ""},
		{"00:00:00", 3, ""},
		{"0000-00-00 00:00:00", 3, ""},

		{"2020-01-01", 3, "2020-04-01"},
		{"2020-01-31", 3, "2020-05-01"},
		{"2020-02-01", 3, "2020-05-01"},
		{"2020-02-28", 3, "2020-05-28"},
		{"2020-02-29", 3, "2020-05-29"},
	}

	for index, test := range tests {
		c := Parse(test.input).AddMonths(test.months)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_AddMonthsNoOverflow(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		months   int
		expected string // 期望值
	}{
		{"", 3, ""},
		{"0", 3, ""},
		{"0000-00-00", 3, ""},
		{"00:00:00", 3, ""},
		{"0000-00-00 00:00:00", 3, ""},

		{"2020-01-01", 3, "2020-04-01"},
		{"2020-01-31", 3, "2020-04-30"},
		{"2020-02-01", 3, "2020-05-01"},
		{"2020-02-28", 3, "2020-05-28"},
		{"2020-02-29", 3, "2020-05-29"},
	}

	for index, test := range tests {
		c := Parse(test.input).AddMonthsNoOverflow(test.months)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SubMonths(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		months   int
		expected string // 期望值
	}{
		{"", 3, ""},
		{"0", 3, ""},
		{"0000-00-00", 3, ""},
		{"00:00:00", 3, ""},
		{"0000-00-00 00:00:00", 3, ""},

		{"2020-01-01", 3, "2019-10-01"},
		{"2020-01-31", 3, "2019-10-31"},
		{"2020-02-01", 3, "2019-11-01"},
		{"2020-02-28", 3, "2019-11-28"},
		{"2020-02-29", 3, "2019-11-29"},
	}

	for index, test := range tests {
		c := Parse(test.input).SubMonths(test.months)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SubMonthsNoOverflow(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		months   int
		expected string // 期望值
	}{
		{"", 3, ""},
		{"0", 3, ""},
		{"0000-00-00", 3, ""},
		{"00:00:00", 3, ""},
		{"0000-00-00 00:00:00", 3, ""},

		{"2020-01-01", 3, "2019-10-01"},
		{"2020-01-31", 3, "2019-10-31"},
		{"2020-02-01", 3, "2019-11-01"},
		{"2020-02-28", 3, "2019-11-28"},
		{"2020-02-29", 3, "2019-11-29"},
	}

	for index, test := range tests {
		c := Parse(test.input).SubMonthsNoOverflow(test.months)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_AddMonth(t *testing.T) {
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

		{"2020-01-01", "2020-02-01"},
		{"2020-01-31", "2020-03-02"},
		{"2020-02-01", "2020-03-01"},
		{"2020-02-28", "2020-03-28"},
		{"2020-02-29", "2020-03-29"},
	}

	for index, test := range tests {
		c := Parse(test.input).AddMonth()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_AddMonthNoOverflow(t *testing.T) {
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

		{"2020-01-01", "2020-02-01"},
		{"2020-01-31", "2020-02-29"},
		{"2020-02-01", "2020-03-01"},
		{"2020-02-28", "2020-03-28"},
		{"2020-02-29", "2020-03-29"},
	}

	for index, test := range tests {
		c := Parse(test.input).AddMonthNoOverflow()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SubMonth(t *testing.T) {
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

		{"2020-01-01", "2019-12-01"},
		{"2020-01-31", "2019-12-31"},
		{"2020-02-01", "2020-01-01"},
		{"2020-02-28", "2020-01-28"},
		{"2020-02-29", "2020-01-29"},
	}

	for index, test := range tests {
		c := Parse(test.input).SubMonth()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SubMonthNoOverflow(t *testing.T) {
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

		{"2020-01-01", "2019-12-01"},
		{"2020-01-31", "2019-12-31"},
		{"2020-02-01", "2020-01-01"},
		{"2020-02-28", "2020-01-28"},
		{"2020-02-29", "2020-01-29"},
	}

	for index, test := range tests {
		c := Parse(test.input).SubMonthNoOverflow()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_AddWeeks(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		weeks    int
		expected string // 期望值
	}{
		{"", 3, ""},
		{"0", 3, ""},
		{"0000-00-00", 3, ""},
		{"00:00:00", 3, ""},
		{"0000-00-00 00:00:00", 3, ""},

		{"2020-01-01", 3, "2020-01-22"},
		{"2020-01-31", 3, "2020-02-21"},
		{"2020-02-01", 3, "2020-02-22"},
		{"2020-02-28", 3, "2020-03-20"},
		{"2020-02-29", 3, "2020-03-21"},
	}

	for index, test := range tests {
		c := Parse(test.input).AddWeeks(test.weeks)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SubWeeks(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		weeks    int
		expected string // 期望值
	}{
		{"", 3, ""},
		{"0", 3, ""},
		{"0000-00-00", 3, ""},
		{"00:00:00", 3, ""},
		{"0000-00-00 00:00:00", 3, ""},

		{"2020-01-01", 3, "2019-12-11"},
		{"2020-01-31", 3, "2020-01-10"},
		{"2020-02-01", 3, "2020-01-11"},
		{"2020-02-28", 3, "2020-02-07"},
		{"2020-02-29", 3, "2020-02-08"},
	}

	for index, test := range tests {
		c := Parse(test.input).SubWeeks(test.weeks)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_AddWeek(t *testing.T) {
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

		{"2020-01-01", "2020-01-08"},
		{"2020-01-31", "2020-02-07"},
		{"2020-02-01", "2020-02-08"},
		{"2020-02-28", "2020-03-06"},
		{"2020-02-29", "2020-03-07"},
	}

	for index, test := range tests {
		c := Parse(test.input).AddWeek()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SubWeek(t *testing.T) {
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

		{"2020-01-01", "2019-12-25"},
		{"2020-01-31", "2020-01-24"},
		{"2020-02-01", "2020-01-25"},
		{"2020-02-28", "2020-02-21"},
		{"2020-02-29", "2020-02-22"},
	}

	for index, test := range tests {
		c := Parse(test.input).SubWeek()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_AddDays(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		days     int
		expected string // 期望值
	}{
		{"", 3, ""},
		{"0", 3, ""},
		{"0000-00-00", 3, ""},
		{"00:00:00", 3, ""},
		{"0000-00-00 00:00:00", 3, ""},

		{"2020-01-01", 3, "2020-01-04"},
		{"2020-01-31", 3, "2020-02-03"},
		{"2020-02-01", 3, "2020-02-04"},
		{"2020-02-28", 3, "2020-03-02"},
		{"2020-02-29", 3, "2020-03-03"},
	}

	for index, test := range tests {
		c := Parse(test.input).AddDays(test.days)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SubDays(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		days     int
		expected string // 期望值
	}{
		{"", 3, ""},
		{"0", 3, ""},
		{"0000-00-00", 3, ""},
		{"00:00:00", 3, ""},
		{"0000-00-00 00:00:00", 3, ""},

		{"2020-01-01", 3, "2019-12-29"},
		{"2020-01-31", 3, "2020-01-28"},
		{"2020-02-01", 3, "2020-01-29"},
		{"2020-02-28", 3, "2020-02-25"},
		{"2020-02-29", 3, "2020-02-26"},
	}

	for index, test := range tests {
		c := Parse(test.input).SubDays(test.days)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_AddDay(t *testing.T) {
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

		{"2020-01-01", "2020-01-02"},
		{"2020-01-31", "2020-02-01"},
		{"2020-02-01", "2020-02-02"},
		{"2020-02-28", "2020-02-29"},
		{"2020-02-29", "2020-03-01"},
	}

	for index, test := range tests {
		c := Parse(test.input).AddDay()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SubDay(t *testing.T) {
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

		{"2020-01-01", "2019-12-31"},
		{"2020-01-31", "2020-01-30"},
		{"2020-02-01", "2020-01-31"},
		{"2020-02-28", "2020-02-27"},
		{"2020-02-29", "2020-02-28"},
	}

	for index, test := range tests {
		c := Parse(test.input).SubDay()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_AddHours(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		hours    int
		expected string // 期望值
	}{
		{"", 3, ""},
		{"0", 3, ""},
		{"0000-00-00", 3, ""},
		{"00:00:00", 3, ""},
		{"0000-00-00 00:00:00", 3, ""},

		{"2020-08-05 13:14:15", 3, "2020-08-05 16:14:15"},
	}

	for index, test := range tests {
		c := Parse(test.input).AddHours(test.hours)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SubHours(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		hours    int
		expected string // 期望值
	}{
		{"", 3, ""},
		{"0", 3, ""},
		{"0000-00-00", 3, ""},
		{"00:00:00", 3, ""},
		{"0000-00-00 00:00:00", 3, ""},

		{"2020-08-05 13:14:15", 3, "2020-08-05 10:14:15"},
	}

	for index, test := range tests {
		c := Parse(test.input).SubHours(test.hours)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_AddHour(t *testing.T) {
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

		{"2020-08-05 13:14:15", "2020-08-05 14:14:15"},
	}

	for index, test := range tests {
		c := Parse(test.input).AddHour()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SubHour(t *testing.T) {
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

		{"2020-08-05 13:14:15", "2020-08-05 12:14:15"},
	}

	for index, test := range tests {
		c := Parse(test.input).SubHour()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_AddMinutes(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		minutes  int
		expected string // 期望值
	}{
		{"", 3, ""},
		{"0", 3, ""},
		{"0000-00-00", 3, ""},
		{"00:00:00", 3, ""},
		{"0000-00-00 00:00:00", 3, ""},

		{"2020-08-05 13:14:15", 3, "2020-08-05 13:17:15"},
	}

	for index, test := range tests {
		c := Parse(test.input).AddMinutes(test.minutes)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SubMinutes(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		minutes  int
		expected string // 期望值
	}{
		{"", 3, ""},
		{"0", 3, ""},
		{"0000-00-00", 3, ""},
		{"00:00:00", 3, ""},
		{"0000-00-00 00:00:00", 3, ""},

		{"2020-08-05 13:14:15", 3, "2020-08-05 13:11:15"},
	}

	for index, test := range tests {
		c := Parse(test.input).SubMinutes(test.minutes)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_AddMinute(t *testing.T) {
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

		{"2020-08-05 13:14:15", "2020-08-05 13:15:15"},
	}

	for index, test := range tests {
		c := Parse(test.input).AddMinute()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SubMinute(t *testing.T) {
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

		{"2020-08-05 13:14:15", "2020-08-05 13:13:15"},
	}

	for index, test := range tests {
		c := Parse(test.input).SubMinute()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_AddSeconds(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		seconds  int
		expected string // 期望值
	}{
		{"", 3, ""},
		{"0", 3, ""},
		{"0000-00-00", 3, ""},
		{"00:00:00", 3, ""},
		{"0000-00-00 00:00:00", 3, ""},

		{"2020-08-05 13:14:15", 3, "2020-08-05 13:14:18"},
	}

	for index, test := range tests {
		c := Parse(test.input).AddSeconds(test.seconds)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SubSeconds(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		seconds  int
		expected string // 期望值
	}{
		{"", 3, ""},
		{"0", 3, ""},
		{"0000-00-00", 3, ""},
		{"00:00:00", 3, ""},
		{"0000-00-00 00:00:00", 3, ""},

		{"2020-08-05 13:14:15", 3, "2020-08-05 13:14:12"},
	}

	for index, test := range tests {
		c := Parse(test.input).SubSeconds(test.seconds)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_AddSecond(t *testing.T) {
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

		{"2020-08-05 13:14:15", "2020-08-05 13:14:16"},
	}

	for index, test := range tests {
		c := Parse(test.input).AddSecond()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SubSecond(t *testing.T) {
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

		{"2020-08-05 13:14:15", "2020-08-05 13:14:14"},
	}

	for index, test := range tests {
		c := Parse(test.input).SubSecond()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_AddMilliseconds(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input        string // 输入值
		milliseconds int
		expected     string // 期望值
	}{
		{"", 3, ""},
		{"0", 3, ""},
		{"0000-00-00", 3, ""},
		{"00:00:00", 3, ""},
		{"0000-00-00 00:00:00", 3, ""},

		{"2020-08-05 13:14:15.222222222", 3, "2020-08-05 13:14:15.225222222 +0800 CST"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC).AddMilliseconds(test.milliseconds)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SubMilliseconds(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input        string // 输入值
		milliseconds int
		expected     string // 期望值
	}{
		{"", 3, ""},
		{"0", 3, ""},
		{"0000-00-00", 3, ""},
		{"00:00:00", 3, ""},
		{"0000-00-00 00:00:00", 3, ""},

		{"2020-08-05 13:14:15.222222222", 3, "2020-08-05 13:14:15.219222222 +0800 CST"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC).SubMilliseconds(test.milliseconds)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_AddMillisecond(t *testing.T) {
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

		{"2020-08-05 13:14:15.222222222", "2020-08-05 13:14:15.223222222 +0800 CST"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC).AddMillisecond()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SubMillisecond(t *testing.T) {
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

		{"2020-08-05 13:14:15.222222222", "2020-08-05 13:14:15.221222222 +0800 CST"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC).SubMillisecond()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_AddMicroseconds(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input        string // 输入值
		microseconds int
		expected     string // 期望值
	}{
		{"", 3, ""},
		{"0", 3, ""},
		{"0000-00-00", 3, ""},
		{"00:00:00", 3, ""},
		{"0000-00-00 00:00:00", 3, ""},

		{"2020-08-05 13:14:15.222222222", 3, "2020-08-05 13:14:15.222225222 +0800 CST"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC).AddMicroseconds(test.microseconds)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SubMicroseconds(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input        string // 输入值
		microseconds int
		expected     string // 期望值
	}{
		{"", 3, ""},
		{"0", 3, ""},
		{"0000-00-00", 3, ""},
		{"00:00:00", 3, ""},
		{"0000-00-00 00:00:00", 3, ""},

		{"2020-08-05 13:14:15.222222222", 3, "2020-08-05 13:14:15.222219222 +0800 CST"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC).SubMicroseconds(test.microseconds)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_AddMicrosecond(t *testing.T) {
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

		{"2020-08-05 13:14:15.222222222", "2020-08-05 13:14:15.222223222 +0800 CST"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC).AddMicrosecond()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SubMicrosecond(t *testing.T) {
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

		{"2020-08-05 13:14:15.222222222", "2020-08-05 13:14:15.222221222 +0800 CST"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC).SubMicrosecond()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_AddNanoseconds(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input       string // 输入值
		nanoseconds int
		expected    string // 期望值
	}{
		{"", 3, ""},
		{"0", 3, ""},
		{"0000-00-00", 3, ""},
		{"00:00:00", 3, ""},
		{"0000-00-00 00:00:00", 3, ""},

		{"2020-08-05 13:14:15.222222222", 3, "2020-08-05 13:14:15.222222225 +0800 CST"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC).AddNanoseconds(test.nanoseconds)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SubNanoseconds(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input       string // 输入值
		nanoseconds int
		expected    string // 期望值
	}{
		{"", 3, ""},
		{"0", 3, ""},
		{"0000-00-00", 3, ""},
		{"00:00:00", 3, ""},
		{"0000-00-00 00:00:00", 3, ""},

		{"2020-08-05 13:14:15.222222222", 3, "2020-08-05 13:14:15.222222219 +0800 CST"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC).SubNanoseconds(test.nanoseconds)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_AddNanosecond(t *testing.T) {
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

		{"2020-08-05 13:14:15.222222222", "2020-08-05 13:14:15.222222223 +0800 CST"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC).AddNanosecond()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_SubNanosecond(t *testing.T) {
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

		{"2020-08-05 13:14:15.222222222", "2020-08-05 13:14:15.222222221 +0800 CST"},
	}

	for index, test := range tests {
		c := Parse(test.input, PRC).SubNanosecond()
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestError_Traveler(t *testing.T) {
	duration := "10x"
	c := Parse("2020-08-05").AddDuration(duration)
	assert.NotNil(t, c.Error, "It should catch an exception in AddDuration()")
}
