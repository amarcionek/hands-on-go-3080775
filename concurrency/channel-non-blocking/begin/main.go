// concurrency/channel-non-blocking/begin/main.go
package main

import (
	"fmt"
	"time"
)

func main() {
	// // declare a signal channel to read from
	// readChan := make(chan struct{})

	// // use a default case in select to prevent blocking
	// select {
	// case <-readChan:
	// 	fmt.Println("received from readChan")
	// default:
	// 	fmt.Println("no signal received")
	// }

	timeChan1 := time.After(200 * time.Millisecond)
	timeChan2 := time.After(400 * time.Millisecond)

	for {
		select {
		case <-timeChan1:
			fmt.Println("1")
			return
		case <-timeChan2:
			fmt.Println("2")
			return
		default:
			fmt.Println("default")
			time.Sleep(250 * time.Millisecond)
		}
	}
}
