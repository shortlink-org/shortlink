SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

# INCLUDE ==============================================================================================================
# Include Makefile
include ../../../../ops/Makefile/common.mk

# INTEGRATION TESTS ====================================================================================================
k6:
	@k6 run ./tests/integrations/k6-link-crud.js
