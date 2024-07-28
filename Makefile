SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

# INCLUDE ==============================================================================================================
# Include Makefile
include $(SELF_DIR)/ops/Makefile/common.mk
include $(SELF_DIR)/ops/Makefile/docs.mk
include $(SELF_DIR)/ops/Makefile/cert.mk
include $(SELF_DIR)/ops/Makefile/proto.mk
include $(SELF_DIR)/ops/Makefile/cli.mk
include $(SELF_DIR)/ops/Makefile/vagrant.mk
include $(SELF_DIR)/ops/Makefile/ansible.mk
include $(SELF_DIR)/ops/Makefile/terraform.mk
include $(SELF_DIR)/ops/Makefile/dev.mk
include $(SELF_DIR)/ops/Makefile/go.mk
include $(SELF_DIR)/ops/Makefile/git.mk
include $(SELF_DIR)/ops/Makefile/k8s/k8s.mk
