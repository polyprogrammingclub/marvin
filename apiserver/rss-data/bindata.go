// Code generated by go-bindata.
// sources:
// rss.xml
// DO NOT EDIT!

package rss_data

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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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

var _rssXml = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x91\xc1\x6a\x03\x21\x10\x86\xef\xfb\x14\xe2\x3d\xbb\x69\x4f\x3d\x18\x03\x25\x04\x0a\x7b\x2a\x9b\x07\x30\xeb\x90\x4a\x8c\xbb\x38\x93\x52\x10\xdf\xbd\xe8\x66\xdd\x6e\x29\xf4\x34\xcc\x3f\xdf\xe8\xef\xaf\xd8\x7f\xdd\x2c\xfb\x04\x8f\x66\x70\x3b\xfe\x54\x6f\x39\x03\xd7\x0f\xda\xb8\xcb\x8e\x9f\xba\xe3\xe6\x85\xb3\xbd\xac\x84\x47\x5c\xb0\xe7\x7a\xcb\x65\x25\xfa\x0f\xe5\x1c\x58\x59\x31\xc6\x98\x20\x43\x16\x64\x08\xf5\x11\x40\x77\xa9\x89\x51\x34\x93\x3a\x11\x1a\xb0\xf7\x66\x24\x33\xb8\x99\x3b\x2c\x52\xa2\x7f\x12\xd3\x8e\x35\xee\x3a\xc3\xad\x71\xd7\x44\x65\xed\x31\x56\x48\xaf\x77\x63\xf5\x41\x51\xb9\xbc\x55\x48\xa7\x51\x2b\x02\x9d\xf1\x15\xf3\x30\x4b\x36\xd1\x5d\xd7\x66\x93\x64\x65\x95\x07\x21\x78\xe5\x2e\xc0\xea\x37\x82\x1b\xc6\x18\xc2\xd9\x0e\xfd\x95\x71\x8f\xb8\x31\x04\x37\xce\xea\x18\x33\x9a\xcf\x49\x92\x2c\xed\x3a\x87\xbf\x32\x28\xd4\xaf\x2c\xfe\xcb\xa1\xec\xcd\x79\x9c\xde\xdb\x75\x14\x85\x18\xef\xe7\x39\x8c\x54\x13\x35\x4b\x8b\xef\x66\x31\x1e\x02\x38\x9d\x5e\x9a\x4b\x25\x9a\xf2\xad\xa2\xf1\x88\xf2\x3b\x00\x00\xff\xff\xf8\xae\xce\xc9\x23\x02\x00\x00"

func rssXmlBytes() ([]byte, error) {
	return bindataRead(
		_rssXml,
		"rss.xml",
	)
}

func rssXml() (*asset, error) {
	bytes, err := rssXmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "rss.xml", size: 547, mode: os.FileMode(420), modTime: time.Unix(1480717470, 0)}
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
	"rss.xml": rssXml,
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
	"rss.xml": &bintree{rssXml, map[string]*bintree{}},
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
