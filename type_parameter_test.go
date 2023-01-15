package gogenerics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestSampleLength(t *testing.T) {
	var res string = Length("OK")
	assert.Equal(t, "OK", res)

	var rate = Length(100)
	assert.Equal(t, 100, rate)

}
