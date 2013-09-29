package main

import (
	"fmt"
	"github.com/twmb/style"
	"time"
)

func main() {
	c := make(chan int)
	w := style.NewWriter()
	w.SetDelay(100 * time.Millisecond)
	go doSending(c)

	fmt.Fprintln(w, "receiving from c")
	value := <-c
	fmt.Fprintf(w, "received %v from c in main\n", value)
}

func doSending(c chan<- int) {
	w := style.NewWriter(style.Bright, style.Cyan)
	w.SetDelay(100 * time.Millisecond)
	fmt.Fprintln(w, "sending to c")
	c <- 1
}
