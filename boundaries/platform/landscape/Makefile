SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

# INCLUDE ==============================================================================================================
# Include Makefile
include ../../../ops/Makefile/common.mk

# LANDSCAPE ============================================================================================================
build: ##@landscape Build landscape
	@landscape2 build \
      --data-file data.yml \
      --settings-file settings.yml \
      --guide-file guide.yml \
      --logos-path logos \
      --output-dir build
