# Include Makefile
include $(SELF_DIR)/ops/Makefile/k8s/helm.mk
include $(SELF_DIR)/ops/Makefile/k8s/ct.mk
include $(SELF_DIR)/ops/Makefile/k8s/k8s.shortlink.mk
include $(SELF_DIR)/ops/Makefile/k8s/k8s.velero.mk
include $(SELF_DIR)/ops/Makefile/k8s/kubeadm.mk
include $(SELF_DIR)/ops/Makefile/k8s/minikube.mk
include $(SELF_DIR)/ops/Makefile/k8s/csi.mk
include $(SELF_DIR)/ops/Makefile/k8s/istio.mk
include $(SELF_DIR)/ops/Makefile/k8s/gitlab.mk
include $(SELF_DIR)/ops/Makefile/k8s/prometheus.mk
include $(SELF_DIR)/ops/Makefile/k8s/devspace.mk
