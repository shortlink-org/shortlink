SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
ROOT_DIR := $(SELF_DIR)

# INCLUDE ==============================================================================================================
# Include Makefile
include $(SELF_DIR)/ops/Makefile/common.mk
include $(SELF_DIR)/ops/Makefile/dev.mk
include $(SELF_DIR)/ops/Makefile/proto.mk
include $(SELF_DIR)/ops/Makefile/tests.mk
include $(SELF_DIR)/ops/Makefile/docs.mk
