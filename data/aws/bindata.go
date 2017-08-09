// Code generated by go-bindata.
// sources:
// cloud.json
// credential.json
// DO NOT EDIT!

package aws

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

var _cloudJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5c\xc1\x8e\xdb\x36\x13\xbe\xef\x53\x10\x3a\xfd\x3f\x10\x0b\x36\x2d\xdb\x69\xae\x01\x5a\x04\x68\x90\x00\x69\x7b\x29\x82\x05\x2d\x71\x1d\x61\x25\x51\xa1\xa8\xcd\x3a\xc1\x3e\x4d\x1f\xa5\x2f\x56\xd8\x5e\xcb\x12\x49\x79\x86\xde\xb5\xb6\x2e\x78\x49\xe4\xe1\x27\xf2\xd3\xc7\x99\x21\x67\x80\xe4\xc7\x15\x21\x41\xc1\x72\x1e\xbc\x21\x01\xfb\x56\x05\xaf\x36\x06\x5e\xdc\x55\xc1\x1b\xf2\xe7\x15\x21\x84\x04\x09\xbf\xdb\x9a\x09\x09\xbe\xb2\xfd\x53\x29\x45\xb2\x7f\x16\x05\x5f\x8a\xfb\xe0\x8a\x90\xcf\xdb\xd7\x25\x5f\xa5\xa2\x38\xcc\xf0\x63\xfb\x27\x21\x41\x26\x62\xa6\x52\x51\x6c\x16\xfb\x23\x95\xab\xb4\x48\xf7\x13\x36\xaf\x6d\xc6\xea\x6a\xc4\x59\xa5\x46\x93\xc3\xe0\x77\x51\xf0\xc3\x8c\x5b\x53\x83\x3a\xcc\xd1\xb6\x2e\xad\xd6\xd8\x6a\x4d\xac\x56\x1e\x3c\x1a\x3f\x6f\xff\x7e\x78\xd5\xff\x35\x1f\xbe\xa4\xe2\xd8\x97\x50\xd4\x97\x50\xeb\x97\x50\xeb\x97\xd0\x18\xcf\xee\x2d\xcb\xd2\x1b\x21\xfb\xd5\xfe\xc6\x31\x6a\xef\x50\x3a\xc7\x9d\x55\xe7\xb8\xb3\x3a\x70\xfc\x20\xf9\x4a\x14\xc7\xf8\x81\x1a\xee\x50\x56\x7e\x86\x86\x3b\xab\x03\xbf\x77\x92\x67\xac\x48\x6c\x04\x79\x8d\x11\xb0\x41\x75\x08\x36\xd6\xa5\xd5\xea\x40\xf0\x57\x51\x24\x76\x01\xf7\xb3\x1d\x17\xf0\xdc\xfc\x3e\xa5\xc5\x8a\x95\x42\x72\x1b\x45\x56\x8e\x2a\x51\xab\x2f\x88\xb0\xef\x42\x3b\x64\xbb\x43\xcb\xfe\x21\x07\xda\xbf\x89\xdb\xb5\x35\xb4\x59\x39\x2a\x84\xc4\x52\x6e\x41\x75\xca\xad\x21\x9d\x72\x6b\x08\xa0\xdc\x23\x65\x6b\xcb\x3b\x7b\xb1\x4e\x0a\xbe\xc6\xcb\x4c\xfb\x65\xa6\xfd\x32\x43\xe1\x75\xe0\x5c\x31\x23\xe1\x77\xe8\xfe\xfd\x97\x20\x1f\x59\x9d\x89\xa3\x94\x9b\x59\x3a\x6c\x1b\xeb\xd2\x6a\x75\xd6\xb5\x8f\xe4\xfb\x3a\x5f\xb2\x14\xa5\x69\x8f\xdb\xf6\xb8\xac\x03\x47\x5e\x8f\x62\x5e\x28\xc9\xb2\x3e\x96\x3f\x4b\x56\xdc\xde\xd4\x52\x41\xb9\xa0\x99\x47\x4f\x07\xcd\x80\x9e\x11\x9a\x01\x9d\x6e\x73\x31\x48\x8b\x4a\xb1\x22\xe6\xd7\x6a\x5d\x72\xcb\xfd\x80\xdf\x2b\x2e\x0b\x96\x5d\x57\xb7\xf5\x86\xac\xa2\x61\x9e\xc6\xb2\xb5\xeb\x09\xaf\x62\x99\x96\xfb\xaf\x31\x01\x31\x53\x7c\x25\xe4\x7a\x33\xfa\x0b\x2f\xb8\x64\x19\xf9\x58\xcb\x52\x54\xad\xbc\x13\x97\x9b\xe9\x27\x87\xa0\x66\xf9\xe6\xb7\x55\x5c\x0b\xa7\x2a\x67\x59\x76\x8c\x93\x06\x78\x02\x27\x8a\xe5\x94\xf3\x24\xad\xf3\xa3\x42\x69\x08\x07\x56\x54\x63\x15\x61\x59\x65\x4c\xae\xf8\x31\x52\x1a\xe0\x09\x9c\x5e\xa3\x38\xe5\x11\xc0\xc9\x04\x0c\xc1\xe9\x1e\x24\xa5\x23\x1c\x58\x45\xba\x9f\xcf\xb1\xb4\x28\xcc\xcb\x80\x38\x10\x7b\xad\x11\x9b\xe2\xbc\x3d\x8f\xc2\x08\x26\x66\x40\x5c\xa2\x70\xae\x31\x9b\xe3\x3c\x3e\x8f\xc2\xc9\x18\xa6\x66\x62\x5c\x76\x73\x6c\x6c\xe7\x18\x47\x6e\x0a\x25\x09\x0b\xe2\x09\xa9\x6b\x1a\x2e\x66\x58\x62\x80\x64\x06\xe0\x09\x31\xb9\x08\xd1\xac\xa0\x9d\x34\x11\x4f\x89\x4a\x34\x2d\x30\x2a\x2d\x90\xa7\x44\x25\xce\xbd\x62\x28\xb3\x9a\x80\x36\xa9\xb7\x22\x2f\x6b\xc5\xc9\x87\x52\xa5\x79\xfa\x9d\x27\xd0\x3e\xa2\xdd\x2b\x06\xd3\xab\x05\xe1\x44\x4d\xdf\x4a\xac\x8b\xc5\x70\x86\xb5\x41\x9c\xb8\xe9\xbb\x89\x74\xb3\x18\xce\xb1\x36\x88\x13\x35\x23\xcb\xe2\x3d\xed\x35\xcc\xcd\x80\x38\x71\x9b\x1a\x27\x00\x92\x1b\x94\xcb\x4c\xc0\x40\x51\x00\xa6\x33\x0b\x62\x98\x28\x80\x33\x9a\x0d\x32\x48\x14\x4c\xe1\x28\xb0\x40\x86\x89\x82\x29\x1c\x05\x16\x88\x5b\x14\xe8\xee\x86\x8c\x82\x15\x05\xb7\xd4\x06\xe9\x1c\x52\x1f\x7f\x27\xef\x1e\xcb\xc5\xea\x99\xb6\x73\x45\x41\xc9\x6c\x10\x34\xad\x53\xe5\x92\x50\xd2\x30\x01\x6d\x4e\xef\x79\x2e\xe4\x9a\x08\x6c\xce\x40\xaa\x25\xc1\x8c\x61\x41\xb8\xf0\xd2\x13\xc6\x74\x8c\xcc\x18\x12\xce\x18\x36\x88\x0b\x37\xdd\xc3\xe6\xb8\xe6\x80\x84\x13\x86\x0d\xe2\xc2\xcc\xc8\x17\x13\x8a\x2b\x9b\x24\x9c\x30\x6c\x10\x17\x6e\x46\x00\xd0\x08\x57\x38\xa5\x14\xf2\x35\x0b\xa2\xcd\xec\x93\x12\x92\xad\x1c\x4e\x27\xb4\xb3\xa5\x70\x2e\xb3\x41\x9c\xc8\x9d\xe8\x6d\x29\x05\xbd\xcd\x06\x71\xa2\x76\xb2\xbb\xa5\x70\xb2\xb5\x41\x9c\xc8\x9d\xec\x6f\x09\xe8\x6f\x16\xc4\x30\xfe\x96\xc0\xfe\x66\x83\x0c\xe1\x6f\x09\xec\x6f\x36\xc8\x30\xfe\x96\xc0\xfe\x66\x83\xb8\xf9\x9b\x4e\xae\xf1\xb7\xa6\xc5\x7d\x5b\x2f\xb9\x2c\xb8\xda\xb6\xb7\x77\x64\x83\x84\xdf\xb0\x3a\x53\xd7\x15\x57\x75\xd9\x98\x09\x09\x0a\x91\xf0\x8a\xab\x96\x89\xb4\x3a\x90\xfb\xee\xef\x6e\xfa\x46\x81\xe0\x8e\xcb\x2a\x15\x45\x75\xbd\x5c\x5f\xf3\xe2\xae\x3d\x61\xc2\xef\x3a\x8d\xfc\xc3\xb4\x87\xf7\x36\x5f\x3a\x09\x69\x38\x6f\xb5\xef\x4d\xad\xb6\x08\xf2\xbf\x25\x57\xec\xff\x3a\xb0\x94\x7c\x23\x5b\x12\xbc\x21\x4a\xd6\xbc\x19\x7c\x78\x85\x5a\x78\x01\x2e\xbc\x38\xc7\xc2\x51\x38\x06\x16\x8e\xc2\xf1\x79\x16\x9e\x80\x0b\x4f\xce\xb3\x70\x04\x2e\x1c\x9d\x67\xe1\x19\xb8\xf0\xec\x3c\x0b\x43\x5e\x1d\x9d\xc7\xab\x67\xe0\x1e\xcf\x7a\xf6\x98\x95\x65\xd5\x49\x00\x5b\xeb\x21\x8d\x8c\x2a\x2e\xef\xb8\xb4\x2f\xa3\x21\x59\xa6\x1a\xdc\x88\xc5\x3a\xb4\x52\x4c\xaa\x51\x27\x43\x05\xe3\x70\x1e\xd2\x9f\x74\xe4\x17\x51\xa9\x1b\x16\xab\x6a\x3f\xdd\x3c\x68\x01\x1e\x9e\x57\x39\x0a\x2a\x47\x9f\x43\x39\x8a\x54\x8e\x5e\x8e\x72\x53\x50\xb9\xe9\x73\x28\x37\x45\x2a\x37\xbd\x1c\xe5\xa0\xc4\x38\xeb\x49\x8c\x8e\xca\x45\x48\xe5\xa2\xcb\x51\x0e\xca\xec\xb3\x9e\xcc\xee\xa8\xdc\x0c\xa9\xdc\xec\x72\x94\x83\x8e\xa6\x59\xcf\xd1\xe4\xa8\xdc\x1c\xa9\xdc\xfc\x72\x94\x83\x6e\x8c\x06\xc2\x51\xb2\x05\x52\xb2\xc5\xf3\x4a\xb6\xe8\x48\xe6\x26\xca\x02\xbc\xcd\x1a\x08\x27\x51\xf4\xb7\xfb\x44\x59\x84\xe3\x61\x44\x79\x7c\xfa\xdc\x54\x60\x5f\xd9\xd9\x4b\x1d\x27\xc5\xf4\x65\xfa\x14\xa3\x0e\x91\x67\xdc\xf6\x06\x0c\xbb\xd3\x0b\x35\x47\xd9\x70\xd1\x47\x1d\xa2\xef\x25\x65\x3b\xbd\xcc\x74\x92\x4d\x5f\xa6\x4f\xb6\xc8\x21\x3e\x69\x38\x31\x26\x1d\x54\xb9\x53\xeb\x64\x47\xe5\x70\x35\x54\xe4\x50\x43\xbd\xb8\x72\xa7\x16\xfa\x8e\xca\xe1\xee\xb3\x91\xc3\x7d\xf6\xc5\x95\x3b\xb5\x53\xe1\xa8\x1c\xee\x3e\x1b\x39\xdc\x67\x5f\x5c\xb9\x53\x5b\x2d\x8e\xca\xe1\x4e\xd5\xe8\x92\xee\xb3\xbe\x57\xe4\x7b\x45\xbe\x57\x74\x29\xca\xf9\x5e\x91\xef\x15\xf9\x5e\xd1\xa5\x28\xe7\x7b\x45\xbe\x57\xf4\xf8\x74\xe8\x15\x6d\xff\xb5\x3a\xb6\x5b\x84\x68\x7b\xf8\x7e\x87\xaf\x3d\x7d\xed\xe9\x6f\x16\xfe\x66\xe1\x6f\x16\xff\xd5\x9b\x85\x71\x88\x3e\xfe\x37\x2f\xfe\x18\xdd\x2b\xe4\x8f\xd1\x7f\xc3\x61\xe0\x8f\x51\xdf\x88\x1c\x5e\x39\xdf\x88\xf4\x8d\x48\xdf\x88\xbc\x14\xe5\x7c\xb9\xe0\xcb\x05\x5f\x2e\x3c\x8b\x64\xbe\x11\x79\x42\x0d\x75\xb5\xff\xf5\x70\xf5\xf0\x4f\x00\x00\x00\xff\xff\x49\xf5\xf6\xde\x73\x53\x00\x00")

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

	info := bindataFileInfo{name: "cloud.json", size: 21363, mode: os.FileMode(420), modTime: time.Unix(1453795200, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _credentialJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x90\xc1\x6a\xc3\x30\x10\x44\xef\xfa\x8a\x41\xa7\x16\xf2\x05\xb9\xe5\x58\x72\xf4\x21\x87\x12\xcc\xda\x9e\x14\x35\xb6\x6c\xb4\x4a\x5a\x37\xe4\xdf\x8b\x36\xa5\x49\x5b\x9a\x8b\x10\x33\x6f\xd9\xc7\x9e\x1c\xe0\xa7\x34\x1e\x43\xc7\xe4\x97\xf0\xf2\xa6\x7e\x51\xc2\x2e\xe8\xd4\xcb\x5c\xb2\xd5\x20\x1f\x63\xc4\x86\x0d\x2a\xa6\x63\x68\xa9\x78\x58\x6d\xaa\xc7\x1f\x64\xbd\x1b\xd3\x20\xb9\x0c\xec\x02\xfb\xee\x52\xda\x57\xfd\x12\xcf\x0e\x00\x4e\xf6\x02\xfe\x55\xc7\xf8\xb5\xaf\x96\xb6\xa5\x6a\xbd\xe7\x5c\x87\xcb\x98\x21\x51\x06\x1a\x62\xf5\xfa\x57\xdb\x4b\xc3\xde\xec\xac\xc6\x9a\x33\x9e\x6e\xfa\x10\xa7\x83\xc9\x64\xbe\x67\x6f\xe9\x79\xf1\xbf\x83\xb2\x4d\xcc\x37\x2a\x7f\x3d\xee\x20\xdf\x32\x95\x31\xb8\x3a\xdd\x17\x72\xc0\xd6\xce\x94\xe5\xe5\x7a\x24\xbf\x3f\x34\x4c\x91\x99\x5a\xb8\xad\x3b\xbb\xcf\x00\x00\x00\xff\xff\x8d\x48\x1c\xa0\xa8\x01\x00\x00")

func credentialJsonBytes() ([]byte, error) {
	return bindataRead(
		_credentialJson,
		"credential.json",
	)
}

func credentialJson() (*asset, error) {
	bytes, err := credentialJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "credential.json", size: 424, mode: os.FileMode(420), modTime: time.Unix(1453795200, 0)}
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
	"credential.json": credentialJson,
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
	"cloud.json": &bintree{cloudJson, map[string]*bintree{}},
	"credential.json": &bintree{credentialJson, map[string]*bintree{}},
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
