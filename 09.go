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

	for received := range RecursiveChannels(0) {
		fmt.Fprintf(w, "received %v from c in main\n", received)
	}
}

func RecursiveChannels(recursionlevel int) <-chan int {
	w := style.NewWriter(style.Bright, style.XTerm256FG(recursionlevel+1))
	w.SetDelay(100 * time.Millisecond)

	returnChan := make(chan int)

	go func() {
		defer close(returnChan)
		if recursionlevel < 10 {
			fmt.Fprintf(w, "recursing and receiving from channel at recursion level %v\n", recursionlevel)
			for received := range RecursiveChannels(recursionlevel + 1) {
				fmt.Fprintf(w, "received value %v from channel at recursion level %v, now sending back my level\n", received, recursionlevel)
			}
		} else {
			fmt.Fprintf(w, "end of recursion: sending to channel at recursion level %v\n", recursionlevel)
		}
		returnChan <- recursionlevel
	}()

	return returnChan
}
