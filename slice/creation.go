package main

import "fmt"

func main() {
	slice := []string{"one", "two", "three"}
	for i, w := range slice {
		fmt.Printf("%d -> %s\n", i+1, w)
	}
}
