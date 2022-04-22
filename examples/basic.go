package main

import (
	"context"
	"log"

	x "github.com/drgomesp/execcutor"
)

func main() {
	// Run returns an error which you can capture and handle
	x.Run(func(ctx context.Context, args []string) error {
		log.Println("my code here")
		return nil
	}, func(ctx context.Context) error {
		log.Println("shutdown")
		return nil
	})
}
