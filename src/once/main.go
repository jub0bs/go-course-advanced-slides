package main

import (
	"fmt"
	"sync"
)

func main() {
	// START OMIT
	var once sync.Once // HL
	var wg sync.WaitGroup
	const n = 10
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(j int) {
			defer wg.Done()
			f := func() { fmt.Printf("Only once: %d\n", j) }
			once.Do(f) // HL
		}(i)
	}
	wg.Wait()
	// END OMIT
}
