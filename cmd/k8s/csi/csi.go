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

func main() {
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

	// TODO: Use cobra
	var (
		endpoint = flag.String("endpoint", "unix:///var/kubelet/plugins/"+csi_driver.DefaultDriverName+"/csi.sock", "CSI endpoint")
		//driverName = flag.String("driver-name", csi_driver.DefaultDriverName, "Name for the driver")
		// TODO: add version package
		//version = flag.Bool("version", false, "Print the version and exit.")
	)

	flag.Parse()

	// Run CSI Driver
	drv, err := csi_driver.NewDriver(*endpoint, s.Log)
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
