package execcutor

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	StatusUnknown = iota - 1
	StatusSuccess
	StatusError
)

type Entrypoint func(ctx context.Context, args ...string) error

// Run an entrypoint function that consumes os.Args
func Run(entrypoint Entrypoint) error {
	ctx := context.Background()
	sig := make(chan os.Signal, 1)

	go handleSignals(sig)

	return entrypoint(ctx, os.Args[1:]...)
}

func handleSignals(sig chan os.Signal) {
	func() {
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(sig)

		<-sig
		log.Println("attempting graceful shutdown...")

		// ...
		os.Exit(StatusUnknown)
	}()
}
