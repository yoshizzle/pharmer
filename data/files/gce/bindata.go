// Code generated by go-bindata.
// sources:
// cloud.json
// DO NOT EDIT!

package gce

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

var _cloudJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x98\x51\x73\xda\x3a\x16\xc7\xdf\xf3\x29\x34\x7e\xda\x9d\x31\x2e\x10\x9a\x64\xfa\x46\x9c\xb4\x4b\x37\x4d\x58\x20\xed\xc3\x4e\x87\x39\xc8\x07\xac\x22\x4b\xbe\x92\x9c\x5c\xda\xc9\x77\xbf\x23\xe3\x18\x63\x6c\x10\x2f\x49\xd0\xf9\xeb\xe8\xa7\x73\xfe\x11\x96\xff\x5c\x10\xe2\x09\x48\xd0\xfb\x44\xbc\x15\x45\xcf\xb7\x03\x28\x5e\xbc\x4f\xe4\xff\x17\x84\x10\xe2\x45\xf8\x92\x8f\x12\xe2\xfd\x05\xde\x05\x21\x3f\x73\x8d\xc2\x15\x93\x42\x97\xba\x3f\xf9\xcf\x32\x60\xf3\xdd\x83\x36\xa8\x04\x79\x9e\x16\x09\x08\xf1\xb8\xa4\x60\x8a\xf8\x2d\xaa\x35\x72\xdc\x90\x50\x66\xc2\x6c\x7c\x32\x95\x99\x89\x49\x08\x4a\x72\x26\x60\x37\xe9\xb7\x14\xb8\x5b\x29\x1f\xca\x74\x07\x41\x9b\x5e\x67\x51\xca\xaa\xa3\xb4\x71\x34\xf2\x8a\xc1\x9f\xf9\xef\x37\xbf\x8d\xfc\x07\x1e\x27\x9f\xc5\x48\xee\x80\x73\xd4\x3e\x79\x9a\x9c\x02\x7d\x45\xbb\x38\xd4\x90\xb6\xa3\x0b\x57\xa4\x10\x85\x51\xc0\x5b\x91\x6c\x0d\x29\xe3\xe4\x96\x67\xcb\xa5\xf6\xc9\x48\xbe\x9e\xac\x20\xdd\xe6\x3c\x64\x2b\x03\xf5\xea\x96\x81\x7a\x81\xcb\xc0\xf2\xdc\x1a\xdf\x67\x4a\xa6\xd8\xbc\xa9\xa9\x09\xc8\x97\x98\x69\x0e\x4c\xf8\xe4\x16\xf9\x8a\x65\xc9\xd1\x5d\x61\x9e\xad\x2c\xae\xdf\x12\xa1\xad\x11\x67\x8f\x0c\x35\x03\xf2\xaf\x19\xb0\x57\x10\xff\x6e\xe9\x49\x0c\x62\x15\x67\x50\x1a\x7c\xab\x3e\xba\x01\xd0\x0c\x0a\xbb\xee\x75\xa5\x32\xbe\x68\x19\xa7\xe7\xa1\x7f\x85\xb4\x95\x7c\x26\xd7\x1b\xe9\x93\x5c\x72\x1a\x57\x48\x65\xe2\x36\xe6\x4a\xf0\x10\xbc\x12\xac\xd3\x97\x27\x0d\x13\xda\x80\xa0\x38\xdb\xa4\xd8\x70\xde\xe8\x75\x96\x1f\x5e\xbd\x8e\x4e\x80\xf3\x1d\x6d\x84\x9a\x2a\x96\xbe\x6f\xe9\x50\x40\xc1\xe0\x4a\xaa\x4d\xee\xb4\x5a\x28\xb5\x49\x7b\xe5\x67\x05\x89\xfd\x1c\x5c\x37\x56\xb6\x40\x10\xbd\x8e\x25\x8d\x40\x45\x9d\x5e\x2b\x47\x8b\x6a\x0f\xa6\x88\x9f\xe2\xb9\x0c\xae\x3f\xba\x02\xf5\x9d\x80\xfa\xe7\x01\xf5\x6b\x40\xd7\x81\x33\xcf\xc0\x89\x67\x70\x1e\xcf\xa0\xde\x30\x67\x9c\x1b\x27\x9c\x9b\xf3\x70\x6e\xea\xfd\xea\x3a\xdb\xe7\xca\xcd\x3f\x57\x67\x1a\xe8\xaa\x46\x74\xe5\x4c\x74\xd9\x77\x43\xda\xd3\xb9\x30\x5d\xd6\x4d\xd4\xeb\x9f\x82\x8a\xd9\x2a\x4e\x30\x39\xee\xe9\x06\x51\x15\xe7\x3f\x6c\x15\x93\x6f\x98\xd8\x8f\x27\x5c\xdd\xbb\x74\xe4\x39\xea\xe9\x06\x91\x23\x4f\xdd\xd5\xfd\x2b\x47\x9e\xa3\xa6\x6e\x10\x39\xf2\xd4\x6d\xfd\xb1\xef\xc8\x73\xdc\xd5\x4d\x2a\x47\xa2\x03\x5f\xf7\xba\x03\x47\xa6\x13\xbe\x6e\x94\x39\x52\x1d\x38\xbb\xdf\xbd\x71\xa0\xa2\x69\x76\xda\xd9\x35\xd1\x01\x51\x38\x7e\x3e\x69\xeb\xc0\x95\xe6\xa4\xaf\x6b\x22\x17\x9a\xba\xa9\x2f\x03\x17\x57\xdb\x85\x4e\xba\xba\x26\x72\xa1\xa9\x5b\xfa\x3a\x70\xf1\xb4\x5d\xe8\xb4\xa7\xeb\x2a\x17\x9e\x43\x43\x0f\x02\x17\x47\xdb\xb5\x1c\x1c\x7d\x20\x73\x61\x3a\xb4\xf3\xcd\xbb\x83\xca\x67\x35\xaa\x30\x42\x61\x18\xf0\x86\x27\xb5\x54\xc9\x17\x16\xa1\xb2\x8b\x7c\x91\x72\xc5\x31\xe4\x32\xab\x7c\x21\x44\x4c\xa7\x1c\x36\x9f\xa5\x4a\xc0\x58\xd5\x2f\x2d\x2b\x8f\x9f\x20\x84\x34\xf9\x23\xaa\x4d\xfe\x67\xf7\x28\x99\xc6\xa0\x12\x54\x01\xa4\xa9\xa6\x32\xc2\x80\xca\xe4\x03\xe5\x99\xbd\x60\x74\x76\x48\x36\x63\xf5\x09\xb4\x71\x5a\x24\xf4\xb9\x53\xb4\x91\x0a\x56\x58\x9f\x56\xcc\x7a\x2b\xf9\x97\x0c\x79\xb4\xff\xfc\xbc\xdb\xc4\xf6\xda\x4d\xa5\x58\xb2\x55\x5e\xa0\xf0\x7e\x3e\x9e\x3c\x7d\xbd\x0f\x67\xf3\xd1\x5d\x85\xc1\x26\x92\x2a\x29\x2e\xec\xf3\x54\xc9\x5f\x48\xcd\x9c\x45\xfb\x9a\xbc\x74\x9f\xf2\xa2\xdb\x78\x3d\x05\x87\x05\xf2\x5d\x23\x48\xde\x09\x32\xde\x8a\x49\x5d\xcd\x44\x9a\xe5\x0d\x31\xf8\xb7\xf1\xca\xc8\x9b\xef\xb8\x91\xe9\xfd\xe4\xfb\x28\xbc\x9f\x0f\xc3\xf0\xe9\xf9\x71\xd6\xbe\x1b\x8d\xea\x85\x51\x9c\x03\xa5\xf6\xca\xd4\xbc\xa5\x42\x34\x6c\xd2\x34\xef\x6b\xba\x9d\x41\x1a\xa7\xec\x6d\x0e\x14\x42\x65\x83\xc5\x5f\xf5\x2b\xc9\x3a\x5b\xa0\x12\x68\x50\x7f\x47\xa5\x9b\xdf\x83\xbc\x6c\x23\x36\x71\x2f\xb8\x0e\xda\xcf\xd0\x5a\x74\xfb\xf2\xa5\xe2\xee\x08\xed\x80\x51\x19\x1e\x38\x2a\xc2\x25\x64\xdc\x4c\x53\xa4\xfb\x73\x0a\xf3\x8f\xd2\x09\x88\x55\xfe\x72\xa7\xd7\x0d\xfa\x83\x41\xd0\x0d\xba\x1f\x2a\x87\x11\x29\xab\x19\x36\xcd\xe8\x36\xe8\x81\xe7\x17\x45\x7c\x94\x11\x86\x2c\x52\xba\x80\xab\x5e\xa8\x05\x2c\xf8\x7b\xc6\x6f\x52\x30\x23\x15\x13\xb9\x1b\xde\xff\x71\xbc\x36\xf9\x83\x5c\xad\xb6\xda\xc6\xa4\x76\xd5\x56\x09\xdf\x06\xee\x50\x1b\x26\xca\xbb\xec\xfb\x92\x1d\xe4\xa0\x0d\xa3\x1a\x41\xd1\x78\x0f\xa0\x1a\x28\xb2\x4f\x30\xe5\x8c\x82\xae\x5e\xbb\x6c\xc5\x85\xb6\x66\xb2\x85\xaa\xd4\xa8\xd7\xf5\xf6\x35\x77\x32\x01\x96\xaf\xbe\x0e\x6c\xb9\xf8\x5e\x05\xa3\x84\x69\xeb\x8d\x50\x0a\xa3\x64\x6e\xd7\x47\x48\x50\xa7\x40\xf1\x81\x2d\x91\x6e\x28\x47\xff\x81\x25\xcc\xe4\xdd\x50\xfe\x74\xcf\xf2\xfe\xd8\x7a\x4b\x1b\x14\xe6\xbb\xe4\x59\x82\x0f\xd6\xf5\xfe\x5d\xe1\x86\xed\x71\x14\x72\xd0\xda\x9f\xa0\x96\x99\xa2\xf8\xbf\x4c\x9a\xbd\xbb\x79\x02\x4d\x06\xb9\xca\x1b\xde\x1f\x54\x95\x02\xcd\xab\x54\xeb\x71\xe5\xf4\xb6\xfe\xef\x2c\x39\x08\x81\xbc\xb5\x93\x43\x8e\xca\xb4\xf5\x5c\xda\xc2\x7a\x11\x2e\x58\xe5\x05\x43\x65\x31\xc9\x19\xdd\x54\x97\x14\x52\x34\x98\xe6\x07\x2e\x62\x29\xd7\x33\xb9\x46\x31\xcc\x4c\x2c\xda\x7c\x33\x59\x00\xb5\x82\xdf\x87\x82\xe9\xd3\xe7\xd9\xc3\x53\xf8\xdf\xe7\xf1\x7c\x3c\x7c\x1c\x85\x6d\x29\x86\xe3\x91\xce\x5b\x7f\x0b\x9a\xd1\x61\x16\x31\xd3\x2a\x2d\x76\x3c\x34\x86\xd1\x43\x51\x2a\x39\x9f\x72\xc4\x74\x24\x0c\xaa\x97\xfc\xfb\xe2\xb2\xd1\xea\xe3\x6c\xc1\x19\x1d\x8d\x6b\x07\xc0\xee\x38\xba\x78\xfb\x27\x00\x00\xff\xff\x20\x24\x55\x4d\xc3\x15\x00\x00")

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

	info := bindataFileInfo{name: "cloud.json", size: 5571, mode: os.FileMode(420), modTime: time.Unix(1453795200, 0)}
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