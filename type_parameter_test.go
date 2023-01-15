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
	// Printing the value of the parameter and returning the value of the parameter.
	fmt.Println(param)
	return param
}

// The function takes two parameters, one of type `k` and one of type `w`
func MultipleParam[k any, w any](param1 k, param2 w) {
	// Printing the value of the two parameters.
	fmt.Println(param1)
	fmt.Println(param2)
}

// > This function takes two values of the same type and returns true if they are equal
func IsEqual[T comparable](val1, val2 T) bool {
	// Checking if the two values are equal.
	if val1 == val2 {
		return true
	} else {
		return false
	}
}

// > The function `SingleParam` takes a single parameter of any type and returns a value of the same
// type
func TestSampleLength(t *testing.T) {
	// Calling the function `SingleParam` with a parameter of type `string` and assigning the return value
	// to a variable of type `string`.
	var res string = SingleParam("OK")
	assert.Equal(t, "OK", res)

	// Calling the function `SingleParam` with a parameter of type `int` and assigning the return value to
	// a variable of type `int`.
	var rate = SingleParam(100)
	assert.Equal(t, 100, rate)
}

// The function takes two parameters of any type and returns nothing
func TestSampleMultipleParam(t *testing.T) {
	// Calling the function `MultipleParam` with two parameters of type `string` and `int` and
	// 		then calling the function `MultipleParam` with two parameters of type `int` and `string`.
	MultipleParam[string, int]("Arif", 100)
	MultipleParam[int, string](100, "Arif")
}

// `IsEqual` takes two parameter and returns a boolean indicating whether or not they are equal
func TestIsEqual(t *testing.T) {
	// Calling the function `IsEqual` with the type `string`, `int`, and `bool` respectively.
	var s = IsEqual[string]("Arif", "Arif")
	var integ = IsEqual[int](110, 110)
	var bools = IsEqual[bool](true, true)

	// Checking if the values are true.
	assert.True(t, s)
	assert.True(t, integ)
	assert.True(t, bools)
}
