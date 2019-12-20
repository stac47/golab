package main

import (
	"fmt"
)

func main() {
	s1 := []string{"one", "two", "three"}
	s2 := []string{"Hello", "World"}
	s1 = s2
	for _, w := range s1 {
		fmt.Println(w)
	}
}
