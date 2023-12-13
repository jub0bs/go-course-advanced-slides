package main

import (
	"fmt"
	"log"
	"strconv"
)

// START OMIT
func main() {
	foo()
}

func foo() {
	defer fmt.Println("important cleanup") // never executed // HL
	bar("zzz")
}

func bar(s string) {
	_, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err) // HL
	}
}

// END OMIT
