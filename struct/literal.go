package main

import "fmt"

type User struct {
	Name string
	Age  int
}

type Superman struct {
	User
	Heroe bool
}

func main() {
	me := Superman{
		User{"stac", 38},
		true,
	}
	stef := Superman{
		Name:  "stef",
		Age:   39,
		Heroe: true,
	}
	fmt.Printf("%v\n", me)
	fmt.Printf("%v\n", stef)
}
