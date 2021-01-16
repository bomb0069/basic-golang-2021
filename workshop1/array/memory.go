package main

import (
	"fmt"
)

func main() {
	var colorSlice = []string{"Red", "Blue", "Green", "Yellow", "Black"}

	for i, v := range colorSlice {
		fmt.Printf("%v at %v\n", v, &colorSlice[i])
	}
}
