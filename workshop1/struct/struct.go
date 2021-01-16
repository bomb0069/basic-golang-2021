package main

import (
	"fmt"
)

type user struct {
	id   int
	name string
}

func main() {
	u1 := user{1, "bomb0069"}
	u2 := user{id: 2, name: "karan"}
	u1.name = "XXX"
	fmt.Printf("%v", u1)
	fmt.Printf("%+v", u2)
}
