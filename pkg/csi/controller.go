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
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strconv"

	"github.com/golang/protobuf/ptypes"

	"github.com/pborman/uuid"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/container-storage-interface/spec/lib/go/csi"
	utilexec "k8s.io/utils/exec"

	"github.com/batazor/shortlink/internal/logger/field"
)

const (
	deviceID           = "deviceID"
	maxStorageCapacity = tib
)

type accessType int

const (
	mountAccess accessType = iota
	blockAccess
)

// ControllerGetVolume gets a specific volume.
// The call is used for the CSI health check feature
// (https://github.com/kubernetes/enhancements/pull/1077) which we do not support yet.
func (d *driver) ControllerGetVolume(ctx context.Context, request *csi.ControllerGetVolumeRequest) (*csi.ControllerGetVolumeResponse, error) {
	panic("implement me")
}

func NewControllerServer(nodeID string) *controllerServer {
	return &controllerServer{
		caps: getControllerServiceCapabilities(
			[]csi.ControllerServiceCapability_RPC_Type{
				csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
				csi.ControllerServiceCapability_RPC_CREATE_DELETE_SNAPSHOT,
				csi.ControllerServiceCapability_RPC_LIST_SNAPSHOTS,
				csi.ControllerServiceCapability_RPC_CLONE_VOLUME,
				csi.ControllerServiceCapability_RPC_EXPAND_VOLUME,
			}),
		nodeID: nodeID,
	}
}

