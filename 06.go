package main

import (
	"fmt"
	"github.com/twmb/style"
)

func main() {
	w := style.NewWriter()
	fmt.Fprintln(w, "calling RecursiveChannels in main")

	channel := RecursiveChannels(0)
	received := <-channel
	fmt.Fprintf(w, "received %v from c in main\n", received)
}

func RecursiveChannels(recursionlevel int) <-chan int {
	w := style.NewWriter(style.Bright, style.XTerm256FG(recursionlevel+1))

	returnChan := make(chan int)

	go func() {
		if recursionlevel < 10 {
			fmt.Fprintf(w, "recursing and receiving from channel at recursion level %v\n", recursionlevel)
			channel := RecursiveChannels(recursionlevel + 1)
			received := <-channel
			fmt.Fprintf(w, "received value %v from channel at recursion level %v, now sending back my level\n", received, recursionlevel)
		} else {
			fmt.Fprintf(w, "end of recursion: sending to channel at recursion level %v\n", recursionlevel)
		}
		returnChan <- recursionlevel
	}()

	return returnChan
}
