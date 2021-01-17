package main

import (
	"fmt"
	"log"
)

func main() {
	var numbers [5]int
	numbers[0] = 1
	numbers[1] = 2
	var colors = [2]string{"Red", "Blue"}
	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	for k, v := range colors {
		fmt.Println(k, " ", v)
	}

	// infinite loop
	//for {

	//}

	// for to 5
	var i = 0
	for i < 5 {
		i++
		log.Printf("Logging %v", i)
	}

}