// CreateVolume creates a new volume from the given request. The function is idempotent.
func (d *driver) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	if err := d.validateControllerServiceRequest(csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME); err != nil {
		d.log.Error(fmt.Sprintf("invalid create snapshot req: %v", req))
		return nil, err
	}

	// Check arguments
	if len(req.GetName()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Name must be provided")
	}

	caps := req.GetVolumeCapabilities()
	if caps == nil {
		return nil, status.Error(codes.InvalidArgument, "Volume Capabilities must be provided")
	}

	// TODO: use real size
	d.log.Info("create volume called", field.Fields{
		"volume_name":         req.Name,
		"method":              "create_volume",
		"volume_capabilities": req.VolumeCapabilities,
	})

	// Keep a record of the requested access types.
	var accessTypeMount, accessTypeBlock bool

	for _, cap := range caps {
		if cap.GetBlock() != nil {
			accessTypeBlock = true
		}
		if cap.GetMount() != nil {
			accessTypeMount = true
		}
	}

	// A real driver would also need to check that the other
	// fields in VolumeCapabilities are sane. The check above is
	// just enough to pass the "[Testpattern: Dynamic PV (block
	// volmode)] volumeMode should fail in binding dynamic
	// provisioned PV to PVC" storage E2E test.

	if accessTypeBlock && accessTypeMount {
		return nil, status.Error(codes.InvalidArgument, "cannot have both block and mount access type")
	}

	var requestedAccessType accessType

	if accessTypeBlock {
		requestedAccessType = blockAccess
	} else {
		// Default to mount.
		requestedAccessType = mountAccess
	}

	// Check for maximum available capacity
	capacity := int64(req.GetCapacityRange().GetRequiredBytes())
	if capacity >= maxStorageCapacity {
		return nil, status.Errorf(codes.OutOfRange, "Requested capacity %d exceeds maximum allowed %d", capacity, maxStorageCapacity)
	}

	// Need to check for already existing volume name, and if found
	// check for the requested capacity and already allocated capacity
	if exVol, err := getVolumeByName(req.GetName()); err == nil {
		// Since err is nil, it means the volume with the same name already exists
		// need to check if the size of existing volume is the same as in new
		// request
		if exVol.VolSize < capacity {
			return nil, status.Errorf(codes.AlreadyExists, "Volume with the same name: %s but with different size already exist", req.GetName())
		}
		if req.GetVolumeContentSource() != nil {
			volumeSource := req.VolumeContentSource
			switch volumeSource.Type.(type) {
			case *csi.VolumeContentSource_Snapshot:
				if volumeSource.GetSnapshot() != nil && exVol.ParentSnapID != volumeSource.GetSnapshot().GetSnapshotId() {
					return nil, status.Error(codes.AlreadyExists, "existing volume source snapshot id not matching")
				}
			case *csi.VolumeContentSource_Volume:
				if volumeSource.GetVolume() != nil && exVol.ParentVolID != volumeSource.GetVolume().GetVolumeId() {
					return nil, status.Error(codes.AlreadyExists, "existing volume source volume id not matching")
				}
			default:
				return nil, status.Errorf(codes.InvalidArgument, "%v not a proper volume source", volumeSource)
			}
		}
		// TODO (sbezverk) Do I need to make sure that volume still exists?
		return &csi.CreateVolumeResponse{
			Volume: &csi.Volume{
				VolumeId:      exVol.VolID,
				CapacityBytes: int64(exVol.VolSize),
				VolumeContext: req.GetParameters(),
				ContentSource: req.GetVolumeContentSource(),
			},
		}, nil
	}

	volumeID := uuid.NewUUID().String()

	vol, err := createHostpathVolume(volumeID, req.GetName(), capacity, requestedAccessType)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create volume %v: %v", volumeID, err)
	}
	d.log.InfoWithContext(ctx, fmt.Sprintf("created volume %s at path %s", vol.VolID, vol.VolPath))

	if req.GetVolumeContentSource() != nil {
		path := getVolumePath(volumeID)
		volumeSource := req.VolumeContentSource
		switch volumeSource.Type.(type) {
		case *csi.VolumeContentSource_Snapshot:
			if snapshot := volumeSource.GetSnapshot(); snapshot != nil {
				err = loadFromSnapshot(capacity, snapshot.GetSnapshotId(), path, requestedAccessType)
				vol.ParentSnapID = snapshot.GetSnapshotId()
			}
		case *csi.VolumeContentSource_Volume:
			if srcVolume := volumeSource.GetVolume(); srcVolume != nil {
				err = loadFromVolume(capacity, srcVolume.GetVolumeId(), path, requestedAccessType)
				vol.ParentVolID = srcVolume.GetVolumeId()
			}
		default:
			err = status.Errorf(codes.InvalidArgument, "%v not a proper volume source", volumeSource)
		}
		if err != nil {
			if delErr := deleteHostpathVolume(volumeID); delErr != nil {
				d.log.InfoWithContext(ctx, fmt.Sprintf("deleting hostpath volume %v failed: %v", volumeID, delErr))
			}
			return nil, err
		}
		d.log.InfoWithContext(ctx, fmt.Sprintf("successfully populated volume %s", vol.VolID))
	}

	topologies := []*csi.Topology{&csi.Topology{
		Segments: map[string]string{TopologyKeyNode: d.cs.nodeID},
	}}

	return &csi.CreateVolumeResponse{
		Volume: &csi.Volume{
			VolumeId:           volumeID,
			CapacityBytes:      req.GetCapacityRange().GetRequiredBytes(),
			VolumeContext:      req.GetParameters(),
			ContentSource:      req.GetVolumeContentSource(),
			AccessibleTopology: topologies,
		},
	}, nil
}

// DeleteVolume deletes the given volume. The function is idempotent.
func (d *driver) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	// Check arguments
	if len(req.GetVolumeId()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume ID must be provided")
	}

	if err := d.validateControllerServiceRequest(csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME); err != nil {
		d.log.InfoWithContext(ctx, fmt.Sprintf("invalid delete volume req: %v", req))
		return nil, err
	}

	volId := req.GetVolumeId()
	if err := deleteHostpathVolume(volId); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete volume %v: %v", volId, err)
	}

	d.log.Info("delete volume called", field.Fields{
		"volume_id": req.VolumeId,
		"method":    "delete_volume",
	})

	return &csi.DeleteVolumeResponse{}, nil
}

