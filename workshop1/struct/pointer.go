package main

import (
	"fmt"
)

type UserX struct {
	Id   int
	Name string
}

func (user *UserX) sleep() {
	user.Name = "XXXXX"
}

func main() {
	u2 := UserX{2, "karan"}
	u2.sleep()
	fmt.Println(u2.Name)
}
