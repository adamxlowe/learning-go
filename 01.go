package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	c <- 1       // send to the channel
	value := <-c // then receive what you just sent
	fmt.Println(value)
}