// ControllerPublishVolume attaches the given volume to the node
func (d *driver) ControllerGetCapabilities(ctx context.Context, req *csi.ControllerGetCapabilitiesRequest) (*csi.ControllerGetCapabilitiesResponse, error) {
	return &csi.ControllerGetCapabilitiesResponse{
		Capabilities: d.cs.caps,
	}, nil
}

func (d *driver) ValidateVolumeCapabilities(ctx context.Context, req *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {

	// Check arguments
	if len(req.GetVolumeId()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume ID cannot be empty")
	}
	if len(req.VolumeCapabilities) == 0 {
		return nil, status.Error(codes.InvalidArgument, req.VolumeId)
	}

	if _, err := getVolumeByID(req.GetVolumeId()); err != nil {
		return nil, status.Error(codes.NotFound, req.GetVolumeId())
	}

	for _, cap := range req.GetVolumeCapabilities() {
		if cap.GetMount() == nil && cap.GetBlock() == nil {
			return nil, status.Error(codes.InvalidArgument, "cannot have both mount and block access type be undefined")
		}

		// A real driver would check the capabilities of the given volume with
		// the set of requested capabilities.
	}

	return &csi.ValidateVolumeCapabilitiesResponse{
		Confirmed: &csi.ValidateVolumeCapabilitiesResponse_Confirmed{
			VolumeContext:      req.GetVolumeContext(),
			VolumeCapabilities: req.GetVolumeCapabilities(),
			Parameters:         req.GetParameters(),
		},
	}, nil
}

// ControllerPublishVolume attaches the given volume to the node
func (d *driver) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	if req.VolumeId == "" {
		return nil, status.Error(codes.InvalidArgument, "ControllerPublishVolume Volume ID must be provided")
	}

	if req.NodeId == "" {
		return nil, status.Error(codes.InvalidArgument, "ControllerPublishVolume Node ID must be provided")
	}

	if req.VolumeCapability == nil {
		return nil, status.Error(codes.InvalidArgument, "ControllerPublishVolume Volume capability must be provided")
	}

	if req.Readonly {
		// TODO(arslan): we should return codes.InvalidArgument, but the CSI
		// test fails, because according to the CSI Spec, this flag cannot be
		// changed on the same volume. However we don't use this flag at all,
		// as there are no `readonly` attachable volumes.
		return nil, status.Error(codes.AlreadyExists, "read only Volumes are not supported")
	}

	d.log.Info("controller publish volume called", field.Fields{
		"volume_id": req.VolumeId,
		"node_id":   req.NodeId,
		"method":    "controller_publish_volume",
	})

	return &csi.ControllerPublishVolumeResponse{
		PublishContext: map[string]string{},
	}, nil
}

// ControllerUnpublishVolume deattaches the given volume from the node
func (d *driver) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	if req.VolumeId == "" {
		return nil, status.Error(codes.InvalidArgument, "ControllerUnpublishVolume Volume ID must be provided")
	}

	d.log.Info("controller unpublish volume called", field.Fields{
		"volume_id": req.VolumeId,
		"node_id":   req.NodeId,
		"method":    "controller_unpublish_volume",
	})

	return &csi.ControllerUnpublishVolumeResponse{}, nil
}

// GetCapacity returns the capacity of the storage pool
func (d *driver) GetCapacity(ctx context.Context, req *csi.GetCapacityRequest) (*csi.GetCapacityResponse, error) {
	d.log.Info("get capacity is not implemented", field.Fields{
		"params": req.Parameters,
		"method": "get_capacity",
	})

	return &csi.GetCapacityResponse{}, nil
}

