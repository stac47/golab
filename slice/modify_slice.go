package main

import (
	"fmt"
)

func f(s *[]string) {
	*s = append(*s, "f")
}

func main() {
	ss := make([]string, 0)
	fmt.Printf("%v", ss)
	f(&ss)
	fmt.Printf("%v", ss)

}
