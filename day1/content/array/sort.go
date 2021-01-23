package main

import (
	"fmt"
	"sort"
)

func main() {
	var intSlice = []int{11, 3, 4, 1, 5, 6, 2}
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	var stringSlice = []string{"wat", "sivarat", "bomb", "karan"}
	sort.Strings(stringSlice)
	fmt.Println(stringSlice)

}
