package migrations

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var __000001_create_links_collection_down_json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _000001_create_links_collection_down_json() ([]byte, error) {
	return bindata_read(
		__000001_create_links_collection_down_json,
		"000001_create_links_collection.down.json",
	)
}

var __000001_create_links_collection_up_json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x53\x4d\x6f\xf2\x30\x0c\xbe\xe7\x57\x58\xd6\x7b\xe4\xc2\x95\x7f\xf0\x4a\xdb\x65\x1f\x27\x84\x50\xda\x58\xc3\xd0\x26\x21\x71\xa6\xa1\x89\xff\x3e\x35\xed\x68\x8a\xd8\x07\x1a\xa7\x54\xb6\x9f\x0f\x57\x7e\x96\x0a\xe0\x5d\x01\x00\x60\x1d\x48\x0b\xe1\x02\xb0\x61\xbb\x8b\xa8\x00\x8e\xb3\xa2\xed\x9a\xe6\xde\x99\xb1\x3f\xeb\xeb\xaf\xba\x61\xa3\xc5\x05\x5c\x0c\xa3\x00\xf8\x6f\x1b\x9d\x7d\xac\x37\xd4\xea\xa2\x0c\x80\x55\x74\xf6\xe9\xe0\xb3\x8c\xab\xb6\x54\xcb\xc0\x93\xbb\x81\xf6\x89\x03\x75\x22\x4b\x4c\xa1\xc1\x19\xe0\x46\xc7\x4d\xf7\xea\x20\x1c\x65\xcd\x06\x57\x05\xc2\x07\xe7\x29\x08\x53\x9c\xe8\x00\x64\xf8\xb4\x74\x26\x1f\x25\xb0\x7d\x29\xe4\xf3\x84\xa1\x58\x07\xf6\xc2\xce\x76\x43\xcf\x0f\x77\xd0\xa6\x28\x50\x11\xc8\xc1\x13\x0c\xa8\x02\x74\x2c\x19\x7a\xbb\x7f\xd7\xed\x68\xae\x13\xee\x09\x2a\xba\x81\xf8\x27\xd5\x75\x06\xfa\xfb\x31\x6b\x2d\xdf\x5b\x10\x6e\x29\x8a\x6e\xfd\x0f\x2e\x46\xbe\xa9\x8f\x11\xff\xa5\x95\xe4\xcd\x4d\xad\x8c\x7c\xbf\xb1\xa2\xce\xbf\xfa\xf7\x78\x1e\xa8\xbc\xdf\x7f\x6b\xe8\x2d\x5f\xef\x34\x56\x7c\xaa\x2f\x07\x96\x22\x45\x3b\x3a\x5c\x3e\xf7\xf9\xa5\x5b\x9c\xab\x0b\x7f\x09\xad\x6e\x4f\x61\x5f\x1b\x12\xcd\x4d\x2c\xb3\x98\x2c\xef\x53\x37\x21\x21\xd1\x64\x91\x55\xb7\x88\x5a\xa9\x8f\x00\x00\x00\xff\xff\x4a\x1b\xf6\xe8\x3c\x04\x00\x00")

func _000001_create_links_collection_up_json() ([]byte, error) {
	return bindata_read(
		__000001_create_links_collection_up_json,
		"000001_create_links_collection.up.json",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"000001_create_links_collection.down.json": _000001_create_links_collection_down_json,
	"000001_create_links_collection.up.json":   _000001_create_links_collection_up_json,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() ([]byte, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"000001_create_links_collection.down.json": &_bintree_t{_000001_create_links_collection_down_json, map[string]*_bintree_t{}},
	"000001_create_links_collection.up.json":   &_bintree_t{_000001_create_links_collection_up_json, map[string]*_bintree_t{}},
}}
