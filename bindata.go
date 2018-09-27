// Code generated by go-bindata.
// sources:
// files/dot.tmpl
// files/map.toml
// DO NOT EDIT!

package main

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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _filesDotTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x93\x41\x6b\xdb\x30\x14\xc7\xcf\xd2\xa7\x78\x98\x1e\xd3\x96\x5c\x27\x34\x18\x64\x94\x5e\xda\x4b\x6f\x21\x8c\xe7\xe8\xd5\x35\x91\x25\xf3\x6c\x77\x0b\x42\xdf\x7d\x58\x72\x32\x7b\x5b\xc8\x21\x37\x3f\xf8\xbf\x3f\xbf\x9f\x25\x3d\x3e\xc2\x13\x39\x62\xec\xc9\x40\x79\x84\xba\x69\x71\xdf\x43\x83\x6d\x4b\xbc\x82\xcd\x2b\xbc\xbc\xbe\xc1\xf7\xcd\xf3\x9b\x34\x75\xc5\xd8\x7e\x40\x08\x0f\x2f\xd8\x50\x8c\x10\xa4\x70\xf4\x93\xd1\x1d\x74\xcf\x03\x29\x29\x9c\x37\x04\xdb\xae\x3f\x5a\xd2\xc5\x7b\x6d\x2d\x99\x15\xfb\xc1\x19\x32\xc5\x0a\xba\x0f\x6c\x49\x97\xfe\xd7\x4e\x49\x29\xba\xa1\xcc\x85\x7b\x3b\x74\x3d\xf1\x8f\xca\xa3\x1d\x3b\x53\xcb\x76\xef\xad\x67\x6d\x90\x0f\x95\xb7\x86\x1c\x7b\xb3\x93\x42\x84\x70\x0f\x8c\xae\x22\xb8\x3b\xd0\x71\x05\x77\x9f\x68\x07\x82\x2f\x1a\x1e\x9e\x3c\xda\x2e\xc6\x31\x04\xe3\x77\x08\x63\x24\xc6\xad\xc5\x92\xac\x2e\x42\xc8\xe1\x09\xbf\xd8\x41\x9c\x0a\xc9\x99\xb4\x97\x82\xa0\xa1\x18\xd7\x0b\x25\x85\x48\x14\xa0\x61\xc1\x21\x45\xfc\x1f\x3f\xee\x7b\xcf\x67\x01\x98\x0c\xde\x3d\x53\xd7\x57\x4c\xe4\xae\xf0\x7f\x1b\xf7\xb3\xc0\x84\x9f\x71\xc7\x21\x46\xb8\xff\x0a\x01\x52\xe6\x06\xb1\xb4\xbf\x30\x9b\xf1\x5d\xf0\x9a\xee\xc4\xdf\x27\xe3\x4d\x45\x5c\xda\x81\xae\x68\x3d\xa7\xf5\xec\x75\xc2\xcf\xb8\x69\x3a\x99\xe5\xd8\x0d\x6a\xb9\x60\x79\x6a\x67\xc6\x0b\x6a\x86\x6c\xfd\x49\x8c\xa5\xa5\x7f\xfc\x88\xda\xb6\x76\x87\x2b\x76\x9b\x3f\x0d\x59\xf1\xec\x91\xb9\xf3\x78\x92\x9c\xa5\x6f\x30\x9d\xb5\x2c\x75\x27\xe4\x2c\x1b\x20\xbd\xcc\x0e\x1b\x52\xe9\x3d\xac\x55\xbe\x3e\x6b\x35\xfd\xec\xb5\x9a\x03\xad\x15\x44\x19\x7f\x07\x00\x00\xff\xff\x73\x55\xba\xc3\x11\x04\x00\x00")

func filesDotTmplBytes() ([]byte, error) {
	return bindataRead(
		_filesDotTmpl,
		"files/dot.tmpl",
	)
}

func filesDotTmpl() (*asset, error) {
	bytes, err := filesDotTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/dot.tmpl", size: 1041, mode: os.FileMode(420), modTime: time.Unix(1538021405, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _filesMapToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\xc1\x4e\xc2\x40\x10\x86\xef\x7d\x8a\x09\x77\x4d\x00\x3d\x72\x28\x18\xbc\x48\x20\x41\xbd\x10\x0e\x43\x3b\x94\x89\xed\x4e\x9d\x9d\x42\xf0\xe9\x4d\xa9\x6d\xe8\x6a\x88\xc7\xd9\xec\xf7\xcf\xff\x65\x1c\x16\x04\x13\x18\xbc\x92\xb7\x41\x14\x6d\x32\xc1\xdc\x6f\xa3\x08\x00\xa0\x19\xee\x87\xdb\xcb\xd4\xfe\x7c\x56\x39\xc1\x42\x76\x9c\x13\xc4\xe9\x91\xd4\xd8\xb3\xcb\x6a\x18\x13\x13\xed\xe8\x66\x0a\xf1\x75\x55\x92\xde\xcd\xd1\x79\x38\xb1\x1d\xda\xa4\x27\x3a\x72\x42\x7e\x70\xf9\x5b\xef\x85\x09\x0c\xfb\x41\xa3\x7e\xd0\x4c\x5c\x42\x6a\xb0\xd4\x0c\x1d\x7f\x91\xde\x84\xc7\x7d\x38\xae\x5b\x1b\xc4\x19\x39\xf3\x80\x2e\x85\x95\x4a\x21\xf5\xd7\x30\x65\xc3\x45\x89\x89\x75\x5a\x3f\x63\xe8\x35\x93\x82\x60\x8a\xc9\x07\x2c\x44\x09\xe6\x4a\x9f\x15\x39\xcb\xcf\x4d\xde\xa5\xc6\x55\xad\x36\x25\x90\x5a\x1b\x9e\xe1\x45\x5c\x46\x7a\x9b\x0b\x7c\xde\x99\x4e\xcd\xe2\x38\xf5\xbf\xc8\x4d\x4a\x39\x1f\x49\x71\x97\x53\xe7\x71\xfd\x16\xca\xac\x2a\x7f\x80\xb7\x32\x45\x6b\x4f\xd2\xec\xbd\x2a\xd2\xc3\x43\x8b\x92\x12\xc6\x1c\x96\xfb\x7d\x77\x96\xdb\x01\x81\xce\x5c\xb4\x2a\x42\x70\xf4\x17\xf8\x10\x5c\xe1\x80\xf6\x2f\xee\xb1\xcf\x4d\xc9\x8c\x14\x56\x98\xb1\x43\x63\x71\x41\xc6\xf8\x3b\x00\x00\xff\xff\xb9\x56\x71\xbf\x28\x03\x00\x00")

func filesMapTomlBytes() ([]byte, error) {
	return bindataRead(
		_filesMapToml,
		"files/map.toml",
	)
}

func filesMapToml() (*asset, error) {
	bytes, err := filesMapTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/map.toml", size: 808, mode: os.FileMode(436), modTime: time.Unix(1538021419, 0)}
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
	"files/dot.tmpl": filesDotTmpl,
	"files/map.toml": filesMapToml,
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
	"files": &bintree{nil, map[string]*bintree{
		"dot.tmpl": &bintree{filesDotTmpl, map[string]*bintree{}},
		"map.toml": &bintree{filesMapToml, map[string]*bintree{}},
	}},
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

