package gogenerics

import (
	"fmt"
	"testing"
)

type Bag[T any] []T

func PrintBag[T any](bag Bag[T]) {
	for _, v := range bag {
		fmt.Println(v)
	}
}

func TestBag(t *testing.T) {
	num := Bag[int]{1, 2, 3, 4, 5}
	PrintBag(num)

	tag := Bag[string]{"Sell", "Buy", "Trade"}
	fmt.Println(tag)
	PrintBag(tag)
}
