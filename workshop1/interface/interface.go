package main

import (
	"fmt"
)

type User interface {
	sleep() int
}

type Member struct {
}

func (m Member) sleep() int {
	return 50
}

func call(u User) {
	fmt.Println(u.sleep())
}

func main() {
	u := Member{}
	call(u)
}
