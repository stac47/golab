package main

import (
	"fmt"
)

type MyEnum int

const (
	A MyEnum = iota
	B
	C
)

func main() {
	var a MyEnum = C
	fmt.Printf("Result = %d\n", a)
}
