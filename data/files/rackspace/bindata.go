// Code generated by go-bindata.
// sources:
// cloud.json
// DO NOT EDIT!

package rackspace

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

var _cloudJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x98\xdd\x6e\xe2\x38\x14\xc7\xef\x79\x8a\xa3\x5c\xcd\x4a\x94\xc1\x06\x52\xe8\x1d\xa2\x19\x06\x0d\x69\x2a\x68\xe7\x43\xab\x5e\x64\xc1\xa2\x51\x83\x8d\x9c\x80\x44\x47\x7d\xa5\x79\x88\x3e\xd9\xca\xe1\x23\x60\x63\xe3\xa4\x68\xb5\x37\xc3\x10\x9f\xfa\xfc\xfe\xe7\x8b\xd8\xbf\x2b\x00\x0e\x0d\xe7\xc4\xb9\x01\x87\x87\x93\x97\x64\x11\x4e\x88\x53\x15\x8f\x09\x5d\x39\x37\xf0\x77\x05\x00\xc0\x99\x92\x95\x53\x01\x78\xca\x56\x38\x99\x45\x8c\x26\xfb\xd5\xdf\xd9\xbf\x00\x4e\xcc\x26\x61\x1a\x31\x2a\x76\xbb\x0d\xe3\x38\x4c\xae\xbe\x30\x9e\xc2\x0f\xc6\xd3\xe7\x2a\x3c\xfc\xac\xc2\xe3\xb8\x9b\x6d\x9f\xd9\x6f\x36\xca\xac\xbf\xfc\xc8\x1f\xbf\x32\x4a\xf2\xdd\xb3\x47\x62\x7d\xfb\xed\x29\xfb\x7c\xab\xea\x5d\xf7\x9e\xa3\x49\x38\x63\x5a\x67\xc1\xe8\xd6\xe8\x4c\xac\x5b\x3b\xbb\x13\xda\x08\xa7\xf0\x3d\xe2\xb3\x88\x46\xa1\xd6\xed\xa0\x6b\x76\x2b\xd6\xad\xdd\x8e\xd7\x53\x4a\xd6\x55\xe8\x2e\x93\x94\x87\x71\x14\x9e\xf2\x38\xfe\x65\xf6\x28\xd6\xad\x3d\x7e\x65\x74\x06\xdf\x18\x9d\x9d\xf2\xf4\xf5\x5b\xdf\xe8\x49\xac\x1f\x7b\xda\x57\x53\x44\x93\x34\xa4\x13\xf2\xb0\x5e\x90\x13\x35\x95\xbc\x2c\x85\x83\xbe\x77\xe7\x8d\xba\x43\x74\x85\x72\x3f\x53\x92\x4c\x78\xb4\xd8\x01\xf6\x09\x25\x3c\x8c\xe1\xc8\x66\x12\xa6\x64\xc6\xf8\xfa\xd0\xe0\x7e\xc9\x17\x2c\x21\x22\x63\xe9\x32\x8c\x61\x4c\xf8\x8a\xf0\xfc\x6f\x56\x93\x85\xf0\x89\x72\x9d\xe1\xfc\xe8\xfb\x34\x4a\x5e\x9c\x1b\xc0\xf5\x93\x41\x93\x89\xb1\x05\x31\xbe\x08\x31\x96\x88\xb1\x4c\x8c\x5c\x3b\xe4\xa6\x05\x72\xf3\x22\xc8\x4d\x09\xb9\x29\x23\xb7\xed\x88\xdb\x16\xc4\xed\x8b\x10\xb7\x25\xe2\x76\xc1\x20\xf7\x02\xff\xfe\xf1\xc1\x33\x05\xb9\xc7\xe6\x8b\x65\x4a\xf4\x41\xde\x19\x04\x8b\x34\x9a\x47\xaf\x64\x5a\xb4\x32\x1a\xb5\xeb\x96\x15\xa4\x3e\xae\x39\xa4\x26\xae\x85\x21\xe5\x5a\xb8\xae\xd9\x31\xa2\x96\x05\xe4\xa1\xd1\x87\x28\xe5\xfc\x23\x3b\xc8\x46\xdd\x02\xf2\xd0\xe8\x43\x90\xc8\x95\x13\x6e\x57\x93\xae\x0d\xa5\x7b\x29\xca\x86\x5c\x96\xe6\xce\x19\x7c\x0e\x8c\xc9\x1e\x7c\x0e\x0c\x89\x16\xab\xa5\x4b\x11\xb5\x94\x2e\xef\x9c\x67\x35\xe4\x7c\xc3\xaa\xcb\x77\x21\x56\xb9\x20\x1b\x75\x99\xb5\xd1\x3c\xcf\x6a\xc8\xfc\x86\x55\x97\xf5\x42\xac\x4a\x5d\xba\x0a\xac\x6b\x01\xdb\x39\x07\xdb\xb9\x04\x2c\x96\xab\xa0\xa3\xc0\x76\x2c\x60\x11\x3e\x47\x7b\x64\x51\x1a\x57\xe9\x26\x84\x15\x5e\xe4\xd6\x8d\xc0\xbe\xe7\x07\xa3\x5f\xc6\x2e\xf3\xc9\x9c\xf1\xb5\xa1\xd1\xb6\x06\xa5\x7f\x9b\xcc\x03\x75\x47\x68\xe8\xad\x3d\xa1\xae\xbd\x8a\x12\xca\x75\x60\x1e\xa6\x3b\x42\x43\x47\xed\x09\x75\x4d\x55\x94\x50\x9e\x01\xe6\x41\xba\xcf\xb2\xa1\x32\xf3\x34\xeb\x8a\xb3\x28\xa3\xd2\xfb\xc8\xfc\x02\xbd\xa3\xc4\x4d\x0b\xca\x23\xa3\x8f\x50\x2a\x5d\x84\xcd\x2d\x1e\xdc\xf9\xde\x43\x77\x08\xdb\x37\x51\xb8\x7f\x1c\xdd\x07\x63\x0f\x56\x18\xc6\x7e\x77\x38\xd4\xa2\x8f\x23\x3a\x8b\x09\x0c\x68\x4a\xe2\xf7\x3f\xf0\x93\x30\xfa\xfe\x07\xbc\xd6\x15\x76\x71\x1d\x56\x0d\xf8\xe4\x42\x8f\x71\x52\x05\x5c\x6b\xf6\x9f\x5f\xff\x3a\xad\x2f\xa0\x3e\x49\xc3\x18\x7a\x31\x5b\x4e\xb7\xa2\x12\xf8\x7e\x78\xa0\xc8\x74\x29\x2f\x04\xca\xe1\xa0\x6d\x1e\x0d\x06\xa1\xbe\x77\x3b\x78\xf4\xb5\x4a\x6f\x45\xb8\xff\x2b\x9d\x48\x79\xa7\x38\x71\xa4\x28\x2b\x74\xd8\x1d\xf5\xbd\xff\xa9\x4e\x84\x95\x93\x88\xa5\x50\xf1\x2b\xb3\xd2\x9f\xab\xb5\xb2\xda\x75\x11\x93\x4f\xa8\x5e\x5e\x97\xec\x7a\xab\xad\x7e\x5e\x5b\x03\x5b\x4b\xd3\x1f\xc0\xb5\xd2\xdc\x4d\xc6\x0e\xa4\xb9\xa5\xa4\x29\x69\xc3\x16\xd2\x2c\x27\xce\x66\x3e\x96\xca\x5c\x63\x93\x39\xb7\xbc\xba\xed\x58\x55\x73\x27\xcf\x99\x16\x52\x06\x8d\x5d\xea\xb6\x67\x92\x0b\x55\x66\xbb\xa8\xc0\xdd\x61\xc6\xa2\x3a\xd5\x51\xba\x53\xb8\xbf\xe4\x7a\x59\xfe\x43\x38\x25\x69\x76\xc3\xb5\x51\xec\x08\x37\x47\xf7\xa8\x79\x2c\xf2\x55\x01\x88\x6a\xd7\xb5\xfc\x22\x40\x0d\x83\xb2\xbe\xb9\xba\xcd\x37\xdb\x5e\xe0\xde\x40\xca\x97\x64\xff\xf4\xad\x72\xf8\xf9\x54\x11\xff\x7b\xfb\x37\x00\x00\xff\xff\xa7\x6d\x30\x67\x12\x16\x00\x00")

func cloudJsonBytes() ([]byte, error) {
	return bindataRead(
		_cloudJson,
		"cloud.json",
	)
}

func cloudJson() (*asset, error) {
	bytes, err := cloudJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cloud.json", size: 5650, mode: os.FileMode(420), modTime: time.Unix(1453795200, 0)}
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
	"cloud.json": cloudJson,
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
	"cloud.json": {cloudJson, map[string]*bintree{}},
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