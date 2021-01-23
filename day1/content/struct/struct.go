package main

import (
	"fmt"
)

type User struct {
	Id int
	UserForNew
}

type User2 struct {
	Id         int
	UserForNew UserForNew
}
type UserForNew struct {
	Name string
}

func NewUser(name string) User {
	return User{1, UserForNew{"bomb0069"}}
}

func main() {
	u1 := User{1, UserForNew{"bomb0069"}}
	u2 := User2{2, UserForNew{"karan"}}
	u3 := NewUser("bomb0069")
	u1.Name = "XXX" //
	fmt.Printf("%+v\n", u1)
	fmt.Printf("%+v\n", u1.Name)
	fmt.Printf("%+v\n", u1.UserForNew.Name)
	fmt.Printf("%+v\n", u2.UserForNew.Name)
	fmt.Printf("%+v\n", u3)

	fmt.Printf("%+v\n", u3.Name)
}
