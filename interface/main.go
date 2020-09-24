package main

import "fmt"

type UserInterface interface {
	Surname() string
	Age() int
}

type User struct {
	name string
	age  int
}

func (user *User) Surname() string {
	return user.name
}

func (user *User) Age() int {
	return user.age
}

type FakeUser struct{}

func (user FakeUser) Surname() string {
	return "Laurent"
}

// As expected, commenting this method will lead to a compilation error
func (user FakeUser) Age() int {
	return 38
}

func Display(user UserInterface) {
	fmt.Printf("%s (%d)\n", user.Surname(), user.Age())
}

type Group struct {
	User1 User
	User2 User
}

func main() {
	user1 := User{"Balibalo", 7}
	user2 := FakeUser{}
	Display(&user1)
	Display(user2)

	g := Group{user1, &user2}
	Display(g.User1)
}
