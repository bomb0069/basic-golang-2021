package main

import (
	"fmt"
)

type User struct {
	Id   int
	Name string
}

func (user *User) sleep() {
	user.Name = "XXXXX"
}

func main() {
	u := User{2, "karan"}
	u.sleep()
	fmt.Println(u.Name)

	u2 := &User{2, "karan"}
	u2.sleep()
	fmt.Println(u2.Name)
}
