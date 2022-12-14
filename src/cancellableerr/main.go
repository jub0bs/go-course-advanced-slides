package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	const timeout = 2 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	printInts(ctx) // this function would typically be spawned as a goroutine
	// do more work (omitted)...
}

// START OMIT
// main is unchanged

func printInts(ctx context.Context) error { // HL
	for i := 0; ; i++ {
		if err := ctx.Err(); err != nil { // HL
			return err
		}
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
	return nil
}

// END OMIT
