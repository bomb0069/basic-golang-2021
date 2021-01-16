package main

import (
	"fmt"
)

type User2 struct {
	Id   int
	Name string
}

func (user User2) printName() {
	fmt.Println(user.Name)
}

func main() {
	u2 := User2{2, "karan"}
	u2.printName()

}
