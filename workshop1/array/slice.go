package main

import "fmt"

func main() {
	// Array
	var colors = [5]string{"Red", "Blue", "Green", "Yellow", "Black"}

	// Slice
	var colorSlice = []string{"Red", "Blue", "Green", "Yellow", "Black"}

	// Make
	var colorMake = make([]string, 5)

	fmt.Println("Colors = ", colors)
	fmt.Println("Colors = ", colorSlice)
	fmt.Println("Colors = ", colorMake)

	a := colors[0:2]
	a = append(a, "XXX")
	a = append(a, "YYY")
	fmt.Println("A = ", a)
	a[0] = "New"
	fmt.Println("Colors,a = ", colors, a)

}
