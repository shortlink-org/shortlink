SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
ROOT_DIR := $(abspath $(SELF_DIR)/../../..)

# INCLUDE ==============================================================================================================
# Include Makefile
include $(ROOT_DIR)/ops/Makefile/common.mk
include $(SELF_DIR)/ops/Makefile/dev.mk
include $(SELF_DIR)/ops/Makefile/proto.mk
