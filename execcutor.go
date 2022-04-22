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

type RunFunc func(ctx context.Context, args []string) error
type ShutdownFunc func(ctx context.Context) error

// Run an entrypoint function that consumes os.Args
func Run(entrypoint RunFunc, shutdownFunc ShutdownFunc) {
	ctx := context.Background()
	sig := make(chan os.Signal, 1)

	go handleSignals(ctx, sig, shutdownFunc)

	if err := entrypoint(ctx, os.Args); err != nil {
		os.Exit(StatusError)
	}

	os.Exit(StatusSuccess)
}

func handleSignals(ctx context.Context, sig chan os.Signal, shutdownFunc ShutdownFunc) {
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(sig)

	<-sig
	log.Println("attempting graceful shutdown...")
	if err := shutdownFunc(ctx); err != nil {
		log.Println(err)
		os.Exit(StatusError)
	}

	os.Exit(StatusUnknown)
}
