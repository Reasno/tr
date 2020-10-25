// Code generated by go-bindata. DO NOT EDIT.
// sources:
// NAME-service/handlers/endpoints.gotemplate (845B)
// NAME-service/handlers/handlers.gotemplate (62B)
// NAME-service/handlers/middlewares.gotemplate (75B)
// NAME-service/svc/client/grpc/client.gotemplate (3.184kB)
// NAME-service/svc/client/http/client.gotemplate (105B)
// NAME-service/svc/endpoints.gotemplate (4.25kB)
// NAME-service/svc/server/endpoints.gotemplate (746B)
// NAME-service/svc/transport_grpc.gotemplate (2.962kB)
// NAME-service/svc/transport_http.gotemplate (106B)

package template

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _handlersEndpointsGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x92\x41\x8b\xdb\x30\x10\x85\xcf\xd6\xaf\x78\x84\x3d\x6c\xa0\x95\xef\x81\x5c\xb6\xd9\xc3\x42\x37\xbb\x74\x43\x7b\x56\xec\x59\x5b\xac\x23\x99\x91\x6c\x53\x84\xfe\x7b\x91\x93\xd8\x5e\x02\xa5\xbd\x59\x33\xf2\x9b\xef\xbd\x51\x9e\xe3\x9b\x2d\x09\x15\x19\x62\xe5\xa9\xc4\xf1\x37\x3c\x77\xce\x49\xec\x5e\xb0\x7f\x39\xe0\x71\xf7\x74\x90\x22\xcf\xf1\x83\xb8\x33\x46\x9b\xea\x7c\x01\x83\x6e\x1a\xd8\x9e\x78\x60\xed\x09\xbe\xd6\x0e\xef\xba\xa1\xf1\xf2\x4f\x62\xa7\xad\xd9\x20\x04\x79\xf9\x8e\x71\xd1\xc0\x4e\x79\x5a\x76\xd3\x39\x46\x21\x5a\x55\x7c\xa8\x8a\x50\x2b\x53\x36\xc4\x4e\x08\x7d\x6a\x2d\x7b\xdc\x8b\x2c\xcf\x71\x48\x53\xde\x88\x7b\x5d\x90\xc8\xda\x23\x56\x21\xc8\xd7\x87\xa7\xf1\xce\xab\xf2\x35\xbe\xc6\xb8\x12\x59\x2a\x7f\x2e\x22\x77\x7d\xb1\x12\x6b\x21\xde\x3b\x53\x60\x4f\xc3\xa3\x29\x5b\xab\x8d\x77\xf7\xee\x2c\x88\xf6\x28\x43\x90\x17\x79\xb9\x57\x27\x8a\x31\x9d\x88\xd7\x70\x7d\x21\xa7\x3f\x10\x46\x9a\x87\xce\x69\x43\xce\xa1\xb4\x27\xa5\x8d\x14\x63\xf5\x17\xab\xf6\xca\x88\x41\xfb\x1a\x27\x5d\x96\x0d\x0d\x8a\xc9\x49\xbc\xd1\xec\x2e\x5f\x76\x2a\x2b\xb2\x2b\xc9\x76\x54\xb9\x88\x5c\xf9\xd6\x67\xfd\x2b\xc5\x34\x35\xeb\x15\xa7\x7c\x42\x60\x65\x2a\xc2\x9d\xc6\x66\x8b\xc9\xc7\x33\xf9\xda\x96\x2e\x85\x20\xb2\x2c\x84\x83\xfd\x6e\x07\x62\xdc\xe9\x8b\xc5\x49\x70\x3b\xba\x7c\x56\x1f\x14\xc2\x4d\x77\xa6\xc8\x42\x20\x53\x26\xb5\x44\x44\x53\x28\x9b\xed\xe7\x94\xc2\x3f\x23\xdd\x0c\xdb\x00\xc0\x5f\x50\xbf\x2c\x20\xe2\x22\x76\x47\x0d\x15\xe9\x1d\xcf\xab\xfa\xcf\x0d\xcc\x76\xce\x3b\x98\x5f\xc9\xd4\x49\xae\x99\x7c\xc7\x06\x53\x4d\x44\xf1\x27\x00\x00\xff\xff\x84\x19\x57\x16\x4d\x03\x00\x00")

func handlersEndpointsGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_handlersEndpointsGotemplate,
		"handlers/endpoints.gotemplate",
	)
}

func handlersEndpointsGotemplate() (*asset, error) {
	bytes, err := handlersEndpointsGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "handlers/endpoints.gotemplate", size: 845, mode: os.FileMode(0644), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x15, 0xb9, 0x3c, 0x15, 0x3e, 0x78, 0x2d, 0x90, 0xd5, 0x74, 0xc0, 0x7, 0x22, 0xbc, 0x3f, 0x5b, 0x70, 0xd1, 0x56, 0xd5, 0x9, 0xf6, 0xe9, 0x64, 0xf2, 0xf5, 0x15, 0x99, 0x47, 0x5, 0xaf, 0x69}}
	return a, nil
}

var _handlersHandlersGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\xc1\x0d\x80\x20\x0c\x05\xd0\xbb\x53\xf4\x4c\xa2\x1d\xc6\x09\x88\x7c\xab\x11\xa9\x69\xeb\x89\xb0\x3b\xaf\x77\x4e\xb4\x03\x24\xba\x86\xfd\xee\x2c\x68\xa2\xcf\x1d\x7c\xe5\x56\x2a\x8c\x03\xef\x57\x73\xc0\x37\x51\x3a\xd5\xe8\xd0\x02\x4a\x3c\xc6\x32\x03\x00\x00\xff\xff\xd6\x21\xab\x2e\x3e\x00\x00\x00")

func handlersHandlersGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_handlersHandlersGotemplate,
		"handlers/handlers.gotemplate",
	)
}

func handlersHandlersGotemplate() (*asset, error) {
	bytes, err := handlersHandlersGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "handlers/handlers.gotemplate", size: 62, mode: os.FileMode(0644), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x1e, 0xcb, 0xd5, 0x72, 0x80, 0xc6, 0xf9, 0x82, 0x4b, 0xe0, 0x8b, 0x90, 0xb8, 0x9b, 0xbc, 0x5d, 0x8d, 0x12, 0xd4, 0x8e, 0x54, 0xf6, 0x72, 0xcb, 0xef, 0xf5, 0x12, 0xd0, 0xe1, 0xb8, 0x41, 0xc8}}
	return a, nil
}

