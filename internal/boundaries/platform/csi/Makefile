SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

# INCLUDE ==============================================================================================================
# Include Makefile
include $(SELF_DIR)/../../../ops/Makefile/common.mk

# CSI TASKS ============================================================================================================

# Change to the latest supported snapshotter version
SNAPSHOTTER_VERSION=v6.2.1

build: ## Build the CSI container
		@echo docker buildx build \
			--platform=linux/amd64,linux/arm64 \
			--provenance=true \
			--sbom=true \
			image ${CI_REGISTRY_IMAGE}-csi:${CI_COMMIT_TAG}
		@docker buildx build \
			--platform=linux/amd64,linux/arm64 \
			--provenance=true \
			--sbom=true \
			--no-cache \
			-t ${CI_REGISTRY_IMAGE}-csi \
			-f ops/dockerfile/csi.Dockerfile .

up: ## Deploy CSI plugin
		@kubectl apply -n default --prune --applyset=external-snapshotter \
			# Apply VolumeSnapshot CRDs
			-f "https://raw.githubusercontent.com/kubernetes-csi/external-snapshotter/${SNAPSHOTTER_VERSION}/client/config/crd/snapshot.storage.k8s.io_volumesnapshotclasses.yaml" \
			-f "https://raw.githubusercontent.com/kubernetes-csi/external-snapshotter/${SNAPSHOTTER_VERSION}/client/config/crd/snapshot.storage.k8s.io_volumesnapshotcontents.yaml" \
			-f "https://raw.githubusercontent.com/kubernetes-csi/external-snapshotter/${SNAPSHOTTER_VERSION}/client/config/crd/snapshot.storage.k8s.io_volumesnapshots.yaml" \

			# Create snapshot controller
			-f "https://raw.githubusercontent.com/kubernetes-csi/external-snapshotter/${SNAPSHOTTER_VERSION}/deploy/kubernetes/snapshot-controller/rbac-snapshot-controller.yaml" \
			-f "https://raw.githubusercontent.com/kubernetes-csi/external-snapshotter/${SNAPSHOTTER_VERSION}/deploy/kubernetes/snapshot-controller/setup-snapshot-controller.yaml" \

			# applying RBAC rules
			-f https://raw.githubusercontent.com/kubernetes-csi/external-provisioner/v2.0.3/deploy/kubernetes/rbac.yaml \
			-f https://raw.githubusercontent.com/kubernetes-csi/external-attacher/v3.0.1/deploy/kubernetes/rbac.yaml \
			-f https://raw.githubusercontent.com/kubernetes-csi/external-snapshotter/v3.0.1/deploy/kubernetes/csi-snapshotter/rbac-csi-snapshotter.yaml \
			-f https://raw.githubusercontent.com/kubernetes-csi/external-resizer/v1.0.1/deploy/kubernetes/rbac.yaml \

			-f /home/batazor/myproejct/shortlink/ops/Helm/csi/templates/

example: ## Deploy examples for CSI
	@kubectl apply -n default --prune --applyset=csi-example \
		-f ops/Helm/csi/examples/
