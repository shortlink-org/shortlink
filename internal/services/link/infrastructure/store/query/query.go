//go:generate protoc -I../../../../../ -I. --gotemplate_out=all=true,template_dir=template:. services/link/domain/link/v1/link.proto

package query
