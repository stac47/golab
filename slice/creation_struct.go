package main

import (
	"fmt"
)

type A struct {
	n int
	m int
}

func main() {
	slice := []A{
		{1, 2},
		{2, 3},
	}

	for _, a := range slice {
		fmt.Printf("%d\n", a.n)
	}
}
