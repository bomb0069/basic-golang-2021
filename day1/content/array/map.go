package main

import (
	"fmt"
)

func main() {

	numbers := make(map[string]int)
	numbers["one"] = 1
	fmt.Println("Numbers =", numbers)

	n1 := numbers["one"]
	fmt.Println("n1 =", n1)

	n2 := numbers["two"]
	fmt.Println("n2 =", n2)

	if e2, found := numbers["two"]; found {
		fmt.Println("Found = ", e2)
	} else {
		fmt.Println("Not Found")
	}
}
