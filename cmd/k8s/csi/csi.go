package main

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	csi_driver "github.com/batazor/shortlink/pkg/csi"
	"github.com/batazor/shortlink/pkg/csi/di"
)

func main() {
	// Init a new service
	s, cleanup, err := di.InitializeSCIDriver()
	if err != nil { // TODO: use as helpers
		var typeErr *net.OpError
		if errors.As(err, &typeErr) {
			panic(fmt.Errorf("address %s already in use. Set GRPC_SERVER_PORT environment", typeErr.Addr.String()))
		}

		panic(err)
	}

	// TODO: Use cobra
	var (
		endpoint          = flag.String("endpoint", "unix://tmp/csi.sock", "CSI endpoint")
		nodeID            = flag.String("nodeid", "", "node id")
		maxVolumesPerNode = flag.Int64("maxvolumespernode", 0, "limit of volumes per node")
	)

	flag.Parse()

	// Run CSI Driver
	driver, err := csi_driver.NewDriver(s.Log, csi_driver.DefaultDriverName, *nodeID, *endpoint, *maxVolumesPerNode)
	if err != nil {
		s.Log.Fatal(fmt.Sprintf("Failed to initialize driver: %s", err.Error()))
	}
	if err := driver.Run(s.Ctx); err != nil {
		s.Log.Fatal(err.Error())
	}

	s.Log.Info("success run CSI plugin")

	// Handle SIGINT and SIGTERM.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	// Stop the service gracefully.
	cleanup()
}
