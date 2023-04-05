#!/bin/sh
set -ex

export PROTOBUF_VERSION=22.2
export basename=protoc-$PROTOBUF_VERSION-linux-x86_64.zip

case "$(uname -sr)" in
    Linux*)
         echo 'Install protobuff for Linux'
         install_protobuf_for_linux
         ;;
    Darwin*)
         echo 'Install protobuff for Mac OS X'
         install_protobuf_for_mac
         ;;
esac

install_protobuf_for_linux() {
  wget https://github.com/protocolbuffers/protobuf/releases/download/v$PROTOBUF_VERSION/$basename
  unzip -o $basename
  rm $basename
  mv -u bin/protoc /home/$USER/.local/bin/protoc
  rm -rf include readme.txt bin
}

install_protobuf_for_mac() {
  brew install protobuf
}
