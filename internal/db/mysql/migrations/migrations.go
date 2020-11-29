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

var __000001_init_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _000001_init_down_sql() ([]byte, error) {
	return bindata_read(
		__000001_init_down_sql,
		"000001_init.down.sql",
	)
}

var __000001_init_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x90\x41\x4b\xf3\x40\x10\x86\xef\xfd\x15\xef\x31\x0b\x3d\xf4\xfb\xb0\xa2\x48\x0f\xdb\x66\xaa\x8b\xe9\xa6\x6c\x66\xc1\x9e\x96\x6d\x12\xe9\x62\x9b\x96\x64\x2b\xfa\xef\xc5\x6a\xd3\x93\x20\xce\x69\x06\x9e\x79\x78\x79\x67\x86\x24\x13\x58\x4e\x33\x82\x9a\x43\xe7\x0c\x7a\x52\x05\x17\xd8\x86\xe6\xa5\x43\x32\x00\x80\x50\xa1\x9f\x75\x68\x7c\xfb\x9e\xfc\xbb\x16\xc0\x89\xd7\x36\xcb\x90\xd2\x5c\xda\x8c\x91\x58\xab\x52\xc7\xb9\x9b\x2a\x7d\xda\x13\x31\x04\x1b\x4b\x42\x0c\x4f\xaa\x63\xbb\xed\x55\xaf\xbe\x2d\x37\xbe\x4d\xfe\x8f\xc7\xa2\x57\x7d\x61\x1b\xdf\x6d\x7e\x81\x55\x75\x57\xb6\xe1\x10\xc3\xbe\x41\xac\xdf\x62\xaf\xbe\x20\x65\x5b\xfb\x58\x57\xce\x47\x20\x86\x5d\xdd\x45\xbf\x3b\x9c\x91\x3e\xf7\xcc\x1a\x43\x9a\x1d\xab\x05\x15\x2c\x17\xcb\xef\xb4\x87\xea\x2f\xcf\xc8\x35\xec\x32\xfd\xac\xf6\x07\xf1\xd2\xa8\x85\x34\x2b\x3c\xd2\x0a\x49\xa8\xc4\x40\x80\xf4\xbd\xd2\x34\x51\x4d\xb3\x4f\xa7\x17\xf7\x83\x34\x05\xf1\xe4\x18\x9f\x6f\x76\xeb\x2b\xcc\xf2\x2c\x93\x4c\xe7\xdb\x8d\x6e\x47\x23\xe7\x83\x2b\xc3\xdd\xe0\x23\x00\x00\xff\xff\x63\xf0\x8e\x72\xce\x01\x00\x00")

func _000001_init_up_sql() ([]byte, error) {
	return bindata_read(
		__000001_init_up_sql,
		"000001_init.up.sql",
	)
}

var __000002_add_test_link_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _000002_add_test_link_down_sql() ([]byte, error) {
	return bindata_read(
		__000002_add_test_link_down_sql,
		"000002_add_test_link.down.sql",
	)
}

var __000002_add_test_link_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\xcf\x3f\x4f\xf3\x30\x10\xc7\xf1\xdd\xaf\xe2\xb7\xa5\x95\x92\x56\xcf\xc3\x46\xa6\x34\x8a\x20\x22\x7f\xa4\xc4\xb0\x46\x4e\xea\xd6\x56\xd3\x38\xb2\xcf\x40\x79\xf5\x88\x96\x8a\x85\x0d\xc4\x4d\x37\x7c\xa5\xcf\x5d\x14\x81\x37\x49\xd5\x26\x29\xcf\xeb\x0a\xd1\x5f\x0d\xdb\x64\x77\x79\x15\x33\x96\x57\x6d\xd6\x70\xe4\x15\xaf\x31\xea\xe9\xe0\x16\xde\x8e\x21\x94\x70\x2a\xc4\x56\xba\xc1\xea\x99\xb4\x99\x96\x0c\x00\x9e\x92\xe2\x31\x6b\xb1\x08\x14\xd1\xec\x6e\xd7\xeb\x5e\x90\x78\x33\x76\x65\x7d\x10\x22\x38\x9e\xee\x85\x53\xff\x3e\xd6\xf2\x84\x59\x5a\x67\x26\x31\xe2\x45\xf6\x4e\x93\x0c\x96\x3f\xf7\xf6\x9a\x94\xef\x57\x83\x39\x5e\xe9\x2f\xf7\xff\xa7\x2b\x86\xc1\x78\x82\xd9\xe1\x52\xff\x86\xfb\x7c\xf8\xde\xbc\xb9\xfe\x2a\xf6\x12\x66\xc2\xa5\x3b\x8b\x51\x84\xa6\x2e\x8a\x4d\x92\x3e\x80\xd7\xa0\xd7\x6e\xb0\x52\x90\xec\xb6\x72\x27\xfc\x48\xdd\xf9\x8c\x98\xb1\xb4\x2e\xcb\x9c\xc7\xec\x3d\x00\x00\xff\xff\x1a\x33\x82\xbf\x0c\x02\x00\x00")

func _000002_add_test_link_up_sql() ([]byte, error) {
	return bindata_read(
		__000002_add_test_link_up_sql,
		"000002_add_test_link.up.sql",
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
	"000001_init.down.sql":          _000001_init_down_sql,
	"000001_init.up.sql":            _000001_init_up_sql,
	"000002_add_test_link.down.sql": _000002_add_test_link_down_sql,
	"000002_add_test_link.up.sql":   _000002_add_test_link_up_sql,
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
	"000001_init.down.sql":          &_bintree_t{_000001_init_down_sql, map[string]*_bintree_t{}},
	"000001_init.up.sql":            &_bintree_t{_000001_init_up_sql, map[string]*_bintree_t{}},
	"000002_add_test_link.down.sql": &_bintree_t{_000002_add_test_link_down_sql, map[string]*_bintree_t{}},
	"000002_add_test_link.up.sql":   &_bintree_t{_000002_add_test_link_up_sql, map[string]*_bintree_t{}},
}}
