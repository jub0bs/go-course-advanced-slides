package main

import (
	"fmt"
	"time"
)

func main() {
	var c chan int
	foo(c)
}

// START OMIT
func foo(c <-chan int) {
	timer := time.NewTimer(500 * time.Millisecond) // HL
	select {
	case i := <-c:
		timer.Stop() // HL
		fmt.Println(i)
	case <-timer.C: // HL
		fmt.Println("timed out")
	}
}

// END OMIT