// ListVolumes returns a list of all requested volumes
func (d *driver) ListVolumes(ctx context.Context, req *csi.ListVolumesRequest) (*csi.ListVolumesResponse, error) {
	d.log.Info("list volumes called", field.Fields{
		"max_entries":        req.MaxEntries,
		"req_starting_token": req.StartingToken,
		"method":             "list_volumes",
	})

	if req.StartingToken != "" {
		_, err := strconv.ParseInt(req.StartingToken, 10, 32)
		if err != nil {
			return nil, status.Errorf(codes.Aborted, "ListVolumes starting token %q is not valid: %s", req.StartingToken, err)
		}
	}

	return &csi.ListVolumesResponse{
		Entries: []*csi.ListVolumesResponse_Entry{},
	}, nil
}

// getSnapshotPath returns the full path to where the snapshot is stored
func getSnapshotPath(snapshotID string) string {
	return filepath.Join(dataRoot, fmt.Sprintf("%s%s", snapshotID, snapshotExt))
}

// CreateSnapshot uses tar command to create snapshot for hostpath volume. The tar command can quickly create
// archives of entire directories. The host image must have "tar" binaries in /bin, /usr/sbin, or /usr/bin.
func (d *driver) CreateSnapshot(ctx context.Context, req *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
	if err := d.validateControllerServiceRequest(csi.ControllerServiceCapability_RPC_CREATE_DELETE_SNAPSHOT); err != nil {
		d.log.InfoWithContext(ctx, fmt.Sprintf("invalid create snapshot req: %v", req))
		return nil, err
	}

	if len(req.GetName()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Name must be provided")
	}

	if len(req.GetSourceVolumeId()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "SourceVolumeId must be provided")
	}

	// Need to check for already existing snapshot name, and if found check for the
	// requested sourceVolumeId and sourceVolumeId of snapshot that has been created.
	if exSnap, err := getSnapshotByName(req.GetName()); err == nil {
		// Since err is nil, it means the snapshot with the same name already exists need
		// to check if the sourceVolumeId of existing snapshot is the same as in new request.
		if exSnap.VolID == req.GetSourceVolumeId() {
			// same snapshot has been created.
			return &csi.CreateSnapshotResponse{
				Snapshot: &csi.Snapshot{
					SnapshotId:     exSnap.Id,
					SourceVolumeId: exSnap.VolID,
					CreationTime:   &exSnap.CreationTime,
					SizeBytes:      exSnap.SizeBytes,
					ReadyToUse:     exSnap.ReadyToUse,
				},
			}, nil
		}
		return nil, status.Errorf(codes.AlreadyExists, "snapshot with the same name: %s but with different SourceVolumeId already exist", req.GetName())
	}

	volumeID := req.GetSourceVolumeId()
	hostPathVolume, ok := hostPathVolumes[volumeID]
	if !ok {
		return nil, status.Error(codes.Internal, "volumeID is not exist")
	}

	snapshotID := uuid.NewUUID().String()
	creationTime := ptypes.TimestampNow()
	volPath := hostPathVolume.VolPath
	file := getSnapshotPath(snapshotID)

	var cmd []string
	if hostPathVolume.VolAccessType == blockAccess {
		d.log.InfoWithContext(ctx, "Creating snapshot of Raw Block Mode Volume")
		cmd = []string{"cp", volPath, file}
	} else {
		d.log.InfoWithContext(ctx, "Creating snapshot of Filsystem Mode Volume")
		cmd = []string{"tar", "czf", file, "-C", volPath, "."}
	}
	executor := utilexec.New()
	out, err := executor.Command(cmd[0], cmd[1:]...).CombinedOutput()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed create snapshot: %v: %s", err, out)
	}

	d.log.InfoWithContext(ctx, fmt.Sprintf("create volume snapshot %s", file))
	snapshot := hostPathSnapshot{}
	snapshot.Name = req.GetName()
	snapshot.Id = snapshotID
	snapshot.VolID = volumeID
	snapshot.Path = file
	snapshot.CreationTime = *creationTime // nolint copylocks
	snapshot.SizeBytes = hostPathVolume.VolSize
	snapshot.ReadyToUse = true

	hostPathVolumeSnapshots[snapshotID] = snapshot // nolint copylocks

	return &csi.CreateSnapshotResponse{
		Snapshot: &csi.Snapshot{
			SnapshotId:     snapshot.Id,
			SourceVolumeId: snapshot.VolID,
			CreationTime:   &snapshot.CreationTime,
			SizeBytes:      snapshot.SizeBytes,
			ReadyToUse:     snapshot.ReadyToUse,
		},
	}, nil
}

