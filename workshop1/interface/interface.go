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

type pui int

func (p pui) sleep() int {
	return int(p)
}

func call(u User) {
	fmt.Println(u.sleep())
}

func main() {
	u := Member{}
	call(u)

	p := pui(70)
	call(p)
}
