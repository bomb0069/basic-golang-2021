package main

import (
	"fmt"
	"io/ioutil"
)

func panicHandler() {
	err := recover()
	if err == "some error" {
		fmt.Println("panicHandler - Try to recover from panic")
	}
	fmt.Println("panicHandler - Do it at the end")
}
func main() {

	deferAtEnd()

	deferWhenError()

}

func deferAtEnd() {
	fmt.Println("deferAtEnd - Start")
	defer panicHandler()
	fmt.Println("deferAtEnd - After defer panicHandler")

	fmt.Println("deferAtEnd - doSomething")

	fmt.Println("deferAtEnd - Start")
}

func deferWhenError() {
	fmt.Println("deferWhenError - Start")
	defer panicHandler()
	fmt.Println("deferWhenError - After defer panicHandler")

	b, err := ioutil.ReadFile("try_panic.go")
	fmt.Println("deferWhenError - After readfile")
	if err != nil {
		fmt.Println("deferWhenError - before panic")
		panic("some error")
		fmt.Println("deferWhenError - after panic")
	}
	fmt.Println(string(b))
	fmt.Println("deferAtEnd - After doSomething")

	fmt.Println("deferWhenError - Start")
}
