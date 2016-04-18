// Code generated by go-bindata.
// sources:
// provider/aws/templates/service/syslog.tmpl
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

var _templatesServiceSyslogTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x58\x5f\x4f\xe3\x38\x10\x7f\xef\xa7\x88\xac\x95\x56\xda\x2b\x29\x05\xad\x4e\x67\xe9\x1e\x7a\xb0\xdd\xe3\x8e\xbd\xad\x5a\x76\xf7\xe1\xc4\x83\x71\xdd\x92\x23\x89\x23\xdb\x81\x05\xd4\xef\x7e\x63\xe7\x4f\x63\xc7\x49\x0b\xa2\x05\x15\xc5\x1e\xcf\x9f\xdf\x8c\x7f\x33\xe1\xf9\x39\x58\xb2\x55\x94\xb2\x00\x49\x26\xee\x23\xca\x50\xb0\xd9\x0c\x9e\x07\x41\x80\x26\x3f\x16\x57\x2c\xc9\x62\xa2\xd8\x94\x8b\x84\xa8\xef\x4c\xc8\x88\xa7\x08\x07\xe8\xe4\x78\x7c\x7c\x74\xfc\x1b\xfc\xa2\xa1\x96\xfd\x9a\xab\x2c\x57\x12\xb6\xf4\x51\x58\xf8\x26\xe2\xfa\x01\x1e\xbf\x93\x38\x67\x8d\x05\x58\x9a\xb3\x95\xd6\xa4\x05\xcb\xc5\xcd\xa0\xfa\xde\x18\xa5\x33\x22\x48\xc2\x14\x58\xed\xd4\x7b\xce\x24\x15\x51\xa6\x4a\xb7\x16\x8f\x32\xe6\xeb\xe0\xdb\xfc\x72\x18\xb0\x70\x1d\x06\xef\x15\xcd\x7e\x51\xb1\xc4\xa3\x11\x6c\xc8\x71\x98\x91\x8c\x09\x25\x48\x14\x93\x2c\x0b\x29\x4f\xf0\x78\x7c\x72\xfa\xf1\xbd\x09\xc3\xa8\xbc\x7a\xcc\x98\xd1\xa5\x44\x94\xae\x91\xe3\xd3\x9c\x49\x9e\x0b\xca\xb6\x2e\x01\x84\x82\xa4\x6b\x16\xbc\xbb\x1b\x06\xef\x40\x6b\x80\x7f\x0f\x42\xf8\x2b\x35\x92\xa5\x52\x10\xca\x33\xb0\x6c\x04\xc2\x7f\x20\x2c\xd8\x9c\x31\x91\x44\xb2\x84\xb4\x81\xcc\x4c\x70\xed\x64\xd4\x30\x52\xee\x4c\x68\x15\x69\x4c\x92\x9b\x25\xc1\x17\xe9\x3d\xbf\x63\xd3\x3c\x2d\x36\x86\x4d\xe1\x6a\x55\x5b\x73\x14\x6d\xe1\xaf\x8f\x36\x76\x37\x96\x9a\x19\xc0\x40\xa3\x8c\xc4\x6d\x1d\xd3\x14\xe3\xbf\x78\xa4\x1d\xfa\xd7\xda\x81\xbd\xd0\x72\x46\x7f\x5c\x11\x10\xd2\x39\x69\xc9\x05\x8e\x19\xdb\x61\xa8\x4a\x8c\xe7\x6c\xed\xf8\xec\xf1\xbc\x3c\x47\x12\xf2\xc4\x53\xf2\x20\x75\xb6\xdd\x23\xd7\x83\xae\x27\x1b\x84\x85\x49\xfa\x84\x52\x9e\xa7\xaa\x13\x4c\xe3\x5b\x29\x74\xb1\xec\x86\xb4\xd4\x26\xd2\x97\x41\xba\x0f\xa2\x44\xa4\x18\x62\xc5\x1a\x59\x7c\x68\x68\x5f\x65\xc0\x8b\x4f\x8f\x0d\x88\xe4\x68\x2d\x78\x9e\xbd\xd0\x1a\x5c\x39\x73\xd9\x4a\x6a\x0a\x2f\xf9\xfa\xb3\x56\xb3\xbd\x95\x3b\x0c\x7f\x78\x41\xb1\x0c\x3c\x9a\x6a\x26\x31\x61\x5f\x16\x17\x16\x37\x2e\xfd\xc0\x39\xe2\x67\x89\x45\x7e\x53\x73\xdc\x34\x8a\x81\x0e\x6d\xb6\x38\x67\x19\x4b\x97\xf2\xab\x5b\x32\xa8\x6d\x09\x9c\x1e\xee\x43\x33\xc0\xaa\x2a\x4a\x89\xb6\xd8\x59\xa4\x9f\x99\x9a\x28\xe5\x2b\x53\x2f\x1b\x99\x1d\xad\x6c\xcf\x0b\x57\x44\x3a\x23\x0a\xbe\x0d\xe1\xd9\xdc\x56\x25\xb3\x8f\xdb\xf6\x2b\x80\x7d\x52\xa7\xef\x12\xf6\xe4\xc1\x4e\x20\xd8\x83\x44\x54\xda\xb7\x30\x34\xda\x55\x07\xe4\xe8\x8c\x2f\xdd\x30\xd0\xe2\xf4\x8f\x9c\xde\x31\x0f\xdb\xf4\x71\xc4\xd1\x3e\x24\x41\x39\xb4\x8d\x9f\x6f\xca\x0e\xaf\xe5\xd5\xd3\xbf\xd9\xe3\xb6\x9f\x8d\xa4\x69\xe0\xe1\x53\x94\x21\x6f\x5e\xec\x7e\x6f\xe1\xd5\x9e\x27\xec\xa3\x3d\xed\xd0\x0a\x71\xa1\x08\xbd\x33\x42\x5e\x35\x7f\x92\x74\x19\x9b\x3b\x88\xa2\x74\xc9\x7e\x86\xb7\xe5\x42\x43\x66\xce\xe3\x96\x89\x9e\x3b\x53\xc8\xdb\xd9\x70\xef\xca\xb5\xd7\x99\x39\x10\x69\x64\xc2\x41\x29\x94\xd0\x7f\xcd\x66\x8a\xae\x60\x87\xe7\xca\x8c\x6b\x1f\xdb\x64\xe3\x65\x27\x7b\x1c\x28\x85\xdd\x78\x3a\xcb\x78\x22\x65\x9e\x30\x2d\x3d\xe3\x71\x44\x1f\xcf\x39\x85\xe7\x56\xc3\x84\xd1\x0a\x66\xca\x72\xc3\x86\xc2\x2d\xc0\xed\xc4\xe3\x29\x63\xa9\x24\xde\x9a\x6c\xb1\x75\x8b\x7e\x3e\xad\x56\x8c\x1a\x3c\x26\x71\xcc\x1f\xda\xfc\xd4\x3d\xea\x14\x6e\x97\x13\xb2\xcf\x99\xa0\x2a\xe1\xb0\x77\xde\x70\x6f\x86\x7b\x73\x9a\x4f\x56\x00\xc8\x1e\xbe\x4f\x8e\x60\xfe\x1e\xff\xea\x2f\x51\x60\xcd\x5b\x2d\x37\x6a\x56\x83\xc9\x48\x91\xb1\xa6\xf7\x0e\xb5\xf4\xe6\xad\x00\xa1\x33\x77\x6d\x75\xe5\x91\x9e\x1c\x9a\x7d\x1a\xf3\x7c\xf9\x40\x14\xbd\xc5\xb3\x5c\x7d\x61\x30\x78\xd3\x73\xa2\x88\x87\xa0\x8c\xbc\x7f\xf6\xf5\xc8\xb6\x4a\xc0\x1c\xdf\x55\x06\x46\xa8\x9a\xf3\xb5\x58\x6b\x10\xf0\x0e\x0c\xaf\x0a\xdd\x8c\x69\x67\x82\x01\xa6\x55\x7b\xea\x8c\xda\x12\x85\xb7\x13\x46\x92\x5e\x59\xc0\x12\x04\x3f\xdd\x43\xae\xe4\x01\xd1\xb1\x26\xce\x0f\xfa\xe7\x70\x70\x99\x4a\x59\x99\x97\x51\x90\xc3\x45\x37\xb8\x61\x86\xb2\x0f\x18\xa3\xcf\xd7\xfe\x3e\x5c\xcb\x74\x64\xc8\xd7\x96\xeb\x33\x15\xa2\x4e\xb4\x9d\xaa\xfc\x60\x36\x02\xd9\xdd\xbf\xab\x8f\x27\x51\xb5\x9e\x37\x71\xa0\x67\xfa\xdf\xcb\x07\xa9\x53\x3d\x7a\x0b\x4f\x3c\x7d\xfe\x25\x9e\x8c\x3c\x65\x5e\x7c\x5c\x86\xef\x5e\xf5\xbc\x88\xec\x6c\x62\x3b\xfa\x80\xc7\xed\x92\xd4\xcb\xb9\x07\x15\xad\xfe\x4c\xd7\xd6\xb4\xaa\x2d\xf3\xf4\x43\x33\x30\xf2\x8e\xc6\xd7\xfd\xd3\xc3\xc5\xe4\x0b\x94\x56\xdd\x83\xcd\x7f\x4b\x06\x9b\x41\x3d\x10\xff\x1f\x00\x00\xff\xff\x5a\x47\x01\x98\x66\x12\x00\x00")

func templatesServiceSyslogTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesServiceSyslogTmpl,
		"templates/service/syslog.tmpl",
	)
}

func templatesServiceSyslogTmpl() (*asset, error) {
	bytes, err := templatesServiceSyslogTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/service/syslog.tmpl", size: 4710, mode: os.FileMode(420), modTime: time.Unix(1461005311, 0)}
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
	"templates/service/syslog.tmpl": templatesServiceSyslogTmpl,
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
	"templates": &bintree{nil, map[string]*bintree{
		"service": &bintree{nil, map[string]*bintree{
			"syslog.tmpl": &bintree{templatesServiceSyslogTmpl, map[string]*bintree{}},
		}},
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

