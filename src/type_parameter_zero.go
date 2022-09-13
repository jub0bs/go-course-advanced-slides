package main

import "fmt"

// START OMIT
func Last[E any](s []E) (E, bool) {
	if len(s) == 0 {
		var zero E // HL
		return zero, false
	}
	return s[len(s)-1], true
}

func main() {
	fmt.Println(Last([]int{})) // 0 false
}

// END OMIT
