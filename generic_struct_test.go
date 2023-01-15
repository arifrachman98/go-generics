package gogenerics

import (
	"fmt"
	"testing"
)

type Data[T any] struct {
	First T
	Last  T
}

func TestData(t *testing.T) {
	data := Data[string]{
		First: "Jake",
		Last:  "Undefeated",
	}

	fmt.Println(data)
}
