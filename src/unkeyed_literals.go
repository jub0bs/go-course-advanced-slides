package main

import "fmt"

// START1 OMIT
type Location struct {
	Lat  float64
	Long float64
}

// END1 OMIT

func main() {
	// START2 OMIT
	strasbourg := Location{48.5734, 7.7521}       // unkeyed literal (avoid!)
	paris := Location{Lat: 48.8566, Long: 2.3522} // keyed literal
	fmt.Println(strasbourg, paris)
	// END2 OMIT
}
