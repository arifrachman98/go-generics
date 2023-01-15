// The package name.
package gogenerics

// Importing the packages fmt, testing, and github.com/stretchr/testify/assert.
import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// `Length` is a function that takes a parameter of type `T` and returns a value of type `T`
func SingleParam[T any](param T) T {
	fmt.Println(param)
	return param
}

// The function takes two parameters, one of type `k` and one of type `w`
func MultipleParam[k any, w any](param1 k, param2 w) {
	fmt.Println(param1)
	fmt.Println(param2)
}

func TestSampleLength(t *testing.T) {
	var res string = SingleParam("OK")
	assert.Equal(t, "OK", res)

	var rate = SingleParam(100)
	assert.Equal(t, 100, rate)

}

func TestSampleMultipleParam(t *testing.T) {
	MultipleParam[string, int]("Arif", 100)
	MultipleParam[int, string](100, "Arif")
}
