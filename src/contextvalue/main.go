package main

import (
	"context"
	"errors"
	"fmt"
	"mymodulepath/auth"
	"os"
)

// START OMIT
func main() {
	ctx := auth.WithToken(context.Background(), "supersecretauthtoken") // HL
	if err := foo(ctx); err != nil {                                    // HL
		fmt.Fprintln(os.Stderr, err)
	}

	ctx = context.Background()
	if err := foo(ctx); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func foo(ctx context.Context) error {
	token, ok := auth.Token(ctx) // HL
	if !ok {
		return errors.New("no auth token")
	}
	// do something interesting with the token
	fmt.Println("token:", token)
	return nil
}

// END OMIT
