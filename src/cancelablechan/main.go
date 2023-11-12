package main

import (
	"context"
	"fmt"
	"time"
)

// START OMIT
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	printInts(ctx) // this function would typically be spawned as a goroutine
	// do more work (omitted)...
}

func printInts(ctx context.Context) error { // HL
	ticker := time.NewTicker(100 * time.Millisecond)
	var i int
	for {
		select {
		case <-ctx.Done(): // HL
			ticker.Stop()
			return ctx.Err()
		case <-ticker.C:
			fmt.Println(i)
			i++
		}
	}
	return nil
}

// END OMIT
