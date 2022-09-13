package main

import (
	"fmt"
)

// START OMIT
type Location struct {
	_    struct{} // HL
	Lat  float64
	Long float64
}

func main() {
	strasbourg := Location{48.5734, 7.7521} // doesn't compile // HL
	fmt.Println(strasbourg)
}

// END OMIT