var _handlersMiddlewaresGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\x41\x0e\x02\x21\x0c\x05\xd0\xbd\xa7\xe8\x9a\x44\x7b\x18\x4f\x40\xec\xb7\x12\x81\x4e\xda\x4e\x66\x41\xb8\xfb\xbc\xb5\xb8\xd0\x1b\x20\xb5\x67\xfa\x19\xc1\x8a\xa9\xf6\x6f\xc9\xbf\x3a\xa5\xc3\x83\x13\xe3\xe8\x35\x11\x3c\x9a\x48\xc7\x55\x1d\xf1\x52\xa3\xaf\x39\x7d\x4c\x40\x85\xf7\x7e\xdc\x01\x00\x00\xff\xff\xcf\x9e\xe9\x81\x4b\x00\x00\x00")

func handlersMiddlewaresGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_handlersMiddlewaresGotemplate,
		"handlers/middlewares.gotemplate",
	)
}

func handlersMiddlewaresGotemplate() (*asset, error) {
	bytes, err := handlersMiddlewaresGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "handlers/middlewares.gotemplate", size: 75, mode: os.FileMode(0644), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xcc, 0xfe, 0x9d, 0x1a, 0xaf, 0x47, 0xe8, 0x97, 0x82, 0x24, 0x50, 0x17, 0xb4, 0x49, 0x73, 0x3b, 0x68, 0xb7, 0xe5, 0x3a, 0x3d, 0xb6, 0x15, 0x9d, 0xb1, 0x8f, 0xc4, 0x27, 0xaf, 0xa7, 0x3c, 0xc1}}
	return a, nil
}

