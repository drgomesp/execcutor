package main

import (
	"context"
	"log"

	x "github.com/drgomesp/execcutor"
)

func main() {
	// Run returns an error which you can capture and handle
	_ = x.Run(func(ctx context.Context, args ...string) error {
		log.Println("and the program runs here...")

		return nil
	})
}
