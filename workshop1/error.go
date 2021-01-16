package demo

import (
	"errors"
	"fmt"
)

func someShort(x int) (int, error) {

	err := errors.New("Normal error")

	return 0, err
}

func some(x int) (result int, err error) {

	err = errors.New("Normal error")

	return
}

func main() {
	err := errors.New("Normal error")
	if err != nil {
		fmt.Println(err)
	}
}
