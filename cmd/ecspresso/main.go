package main

import (
	"context"
	"errors"
	"os"
	"os/signal"

	"github.com/kayac/ecspresso/v2"
)

var Version string

func main() {
	ecspresso.Version = Version
	ctx, stop := signal.NotifyContext(context.Background(), trapSignals...)
	defer stop()

	exitCode, err := ecspresso.CLI(ctx, ecspresso.ParseCLIv2)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			ecspresso.LogWarn("Interrupted")
		} else {
			ecspresso.LogError("FAILED. %s", err)
		}
	}
	os.Exit(exitCode)
}
