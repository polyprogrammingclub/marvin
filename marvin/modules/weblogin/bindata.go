// Code generated by go-bindata.
// sources:
// layout.html.tmpl
// assets/styles.css
// templates/factoid-info.html
// templates/factoid-list.html
// templates/home.html
// templates/logs-index.html
// DO NOT EDIT!

package weblogin

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

var _layoutHtmlTmpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x56\x5b\x73\xda\x38\x14\x7e\xcf\xaf\x50\xb5\x7d\xcb\x62\x0d\xd0\xa6\x6c\xc7\x78\x97\x84\xb4\x4d\x43\x2e\x5b\x72\x99\xed\xdb\xc1\x3e\xb6\x05\xb2\x44\x24\x99\xe0\xf5\xf8\xbf\xef\xf8\x06\x84\xd2\x76\xa7\xbc\x58\x3e\xe7\xf3\x77\x6e\x9f\xce\x90\xe7\x24\xc0\x90\x4b\x24\x54\x40\xa6\x52\x4b\x49\x51\x1c\xb9\xaf\xc6\x37\x67\x77\xff\xdc\x9e\x93\xd8\x26\xc2\x3b\x72\xdb\x07\x42\xe0\x1d\x11\x42\x48\x9e\x93\x99\x50\xfe\x82\x50\xcb\xad\x40\x4a\x1c\x52\x14\x6e\x75\xf6\xf2\x9c\x87\xc4\xb9\x2b\xcf\x45\x91\xe7\xdb\x13\x0a\x83\x45\x71\x05\x7a\xc5\x25\x79\xc4\x19\xb9\x90\x16\x75\x08\x3e\xe6\x39\xca\xa0\x28\x5c\xd6\x32\x10\x94\x41\x99\xc9\xcb\x58\x01\x86\x90\x0a\xdb\x31\x36\x13\x68\xca\xa0\x9d\x06\xe3\x0a\x2e\x17\x44\xa3\x18\xd2\xda\x19\x23\x5a\x4a\x62\x8d\xe1\x90\xc6\xd6\x2e\xcd\x7b\xc6\x12\x58\xfb\x81\x74\x66\x4a\x59\x63\x35\x2c\xcb\x17\x5f\x25\x6c\x63\x60\x7d\xa7\xef\x9c\x30\xdf\x98\xad\xcd\x49\xb8\x74\x7c\x63\x68\x15\xa7\xfe\x71\x69\x31\xd2\xdc\x66\x43\x6a\x62\xe8\x0f\xde\x74\xba\x4f\x83\xe4\xee\xf3\xcd\x68\xba\x1e\xcc\xbb\xa3\xf4\x18\xde\x3e\x8e\x1f\xe4\x2d\xef\x89\xc5\x87\xf0\xf9\xf9\x7c\x04\x83\x78\x3c\x0e\xe6\x5f\xc5\x72\x82\xd1\x3a\x9e\x3f\x5c\x9d\x77\xc3\x68\xfe\x78\xfb\x31\x59\xfc\x6b\xde\x51\xe2\x6b\x65\x8c\xd2\x3c\xe2\x72\x48\x41\x2a\x99\x25\x2a\x35\xb4\x6e\xb7\x6b\x7c\xcd\x97\x96\x18\xed\x6f\xcb\xf1\x03\x39\x37\x8e\x2f\x54\x1a\x84\x02\x34\x56\xb5\xc0\x1c\xd6\x4c\xf0\x99\x61\xf3\xa7\x14\x75\xc6\x7a\x4e\xcf\xe9\x37\x2f\x55\x2d\xf3\x9f\x94\x72\x71\xf2\xe1\xed\xcd\xe5\xf9\xd9\xe4\xc1\x5e\xb2\xd3\xc9\xf1\x80\x4f\x27\xe3\xf3\x4f\xea\x79\x3a\x0a\xef\xd5\xbb\x93\xaf\x93\x3f\x8e\x17\x1f\x47\xd1\xdd\x97\x80\x9f\x66\x17\xd7\x97\x9f\xe1\xe9\xee\xf6\x13\xfb\xfb\xe1\x7a\xda\x7d\x18\xcf\xbe\x5f\x8a\xcb\xea\x32\xbc\x9d\x91\xbd\x9c\x50\xa8\xa4\x35\x4e\xa4\x54\x24\x10\x96\xdc\x54\x25\xf9\xc6\xfc\x19\x42\xc2\x45\x36\x1c\x09\x8c\x34\x66\x70\x3c\x05\x69\x8e\xa7\x67\xf4\x9b\x91\x7b\x3f\x97\x03\x03\x63\xd0\x1a\x56\x7b\xaa\xd9\xb6\xa2\xee\x1c\xd6\xdd\x56\x6f\xa5\x8c\x5b\x8c\xcb\xea\xfb\xe0\xce\x54\x90\x79\x47\x5b\xb8\x84\xd5\x0c\x74\x2b\x4f\x57\xc2\x8a\xf8\x02\x8c\x19\x36\x1e\x52\x3f\x3a\x5c\xae\x50\x1b\x6c\x73\x0e\xf8\x06\xe7\x2b\x69\x81\x4b\xd4\x9d\x50\xa4\x3c\x38\x80\x68\x28\xca\x0c\x50\x37\xfe\x0a\x03\x7b\x88\x99\x06\x19\x6c\x4a\xa7\x5e\x7d\xfd\x5c\x06\x0d\x27\x0b\xf8\x6a\x87\x9e\x07\xc3\x4d\xfe\x9b\x64\x84\x80\xa5\xc1\x36\xed\xf6\x7d\x37\x6a\x2a\x76\xc2\xb6\x40\x09\xab\x1d\x4c\x9e\x93\x67\x6e\xe3\xb6\x89\xaf\xfd\x54\x6b\x94\x96\xbc\x1f\x12\xe7\xba\xfa\xe0\xac\xb1\x54\x7e\x0d\x32\xc2\xd6\x73\x61\x31\x31\xed\x5c\x76\x08\x79\x48\xf0\xa9\x04\x25\xb8\x25\xec\xec\xe1\x5c\xc1\xdb\xe4\xc0\xb7\x7c\x85\xd4\x73\xa1\x69\x48\x9e\x3b\xf7\x5f\x26\x45\x41\xbd\x3c\xaf\x68\xca\x25\x04\x9e\xcb\x04\xf7\xf6\x82\x75\x48\xb9\xbf\x0e\xb1\xff\x2a\xdd\x56\x6a\x4d\x39\xb5\xe5\xe5\xf6\xab\x87\x94\x8a\xff\xd1\xec\xf6\xa8\x79\x14\x5b\xba\x1f\xb0\x6c\x96\xd3\xb4\xf8\xde\xa0\x3e\x58\xc9\x0b\x43\x65\xe4\x49\x44\x62\x2c\x19\x87\xb4\xdf\xa3\xe4\x99\x07\x36\xae\x8f\xd5\x36\xca\xf3\x5d\x52\x67\xb4\x02\x0b\x9a\xf4\x7b\x45\x41\xbf\x21\xab\x7e\x46\xfb\x06\xed\x8f\x3e\x24\xdd\xf5\xef\xe4\xb0\xfb\xe4\x4d\x51\x90\xde\x77\xdd\xdd\xde\xa0\x28\xc8\x9b\x35\x25\xec\x40\x29\xed\x94\x7e\xa3\x9e\x11\xe0\x2f\xfe\xda\x23\xd9\x0e\xec\x65\x5f\x7e\x51\x0c\x4c\x41\x6a\x63\x56\x85\x62\xc6\x82\xb6\xd4\x9b\xa8\x88\x5c\xc8\x1f\x6a\xe2\xe0\xdc\x77\xaf\x69\x7d\x74\x99\x84\x55\xb9\x75\x36\x42\x6a\x2f\xd8\xeb\x48\xa8\x19\x88\xea\x62\xd5\x72\x6a\xf6\x52\xb9\x55\x50\x5a\x4a\x9c\x53\x15\x64\x63\xb0\x50\xad\xb1\xb8\xeb\x5d\x71\x63\xb8\x8c\x48\x83\x78\xe5\xb2\xb8\x5b\x6d\xb4\x7d\x45\xba\xac\x5e\x76\x2e\xab\xff\x12\x6c\x3c\xff\x05\x00\x00\xff\xff\x3e\xd8\xcc\xbe\x49\x08\x00\x00"

