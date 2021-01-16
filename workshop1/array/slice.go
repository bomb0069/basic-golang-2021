package main

import "fmt"

func main() {
	var colors = [3]string{"Red", "Blue", "Green"}
	fmt.Println("Colors = ", colors)

	a := colors[0:2]
	fmt.Println("A = ", a)
	a[0] = "New"
	fmt.Println("Colors,a = ", colors, a)

}
