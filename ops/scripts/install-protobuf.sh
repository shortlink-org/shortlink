#!/bin/sh
set -ex

PROTOBUF_VERSION=3.10.0
basename=protobuf-$PROTOBUF_VERSION

wget https://github.com/google/protobuf/releases/download/v$PROTOBUF_VERSION/$basename.tar.gz
tar -xzvf $basename.tar.gz
cd protobuf-$PROTOBUF_VERSION && ./configure --prefix=/usr && make -j2 && sudo make install
