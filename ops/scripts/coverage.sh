#!/bin/sh
set -ex

echo "" > coverage.txt

for d in $(go list ./... | grep -v vendor | grep -v node_modules); do
    go test -race -coverprofile=profile.out -covermode=atomic "${d}"
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done
