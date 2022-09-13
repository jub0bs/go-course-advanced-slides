package main

import (
	"fmt"
	"io"
	"os"
)

// START OMIT
func main() {
	var (
		stdout io.Writer = os.Stdout
		stderr io.Writer = os.Stderr
	)
	fmt.Println(stdout == stderr) // HL
	fmt.Println(stdout != stderr) // HL
}
// END OMIT
