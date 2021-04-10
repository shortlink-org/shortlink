//go:generate protoc -I. -I../../..  -I../../../third_party/googleapis --proto_path=src --go_out=Mpkg/api/http-chi/api.proto=.:. --go_opt=paths=source_relative http-api.proto
//go:generate protoc -I. -I../../..  -I../../../third_party/googleapis --jsonschema_out=./jsonschema http-api.proto

package http_chi
