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

	go doSending(c, 0)

	fmt.Fprintln(w, "receiving from c in main")
	value := <-c
	fmt.Fprintf(w, "received %v from c in main\n", value)
}

func doSending(c chan int, recursionlevel int) {
	w := style.NewWriter(style.Bright, style.XTerm256FG(recursionlevel+1))
	w.SetDelay(100 * time.Millisecond)

	if recursionlevel == 10 {
		fmt.Fprintf(w, "sending to c at recursion level %v\n", recursionlevel)
		c <- recursionlevel
	} else {
		fmt.Fprintf(w, "recursing and receiving from channel at recursion level %v\n", recursionlevel)
		doSending(c, recursionlevel+1)
		received := <-c
		fmt.Fprintf(w, "received value %v from channel at recursion level %v, now sending back my level\n", received, recursionlevel)
		c <- recursionlevel
	}
}
