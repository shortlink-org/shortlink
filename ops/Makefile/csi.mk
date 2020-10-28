# CSI TASKS ============================================================================================================

# Change to the latest supported snapshotter version
SNAPSHOTTER_VERSION=v3.0.1

csi-build: ## Build the CSI container
		@echo docker build image ${CI_REGISTRY_IMAGE}-csi:${CI_COMMIT_TAG}
		@docker build -t ${CI_REGISTRY_IMAGE}-csi:${CI_COMMIT_TAG} -f ops/dockerfile/csi.Dockerfile .

csi-up: ## Deploy CSI plugin
		# Apply VolumeSnapshot CRDs
		@kubectl apply -f "https://raw.githubusercontent.com/kubernetes-csi/external-snapshotter/${SNAPSHOTTER_VERSION}/client/config/crd/snapshot.storage.k8s.io_volumesnapshotclasses.yaml"
		@kubectl apply -f "https://raw.githubusercontent.com/kubernetes-csi/external-snapshotter/${SNAPSHOTTER_VERSION}/client/config/crd/snapshot.storage.k8s.io_volumesnapshotcontents.yaml"
		@kubectl apply -f "https://raw.githubusercontent.com/kubernetes-csi/external-snapshotter/${SNAPSHOTTER_VERSION}/client/config/crd/snapshot.storage.k8s.io_volumesnapshots.yaml"

		# Create snapshot controller
		@kubectl apply -f "https://raw.githubusercontent.com/kubernetes-csi/external-snapshotter/${SNAPSHOTTER_VERSION}/deploy/kubernetes/snapshot-controller/rbac-snapshot-controller.yaml"
		@kubectl apply -f "https://raw.githubusercontent.com/kubernetes-csi/external-snapshotter/${SNAPSHOTTER_VERSION}/deploy/kubernetes/snapshot-controller/setup-snapshot-controller.yaml"

		# applying RBAC rules
		@kubectl apply -f https://raw.githubusercontent.com/kubernetes-csi/external-provisioner/v2.0.3/deploy/kubernetes/rbac.yaml
		@kubectl apply -f https://raw.githubusercontent.com/kubernetes-csi/external-attacher/v3.0.1/deploy/kubernetes/rbac.yaml
		@kubectl apply -f https://raw.githubusercontent.com/kubernetes-csi/external-snapshotter/v3.0.1/deploy/kubernetes/csi-snapshotter/rbac-csi-snapshotter.yaml
		@kubectl apply -f https://raw.githubusercontent.com/kubernetes-csi/external-resizer/v1.0.1/deploy/kubernetes/rbac.yaml

		@kubectl apply -f /home/batazor/myproejct/shortlink/ops/Helm/csi/templates/
