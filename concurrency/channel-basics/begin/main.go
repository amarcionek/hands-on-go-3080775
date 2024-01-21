// concurrency/channels/begin/main.go
package main

import (
	"fmt"
)

// sum calculates and prints the sum of numbers
func sum(nums []int, ch chan<- int) {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	//fmt.Println("Result:", sum)
	ch <- sum
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	ch := make(chan int)

	// invoke the sum function as a goroutine
	go sum(nums, ch)

	// force main thread to sleep
	//time.Sleep(100 * time.Millisecond)

	result := <-ch

	fmt.Println("Result:", result)

	// Buffered channel
	ch2 := make(chan string)

	ch2 <- "Adam"
	ch2 <- "David"

	fmt.Println(<-ch2)
	fmt.Println(<-ch2)

}
