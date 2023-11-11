#!/bin/bash -eu

# install go-fuzz-build
go install github.com/AdamKorcz/go-118-fuzz-build@latest
go get github.com/AdamKorcz/go-118-fuzz-build/testing

compile_native_go_fuzzer github.com/shortlink-org/shortlink/internal/pkg/batch -func FuzzBatch -tags fuzz fuzz_batch
