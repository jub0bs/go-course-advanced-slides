package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	printInts(ctx) // this function would typically be spawned as a goroutine
	// do more work (omitted)...
}

// START OMIT
// main is unchanged

func printInts(ctx context.Context) error {
	if deadline, exists := ctx.Deadline(); exists { // HL
		estimatedTimeToCompletion := 3 * time.Second
		if deadline.Sub(time.Now()) < estimatedTimeToCompletion { // HL
			return context.DeadlineExceeded
		}
	}
	ticker := time.NewTicker(100 * time.Millisecond)
	var i int
	for {
		select {
		case <-ctx.Done():
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
