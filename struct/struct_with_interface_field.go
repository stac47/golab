package main

import (
	"fmt"
)

type T struct {
	A interface{}
}

func main() {
	s := []T{
		{1},
		{"Hello"},
	}
	for _, t := range s {
		fmt.Printf("%T\n", t.A)
	}
}
