package main

import (
	"fmt"
)

type User struct {
	Name string
}

type User2 struct {
	Id int
	User
}

func (user User) printName() {
	fmt.Println("User := ", user.Name)
}

func (user User2) printName() {
	fmt.Println("User 2 := ", user.Name)
}

func main() {
	u2 := User2{2, User{"karan"}}
	u2.printName()

}
