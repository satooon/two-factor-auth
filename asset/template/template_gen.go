// Code generated by go-bindata.
// sources:
// layout/footer.tmpl
// layout/header.tmpl
// layout/layout.tmpl
// view/admin/index.tmpl
// view/admin/logged_in.tmpl
// view/admin/two_factor_auth_key.tmpl
// view/admin/user.tmpl
// view/auth/index.tmpl
// view/auth/logged_in.tmpl
// view/auth/two_factor_auth.tmpl
// view/index.tmpl
// DO NOT EDIT!

package template

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

var _layoutFooterTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xce\x41\x6a\xc5\x20\x10\x06\xe0\xbd\xa7\x18\xdc\x37\xb6\xdb\xd6\xe7\xae\x27\xe8\x09\xc4\x99\xf7\x10\x82\x09\x8e\x0d\x09\xc3\xdc\xbd\x44\x9b\x42\x17\x85\xba\x10\x94\xff\xff\x66\x44\x00\xe9\x9e\x0b\x81\xbd\x2f\x4b\xa3\x6a\x41\xd5\x78\x6e\xc7\x4c\xd0\x8e\x95\x6e\xb6\xd1\xde\x5c\x62\xb6\xc1\x00\x00\x4c\x23\x07\xd2\x5f\xe7\x59\x23\x62\x2e\x8f\x57\x78\x79\x5e\xf7\xb7\xfe\xad\xc6\xbb\x6e\x04\x63\xfc\x77\x21\xcd\x91\xf9\x76\x8d\x19\x98\x48\x8d\xe5\x41\x30\xbd\xd7\xba\x54\xd5\x1f\xd3\x63\xde\xae\x46\x9c\xa9\x36\xe8\xf7\x13\x9e\xf1\x6a\x83\xc8\xa4\xea\x1d\xe6\xed\x82\xa8\xa0\xaa\xf9\xad\x7e\x7c\xa6\x44\xcc\xff\x70\x79\x24\xff\x86\xbd\x1b\x8b\x07\x23\x02\x54\x10\x54\xbf\x02\x00\x00\xff\xff\xf7\xea\xa0\xc3\x3c\x01\x00\x00")

func layoutFooterTmplBytes() ([]byte, error) {
	return bindataRead(
		_layoutFooterTmpl,
		"layout/footer.tmpl",
	)
}

