// Code generated for package configs by go-bindata DO NOT EDIT. (@generated)
// sources:
// configs/config.yaml
package configs

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

var _configsConfigYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x93\xcf\x4f\x1b\x47\x14\xc7\xef\x96\xfc\x3f\x8c\xe4\x73\xcc\x82\x03\x26\x73\x2a\x49\x48\x4b\x14\x52\xab\x76\xc4\xb1\x1a\x7b\x1f\xcb\xa2\x5d\xcf\x30\x33\x0b\xa6\x27\x12\x99\x46\x8d\x4c\x9d\xb6\x21\xa4\x2e\x6a\x6b\x35\xa8\x48\x6d\x0c\x87\x36\x4d\xb1\x13\xfe\x19\x66\x7f\x9c\xfc\x2f\x54\x3b\xe3\x1a\x43\x7b\xdb\xf9\xbc\xb7\x6f\xbe\xef\xfb\xde\xe4\x92\xdd\xbd\xe8\x7d\x2f\x7c\xf9\xf4\xa2\xff\x36\x9b\xc9\x85\x87\x7b\xea\x59\xd7\x40\x1c\xf7\x3e\x44\xef\x7b\x1f\x2f\x3d\x8c\x3a\xcd\xf8\xfc\x79\xdc\x6d\x85\xc7\x5d\x35\x68\x0f\x07\xad\xa4\xff\x2a\xee\xbd\x5e\x93\x92\x45\x3f\x7c\xa3\x9e\xff\x1e\xfd\x76\xa2\xda\xbf\x0c\x07\x2d\xb5\xfb\x38\xee\xbd\x8b\x4f\xfa\xaa\xfd\x52\x7d\xdb\x52\x5f\x7e\xaf\x76\x8f\xa2\x4e\x33\x3c\xdc\x51\xaf\x7f\x0d\x5b\x8f\xa3\xfe\x9b\xf0\xe0\x6d\x72\xf0\x47\x36\x53\x06\xbe\x09\x1c\x67\x33\x08\x7d\x16\xd4\x97\xa9\x0d\x18\xd9\x50\x0d\x9c\x94\x7c\x22\x25\x2b\x51\x2e\x31\x9a\xb7\x2c\x4b\xe7\x00\xb1\x2b\xae\x0f\x34\x90\x18\xcd\x69\xb4\xc2\x5d\x09\x57\x58\x36\x93\x5b\x60\xec\x4a\x03\x46\x6b\x78\xd2\x4e\xba\x7f\x86\xfb\xa7\xc9\xd3\xf6\x58\xe7\x48\xd5\xd5\x90\xc9\x57\x67\x2f\xa2\x17\xc7\xe1\xc1\x91\x3a\x3f\x50\x6f\x5e\xa9\x27\xc7\xf1\x5f\x27\xea\x43\x33\x9b\x59\x60\x4c\x6b\xbe\x0b\xab\x24\xf0\x64\x89\x38\x50\x76\xbf\x00\x8c\xa6\xb5\xa6\x65\xd2\x98\x44\x9a\x3d\xa0\x4e\x99\x6c\x42\x89\xc8\x35\x8c\x84\xa4\x9c\x38\x30\xe5\x51\x47\x8c\x82\xf7\x5c\x0f\x1e\x12\x1f\x30\x22\x8c\x4d\xb0\xc5\x86\xc4\x28\xef\x51\xed\x48\xee\xe2\xdd\xb3\x8b\xc1\xcf\x66\x56\xe1\x57\x3b\xc9\xe1\x8e\xe9\x33\x0d\x3e\x62\x1e\x25\xf6\x7f\x2f\x09\x34\x17\x13\x29\xda\xf4\x47\xdc\xc3\x28\x1d\x1f\x9e\x9a\x9a\x9e\x29\xe6\xad\xbc\x95\x9f\xc6\xa9\xd3\x53\x42\x12\xe9\xd6\x2e\x7f\x58\xf2\x89\x03\xcb\xa4\x61\x1a\x9a\x45\x28\x87\x96\x6f\x5f\x0b\x2f\x78\x1e\xdd\x5a\x6c\x48\xa1\x8d\x41\xe8\x06\xca\xaf\x33\x67\xe2\x1b\x2e\x0f\xac\xee\xe8\x31\x85\xfb\xa7\xe1\x5e\x4f\x9d\x7d\x67\x9a\x18\x0e\x3a\xf1\xf9\x8f\xe1\xd7\x47\x63\x1e\x75\x9a\xea\x7c\x37\xe9\xf6\xd5\x4f\x67\x51\x77\x47\xb5\x9f\x84\xfb\xa7\xd9\xcc\x5d\x22\xa1\x4a\x04\x98\x19\xdc\xae\x6c\x33\xc0\xc8\xdf\x16\x1b\x9e\x16\x25\x80\xd7\xb5\x93\x9c\x52\x99\x92\x12\x11\x62\x8b\x72\x1b\xa3\xb5\xda\x4c\xe1\xe6\xec\x5c\x71\x5e\xef\x17\x15\x32\x9d\x4f\xda\xb8\x65\xe5\x8b\xd3\xb8\x50\xb0\x8a\xa6\xa6\x19\x45\xd5\xa3\xce\xe7\x02\xf8\xa6\x5b\x83\x94\x57\x48\xd5\x83\x12\x87\x55\xb7\x31\x0a\xa6\xf4\xce\x1a\xe1\x02\x24\x46\x81\x5c\x9d\x37\xf7\x71\xa1\x77\x12\xa3\x0a\x0f\x60\xb4\x12\x4b\xb6\x07\x77\x68\xbd\x2e\x26\xd6\xe4\x53\x06\xf5\x11\x2b\x98\xd5\x35\xcf\xe5\xfe\x4a\x25\xea\x34\x8d\x2d\x49\xf7\xef\xe1\xa0\x93\xcd\xdc\x5f\xa9\xe8\x86\xcb\x50\xe3\xe9\x6d\x60\xdb\xdb\xb5\xf5\xed\x14\x2d\x09\x11\x00\x37\x92\x6e\x4c\xe8\x5d\x6c\x30\x97\x03\x46\xc5\x19\xcb\x54\x5f\xf4\x89\xeb\x8d\xcb\x66\x33\xfa\x8c\x2f\xcd\x10\xbe\x64\xf9\x8d\x8d\x7c\x8d\xfa\xba\x11\xfd\xfa\x6e\xce\xcd\xfe\xeb\xab\xb1\x65\x66\x7e\x6e\xb6\x58\x2c\xdc\xba\xf5\xd1\x44\xea\xd8\x63\xc9\x1d\x2a\x60\xcb\xa6\x0c\x68\x75\xdd\x59\x37\x02\xcb\xe5\x07\x69\xc8\xb8\x71\x8f\x53\xff\xff\xcb\x54\xe8\x78\x7f\xae\x87\xff\x09\x00\x00\xff\xff\x3b\xb6\x8e\xbf\xa9\x04\x00\x00")

func configsConfigYamlBytes() ([]byte, error) {
	return bindataRead(
		_configsConfigYaml,
		"configs/config.yaml",
	)
}

func configsConfigYaml() (*asset, error) {
	bytes, err := configsConfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "configs/config.yaml", size: 1193, mode: os.FileMode(438), modTime: time.Unix(1598408506, 0)}
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
	"configs/config.yaml": configsConfigYaml,
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
	"configs": &bintree{nil, map[string]*bintree{
		"config.yaml": &bintree{configsConfigYaml, map[string]*bintree{}},
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
