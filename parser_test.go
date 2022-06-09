package carbon

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_Parse(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		output string // 期望值
	}{
		{"", ""},
		{"0", ""},
		{"0000-00-00", ""},
		{"00:00:00", ""},
		{"0000-00-00 00:00:00", ""},

		{"2020-08-05 13:14:15", "2020-08-05 13:14:15 +0800 CST"},
		{"2020-08-05", "2020-08-05 00:00:00 +0800 CST"},
		{"2020-08-05T13:14:15+08:00", "2020-08-05 13:14:15 +0800 CST"},
		{"2020-08-05T13:14:15.999+08:00", "2020-08-05 13:14:15.999 +0800 CST"},
		{"2020-08-05T13:14:15.999999+08:00", "2020-08-05 13:14:15.999999 +0800 CST"},
		{"2020-08-05T13:14:15.999999999+08:00", "2020-08-05 13:14:15.999999999 +0800 CST"},
		{"20200805131415", "2020-08-05 13:14:15 +0800 CST"},
		{"20200805131415.999", "2020-08-05 13:14:15.999 +0800 CST"},
		{"20200805131415.999999", "2020-08-05 13:14:15.999999 +0800 CST"},
		{"20200805131415.999999999", "2020-08-05 13:14:15.999999999 +0800 CST"},
		{"20200805", "2020-08-05 00:00:00 +0800 CST"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).Parse(test.input)
		assert.Nil(c.Error)
		assert.Equal(test.output, c.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := Parse(test.input, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.output, c.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ParseByFormat(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		format string // 输入参数
		output string // 期望值
	}{
		{"", "Y|m|d H:i:s", ""},
		{"0", "Y|m|d H:i:s", ""},
		{"0000-00-00", "Y|m|d H:i:s", ""},
		{"00:00:00", "Y|m|d H:i:s", ""},
		{"0000-00-00 00:00:00", "Y|m|d H:i:s", ""},

		{"2020|08|05 13:14:15", "Y|m|d H:i:s", "2020-08-05 13:14:15"},
		{"It is 2020-08-05 13:14:15", "\\I\\t \\i\\s Y-m-d H:i:s", "2020-08-05 13:14:15"},
		{"今天是 2020年08月05日13时14分15秒", "今天是 Y年m月d日H时i分s秒", "2020-08-05 13:14:15"},
		{"上次上报时间:2020-08-05 13:14:15，请每日按时打卡", "上次上报时间:Y-m-d H:i:s，请每日按时打卡", "2020-08-05 13:14:15"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).ParseByFormat(test.input, test.format)
		assert.Nil(c.Error)
		assert.Equal(test.output, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := ParseByFormat(test.input, test.format, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.output, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_ParseByLayout(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input  string // 输入值
		layout string // 输入参数
		output string // 期望值
	}{
		{"", "2006-01-02", ""},
		{"0", "2006-01-02", ""},
		{"0000-00-00", "2006-01-02", ""},
		{"00:00:00", "15:04:05", ""},
		{"0000-00-00 00:00:00", "2006-01-02 15:04:05", ""},

		{"2020|08|05 13:14:15", "2006|01|02 15:04:05", "2020-08-05 13:14:15"},
		{"It is 2020|08|05 13:14:15", "It is 2006|01|02 15:04:05", "2020-08-05 13:14:15"},
		{"今天是 2020年08月05日13时14分15秒", "今天是 2006年01月02日15时04分05秒", "2020-08-05 13:14:15"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).ParseByLayout(test.input, test.layout)
		assert.Nil(c.Error)
		assert.Equal(test.output, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := ParseByLayout(test.input, test.layout, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.output, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestError_Parse(t *testing.T) {
	date, layout, format, timezone := "2020-08-50", "2006-01-02", DateLayout, "xxx"
	assert.NotNil(t, Parse(date).Error, "It should catch an exception in Parse()")
	assert.NotNil(t, ParseByLayout(date, layout, timezone).Error, "It should catch an exception in ParseByLayout()")
	assert.NotNil(t, ParseByFormat(date, format, timezone).Error, "It should catch an exception in ParseByFormat()")
}
