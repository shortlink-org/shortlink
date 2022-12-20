package handle_signal

import (
	"os"
	"os/signal"
	"syscall"
)

func WaitExitSignal() os.Signal {
	sigs := make(chan os.Signal, 3)
	signal.Notify(
		sigs,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	return <-sigs
}
