package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	doSending(c)
	value := <-c
	fmt.Println(value)
}

func doSending(c chan<- int) {
	c <- 1
}
