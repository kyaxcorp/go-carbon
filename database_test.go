package carbon

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCarbon_Scan(t *testing.T) {
	c := NewCarbon()
	err := c.Scan(time.Now())
	assert.Nil(t, err)
	assert.Equal(t, c.ToDateTimeString(), Now().ToDateTimeString())
}

func TestCarbon_Value(t *testing.T) {
	c := Now()
	v, err := c.Value()
	assert.Nil(t, err)
	assert.Equal(t, v, c.time)
}

func TestError_Scan(t *testing.T) {
	c, v := NewCarbon(), "xxx"
	err := c.Scan(v)
	assert.Equal(t, err, fmt.Errorf("can not convert %v to carbon", v))
}

func TestError_Value(t *testing.T) {
	c := Parse("")
	v, err := c.Value()
	assert.Nil(t, err)
	assert.Equal(t, v, nil)
}