// DeleteSnapshot will be called by the CO to delete a snapshot.
func (d *driver) DeleteSnapshot(ctx context.Context, req *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {
	// Check arguments
	if len(req.GetSnapshotId()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Snapshot ID must be provided")
	}

	if err := d.validateControllerServiceRequest(csi.ControllerServiceCapability_RPC_CREATE_DELETE_SNAPSHOT); err != nil {
		d.log.InfoWithContext(ctx, fmt.Sprintf("invalid delete snapshot req: %v", req))
		return nil, err
	}
	snapshotID := req.GetSnapshotId()
	d.log.InfoWithContext(ctx, fmt.Sprintf("deleting snapshot %s", snapshotID))
	path := getSnapshotPath(snapshotID)
	os.RemoveAll(path)
	delete(hostPathVolumeSnapshots, snapshotID)
	return &csi.DeleteSnapshotResponse{}, nil
}

// ListSnapshots returns the information about all snapshots on the storage
// system within the given parameters regardless of how they were created.
// ListSnapshots shold not list a snapshot that is being created but has not
// been cut successfully yet.
func (d *driver) ListSnapshots(ctx context.Context, req *csi.ListSnapshotsRequest) (*csi.ListSnapshotsResponse, error) {
	if err := d.validateControllerServiceRequest(csi.ControllerServiceCapability_RPC_LIST_SNAPSHOTS); err != nil {
		d.log.InfoWithContext(ctx, fmt.Sprintf("invalid list snapshot req: %v", req))
		return nil, err
	}

	// case 1: SnapshotId is not empty, return snapshots that match the snapshot id.
	if len(req.GetSnapshotId()) != 0 {
		snapshotID := req.SnapshotId
		if snapshot, ok := hostPathVolumeSnapshots[snapshotID]; ok {
			return convertSnapshot(snapshot), nil // nolint copylocks
		}
	}

	// case 2: SourceVolumeId is not empty, return snapshots that match the source volume id.
	if len(req.GetSourceVolumeId()) != 0 {
		for _, snapshot := range hostPathVolumeSnapshots { // nolint copylocks
			if snapshot.VolID == req.SourceVolumeId {
				return convertSnapshot(snapshot), nil // nolint copylocks
			}
		}
	}

	var snapshots []csi.Snapshot
	// case 3: no parameter is set, so we return all the snapshots.
	sortedKeys := make([]string, 0)
	for k := range hostPathVolumeSnapshots {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)

	for _, key := range sortedKeys {
		snap := hostPathVolumeSnapshots[key] // nolint copylocks
		snapshot := csi.Snapshot{
			SnapshotId:     snap.Id,
			SourceVolumeId: snap.VolID,
			CreationTime:   &snap.CreationTime,
			SizeBytes:      snap.SizeBytes,
			ReadyToUse:     snap.ReadyToUse,
		}
		snapshots = append(snapshots, snapshot)
	}

	var (
		ulenSnapshots = int32(len(snapshots))
		maxEntries    = req.MaxEntries
		startingToken int32
	)

	if v := req.StartingToken; v != "" {
		i, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return nil, status.Errorf(
				codes.Aborted,
				"startingToken=%d !< int32=%d",
				startingToken, math.MaxUint32)
		}
		startingToken = int32(i)
	}

	if startingToken > ulenSnapshots {
		return nil, status.Errorf(
			codes.Aborted,
			"startingToken=%d > len(snapshots)=%d",
			startingToken, ulenSnapshots)
	}

	// Discern the number of remaining entries.
	rem := ulenSnapshots - startingToken

	// If maxEntries is 0 or greater than the number of remaining entries then
	// set maxEntries to the number of remaining entries.
	if maxEntries == 0 || maxEntries > rem {
		maxEntries = rem
	}

	var (
		i       int
		j       = startingToken
		entries = make(
			[]*csi.ListSnapshotsResponse_Entry,
			maxEntries)
	)

	for i = 0; i < len(entries); i++ {
		entries[i] = &csi.ListSnapshotsResponse_Entry{
			Snapshot: &snapshots[j],
		}
		j++
	}

	var nextToken string
	if j < ulenSnapshots {
		nextToken = fmt.Sprintf("%d", j)
	}

	return &csi.ListSnapshotsResponse{
		Entries:   entries,
		NextToken: nextToken,
	}, nil
}

