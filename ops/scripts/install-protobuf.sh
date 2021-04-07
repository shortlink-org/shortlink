#!/bin/sh
set -ex

# TODO: refactoring this script
PROTOBUF_VERSION=3.15.7
basename=protoc-$PROTOBUF_VERSION-linux-x86_64.zip

wget https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOBUF_VERSION/$basename
unzip -o $basename
rm $basename
rm -rf include
rm -rf readme.txt
mv bin/protoc /home/$USER/.local/bin/protoc
