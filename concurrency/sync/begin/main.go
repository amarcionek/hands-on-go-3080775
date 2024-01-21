// concurrency/sync/begin/main.go
package main

import (
	"fmt"
	"sync"
)

func main() {
	// given a list of names
	names := []string{"Adam", "David", "Marcionek"}

	// initialize a map to store the length of each name
	namesMap := make(map[string]int)

	// initialize a wait group for the goroutines that will be launched
	var wg sync.WaitGroup

	var mu sync.Mutex

	// launch a goroutine for each name we want to process
	for _, name := range names {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			namesMap[name] = len(name)
		}(name)
	}

	// wait for all goroutines to finish
	wg.Wait()

	// print the map
	fmt.Println(namesMap)
}
