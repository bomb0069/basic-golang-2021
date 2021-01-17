package main

import (
	"fmt"
	"sort"
)

func main() {
	var intSlice = []int{11, 3, 4, 1, 5, 6, 2}

	fmt.Println(intSlice)

	sort.Ints(intSlice)
	fmt.Println(intSlice)
}
