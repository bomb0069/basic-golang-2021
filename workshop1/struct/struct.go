package main

import (
	"fmt"
)

type User struct {
	Id   int
	Name string
}

func New(id int, name string) User {
	return User{1, "bomb0069"}
}

func main() {
	u1 := User{1, "bomb0069"}
	u2 := User{Id: 2, Name: "karan"}
	u3 := New(1, "bomb0069")
	u1.Name = "XXX"
	fmt.Printf("%v", u1)
	fmt.Printf("%+v", u2)
	fmt.Printf("%+v", u3)
}
