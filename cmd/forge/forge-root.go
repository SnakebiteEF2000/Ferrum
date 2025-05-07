package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

type Forge struct {
	InstanceName  string
	State         int
	IsSiteManager bool
	ctx           context.Context
}

var forge = Forge{
	InstanceName:  "randomize",
	State:         0,
	IsSiteManager: false,
}

func run() int {
	rootCtx, rootCancel := context.WithCancel(context.Background())
	defer rootCancel()

	forge.ctx = rootCtx

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	_, _ = OpenBadgerCon()

	go func() {
		sig := <-sigChan
		slog.Warn("signal cought", "signal", sig)
		rootCancel()
	}()

	return 0
}