func layoutHtmlTmplBytes() ([]byte, error) {
	return bindataRead(
		_layoutHtmlTmpl,
		"layout.html.tmpl",
	)
}

func layoutHtmlTmpl() (*asset, error) {
	bytes, err := layoutHtmlTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "layout.html.tmpl", size: 2121, mode: os.FileMode(420), modTime: time.Unix(1481242895, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsStylesCss = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x53\xd1\x6e\xd3\x30\x14\x7d\xef\x57\x5c\x6d\x0f\x93\xba\xa5\x4d\xbb\x6e\x94\x4c\x4c\x1a\x05\x09\x24\xc4\x03\x9b\x78\x41\x3c\xdc\xd8\x37\x89\x35\xc7\x37\xb3\x9d\x75\x05\xed\xdf\x91\x9d\xb4\x2b\xa5\x9d\x78\x8c\x7d\xcf\xb9\xc7\xe7\x9c\x8c\x87\xf0\x89\x50\x92\x85\xe1\x78\x90\xb3\x5c\xc1\xef\x01\x00\x40\x8e\xe2\xbe\xb4\xdc\x1a\x99\x08\xd6\x6c\x33\x38\x2e\xa8\x10\x85\xb8\x1a\x3c\x0f\x46\x06\x1f\x73\xb4\x89\xc1\xc7\x6b\xad\xae\x55\x5d\xfe\xa8\x48\x95\x95\x7f\x77\x74\x3e\x3d\xfa\xd9\x53\xd4\x68\x4b\x65\x12\x1b\x2e\x32\xb8\x68\x9e\xae\xe2\xb1\x54\xae\xd1\xb8\xca\x40\x19\xad\x0c\x25\xb9\x66\x71\xff\x3f\xac\xa7\x80\x3d\xf3\x01\x8a\x70\xd5\xa0\x94\xca\x94\x89\xa6\xc2\x67\x90\x5e\x6d\x2b\xe9\xce\xa2\x90\xe7\xc1\x78\x08\x77\xdc\x00\x17\x20\xd8\x78\x32\x3e\x18\x30\x6a\xb0\xa4\xa4\xea\x0c\xa9\x26\xaf\xaf\x7b\xfe\x7b\xdc\xd5\xa8\x75\x8f\x28\xd8\xf8\xc4\xa9\x5f\x94\xc1\x64\xbe\x7e\xf8\xda\xc7\xf9\x7c\xbe\x47\xd6\x66\x2c\x62\x0b\xac\x95\x5e\x65\x70\x72\xa3\xa9\xb4\xb4\x42\xb8\x45\xe3\xe0\x76\x71\x72\x06\x0e\x8d\x4b\x1c\x59\x55\x04\x0d\xe1\x21\xb7\x24\xbc\x62\x93\xc1\x02\xb5\xe6\xd6\xbb\x40\xf4\x9e\xd9\x3b\x6f\xb1\x81\x0f\x2c\xe2\xc9\x82\x9b\x55\x4c\x03\xa6\xe9\x64\x92\x4c\xd3\xc9\x25\xdc\x2d\x95\xf7\x64\xcf\xe0\xb3\x11\xa3\x30\xf4\x45\x09\x32\x8e\x32\x58\x2c\xe0\xc6\x7b\xab\xf2\x36\x70\xc3\xf9\x28\x0d\x0e\x0d\xc6\xc3\x01\x0c\xb7\x16\xc5\xcf\xaf\xec\xe1\xa1\x55\x9e\x00\x35\x59\xef\xce\x20\x6f\x3d\x88\xd6\x79\xae\x01\x8d\x84\x8a\x74\x53\xb4\x1a\x0c\x7b\x72\x50\xb0\x85\x82\xf5\xbd\x03\x4b\x18\xe2\x02\x5f\x11\x48\x16\x6e\x14\xd8\xbe\xd1\x43\xab\x2c\x39\x40\xc8\xd1\x51\x24\xa8\x59\xaa\x42\x91\x05\xa1\xd1\xc5\xb1\xa8\x05\x16\x5c\xd7\x6c\xc0\xf9\x95\xee\x89\x43\x0a\x7e\xd5\x90\x8b\x81\xe6\x2e\x11\x9d\xd6\x18\x4d\x5f\x8f\x0c\xa6\x69\xe7\x77\x97\x41\xf7\xdd\xd5\x25\x67\x2b\xc9\x66\x30\x69\x9e\xc0\xb1\x56\x12\x8e\x89\xe8\xe5\x26\xe6\x95\x2c\x95\xf4\xd5\xa6\xd5\xfd\x8d\x45\xa9\x5a\x97\xc1\x79\x57\xb1\xed\xe5\xd5\x2c\xee\xef\x23\xf7\xdc\xf4\xe5\xec\x0f\x72\xf6\x9e\xeb\x4d\x39\xb7\x91\x4d\xa6\xd1\xf9\x44\x54\x4a\xcb\x6d\x8e\x35\x24\xdd\x05\x08\x96\x14\x07\xf7\xab\x8a\xcd\x0f\x2d\x20\x03\x6d\x03\xae\x41\x41\x90\x93\x5f\x12\x19\xa8\x5b\xed\x55\xa3\x09\x7a\xb2\x7f\x4c\x3c\x85\x5d\x4b\xb7\x9f\x94\x5c\xbc\xec\xf8\x8e\x56\x61\xa8\xce\x2e\x47\x22\xd1\x94\x64\xb7\x25\x46\x4b\xd7\xff\x87\xa0\xd9\x7c\x36\xdb\x79\xd5\x1a\xd4\xfb\xf8\xfa\xec\x12\xad\x09\xad\x3a\xb4\x01\xf1\xf2\x4d\x3a\x3f\x84\xda\x59\xb1\x7f\x58\x99\x82\x0f\xf2\x4f\xf2\x79\xfa\x96\xf6\x42\x76\xc8\x5f\x26\xc7\x43\xf8\x68\xe4\xe6\xc7\x82\xe1\xf8\x4f\x00\x00\x00\xff\xff\xe4\x31\x71\x04\x9b\x05\x00\x00"

func assetsStylesCssBytes() ([]byte, error) {
	return bindataRead(
		_assetsStylesCss,
		"assets/styles.css",
	)
}

func assetsStylesCss() (*asset, error) {
	bytes, err := assetsStylesCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/styles.css", size: 1435, mode: os.FileMode(420), modTime: time.Unix(1481243193, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesFactoidInfoHtml = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x8f\xcd\x6a\x2b\x31\x0c\x85\xf7\x7e\x0a\x31\xfb\x9b\x21\xdb\x8b\xeb\x4d\xa1\x50\x28\x5d\xf4\xe7\x01\x84\x7d\xd2\x31\x9d\xd8\xc3\xc8\x4d\x09\x46\xef\x5e\xe2\x78\x52\xba\x3b\xfa\xfc\x49\x96\x6a\x0d\x38\xc4\x04\x1a\xa4\x9c\x67\xc8\xa0\x6a\x6c\x8b\xce\x18\x3b\xf6\x54\x2b\x52\x50\x35\xbf\xb6\xcf\xa9\x20\x95\xa6\x87\x78\x22\x3f\xb3\xc8\x5d\xc3\x1c\x13\xd6\xc1\xfd\xe1\x0b\x7f\xe0\xdf\x04\x0e\xed\x85\x88\xc8\x4e\x7b\xf7\xc0\xbe\xe4\x18\xfe\x53\xad\xbb\x9e\x9f\xf9\x08\x55\x3b\x4e\x7b\x67\xec\x18\xe2\xc9\x99\xab\xbe\xb8\xa7\xec\x3f\xd1\xe4\x78\xa0\xdd\xa3\x5c\x6b\xd5\x33\xa4\x56\xcc\x02\xd5\x94\xfb\xa6\x76\x5c\xdc\xad\x8f\xa5\x10\x42\x2c\xed\x9f\x4b\xf5\x2e\x58\x55\x29\xa6\x0d\xdc\x4f\x9c\x12\x66\x55\xe2\xb2\xb1\xb7\x78\x84\x14\x3e\x2e\x7d\x5a\x1f\xb7\x62\xbb\x49\xf2\xd7\xea\x31\x38\xeb\x73\x80\xab\x75\xf7\xc2\xdf\xaf\x8d\x5d\x3a\x1a\xb4\xe3\xb2\xe2\x76\x48\xdf\xed\x27\x00\x00\xff\xff\x0f\x6b\x9d\x95\x73\x01\x00\x00"

func templatesFactoidInfoHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesFactoidInfoHtml,
		"templates/factoid-info.html",
	)
}

func templatesFactoidInfoHtml() (*asset, error) {
	bytes, err := templatesFactoidInfoHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/factoid-info.html", size: 371, mode: os.FileMode(420), modTime: time.Unix(1481185646, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesFactoidListHtml = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x55\x4f\x6f\xdc\xb6\x13\xbd\xeb\x53\x4c\xf6\x10\xfc\x7e\xc0\x4a\x42\xfe\x9c\x5c\x46\x68\x10\xa0\x80\x01\xa7\x87\x3a\x46\xd1\x53\x3d\x2b\x8d\x56\xac\xf9\x47\xe5\x50\xbb\x16\x04\x7e\xf7\x82\x5c\x6a\xbd\x5b\x27\x97\xde\xa4\x21\x67\xde\x9b\x37\x4f\xa3\x65\xe9\xa8\x97\x86\x60\xc3\x7e\x56\xc4\x9b\x10\x0a\x91\x1e\x9b\x62\x52\x55\x8f\xad\xb7\xb2\x2b\x95\x64\x0f\x4b\x51\x84\x57\x41\x25\x61\x29\x00\x00\xe2\x5b\x99\x32\x6f\xc0\x58\x43\x3f\xa5\xa8\x46\xb7\x97\xa6\x54\xd4\xfb\x1b\x78\x3f\x3e\x5f\x45\x9d\xdc\x0f\x6b\x38\x14\xaf\xea\x56\x0a\xd9\x97\xda\x76\xb2\x97\xd4\x65\x98\x5e\x59\xf4\x37\x90\x52\x63\x96\xa8\x33\xdb\x65\x21\xd3\x85\x50\xbc\x74\xd4\x5a\xe3\xc9\xf8\xd4\x52\x27\x0f\xd0\x2a\x64\xfe\x94\xc2\x28\x0d\xb9\x4d\x73\x15\x1f\x71\x4f\xe5\x40\xd8\xa5\x93\x88\x25\x86\x77\xcd\x2f\x27\x52\x2c\xea\xe1\x5d\x23\x58\xa3\x52\xcd\x67\x38\xe2\x0c\xde\x02\x7b\xeb\x08\xd0\x74\xe0\xc8\x3b\x49\x07\x02\x4f\xcf\x5e\xd4\xa7\x7b\x85\xa8\x3b\x79\xb8\x46\xd9\x71\xd9\xa2\x52\x76\xf2\xf0\xf2\x58\x4a\xd3\xdb\x33\xe8\xc7\xe6\x81\xa5\xd9\xc3\x05\xf4\xc7\x7c\x36\x36\xf7\x38\x83\x68\x6d\x47\xcd\x9b\xac\x97\x41\x4d\xa2\x4e\x21\x90\x06\xd0\xcc\xd0\x0e\x68\x0c\x29\x38\x0e\xe4\x08\xbe\xa2\x3b\x48\x03\x92\x23\x65\x37\x19\x40\xc8\xa9\x15\xdc\xf6\xe0\x07\x5a\xdf\xc1\xe3\x13\x31\xa0\xdb\x4f\x9a\x8c\xe7\x6d\x42\x1d\x9d\x3d\xc8\x8e\xe2\x45\x0d\xd8\x7b\x72\x29\x27\xe2\x6e\x81\x47\x6c\xa9\x64\x1a\xd1\xa1\xa7\xae\x82\x6f\x03\xad\x04\x57\x56\x82\x74\xa3\x27\xf6\xa2\x26\xdd\xc0\x8e\x00\x7d\xaa\x70\x20\x37\xc3\x8e\xf6\xd2\x98\xd8\xaf\x4d\x5c\x4e\x06\x21\x66\xdc\x53\xac\x26\x39\x32\x47\xc5\x16\xa2\x5a\xd4\x01\xc2\x0e\xcd\xbe\x5c\x7b\x10\xf5\x78\x56\xe7\x0f\x3b\x41\x8b\xe6\x74\xfd\xaa\x55\xd8\xcd\xc0\x38\x47\x9c\x13\xbb\x9f\xf5\x49\x96\xf5\x7c\x4f\x1e\x5e\x2b\x5a\xc1\x67\x86\x1d\xf5\xd6\xd1\xf6\x2c\xc4\x59\x9f\x04\x8b\x0c\xc6\x3a\x8d\xea\xc4\xe4\xbf\xce\xfc\x77\x27\xfd\x0f\xa7\xfe\xc0\x49\xfe\x7f\x51\x77\xa4\x49\xef\xc8\x5d\xf2\x86\xec\xfa\x55\xfb\xd6\x6a\x1d\x0d\x1a\xdd\x4a\xfe\x62\xf4\x97\xb2\xdd\xf6\x30\xdb\x09\x0c\x51\x77\x65\x87\x5c\x2b\x26\xe7\xaf\x06\x0c\x1d\x95\x34\xc4\x5b\xe8\x2c\xf8\x41\xf2\xcd\xa9\xeb\xd1\x51\xf3\x03\x76\xf4\x8c\x7a\x54\x04\x6f\xdf\x15\x8f\x8f\x8f\xc5\xd7\x49\x79\x19\x6b\x14\x5f\x4e\xe5\x63\x34\xd3\x15\x75\xac\xf3\x3d\x5a\x3a\x66\xc5\x2a\x7a\x4d\x87\x11\x9d\xe7\x2d\xfc\x35\xb1\x07\xec\xa2\x2f\xde\xbe\xcf\x8e\x39\x5a\xf7\xc4\xd0\x5b\x07\xa8\xd4\xaa\x01\xc7\xcf\x23\x7f\x0c\x25\xec\x26\x0f\xc6\xfa\x78\x29\xe1\x5d\x7a\x8a\xb7\x71\xaa\x7e\xa0\x19\x50\x1d\x71\x66\x98\xf2\x00\xc8\x78\xe9\x68\xf5\xe7\x9b\xcb\x89\x17\x62\xf8\xb0\xae\x0b\xb8\x93\xd1\xee\xc3\x87\xa6\x10\x93\x5a\x7d\x70\xb9\xe0\x36\x71\x63\x81\x43\xb3\x27\xa8\xe2\x6d\x08\xa1\x00\x10\x4a\x36\x02\x61\x70\xd4\x7f\xda\xd4\x2b\x9d\x7a\x59\x64\x0f\xf4\x37\x54\xf7\xad\x1d\xe9\x4b\xfe\xbc\x37\x9b\x10\xfe\x5c\x16\x52\x4c\x21\x2c\xcb\xd5\x61\x0c\xa4\x85\x58\x2f\x4b\x95\x59\xfd\x8a\x9a\x42\xc8\x9e\x83\x6c\xa6\x65\x81\xcb\x73\x08\x21\xcf\x22\x61\x1a\xfa\x0e\xa6\xe0\x11\xcd\xda\x94\xe4\x32\xaf\x9b\xd2\x1a\x35\x6f\x9a\x57\x44\x44\x1d\xef\x37\x99\x8f\xa8\x71\x25\x90\x10\xaa\x5b\xbe\xb3\xed\x13\x75\x21\xfc\x4f\xa5\x87\xff\xaf\xab\x3c\xd3\xbc\x44\xbb\xfa\x25\x6c\x9a\xf8\x0a\xe7\x3f\xc4\x6e\x86\x65\xa9\xee\x90\xfd\x03\x93\x0b\x21\xee\x99\x1c\xf8\x26\x35\xb1\x47\x3d\x9e\xe9\xac\xd5\x47\x47\x6b\x71\xb6\x93\x6b\x69\xd3\xac\xc2\x54\xbf\xe1\xf1\x3e\xc5\xce\xa2\x9c\x0d\x2a\x6a\x25\xd3\x08\xc9\x74\x71\x74\xa2\x9e\x5e\x16\x7e\x6e\xe0\x9f\x00\x00\x00\xff\xff\xd1\x7d\x15\x39\x5d\x07\x00\x00"

func templatesFactoidListHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesFactoidListHtml,
		"templates/factoid-list.html",
	)
}

