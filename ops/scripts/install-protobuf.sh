#!/bin/sh
set -ex

PROTOBUF_VERSION=3.10.0
basename=protoc-$PROTOBUF_VERSION-linux-x86_64.zip

wget https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOBUF_VERSION/$basename
unzip $basename
mv bin/protoc /usr/bin/protoc
