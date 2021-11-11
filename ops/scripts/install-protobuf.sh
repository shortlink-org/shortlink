#!/bin/sh
set -ex

# TODO: refactoring this script
PROTOBUF_VERSION=3.19.1
basename=protoc-$PROTOBUF_VERSION-linux-x86_64.zip

wget https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOBUF_VERSION/$basename
unzip -o $basename
rm $basename
mv bin/protoc /home/$USER/.local/bin/protoc
rm -rf include readme.txt bin