var _svcClientGrpcClientGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x56\x4d\x6f\xdb\x38\x10\x3d\x8b\xbf\x62\xd6\x08\x16\x52\xa0\xd0\xf7\x2c\x7c\xa9\xd3\x2d\xba\xd8\xa6\x46\x1a\x74\x0f\x45\x51\x30\xd4\x58\x26\x2c\x93\x2a\x49\x3b\x31\x04\xfd\xf7\xc5\x90\x94\x23\x27\x8e\xdb\x43\x10\x8b\xf3\x38\x1f\xef\x0d\x39\x9c\x4e\x61\x6e\x2a\x84\x1a\x35\x5a\xe1\xb1\x82\x87\x3d\x78\xbb\x75\x8e\xc3\xcd\x67\xb8\xfd\x7c\x0f\xef\x6f\x3e\xde\x73\x36\x9d\xc2\x1d\xda\xad\xd6\x4a\xd7\x11\x00\x8f\xaa\x69\xc0\xec\xd0\x3e\x5a\xe5\x11\xfc\x4a\x39\x58\xaa\x06\x03\xf8\x2b\x5a\xa7\x8c\xbe\x86\xae\xe3\xe9\x77\xdf\x8f\x0c\x70\x23\x3c\x8e\xad\xf4\xdd\xf7\x8c\x20\x0b\x21\xd7\xa2\x46\xa8\x6d\x2b\xa1\xb5\x66\xa7\x2a\x74\x20\xa0\xbe\x5b\xcc\x41\x36\x0a\xb5\x87\xa5\xb1\xe0\x57\x48\x0e\xbe\xa0\xdd\x29\x89\xfc\x56\x6c\xb0\xef\xc1\xa5\x4f\xd6\x8e\xdc\x30\xa6\x36\xad\xb1\x1e\x72\x96\x4d\xa4\xd1\x1e\x9f\xfc\x84\x65\x93\xda\x98\xba\x41\x5e\x9b\x46\xe8\x9a\x1b\x5b\x4f\x09\xfd\xb6\x65\xba\x41\x2f\x2a\xe1\x45\x80\x28\xbf\xda\x3e\x70\x69\x36\xd3\x76\x5d\x4f\xd1\x5a\x63\xdd\x84\x1d\x5b\x6a\x73\xb5\x56\x7e\x4a\x7f\xa8\xab\xd6\x28\x4d\x81\xc9\x97\xb7\x42\xbb\x90\xd4\x1b\xf8\x03\x20\x25\xc5\xb2\xe9\x14\xee\x89\xe6\x54\x32\xcb\x26\x5d\xc7\x3f\x86\xca\x16\xc2\xaf\xe0\xaa\xef\x61\xea\x76\x54\x40\xfb\x00\x64\x5c\xbc\x3b\x36\x4f\x58\x11\x38\xbe\xc5\x47\xb0\xe8\xb7\x56\x3b\x10\x7a\x20\x0d\x1e\x84\x5c\xc7\x26\x38\xa6\x5b\x1a\xad\x51\x7a\x65\x34\x87\x8f\x1e\x94\x23\xf2\xc9\x8f\x45\xd7\x1a\xed\xd4\x83\x6a\x94\xdf\x83\x59\x06\x55\xa4\x68\x1a\xb4\xe0\x0d\x54\x4a\x34\x25\x08\x5d\x41\x23\x3c\x5a\x90\x8d\x71\x58\x46\xd0\xb3\x4f\xb6\xdc\x6a\x49\x39\xe5\xb4\x08\x97\x54\x2f\x9f\x87\xd0\x73\xa3\x75\x09\xa6\x25\x9c\x03\xce\xd3\xf2\xe7\xb0\x50\x40\xde\x3e\xf0\x57\x3d\x40\x5f\x68\x4b\x08\x8a\x14\xd0\xb1\x6c\x27\x2c\x48\x99\xaa\x99\x1b\xbd\x54\x35\x63\x19\x35\xd1\x8f\x12\x96\x70\x3d\x03\x2b\x74\x8d\x87\x38\x1d\xcb\x32\xb4\x96\x0c\xcb\xfc\x4f\x29\x0b\x96\x65\x6a\x49\x0e\xe1\x8f\x19\x68\xd5\x04\x44\x16\x19\xa4\xef\x14\xcc\xf1\xff\xac\x68\x73\xb4\xb6\x84\x89\x14\x5a\x1b\x0f\xa2\x6d\x9b\x7d\xf2\x3c\x21\x47\x3d\xcb\x7a\xc6\x32\x39\x2a\xc4\x51\xa4\x6f\xdf\x8f\xda\xe2\xa8\x52\x0a\x77\xca\xfa\x0e\x97\xc6\x62\x4e\xc9\xa4\xb6\xfe\x2a\x9a\x2d\xba\x7b\xf3\xe1\x6e\x31\xff\x94\xba\x35\x97\x92\xaf\x50\x54\x68\x5d\x51\x94\x14\x3e\xeb\xba\x2b\x78\x54\x7e\x05\x17\x1e\x29\x38\xef\x7b\x96\x8d\x56\xdb\x75\x4d\x64\x92\xe9\xc2\x23\x4f\x67\x32\xf2\x4b\xd1\x08\x19\x39\xbb\x50\x03\x68\x50\xe1\x13\xfa\x95\xa9\x5c\x04\x06\xee\xbb\xee\xde\xfc\x6b\x1e\xd1\xc2\x85\x4a\x22\xbd\x4f\xa7\x01\x86\x63\xc1\x87\x95\xb0\x2b\xf0\x4b\x61\xde\xde\x38\x83\x63\x46\x6e\xf1\x31\x92\x92\xc7\xbd\xc4\x88\x2e\xd3\xef\x49\xd7\x0d\x35\xf5\x3d\xef\xba\x71\xbe\x71\x71\x32\x86\xaa\x97\x8b\xef\xb5\x34\x15\x12\xa9\x23\xeb\x1d\xfe\xdc\xa2\xf3\x03\xe6\x06\x4f\x62\xc2\x09\xc1\x01\x14\x1a\xf6\x83\x09\xe4\x5e\x28\x3e\x98\xef\xf7\xed\x90\x48\xd7\x0f\xd8\xa3\x16\xe1\x9c\xa7\xf5\xe2\x40\x55\x5e\x84\x95\xa4\x08\xea\x2a\xa9\x98\x7e\x0d\x3f\xd8\xd0\xa9\x6e\x27\x0f\x7b\x5d\x47\x80\xb1\x86\x2f\x05\xa4\x0b\x23\xb8\x7b\xc5\xfd\x35\x00\x9c\x13\xb5\x7c\x8e\x9d\xf5\x25\x1d\x10\x16\xef\x76\x22\x07\xa2\x4a\x10\xe9\x62\xe7\x73\x88\x53\xe3\x2c\xb3\x74\x1d\x09\x38\xbe\x2d\x79\xdc\x31\x40\xfe\xa6\xfb\xc5\xaf\x44\xb8\xc9\x76\x68\xbd\x03\x41\x7e\xc3\x1d\x77\xa2\x0e\xb0\x48\x87\xd6\x1b\x10\xb0\x75\x68\xaf\x2a\xb3\x11\x4a\xbf\x01\x8d\x31\x38\x2c\xac\xda\x08\xab\x9a\x3d\xed\x59\x6e\x1b\x50\x1a\x44\xba\x74\xd2\x1d\x77\xb6\x90\xfc\x07\xa4\x43\xcc\xe7\xf1\x7f\x19\x5a\xfc\x2e\x24\xa3\xb4\x47\xbb\x14\x12\xbb\xbe\x80\x7c\xf4\x35\xbe\xe8\x62\xde\xd7\xb3\xe7\x7d\x3c\xbf\xfc\x75\xcb\x15\x87\x0e\x09\x0e\x06\xc5\x0e\xfd\xf3\x42\xb9\x78\x18\x7e\x4b\xb9\x73\xe7\xe6\xa4\x70\x71\x43\x42\xbc\xa5\xdb\xaf\x35\x89\x01\x82\x80\x67\x44\x0e\xa8\xdf\x12\xee\x5c\x1d\xa7\x74\x1b\x32\xf8\x4d\xd5\x7e\x86\x19\x94\xf2\x39\xa1\x58\x30\xbc\x21\xd8\xcf\x57\x72\x31\xbf\x6f\xf1\x68\xda\x81\xf3\x76\x2b\x3d\x05\x4b\x83\x00\xbe\x7d\x77\xde\x2a\x5d\xa7\x93\x39\x9e\x36\x51\x18\xaa\x3b\x7c\x05\x01\x36\xa6\x52\x4b\x85\x2e\xce\xee\xc3\xb3\x80\x26\x69\x88\x76\xb4\x9f\xb6\xe6\x97\xe3\x04\x8a\x58\x2e\x8b\x6c\xce\xfd\xd3\x30\xa7\xbe\xa0\xae\xf2\x35\xee\xc3\x70\x8f\x19\x15\xc7\xce\xba\x43\xad\xc1\xad\x81\x53\x8e\xc3\x40\x36\xc3\x94\x83\x19\x90\x4b\x36\x1e\xd1\x34\xf6\xfa\x14\xff\xdc\xac\x0c\xb9\x0c\xe4\x14\x70\x6a\xea\x8e\xbb\xf3\x45\x76\xd2\x3f\xbd\x6e\x86\x4d\x05\x97\xc3\xcb\x91\x7f\xba\x29\x5e\x22\x42\xf2\x34\x27\x5b\xa1\xc6\xca\x64\xc3\x13\x65\xfd\xfc\x44\x09\xe9\x85\xe9\xa8\x96\xb0\x2b\xc1\x04\x9b\xf4\x4f\x3c\x54\x93\xaf\x0b\x9e\xa7\xdc\xff\x22\x63\x1c\xa4\xd1\xf1\x8c\x1e\x23\xc4\x77\xf8\x2c\x61\x5d\xc2\x2e\x4c\x90\x3e\x3c\x4b\xe2\x23\x27\x42\xc7\xcf\x9c\xcb\x4d\x05\x33\x38\x14\xf0\x8f\x51\x3a\xbf\xdc\x54\xe5\xf3\xd2\x82\xf6\x44\xaf\x9c\xf3\xa2\x18\xdc\x25\x66\xa4\x7f\x8a\xec\xff\x1f\x00\x00\xff\xff\x00\xce\x0e\xa6\x70\x0c\x00\x00")

func svcClientGrpcClientGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_svcClientGrpcClientGotemplate,
		"svc/client/grpc/client.gotemplate",
	)
}

func svcClientGrpcClientGotemplate() (*asset, error) {
	bytes, err := svcClientGrpcClientGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "svc/client/grpc/client.gotemplate", size: 3184, mode: os.FileMode(0644), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x63, 0x72, 0x1f, 0xe5, 0x3a, 0x45, 0x1, 0x91, 0xd8, 0x5b, 0xa8, 0x47, 0x45, 0x45, 0x98, 0xee, 0x0, 0xf5, 0xc1, 0x3c, 0x43, 0xf0, 0x86, 0x3c, 0xec, 0xbe, 0x2d, 0x84, 0xed, 0x1a, 0x17, 0x6c}}
	return a, nil
}

