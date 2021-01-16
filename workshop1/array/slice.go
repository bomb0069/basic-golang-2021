package main

import "fmt"

func main() {
	var colors = [5]string{"Red", "Blue", "Green", "Yellow", "Black"}
	fmt.Println("Colors = ", colors)

	a := colors[0:2]
	a = append(a, "XXX")
	a = append(a, "YYY")
	fmt.Println("A = ", a)
	a[0] = "New"
	fmt.Println("Colors,a = ", colors, a)

}
