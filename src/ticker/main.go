package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	defer signal.Reset(os.Interrupt)
	foo(ctx)
}

// START OMIT
func foo(ctx context.Context) error {
	ticker := time.NewTicker(500 * time.Millisecond) // HL
	for {
		select {
		case <-ctx.Done():
			ticker.Stop() // HL
			return ctx.Err()
		case v := <-ticker.C: // HL
			fmt.Println(v)
			return nil
		}
	}
}

// END OMIT
