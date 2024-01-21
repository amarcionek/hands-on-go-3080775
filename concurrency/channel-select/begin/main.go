// concurrency/channel-select/begin/main.go
package main

import (
	"fmt"
	"time"
)

func main() {
	// declare an empty struct channel for signaling
	sigCh := make(chan struct{})

	// declare a timer channel
	timCh := time.After(2 * time.Second)

	// launch a goroutine to signal after 1 second
	go func() {
		time.Sleep(1 * time.Second)
		sigCh <- struct{}{}
	}()

	// wait for a signal on either channel
	select {
	case <-sigCh:
		fmt.Println("received from sig")
	case <-timCh:
		fmt.Println("received from tim")
	}
}