func templatesFactoidListHtml() (*asset, error) {
	bytes, err := templatesFactoidListHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/factoid-list.html", size: 1885, mode: os.FileMode(420), modTime: time.Unix(1481242988, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHomeHtml = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x74\x53\x5f\x6b\xdb\x30\x10\x7f\xef\xa7\x38\xfc\x9c\x45\x0b\xf4\x69\xb8\x82\x32\x18\x2b\xcc\x50\x58\xd9\x6b\xb9\xc8\x67\xeb\x16\x59\x32\xd2\xd9\x59\x16\xf2\xdd\x87\x6c\x27\x69\x3b\xf2\x26\xee\xcf\xef\xcf\xdd\xe9\x78\xac\xa9\x61\x4f\x50\x98\xe0\x85\xbc\x14\xa7\xd3\x5d\x59\xf3\x08\xc6\x61\x4a\x0f\x53\x18\xd9\x53\x2c\xf4\x1d\x00\xc0\xdb\x5c\x8f\x2d\x7d\xb2\x84\xf5\x25\x3b\x55\xd8\x8d\xfe\x1e\x3a\x2a\x95\xdd\xe8\x32\x75\xe8\x9c\xe6\x04\x7b\x4b\x91\x40\x2c\x81\x25\x8c\x02\x9c\x4a\x35\x27\x67\x5c\x55\xf3\xa8\xef\xe6\x77\xaf\x5f\x2c\x27\xe0\x34\xd5\xef\x69\x0b\xec\x85\x62\x83\x86\xa0\x09\x11\x2a\x8c\x23\xfb\x75\xa9\xfa\xa5\xd9\xde\xeb\x12\xc1\x46\x6a\x1e\x0a\xd5\xa0\x91\xc0\x75\x2a\xf4\xb7\xe5\x55\x2a\xd4\xa5\xb2\xf7\xfa\x0c\xff\x8b\x69\x0f\x29\x0c\xd1\xd0\x0a\x2c\x27\x09\xf1\xb0\x02\xf4\x35\xe0\x50\xb3\x80\x0b\x6d\x9a\x98\xce\x58\xb7\xb8\x72\x61\xa1\x7f\x84\xf6\x06\x47\x70\xf5\x0c\x16\x9a\xc9\xcb\x4f\x87\x66\x07\xc6\xa2\xf7\xe4\x12\xf4\x98\x64\x8a\x6f\x3e\xef\xa0\xa3\x94\xb0\x25\x70\xdc\xb1\xac\xcb\x6d\xbc\xce\xf4\x25\xc0\x98\xe1\xfa\xc8\x23\x0a\x5d\x00\x56\x70\x08\x03\x74\x43\x12\xd8\x12\x5c\x65\x05\x1c\xc4\xaa\x94\xc9\x54\x12\x8c\x52\xe8\xc4\xad\xa7\x1a\xd8\xc3\x9e\xc5\xce\x42\xb2\xe6\x5b\xce\xd8\x8f\x2c\x94\x0a\xfd\xbc\x70\x7e\x9d\x39\xe1\x29\x27\x50\x38\xf8\xff\x3d\x57\xe8\xb3\x83\x79\x3f\xb0\x40\x80\x84\x2c\x33\x7e\x54\xbf\x82\x10\x41\x70\x47\x80\xd0\x0f\x5b\xc7\xc6\x1d\x00\x47\x64\x87\x5b\x47\x4b\x77\x6e\xc6\x8f\x9d\xef\x35\x57\x21\x4e\x65\x66\x3e\xba\xb3\x9a\xc1\xbd\xb9\x49\xc7\x79\x49\x17\xf7\x4f\x5e\x22\xe6\x9e\x96\x64\xd9\xc9\x73\x35\xdd\xa8\x9f\x26\x6a\x71\xcc\xaa\x4c\x88\x91\x4c\xb6\x5a\x2a\xc7\xef\xe1\x2e\x93\xb2\x22\xfd\x17\xa5\xf6\xbc\xe3\x35\x39\x32\x12\x22\x76\xb8\x36\xa1\x9b\x62\xaa\xc2\x3f\xdc\xf1\x5f\x7a\x7d\x6c\x1a\x8e\x1d\xd5\xaf\x15\xfe\x0e\x91\x85\xf3\x74\xab\xc7\x2a\x4f\x11\xfa\xe0\x1c\xfb\xf6\xca\x53\xaa\x6c\x60\xf9\x16\xc7\x23\xf9\xfa\x74\xfa\x17\x00\x00\xff\xff\x30\xb9\x6f\x2f\xae\x03\x00\x00"

func templatesHomeHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesHomeHtml,
		"templates/home.html",
	)
}

