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
	for v := range time.Tick(500 * time.Millisecond) { // HL
		fmt.Println(v)
	}
}

// END OMIT
