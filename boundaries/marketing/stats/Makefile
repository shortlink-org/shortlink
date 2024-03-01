SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

# INCLUDE ==============================================================================================================
# Include Makefile
include ../../../ops/Makefile/common.mk

# BUILD ================================================================================================================
build: ### Build the project using CMake
	@mkdir build && cd build
	@conan install ..
	@cmake .. -DCMAKE_EXPORT_COMPILE_COMMANDS=1 # generates compile_commands.json
	@ln -s compile_commands.json ../compile_commands.json # link compile_commands.json to home directory

# BAZEL ================================================================================================================
bazel_build: ### Build the project using Bazel
	@conan install --install-folder=deps --build=missing .
	@bazel build //...
