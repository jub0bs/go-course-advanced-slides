package main

import "fmt"

// START OMIT
func Equal[E comparable](s1, s2 []E) bool { // HL
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func main() {
	s1 := []int{4, 8, 15, 16, 23, 42}
	s2 := []string{"foo", "bar", "baz"}
	fmt.Println(Equal(s1, s1), Equal(s2, s2))
	// fmt.Println(Equal(s1, s2)) // compilation error
}
// END OMIT
