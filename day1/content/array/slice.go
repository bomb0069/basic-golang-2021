package main

import "fmt"

func main() {
	// Array
	var colors = [5]string{"Red", "Blue", "Green", "Yellow", "Black"}

	// Slice
	var colorSlice = []string{"Red", "Blue", "Green", "Yellow", "Black"}

	// Make
	var colorMake = make([]string, 5)

	fmt.Println("colors = 	", colors)
	fmt.Println("colorSlice = 	", colorSlice)
	fmt.Println("colorMake = 	", colorMake)

	a := colors[0:2]
	fmt.Println("a = ", a)

	a = append(a, "XXX")
	fmt.Println("a = ", a)
	a = append(a, "YYY")
	fmt.Println("a = ", a)
	a[0] = "New"
	fmt.Println("Colors,a = ", colors, a)

}
