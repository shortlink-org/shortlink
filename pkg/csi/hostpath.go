package csi_driver

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang/glog"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/kubernetes/pkg/volume/util/volumepathhandler"
	utilexec "k8s.io/utils/exec"

	"github.com/batazor/shortlink/internal/logger"
)

const (
	kib    int64 = 1024
	mib    int64 = kib * 1024
	gib    int64 = mib * 1024
	gib100 int64 = gib * 100
	tib    int64 = gib * 1024
	tib100 int64 = tib * 100
)

type hostPath struct {
	name              string
	nodeID            string
	version           string
	endpoint          string
	ephemeral         bool
	maxVolumesPerNode int64

	log logger.Logger

	ids *identityServer
	ns  *nodeServer
	cs  *controllerServer
}

type hostPathVolume struct {
	VolName       string     `json:"volName"`
	VolID         string     `json:"volID"`
	VolSize       int64      `json:"volSize"`
	VolPath       string     `json:"volPath"`
	VolAccessType accessType `json:"volAccessType"`
	ParentVolID   string     `json:"parentVolID,omitempty"`
	ParentSnapID  string     `json:"parentSnapID,omitempty"`
	Ephemeral     bool       `json:"ephemeral"`
}

type hostPathSnapshot struct {
	Name         string              `json:"name"`
	Id           string              `json:"id"`
	VolID        string              `json:"volID"`
	Path         string              `json:"path"`
	CreationTime timestamp.Timestamp `json:"creationTime"`
	SizeBytes    int64               `json:"sizeBytes"`
	ReadyToUse   bool                `json:"readyToUse"`
}

var (
	vendorVersion = "dev"

	hostPathVolumes         map[string]hostPathVolume
	hostPathVolumeSnapshots map[string]hostPathSnapshot
)

const (
	// Directory where data for volumes and snapshots are persisted.
	// This can be ephemeral within the container or persisted if
	// backed by a Pod volume.
	dataRoot = "/csi-data-dir"

	// Extension with which snapshot files will be saved.
	snapshotExt = ".snap"
)

func init() {
	hostPathVolumes = map[string]hostPathVolume{}
	hostPathVolumeSnapshots = map[string]hostPathSnapshot{}
}

func getSnapshotID(file string) (bool, string) {
	glog.V(4).Infof("file: %s", file)
	// Files with .snap extension are volumesnapshot files.
	// e.g. foo.snap, foo.bar.snap
	if filepath.Ext(file) == snapshotExt {
		return true, strings.TrimSuffix(file, snapshotExt)
	}
	return false, ""
}

func discoverExistingSnapshots() {
	glog.V(4).Infof("discovering existing snapshots in %s", dataRoot)
	files, err := ioutil.ReadDir(dataRoot)
	if err != nil {
		glog.Errorf("failed to discover snapshots under %s: %v", dataRoot, err)
	}
	for _, file := range files {
		isSnapshot, snapshotID := getSnapshotID(file.Name())
		if isSnapshot {
			glog.V(4).Infof("adding snapshot %s from file %s", snapshotID, getSnapshotPath(snapshotID))
			hostPathVolumeSnapshots[snapshotID] = hostPathSnapshot{
				Id:         snapshotID,
				Path:       getSnapshotPath(snapshotID),
				ReadyToUse: true,
			}
		}
	}
}

func (hp *hostPath) Run(ctx context.Context) error {
	// Create GRPC servers
	hp.ids = NewIdentityServer(hp.name, hp.version, hp.log)
	hp.ns = NewNodeServer(hp.nodeID, hp.ephemeral, hp.maxVolumesPerNode)
	hp.cs = NewControllerServer(hp.ephemeral, hp.nodeID)

	discoverExistingSnapshots()
	s := NewNonBlockingGRPCServer()
	s.Start(hp.endpoint, hp.ids, hp.cs, hp.ns)
	s.Wait()

	return nil
}

func getVolumeByID(volumeID string) (hostPathVolume, error) {
	if hostPathVol, ok := hostPathVolumes[volumeID]; ok {
		return hostPathVol, nil
	}
	return hostPathVolume{}, fmt.Errorf("volume id %s does not exist in the volumes list", volumeID)
}

