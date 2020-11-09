/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package csi_driver

import (
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc"

	"github.com/batazor/shortlink/internal/logger"
)

const (
	// DefaultDriverName defines the name that is used in Kubernetes and the CSI
	// system for the canonical, official name of this plugin
	DefaultDriverName = "shrts.csi.k8s.io"
)

type identityServer struct {
	name string
	log  logger.Logger
}

type nodeServer struct {
	nodeID            string
	maxVolumesPerNode int64
}

type controllerServer struct {
	caps   []*csi.ControllerServiceCapability
	nodeID string
}

type driver struct {
	name              string
	nodeID            string
	endpoint          string
	maxVolumesPerNode int64

	log logger.Logger
	srv *grpc.Server

	ids *identityServer
	ns  *nodeServer
	cs  *controllerServer

	// ready defines whether the driver is ready to function. This value will
	// be used by the `Identity` service via the `Probe()` method.
	readyMu sync.Mutex // protects ready
	ready   bool
}

// NewDriver returns a CSI plugin that contains the necessary gRPC
// interfaces to interact with Kubernetes over unix domain sockets for
// managing ShortLink Storage
func NewDriver(log logger.Logger, driverName, nodeID, endpoint string, maxVolumesPerNode int64) (*driver, error) {
	if driverName == "" {
		return nil, errors.New("no driver name provided")
	}

	if nodeID == "" {
		return nil, errors.New("no node id provided")
	}

	if endpoint == "" {
		return nil, errors.New("no driver endpoint provided")
	}

	if err := os.MkdirAll(dataRoot, 0750); err != nil {
		return nil, fmt.Errorf("failed to create dataRoot: %v", err)
	}

	log.Info(fmt.Sprintf("Driver: %v ", driverName))
	log.Info(fmt.Sprintf("Version: %s", vendorVersion))

	return &driver{
		name:              driverName,
		nodeID:            nodeID,
		endpoint:          endpoint,
		maxVolumesPerNode: maxVolumesPerNode,
		log:               log,
	}, nil
}
