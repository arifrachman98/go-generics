package gogenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Range int

type Number interface {
	~int | int16 | int32 | int64 |
		float32 | float64
}

func Min[T Number](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestMin(t *testing.T) {
	assert.Equal(t, 100, Min[int](100, 200))
	assert.Equal(t, int64(100), Min[int64](int64(100), int64(200)))
	assert.Equal(t, float64(0.33), Min[float64](0.33, 0.78500))
	assert.Equal(t, Range(1000), Min[Range](Range(1000), Range(50000)))
}
