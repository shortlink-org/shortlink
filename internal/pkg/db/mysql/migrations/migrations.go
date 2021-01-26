// Code generated for package migrations by go-bindata DO NOT EDIT. (@generated)
// sources:
// migrations/000001_init.down.sql
// migrations/000001_init.up.sql
// migrations/000002_add_test_link.down.sql
// migrations/000002_add_test_link.up.sql
// migrations/migrations_test.go
package migrations

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var __000001_initDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _000001_initDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000001_initDownSql,
		"000001_init.down.sql",
	)
}

func _000001_initDownSql() (*asset, error) {
	bytes, err := _000001_initDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000001_init.down.sql", size: 0, mode: os.FileMode(436), modTime: time.Unix(1601856486, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000001_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x90\x41\x4b\xf3\x40\x10\x86\xef\xfd\x15\xef\x31\x0b\x3d\xf4\xfb\xb0\xa2\x48\x0f\xdb\x66\xaa\x8b\xe9\xa6\x6c\x66\xc1\x9e\x96\x6d\x12\xe9\x62\x9b\x96\x64\x2b\xfa\xef\xc5\x6a\xd3\x93\x20\xce\x69\x06\x9e\x79\x78\x79\x67\x86\x24\x13\x58\x4e\x33\x82\x9a\x43\xe7\x0c\x7a\x52\x05\x17\xd8\x86\xe6\xa5\x43\x32\x00\x80\x50\xa1\x9f\x75\x68\x7c\xfb\x9e\xfc\xbb\x16\xc0\x89\xd7\x36\xcb\x90\xd2\x5c\xda\x8c\x91\x58\xab\x52\xc7\xb9\x9b\x2a\x7d\xda\x13\x31\x04\x1b\x4b\x42\x0c\x4f\xaa\x63\xbb\xed\x55\xaf\xbe\x2d\x37\xbe\x4d\xfe\x8f\xc7\xa2\x57\x7d\x61\x1b\xdf\x6d\x7e\x81\x55\x75\x57\xb6\xe1\x10\xc3\xbe\x41\xac\xdf\x62\xaf\xbe\x20\x65\x5b\xfb\x58\x57\xce\x47\x20\x86\x5d\xdd\x45\xbf\x3b\x9c\x91\x3e\xf7\xcc\x1a\x43\x9a\x1d\xab\x05\x15\x2c\x17\xcb\xef\xb4\x87\xea\x2f\xcf\xc8\x35\xec\x32\xfd\xac\xf6\x07\xf1\xd2\xa8\x85\x34\x2b\x3c\xd2\x0a\x49\xa8\xc4\x40\x80\xf4\xbd\xd2\x34\x51\x4d\xb3\x4f\xa7\x17\xf7\x83\x34\x05\xf1\xe4\x18\x9f\x6f\x76\xeb\x2b\xcc\xf2\x2c\x93\x4c\xe7\xdb\x8d\x6e\x47\x23\xe7\x83\x2b\xc3\xdd\xe0\x23\x00\x00\xff\xff\x63\xf0\x8e\x72\xce\x01\x00\x00")

func _000001_initUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000001_initUpSql,
		"000001_init.up.sql",
	)
}

func _000001_initUpSql() (*asset, error) {
	bytes, err := _000001_initUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000001_init.up.sql", size: 462, mode: os.FileMode(436), modTime: time.Unix(1601856496, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000002_add_test_linkDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _000002_add_test_linkDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000002_add_test_linkDownSql,
		"000002_add_test_link.down.sql",
	)
}

