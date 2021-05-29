SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

# BASE CONFIG ==========================================================================================================
.SILENT: ;               # no need for @
.ONESHELL: ;             # recipes execute in same shell
.NOTPARALLEL: ;          # wait for this target to finish
.EXPORT_ALL_VARIABLES: ; # send all vars to shell
default: help;           # default target
Makefile: ;              # skip prerequisite discovery

# HELP =================================================================================================================
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

# INCLUDE ==============================================================================================================
# Include Makefile
include $(SELF_DIR)/ops/Makefile/common.mk
include $(SELF_DIR)/ops/Makefile/cert.mk
include $(SELF_DIR)/ops/Makefile/proto.mk
include $(SELF_DIR)/ops/Makefile/cli.mk
include $(SELF_DIR)/ops/Makefile/vagrant.mk
include $(SELF_DIR)/ops/Makefile/ansible.mk
include $(SELF_DIR)/ops/Makefile/terraform.mk
include $(SELF_DIR)/ops/Makefile/docker.mk
include $(SELF_DIR)/ops/Makefile/go.mk
include $(SELF_DIR)/ops/Makefile/k8s/k8s.mk
include $(SELF_DIR)/ops/Makefile/ui.mk
