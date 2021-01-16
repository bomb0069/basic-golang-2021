package main

import (
	"fmt"
)

type Stringer interface {
	String() string
}

type Employee struct {
	Name string
	Age  int
}

func (emp Employee) String() string {
	return fmt.Sprintf("%s (%d)", emp.Name, emp.Age)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	emp := Employee{"bomb", 15}
	describe(emp)

}