var _svcClientHttpClientGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xcb\xb1\x0d\xc2\x30\x10\x05\xd0\x3e\x53\x5c\x1d\x09\xdf\x10\x34\x29\x91\xc8\x02\x56\xf8\x98\x88\xc3\x67\x9d\x3f\x95\xe5\xdd\x69\x18\xe0\x8d\xa1\xab\xdc\x01\x29\x7e\x61\x7c\x7b\xd7\x82\x5a\xfc\x7d\x52\x5f\x64\x63\xe4\xda\x9b\x07\x95\xf8\x34\xcb\x44\x2a\x2e\x4f\x0f\x39\xfc\x01\x59\x75\xce\x65\x8c\x23\x9b\x49\xda\xf6\xfd\xb6\xc1\x1a\x22\x5d\xed\x44\xe5\xfe\x27\x92\xe6\x5c\x7e\x01\x00\x00\xff\xff\x0b\x3c\x4c\x9e\x69\x00\x00\x00")

func svcClientHttpClientGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_svcClientHttpClientGotemplate,
		"svc/client/http/client.gotemplate",
	)
}

func svcClientHttpClientGotemplate() (*asset, error) {
	bytes, err := svcClientHttpClientGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "svc/client/http/client.gotemplate", size: 105, mode: os.FileMode(0644), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa1, 0xf0, 0x36, 0xf9, 0x16, 0xea, 0x9d, 0x4e, 0x73, 0x64, 0xc5, 0xad, 0xb3, 0x1b, 0x4, 0xe, 0xd8, 0xc8, 0x1e, 0xf7, 0x7a, 0x39, 0x40, 0x4c, 0xb2, 0x12, 0x83, 0x35, 0xca, 0x82, 0x6f, 0xd0}}
	return a, nil
}

