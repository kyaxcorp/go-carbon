package carbon

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_CreateFromTimestamp(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		timestamp int64  // 输入参数
		expected  string // 期望值
	}{
		{-1, "1970-01-01 07:59:59 +0800 CST"},
		{0, "1970-01-01 08:00:00 +0800 CST"},
		{1, "1970-01-01 08:00:01 +0800 CST"},
		{1649735755, "2022-04-12 11:55:55 +0800 CST"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).CreateFromTimestamp(test.timestamp)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := CreateFromTimestamp(test.timestamp, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_CreateFromTimestampMilli(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		timestamp int64  // 输入参数
		expected  string // 期望值
	}{
		{-1, "1970-01-01 07:59:59.999 +0800 CST"},
		{0, "1970-01-01 08:00:00 +0800 CST"},
		{1, "1970-01-01 08:00:00.001 +0800 CST"},
		{1649735755981, "2022-04-12 11:55:55.981 +0800 CST"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).CreateFromTimestampMilli(test.timestamp)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := CreateFromTimestampMilli(test.timestamp, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_CreateFromTimestampMicro(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		timestamp int64  // 输入参数
		expected  string // 期望值
	}{
		{-1, "1970-01-01 07:59:59.999999 +0800 CST"},
		{0, "1970-01-01 08:00:00 +0800 CST"},
		{1, "1970-01-01 08:00:00.000001 +0800 CST"},
		{1649735755981566, "2022-04-12 11:55:55.981566 +0800 CST"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).CreateFromTimestampMicro(test.timestamp)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := CreateFromTimestampMicro(test.timestamp, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_CreateFromTimestampNano(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		timestamp int64  // 输入参数
		expected  string // 期望值
	}{
		{-1, "1970-01-01 07:59:59.999999999 +0800 CST"},
		{0, "1970-01-01 08:00:00 +0800 CST"},
		{1, "1970-01-01 08:00:00.000000001 +0800 CST"},
		{1649735755981566000, "2022-04-12 11:55:55.981566 +0800 CST"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).CreateFromTimestampNano(test.timestamp)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := CreateFromTimestampNano(test.timestamp, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_CreateFromDateTime(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		year, month, day, hour, minute, second int    // 输入参数
		expected                               string // 期望值
	}{
		{2020, 1, 1, 13, 14, 15, "2020-01-01 13:14:15"},
		{2020, 1, 31, 13, 14, 15, "2020-01-31 13:14:15"},
		{2020, 2, 1, 13, 14, 15, "2020-02-01 13:14:15"},
		{2020, 2, 28, 13, 14, 15, "2020-02-28 13:14:15"},
		{2020, 2, 29, 13, 14, 15, "2020-02-29 13:14:15"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).CreateFromDateTime(test.year, test.month, test.day, test.hour, test.minute, test.second)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := CreateFromDateTime(test.year, test.month, test.day, test.hour, test.minute, test.second, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_CreateFromDateTimeMilli(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		year, month, day, hour, minute, second, millisecond int    // 输入参数
		expected                                            string // 期望值
	}{
		{2020, 1, 1, 13, 14, 15, 999, "2020-01-01 13:14:15.999 +0800 CST"},
		{2020, 1, 31, 13, 14, 15, 999, "2020-01-31 13:14:15.999 +0800 CST"},
		{2020, 2, 1, 13, 14, 15, 999, "2020-02-01 13:14:15.999 +0800 CST"},
		{2020, 2, 28, 13, 14, 15, 999, "2020-02-28 13:14:15.999 +0800 CST"},
		{2020, 2, 29, 13, 14, 15, 999, "2020-02-29 13:14:15.999 +0800 CST"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).CreateFromDateTimeMilli(test.year, test.month, test.day, test.hour, test.minute, test.second, test.millisecond)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := CreateFromDateTimeMilli(test.year, test.month, test.day, test.hour, test.minute, test.second, test.millisecond, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_CreateFromDateTimeMicro(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		year, month, day, hour, minute, second, microsecond int    // 输入参数
		expected                                            string // 期望值
	}{
		{2020, 1, 1, 13, 14, 15, 999999, "2020-01-01 13:14:15.999999 +0800 CST"},
		{2020, 1, 31, 13, 14, 15, 999999, "2020-01-31 13:14:15.999999 +0800 CST"},
		{2020, 2, 1, 13, 14, 15, 999999, "2020-02-01 13:14:15.999999 +0800 CST"},
		{2020, 2, 28, 13, 14, 15, 999999, "2020-02-28 13:14:15.999999 +0800 CST"},
		{2020, 2, 29, 13, 14, 15, 999999, "2020-02-29 13:14:15.999999 +0800 CST"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).CreateFromDateTimeMicro(test.year, test.month, test.day, test.hour, test.minute, test.second, test.microsecond)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := CreateFromDateTimeMicro(test.year, test.month, test.day, test.hour, test.minute, test.second, test.microsecond, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_CreateFromDateTimeNano(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		year, month, day, hour, minute, second, nanosecond int    // 输入参数
		expected                                           string // 期望值
	}{
		{2020, 1, 1, 13, 14, 15, 999999999, "2020-01-01 13:14:15.999999999 +0800 CST"},
		{2020, 1, 31, 13, 14, 15, 999999999, "2020-01-31 13:14:15.999999999 +0800 CST"},
		{2020, 2, 1, 13, 14, 15, 999999999, "2020-02-01 13:14:15.999999999 +0800 CST"},
		{2020, 2, 28, 13, 14, 15, 999999999, "2020-02-28 13:14:15.999999999 +0800 CST"},
		{2020, 2, 29, 13, 14, 15, 999999999, "2020-02-29 13:14:15.999999999 +0800 CST"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).CreateFromDateTimeNano(test.year, test.month, test.day, test.hour, test.minute, test.second, test.nanosecond)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := CreateFromDateTimeNano(test.year, test.month, test.day, test.hour, test.minute, test.second, test.nanosecond, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_CreateFromDate(t *testing.T) {
	assert := assert.New(t)

	clock := Now(PRC).ToTimeString()
	tests := []struct {
		year, month, day int    // 输入参数
		expected         string // 期望值
	}{
		{2020, 1, 1, "2020-01-01 " + clock},
		{2020, 1, 31, "2020-01-31 " + clock},
		{2020, 2, 1, "2020-02-01 " + clock},
		{2020, 2, 28, "2020-02-28 " + clock},
		{2020, 2, 29, "2020-02-29 " + clock},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).CreateFromDate(test.year, test.month, test.day)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := CreateFromDate(test.year, test.month, test.day, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_CreateFromDateMilli(t *testing.T) {
	assert := assert.New(t)

	clock := Now(PRC).ToTimeString()
	tests := []struct {
		year, month, day, millisecond int    // 输入参数
		expected                      string // 期望值
	}{
		{2020, 1, 1, 999, "2020-01-01 " + clock + ".999"},
		{2020, 1, 31, 999, "2020-01-31 " + clock + ".999"},
		{2020, 2, 1, 999, "2020-02-01 " + clock + ".999"},
		{2020, 2, 28, 999, "2020-02-28 " + clock + ".999"},
		{2020, 2, 29, 999, "2020-02-29 " + clock + ".999"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).CreateFromDateMilli(test.year, test.month, test.day, test.millisecond)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeMilliString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := CreateFromDateMilli(test.year, test.month, test.day, test.millisecond, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeMilliString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_CreateFromDateMicro(t *testing.T) {
	assert := assert.New(t)

	clock := Now(PRC).ToTimeString()
	tests := []struct {
		year, month, day, microsecond int    // 输入参数
		expected                      string // 期望值
	}{
		{2020, 1, 1, 999999, "2020-01-01 " + clock + ".999999"},
		{2020, 1, 31, 999999, "2020-01-31 " + clock + ".999999"},
		{2020, 2, 1, 999999, "2020-02-01 " + clock + ".999999"},
		{2020, 2, 28, 999999, "2020-02-28 " + clock + ".999999"},
		{2020, 2, 29, 999999, "2020-02-29 " + clock + ".999999"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).CreateFromDateMicro(test.year, test.month, test.day, test.microsecond)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeMicroString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := CreateFromDateMicro(test.year, test.month, test.day, test.microsecond, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeMicroString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_CreateFromDateNano(t *testing.T) {
	assert := assert.New(t)

	clock := Now(PRC).ToTimeString()
	tests := []struct {
		year, month, day, nanosecond int    // 输入参数
		expected                     string // 期望值
	}{
		{2020, 1, 1, 999999999, "2020-01-01 " + clock + ".999999999"},
		{2020, 1, 31, 999999999, "2020-01-31 " + clock + ".999999999"},
		{2020, 2, 1, 999999999, "2020-02-01 " + clock + ".999999999"},
		{2020, 2, 28, 999999999, "2020-02-28 " + clock + ".999999999"},
		{2020, 2, 29, 999999999, "2020-02-29 " + clock + ".999999999"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).CreateFromDateNano(test.year, test.month, test.day, test.nanosecond)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeNanoString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := CreateFromDateNano(test.year, test.month, test.day, test.nanosecond, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeNanoString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_CreateFromTime(t *testing.T) {
	assert := assert.New(t)

	date := Now(PRC).ToDateString()
	tests := []struct {
		hour, minute, second int    // 输入参数
		expected             string // 期望值
	}{
		{0, 0, 0, date + " 00:00:00"},
		{00, 00, 15, date + " 00:00:15"},
		{00, 14, 15, date + " 00:14:15"},
		{13, 14, 15, date + " 13:14:15"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).CreateFromTime(test.hour, test.minute, test.second)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := CreateFromTime(test.hour, test.minute, test.second, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_CreateFromTimeMilli(t *testing.T) {
	assert := assert.New(t)

	date := Now(PRC).ToDateString()
	tests := []struct {
		hour, minute, second, millisecond int    // 输入参数
		expected                          string // 期望值
	}{
		{0, 0, 0, 999, date + " 00:00:00.999"},
		{00, 00, 15, 999, date + " 00:00:15.999"},
		{00, 14, 15, 999, date + " 00:14:15.999"},
		{13, 14, 15, 999, date + " 13:14:15.999"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).CreateFromTimeMilli(test.hour, test.minute, test.second, test.millisecond)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeMilliString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := CreateFromTimeMilli(test.hour, test.minute, test.second, test.millisecond, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeMilliString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_CreateFromTimeMicro(t *testing.T) {
	assert := assert.New(t)

	date := Now(PRC).ToDateString()
	tests := []struct {
		hour, minute, second, microsecond int    // 输入参数
		expected                          string // 期望值
	}{
		{0, 0, 0, 999999, date + " 00:00:00.999999"},
		{00, 00, 15, 999999, date + " 00:00:15.999999"},
		{00, 14, 15, 999999, date + " 00:14:15.999999"},
		{13, 14, 15, 999999, date + " 13:14:15.999999"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).CreateFromTimeMicro(test.hour, test.minute, test.second, test.microsecond)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeMicroString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := CreateFromTimeMicro(test.hour, test.minute, test.second, test.microsecond, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeMicroString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestCarbon_CreateFromTimeNano(t *testing.T) {
	assert := assert.New(t)

	date := Now(PRC).ToDateString()
	tests := []struct {
		hour, minute, second, nanosecond int    // 输入参数
		expected                         string // 期望值
	}{
		{0, 0, 0, 999999999, date + " 00:00:00.999999999"},
		{00, 00, 15, 999999999, date + " 00:00:15.999999999"},
		{00, 14, 15, 999999999, date + " 00:14:15.999999999"},
		{13, 14, 15, 999999999, date + " 13:14:15.999999999"},
	}

	for index, test := range tests {
		c := SetTimezone(PRC).CreateFromTimeNano(test.hour, test.minute, test.second, test.nanosecond)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeNanoString(), "Current test index is "+strconv.Itoa(index))
	}

	for index, test := range tests {
		c := CreateFromTimeNano(test.hour, test.minute, test.second, test.nanosecond, PRC)
		assert.Nil(c.Error)
		assert.Equal(test.expected, c.ToDateTimeNanoString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestError_Creator(t *testing.T) {
	year, month, day, hour, minute, second, timestamp, timezone := 2020, 8, 5, 13, 14, 15, int64(1577855655), "xxx"

	c := SetTimezone(timezone)
	assert.NotNil(t, c.CreateFromDateTime(year, month, day, hour, minute, second).Error, "It should catch an exception in CreateFromDateTime()")
	assert.NotNil(t, c.CreateFromDate(year, month, day).Error, "It should catch an exception in CreateFromDate()")
	assert.NotNil(t, c.CreateFromTime(hour, minute, second).Error, "It should catch an exception in CreateFromTime()")
	assert.NotNil(t, c.CreateFromTimestamp(timestamp).Error, "It should catch an exception in CreateFromTime()")
	assert.NotNil(t, c.CreateFromTimestampMilli(timestamp).Error, "It should catch an exception in CreateFromTimestampMilli()")
	assert.NotNil(t, c.CreateFromTimestampMicro(timestamp).Error, "It should catch an exception in CreateFromTimestampMicro()")
	assert.NotNil(t, c.CreateFromTimestampNano(timestamp).Error, "It should catch an exception in CreateFromTimestampNano()")
}
