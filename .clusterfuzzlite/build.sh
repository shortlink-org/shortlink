#!/bin/bash -eu

go mod init "github.com/shortlink-org/shortlink"
export FUZZ_ROOT="github.com/shortlink-org/shortlink"

compile_go_fuzzer github.com/shortlink-org/shortlink/internal/pkg/batch FuzzBatch fuzz_batch
