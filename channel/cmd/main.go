package main

import (
	"fmt"
	"time"
)

func main() {
	var Ball int
	table := make(chan int)
	for i := 0; i < 100; i++ {
		go player(i, table)
	}

	table <- Ball
	fmt.Printf("table %v <- Ball %v \n", table, Ball)
	time.Sleep(1 * time.Second)
	fmt.Printf("%v \n", <-table)
}

func player(play int, table chan int) {
	for {
		ball := <-table
		ball++
		fmt.Printf("%v - %v => %v \n", table, play, ball)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}
