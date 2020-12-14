//go:generate protoc -I. --go_out=Minternal/metadata/domain/rpc.proto=.:. --go_opt=paths=source_relative rpc.proto

package rpc
