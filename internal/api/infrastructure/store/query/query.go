//go:generate protoc -I../../../../../ -I. --gotemplate_out=all=true,template_dir=template:. internal/api/domain/link/link.proto

package query
