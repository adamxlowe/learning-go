package main

import (
	"fmt"
	"github.com/twmb/style"
	"time"
)

func main() {
	w := style.NewWriter()
	w.SetDelay(100 * time.Millisecond)

	fmt.Fprintln(w, "calling RecursiveChannels in main")
	channel := RecursiveChannels(0)
	received := <-channel
	fmt.Fprintf(w, "received %v from c in main\n", received)
}

func RecursiveChannels(recursionlevel int) <-chan int {
	w := style.NewWriter(style.Bright, style.XTerm256FG(recursionlevel+1))
	w.SetDelay(100 * time.Millisecond)

	returnChan := make(chan int)

	if recursionlevel == 10 {
		fmt.Fprintf(w, "sending to channel at recursion level %v\n", recursionlevel)
		returnChan <- 1
		return returnChan
	}

	fmt.Fprintf(w, "recursing and receiving from channel at recursion level %v\n", recursionlevel)
	channel := RecursiveChannels(recursionlevel + 1)
	received := <-channel
	fmt.Fprintf(w, "received value %v from channel at recursion level %v, now sending back my level\n", received, recursionlevel)

	return returnChan
}
