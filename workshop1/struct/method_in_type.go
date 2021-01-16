package main

import (
	"fmt"
)

type bomb int

func (h bomb) sleep() {
	fmt.Printf("Sleeping")
}

func main() {
	var my bomb
	my.sleep()
}
