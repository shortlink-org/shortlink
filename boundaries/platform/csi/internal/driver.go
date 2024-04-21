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

package internal

import (
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc"

	"github.com/shortlink-org/shortlink/pkg/logger"
)

const (
	// DefaultDriverName defines the name that is used in Kubernetes and the CSI
	// system for the canonical, official name of this plugin
	DefaultDriverName = "shrts.csi.k8s.io"
)

type identityServer struct {
	log  logger.Logger
	name string
}

type nodeServer struct {
	nodeID            string
	maxVolumesPerNode int64
}

type controllerServer struct {
	nodeID string
	caps   []*csi.ControllerServiceCapability
}

type driver struct {
	log               logger.Logger
	srv               *grpc.Server
	ids               *identityServer
	ns                *nodeServer
	cs                *controllerServer
	name              string
	nodeID            string
	endpoint          string
	maxVolumesPerNode int64
	mu                sync.Mutex
	ready             bool
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

	if err := os.MkdirAll(dataRoot, 0o750); err != nil { //nolint:mnd
		return nil, fmt.Errorf("failed to create dataRoot: %w", err)
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