var _svcEndpointsGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x57\x5d\x6f\xdb\x38\x16\x7d\x96\x7e\xc5\xad\x90\x45\xec\x42\x61\xde\x53\xe4\x61\xb7\xcd\xee\x06\xd8\x7e\xa0\xc9\xee\x3e\x14\x45\x41\x4b\x57\x16\x11\x8a\x64\x49\xca\x8e\x47\xd0\x7f\x1f\x5c\x52\x92\xe5\x5a\xed\x74\xe6\x71\x30\x0f\x41\x6c\x7e\x1c\x9e\x7b\xce\xbd\x97\xf4\xf5\x35\xbc\xd6\x25\xc2\x16\x15\x5a\xee\xb1\x84\xcd\x01\xbc\x6d\x9d\x63\xf0\xe6\x3d\xbc\x7b\xff\x08\x77\x6f\xee\x1f\x59\x7a\x7d\x0d\x1f\xd1\xb6\x4a\x09\xb5\x8d\x0b\x60\x2f\xa4\x04\xbd\x43\xbb\xb7\xc2\x23\xf8\x5a\x38\xa8\x84\xc4\xb0\xf8\x7f\x68\x9d\xd0\xea\x06\xba\x8e\x0d\x9f\xfb\x7e\x36\x01\x6f\xb8\xc7\xf9\x2c\x7d\xef\xfb\x34\x35\xbc\x78\xe2\x5b\x04\xb7\x2b\x52\x5a\xff\x38\xc2\x42\xa1\x95\xe7\x42\x39\x68\xd0\xd7\xba\x74\xe0\x35\x34\xfc\x09\x41\xa8\x52\xec\x44\xd9\x72\x09\xa8\x4a\xa3\x85\xf2\x0e\x2a\xab\x1b\x70\x68\x77\xa2\x40\x97\x13\x92\xc5\xaf\x2d\x3a\x0f\x5c\x95\x60\xd1\x19\xad\x1c\x82\x3f\x18\x0c\x48\xb4\x94\x82\xd0\x0e\x8f\x28\x39\x70\x07\x7b\x94\x92\xfe\xa3\x2a\x74\x89\xd6\x11\x00\xe1\x95\x38\x7c\xaf\xb4\x1d\x36\x06\xb4\x3c\x0c\x70\x12\xa7\x02\xdd\x5a\x70\xad\x31\xda\x92\xb8\xde\x72\xe5\xe8\x33\x1d\x27\xb8\x14\xbf\x70\x2f\xb4\x22\xb4\x4a\xdb\x86\x7b\xc7\xd2\x54\x34\x61\xc5\x2a\x4d\xb2\xaa\xf1\x59\x9a\x64\x14\x39\x3e\xfb\x2c\x4d\x93\x6c\x2b\x7c\xdd\x6e\x58\xa1\x9b\xeb\xad\xbe\x7a\x12\xfe\x9a\xfe\x46\xc6\xb4\xc4\x6c\x20\xeb\x3a\xf6\xe1\x1f\xf7\x01\xe8\x03\xf7\x35\x5c\xf5\x7d\x96\xae\x83\xa0\x77\x93\x44\x85\x96\x12\x0b\xef\x46\xae\xbe\x9e\x85\x0e\xbe\xe6\x1e\x0a\xdd\x18\x0a\x8c\x2b\xe0\x65\x39\xea\xc9\xe0\xde\x5f\x3a\x02\x6b\x90\x2b\x4f\xf2\x6d\x10\x5a\x87\x25\xe9\xc4\xa1\x46\x69\xd0\x82\xf3\xb6\x2d\x7c\x4e\xd3\xc3\x51\xcb\x27\x09\xe5\x35\x70\x82\x73\x42\x6d\x25\x82\xe1\x96\x37\xe8\xd1\x52\x2a\xd1\xf8\xbd\x02\x1e\x1d\xb2\x39\x08\x7f\xe9\xe8\xb0\xaa\x95\x41\xe9\xaa\x55\x05\xa9\x38\x50\x56\x48\x42\x6b\xd0\x26\x64\x34\x68\xda\x6b\xd0\x5e\x8d\x07\x12\xe0\x86\x3b\xe1\x18\xfc\x53\x5b\xc0\x67\xde\x18\x89\x39\x1c\x74\x0b\x8d\xd8\xd6\x1e\x0c\x77\xe4\xf2\x4c\x2a\x22\x38\x1d\x14\xcf\x31\x56\x97\x6d\x81\x41\x06\xae\xa0\xf6\xde\xb0\x7f\x73\x55\x4a\xe2\xb8\x17\xbe\x06\xe4\x45\x3d\x24\x2b\xac\xc6\xd3\xd7\xb0\x17\x16\x4b\x68\x4d\x04\x75\x06\x0b\x51\x89\x02\x0c\xf7\x35\x83\xd5\x7d\xe0\x27\x1c\xe1\x6f\xf8\x46\x1e\x80\x43\x23\x9c\x8f\x89\x0e\x25\x3a\xb1\x55\xb4\x55\xa8\x9d\x7e\xc2\x20\xe5\x43\xb4\x65\x2a\x8c\x40\x11\x4f\xcd\x8e\x66\x10\xc4\xa8\x24\x5b\xcf\xd5\x2d\xa4\x40\xe5\x4f\xd5\x9d\x19\x77\xac\x31\x79\xa0\x4a\x8c\x70\x58\xfe\xc8\x46\xaa\x86\xa8\x95\x20\x85\x1b\x8c\x69\x75\xe4\x2b\x94\x47\x5b\x71\x4a\xa8\x65\x27\x08\x6c\x3a\x6c\xb9\xce\x5b\x17\x3b\xd2\x50\x58\xd7\xc1\x87\x77\xb8\x7f\x3d\xc4\x53\xe8\x66\x23\x54\xd0\xa9\x19\x28\xce\x8c\xcd\x87\x6e\xe0\x5b\xab\x40\x84\x4c\x26\x82\x05\x97\x12\x6d\x4c\xe6\x81\x2c\x4b\x43\x38\x67\x82\x76\x69\xd7\x59\xae\xb6\x08\x17\x02\x6e\x6e\x81\x8d\xeb\xdf\x46\x33\xfa\x3e\x4d\xba\xee\x42\xb0\x77\xbc\xc1\xbe\x1f\xf7\x03\xc0\x14\x04\x1b\x07\xd3\xae\xbb\xa2\xd1\xbe\x4f\xfb\xd3\x5a\xfd\x89\x43\x28\x3b\x61\x35\x63\xb8\x86\xd9\xb9\xab\xc2\x3f\xc3\xd0\x47\xd8\xeb\xf8\x3f\xa7\x6c\x78\x69\x36\xac\xeb\xfe\xa5\x69\x19\x5c\x08\xf6\x31\x76\xc9\xc7\x83\xc1\x61\xeb\x1a\x56\xe7\x8b\x62\xfb\x9c\xad\xca\x01\xad\xd5\x76\x0d\x5d\x9a\x24\x63\x7b\x0d\x83\x44\x18\xd9\x82\x06\xc4\x89\x38\xac\xd3\x24\x11\x55\x58\xfa\xe2\x16\x94\x90\x01\x23\x19\x5c\x51\x42\x06\x98\x34\x49\xfa\x74\x1a\x1d\x4f\x60\x3f\xc3\x6d\x9d\x13\x4a\x9a\xf4\x69\xd7\x45\x79\x49\xdc\xb7\x54\x52\x73\x85\x43\xd1\x5e\x78\x0c\x0a\x47\xdf\xe6\xa2\x5f\x78\x5c\xd2\x3d\x0a\x4f\x60\x4b\x21\x3a\x08\xf4\xe6\x7b\xe3\x8a\x87\x50\x83\xeb\xf3\x24\x38\x09\x9e\xb0\x97\xad\x1b\x6f\xb3\xa9\x86\x3a\x32\x6a\xba\xd7\x66\xc3\xd1\x84\x99\x3b\x84\xfe\x95\x22\x1a\x30\x96\x34\x3c\x4b\x82\xb0\x6f\x37\x19\xea\xd8\x37\xc9\x15\x18\xc5\x55\x0b\x5e\x2e\xb9\x19\xfd\x9c\x66\x76\x83\x49\x71\x38\xa8\x1f\xbd\x9a\x7b\xf6\x7f\xcb\xcd\xdf\xa5\xbc\x7b\x2e\xd0\x78\xd8\x5b\x6e\x5c\x6c\xb3\x93\x7a\x95\x40\x59\xd2\x1d\x33\xd4\xe7\xb1\x60\x83\xbd\xa1\x3f\x2d\x5c\x9c\xec\xad\x28\x4b\x89\x7b\x6e\xe3\xfb\xe5\xbf\x6e\x7c\xd1\xd0\x5d\x6e\x8c\x3c\x50\x9b\xa1\xd6\xe9\x09\xbc\x99\x56\x87\xbb\x01\x77\x68\x0f\x93\x95\x54\x56\xd4\x45\xc6\xdb\x92\xf0\xde\x1b\xba\x39\xa8\x7b\xe6\xb3\xe6\x55\x70\x45\x37\x27\xdd\x37\x58\xd2\xb6\xcd\x01\x14\x79\x10\x6f\x54\x7c\x2e\x64\x5b\x62\x19\x1f\x33\x1b\x24\x0a\x14\xb3\xc1\x92\x9d\xa9\xb1\x3a\x72\xca\x21\x7b\xf0\xdc\xb7\x2e\xcb\x21\xfb\x20\xd4\x36\x5b\xa7\x63\x7b\x78\x39\xeb\x0f\xdf\xdb\x0f\x0b\xaa\xe4\x47\x36\x8c\x31\xe7\xad\x50\xdb\x90\x4e\x42\x0d\xc3\x37\xb7\xd0\x70\xf3\x29\x4e\x7d\x8e\xf2\x77\x3d\xd9\x4f\x6d\xed\xb7\xda\x57\x92\x64\xb3\x8c\xca\x6e\xa0\xeb\xf3\x61\x6b\xb4\x3f\xe9\xd3\x34\x21\x37\xbe\x10\x95\x90\xbe\x01\x72\xa2\xd5\xc5\x36\xf2\x25\x07\xfd\x44\xd3\x23\xb1\x4f\xf8\xfc\xf9\x15\xbc\xd0\x4f\x31\x15\x0d\x57\xa2\x58\x55\x8d\x67\x0f\xc6\x0a\xe5\xab\x55\x76\x37\x42\x4c\x0e\x5e\xfe\xcd\x5d\x42\xa9\xd1\x81\xd2\x1e\xf0\x59\x38\xff\x0a\x1c\xe2\xdc\xf8\x29\x77\x1c\xdb\xea\x8c\x48\xad\xd7\x43\x93\x2a\x51\xa2\xc7\xd5\xc8\x20\xcc\x1d\x03\x10\xaa\x38\xd2\x9f\xe4\xfb\x79\xa1\x44\x15\x20\x6e\x6f\xe1\x44\xb2\xa1\xd2\x16\x5b\x2d\xdc\xce\x98\xaf\x16\x97\xac\xc7\xd2\x3b\x91\x3c\x96\xdd\x7f\xf8\x06\x25\x96\xc7\x6c\x88\x8f\xff\x2d\xfa\x31\x77\xe7\x2f\xba\x98\xc2\xfb\x1a\xd5\x34\xab\x67\xe9\x3a\x80\xc5\xac\xcb\x63\x95\x0d\x85\xd0\xc6\xc5\x10\x7f\x51\xf0\xf8\xb3\x44\x14\xf4\xb0\xb1\xa2\x88\x2f\xce\x19\x87\x5a\x14\x75\xd8\xea\x50\x2d\x51\x18\x6e\xf3\x61\xf7\xf8\x96\xd1\x76\xb8\xcb\xcf\xa3\x0a\xed\x36\x26\x70\x7e\xde\x99\x17\x9a\x75\xfa\xbd\xb8\xfe\x70\x6f\x3a\x23\x95\x0f\x71\x06\xc5\x2d\x16\x28\x76\xf1\xd5\x17\x42\xfc\xe6\x31\xcd\xe0\x01\x71\x82\x99\xa1\x84\x89\xf1\x31\x3a\xd5\x3d\x11\xa5\x8c\x2c\xd1\x73\x21\xc3\xc3\x71\x2c\xa7\xf0\x9b\x64\x78\xf0\x72\x29\xfc\x81\xfd\xa8\x85\x9c\xc4\x3e\xef\x24\xbf\x5b\xd1\xbf\xfa\xcc\x9f\xa7\xcf\x9c\x6c\xcb\x97\x1f\x81\xdf\x6b\x3b\xbf\x06\x00\x00\xff\xff\x74\xb2\xce\x07\x9a\x10\x00\x00")

func svcEndpointsGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_svcEndpointsGotemplate,
		"svc/endpoints.gotemplate",
	)
}

func svcEndpointsGotemplate() (*asset, error) {
	bytes, err := svcEndpointsGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "svc/endpoints.gotemplate", size: 4250, mode: os.FileMode(0644), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x41, 0x2d, 0x18, 0xaf, 0x17, 0xfe, 0xa0, 0x91, 0xab, 0x63, 0xd2, 0xa4, 0xc4, 0xff, 0xae, 0x35, 0x3b, 0xa8, 0x91, 0xeb, 0x2a, 0xb, 0x54, 0xee, 0x1e, 0xd0, 0x4c, 0x36, 0xb6, 0xce, 0x17, 0x1}}
	return a, nil
}

var _svcServerEndpointsGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x92\x41\x8b\xfa\x30\x10\xc5\xcf\xc9\xa7\x18\xc4\x83\xc2\xff\x9f\xde\x85\x5e\x84\x3d\x2c\xac\x22\x28\xec\x39\x36\xb3\x36\x68\x93\x92\xc4\xf6\x30\xf4\xbb\x2f\xa9\x36\xad\x2b\xbb\xec\x5e\x67\x26\xf3\x7e\xef\x4d\x6a\x59\x9c\xe5\x09\xc1\xa3\x6b\xd0\x71\xae\xab\xda\xba\x00\x0b\xce\xb2\x0c\x0e\xa5\xf6\xb0\x47\xd7\xe8\x02\x39\xab\x8f\x30\x23\x12\xbb\xf5\x6b\x3f\xb3\x93\xa1\x84\xff\x5d\x37\xe3\x2c\x96\x1f\x8b\x90\xf9\xa6\xf8\xa6\x53\x4a\xa3\x2e\xe8\xfc\x8c\x2f\x39\xff\xb8\x9a\x02\xb6\xd8\xbe\x18\x55\x5b\x6d\x82\x5f\xf8\x9b\x1e\xd4\x47\x41\x24\xee\xea\x62\x2b\x2b\xec\xba\x7d\x4f\xb9\x04\xdf\x14\x22\xbd\x00\xea\x61\xd7\x57\xaf\x0d\x7a\x0f\xca\x56\x52\x1b\xc1\xfb\xea\xbb\x93\xf5\x60\x01\x5a\x1d\x4a\xa8\xb4\x52\x17\x6c\xa5\x43\x2f\x60\x8f\x08\x03\x4f\x36\xed\x9c\x2c\x67\x03\x49\x9e\x46\x44\x5c\x77\xdf\x36\x80\x2e\x6f\x42\x03\x4e\x92\x67\x8d\x74\x31\x47\x22\x27\xcd\x09\x61\xae\x61\x95\x43\x32\xb4\xc1\x50\x5a\xe5\x63\x24\x9c\x31\xa2\x83\x7d\xb3\x2d\x3a\x98\xeb\xbb\xd7\xb4\x30\xef\xed\x6e\xe4\x19\x89\x9e\xba\x23\x05\x23\x42\xa3\xe2\xb6\x48\x84\x29\x9d\x55\xfe\x18\x17\xfd\x1a\xe9\x49\x6c\x05\x00\xf0\x03\xea\xbf\x09\x44\x37\xc9\xdf\xe3\x05\x8b\x80\x0a\xc6\x9b\xfd\xf1\x14\xa3\x9d\x2f\xc7\x18\xff\x4d\x1a\x89\xf6\x1d\x86\xab\x33\x90\x6a\xbc\xe3\x9f\x01\x00\x00\xff\xff\xc8\xfd\xc4\xdc\xea\x02\x00\x00")

func svcServerEndpointsGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_svcServerEndpointsGotemplate,
		"svc/server/endpoints.gotemplate",
	)
}

func svcServerEndpointsGotemplate() (*asset, error) {
	bytes, err := svcServerEndpointsGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "svc/server/endpoints.gotemplate", size: 746, mode: os.FileMode(0644), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x56, 0x26, 0x44, 0x3a, 0x6b, 0x62, 0x34, 0x94, 0x55, 0xb7, 0xb3, 0x8c, 0x52, 0x9f, 0xd3, 0x9c, 0x3d, 0x4e, 0x23, 0x7f, 0xc1, 0x7b, 0xa7, 0xf8, 0xc2, 0xce, 0x27, 0x49, 0x72, 0x4d, 0x9c, 0x79}}
	return a, nil
}

