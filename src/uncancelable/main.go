package main

import (
	"fmt"
	"sync"
	"time"
)

// START OMIT
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		printInts()
	}()
	// do more work (omitted)...
	wg.Wait()
}

func printInts() {
	for i := 0; ; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

// END OMIT
