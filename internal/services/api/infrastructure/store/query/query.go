//go:generate protoc -I../../../../../ -I. --gotemplate_out=all=true,template_dir=template:. services/api/domain/link/link.proto

package query
