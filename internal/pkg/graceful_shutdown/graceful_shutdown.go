package graceful_shutdown

import (
	"os"
	"os/signal"
	"syscall"
)

func GracefulShutdown() os.Signal {
	shutdownSignals := make(chan os.Signal, 1)

	signal.Notify(
		shutdownSignals,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)

	return <-shutdownSignals
}
