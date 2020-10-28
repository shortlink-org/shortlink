package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	csi_driver "github.com/batazor/shortlink/pkg/csi"
	"github.com/batazor/shortlink/pkg/csi/di"
)

// TODO: Use cobra
var (
	endpoint          = flag.String("endpoint", "unix://tmp/csi.sock", "CSI endpoint")
	driverName        = flag.String("drivername", "hostpath.csi.k8s.io", "name of the driver")
	nodeID            = flag.String("nodeid", "", "node id")
	ephemeral         = flag.Bool("ephemeral", false, "publish volumes in ephemeral mode even if kubelet did not ask for it (only needed for Kubernetes 1.15)")
	maxVolumesPerNode = flag.Int64("maxvolumespernode", 0, "limit of volumes per node")
	showVersion       = flag.Bool("version", false, "Show version.")
	// Set by the build process
	version = ""
)

func init() {
	flag.Set("logtostderr", "true")
}

func main() {
	flag.Parse()

	// Create a new context
	ctx := context.Background()

	// Init a new service
	s, cleanup, err := di.InitializeSCIDriver(ctx)
	if err != nil { // TODO: use as helpers
		if r, ok := err.(*net.OpError); ok {
			panic(fmt.Errorf("address %s already in use. Set GRPC_SERVER_PORT enviroment", r.Addr.String()))
		}

		panic(err)
	}

	// Run CSI Driver
	drv, err := csi_driver.NewDriver(*endpoint, *nodeID, s.Log)
	if err != nil {
		s.Log.Fatal(err.Error())
	}

	if err := drv.Run(ctx); err != nil {
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