func _000002_add_test_linkDownSql() (*asset, error) {
	bytes, err := _000002_add_test_linkDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000002_add_test_link.down.sql", size: 0, mode: os.FileMode(436), modTime: time.Unix(1601856489, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000002_add_test_linkUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\xcf\x3f\x4f\xf3\x30\x10\xc7\xf1\xdd\xaf\xe2\xb7\xa5\x95\x92\x56\xcf\xc3\x46\xa6\x34\x8a\x20\x22\x7f\xa4\xc4\xb0\x46\x4e\xea\xd6\x56\xd3\x38\xb2\xcf\x40\x79\xf5\x88\x96\x8a\x85\x0d\xc4\x4d\x37\x7c\xa5\xcf\x5d\x14\x81\x37\x49\xd5\x26\x29\xcf\xeb\x0a\xd1\x5f\x0d\xdb\x64\x77\x79\x15\x33\x96\x57\x6d\xd6\x70\xe4\x15\xaf\x31\xea\xe9\xe0\x16\xde\x8e\x21\x94\x70\x2a\xc4\x56\xba\xc1\xea\x99\xb4\x99\x96\x0c\x00\x9e\x92\xe2\x31\x6b\xb1\x08\x14\xd1\xec\x6e\xd7\xeb\x5e\x90\x78\x33\x76\x65\x7d\x10\x22\x38\x9e\xee\x85\x53\xff\x3e\xd6\xf2\x84\x59\x5a\x67\x26\x31\xe2\x45\xf6\x4e\x93\x0c\x96\x3f\xf7\xf6\x9a\x94\xef\x57\x83\x39\x5e\xe9\x2f\xf7\xff\xa7\x2b\x86\xc1\x78\x82\xd9\xe1\x52\xff\x86\xfb\x7c\xf8\xde\xbc\xb9\xfe\x2a\xf6\x12\x66\xc2\xa5\x3b\x8b\x51\x84\xa6\x2e\x8a\x4d\x92\x3e\x80\xd7\xa0\xd7\x6e\xb0\x52\x90\xec\xb6\x72\x27\xfc\x48\xdd\xf9\x8c\x98\xb1\xb4\x2e\xcb\x9c\xc7\xec\x3d\x00\x00\xff\xff\x1a\x33\x82\xbf\x0c\x02\x00\x00")

func _000002_add_test_linkUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000002_add_test_linkUpSql,
		"000002_add_test_link.up.sql",
	)
}

