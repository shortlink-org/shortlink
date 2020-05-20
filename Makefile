SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

# PROJECT_NAME defaults to name of the current directory.
# should not to be changed if you follow GitOps operating procedures.
PROJECT_NAME := shortlink

CI_COMMIT_TAG := latest

DOCKER_USERNAME := "batazor"

# Export such that its passed to shell functions for Docker to pick up.
export PROJECT_NAME

# Include Makefile
include $(SELF_DIR)/ops/Makefile/common.mk
include $(SELF_DIR)/ops/Makefile/vagrant.mk
include $(SELF_DIR)/ops/Makefile/ansible.mk
include $(SELF_DIR)/ops/Makefile/docker.mk
include $(SELF_DIR)/ops/Makefile/gitlab.mk
include $(SELF_DIR)/ops/Makefile/go.mk
include $(SELF_DIR)/ops/Makefile/istio.mk
include $(SELF_DIR)/ops/Makefile/k8s.mk
include $(SELF_DIR)/ops/Makefile/minikube.mk
include $(SELF_DIR)/ops/Makefile/ui.mk

# HELP =================================================================================================================
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