var _svcTransport_grpcGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x56\x4d\x6f\xdb\x38\x10\x3d\x8b\xbf\x62\xd6\x28\x16\x56\xe1\xd0\x7b\x0e\x90\x4b\x93\x6e\x5b\xec\xe6\x03\x59\xa3\x7b\x28\x8a\x82\x96\xc6\x12\x61\x89\x54\x48\xca\x89\x97\xd0\x7f\x5f\x0c\xf5\x61\x39\x76\x1c\x1f\x0c\x58\xe4\xe3\xcc\x9b\xf7\x86\x23\xcd\xe7\x70\xad\x53\x84\x0c\x15\x1a\xe1\x30\x85\xe5\x16\x9c\xa9\xad\xe5\x70\x73\x0f\x77\xf7\x0b\xf8\x7c\xf3\x6d\xc1\xd9\x7c\x0e\x8f\x68\x6a\xa5\xa4\xca\x5a\x00\x3c\xcb\xa2\x00\xbd\x41\xf3\x6c\xa4\x43\x70\xb9\xb4\xb0\x92\x05\x06\xf0\x77\x34\x56\x6a\x75\x09\xde\xf3\xee\x7f\xd3\x8c\x36\xe0\x46\x38\x1c\xef\xd2\x73\xd3\x30\x56\x89\x64\x2d\x32\x04\xbb\x49\x18\xe1\x17\x7d\x58\xa8\x8c\xde\xc8\x14\x2d\x58\x34\x1b\x34\x17\x56\xa6\x08\x4b\xa9\x52\xa9\x32\x0b\x2b\x6d\xc0\xe5\x08\xd9\xe3\xc3\x35\x38\x23\x94\xad\xb4\x71\x81\xcb\x37\x07\xb5\x93\x85\xfc\x0f\x6d\x80\x0c\xbb\xf3\xcc\x54\x09\xff\x27\x84\xe3\x8c\xc9\x92\x16\x61\xca\xa2\x89\x42\x37\xcf\x9d\xab\x26\x2c\x9a\x24\x5a\x39\x7c\x71\x13\xc6\xa2\x49\xa6\x75\x56\x20\xcf\x74\x21\x54\xc6\xb5\xc9\x42\x88\x79\x89\x4e\xa4\xc2\x09\xc2\xd0\xc2\x90\x01\x26\x99\x74\x79\xbd\xe4\x89\x2e\xe7\x99\xbe\x58\x4b\x37\xa7\xdf\x3e\x05\x3a\xd6\x97\x4a\x6c\x64\x82\x2c\xaa\x96\x30\xf1\x9e\x3f\x7c\xfa\x16\x68\x3d\x08\x97\xc3\x45\xd3\x4c\x58\x1c\x74\xb9\x15\x6b\xfc\xf2\xf8\x70\xdd\xb2\x87\x52\xac\xd1\x82\x00\x8b\x0e\xf4\x0a\x50\xa5\x95\x96\xca\x59\x10\x1b\x21\x0b\xb1\x2c\x10\x04\xed\x07\x79\xbc\xe7\x5d\x1a\x7e\x27\x4a\x6c\x9a\x5e\x82\x55\xad\x92\x57\x91\xa7\xbb\x50\x9f\xfb\x7f\x33\xd0\x95\x93\x5a\x59\xe0\x9c\xef\xd5\xdb\x89\x79\x1f\xb6\x63\xa8\x96\xfc\x8d\x5c\xe0\x59\x64\x47\x58\x0b\x97\x57\xf0\xe3\xe7\xdb\xc1\x3c\x8b\xa2\x63\xbb\x9f\x70\xa5\x0d\x4e\x7b\x07\x16\xfa\xba\xb5\x2b\x9e\xb1\xa8\x79\x9d\xe3\x0a\x44\x55\xa1\x4a\xa7\x7b\xcb\x43\x39\x9c\xf3\x98\x45\x06\x5d\x6d\x14\xfc\x4e\xd9\xda\x1c\x3e\xd8\xe3\x3d\x2c\xf4\xdf\xfa\x19\x0d\xec\x95\x04\x4d\xc3\x22\xef\x8d\x50\x19\xc2\x07\x49\x85\x0c\xfb\xb7\xe8\x72\x9d\x5a\x42\x44\xde\xf7\xc7\x3f\xc8\x4e\x8b\x4b\xd8\x2f\xe9\x0e\x9f\x3b\xd5\x59\x14\x45\x83\xf2\xdc\xfb\xe1\x48\x6f\xc2\x8c\x10\x37\x98\xe8\x34\x98\x35\x42\x3c\xe2\x53\x8d\xb6\x05\x7c\x56\x47\x01\xb6\xd2\xca\x62\x40\xec\x29\xc1\x39\xa7\x45\xd2\xce\xfb\x0b\xea\x22\x62\xde\xb0\x26\xb4\xdc\x4e\x10\x90\x65\x55\x60\x89\xd4\x15\x74\xa3\xbc\xff\xa2\x83\x14\xc7\xbd\x96\xca\xa1\x59\x89\x04\x99\xdb\x56\x38\x8e\x63\x9d\xa9\x13\x07\x9e\xbd\xaf\xdf\x11\xf9\x00\x5e\xe9\xf7\x55\xa8\xb4\x40\xc3\x76\xe4\x5b\xe6\x5d\x98\x30\x24\x46\xd9\x9d\xde\x15\x72\x7e\x0d\xef\x52\x0d\xb7\x68\x6a\xe1\xe3\x2e\x55\xbc\x0b\x3f\xb0\x9f\x26\xee\x05\xba\xe1\xc2\xbb\xae\x9d\x81\xc1\x27\xf8\x18\xee\xcd\x0e\xdf\x39\xba\xd8\x56\x3d\xa9\x18\xa6\x87\xa0\xd6\xd5\x11\x6a\x06\x68\x8c\xa6\xe4\x2c\xfa\x45\xa1\xab\xb0\x42\xb4\xa9\xa7\x0e\xf4\x6c\xaf\x14\x75\x0b\x71\x0b\x5c\x62\x16\xc9\x55\x38\xf4\xdb\x15\x28\x59\x50\xa8\xfe\x86\x28\x59\x84\x78\xe1\xa2\x75\x6b\x06\x2b\x7e\x0e\xb5\x78\x46\xc7\x59\xc3\xbc\x6f\x8d\x22\x9b\x3a\xa9\xdb\xae\x7e\x5f\xe7\xf9\x1c\x4e\x5d\x00\x90\x34\xf0\x5e\x0d\xfb\xf6\x40\x87\xf8\x93\x8c\x72\xb9\x70\x64\xc3\x06\x0d\x8d\xcb\xd0\xe8\xed\x90\x3c\xec\x37\xd3\x45\x76\x1a\x04\xd4\x16\xcd\x45\xaa\x4b\x21\xd5\x29\x30\x87\x07\x23\x4b\x61\x64\xb1\xa5\x23\xab\xba\x00\xa9\xc2\xa4\x1e\xcd\xdc\x53\x75\x4c\x7f\x1d\x76\x09\xd5\xf2\x88\x4f\xbb\xae\xf4\xd4\x12\xa3\xa7\xb1\xf5\xd4\x52\x97\x57\xfd\x99\x63\xf6\x1c\xb4\xd7\xc8\xcf\xa7\x13\x4e\xb5\xe3\xe5\x2c\xa7\x4e\x4e\xa2\xa3\x56\xb5\x27\x7a\xc8\x5b\x5e\xbd\xef\x42\x97\x22\x78\x76\xc2\xd9\xaa\xd8\x9e\x65\xd5\xc9\x42\x8e\x79\x35\x30\x38\xd3\x2c\x5b\x91\x8a\xfd\xa9\xf3\x6e\xd3\xc8\x2f\x5b\x1d\x33\xec\x2b\x16\x15\x1a\xcb\xda\x1a\x0e\xde\x96\xc7\x67\x51\x99\x0e\x48\x7e\x7b\x13\xbf\x06\x10\x5d\x9a\xa8\xeb\x19\x6c\x02\xe5\xd0\x04\x65\x1a\x66\x84\x5c\xc1\x66\x3c\x33\xda\x0f\x1c\x84\x35\x6e\x83\xdb\x69\x4a\x1f\x9b\xda\xe5\x24\x71\x9f\x85\x06\x74\x29\x1c\x4c\xd7\x31\x3c\xe7\x32\xc9\x03\xb4\x28\xa0\x20\xbb\xba\x28\x42\xa5\xe1\xa5\x43\xdf\x67\xfc\x5a\x28\xad\x64\x22\x8a\xaf\x28\x52\x34\x7f\xe1\x96\x3e\x7f\x5c\x97\xc8\xea\xb6\x65\xa4\x83\x44\x28\x58\x62\x1f\x22\x49\xd0\x5a\x4c\x29\x37\x4a\x97\xa3\xe9\x32\xd3\x3e\x49\x71\x35\xd4\xfa\xaf\x74\xf9\x77\x51\xd4\xd8\x8e\x44\xaa\xf5\xc7\x1f\x3f\xe3\x77\x81\x6f\xb0\x9b\xae\xe3\x5d\x84\xf0\x6e\x1d\xac\x4b\xdc\x0b\x6b\xd8\xff\x01\x00\x00\xff\xff\x71\x92\xdd\x9a\x92\x0b\x00\x00")