// ControllerExpandVolume is called from the resizer to increase the volume size.
func (d *driver) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest) (*csi.ControllerExpandVolumeResponse, error) {
	volID := req.GetVolumeId()
	if len(volID) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume ID must be provided")
	}

	capRange := req.GetCapacityRange()
	if capRange == nil {
		return nil, status.Error(codes.InvalidArgument, "Capacity range not provided")
	}

	capacity := int64(capRange.GetRequiredBytes())
	if capacity >= maxStorageCapacity {
		return nil, status.Errorf(codes.OutOfRange, "Requested capacity %d exceeds maximum allowed %d", capacity, maxStorageCapacity)
	}

	exVol, err := getVolumeByID(volID)
	if err != nil {
		// Assume not found error
		return nil, status.Errorf(codes.NotFound, "Could not get volume %s: %v", volID, err)
	}

	if exVol.VolSize < capacity {
		exVol.VolSize = capacity
		if err := updateHostpathVolume(volID, exVol); err != nil {
			return nil, status.Errorf(codes.Internal, "Could not update volume %s: %v", volID, err)
		}
	}

	return &csi.ControllerExpandVolumeResponse{
		CapacityBytes:         exVol.VolSize,
		NodeExpansionRequired: true,
	}, nil
}

func convertSnapshot(snap hostPathSnapshot) *csi.ListSnapshotsResponse { // nolint copylocks
	entries := []*csi.ListSnapshotsResponse_Entry{
		{
			Snapshot: &csi.Snapshot{
				SnapshotId:     snap.Id,
				SourceVolumeId: snap.VolID,
				CreationTime:   &snap.CreationTime,
				SizeBytes:      snap.SizeBytes,
				ReadyToUse:     snap.ReadyToUse,
			},
		},
	}

	rsp := &csi.ListSnapshotsResponse{
		Entries: entries,
	}

	return rsp
}

func (d *driver) validateControllerServiceRequest(c csi.ControllerServiceCapability_RPC_Type) error {
	if c == csi.ControllerServiceCapability_RPC_UNKNOWN {
		return nil
	}

	for _, cap := range d.cs.caps {
		if c == cap.GetRpc().GetType() {
			return nil
		}
	}

	return status.Errorf(codes.InvalidArgument, "unsupported capability %s", c)
}

func getControllerServiceCapabilities(cl []csi.ControllerServiceCapability_RPC_Type) []*csi.ControllerServiceCapability {
	var csc []*csi.ControllerServiceCapability

	for _, cap := range cl {
		csc = append(csc, &csi.ControllerServiceCapability{
			Type: &csi.ControllerServiceCapability_Rpc{
				Rpc: &csi.ControllerServiceCapability_RPC{
					Type: cap,
				},
			},
		})
	}

	return csc
}
