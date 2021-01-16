package demo

import (
	"errors"
	"fmt"
)

func someFmt(x int) (int, error) {

	return 0, fmt.Errorf("Normal error")
}

func someShort(x int) (int, error) {

	err := errors.New("Normal error")

	return 0, err
}

func some(x int) (result int, err error) {

	err = errors.New("Normal error")

	return
}

func main() {
	_, err := some(5)
	if err != nil {
		fmt.Println(err)
	}
}
