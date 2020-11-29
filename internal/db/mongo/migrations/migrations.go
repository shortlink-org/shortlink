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
		return nil, fmt.Errorf("Read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %w", name, err)
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

var __000001_create_links_collection_up_json = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x52\x4d\x6b\x84\x30\x10\xbd\xe7\x57\x0c\x43\x8f\x5e\xf6\xba\xff\xa0\xd0\x5e\xfa\x71\x12\x29\x51\x87\x9a\xdd\x98\xb8\xc9\xa4\x54\x8a\xff\xbd\x24\xda\x6d\x14\xa1\x94\xf6\xa4\xbc\x79\xf3\x5e\x66\xe6\x95\x02\xe0\x43\x00\x00\x60\xe3\x48\x32\xe1\x11\x50\x2b\x73\xf6\x28\x00\xa6\x22\x2b\x5b\xad\xef\x6d\xfb\x5d\x2f\x66\xfc\x4d\x6a\xd5\x4a\xb6\x0e\x8f\x0b\x15\x00\x6f\x4e\xde\x9a\xc7\xa6\xa3\x5e\x66\x30\x00\xd6\xde\x9a\xa7\x71\x48\x36\xb6\x3e\x51\xc3\x8b\x4e\xaa\x3a\xba\x04\xe5\x28\x9a\x94\x18\x9c\xc6\x02\xb0\x93\xbe\xc3\x2a\x23\x0d\xce\x0e\xe4\x58\x91\x5f\x49\x03\xa4\x8e\x35\xb4\x71\xf4\xec\x94\x79\xcd\x1c\x13\xa3\x25\xdf\x38\x35\xb0\xb2\x26\x92\x9e\x1f\xee\xa0\x0f\x9e\xa1\x26\xe0\x71\x20\x58\xba\xb2\xa6\x29\x57\x98\x5f\xf8\x77\xdf\x28\xf3\x3b\xe3\x59\xa0\xa6\x7f\x30\xff\x92\xfa\xf1\x01\x62\xfb\x37\x7f\xa7\x6d\x5a\x52\x98\x6e\x4d\x4b\xef\xe9\x4e\xeb\xcc\xa8\x2b\x5e\x2e\x2a\x59\x44\xce\x34\xee\x1f\xf6\xb0\xb7\xf5\x83\xd8\xd9\x0d\x1a\xd9\x5f\x93\xfc\xd2\x12\x4b\xa5\x7d\x1e\xb4\x60\xd4\x25\x44\x06\xbb\x40\xab\x41\xaa\x38\x88\xa8\xc4\x67\x00\x00\x00\xff\xff\x12\xa4\x6c\x29\x19\x03\x00\x00")

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
	"000001_create_links_collection.down.json": {_000001_create_links_collection_down_json, map[string]*_bintree_t{}},
	"000001_create_links_collection.up.json":   {_000001_create_links_collection_up_json, map[string]*_bintree_t{}},
}}
