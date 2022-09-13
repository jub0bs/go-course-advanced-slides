package main

import (
	"fmt"
	"sync"
)

// START OMIT
func main() {
	const n = 8
	results := make([]int, n) // HL
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		i := i
		go func() {
			defer wg.Done()
			results[i] = square(i) // HL
		}()
	}
	wg.Wait()
	fmt.Println(results)

}

func square(i int) int {
	return i * i
}

// END OMIT
