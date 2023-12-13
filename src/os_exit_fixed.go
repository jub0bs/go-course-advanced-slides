package main

import (
	"fmt"
	"log"
	"strconv"
)

// START OMIT
func main() {
	if err := foo(); err != nil {
		log.Fatal(err)
	}
}

func foo() error {
	defer fmt.Println("important cleanup")
	return bar("zzz")
}

func bar(s string) error {
	if _, err := strconv.Atoi(s); err != nil {
		return err // HL
	}
	return nil
}

// END OMIT