func svcTransport_grpcGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_svcTransport_grpcGotemplate,
		"svc/transport_grpc.gotemplate",
	)
}

func svcTransport_grpcGotemplate() (*asset, error) {
	bytes, err := svcTransport_grpcGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "svc/transport_grpc.gotemplate", size: 2962, mode: os.FileMode(0644), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa2, 0xc7, 0xff, 0xf, 0x16, 0xfd, 0xd8, 0x43, 0xfc, 0x25, 0x2b, 0x9c, 0xfe, 0x3c, 0xf, 0x3a, 0x3e, 0xc7, 0x21, 0xdf, 0x32, 0x41, 0x25, 0x39, 0x7d, 0x4b, 0xb6, 0xc2, 0xc, 0x27, 0x81, 0x3c}}
	return a, nil
}

var _svcTransport_httpGotemplate = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xcb\xb1\x0d\xc2\x30\x10\x05\xd0\x3e\x53\x5c\x1d\x09\xdf\x1a\x29\x91\xe2\x05\xac\xf0\x31\x08\x93\xb3\xce\x1f\x1a\xeb\x76\xa7\x61\x80\x37\xa7\xae\xb2\x03\x52\xed\x42\xff\x8c\xa1\x15\x67\xb5\xd7\x93\xfa\x20\x3b\xbd\x9c\xa3\x9b\x53\x89\x77\x6f\x85\x18\xa9\x9a\xdc\xcd\xe5\xb0\x1b\x64\xd5\x88\x65\xce\xa3\xb4\x26\x69\xcb\xf9\xba\xa1\x75\x78\xda\xe1\x5f\x78\xfe\x1b\x49\x11\xcb\x2f\x00\x00\xff\xff\xdd\x3a\x4a\x8f\x6a\x00\x00\x00")

func svcTransport_httpGotemplateBytes() ([]byte, error) {
	return bindataRead(
		_svcTransport_httpGotemplate,
		"svc/transport_http.gotemplate",
	)
}

func svcTransport_httpGotemplate() (*asset, error) {
	bytes, err := svcTransport_httpGotemplateBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "svc/transport_http.gotemplate", size: 106, mode: os.FileMode(0644), modTime: time.Unix(1464111000, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x52, 0x57, 0x56, 0xc6, 0xb4, 0xe5, 0x1f, 0xf4, 0x1d, 0xa5, 0xda, 0x23, 0xea, 0x8f, 0xfb, 0xff, 0xae, 0x4b, 0x12, 0xe4, 0xf6, 0xbf, 0x11, 0xa6, 0x4, 0x83, 0x53, 0xfd, 0xbf, 0xce, 0x4a, 0x47}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"handlers/endpoints.gotemplate":     handlersEndpointsGotemplate,
	"handlers/handlers.gotemplate":      handlersHandlersGotemplate,
	"handlers/middlewares.gotemplate":   handlersMiddlewaresGotemplate,
	"svc/client/grpc/client.gotemplate": svcClientGrpcClientGotemplate,
	"svc/client/http/client.gotemplate": svcClientHttpClientGotemplate,
	"svc/endpoints.gotemplate":          svcEndpointsGotemplate,
	"svc/server/endpoints.gotemplate":   svcServerEndpointsGotemplate,
	"svc/transport_grpc.gotemplate":     svcTransport_grpcGotemplate,
	"svc/transport_http.gotemplate":     svcTransport_httpGotemplate,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"handlers": &bintree{nil, map[string]*bintree{
		"endpoints.gotemplate":   &bintree{handlersEndpointsGotemplate, map[string]*bintree{}},
		"handlers.gotemplate":    &bintree{handlersHandlersGotemplate, map[string]*bintree{}},
		"middlewares.gotemplate": &bintree{handlersMiddlewaresGotemplate, map[string]*bintree{}},
	}},
	"svc": &bintree{nil, map[string]*bintree{
		"client": &bintree{nil, map[string]*bintree{
			"grpc": &bintree{nil, map[string]*bintree{
				"client.gotemplate": &bintree{svcClientGrpcClientGotemplate, map[string]*bintree{}},
			}},
			"http": &bintree{nil, map[string]*bintree{
				"client.gotemplate": &bintree{svcClientHttpClientGotemplate, map[string]*bintree{}},
			}},
		}},
		"endpoints.gotemplate": &bintree{svcEndpointsGotemplate, map[string]*bintree{}},
		"server": &bintree{nil, map[string]*bintree{
			"endpoints.gotemplate": &bintree{svcServerEndpointsGotemplate, map[string]*bintree{}},
		}},
		"transport_grpc.gotemplate": &bintree{svcTransport_grpcGotemplate, map[string]*bintree{}},
		"transport_http.gotemplate": &bintree{svcTransport_httpGotemplate, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
