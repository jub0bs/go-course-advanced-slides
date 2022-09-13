package main

import (
	"fmt"
)

// START OMIT
func main() {
	ch := make(chan int)
	ch <- 1
	fmt.Println(<-ch)
}

// END OMIT
