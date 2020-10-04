// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package statsviz

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// assets statically implements the virtual filesystem provided to vfsgen.
var assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2020, 10, 2, 22, 33, 52, 264833589, time.UTC),
		},
		"/app.js": &vfsgen۰CompressedFileInfo{
			name:             "app.js",
			modTime:          time.Date(2020, 10, 3, 0, 9, 18, 356511442, time.UTC),
			uncompressedSize: 1622,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x84\x54\x4d\x6f\xdb\x38\x10\xbd\xe7\x57\xcc\x12\x39\xd0\x8e\x43\x05\x7b\xc8\xc1\x5e\x2d\xb0\x71\x72\xc8\x02\x6d\x83\x18\x45\x8e\x01\x4d\x4d\x6c\x22\x14\x29\x68\xa8\x28\x6e\xe1\xff\x5e\x90\xb4\x65\xd9\x51\x5a\x5d\x38\x20\xdf\xbc\x79\xf3\x25\xfe\xd2\x58\xe5\xb5\xb3\xc0\x47\xf0\xf3\x0c\x00\xa0\xbb\x39\xe7\xba\xd8\x5f\x86\xaf\x46\xdf\xd4\x16\x0a\xa7\x9a\x12\xad\x17\x2b\xf4\x77\x06\x83\x79\xb3\xb9\x2f\x02\x78\x16\xb1\xdb\xb3\x63\x9e\x65\xa3\x4d\xf1\x84\x4b\x72\xea\x15\xfd\xf7\xc7\x7b\xde\x67\x7d\x93\x35\x18\xa7\x20\x87\x56\xdb\xc2\xb5\xc2\x38\x25\x83\xdf\x04\x5a\x7a\xae\x6a\xe7\x21\x07\xd6\xd2\x94\xcd\x3a\x1f\xfd\x02\xdc\x38\x25\xc2\xab\x53\xce\x40\x9e\xe7\xc0\xd6\xde\x57\x34\x65\x7d\xf2\xf0\xb5\xf4\xdc\xd4\x3a\x91\x1c\xb1\x6c\x4f\x33\xdb\xc7\xbb\x00\x96\x65\x0c\x2e\x82\x2e\xb1\x76\xe4\x77\x66\x25\xfd\xda\xca\x12\x03\xa0\x25\xd6\x4f\x56\x39\x4b\x1e\x0a\xe9\xe5\x23\x7a\xb4\x41\xff\x02\x95\xb3\x05\x41\x0e\xd7\x57\xb3\x93\x92\x34\x55\x21\x3d\x2e\xbc\xf4\xc4\x4b\x2c\xa3\xd1\xd7\x4d\xe1\x42\x54\x0d\xad\x6f\xa5\x97\xdc\x62\x0b\xb7\xd2\x23\x1f\x4d\xa0\x83\xef\x38\xf7\xf5\x68\xb4\xd0\xf4\x20\x1b\xc2\x82\x8f\x4e\x6b\x90\x12\xec\x65\xde\x99\x06\x93\x6c\xc8\x77\x41\xc9\x68\x85\x7c\x28\x93\x81\x88\x95\x71\x9e\x20\xcf\xc1\x36\xc6\x9c\x06\x0d\x98\xc4\x69\xd0\xae\xfc\x9a\x8f\xe0\x1f\xf8\xfb\x14\x35\x20\xef\xb8\x39\x87\x7a\x68\xab\xfd\xdc\x48\xa2\x85\xfe\x81\x87\xc2\x89\x2f\x58\x8a\x9b\x4d\xb8\xec\x4b\xdc\x67\x87\x26\x48\x1c\x88\xba\x46\x59\x4d\xe1\x9c\xb3\x60\x5c\x86\x54\xd8\x68\xf2\x01\xb5\x8c\xc4\x14\x81\xcb\x0d\x05\x7b\x10\xbb\x3d\x09\xdd\x68\xa1\x6a\x94\x1e\x1f\x42\x8d\xb8\xab\x3c\x4d\x62\xa5\x27\x51\xd1\x68\x36\xd4\x8d\x2c\x8b\x92\xdf\x17\x4a\x1a\x3c\x12\x9d\x65\xf1\x28\xb5\x9d\x82\x75\x2d\x5c\xc2\xf5\xd5\xe4\xc3\xab\x7c\x8f\xaf\x47\x0f\x7d\x61\x8d\x16\x69\xf4\x92\xa8\x6c\x9c\x22\x8d\xb3\xd8\xc0\xa4\xef\xb0\xc5\xf1\xcc\xc6\xf0\x84\xcb\x45\xdc\x5d\x50\xd2\x98\xa5\x54\xaf\x04\xe3\x2c\x3d\x07\xb9\x69\xb1\x21\x87\x30\xa7\x1d\x98\x0f\xec\xfd\x8e\x3b\x6c\x8b\x33\x28\x8c\x5b\x71\xf6\x9f\xf7\x58\x56\x5e\xdb\x15\xcc\x9d\xb5\x18\x37\x44\x08\xc1\xf6\xcd\x4c\xfe\xc2\x59\x57\xa1\x85\x3c\xfc\xa9\xf2\x7f\x7b\xa5\x39\x62\x5b\x34\x4a\x21\xd1\x4b\x63\xcc\x66\xcf\x87\x05\xeb\x95\xfb\x30\x4b\x9f\x4d\x79\xbf\x9d\x5d\x70\x65\x1c\x85\x96\xe0\x1b\x5a\xff\x3b\x01\xa9\x18\xf3\x00\x2f\x7a\x19\x4d\x81\x4d\x92\x73\x5f\x4b\x62\x27\xb4\x05\x67\x73\xa3\x03\x75\xf2\xfc\x8b\x8d\x86\x75\x60\x5d\xbb\x3a\xe8\x48\xe7\x9f\x74\xdc\x05\x58\x8a\x1d\xac\xcf\xb2\x2b\x91\x48\xae\x3e\xc9\xaf\xff\xb7\xfa\x7f\xf1\xed\xab\xa8\x64\x4d\xc8\x23\x52\xc4\x91\xe9\x66\x66\x1b\x7b\xfc\x2b\x00\x00\xff\xff\xb5\x62\xb7\xa5\x56\x06\x00\x00"),
		},
		"/buffer.js": &vfsgen۰CompressedFileInfo{
			name:             "buffer.js",
			modTime:          time.Date(2020, 9, 13, 22, 23, 8, 683400866, time.UTC),
			uncompressedSize: 1561,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x84\x54\xcb\x6e\xdb\x30\x10\xbc\xfb\x2b\x06\xb9\x44\x42\x6c\x2b\xe7\xa6\x2a\xd0\xa6\x3d\xe7\x12\xa0\x87\xa2\x28\xd6\xd2\xca\x22\x40\x93\x04\xb9\x72\x60\x14\xfe\xf7\x82\x7a\xd8\x96\xc3\xb8\x3e\x99\xcb\xe1\x72\x76\x66\xa8\xa2\xc0\xb7\xae\x69\xd8\xa3\xe6\x4a\x93\xe7\x00\x69\x79\xaa\x55\x9a\x42\x58\x2f\xf6\xe4\xa7\x4a\x89\xac\xe9\x4c\x25\xca\x1a\x64\x39\xfe\x2e\x00\x0c\xb0\x09\x31\x94\xfa\xb2\x35\x41\x7c\x57\x89\xf5\x99\x66\xb3\x44\x45\x2e\xbf\xd8\x8f\x3f\xd5\x20\xab\xc8\x61\x05\xcd\x06\x9f\xf1\x78\x0d\x98\x1a\x59\xcd\xeb\x1f\xde\x5b\x9f\xdd\x9d\xf1\xbb\x2e\x08\x36\x0c\x67\x83\x12\xb5\xe7\xbb\xfc\x69\x76\xf8\xb8\x98\x2d\x8b\x02\xaf\x2f\xdf\x5f\x32\xf2\x3a\xff\x84\x3d\xe9\x8e\xd1\x05\x65\xb6\x78\x3d\x38\xae\xbf\x7a\x4f\x07\x78\x92\x96\x3d\xa4\x25\x83\xa1\xd2\xb2\xe7\x59\x1f\x69\x55\x58\xff\xd9\x74\x0d\x4a\x18\x7e\x1b\x60\x71\x8c\xab\xeb\x07\x9c\xb3\x01\x25\x1e\x53\x5b\x71\x86\x32\x4e\x92\xda\x8c\x63\x96\x51\xb3\xf3\xe6\xf1\xf4\x4f\x53\x90\x2c\x25\x66\x7f\x56\xb3\xd9\x4a\x9b\xe5\x28\xcb\xb4\xa2\xd2\x7a\xfb\x86\xfb\x67\x32\xc6\x0a\x2a\xd2\x7a\xea\x68\x0d\xc8\x80\x77\x4e\x0e\xa3\xa3\xf7\xd7\x9a\x5e\xae\x3c\x4b\xe7\xcd\x59\x91\x5f\xa7\x99\x7f\xa7\x68\xbb\x2e\xb4\x99\x93\x0f\x89\xf7\x62\x7d\x29\xcf\x0a\xa4\xc8\x17\x05\x76\x76\xcf\xa8\x49\x08\x62\xfb\xc0\x6e\x78\xbb\x55\x26\x5a\x69\x9b\xa1\xd0\x73\x5f\x82\x9b\x86\xab\x98\x0d\x7d\x40\xad\x42\x45\xbe\x56\x66\x9b\xea\x19\x4f\x55\xe4\x56\xd1\x14\xab\x6b\x0e\x02\xd6\xbc\x63\x23\x21\xa1\xdf\x38\xef\xba\xb2\xee\xf0\x53\x49\xab\x4c\xf6\xb8\xbc\x70\x6e\x75\xb6\xf8\x2a\x15\xd7\xc9\x38\xe1\x6e\x66\x37\xa5\x30\x4a\x38\xf9\x20\x72\x0f\x0f\xc9\xd8\x4c\xc1\xb8\xa9\xff\x05\xf5\x84\xfc\x33\xcb\xff\x47\x7b\x06\x76\x36\xa4\x48\x15\x05\x82\x56\x15\x8f\xe0\x00\x1a\xd7\xa3\x95\xd1\x10\x4d\x12\x0d\x89\x96\x3b\xab\x8c\x04\x38\xcf\x81\x8d\x40\x99\x0b\xbf\xd7\xa7\xa6\x7d\x87\x2c\x31\x42\x51\xe0\x99\x5c\x7f\xa6\x56\x3b\x36\x21\x7e\xca\xc6\x9b\x06\x02\x5c\x8f\xf7\x8f\xd9\xea\x73\x46\x7b\x52\x9a\x36\x9a\xdf\x09\x17\xe9\x8d\x92\x4d\xf2\xa6\x64\x1b\xde\xfa\x0c\x76\x53\x39\xcd\x82\x20\xe4\xe5\x14\x91\xe8\xcd\xea\xfd\xd7\xe2\xfa\x09\xae\x87\xd1\xfb\xb3\xcb\xb1\xc5\x03\xe6\x39\x1c\xc4\x3f\x2e\x86\x2b\xc7\x0e\xc3\x73\x7f\x5a\x1c\xb3\x3c\x7f\xfa\x17\x00\x00\xff\xff\x93\x89\x21\xf8\x19\x06\x00\x00"),
		},
		"/index.html": &vfsgen۰CompressedFileInfo{
			name:             "index.html",
			modTime:          time.Date(2020, 10, 2, 22, 36, 8, 120825946, time.UTC),
			uncompressedSize: 659,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x8c\x92\x41\x8b\xdb\x30\x10\x85\xef\xfe\x15\x13\xdd\x6d\x91\x5b\x09\xb2\x2f\x4d\x8e\xad\x4b\x1b\x0a\x3d\x2a\xd2\x38\x52\x90\x25\x21\x8d\x0d\xe9\xaf\x2f\xaa\x9d\x5d\xb3\xbb\x90\x3d\x0d\xf6\x7b\xf3\x66\xe6\x43\x62\x77\xec\xbf\x9e\xff\xfc\x38\x81\xa1\xd1\x75\x95\x28\x05\x9c\xf4\xd7\x96\xa1\x67\x5d\x55\x09\x83\x52\x77\x15\x00\x80\x20\x4b\x0e\xbb\x5f\x24\x29\xcf\xf6\xaf\xe0\xcb\x77\xb5\x88\x23\x92\x04\x65\x64\xca\x48\x2d\x9b\x68\xa8\xbf\xb0\x6e\x23\x79\x39\x62\xcb\x34\x66\x95\x6c\x24\x1b\x3c\x03\x15\x3c\xa1\xa7\x96\x3d\x22\xc1\x7a\xc2\x34\x48\x85\x0c\xf8\xda\xbc\xab\xeb\x63\x0f\xdf\xfb\x33\xfc\x3c\x7d\xeb\x7f\x9f\x0e\x40\xc6\x66\x50\x61\x1c\xd1\x13\x5c\x91\x32\x24\x8c\x4e\x2a\xd4\x60\x3d\x68\x9c\xd1\x85\x58\xc4\xba\x5e\x33\x96\x99\x90\x93\x6a\x99\x21\x8a\xf9\xc0\xb9\xd2\xbe\x89\x2e\x50\xe3\xee\xbc\x54\x77\xaf\x9d\x24\xcc\xd4\x8c\xd6\x37\xb7\xcc\x3a\xc1\x97\xbe\xae\x12\x7c\xa1\x50\x89\x4b\xd0\xf7\xd7\xc5\xe0\x11\x86\x31\x28\x73\xcb\xcd\xd5\x92\x99\x2e\x8d\x0d\xcb\x1f\x9e\x50\xba\x9a\xec\x88\xf0\xb2\x8b\xb6\x33\x58\xdd\xb2\x92\x88\x69\x45\xf4\x5f\x31\xfb\x0d\x5b\xb3\x5f\xfd\x5c\xdb\xf9\xc1\x78\xd3\x1b\xcb\x7e\x1b\x69\x7b\x62\xc3\x2f\xd3\x30\x60\x7a\x73\xc5\x07\xbe\x5c\xe6\x3d\xb7\x4d\xf6\xb9\x27\xc4\xcf\x24\xc9\x18\xdf\xb1\x5d\x98\x16\xc8\xe5\x0d\xfe\x0b\x00\x00\xff\xff\x73\x11\x9b\xf3\x93\x02\x00\x00"),
		},
		"/opts.js": &vfsgen۰CompressedFileInfo{
			name:             "opts.js",
			modTime:          time.Date(2020, 10, 2, 8, 12, 44, 180745210, time.UTC),
			uncompressedSize: 19279,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xd4\x3c\xfd\x73\xdb\x36\xb2\x3f\xd7\x7f\xc5\x56\xd3\xab\xa8\x98\xa6\x24\x27\xee\xbb\xc8\x51\x32\x49\x9a\x6b\x32\xef\xfa\x9a\x57\xfb\xf5\xde\x9d\xc7\x33\x81\x48\x48\x44\x4d\x01\x3c\x02\x94\xa9\xb8\xfa\xdf\x6f\x16\x20\x25\x7e\x80\x94\xec\xf4\x23\xcd\x4c\x24\x91\xdc\x5d\x00\xbb\x8b\xfd\xc2\xd2\xc3\x21\x88\x58\x49\xa0\x59\x2c\x12\x25\x41\x85\x14\xe2\x48\x28\x89\xb7\x99\xe0\xf2\x68\x45\x12\x03\x32\x05\x67\x9e\x72\x1f\xef\x82\x33\x80\xbb\xa3\xa3\x23\x00\x80\xed\x3d\x25\x44\xa4\x58\x2c\xdf\x47\xe9\x82\x71\x07\x71\x10\x0a\xf2\x7f\x5b\x38\xc6\x99\x72\x52\x57\x13\x75\x21\x20\x8a\x94\xc1\xf0\x5f\x44\x95\x9e\x04\x4c\x21\xf5\x12\x21\x94\xf7\xef\x94\x26\xeb\x0b\x1a\x51\x5f\x89\xc4\xe9\x79\xe9\x89\x58\xd1\xa4\x37\x38\x3f\x6a\x20\x2a\xe5\x6b\x3c\x3f\x4d\xa4\x48\x14\x12\x09\x84\x9f\x2e\x29\x57\x9e\x9f\x50\xa2\xe8\x9b\x88\xe2\x95\xd3\x0b\xd8\x0a\x49\x94\x29\x28\xe5\x7b\x7e\x44\xa4\xfc\x1f\xb2\xa4\x30\x85\x5e\xbe\xaa\x5e\x13\x4c\xd1\x4c\xbd\x16\x5c\x51\x8e\x63\xf4\x9c\xcc\x5d\x0f\x2c\x60\x52\xad\x23\xea\xc5\x82\x71\x45\x93\x37\x2b\xca\x35\x2b\x7b\x5c\x70\xda\x01\x2d\x99\xe6\xd5\x14\x7a\x64\x26\x45\x94\xaa\x76\xe0\x19\xf1\x6f\x16\x89\x48\x79\x80\xe0\xc9\x62\x46\x9c\x91\x3b\x72\x4f\xcf\xce\xdc\x91\x37\xae\x4f\x09\xf9\xea\x91\x38\xa6\x3c\x78\x1d\xb2\x28\x70\x94\xf2\xeb\x6c\x4c\x3d\x49\x13\x46\xa5\x66\x1e\x8a\x29\xbf\xf6\x96\x24\x76\x1c\xe9\x02\x1b\xc0\xf4\x79\x4d\x68\xf8\x8f\xcd\xc1\x61\x30\x9d\xc2\x68\x00\x09\x55\x69\xc2\x6b\x94\x77\x42\xba\x9f\x58\xcc\x8a\x0f\x91\x4c\x0e\x59\x13\xce\xa5\x81\xfd\xd2\x0e\x7c\xb8\x88\x6a\x08\x7b\xa5\x54\x81\x6f\x11\xd4\xc8\x26\xa6\x0a\xa2\x2f\x22\x91\xc0\x14\xa4\xf9\xd5\x01\x19\x30\x19\x47\x64\xad\x61\x65\x28\x6e\xe1\x05\xf0\x34\x8a\x60\xd2\xba\x1a\x8b\x3e\x58\x78\x6f\x84\x09\x4a\x55\x1f\x6d\xea\x9a\xb3\xdd\xe4\x21\x0b\xe8\x25\x8b\xa5\x33\xb0\xa8\xc9\x4e\x77\x77\xf3\x6d\x9b\xdf\x4e\x17\xbd\xb9\x48\xde\x10\x3f\x74\x1c\xa5\xda\x55\xf0\x60\x35\x6c\xe1\x5b\xdb\x3c\x36\x35\xae\x6c\x5a\x16\x8e\x5c\xbf\xd7\xc2\x51\x3e\xbf\xf3\xb2\x71\x07\x4a\x6d\x24\xcd\x20\x57\xec\xfa\xfc\x50\xfe\x1c\xaa\x57\x7b\xf8\x65\xd4\x2e\x08\xf4\x6e\xfb\x3b\x93\x8a\x72\x9a\x38\xbd\xa5\x48\x25\x8d\x28\x59\xd1\x9e\x8b\x4e\xa6\xd5\xca\x7c\x59\x18\x78\x2f\x12\xfe\x0d\x0d\x06\x2d\x3c\x19\x0e\xbf\xc0\x55\xaa\xd7\x1a\xd8\xb9\x8b\xe8\x5c\x4d\xe0\x64\x3c\x72\x41\x89\x58\xff\xda\x0c\xec\x8b\xdf\xa9\xb0\x65\x75\xdd\xdb\xa0\x6b\x75\x14\x8d\x4c\xc7\xea\x76\x0a\xb4\x67\xaf\xd9\xe6\x57\xe2\xf2\x4e\x23\xb7\xab\x4f\xeb\x5c\xf2\x05\x97\x0a\xee\x00\xb9\xa2\x19\xe2\x02\x0b\x32\xd8\x94\x1c\x68\x6d\xd0\xe1\x10\x54\xc8\x24\x30\x09\x21\x4d\x28\x28\x01\x21\xe1\x41\x44\x51\x2a\xe8\xd9\x19\x89\xc0\x60\xc2\xd6\x3e\x32\x89\x93\xa8\xd3\xe1\x42\xc1\x02\x2d\x3f\x38\x3e\xe1\x30\xa3\x3a\xe4\x58\xb2\x8f\x34\x80\xd9\x1a\x02\xc1\xf8\x02\x96\x22\xa1\xa0\x59\x36\xd4\x6a\x01\x52\x11\x45\x41\x25\x84\xe7\xc4\x55\x42\xfc\x1b\xc6\x17\x83\x1a\xfd\x2f\x50\x4d\x70\x61\xf0\x1c\x46\x8d\x87\x5f\xec\x02\x84\x96\x1d\x59\xf3\x88\x35\x68\x4d\x78\xaa\x19\x07\xc7\xd0\x8b\xb3\xda\x16\x68\x20\x28\x11\xc3\x14\x59\xbc\x0f\xbc\x1e\x56\xf4\xe0\x18\x52\xf4\x35\x97\xe2\x27\x12\x39\x46\x54\xbd\xac\x37\xf0\x94\xf8\x1b\xcb\x68\xe0\x9c\x0e\x90\xa6\x0b\x35\x48\x2d\xce\xde\xba\x01\x88\xae\xa6\x2e\x0b\x14\x40\xc1\x7d\x98\xa7\x89\x0a\x69\x82\x42\x40\x19\x91\x38\x8e\xd6\x28\x0b\xbd\x10\xa9\x05\x1d\x64\x10\xb0\x40\x3f\xf6\x43\xc2\x17\xf4\xe8\x53\x8c\xd7\xc1\x61\x43\xc3\x68\x59\x49\x19\x13\xd5\x6e\x11\x7e\x1d\x0d\xfe\xbd\x34\xf9\x20\x8d\x2e\x6b\xf6\x61\x1a\x5d\x66\x6b\xf6\x13\x89\x34\x67\x31\x18\xbf\x1a\x5d\x5f\xb1\x20\x6b\xf1\x08\x08\xbe\xae\x80\xb3\x02\xbc\xcd\x83\xd8\x14\x5a\x8f\xb8\x55\xda\x75\x7e\xd5\xd0\xcc\x86\x1f\xca\xb7\xdd\xf7\x44\x85\x9e\x8e\xa4\x9c\xd4\x5b\x91\xe8\x52\xbc\x17\xd2\x41\xa2\x2e\xf4\xb3\xfe\x60\x60\xdd\x64\x0d\x6a\x66\x4f\xda\x89\xad\x35\x31\xe9\x49\x9f\x44\xb4\x83\x60\xd3\x13\x58\x6c\x71\x1e\x3e\x55\x15\x32\x14\xe2\x46\x4e\x6c\xdb\x81\x33\xe5\x36\xfd\x42\x61\xc6\xad\x8f\x2e\x70\x9a\x13\xb8\xb2\x2e\x18\xb3\xad\x1b\xba\xee\x08\x1d\x0a\x67\x20\x90\xc7\x62\xe1\xf4\x0b\x92\x7d\x83\x69\x67\xe4\xa6\x71\xf7\xda\x3e\x39\xbd\x61\xbb\x66\xc7\x82\xec\xbe\xb3\xd3\x34\xfb\x06\xf5\x81\xd3\xdb\xec\x2e\x37\x86\xc4\x26\x4f\x6a\x87\x43\xf8\x91\x2e\xc5\x8a\x9a\x8c\x38\xa2\x0b\xca\x03\x20\x6a\x9b\x1f\xc3\x4c\x28\x25\x96\x40\x78\xa0\x3d\xb6\x84\x15\x89\x52\x2a\x01\xf7\x71\x48\x21\x61\x8b\x50\x79\xd5\x04\xd9\x50\x79\x29\xf3\x64\x24\xcf\x93\xef\x60\x9b\xd5\xb8\xc6\xc4\xc2\x14\xee\x60\x97\x2c\xbc\xc6\x98\x7f\x92\x27\x0c\x98\xd5\xc1\xe9\x93\xa7\x2e\x8c\x9f\x7e\xe3\xc2\xc8\x7b\x7a\x3a\xe8\xb9\xe0\xe7\x30\xb3\x88\xf8\x37\x3d\xd8\x68\x17\x7e\xb7\x29\xdb\x40\xdc\xb9\x66\x06\x6f\xca\xb6\xc0\x9a\x96\x37\x13\x72\x83\xd7\x95\x90\x1b\x18\x4b\x4a\x6e\x50\x4d\xee\x86\xb1\x90\x97\x68\xce\x3a\xbd\xf4\x84\xf1\x88\x71\x5a\x4f\xf6\x76\x59\xde\xd7\x5f\xdb\xf0\x49\x10\x38\x5b\x98\x46\xee\xfa\x5e\xc7\x5e\x52\xb2\x05\x77\xb6\xc8\x9a\xaf\xae\x2d\x22\xa7\x99\x7a\x19\xb1\x05\x9f\x40\x0f\x8d\x4b\xaf\xa9\xc0\x95\xbc\xb0\x08\x78\x9b\x60\xb9\xb5\x6d\x07\x28\x9c\xc9\xa4\x94\x2d\xba\x16\x37\x87\x01\xea\xa8\xf9\x40\x87\xab\x96\xfb\x1f\xdf\xf1\x80\x66\x13\x18\x8f\x2c\x0f\x67\x22\xbb\x08\x49\x20\x6e\x27\xd0\x3b\x8d\x33\xc0\xff\xe3\x51\x9c\x41\x25\xfb\x3c\x1b\x58\x26\xe2\x79\x86\x6b\xdd\x41\xe8\x70\xa8\xe3\x50\x30\x3e\xd9\xe8\x21\x2c\x49\x72\x43\x13\x69\x09\x35\x59\x90\x27\xd7\x5b\xc9\x54\x14\xe9\x65\x14\x69\x5d\x32\x04\x1a\xba\x34\x17\x09\x3a\x40\x05\x0c\xa6\x30\x3a\x07\x06\xcf\x72\x8a\x5e\x44\xf9\x42\x85\xe7\xc0\x8e\x8f\x9b\x9e\xd1\xc0\x5c\xb1\xeb\xd6\x6c\xcf\x32\x55\xb1\xa2\x49\xb7\xc6\x17\x25\xa8\x32\xae\xc1\xca\xc7\xc1\x8b\x79\x24\x6e\x71\xa0\x15\x93\x6c\x16\x51\x4b\xec\x85\x9b\xa1\xb0\x30\x8c\x2b\x51\xd8\x97\x94\x07\xd2\x46\xbb\x9c\xac\x17\x6c\xb4\x88\x05\xad\xd2\x50\xcb\x26\x2f\x94\xa0\x69\x32\xf1\x07\xcd\x98\xb2\x52\x3e\x28\x63\x81\xea\xae\xaa\x87\x19\x8d\xcc\xaf\x9b\x78\x2d\xd9\x6b\x27\x9e\x8b\xca\xaa\x82\xba\xa4\x94\xaf\x12\x17\xa7\x59\x78\x5f\xb9\x74\xe5\x50\x69\x1c\x10\x45\x0f\x48\xa0\x6a\xb9\x93\xd5\x12\xe6\xc1\x07\x06\x7e\x73\x91\x2c\x75\x2d\x0b\x2f\x22\x1c\x02\xe3\xa1\x5d\x56\x61\xe2\xa3\x6d\xd6\x30\xe8\xfd\x4a\xe1\xc5\x64\x5f\x90\x31\xc9\x97\x5c\x73\x98\x16\x7f\x59\x2d\xfb\xa4\x4b\xc2\x5f\xad\x15\x95\xce\x0c\x3f\x5d\x90\xb8\x57\xe7\x24\x92\xd4\x85\x00\xc3\xad\x71\x99\x87\x86\x7f\x2a\x4c\xa8\x0c\x61\x8a\xc0\x2f\xd0\x90\x8d\x00\xed\xd9\xe9\x93\x92\xa4\x31\xf6\xd5\x91\x1a\x99\xe5\xb4\x07\xf0\x2c\xc7\xac\x4b\x25\xe7\x87\x86\x82\x63\xe8\xc3\xab\xbe\x95\x6f\x66\xf4\x94\x33\x6d\x92\x24\xab\x10\x79\x01\x57\xfd\x9b\x57\x7d\x17\xfa\xdf\xeb\xcf\xef\xf4\xe7\xa5\xfe\x7c\xaf\x3f\xdf\xe8\xcf\x7f\xe9\xcf\x7f\xbe\xea\x5f\x57\xd0\x27\x70\xd5\xff\x6f\x66\xf0\xcd\xd7\x77\xe6\xeb\xd2\x7c\xbd\x37\x5f\x6f\xcc\xd7\xbf\xcc\xd7\x3f\xd9\xab\x7e\x29\xf0\x46\xb5\x4e\x61\x0a\x27\xe3\xf3\xda\x9c\x13\x64\xe4\x08\x1e\x3d\x82\x20\x2e\x31\x29\x10\x35\x4e\x18\x16\x0c\xa7\x39\xa3\xaa\x0a\x79\x7c\x9c\x96\xd8\x02\xb7\x21\x8b\x68\xce\x64\x13\x0e\xd7\xf9\xfd\x08\x92\x01\x0c\x21\x81\xe7\x05\x41\xf4\xd0\x29\x3c\x33\x3c\xcc\x8d\x30\x9c\xc0\xb8\xbc\x47\xcb\xd2\xd8\xa6\xa2\x41\x3c\xd0\x92\xe9\x63\xce\x8a\xc8\x57\xe9\xb5\x5d\xa5\x74\x5c\xf5\xb7\xa5\x8e\x4d\x56\x65\x49\xe7\x74\x4b\x2a\xb7\x72\x41\x25\x29\x1d\x54\x08\x19\x86\x99\x0d\xf9\x83\x39\xc7\x28\x05\x45\xc2\xbf\x99\x68\xa4\x9d\xa2\xcf\x85\x9f\x36\x36\x4e\x9c\x08\xf4\xb1\xdf\x94\x22\xc6\xdd\x4f\x1d\x21\x34\x50\xd0\x00\x4f\x72\xd5\xb7\x61\xc9\x35\xf7\xeb\x38\x37\x14\x03\x08\x25\x6b\xfe\xb8\x14\x43\x57\xe7\x9a\x93\xdb\x9c\xd7\x98\xb6\xf0\xff\xce\x38\x2d\x8e\x62\x2c\x5c\x3b\xd0\x56\x04\x09\xb9\x9d\xa0\x0e\xb6\x85\xe6\x85\x01\xf4\x55\x5e\x36\x3a\xef\x00\x5b\xf8\x7a\x9f\x29\x82\xba\x42\xa4\xfa\xee\xb5\xec\x02\x5f\x32\x13\x2b\xec\xf2\xd2\xd1\x75\x27\x3c\xc9\x6a\xf0\xdb\x5f\x25\xd5\xec\x24\xb1\x1e\x69\xfc\x6d\x16\x98\x9a\xfc\x4f\x5e\xf5\x67\xfd\x6b\x6f\xc9\xb8\x0b\xfd\x59\xbf\xa2\x66\x2d\x84\xc6\x9d\x84\x48\x56\x25\x64\xa5\x54\x0d\x77\x5c\xe0\x30\x45\x16\xee\x62\x1d\x78\x06\xdc\x84\x3c\x7b\x12\x27\x05\x9a\x31\x0b\xbf\xbd\xd4\x5b\xd8\x58\x25\xe1\x59\xce\xf9\x5f\x7e\x41\xb4\xe7\x86\xaf\x5d\x43\xe4\xc3\x28\xc6\x53\xda\x4e\x7d\x73\xb4\x67\x8e\x59\x95\x65\x4a\xea\x64\xbe\x93\xd9\x1a\x59\x65\xde\x8c\x2e\x18\x7f\x4f\x54\xe8\xec\x01\xc4\x60\xeb\x52\x38\x99\x0b\xeb\xd1\x1e\x50\x4c\x4d\x72\xd0\xf1\x1e\x50\x49\x15\xee\xb7\x6f\x89\x0c\x9d\xab\x33\x17\xc6\xa3\xeb\x03\x88\xff\x83\x05\x0a\xbd\xde\x78\x0f\x71\x95\x88\x1b\x7a\x91\x67\x86\xfd\x45\x42\xd7\xfd\x43\x30\x9c\x83\xd3\xe2\xcd\xfd\x9c\xfc\x82\xaa\x0b\xf6\x91\xee\x37\x2b\xb7\xb8\xc0\x09\x3c\x3d\xab\x65\x26\x21\xc5\xec\x78\x02\x4f\xca\x29\xcb\x66\x4f\x40\xf1\x93\xce\xaf\xd1\x0b\x48\x16\x64\x79\xc9\xa0\x9a\xde\xae\x74\x5d\xaa\x07\x27\x27\x50\x0a\x94\x50\xab\x8d\x21\x80\x2f\xa7\x90\xf2\x80\xce\x19\x6f\x9e\x18\x18\xe4\x92\x37\xc9\x8d\x07\x8e\x66\x4a\x5c\x0d\x4d\xdc\x74\x2f\xbf\xa7\x5d\x5e\x6f\x82\xa4\xdd\x16\xa6\xe6\xb9\x46\xac\xe4\xb8\xe2\x98\x14\x53\x11\x9d\x40\xef\x2d\x25\x71\xc9\x19\x78\x9e\xb7\xe5\xfe\xee\xae\x9f\x87\x6c\x3b\x27\x57\xf2\x4d\xda\x05\x34\x8a\x2f\x35\x07\xe1\x36\x23\x6a\x5b\xad\xa2\x09\x57\x3b\xf4\x2f\x01\x94\x6a\x2d\xd2\x5a\xff\xb9\xdb\x54\xa9\x35\xed\x4b\x44\x66\x34\xca\x99\xf0\x32\x8a\x84\x6f\x49\x53\xa5\x29\x7c\xf5\x66\x96\x67\x3a\x74\x98\x6c\x23\x88\x16\x00\x39\x69\x68\x99\x65\x18\xbd\xa5\x26\xd0\x4b\x68\xd0\x56\x23\xb0\x7a\xd0\x5d\x24\x50\xf5\xdc\x95\xe7\xec\x23\x9d\xc0\x63\xfb\xc3\x39\x8b\xa2\xd6\x71\x37\xad\x05\xad\xfd\x1c\xbd\x58\xcb\xcf\x80\x9f\xb3\x28\xa5\x7f\x08\x43\x5b\x06\xfe\x14\x8e\xbe\x0b\x22\xfa\x19\xb0\x74\x91\x50\xca\xff\x10\x9e\xb6\x8d\xfc\x49\x4c\xe5\xa9\xfc\x1c\xb8\x2a\x12\xc2\x17\x7f\x8c\xaa\xb6\x0e\xdd\xce\xd7\x92\xf5\x25\x99\xc5\xf6\xb6\x72\x44\xa7\x59\x24\xc2\xcc\x3d\x26\x3e\xd5\x55\x19\xbc\xd6\xbd\x3f\x2b\xbc\x9a\x8b\x64\x49\xd4\xcb\x8c\xc9\x4b\xb6\xa4\x52\x91\x65\xec\xac\x06\x83\xe6\xec\x12\xa1\x88\xa2\x13\x38\xb3\x14\x26\xcd\xb2\xff\x3a\xba\xa7\x62\xe4\x42\xc7\xd0\xf9\x13\x56\xd0\x4c\x19\x5d\x18\xd9\x56\x60\x66\xf9\x74\xb4\x8f\xcd\x45\xfe\x35\x1c\xda\xb8\x93\xdf\x93\x40\x60\xc1\x56\x94\xc3\xff\x71\x96\x01\x8d\x85\x1f\x82\x2a\x03\x41\x9c\x30\x8c\xa2\x17\x05\x31\x12\x09\xbe\x00\xc2\x81\x64\x4c\x7a\xf0\x4e\x19\x08\x39\x29\x00\xe0\x04\x3e\x84\xe1\x64\xb9\x9c\x48\xf9\x01\x23\x1d\xa9\x0f\x36\x09\x2c\x53\x74\xcb\x11\x05\x31\x87\xb3\x5c\xad\x0d\xbc\xe0\xd1\x1a\x3e\x20\xb8\x50\x21\x4d\x6e\x99\xa4\xd5\x90\xcb\x26\xdf\x6a\xfc\x8f\xb1\x56\x00\x53\xe0\xf4\x16\xbe\x25\x8a\x62\xd2\xf0\x48\xd7\x6d\x06\xd5\xe2\x05\x26\x1d\x81\x8e\x58\xa8\x2f\x78\x20\x9d\x41\xf5\xb1\x4e\x07\x3d\x25\x2e\x54\xc2\xf8\xc2\x19\x78\x31\x09\x2e\x14\x49\x94\x73\xea\x42\x7f\xd4\x1f\x54\xeb\x3f\x12\xfe\x02\x67\x18\xc4\x8d\x5a\x2a\x3e\x52\xda\xca\x3c\x38\x54\x18\x16\x53\x79\x2b\xd2\x44\x3a\x83\xce\x51\xcb\x98\xcb\x65\x81\xf9\x3d\xe3\x29\xea\xcc\x61\xb8\x45\x79\x22\x84\x63\xe8\x4f\xfa\x70\x8c\x94\x8a\x9f\xf9\x44\x1b\xd1\x2e\x25\x6a\x49\x62\x4b\xda\x9e\x17\x39\x17\x91\x98\x91\x08\xf3\xb3\xe1\x92\x64\xcd\x3a\x25\x1a\x0c\x9d\x29\x38\xbe\x48\xb9\x72\x31\x77\x7b\x8d\xbf\xec\x85\x4b\x6e\xca\x8f\x1a\x16\x86\x5b\x60\x4b\x31\x99\x44\x4b\xc1\x1b\x5d\x03\xdb\x5a\xd4\xe9\xd9\x13\x38\x01\xe7\xf4\x09\x3c\xd2\x44\x6b\xa9\xc7\x16\x76\x81\xb0\x8f\x47\x08\x3b\x7e\x72\xb6\x07\x78\x86\xc0\xa3\x6f\x34\xf0\xd3\xc7\x3b\xe0\x3a\x74\x9c\x26\x71\xed\x9c\xa2\x34\xb3\xc7\x4f\x35\x81\xd3\x53\xfb\x68\xe5\x79\xfd\x97\x81\x1c\x75\x41\xea\x49\x3d\x39\x43\xc8\x6f\xfe\x6a\x07\xcc\x25\xff\x41\x9f\xb1\x7c\x75\x97\x6c\x5c\xf8\xea\x6e\xa1\x3f\x67\x1b\x17\xc6\x83\x0f\xbf\x42\x15\xf7\x3e\x95\x19\xd3\x65\x5b\xd4\x67\xda\xcf\xf5\xbf\x45\xa8\xa9\x06\xbe\x7a\xdc\xd5\x00\xf0\xbf\x6a\x2d\x0b\xc0\x27\xd7\x1d\x24\xd9\xf7\x8c\x9b\x9e\x0d\x5d\xfd\xf0\x32\x6f\xc9\xf8\x79\x07\x38\xc9\x6a\xe0\x24\x6b\x21\x9f\x8b\x98\xfa\xea\x56\xa3\xcc\x66\x22\xf3\x74\xc6\x09\x43\x70\x34\xa5\x13\x3d\x7e\x67\x9d\x06\xf1\xc3\x1d\xbe\xc9\x4c\x91\xc0\x76\x0e\x6b\x9c\x03\x9c\x40\xf9\x86\x26\xda\xba\x8a\x62\x1b\xc1\x14\x4e\xde\xf1\x39\xe3\x4c\xad\x5b\xc0\x35\x27\xb7\x4d\x32\xff\xd6\x6c\xed\x3a\x05\x2f\x91\xd6\xc5\xd9\x25\xc9\x9c\xe2\x9e\xbb\xbd\xe5\xe9\x96\x1d\x87\xa7\x51\xe4\x02\x12\x1d\xb4\x95\x02\xda\x96\xa1\x55\x61\xd7\xbc\xb3\xfe\x49\xfb\xd3\x8c\xed\x39\xa4\xd7\xcd\x24\xef\x85\xac\x16\x72\x8a\x6a\x5c\xc6\xae\x0f\xab\xe8\xe4\x24\xf4\xd7\x89\x11\xf1\x79\x7b\xf1\x48\x4f\xae\x3a\x57\x17\xd6\xfb\xa6\x5a\x2a\x0d\xe7\x0c\xd5\xb2\xc0\x49\x5e\xad\xbb\x8a\x64\x85\x33\xca\xd1\x2c\xce\xa8\xa5\x4b\xe8\xe2\x86\xc5\x40\x97\xb1\x5a\xeb\xb8\xc2\x1c\x73\x53\xb9\x17\xb5\xe8\x88\xea\x82\xd9\x1c\xb0\xce\xb5\xe1\x6a\x57\xbb\x4b\x7f\x5d\x08\x67\xd0\x3d\x9e\xaf\x32\x6f\xeb\x6b\x60\xda\xe1\x77\x0e\xa3\xf3\x23\xf5\x95\x83\xf2\x76\xf5\x34\x5d\x23\x74\xf3\x15\x1e\x40\xc3\x44\xe9\x0f\xa5\xb2\xe9\xd8\x1e\x0f\xab\x95\x69\x1f\x16\xa5\x4b\x7e\x12\xb2\x45\x18\xa1\x51\x31\x9d\x24\xa1\x58\xd1\x84\x06\x90\x01\xe3\x01\xcd\xaa\xee\xdf\xa0\xbc\x2d\x30\x1e\xd4\x20\x72\x36\x76\xa1\xf8\xaf\xfb\xca\x5b\x9b\x41\x70\x26\x2e\x84\xd1\x6b\x3d\x2a\xfe\xba\x60\x1f\xe9\x6b\x1c\xcb\x05\x3f\x4d\x92\x77\x41\x66\x7e\x5c\x8a\xb8\xbd\x63\xa4\xae\xfd\x48\xf6\xde\x2f\x6e\x14\xb3\xb8\xdf\x5b\x01\xa5\x19\xef\x47\x6c\x6f\x31\x29\x06\xef\x68\x31\x69\xed\x30\xd9\xe2\xb6\x76\x98\x7c\xe6\xed\x23\x45\x2d\xb6\x37\x1e\x8d\xfe\xf2\xe0\x36\x90\x1a\x3f\xb7\x62\x79\x18\x4b\x77\xe8\x7f\x56\xae\x0e\x87\x70\xf9\xc3\xb7\x3f\x4c\x20\x10\x54\xf2\xbe\x82\x39\xa5\x51\xde\x14\xe6\x79\x70\x4b\x21\x10\x10\x92\x15\x85\xbe\x39\x0e\xd3\x9c\xc1\x75\x17\xc7\x3b\x7d\x1b\xcd\xb2\xdf\x70\x21\x14\xb7\x14\x37\x1b\xda\x95\x79\x1a\x45\xb9\x2c\x81\x44\x91\xb8\x95\xa0\x3b\x49\x88\x4f\x4d\xe7\xa9\x0a\x09\x87\x23\x6b\x33\x2c\x51\x1e\xfc\x83\x42\x9c\x88\x19\x99\x45\x6b\x08\x04\xce\x58\x52\x0a\x84\xaf\x21\xa1\x24\x82\x80\xcd\xe7\x34\xa1\xdc\xa7\x30\xa3\x3e\x49\x25\x85\x5b\xda\x4f\xa8\x9d\x22\x89\x6e\x18\xe6\xad\x33\x91\x2a\x20\x28\x04\x3f\xa1\x31\xe1\xfe\x1a\xd3\x51\x29\x96\x14\xe6\x09\x31\x46\x44\xcc\x81\x40\xcc\x32\x1a\x79\xf0\x36\x5f\x90\x8d\x28\x9b\xeb\x75\x06\x74\xc5\x7c\xfa\x1e\xe1\x7f\x24\x8a\x89\xbc\xfd\x58\x02\xe1\xc1\x50\x24\xc8\xd9\x8f\x42\x2c\x81\x71\x17\xde\x81\x4c\x65\x4c\x7d\xdd\xb5\x67\xa3\x99\xf3\x4b\x18\xd2\x5b\x4b\x8d\x73\xa7\xc6\x7a\x18\x66\xc3\xad\xe6\x08\xcd\x88\xaf\xa2\x35\x2c\x89\xc2\xcc\x3d\x24\xca\x46\xb4\x4e\x8d\x06\x25\xb1\x79\xad\xfb\x6f\x3c\x1a\xc1\x10\x5a\x94\x01\x8e\xa1\x67\xdb\x9c\x07\xf9\x80\xc7\xf5\xee\x2e\xbd\x7b\x1b\x36\xbb\xd2\x57\x54\x98\x35\x4b\x2b\x4f\x0d\x6e\xbb\x57\x3b\x5b\x90\xb6\xdc\xd8\xd3\x84\xf4\xd0\x97\x26\xaa\x66\x78\xef\xdb\x35\x0d\x13\xd3\x89\xb1\xb1\x72\xe1\x61\xef\xae\xb4\x4e\xb4\xed\x3d\x9a\xae\xa9\xda\x70\x36\x83\x87\x75\x32\xe9\x90\xd6\x38\x7d\xf8\x72\xba\xeb\x61\xf2\x58\x90\xc1\x2f\xbf\x14\x71\x40\xf5\x99\x12\xb1\x2d\xf8\x2d\xe8\x54\xa9\x9c\x5b\xe1\x2e\x75\x0f\x76\x99\xa2\x05\x4e\xc7\xaf\x41\x33\x35\x2c\xa7\x65\x2d\xa9\xa5\xc1\xbd\xcd\x8f\x5a\x9d\x5a\x9e\x18\x64\x03\xfc\xac\x19\x95\x36\x2a\x8d\x3e\xf9\x7c\xa1\xd7\x6d\x08\x79\xbb\x7a\xa3\x47\xbd\x97\xf5\x06\x70\x62\x66\x65\xc9\x6c\x6a\x4a\x62\xef\x14\xfb\x7f\xdd\x2a\x56\x8a\xe6\x71\xb0\x41\xb3\x4f\xac\x85\x68\xc1\x91\x12\x01\x7d\x6b\xd7\xe9\xde\xed\x80\x74\x50\x4d\xf8\x22\xa2\xb9\x09\x3b\xfa\xed\x13\xec\x3c\x8f\x29\xa4\xb0\x7d\xd3\xa5\xac\x3d\xe6\x95\x97\x36\x81\x48\xdf\xe8\xa5\x5e\xf5\x3c\x12\x22\x71\xf6\x4e\x04\x4e\xf4\x90\xbb\xd7\x68\x46\xad\xe4\x15\xd2\x36\x63\x3c\x32\x2b\x3f\x64\x43\x77\x74\x02\xda\xc5\x6b\xba\x02\x4b\xcf\x54\xd6\x29\xf8\xfa\x80\xfb\x65\xdf\x92\xe3\xec\x29\x56\x89\x58\xe5\xf5\x77\xd3\x44\x6e\xb5\x80\x95\x58\xcf\xbc\x03\xde\x52\xc0\xca\x4f\xb5\xdb\x93\xeb\xac\xd1\x5e\xd5\xa8\x0d\x74\x42\x6c\xf6\xbf\xb5\xe8\xfe\x61\x3d\x95\xbb\xce\x80\x53\x6b\x67\xc0\x45\x29\x16\x84\xb7\xa6\x70\xfc\x49\x9d\x02\xe6\xc8\x7f\x02\x77\xe5\xce\xb5\x4a\x9b\x9b\xbd\x95\xa0\x56\xb4\x76\x6b\x45\x53\x5b\x4a\x6b\x6f\x13\x58\xa4\x4a\xd1\xa4\xc1\xdf\x35\x06\x47\xf6\x0e\x3a\x7b\x5f\x41\xeb\xa9\x51\xd6\xbf\xe7\x71\x53\x4c\x54\x88\x1a\xad\x55\x59\x97\xd3\xda\x4f\xff\xda\xb8\x56\x9f\xc4\xda\x72\x74\x55\x9c\x77\xee\xcc\xab\x25\xd8\x1b\x3e\xea\x3c\xf2\x2a\x75\xc5\xb4\x16\xbe\xda\x9b\x64\xea\xa1\xc0\x01\x0d\x33\xb6\x62\x73\x77\x07\x74\xa3\xd7\x2c\x8f\x27\x90\xb3\xfb\x8a\x68\xf9\xa9\x40\xf1\x0a\x50\xcf\xa4\x64\xe5\x6e\x69\x93\x8c\x95\xfa\xa5\x91\x19\xe6\x46\xfb\x3b\x42\x2d\xe4\x3f\xa0\xb3\x9a\x98\xbc\xef\xab\xbb\xdc\x77\xe1\xd5\x46\x8f\xb2\xbd\xa5\x44\xbc\xc9\xdb\x9c\xb6\xf7\xf4\xe5\xa6\x08\xec\xb7\xb7\xcd\xf5\xe6\xc3\xfe\x89\xe8\x4a\x78\xdd\x4d\xe5\xee\xee\x5d\x90\xe5\xaf\x80\xea\xf7\xdc\xba\x69\x21\xa1\x9f\xad\x64\xb6\xef\x87\xf6\xd7\x7b\xa9\xa0\xa4\x7e\x86\xe7\x53\x18\x61\x72\xff\x33\x3c\x6b\xcb\x52\x0e\xa9\x83\x1a\xed\xab\x13\xb8\xfa\xf9\xfa\xa1\xf5\xcd\xcd\xd1\xe1\x77\xad\x4e\xab\xd2\xb1\xa5\x37\x94\xe9\xd8\xd2\xed\xc9\x27\xba\x41\xb9\xff\xe2\x45\xbf\xc5\x83\xec\x2b\x10\xe2\xbf\x47\xc3\xcf\xdb\xea\xe8\x92\xed\x9f\xd7\xe0\xd8\xde\x59\xef\xd2\x65\x96\x67\x38\xfa\x0f\x18\x7c\xfd\x35\xdc\xc7\x0a\xfd\x56\x3b\x2a\x4f\x29\x9e\x98\x46\xc3\x7b\xef\xa7\x2a\xfa\xe7\xb8\x9b\x7e\x93\xed\x73\xef\x26\x97\xb6\x40\xe0\xcf\xd7\x00\xb3\x7e\xe8\x0a\x76\x7f\xc2\x89\x0d\x3a\xff\x8e\xc8\xf3\xd6\x62\xd4\x49\xf5\x1d\x9e\x16\xbd\xe8\xf7\xcf\xef\xaf\x4c\xa5\xce\x9c\x86\x8f\x60\xd7\xbb\x5e\x1d\xdb\xdf\x1b\xf9\xf4\xf6\x9d\x86\x46\x87\x94\xc4\x13\xd3\x9d\xbb\x43\x9a\xad\xf5\x7c\xcc\xfd\xd3\x9c\xc2\xd1\xc6\xc1\x5d\xff\x9f\x00\x00\x00\xff\xff\x3c\xa5\xd1\x02\x4f\x4b\x00\x00"),
		},
		"/stats.js": &vfsgen۰CompressedFileInfo{
			name:             "stats.js",
			modTime:          time.Date(2020, 10, 3, 0, 1, 2, 704539327, time.UTC),
			uncompressedSize: 3951,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xa4\x57\x5f\x6f\xdb\x36\x10\x7f\xf7\xa7\xb8\x3d\x0c\x95\x1b\x57\x96\x93\xee\x25\xae\x07\xb4\x29\xd6\x15\x58\x51\x60\xee\x9b\x61\x0c\xb4\x74\x8a\xb8\x49\xa4\x41\x52\x69\xdc\x22\xdf\x7d\x38\x92\x92\x49\x59\x6a\x30\xcc\x08\x62\x8b\xf7\x97\xbf\xfb\x1d\x79\x5a\x2e\x41\x1b\x66\x34\x54\xb2\x2e\x34\x98\x0a\xa1\x60\x86\x01\x13\x05\x94\xad\xc8\x0d\x97\x02\x8c\x84\x46\x16\xbc\x3c\x01\x37\xe9\xec\x81\x29\x6f\xb3\x81\xa4\xd7\x49\xe6\xf0\x7d\x06\x00\x40\xe2\x06\x36\xf0\xfd\x69\x3d\xb3\x0b\xb9\x14\xda\x00\x2f\x1e\xbf\xf0\x06\xc9\x28\x5b\xc7\xeb\xbf\x23\x3b\xbe\xad\x6b\x99\xc3\x06\x56\x23\xb2\xed\x89\xac\xae\x47\x24\x1f\x8b\x1a\x61\x03\x37\x63\x22\xd1\x6a\x92\xbd\x8e\xb2\x10\x6d\xb3\x45\xc5\x51\x93\x8a\x93\x9e\x85\x46\x1a\x56\x3b\x31\x65\x02\x57\xb1\xfa\x1a\x96\x4b\x30\x76\x0f\x57\x20\x4d\x85\x0a\xb4\x15\xce\xfa\x7d\x5b\xe8\x36\x1e\x08\xfa\x38\x85\x5b\x10\xf8\x15\xde\x2a\xc5\x4e\x49\x10\x64\xbe\x20\x8f\x5f\x3e\xbf\xff\x7c\x0b\x0a\x05\x6b\x90\x90\xb6\x11\x5c\x12\xbd\x9b\x9a\x69\xf3\xe1\x2e\xf2\x33\x5f\xcc\x7a\xf1\xe1\xb4\xe5\xdf\x30\x94\xbe\x76\xbe\x77\xd9\x1e\x1a\x2e\x60\xb7\xda\x43\xc3\x1e\x61\x77\xbd\x87\xbc\x66\x5a\x83\xe6\xdf\x50\xc3\xee\x66\x0f\x35\x7f\x40\x90\x87\xbf\x31\x37\xda\x7a\xec\xea\xd6\xa4\x5c\x70\x03\x9b\x33\x0f\x92\x43\x5b\xd6\x28\xe6\xc1\x06\x1d\x72\xf8\x68\x14\x7b\xd7\x96\x25\xaa\x3b\x76\x64\x39\x37\x27\xaa\x58\x66\x21\xbb\xce\x7e\x06\x59\x3a\x1d\x48\x8e\x0a\x19\xd5\x9a\x19\x2c\xe6\x70\xb0\x36\x16\xb7\xa3\xe4\xc2\xe8\xd9\xc0\xf3\xa1\x2d\x73\x5b\x29\x17\x1a\xae\xba\x24\xe0\xe5\x58\xd4\x39\x2c\x61\x95\xb9\xb8\xa2\x6d\x0e\xa8\x28\x34\xcb\x4d\xcb\xea\x28\x4a\x1f\xa6\x94\x0a\x92\x1a\x0d\x70\xcb\x4c\xe0\xf0\x26\x24\xc2\x1a\xf8\xd5\x55\xb8\x61\xfa\x90\xa3\xd4\xd5\x68\xc7\xf7\xb0\xb1\xc0\xbb\x44\x7c\x76\x0b\x9f\xf8\x7c\xdd\x1b\x3e\xfd\x30\xa4\x75\xe9\xea\x98\xd6\x28\xee\x4d\x35\x1d\xd9\xa9\xfd\xd7\xc8\x5d\x59\x97\x4b\xc7\x11\x02\x86\xda\x9d\xb8\x05\x0a\x6b\x7c\x60\xc2\xc0\x87\x3b\x47\x40\xcf\x00\x4f\x3c\xd8\xb8\xc8\xfe\xf1\xec\xe9\x4e\x0a\xc3\xb8\x00\x2e\x0a\x7c\xc4\x22\x24\xd7\x02\x4c\xc5\x35\xd0\x9f\xe0\x86\xb3\x9a\x7f\xc3\x02\x58\x69\x50\x81\xc2\x1c\x8f\x96\x52\x3e\x89\x92\x2b\x6d\xa0\x41\xad\xd9\x3d\xa6\x3e\xb8\x75\xb6\xb5\x44\xdd\x84\xcc\x8f\xf8\x79\x17\x2a\x05\x4c\xb5\x18\x85\xf0\x8d\xa1\xfe\x3c\xe0\x61\x16\xe9\xb1\xd5\x55\xd2\xa3\x9f\xda\x08\x17\x30\xbb\xe4\xfa\x4c\xda\x63\xc1\x0c\xfe\x61\x81\x4b\x1a\x6c\xb6\x74\x72\x5e\x76\x90\x60\x42\x7e\x91\x5b\xcc\xa5\x28\xec\xe9\x93\x65\x19\xbc\x0c\xbf\xce\x81\x68\x13\xd4\x97\x9f\x98\xa9\xd2\xb2\x96\x52\xf5\x8e\xd3\x4f\xd8\xa4\x2e\x18\x2c\x63\xa7\x1d\x6c\x9d\x07\x57\x4b\x8f\xec\x7b\x66\x30\x31\x3e\x52\xa8\xc9\x4b\x48\xbc\xe6\x4f\x31\x09\x76\xe1\x83\x87\x10\x5e\xc1\x6a\x3f\xca\xd9\x4e\xcf\x42\xe8\x1e\x22\xe8\xfa\x9f\xcb\x25\xfc\x89\x8d\x7c\x40\x28\x95\x6c\x7a\x86\x12\x09\x99\xa5\x2d\xad\x58\x8a\x1a\xd6\x1c\x35\x7c\xad\x78\x5e\x01\x53\x08\x47\xc5\xa5\x02\x23\x43\x57\xa4\xdc\x70\xc1\x9b\xb6\x39\x1b\x01\x17\xf0\xc2\xb5\xef\x8b\x34\xc2\xa4\xa1\xc3\xa1\x23\x7b\xd7\xe0\xfe\xe2\xda\xa7\x7f\x1d\xda\x72\x97\xed\xd7\x43\x93\xfb\xfc\x7c\xa9\x5d\x32\x6d\x01\x62\xd0\x3e\x67\xbe\xc1\x1b\x10\xa3\xb4\x23\xd8\x23\xb0\xf9\x1e\x7e\x75\xe9\x0d\x55\xed\x05\xa0\x90\xfd\xb3\x8e\x96\x9f\x62\x1a\xfb\x2c\xf9\xd8\x79\x14\xe5\xa6\x8f\x35\xcf\x31\xc9\x16\xce\xc6\x17\xe9\xa9\x6b\x39\xaa\xdf\x7b\x77\xcb\x9d\x9b\xcd\xe8\x05\x8c\x51\x7b\x1c\x46\x4b\x01\xa3\xe7\xe7\xfb\x94\x8a\x32\x65\xd4\x0f\x07\xde\x30\x62\x7a\x2f\x0c\xb8\x34\x62\xbf\x3d\xe9\x29\xeb\xed\x49\xff\xd8\x96\x46\x8c\x29\x63\x92\x3d\x63\x4d\x53\xc8\xa4\x39\x09\xc3\x66\x23\xb5\x77\xf6\x78\x89\x95\xdd\x5a\xa8\x39\x7e\xaa\x0c\x6a\xe5\x7b\x32\xac\x54\x58\x1d\x85\xa6\x55\x62\xa2\x48\xce\x36\x19\xba\xd4\x44\x8e\xc8\xa3\xe0\x06\x9b\xa8\xea\x34\xd3\xf0\xc6\xcd\x91\xf1\x99\xe5\xe7\xbf\xf1\x88\xd6\x75\xe7\x6e\x1d\x1d\x08\x76\x58\x3b\xd6\xd2\x5c\xfa\xac\xdc\x1c\x77\xbe\x1d\xa2\x91\x2d\x28\x0e\x29\xd2\x38\xb4\x71\x69\x4c\x74\xeb\xca\x36\xe5\x66\x38\xf8\x8d\x74\xa8\xf5\x67\xaf\xe0\x78\x1a\x18\xee\x63\xfc\x88\x73\x15\xd5\xe4\xc5\x34\xec\x78\xb9\xaf\xc3\xe9\xf2\xe2\xfb\x25\x70\xe7\xe5\xc1\x8e\x2e\x44\xab\x3e\x39\x7f\x65\x65\x93\xc9\x75\x26\xd7\x43\x93\xd5\xb3\x26\x37\x43\x93\xeb\x67\x4d\x5e\x0f\x4d\x6e\xa6\xab\xef\x49\x7a\x89\xfd\xad\xfd\xbf\x88\xd6\x7d\x80\xdb\xee\xc7\xe2\x99\xbb\x39\xe8\xb7\xcb\x79\x81\xaa\x60\x07\x99\x8f\xa2\xe0\xf9\xa0\x14\xd1\xe0\x30\x8f\x6f\x84\x5c\xb6\xee\x16\x99\x56\xff\x9f\x43\x89\x1b\x19\x28\x39\x1a\x8c\xbb\x79\x24\x3e\xff\xc3\xd4\x1d\x4f\xf9\x7a\xe0\x84\xd2\x74\x22\x52\x4e\x3f\xd9\xb9\x5c\xc3\x2b\xf7\xf8\x9b\xc2\xb0\x4d\x62\xfe\xd2\x4b\xcb\x80\x5b\xf6\x25\x83\x5e\x1b\x89\x45\x5c\xd0\xd5\x2b\xd5\x09\x04\x3e\xa0\x82\xbc\x62\xe2\x1e\x41\x4b\x5a\x56\x08\xa1\x2f\x5d\xc9\xb6\x2e\xe0\x80\x20\x24\x08\xc4\x82\xde\x82\x14\xe6\x0a\x99\x41\xe0\x06\x90\xe5\x95\xe5\x78\x3a\x1b\x9b\x83\x33\x7f\xbe\x66\xc3\xa3\xf8\xcc\x61\x2b\x0f\x11\x09\xe6\x95\x09\xab\xeb\x11\xab\x09\xd5\x1b\xaf\xea\x10\x8d\x8f\x4c\xcf\xdf\x66\x3d\x7b\x4a\xe6\xf3\xf5\xbf\x01\x00\x00\xff\xff\x7f\xc5\xc2\x07\x6f\x0f\x00\x00"),
		},
		"/statsviz.css": &vfsgen۰CompressedFileInfo{
			name:             "statsviz.css",
			modTime:          time.Date(2020, 9, 24, 21, 23, 45, 192276676, time.UTC),
			uncompressedSize: 1027,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\x9c\x52\xd1\x6e\x9c\x30\x10\x7c\xe7\x2b\x56\xa9\xaa\x48\x51\x4c\xb8\xab\xd2\x07\xdf\xd7\xf8\xf0\x02\xdb\x2e\x36\xb2\xf7\x02\xa7\x2a\xff\x5e\x19\x0e\x8e\x72\xf4\x25\x0f\x58\xd8\xde\xd9\x19\xcf\x6c\x23\x2d\xbf\xc2\xd9\xdb\xeb\x2b\xe4\x75\x20\xab\x4a\xef\xc4\x90\xc3\x00\x7f\x32\x00\x80\x06\xa9\x6e\x44\xc3\xa1\x28\xbe\x9f\xc6\x93\xd6\x84\x9a\x9c\x86\xe2\x94\x7d\x66\xd9\x3e\xca\x52\xec\xd8\x5c\x35\xa4\xdb\x09\x36\xd6\x09\xb6\x1d\x1b\x41\x55\x7a\xbe\xb4\x2e\x6a\x68\xcd\x30\xa2\xd1\x09\x1c\x8b\xa2\x1b\xf6\xaa\x83\xef\xa3\x86\x43\x15\xe6\xef\x56\x64\x3a\x0d\x87\x6e\x48\xdf\x1e\xcc\x04\x34\x51\xc3\x53\x0e\xf9\xd3\x7a\x1d\x85\x7f\x6b\xd0\xd8\x45\xb0\xe0\x20\xca\x30\xd5\x4e\x43\x89\x4e\x30\x4c\x45\xdb\x87\xf5\x64\xa5\xd1\x70\x38\x1e\x17\xa9\xb3\x1f\xe9\x04\xcc\x45\xfc\xe4\x4b\x24\x57\x33\xaa\x28\x46\xb6\xa6\x54\x8c\x37\xec\xaf\x4b\x14\xaa\xae\xb3\x03\x1a\x62\x67\x4a\x54\x67\x94\x1e\xd1\x9d\xd6\x94\x3f\xb7\xfe\xe7\xef\x01\xdb\x91\xcb\x72\x4e\x8e\xc9\x21\x58\xbb\xe5\x9a\x2e\x1e\x92\x4b\x5b\xc6\x41\xb6\x70\x6d\x2a\x79\x4c\xf1\xcc\xbe\xfc\x3d\x81\x16\xa5\xcf\xcf\x5b\x6e\xd9\xe7\x56\x2b\x74\x4b\x4e\xcd\x16\x4e\x61\xaf\x5b\xa8\x64\xcc\xff\xcd\x4a\x7f\xaa\x62\xdf\x6b\x08\xbe\x5f\x9d\xf5\x21\x0d\x42\x5a\xff\x71\xec\xc7\x7d\x9c\xde\x5e\x20\xa2\x80\x34\x08\xf7\x44\xc7\xb2\x97\xb7\xb1\xc0\x7f\x60\x98\x5a\x7f\x50\xa4\x33\xe3\x9e\xb2\xe5\x85\x69\xa7\xa1\x80\x02\xde\xe7\x54\xc6\x01\xba\x77\x41\x66\xea\x22\xc5\xd3\xa6\x7b\x43\xd6\xa6\x60\x77\x9a\xcf\xd1\x4d\x29\x29\xc6\x4a\xf4\x6d\x9e\xb6\x13\x9a\xee\xbe\x44\xfb\x20\x3e\xfb\xfc\x1b\x00\x00\xff\xff\x70\x11\x3b\x47\x03\x04\x00\x00"),
		},
		"/ui.js": &vfsgen۰CompressedFileInfo{
			name:             "ui.js",
			modTime:          time.Date(2020, 10, 3, 0, 10, 38, 920506910, time.UTC),
			uncompressedSize: 2440,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xb4\x55\xcd\x8e\xdc\x36\x0c\xbe\xcf\x53\xb0\x87\xc5\x78\x10\xc7\x1e\xa7\x0d\xd0\xda\x99\x1e\xba\x45\x37\x05\x52\x60\x81\xed\x6d\xb0\x07\xad\x4d\x8f\xd5\xca\x92\x20\xd1\x89\x8d\xcd\xbc\x7b\xa1\xb1\xc7\x91\x7f\x52\x14\xd8\x46\x27\x8b\xfc\x48\x7e\xa4\x48\x3a\x8e\xa1\xe1\x50\x29\x51\x58\xa0\x0a\xa1\xb1\x68\x80\x4b\x42\x53\xb2\x1c\xc1\x12\x23\xdc\x7c\x64\xc6\xa1\x0e\x10\x94\x8d\xcc\x89\x2b\x09\xc1\x0e\x9e\x37\x00\x00\x4e\x57\xc3\x01\x9e\xcf\xd9\xe6\x22\x10\x48\xa0\x59\x63\xb1\x80\x03\x94\x4c\x58\x1c\x14\x75\xc4\xed\xfd\xa8\xf0\x1d\x81\x41\x6a\x8c\x1c\xcc\x32\x38\x0f\x78\x52\xa7\x93\xc0\x8b\xcd\xc2\x64\x0c\xf1\xdd\xdc\x4a\x0b\x45\x16\x0e\x20\x1b\x21\x86\xd0\xa3\x69\x85\x4c\xff\xca\x88\x05\x05\x23\x76\x4d\xc1\x9d\x81\xc1\x71\x14\xb8\xf3\x3c\xb9\xb9\xd3\xa6\xe0\x2c\x23\xe7\xe7\xb8\x7f\x0c\x17\x80\xce\x07\x24\x2b\x00\xea\x34\xa6\xb0\xb5\x39\x23\x42\xb3\x5d\x02\x24\xab\x1d\xc0\x39\x00\x26\x84\xca\xb7\x13\xc8\x39\xfc\x5f\x29\xbe\x79\x29\x45\xdb\xd9\x6f\x4a\xf0\xfb\x97\x12\xe4\x85\xc0\x6f\xca\xf0\x87\x17\x33\x94\xaf\x1b\xfb\x75\x8e\x8f\x97\xaf\x73\xdf\xc9\x71\x0c\x15\x91\xb6\x69\x1c\xbb\x3e\x17\x5d\x94\xab\x3a\xfe\x8b\x7d\x64\x36\x37\x5c\x53\x6c\xb0\x44\x83\x32\xc7\x58\xb0\x4e\x35\x34\xce\xa4\x0b\xf5\xe1\x22\x72\xd3\x3a\x7a\x27\x4e\xc2\x31\x79\x8f\x4c\x7b\x3c\x5b\xd6\x72\x9b\xce\x6a\x73\xc5\x12\xaf\x71\x96\x13\xf1\xfc\xef\x52\x99\x9a\x51\x0a\xdb\x9b\xf7\xe9\xcd\x1f\xe9\xcd\x83\x87\xf1\xf2\xe9\xfe\xcd\xf5\x53\x47\x68\x57\x7c\xdb\xa6\x2c\x79\x9b\xc2\xf6\x97\x99\x32\x8e\xa7\xb1\x61\xa6\xc7\x56\x2b\x89\x92\x46\xc0\xc3\xef\x3e\xad\xbe\xb6\xe3\x86\xca\x0d\x32\xc2\xfb\x61\x83\x7c\xd9\x38\x4a\x93\x0d\x2f\x6f\x1e\x02\x0a\xb2\xfe\xea\xb8\xef\xdf\xc1\x3d\x47\x70\x79\xd0\x6d\x38\xdb\x33\xa1\x57\xfc\x5d\x76\x7d\xce\xe9\x66\x2a\x18\xe1\x6f\x46\xd5\x7f\xf2\x1a\x2d\xb1\x5a\x07\xd3\x28\xfd\x82\xf2\x5b\x61\x34\xbd\xbb\xfd\xc0\x25\xda\xc5\x4e\xcb\x95\xb4\x04\xa7\xdc\xa5\xe2\x16\xb9\x8d\x04\xb3\x74\x77\x6b\xb3\x19\xa4\xe6\xf2\x92\xaf\xdf\xf4\xc7\xfd\xe3\x02\xc6\xda\x25\xcc\xbf\x44\x02\xe5\x89\x2a\x78\x0d\xc9\xe3\x50\xd2\x6b\xf7\xd9\x8a\x69\x74\xc6\x47\x5f\x53\x2a\x03\x81\x53\xbb\x9f\xcb\x3e\x04\x09\x07\xc7\x77\xf0\x93\x01\x87\x77\x20\x33\xe0\xaf\x5e\xed\x66\xfd\xe2\x8c\x8a\x1e\x7d\xe4\x1e\xd3\xa1\x25\x6e\x05\xab\x35\xdc\xdd\x82\x6b\x55\x0b\x9f\x2a\x9e\x57\xc0\x0c\x82\x6b\x7f\x55\xc2\x93\x6a\x64\x61\x27\x56\xbc\x84\xa0\x80\x77\x43\x31\x3e\x7f\x86\x02\x7e\xee\x53\xde\xad\xac\x88\x5c\x49\xe2\xb2\xc1\x69\xe4\xf3\x66\x72\xed\x93\x8e\x74\x63\xab\xe0\xf9\x6b\x3b\x42\x70\x89\x2b\x0b\xa2\xdd\xa7\x50\xac\x88\x93\x55\x71\x67\xb0\x4c\x61\xab\x99\x5e\xdd\x36\xdd\x3e\x85\xfd\x8a\x38\x49\x21\x59\x8a\x1d\xa3\x74\x25\xe7\x3e\x6f\xa1\x4c\x0a\x5b\x73\x7a\x0a\xde\xbe\x0d\x21\x79\xf3\x63\x08\xc9\x4f\xc9\x6e\x25\xaa\x3b\x9f\x78\x41\xd5\x6a\x14\x77\x0a\x66\x2b\x57\x02\x25\x4f\xee\xb3\x50\xb4\xe2\xe6\x3c\x2d\xf1\x6e\xb3\x54\x0c\x7f\xef\xbe\xde\x99\x3f\x24\x75\xd4\xe8\x62\x75\xa4\xdb\x87\x9c\x09\xec\x87\xda\x7f\xe1\x2f\xa3\x1a\x8d\x4d\x3b\x99\xb0\x6c\x3e\xf9\x06\x59\xfe\x9f\x46\xdf\x27\x36\x50\xae\xb3\xcd\x39\xd8\xed\xb2\x7f\x02\x00\x00\xff\xff\x81\x18\xe2\xd3\x88\x09\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/app.js"].(os.FileInfo),
		fs["/buffer.js"].(os.FileInfo),
		fs["/index.html"].(os.FileInfo),
		fs["/opts.js"].(os.FileInfo),
		fs["/stats.js"].(os.FileInfo),
		fs["/statsviz.css"].(os.FileInfo),
		fs["/ui.js"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}