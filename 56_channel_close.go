package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// close the channel meaning no more values will be sent 
	// consider that only sender should close a channel 
	// also closing channel is not nessecarly 
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// this for loop reads from channel until it will be closed 
	for i := range c {
		fmt.Println(i)
	}
}
