// The package name.
package gogenerics

// Importing the packages fmt, testing, and github.com/stretchr/testify/assert.
import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Employee is an interface that has a method called GetName that returns a string.
// @property {string} GetName - This is a method that returns the name of the employee.
type Employee interface {
	GetName() string
}

// The function GetName takes a parameter of type Employee and returns a string
func GetName[T Employee](param T) string {
	return param.GetName()
}

// `Manager` is an interface that has two methods: `GetName()` and `GetManagerName()`.
// @property {string} GetName - The name of the employee
// @property {string} GetManagerName - The name of the manager.
type Manager interface {
	GetName() string
	GetManagerName() string
}

// @property {string} Name - The name of the manager
type MyManager struct {
	Name string
}

// Defining a method on the type `MyManager`.
func (m *MyManager) GetName() string {
	return m.Name
}

// Defining a method on the type `MyManager`.
func (m *MyManager) GetManagerName() string {
	return m.Name
}

// The VicePresident interface is a type that has two methods: GetName() and GetVicePresident().
// @property {string} GetName - Returns the name of the president
// @property {string} GetVicePresident - Returns the name of the Vice President
type VicePresident interface {
	GetName() string
	GetVicePresidentName() string
}

// @property {string} Name - The name of the Vice President
type MyVicePresident struct {
	Name string
}

// Defining a method on the type `MyVicePresident`.
func (m *MyVicePresident) GetName() string {
	return m.Name
}

// Defining a method on the type 'MyVicePresident'
func (m *MyVicePresident) GetVicePresidentName() string {
	return m.Name
}

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

func TestGetName(t *testing.T) {
	assert.Equal(t, "Kinda", GetName[VicePresident](&MyVicePresident{Name: "Kinda"}))
	assert.Equal(t, "Arif", GetName[Manager](&MyManager{Name: "Arif"}))
}
