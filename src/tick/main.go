package main

import (
	"fmt"
	"time"
)

func main() {
	foo()
}

// START OMIT
func foo() {
	c := time.Tick(500 * time.Millisecond)
	for v := range c {
		fmt.Println(v)
	}
}

// END OMIT
