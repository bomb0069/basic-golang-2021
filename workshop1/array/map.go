package main

import (
	"fmt"
)

func main() {
	numbers := make(map[string]int)
	numbers["one"] = 1
	if e2, found := numbers["two"]; found {
		fmt.Println("Found=", e2)
	} else {
		fmt.Println("Not Found")
	}
}
