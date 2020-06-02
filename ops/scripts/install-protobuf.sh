#!/bin/sh
set -ex

# TODO: refactoring this script
PROTOBUF_VERSION=3.12.2
basename=protoc-$PROTOBUF_VERSION-linux-x86_64.zip

wget https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOBUF_VERSION/$basename
unzip -o $basename
mv bin/protoc /usr/bin/protoc
rm $basename
rm -rf include/google
rm -rf readme.txt