func templatesHomeHtml() (*asset, error) {
	bytes, err := templatesHomeHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/home.html", size: 942, mode: os.FileMode(420), modTime: time.Unix(1481242970, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesLogsIndexHtml = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x53\x5b\x6b\xdb\x4c\x10\x7d\xf7\xaf\x18\x1c\x3e\xbe\x27\x49\xb6\x13\x52\x90\x15\x51\xda\x3e\xb4\xd0\x94\x40\xa1\xd0\xc7\x91\x76\x24\x6d\xb3\xda\x15\xbb\x2b\xd9\x66\xd1\x7f\x2f\xab\x4b\xe2\x38\x29\x26\xbd\xd8\x60\xac\xb9\x9c\x33\x73\xe6\xc8\x39\x46\x05\x97\x04\x4b\x63\x0f\x82\xcc\xb2\xef\x17\xc9\xf0\x37\x5d\x84\x79\x85\x52\x92\x08\x32\xb5\x0f\x72\x25\x2d\x72\x49\x1a\xdc\x02\x00\x80\x71\xd3\x08\x3c\xc4\x50\x08\xda\x6f\x87\xdf\xa0\x10\x6a\x17\x6b\xb5\x83\x9d\xc6\x66\xfb\xa3\x35\x96\x17\x87\xa1\x93\xa4\x8d\xc1\x34\x98\x53\x90\x91\xdd\x11\xc9\xed\xa2\xff\x15\x41\x0a\x8c\x77\x13\x4d\x8d\xba\xe4\x32\x86\xcb\x66\x7f\xb6\xe3\xb8\x2f\xc3\xfc\xbe\xd4\xaa\x95\x2c\xc8\x95\x50\x3a\x86\x8b\x62\xe5\xbf\xdb\x31\xad\x34\x23\x1d\x68\x64\xbc\x35\x31\x5c\x7b\x74\x1f\x6f\x90\x31\x2e\xcb\x18\x56\xe1\xba\xd9\x8f\xf1\x33\xac\x78\x2a\x48\x26\x54\x7e\x3f\xc2\xed\x38\xb3\x55\x0c\xeb\xd5\xea\xbf\x31\x30\xcf\x72\xbd\xb9\xbc\xcc\xde\x6c\x8f\x56\x0c\x34\x2f\x2b\xeb\x6b\xcf\x73\x86\x52\x05\x35\xea\x8e\xcb\x89\x5b\x35\x98\x73\x7b\xf0\x63\x5f\x3f\x25\xda\x6c\x36\xaf\x40\x9b\x77\xf9\xcd\xee\xb8\xe3\x86\x5b\x62\x2f\xa3\x24\xd1\xe4\x2b\xe7\x48\xb2\xbe\x5f\x3c\x7a\x6f\xb2\xc8\x60\x3e\xaf\x69\x2e\xd0\x98\x9b\xe5\x03\xd5\x32\x1d\x00\x8f\x73\x0d\x96\x14\x54\x84\xcc\x67\x93\x6a\x9d\xbe\x1f\x67\x84\xcf\xaa\x34\x49\x54\xad\xd3\xc4\xd4\x28\x44\xfa\x8e\x72\x6c\x0d\xc1\x57\x81\xf9\x3d\x34\x9a\x6a\xde\xd6\xc0\x0d\x64\x5c\x08\x62\xd0\x90\x86\xd6\x90\x4e\xa2\xb1\x3e\x89\x18\xef\xd2\xc5\x48\x58\x6d\xd2\xbb\x36\x13\x3c\x87\x09\xde\x43\x6f\x9e\x4f\xf3\xa2\x40\xd3\xd4\xfe\xe3\x9c\x46\x59\x12\x84\x33\x4c\xdf\x3f\xe4\x06\x9c\x41\x9a\x9b\xe5\x60\xcb\x18\x9c\x13\x24\x21\xbc\xa5\x3a\x23\x6d\xfa\xfe\x08\x69\xee\x38\x89\x54\x57\x69\x82\x50\x69\x2a\x6e\x96\x91\x50\xa5\x89\x9c\x0b\x3f\x7d\xf0\xad\x17\xce\x85\x5f\xb0\xa6\xbe\x4f\x22\x4c\x93\xa8\xba\x3a\xe9\x6d\x52\xe7\xc2\xbb\x56\x37\xca\x50\xf8\x0d\x45\x3b\x94\x36\x27\x55\xd1\x13\xd2\x93\xc7\xf9\xa4\x27\x29\xe7\x80\x17\x10\x7e\x57\xad\xbe\xd3\xbc\x43\x4b\xf3\xfe\x10\xcc\xd5\x5e\xe2\x31\xf7\xc7\x1a\x07\x30\xa9\xfc\x12\xe1\x2b\x15\x9f\x59\x9d\xe3\x05\x48\x65\x21\xfc\x88\xe6\x76\x30\x7a\xdf\x1f\xbd\x7f\xe3\xe2\xff\xfa\x40\x09\xd5\xa9\x37\x36\xa0\xa6\x61\x1a\xec\x90\x0b\xcc\x04\x41\xa1\x34\xd8\x8a\x1b\x98\xf4\x81\x6c\x32\x3c\xb7\xc0\x14\x19\xf9\xbf\x85\x0a\x3b\x82\xb7\xe3\xc8\x61\x12\x51\x9d\x3e\xbf\xef\x5f\x71\x01\x90\x64\x8f\xa7\x3d\x36\x42\x30\xa4\xfc\x0b\x3e\x46\x27\xe1\x7e\x06\x00\x00\xff\xff\x08\xf5\xb3\xc6\x80\x06\x00\x00"

func templatesLogsIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesLogsIndexHtml,
		"templates/logs-index.html",
	)
}

