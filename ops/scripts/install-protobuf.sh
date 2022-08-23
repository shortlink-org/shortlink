#!/bin/sh
set -ex

export PROTOBUF_VERSION=21.5
export basename=protoc-$PROTOBUF_VERSION-linux-x86_64.zip

wget https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOBUF_VERSION/$basename
unzip -o $basename
rm $basename
mv -u bin/protoc /home/$USER/.local/bin/protoc
rm -rf include readme.txt bin