func getVolumeByName(volName string) (hostPathVolume, error) {
	for _, hostPathVol := range hostPathVolumes {
		if hostPathVol.VolName == volName {
			return hostPathVol, nil
		}
	}
	return hostPathVolume{}, fmt.Errorf("volume name %s does not exist in the volumes list", volName)
}

func getSnapshotByName(name string) (hostPathSnapshot, error) {
	for _, snapshot := range hostPathVolumeSnapshots {
		if snapshot.Name == name {
			return snapshot, nil
		}
	}
	return hostPathSnapshot{}, fmt.Errorf("snapshot name %s does not exist in the snapshots list", name)
}

// getVolumePath returns the canonical path for hostpath volume
func getVolumePath(volID string) string {
	return filepath.Join(dataRoot, volID)
}

// createVolume create the directory for the hostpath volume.
// It returns the volume path or err if one occurs.
func createHostpathVolume(volID, name string, cap int64, volAccessType accessType, ephemeral bool) (*hostPathVolume, error) {
	path := getVolumePath(volID)

	switch volAccessType {
	case mountAccess:
		err := os.MkdirAll(path, 0777)
		if err != nil {
			return nil, err
		}
	case blockAccess:
		executor := utilexec.New()
		size := fmt.Sprintf("%dM", cap/mib)
		// Create a block file.
		out, err := executor.Command("fallocate", "-l", size, path).CombinedOutput()
		if err != nil {
			return nil, fmt.Errorf("failed to create block device: %v, %v", err, string(out))
		}

		// Associate block file with the loop device.
		volPathHandler := volumepathhandler.VolumePathHandler{}
		_, err = volPathHandler.AttachFileDevice(path)
		if err != nil {
			// Remove the block file because it'll no longer be used again.
			if err2 := os.Remove(path); err2 != nil {
				glog.Errorf("failed to cleanup block file %s: %v", path, err2)
			}
			return nil, fmt.Errorf("failed to attach device %v: %v", path, err)
		}
	default:
		return nil, fmt.Errorf("unsupported access type %v", volAccessType)
	}

	hostpathVol := hostPathVolume{
		VolID:         volID,
		VolName:       name,
		VolSize:       cap,
		VolPath:       path,
		VolAccessType: volAccessType,
		Ephemeral:     ephemeral,
	}
	hostPathVolumes[volID] = hostpathVol
	return &hostpathVol, nil
}

// updateVolume updates the existing hostpath volume.
func updateHostpathVolume(volID string, volume hostPathVolume) error {
	glog.V(4).Infof("updating hostpath volume: %s", volID)

	if _, err := getVolumeByID(volID); err != nil {
		return err
	}

	hostPathVolumes[volID] = volume
	return nil
}

// deleteVolume deletes the directory for the hostpath volume.
func deleteHostpathVolume(volID string) error {
	glog.V(4).Infof("deleting hostpath volume: %s", volID)

	vol, err := getVolumeByID(volID)
	if err != nil {
		// Return OK if the volume is not found.
		return nil
	}

	if vol.VolAccessType == blockAccess {
		volPathHandler := volumepathhandler.VolumePathHandler{}
		// Get the associated loop device.
		device, err := volPathHandler.GetLoopDevice(getVolumePath(volID))
		if err != nil {
			return fmt.Errorf("failed to get the loop device: %v", err)
		}

		if device != "" {
			// Remove any associated loop device.
			glog.V(4).Infof("deleting loop device %s", device)
			if err := volPathHandler.RemoveMapPath(device); err != nil {
				return fmt.Errorf("failed to remove loop device %v: %v", device, err)
			}
		}
	}

	path := getVolumePath(volID)
	if err := os.RemoveAll(path); err != nil && !os.IsNotExist(err) {
		return err
	}
	delete(hostPathVolumes, volID)
	return nil
}

