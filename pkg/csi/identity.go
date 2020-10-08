package csi_driver

import (
	"context"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/golang/protobuf/ptypes/wrappers"

	"github.com/batazor/shortlink/internal/logger/field"
)

// GetPluginInfo returns metadata of the plugin
func (d *Driver) GetPluginInfo(ctx context.Context, req *csi.GetPluginInfoRequest) (*csi.GetPluginInfoResponse, error) {
	resp := &csi.GetPluginInfoResponse{
		Name:          d.name,
		VendorVersion: "1.0.0",
	}

	d.log.InfoWithContext(ctx, "get plugin info called", field.Fields{
		"response": resp,
		"method":   "get_plugin_info",
	})

	return resp, nil
}

// GetPluginCapabilities returns available capabilities of the plugin
func (d *Driver) GetPluginCapabilities(ctx context.Context, req *csi.GetPluginCapabilitiesRequest) (*csi.GetPluginCapabilitiesResponse, error) {
	resp := &csi.GetPluginCapabilitiesResponse{
		Capabilities: []*csi.PluginCapability{
			{
				Type: &csi.PluginCapability_Service_{
					Service: &csi.PluginCapability_Service{
						Type: csi.PluginCapability_Service_CONTROLLER_SERVICE,
					},
				},
			},
			{
				Type: &csi.PluginCapability_Service_{
					Service: &csi.PluginCapability_Service{
						Type: csi.PluginCapability_Service_VOLUME_ACCESSIBILITY_CONSTRAINTS,
					},
				},
			},
			{
				Type: &csi.PluginCapability_VolumeExpansion_{
					VolumeExpansion: &csi.PluginCapability_VolumeExpansion{
						Type: csi.PluginCapability_VolumeExpansion_ONLINE,
					},
				},
			},
		},
	}

	d.log.InfoWithContext(ctx, "get plugin capabitilies called", field.Fields{
		"response": resp,
		"method":   "get_plugin_capabilities",
	})

	return resp, nil
}

// Probe returns the health and readiness of the plugin
func (d *Driver) Probe(ctx context.Context, req *csi.ProbeRequest) (*csi.ProbeResponse, error) {
	d.log.InfoWithContext(ctx, "probe called", field.Fields{
		"method": "probe",
	})
	d.readyMu.Lock()
	defer d.readyMu.Unlock()

	return &csi.ProbeResponse{
		Ready: &wrappers.BoolValue{
			Value: d.ready,
		},
	}, nil
}
