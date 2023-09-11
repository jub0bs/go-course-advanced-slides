package main

import (
	"fmt"
)

// START OMIT
func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Value from recovered panic:", r)
		}
	}()
	fmt.Println("work begins")
	panic("k8s goes brrrr...")
}

// END OMIT
