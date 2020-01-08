package main

import "fmt"

type User interface {
	GetName() string
}

type UserImpl struct {
	name string
}

func (u UserImpl) GetName() string {
	return u.name
}

func NewUser(t int) *User {
	var user User = nil
	if t == 0 {
		user = UserImpl{"lolo"}
	}
	return &user
}

func main() {
	fmt.Printf("%v\n", *NewUser(1))
	fmt.Printf("%v\n", *NewUser(1))
}