// hostPathIsEmpty is a simple check to determine if the specified hostpath directory
// is empty or not.
func hostPathIsEmpty(p string) (bool, error) {
	f, err := os.Open(p)
	if err != nil {
		return true, fmt.Errorf("unable to open hostpath volume, error: %v", err)
	}
	defer f.Close()

	_, err = f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err
}

// loadFromSnapshot populates the given destPath with data from the snapshotID
func loadFromSnapshot(size int64, snapshotId, destPath string, mode accessType) error {
	snapshot, ok := hostPathVolumeSnapshots[snapshotId]
	if !ok {
		return status.Errorf(codes.NotFound, "cannot find snapshot %v", snapshotId)
	}
	if snapshot.ReadyToUse != true {
		return status.Errorf(codes.Internal, "snapshot %v is not yet ready to use.", snapshotId)
	}
	if snapshot.SizeBytes > size {
		return status.Errorf(codes.InvalidArgument, "snapshot %v size %v is greater than requested volume size %v", snapshotId, snapshot.SizeBytes, size)
	}
	snapshotPath := snapshot.Path

	var cmd []string
	switch mode {
	case mountAccess:
		cmd = []string{"tar", "zxvf", snapshotPath, "-C", destPath}
	case blockAccess:
		cmd = []string{"dd", "if=" + snapshotPath, "of=" + destPath}
	default:
		return status.Errorf(codes.InvalidArgument, "unknown accessType: %d", mode)
	}
	executor := utilexec.New()
	out, err := executor.Command(cmd[0], cmd[1:]...).CombinedOutput()
	if err != nil {
		return status.Errorf(codes.Internal, "failed pre-populate data from snapshot %v: %v: %s", snapshotId, err, out)
	}
	return nil
}

// loadFromVolume populates the given destPath with data from the srcVolumeID
func loadFromVolume(size int64, srcVolumeId, destPath string, mode accessType) error {
	hostPathVolume, ok := hostPathVolumes[srcVolumeId]
	if !ok {
		return status.Error(codes.NotFound, "source volumeId does not exist, are source/destination in the same storage class?")
	}
	if hostPathVolume.VolSize > size {
		return status.Errorf(codes.InvalidArgument, "volume %v size %v is greater than requested volume size %v", srcVolumeId, hostPathVolume.VolSize, size)
	}
	if mode != hostPathVolume.VolAccessType {
		return status.Errorf(codes.InvalidArgument, "volume %v mode is not compatible with requested mode", srcVolumeId)
	}

	switch mode {
	case mountAccess:
		return loadFromFilesystemVolume(hostPathVolume, destPath)
	case blockAccess:
		return loadFromBlockVolume(hostPathVolume, destPath)
	default:
		return status.Errorf(codes.InvalidArgument, "unknown accessType: %d", mode)
	}
}

func loadFromFilesystemVolume(hostPathVolume hostPathVolume, destPath string) error {
	srcPath := hostPathVolume.VolPath
	isEmpty, err := hostPathIsEmpty(srcPath)
	if err != nil {
		return status.Errorf(codes.Internal, "failed verification check of source hostpath volume %v: %v", hostPathVolume.VolID, err)
	}

	// If the source hostpath volume is empty it's a noop and we just move along, otherwise the cp call will fail with a a file stat error DNE
	if !isEmpty {
		args := []string{"-a", srcPath + "/.", destPath + "/"}
		executor := utilexec.New()
		out, err := executor.Command("cp", args...).CombinedOutput()
		if err != nil {
			return status.Errorf(codes.Internal, "failed pre-populate data from volume %v: %v: %s", hostPathVolume.VolID, err, out)
		}
	}
	return nil
}

func loadFromBlockVolume(hostPathVolume hostPathVolume, destPath string) error {
	srcPath := hostPathVolume.VolPath
	args := []string{"if=" + srcPath, "of=" + destPath}
	executor := utilexec.New()
	out, err := executor.Command("dd", args...).CombinedOutput()
	if err != nil {
		return status.Errorf(codes.Internal, "failed pre-populate data from volume %v: %v: %s", hostPathVolume.VolID, err, out)
	}
	return nil
}