func _000002_add_test_linkUpSql() (*asset, error) {
	bytes, err := _000002_add_test_linkUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000002_add_test_link.up.sql", size: 524, mode: os.FileMode(436), modTime: time.Unix(1601857839, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations_testGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x52\x51\x6b\xdb\x30\x10\x7e\xd6\xfd\x8a\x9b\xa0\x60\xaf\xc6\x6e\xfb\x38\xf0\x5e\xc6\x0a\x83\x6e\x30\x52\xd8\x43\x08\x45\x71\x64\x47\xc4\x96\x5c\xe9\xcc\x12\x52\xff\xf7\x71\xb6\x5c\x6f\xb0\x97\xc4\xfa\xee\xf4\x7d\xdf\x7d\xa7\xa2\xc0\xdb\xfd\x60\xda\x03\x0e\xd6\x10\x1e\x14\xa9\xbd\x0a\x3a\xeb\x2e\xe1\xb5\x05\xe8\x55\x75\x52\x8d\xc6\xce\x34\x5e\x91\x71\x36\x00\x98\xae\x77\x9e\x30\x01\x21\xeb\x8e\x24\x08\xe9\x02\xff\xf6\x8a\x8e\x45\x6d\x5a\xcd\x1f\x0c\x04\xf2\xc6\x36\x53\x8d\x74\x20\x63\x1b\x09\x20\x64\x63\xe8\x38\xec\xf3\xca\x75\x45\x20\xaf\xa9\x3a\xfa\x62\xaa\xd7\x97\x42\x85\xa0\x3d\x49\x48\x01\xea\xc1\x56\xf8\xac\x03\x7d\x39\xea\xea\xf4\x7d\x31\x90\x10\x7e\x8c\x6c\xf9\x73\x8a\x57\x10\x2c\x19\x32\xd4\xde\xe3\xa7\x12\x1b\x4d\x8f\x0c\x6c\x7e\x3e\x25\x29\x88\x99\x30\xff\x61\xda\x84\xa6\x9e\x14\x40\x78\xdd\xe8\xf3\x93\x09\xc4\x17\x3a\xd5\x6f\x67\xa7\x3b\x63\xe9\x3a\x02\x88\xda\x79\x7c\xc9\x90\x89\xb9\xc3\x2b\xdb\xe8\xe9\x14\x58\x4f\x04\x06\xe3\x70\xf9\xa6\x6f\x0d\x25\x5c\xcc\x50\xbe\xc8\x14\xc4\x4a\xbf\x0d\xdb\xbb\xdd\x0e\x6f\x4b\xbc\x07\xb1\x10\x9f\xf4\x25\xc3\xca\x0d\x96\x56\xee\xd5\x10\xf3\x9b\x3a\xd6\xcb\x12\xef\xf1\xed\x2d\x9e\x3e\xe3\xc3\x54\x5e\x66\x7a\x54\xf3\x50\x75\x47\xf9\xa6\xf7\xc6\x52\x9d\xc8\xf7\x45\xe1\x6f\x43\x47\xb4\x43\xb7\xd7\x1e\x6f\x02\xaa\x80\x37\x87\x0c\xf7\x03\xa1\x3e\xf7\xba\x22\xbc\x43\xe7\xf1\x41\x66\x7f\x39\x4a\xd9\xfe\xc8\x5e\xc7\xb8\x80\x7f\xe2\xc4\x64\xbb\x9b\xe7\x9e\xa2\x74\x7e\xcd\x9f\x67\x59\x8a\xd7\x11\x44\xdc\xc6\xf2\x1c\xf2\x5f\xaa\x3d\x25\x32\x2f\x64\x86\xcc\x9b\x30\x88\x0b\x97\xb1\xb5\x43\x17\x72\x56\xfa\x66\x6b\x37\x2f\x33\x2a\x4c\x7f\x4b\x2e\x8c\x7f\x28\xd1\x9a\x76\x8e\xc2\x6b\x1a\xbc\x65\x78\x32\x3e\x37\xbd\x8b\x7e\x3d\xd3\x24\x94\x72\x92\x32\x0f\xaf\xad\x9c\xaf\xcd\x96\x4b\x54\x7d\xaf\xed\x21\x89\x2f\x68\x6a\x5d\x78\x22\xb3\x35\x2d\x88\x31\x85\xff\x88\xaf\x1d\xd9\x6c\x60\x84\x05\x8b\x84\x7c\x79\x84\x3f\x01\x00\x00\xff\xff\xc1\x01\x03\x0f\x63\x03\x00\x00")

func migrations_testGoBytes() ([]byte, error) {
	return bindataRead(
		_migrations_testGo,
		"migrations_test.go",
	)
}

func migrations_testGo() (*asset, error) {
	bytes, err := migrations_testGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations_test.go", size: 867, mode: os.FileMode(436), modTime: time.Unix(1611628629, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
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
var _bindata = map[string]func() (*asset, error){
	"000001_init.down.sql":          _000001_initDownSql,
	"000001_init.up.sql":            _000001_initUpSql,
	"000002_add_test_link.down.sql": _000002_add_test_linkDownSql,
	"000002_add_test_link.up.sql":   _000002_add_test_linkUpSql,
	"migrations_test.go":            migrations_testGo,
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
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"000001_init.down.sql":          {_000001_initDownSql, map[string]*bintree{}},
	"000001_init.up.sql":            {_000001_initUpSql, map[string]*bintree{}},
	"000002_add_test_link.down.sql": {_000002_add_test_linkDownSql, map[string]*bintree{}},
	"000002_add_test_link.up.sql":   {_000002_add_test_linkUpSql, map[string]*bintree{}},
	"migrations_test.go":            {migrations_testGo, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