func layoutFooterTmpl() (*asset, error) {
	bytes, err := layoutFooterTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "layout/footer.tmpl", size: 316, mode: os.FileMode(420), modTime: time.Unix(1520565142, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _layoutHeaderTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\x4d\x5b\x0e\xc2\x30\x0c\xfb\xcf\x29\xac\x1e\x80\xc2\x2f\x74\xbd\x4b\xb5\x04\x98\x34\x55\x15\xe9\xc7\xa6\x28\x77\x47\x74\xcc\x5f\xb6\xe5\x87\x19\x58\x9e\x4b\x15\x84\xb7\x14\x96\x4f\x80\x3b\x25\xed\xfb\x2a\xe8\x7b\x93\x29\x74\xd9\x7a\x9c\x55\x43\x26\x00\xb8\x1c\x39\xd8\x50\x3f\xb4\xc2\xbc\xd4\xd7\x1d\xb7\x6b\xdb\x1e\xc3\x76\x4a\x71\x6c\x64\xa2\xf4\x2f\xcc\x6b\x51\x9d\xce\x9b\x4c\x29\x1e\x34\x93\x19\xa4\x32\xdc\xbf\x01\x00\x00\xff\xff\x57\x43\x1e\x8f\x8e\x00\x00\x00")

func layoutHeaderTmplBytes() ([]byte, error) {
	return bindataRead(
		_layoutHeaderTmpl,
		"layout/header.tmpl",
	)
}

func layoutHeaderTmpl() (*asset, error) {
	bytes, err := layoutHeaderTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "layout/header.tmpl", size: 142, mode: os.FileMode(420), modTime: time.Unix(1520317669, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _layoutLayoutTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x93\xd1\x6e\xda\x3c\x18\x86\xcf\xb9\x0a\xff\x3e\xad\x6a\x03\x81\x00\xbf\x02\x52\xd6\x92\xac\x64\xa5\x09\x94\x12\x7a\x66\x12\x27\x71\x88\xe3\xd4\x76\x28\x59\xd5\x7b\x9f\x80\x6d\x95\xa6\x76\xaa\xb4\x23\xdb\xaf\xbf\x83\xe7\x91\xbe\xd7\xfa\xef\xfa\xee\xea\x7e\xe3\x4f\x41\xa6\x79\x31\x69\x59\xc7\x03\x14\xa4\x4c\xc7\x30\x27\xf0\x18\x50\x12\x4f\x5a\x00\x00\x60\x71\xaa\x09\x88\x32\x22\x15\xd5\x63\x58\xeb\xe4\x72\x08\x27\xad\xf3\x5f\xc1\xca\x1d\x90\xb4\x18\x43\xa5\x9b\x82\xaa\x8c\x52\x0d\x41\x26\x69\x32\x86\x99\xd6\x95\xfa\x1f\x63\x4e\x0e\x51\x5c\xa2\xad\x10\x5a\x69\x49\xaa\xe3\x23\x12\x1c\xff\x0e\x70\x0f\xb5\x51\x1b\x47\x4a\xbd\x65\x88\xb3\x12\x45\x4a\x41\xc0\x4a\x4d\x53\xc9\x74\x33\x86\x2a\x23\xc6\xb0\x77\xe9\x96\x7d\x63\xd8\x3b\x3c\x05\x1d\x22\xd6\xa1\x7d\xd1\xee\x0f\x17\xa1\x7f\xf0\x53\x33\x69\x7a\x37\xeb\xfd\xfd\x3c\x6b\x4f\xbb\xa6\x11\x72\x27\x9a\x15\x4b\xfb\x99\xb9\xa9\x63\xaf\x71\x6c\xb3\xa5\x39\x0b\x39\x04\x91\x14\x4a\x09\xc9\x52\x56\x8e\x21\x29\x45\xd9\x70\x51\x2b\xf8\x53\x58\x45\x92\x55\x1a\x28\x19\xbd\x49\x44\x22\xa6\x28\x7f\xaa\xa9\x6c\x4e\xf0\xe7\xeb\xa5\x81\xba\xa8\x83\x54\xc1\xf8\x09\x38\x7f\x97\xd7\x9b\x19\xa2\x7b\xed\xe9\x9b\xdd\x7e\x73\xe3\x19\xab\xe9\xfc\x3b\xbf\x1d\x78\x57\xbb\x85\xc4\x72\x3a\xc2\x41\x95\x9a\xc4\x7e\x74\x67\xcf\xce\xf5\xed\xc3\xdc\xc6\x6e\xe5\x3a\xce\xc8\xc8\xc2\xca\xed\x7b\xbb\xf9\xc7\xbc\x16\x3e\xb3\xfe\x0d\x3c\x2e\x73\x85\xa2\x42\xd4\x71\x52\x10\x49\x4f\xf4\x24\x27\x07\x5c\xb0\xad\xc2\x95\xa8\x2a\x2a\x51\xae\x70\x07\x75\xba\x68\x84\x6b\x1e\xff\x0a\x3f\x36\xb2\xab\xf9\x36\xcd\x46\x5f\x2e\x36\x9d\xc0\xd3\x7b\x63\x51\x0e\xd6\x06\x4f\xfd\x43\xb6\x1a\x79\x78\x19\x05\xca\xf6\x07\xd9\x8a\x6d\x43\x63\x94\x0f\x12\xb2\x73\x7c\xb5\xdb\x87\xb5\xda\x27\xa4\xbd\xed\x05\xff\x64\xf4\xd9\x7d\xca\xff\x5c\xa7\xf7\x5d\x66\x8f\x0b\x73\x59\xd1\x3c\xeb\xad\xda\xdd\x78\x98\xdf\x69\x73\xff\x6d\xfa\x35\xa1\x78\x16\xb8\x6c\xb1\x58\x06\xc1\x61\x99\x38\xeb\x8a\x75\x6e\x9f\xea\x87\xd8\x6e\xf2\x15\x91\xfd\x8b\x81\xe9\x3f\x5c\xf1\x4d\xf1\x19\x17\x0b\x9f\xfb\x64\x6d\x45\xdc\x4c\x5a\x2f\x2f\x40\x53\x5e\x15\x44\x53\x00\x23\x51\x6a\x5a\x6a\x08\x10\x78\x7d\x6d\x59\xf8\x3c\x62\xe1\x53\x33\x7f\x04\x00\x00\xff\xff\x19\x3a\x76\x04\xa9\x03\x00\x00")

func layoutLayoutTmplBytes() ([]byte, error) {
	return bindataRead(
		_layoutLayoutTmpl,
		"layout/layout.tmpl",
	)
}

func layoutLayoutTmpl() (*asset, error) {
	bytes, err := layoutLayoutTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "layout/layout.tmpl", size: 937, mode: os.FileMode(420), modTime: time.Unix(1518660036, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewAdminIndexTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x92\xc1\x6a\xc3\x30\x0c\x86\xef\x79\x0a\xa1\x7b\x97\x17\x48\x72\xdf\x6d\xb0\x07\x18\x4e\xac\xb4\x02\x5b\x32\x8e\xd2\x51\x42\xde\x7d\x24\x4d\x9b\x76\x50\x76\xd8\xc5\xc8\xf6\xaf\xff\xd7\x07\x9a\x26\xf0\xd4\xb3\x10\x60\xa7\x62\x24\x86\x30\xcf\x45\x31\x4d\x60\x14\x53\x70\x46\x80\x27\x72\x9e\x32\xc2\xdb\xfa\x55\x79\x3e\x43\x17\xdc\x30\xd4\x6b\x8f\x63\xa1\x7c\xe8\xc3\xc8\x1e\x9b\x02\x00\xa0\xea\x35\xc7\x9b\x64\xa9\x0f\xc7\xac\x63\x42\x70\x9d\xb1\x4a\x8d\xa5\xf3\x91\xa5\x1c\xf8\x28\x5f\x2c\x08\x91\xec\xa4\xbe\xc6\xa4\x83\x6d\x1e\xab\xcf\x43\xd2\x83\xcd\x2e\x58\x45\xc1\xb5\x14\xa0\xd7\x5c\x63\x74\x1c\xb0\x59\xce\xaa\x5c\x9f\x7f\x49\x59\xd2\x68\x60\x97\x44\x35\xd2\x2a\x7e\xb2\x5f\x68\xb2\x06\x04\xf6\x9b\x17\x88\x8b\x74\xab\x53\x70\x1d\x9d\x34\x78\xba\x47\xed\xa3\x96\x9e\xcf\xff\x98\x3c\xb9\x61\xf8\xd6\xec\xb1\xf9\xd8\xaa\xbf\x09\xee\x3d\xaf\x21\x76\xc9\x15\x64\xbf\x3f\xc1\xec\xe9\x2f\x81\xda\xd1\x4c\x65\x4b\x1e\xc6\x36\xb2\xdd\x73\x5b\x13\x68\x4d\x0e\x29\x73\x74\xf9\x82\xcd\x27\x1f\xe5\x5d\xaa\xf2\xda\xb4\x2d\x45\xb9\x8c\xd7\x14\x9b\xf1\xf3\x86\xf5\xaa\xb6\x6f\xd8\x34\x01\x89\x87\x79\xfe\x09\x00\x00\xff\xff\xef\xb0\x36\xe3\x9e\x02\x00\x00")

func viewAdminIndexTmplBytes() ([]byte, error) {
	return bindataRead(
		_viewAdminIndexTmpl,
		"view/admin/index.tmpl",
	)
}

func viewAdminIndexTmpl() (*asset, error) {
	bytes, err := viewAdminIndexTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "view/admin/index.tmpl", size: 670, mode: os.FileMode(420), modTime: time.Unix(1520565197, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewAdminLogged_inTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x94\xcf\x8e\x9b\x30\x10\xc6\xef\x3c\xc5\xc8\x77\x87\xa4\xd2\x5e\x5a\x12\x69\x2f\x3d\xb6\x55\xb7\x7b\x8e\x0c\x1e\xc0\x2d\x1e\x23\x7b\x9c\x04\x21\xde\xbd\x82\x90\x5d\xf2\xa7\xb7\x6e\x2e\xb6\x33\xcc\x37\xf3\xfd\x46\x76\xdf\x83\xc6\xd2\x10\x82\x28\x1c\x31\x12\x0b\x18\x86\x24\xe9\x7b\x60\xb4\x6d\xa3\x18\x41\xd4\xa8\x34\x7a\x01\xab\x29\x94\x05\xee\x1a\x04\xee\x5a\xdc\x0a\xc6\x13\xa7\x45\x08\x62\x97\x00\x00\x90\x3a\x40\x3f\xed\xc6\x9f\x55\xbe\x32\x24\x73\xc7\xec\xec\x67\xd8\xac\xdb\xd3\x97\x29\x38\x24\x59\x3a\xa9\xec\x92\x24\xd3\xe6\x00\x45\xa3\x42\xd8\x4e\x2d\x28\x43\xe8\x65\xd9\x44\xa3\x67\xd1\x6c\x54\x9d\xbf\x20\x75\xc8\x95\x87\xf3\x22\xf1\xd4\x2a\xd2\x32\x58\xc8\x2b\xa9\x95\xff\x73\x09\x4c\xfb\xdf\x31\xb0\x29\x3b\x39\x1b\x93\x48\x17\xc5\x49\xb5\x74\xde\x5e\x64\x6d\x27\x3f\x81\xed\x64\x53\xc9\xb5\x00\x55\xb0\x71\xb4\x15\xa9\xd2\xd6\x50\x1a\x4c\x45\x7b\x17\x59\x80\x45\xae\x9d\xde\x8a\xd6\x05\x5e\x48\x4d\x72\x79\x64\x76\x74\x11\xcc\x99\x20\x67\x92\x2e\x72\x63\x08\xa5\xa1\xd2\xc1\xa5\x4a\xb0\x72\x2d\x76\x2f\xa6\xa2\xef\x91\xb3\xf4\x9c\xb9\xe8\x2c\x1d\x5b\x9b\xbd\xa7\xa4\x0e\xbb\xe4\xbc\x5f\x90\x6a\x4c\x60\x59\x79\x17\xdb\xa5\x25\x05\xb5\xc7\xf2\xad\x71\x3e\xba\x7d\xa9\x0a\x76\x7e\xaf\x22\xd7\xe2\x3e\x59\x1a\x46\x0b\x37\x67\x79\xf6\x0f\x65\x83\x27\x59\xb8\x26\x5a\x02\xd5\x98\x8a\xa6\x68\x90\x81\x95\xbf\xb3\xbf\xe8\x4d\xcb\x31\x13\x8e\x72\xb3\x5e\xdf\x0d\x21\x47\x3e\x22\xd2\x4d\xfa\x24\x51\x3f\xbd\x8d\x23\x97\x1b\xb1\x7b\xc1\x22\x7a\xc3\x5d\x96\xd6\x4f\x37\xd5\x52\x6d\x0e\x37\x7f\xb5\xd7\xc9\xbf\x8e\xee\xeb\x64\xfd\x39\x72\x0d\x01\x99\x0d\x55\x59\xda\x2e\x31\xab\xf7\x43\xdf\x9b\x12\x56\x3f\x5d\x83\xab\x1f\xe8\xad\x09\xc1\x38\x7a\x1e\x29\xbe\x06\xf4\xe3\xa8\x5e\xdb\x61\xb8\xae\x78\x43\x3b\x86\xf1\x8e\x7c\x1c\xe2\xff\x84\xf9\x21\xea\xd1\xe4\x3d\xe6\x7f\xa0\x7e\x84\xfb\x1b\x1e\x61\x04\x00\xe3\x65\x81\xd8\x5e\x91\x7e\x40\x1b\x49\xcf\x38\xe7\x02\xf3\x72\xfd\xf0\x94\xce\xf1\xfb\xc3\xd3\xf7\x80\xa4\x61\x18\xfe\x06\x00\x00\xff\xff\xea\x6b\xa7\x55\xb5\x04\x00\x00")

func viewAdminLogged_inTmplBytes() ([]byte, error) {
	return bindataRead(
		_viewAdminLogged_inTmpl,
		"view/admin/logged_in.tmpl",
	)
}

func viewAdminLogged_inTmpl() (*asset, error) {
	bytes, err := viewAdminLogged_inTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "view/admin/logged_in.tmpl", size: 1205, mode: os.FileMode(420), modTime: time.Unix(1520578740, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewAdminTwo_factor_auth_keyTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\x59\x8f\xdb\xb8\x0f\x7f\xcf\xa7\x20\xfc\xc7\x7f\xa7\x05\xaa\x68\x26\xd3\x74\x5b\xe7\x28\xf6\x00\xfa\x50\x14\xbd\xf6\x7d\x40\x5b\xb4\xa3\x8e\x2c\x09\x12\x9d\x4c\x36\xc8\x77\x5f\xc8\x4e\x9a\x64\x92\x1e\xd8\xf5\x4b\x6c\x4a\x26\x7f\x07\x19\x6b\xb3\x01\x45\x95\xb6\x04\x59\xe9\x2c\x93\xe5\x0c\xb6\xdb\xc1\x60\xb3\x01\xa6\xc6\x1b\x64\x82\x6c\x41\xa8\x28\x64\x30\xec\x96\xa6\x91\xd7\x86\x80\xd7\x9e\x66\x19\xd3\x03\xcb\x32\xc6\x6c\x3e\x00\x00\x70\x06\x36\xdd\x4d\xba\x4a\xd7\x5a\xa6\x20\x02\x45\xe2\x1c\x9a\xb5\xd8\x45\x26\x5f\xb7\x18\x1d\x59\x74\xf9\x72\xb0\xce\xd2\x61\xc5\xa3\x52\xda\xd6\x39\x5c\xf7\xb1\xed\x60\xd0\xbf\x70\x54\xa0\xc1\x50\x6b\x2b\x0a\xc7\xec\x9a\x1c\x6e\xae\xfd\xc3\x59\x02\x61\xa8\xe2\x1c\x6e\x4f\xd7\x5c\xd4\xac\x9d\xcd\x21\x90\x41\xd6\x4b\xda\x17\xe9\x6b\xe4\x05\x55\x2e\xd0\x09\x97\x4e\x9c\x7c\x4f\xea\xc9\x81\xcd\xd3\xc9\x19\x63\x6d\xcb\x40\x4d\xb7\xff\x12\xeb\x02\xcb\xfb\x3a\xb8\xd6\x2a\x51\x3a\xe3\x42\x0e\xff\x2b\x50\xe1\x78\x7c\x9c\xa9\x8f\x8f\x46\xa3\x43\x50\xe9\xe8\x0d\xae\x73\x28\x8c\x2b\xef\x0f\xf1\xca\x38\xe4\x1c\x12\xd3\x63\x69\x2d\x89\x05\xe9\x7a\xc1\x39\x8c\x46\xc7\xf4\x77\xba\xf5\xca\x88\x53\x69\x92\xa1\x02\x8d\xae\x6d\x0e\x25\x9d\xe2\xbe\x9c\x6d\xa5\x15\x2f\x1e\x07\x0b\x17\x54\xb2\x1e\x95\x6e\x63\x0e\xe3\xeb\xff\xef\x15\x9e\xca\xce\xef\xf9\x60\x30\x55\x7a\x09\xa5\xc1\x18\x67\x5d\xf3\xa1\xb6\x14\x44\x65\x5a\xad\x76\xed\x34\x75\x66\xfe\x35\xe5\xd4\xe8\xc3\x43\x17\xf0\xa7\xcf\xe9\xfa\xd3\xad\xac\x71\xa8\xe0\xb7\x96\x17\x64\x59\x97\x98\x8c\x06\xf4\x7e\x78\xfa\xb2\x7c\xf4\xf6\x31\x9a\xe0\x56\xd9\x79\xee\x29\xc2\x22\x50\x35\xcb\x16\xcc\x3e\xe6\x52\x6a\x6e\x2d\xc5\x21\x7a\x6f\x68\x58\xba\x46\x7e\xf1\x12\xbd\x97\xb5\x73\xb5\x21\x81\x07\x08\x2e\x48\xad\x6e\x5f\xbe\x7c\xfe\xea\xd7\x17\xd7\xe3\xd7\x0d\xcf\x5e\x66\xd0\xe9\x30\xcb\xf6\xae\x6a\xdb\x59\xd6\x7b\xeb\x96\x14\x2a\xe3\x56\xf9\x42\x2b\x45\x76\x72\x68\x99\xbc\x0d\xe6\xc9\x1e\x81\xd1\xf6\xbe\xc1\x7b\x0a\xc3\x33\x2c\x18\x23\x71\x94\x71\x81\x81\x94\x2c\x50\xd5\x14\x25\x59\xd1\xc6\x84\x31\xb2\x0b\x24\x4c\xa8\x87\x71\x59\x3f\x05\xeb\x44\x20\x4f\xc8\x93\xde\xcd\x9b\xdb\xb1\x7f\x98\xec\xfc\x7e\x9e\x1a\xa4\xef\x99\xfc\x95\x7f\x38\xc2\x22\xa2\xfe\x9b\xf2\x9d\x77\x93\x6c\x3e\x95\xf8\x6d\xd9\xae\xf6\xa0\x13\xdb\x61\xaf\x51\x87\xb4\xc3\xd2\x81\x92\x8a\x18\xb5\x89\xaf\xb5\x9a\x95\xae\xd9\x6f\x42\xab\x82\xd3\x2a\x91\x8b\xc3\x13\x55\x47\xbf\xf8\x12\x1b\x8f\xba\xb6\x5a\xcd\xde\xbd\xfd\x4b\xbc\xe7\x05\x05\x51\x1b\x57\xa0\x11\x68\x8c\x28\x9d\xf0\x81\x6d\x10\x7e\x2d\x3e\x60\xe0\xdf\x93\x12\xe2\x1d\x86\xd1\xf8\x66\x2c\x6e\xae\xe6\x53\xdd\xd4\x80\x86\x67\x57\x6f\x88\x41\x33\x38\x0b\x6f\xba\xc2\xf0\xc1\xe0\xfa\x0a\x62\x28\xbf\x0d\x5e\x5b\x36\x92\xec\x5d\x1b\xf7\x1a\xeb\x06\xd3\x4f\x4d\x96\x82\x2e\xd3\x5a\xb7\x70\xb7\xa2\xe2\x6e\x17\x1c\x7a\x5b\x5f\xf5\x83\x33\xcb\x6e\xc6\xb7\x99\x3c\x97\x6e\x2a\x95\x5e\x1e\x75\xbf\x3c\x6e\xff\x9f\x9a\x85\xf7\x9e\x2c\xf0\x82\x00\xcf\x66\x01\xd0\x2a\xe8\xe6\x24\xad\x7f\xfc\x04\xa5\x53\xf4\xdf\x07\x24\x29\x99\xd4\xca\x14\x32\xe6\x9d\x10\xf2\x8b\xaf\x27\x05\x46\x7a\xf1\xfc\xd9\x66\x33\xfc\xf8\x69\xbb\xcd\xbe\x4b\xf4\x71\xa9\xca\x85\x46\xa4\x76\xf3\x97\x2a\xfa\xf9\xfb\xf0\x0c\x1a\xb4\x2d\x1a\xb3\x86\xee\x0f\xab\xe3\x54\x39\x63\xdc\x4a\xdb\x1a\x22\x95\x81\xb8\x67\x78\x46\xaa\x17\x13\x0b\x32\xf3\xcf\xfd\xbe\x3f\x9c\xa2\xa9\xec\x43\x17\x18\x5a\xdf\xf2\xd1\xb7\x2f\x3b\x81\x99\x26\x21\x38\x93\x41\x20\x54\xce\x9a\xf5\x2c\xdb\xdf\x65\xb0\x44\xd3\xd2\x2c\xdb\x6c\x86\x6f\x69\x3d\xec\xab\xfd\x48\x8c\x7f\xe1\xfa\x07\x43\x18\x09\x6a\xe2\x4e\x87\x44\x1b\xaa\xe0\x9a\x6f\x74\xc2\x8f\x4c\x4f\xbc\x00\xcb\xb4\x7d\x96\x49\x54\x8d\xb6\x92\x57\xee\xae\xc2\x92\x5d\xb8\x4b\x09\xe5\x12\x8d\x56\xc8\x94\x41\x43\xbc\x70\x6a\x96\x79\x17\xf9\x92\x5d\x3f\xeb\xeb\xc1\x16\xa8\x5c\x48\x9f\x07\x45\xd9\xfc\xbb\xd6\x9c\xd9\x63\xdb\xa6\x48\xa7\x96\x8b\x06\x59\x6c\x68\x9f\xf6\x1c\xe6\x79\x47\x76\xe1\xa2\x65\x76\x76\x9f\xb0\x60\x0b\x05\x5b\xe1\x83\x6e\x30\xac\xb3\xf9\xe7\xb6\x68\x34\x4f\x65\xbf\xed\xb1\xaf\xa9\xfe\x25\x63\xa7\x32\x7d\xe4\x76\x25\x4f\xcf\x5d\x95\x73\x7c\x38\x77\x6d\x36\x40\x56\xc1\x76\xfb\x4f\x00\x00\x00\xff\xff\x5a\x32\x4f\x12\xb4\x09\x00\x00")

func viewAdminTwo_factor_auth_keyTmplBytes() ([]byte, error) {
	return bindataRead(
		_viewAdminTwo_factor_auth_keyTmpl,
		"view/admin/two_factor_auth_key.tmpl",
	)
}

func viewAdminTwo_factor_auth_keyTmpl() (*asset, error) {
	bytes, err := viewAdminTwo_factor_auth_keyTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "view/admin/two_factor_auth_key.tmpl", size: 2484, mode: os.FileMode(420), modTime: time.Unix(1520565202, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewAdminUserTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x94\xcb\x92\x9b\x3a\x10\x86\xf7\x3c\x45\x97\xf6\x3a\xcc\xc9\x54\xcd\x22\x01\x76\x59\xe6\x52\x35\x0f\x90\x12\xa8\x01\x25\xba\x45\x6a\x79\xc6\x45\xf9\xdd\x53\xc8\x10\xdb\x8c\xc9\x66\xbc\x41\xa6\xa5\xaf\xff\xfe\x5b\xf4\x34\x81\xc4\x5e\x59\x04\xd6\x39\x4b\x68\x89\xc1\xe9\x54\x14\xd3\x04\x84\xc6\x6b\x41\x08\x6c\x44\x21\x31\x30\xf8\x2f\x87\xaa\x48\x47\x8d\x40\x47\x8f\x35\x23\x7c\xa5\xb2\x8b\x91\x35\x05\x00\x80\x15\x07\x98\xf2\x6a\xfe\x19\x11\x06\x65\x79\xeb\x88\x9c\xf9\x08\xff\x3f\xf8\xd7\x4f\x39\x78\x2a\xaa\x32\x53\x9a\xa2\xa8\xa4\x3a\x40\xa7\x45\x8c\x75\x96\x20\x94\xc5\xc0\x7b\x9d\x94\x5c\xa0\xd5\x4c\x5d\x76\x58\x71\x68\x45\x80\xf3\x83\xe3\xab\x17\x56\xf2\x68\xa0\x1d\xb8\x14\xe1\xd7\x1a\xc8\xeb\x9f\x29\x92\xea\x8f\x7c\x29\x8c\xa3\x5d\x89\x99\xda\xbb\x60\x56\xac\x39\xf2\x0f\x60\x8e\x5c\x0f\xfc\x81\x81\xe8\x48\x39\x5b\xb3\x52\x48\xa3\x6c\x19\xd5\x60\x7f\xb8\x44\x0c\x0c\xd2\xe8\x64\xcd\xbc\x8b\x74\x85\xca\xb8\x36\x11\x39\xbb\x02\x5b\xb2\xd0\x92\xe5\x2e\x91\x56\x16\xb9\xb2\xbd\x83\x35\x4b\x34\xfc\x81\x35\xcf\x6a\xb0\xdf\x12\x55\xe5\xf9\xe4\x95\xb2\x72\x96\xb6\xd4\x5e\x5a\x71\x68\x8a\xf3\x5a\xaa\xc3\x56\xff\x46\x6a\x8a\x18\xce\x7a\x93\xdf\xc8\x5d\x95\xcd\xc7\xf8\x10\x5c\xf2\xdb\x0a\xc6\xc7\xe6\x2b\xbe\xc0\xcc\x80\x99\x01\xc9\x57\xe5\xf8\xb8\xd9\x75\xd5\xae\x5d\x54\xde\xa8\x45\x8b\x1a\x7a\x17\x6a\xa6\xac\x4f\xf4\xd9\x08\xa5\xd9\x75\xab\x83\xd3\x3c\xef\x62\x4d\x0e\x56\x65\xfe\x77\x87\x95\x01\xcb\x95\xc3\x33\xc7\x0a\x83\x35\xbb\x61\x66\x3d\x0b\x18\x3a\xa7\xb9\x91\xfc\x89\x81\x92\xb7\x0a\xbc\x16\x1d\x8e\x4e\x4b\x0c\x35\x5b\xde\x05\xfc\x9d\x54\x40\xb9\x29\xb6\xbc\xb1\xfc\x7d\xf5\x7f\x17\x31\xbe\xb8\x20\xf7\x2c\x58\xe3\xfb\x2e\x6c\x53\x07\xf7\x72\x27\xf1\xbe\xca\x8b\x27\xf7\x4f\xbd\xb1\xda\xff\x95\x7c\x76\xdb\x6f\x4b\xb8\x76\xfc\xca\xe8\x4b\xa9\x37\x5e\x5f\x5e\xaf\x76\x83\x51\x56\xa3\x1d\x68\xac\xd9\x3f\x55\x45\x23\xb4\x5e\xb3\xce\x43\x87\x9b\x44\x28\x59\xf3\x45\x59\x65\x92\x01\xd7\xc3\x13\x74\xa3\x08\xa2\x23\x0c\xb1\x2a\xf3\x89\x1d\x77\xde\xb6\x75\xaf\xdb\xef\xb9\x00\xf7\x27\x82\x0f\xca\x88\x70\x64\xcd\x73\x6a\x8d\xba\xf3\xf9\xdf\xc9\x7b\x3b\x11\x72\x68\x79\xdc\x0e\xea\xde\x39\xba\x0c\xea\x69\x02\xb4\x12\x4e\xa7\x3f\x01\x00\x00\xff\xff\xe7\xb1\x29\x2e\xe5\x05\x00\x00")

func viewAdminUserTmplBytes() ([]byte, error) {
	return bindataRead(
		_viewAdminUserTmpl,
		"view/admin/user.tmpl",
	)
}

func viewAdminUserTmpl() (*asset, error) {
	bytes, err := viewAdminUserTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "view/admin/user.tmpl", size: 1509, mode: os.FileMode(420), modTime: time.Unix(1520580378, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewAuthIndexTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x92\xc1\x6e\x83\x30\x0c\x86\xef\x3c\x85\xe5\x7b\xc7\x0b\x00\xf7\xdd\x26\xed\x01\xa6\x40\x4c\xb1\x94\x38\x51\x62\x3a\x55\x88\x77\x9f\xa0\xb4\x94\x49\xd5\x0e\xbb\x44\x4e\xf2\xfb\xff\xfd\x49\x9e\x26\xb0\xd4\xb3\x10\x60\x17\x44\x49\x14\x61\x9e\x8b\x62\x9a\x40\xc9\x47\x67\x94\x00\x07\x32\x96\x12\xc2\xdb\xfa\x55\x59\xbe\x40\xe7\x4c\xce\xf5\xda\x63\x58\x28\x9d\x7a\x37\xb2\xc5\xa6\x00\x00\xa8\xfa\x90\xfc\x5d\xb2\xd4\xa7\x73\x0a\x63\x44\x30\x9d\x72\x90\x1a\x4b\x33\xea\x50\x66\x3e\xcb\x17\x0b\x82\x27\x1d\x82\xad\x31\x86\xac\x9b\xc5\x6a\xf3\x14\xf4\xe4\xb2\x0b\x56\x91\x33\x2d\x39\xe8\x43\xaa\xd1\x1b\x76\xd8\x2c\x67\x55\xae\xcf\xbf\xa4\x2c\x71\x54\xd0\x6b\xa4\x1a\x69\x15\x1f\xec\x17\x98\x14\x1c\x02\xdb\xcd\x0b\xc4\x78\xba\xd7\xd1\x99\x8e\x86\xe0\x2c\x3d\xa2\xf6\x51\x4b\xcb\x97\x7f\x4c\x1e\x4d\xce\xdf\x21\x59\x6c\x3e\xb6\xea\x6f\x82\x47\xcf\x6b\x88\x5d\x72\x03\xd9\xef\x07\x98\x3d\xfd\x25\x50\x3b\xaa\x06\xd9\x92\xf3\xd8\x7a\xd6\x47\x6e\xab\x02\xad\xca\x29\x26\xf6\x26\x5d\xb1\xf9\xe4\xb3\xbc\x4b\x55\xde\x9a\xb6\x9d\x28\x97\xf1\x9a\x62\x33\x3e\x2e\x58\x1f\x82\xee\x0b\x36\x4d\x40\x62\x61\x9e\x7f\x02\x00\x00\xff\xff\x2e\x2f\xc5\xf9\x9d\x02\x00\x00")

func viewAuthIndexTmplBytes() ([]byte, error) {
	return bindataRead(
		_viewAuthIndexTmpl,
		"view/auth/index.tmpl",
	)
}

func viewAuthIndexTmpl() (*asset, error) {
	bytes, err := viewAuthIndexTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "view/auth/index.tmpl", size: 669, mode: os.FileMode(420), modTime: time.Unix(1520565208, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewAuthLogged_inTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x93\xc1\x8e\x9b\x30\x10\x86\xef\x3c\xc5\xc8\xf7\x09\x49\xa5\xbd\xb4\x90\x1e\xda\x73\x5b\x69\x1f\xa0\x32\x30\x80\x5b\x7b\x8c\xec\x21\x1b\x64\xf1\xee\x15\x84\xec\xd2\x84\x8b\x07\x86\xff\xf3\x3f\xbf\xe5\x94\xa0\xa1\xd6\x30\x81\xaa\x3d\x0b\xb1\x28\x98\xe7\x2c\x4b\x09\x84\xdc\x60\xb5\x10\xa8\x9e\x74\x43\x41\xc1\x61\x6d\x15\x51\x26\x4b\x20\xd3\x40\xa5\x12\xba\x4a\x5e\xc7\xa8\xce\x19\x00\x00\xeb\x0b\xa4\xb5\x5a\x1e\xa7\x43\x67\x18\x2b\x2f\xe2\xdd\x67\x38\x1d\x87\xeb\x97\xb5\x39\x67\x45\xbe\x52\xce\x59\x56\x34\xe6\x02\xb5\xd5\x31\x96\xab\x05\x6d\x98\x02\xb6\x76\x34\xcd\x06\x2d\x16\xea\xf6\x07\xeb\x4b\xa5\x03\xdc\x16\xa4\xeb\xa0\xb9\xc1\xe8\xa0\xea\xb0\xd1\xe1\xef\xbd\xb1\xd6\x7f\xc6\x28\xa6\x9d\x70\x1b\x0c\x89\xef\xc4\x95\xda\xfa\xe0\xee\x58\x37\xe1\x27\x70\x13\xda\x0e\x8f\x0a\x74\x2d\xc6\x73\xa9\x72\x3d\x4a\x9f\x47\xd3\xf1\x6f\x3f\x8a\x02\x47\xd2\xfb\xa6\x54\x83\x8f\xb2\x23\xad\xb4\x6a\x14\xf1\x7c\xe7\x55\xc2\x50\x09\xa3\x1f\xc5\x1a\x26\x34\xdc\x7a\xb8\x6f\x12\x1d\x1e\xd5\xf9\xd5\x74\xfc\x73\x94\x22\xbf\x29\x77\xc6\xf2\xc5\xd9\x36\x7a\xce\xfa\x72\xce\x6e\xf5\x2e\x28\x6b\xa2\x60\x17\xfc\x38\xec\x7c\xa4\x14\x34\x77\x04\x87\x5f\xc1\x5f\xa7\x6f\x9e\xdb\x57\x6b\x6a\x9a\xe7\x0f\xb4\x86\x3e\x50\x5b\xaa\xaf\xac\x1d\x95\x29\x1d\x7e\x68\x47\xf3\xac\x9e\xb9\x68\x84\x1c\x3c\xbc\xe3\x2d\x19\x68\x2d\x5d\xb1\xf6\x76\x74\x0c\xda\x9a\x8e\xd7\x6e\xc4\x28\x3a\x3c\x25\xb3\xb3\xdd\xe0\xa2\x84\x37\x3c\x1d\x8f\x4f\xc7\x53\x91\xbc\x11\xf1\x83\x7c\x45\xf4\x2f\xef\x07\x55\xe1\x49\x9d\xdf\x8d\x17\x79\xff\xf2\xb0\x5d\xde\x98\xcb\xc3\xa7\xe1\x49\xfd\x9d\x62\x1d\xcc\xb0\x0c\xb3\x40\x86\x7d\xfa\x7a\x9f\x28\x71\xb3\xe5\xb7\x81\xb7\xe5\xff\x0b\xd2\x7a\x2f\x1f\x17\x24\x25\x20\x6e\x60\x9e\xff\x05\x00\x00\xff\xff\x01\x5c\x56\x87\x5d\x03\x00\x00")

func viewAuthLogged_inTmplBytes() ([]byte, error) {
	return bindataRead(
		_viewAuthLogged_inTmpl,
		"view/auth/logged_in.tmpl",
	)
}

func viewAuthLogged_inTmpl() (*asset, error) {
	bytes, err := viewAuthLogged_inTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "view/auth/logged_in.tmpl", size: 861, mode: os.FileMode(420), modTime: time.Unix(1520318596, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewAuthTwo_factor_authTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x90\x41\x8e\xab\x30\x0c\x86\xf7\x9c\xc2\xf2\xbe\x8f\x0b\x00\x9b\x77\x84\x39\x40\x65\x88\x99\x46\x4a\xec\x28\x75\x3a\xaa\x22\xee\x3e\x22\x43\x55\x66\x96\x4e\x3e\x7f\xb6\xff\x5a\xc1\xf1\xea\x85\x01\x17\x15\x63\x31\x84\x6d\xeb\xba\x5a\xc1\x38\xa6\x40\xc6\x80\x37\x26\xc7\x19\xe1\x5f\xfb\x1a\x9c\x7f\xc0\x12\xe8\x7e\x1f\x5b\x0f\x79\xe1\x7c\x59\x43\xf1\x0e\xa7\x0e\x00\x60\x58\x35\x47\xa0\xc5\xbc\xca\x88\x3d\x15\xbb\xf5\xf6\xa5\xd7\x95\x16\xd3\x7c\x6d\xf5\x83\x82\x77\x64\x8c\x10\xd9\x6e\xea\x46\x4c\x7a\xb7\x43\xd0\x24\xa7\x31\xbb\xef\xf2\x99\xb5\xa4\x13\xd0\xa0\x40\x33\x07\x58\x35\xef\xbb\x38\xc6\xe9\xbf\x3a\x1e\xfa\xf6\xfc\x07\xf5\x92\x8a\x81\x3d\x13\x8f\x28\x25\xce\xfb\x45\x67\xff\x7e\x4b\xd6\x80\x20\x14\xf9\xa5\x7b\xaf\xd3\x3b\xff\x38\x95\x73\x31\x53\x79\x09\x66\x13\x98\x4d\x2e\x29\xfb\x48\xf9\x89\xd3\x47\x99\xa3\xb7\xa1\xff\xc1\x8e\x58\xfa\x7d\xce\xd4\x1d\xaa\xdf\x19\xaf\xaa\xf6\xce\xb8\x56\x60\x71\xb0\x6d\xdf\x01\x00\x00\xff\xff\x78\x9b\xc2\xae\xa0\x01\x00\x00")

func viewAuthTwo_factor_authTmplBytes() ([]byte, error) {
	return bindataRead(
		_viewAuthTwo_factor_authTmpl,
		"view/auth/two_factor_auth.tmpl",
	)
}

func viewAuthTwo_factor_authTmpl() (*asset, error) {
	bytes, err := viewAuthTwo_factor_authTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "view/auth/two_factor_auth.tmpl", size: 416, mode: os.FileMode(420), modTime: time.Unix(1520499780, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewIndexTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x92\xdd\x69\xc3\x30\x10\xc7\xdf\x3d\xc5\x71\xef\xa9\x17\xb0\x3d\x43\x37\x28\xb2\x75\x4e\x0e\x64\x9d\x90\xce\x29\xc1\xf8\x25\x5d\xa5\x1b\x94\x4e\x94\x45\x8a\x1d\x11\x27\x85\xd0\x87\xbe\x1c\x42\xfa\x7f\xfc\x84\x34\x4d\x60\xa9\x67\x4f\x80\x9d\x78\x25\xaf\x08\xf3\x5c\x14\x95\xe5\x23\x74\xce\xa4\x54\xaf\x07\x86\x3d\xc5\x5d\xef\x46\xb6\xd8\x14\x00\x00\x55\x2f\x71\x00\xd3\x29\x8b\xaf\xb1\x34\xa3\x1e\xca\xc4\x7b\xff\xc6\x1e\x61\x20\x3d\x88\xad\x31\x48\xd2\xac\x5f\x3d\x77\xa9\x8b\x7d\xb7\x8f\x32\x86\x3b\xc1\x2a\x72\xa6\x25\x07\xbd\xc4\x1a\x07\xc3\x0e\x9b\x65\x56\xe5\xba\xfd\x4b\xca\x3e\x8c\x0a\x7a\x0a\x54\x23\xad\xe2\x87\xf8\x85\x3c\x8a\x43\x60\x9b\xb3\x20\x38\xd3\xd1\x41\x9c\xa5\x5b\xfc\x86\x57\x5a\x3e\xfe\x83\x36\x98\x94\xde\x25\x5a\x6c\x5e\xf3\xea\x6f\xea\x9b\xe7\x39\xf8\x26\x79\x80\xdf\xda\x9e\x5e\xa0\x1d\x55\xc5\xe7\xa6\x34\xb6\x03\xeb\xad\xa7\x55\x0f\xad\xfa\x5d\x88\x3c\x98\x78\xc2\xe6\x72\xfe\xbe\x9c\x3f\x2f\x1f\x5f\xd7\x59\x95\x57\x77\x7e\xed\x72\xe1\x6a\x8a\xdc\x50\x4c\x13\x28\x0d\xc1\x19\x25\xc0\x5e\x44\x29\x22\xbc\xac\x5f\x67\x9a\x80\xbc\x85\x79\xfe\x09\x00\x00\xff\xff\x28\x69\x18\xbd\x5c\x02\x00\x00")

func viewIndexTmplBytes() ([]byte, error) {
	return bindataRead(
		_viewIndexTmpl,
		"view/index.tmpl",
	)
}

func viewIndexTmpl() (*asset, error) {
	bytes, err := viewIndexTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "view/index.tmpl", size: 604, mode: os.FileMode(420), modTime: time.Unix(1519789277, 0)}
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
	"layout/footer.tmpl": layoutFooterTmpl,
	"layout/header.tmpl": layoutHeaderTmpl,
	"layout/layout.tmpl": layoutLayoutTmpl,
	"view/admin/index.tmpl": viewAdminIndexTmpl,
	"view/admin/logged_in.tmpl": viewAdminLogged_inTmpl,
	"view/admin/two_factor_auth_key.tmpl": viewAdminTwo_factor_auth_keyTmpl,
	"view/admin/user.tmpl": viewAdminUserTmpl,
	"view/auth/index.tmpl": viewAuthIndexTmpl,
	"view/auth/logged_in.tmpl": viewAuthLogged_inTmpl,
	"view/auth/two_factor_auth.tmpl": viewAuthTwo_factor_authTmpl,
	"view/index.tmpl": viewIndexTmpl,
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
	"layout": &bintree{nil, map[string]*bintree{
		"footer.tmpl": &bintree{layoutFooterTmpl, map[string]*bintree{}},
		"header.tmpl": &bintree{layoutHeaderTmpl, map[string]*bintree{}},
		"layout.tmpl": &bintree{layoutLayoutTmpl, map[string]*bintree{}},
	}},
	"view": &bintree{nil, map[string]*bintree{
		"admin": &bintree{nil, map[string]*bintree{
			"index.tmpl": &bintree{viewAdminIndexTmpl, map[string]*bintree{}},
			"logged_in.tmpl": &bintree{viewAdminLogged_inTmpl, map[string]*bintree{}},
			"two_factor_auth_key.tmpl": &bintree{viewAdminTwo_factor_auth_keyTmpl, map[string]*bintree{}},
			"user.tmpl": &bintree{viewAdminUserTmpl, map[string]*bintree{}},
		}},
		"auth": &bintree{nil, map[string]*bintree{
			"index.tmpl": &bintree{viewAuthIndexTmpl, map[string]*bintree{}},
			"logged_in.tmpl": &bintree{viewAuthLogged_inTmpl, map[string]*bintree{}},
			"two_factor_auth.tmpl": &bintree{viewAuthTwo_factor_authTmpl, map[string]*bintree{}},
		}},
		"index.tmpl": &bintree{viewIndexTmpl, map[string]*bintree{}},
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