func templatesLogsIndexHtml() (*asset, error) {
	bytes, err := templatesLogsIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/logs-index.html", size: 1664, mode: os.FileMode(420), modTime: time.Unix(1481243255, 0)}
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
	"layout.html.tmpl":            layoutHtmlTmpl,
	"assets/styles.css":           assetsStylesCss,
	"templates/factoid-info.html": templatesFactoidInfoHtml,
	"templates/factoid-list.html": templatesFactoidListHtml,
	"templates/home.html":         templatesHomeHtml,
	"templates/logs-index.html":   templatesLogsIndexHtml,
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
	"assets": &bintree{nil, map[string]*bintree{
		"styles.css": &bintree{assetsStylesCss, map[string]*bintree{}},
	}},
	"layout.html.tmpl": &bintree{layoutHtmlTmpl, map[string]*bintree{}},
	"templates": &bintree{nil, map[string]*bintree{
		"factoid-info.html": &bintree{templatesFactoidInfoHtml, map[string]*bintree{}},
		"factoid-list.html": &bintree{templatesFactoidListHtml, map[string]*bintree{}},
		"home.html":         &bintree{templatesHomeHtml, map[string]*bintree{}},
		"logs-index.html":   &bintree{templatesLogsIndexHtml, map[string]*bintree{}},
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