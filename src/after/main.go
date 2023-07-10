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
	select {
	case i := <-c:
		fmt.Println(i)
	case <-time.After(500 * time.Millisecond): // HL
		fmt.Println("timed out")
	}
}

// END OMIT
