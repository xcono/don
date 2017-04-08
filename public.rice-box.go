package main

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    `style.css`,
		FileModTime: time.Unix(1491620829, 0),
		Content:     string([]byte{0x62, 0x6f, 0x64, 0x79, 0x20, 0x7b, 0xa, 0x9, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x3a, 0x20, 0x30, 0x3b, 0xa, 0x9, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x3a, 0x20, 0x73, 0x61, 0x6e, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x69, 0x66, 0x3b, 0xa, 0x7d, 0xa, 0xa, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x20, 0x7b, 0xa, 0x9, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x3a, 0x20, 0x23, 0x32, 0x32, 0x32, 0x3b, 0xa, 0x9, 0x70, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x20, 0x2e, 0x35, 0x65, 0x6d, 0x3b, 0xa, 0x9, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x2d, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x3a, 0x20, 0x31, 0x65, 0x6d, 0x3b, 0xa, 0x9, 0x74, 0x65, 0x78, 0x74, 0x2d, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x3a, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x3b, 0xa, 0x9, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2d, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x3a, 0x20, 0x35, 0x70, 0x78, 0x20, 0x73, 0x6f, 0x6c, 0x69, 0x64, 0x20, 0x23, 0x61, 0x66, 0x61, 0x3b, 0xa, 0x7d, 0xa, 0xa, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x20, 0x3e, 0x20, 0x68, 0x31, 0x20, 0x7b, 0xa, 0x9, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x3a, 0x20, 0x30, 0x3b, 0xa, 0x9, 0x74, 0x65, 0x78, 0x74, 0x2d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x20, 0x75, 0x70, 0x70, 0x65, 0x72, 0x63, 0x61, 0x73, 0x65, 0x3b, 0xa, 0x7d, 0xa, 0xa, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x20, 0x3e, 0x20, 0x68, 0x31, 0x20, 0x3e, 0x20, 0x61, 0x3a, 0x6c, 0x69, 0x6e, 0x6b, 0x2c, 0x20, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x20, 0x3e, 0x20, 0x68, 0x31, 0x20, 0x3e, 0x20, 0x61, 0x3a, 0x76, 0x69, 0x73, 0x69, 0x74, 0x65, 0x64, 0x20, 0x7b, 0xa, 0x9, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x3a, 0x20, 0x23, 0x61, 0x66, 0x61, 0x3b, 0xa, 0x9, 0x74, 0x65, 0x78, 0x74, 0x2d, 0x64, 0x65, 0x63, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x6e, 0x6f, 0x6e, 0x65, 0x3b, 0xa, 0x7d, 0xa, 0xa, 0x2e, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x20, 0x7b, 0xa, 0x9, 0x6d, 0x61, 0x78, 0x2d, 0x77, 0x69, 0x64, 0x74, 0x68, 0x3a, 0x20, 0x31, 0x30, 0x30, 0x30, 0x70, 0x78, 0x3b, 0xa, 0x9, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x3a, 0x20, 0x31, 0x65, 0x6d, 0x20, 0x61, 0x75, 0x74, 0x6f, 0x3b, 0xa, 0x9, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x3a, 0x20, 0x72, 0x67, 0x62, 0x61, 0x28, 0x30, 0x2c, 0x20, 0x30, 0x2c, 0x20, 0x30, 0x2c, 0x20, 0x2e, 0x31, 0x29, 0x3b, 0xa, 0x9, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x3a, 0x20, 0x31, 0x70, 0x78, 0x20, 0x73, 0x6f, 0x6c, 0x69, 0x64, 0x20, 0x72, 0x67, 0x62, 0x61, 0x28, 0x30, 0x2c, 0x20, 0x30, 0x2c, 0x20, 0x30, 0x2c, 0x20, 0x2e, 0x35, 0x29, 0x3b, 0xa, 0x9, 0x70, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x20, 0x31, 0x65, 0x6d, 0x20, 0x32, 0x65, 0x6d, 0x3b, 0xa, 0x9, 0x62, 0x6f, 0x78, 0x2d, 0x73, 0x69, 0x7a, 0x69, 0x6e, 0x67, 0x3a, 0x20, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2d, 0x62, 0x6f, 0x78, 0x3b, 0xa, 0x7d, 0xa, 0xa, 0x23, 0x66, 0x69, 0x6e, 0x64, 0x2d, 0x66, 0x65, 0x65, 0x64, 0x20, 0x7b, 0xa, 0x9, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x3a, 0x20, 0x31, 0x65, 0x6d, 0x20, 0x61, 0x75, 0x74, 0x6f, 0x3b, 0xa, 0x9, 0x74, 0x65, 0x78, 0x74, 0x2d, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x3a, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x3b, 0xa, 0x9, 0x66, 0x6c, 0x65, 0x78, 0x2d, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x20, 0x72, 0x6f, 0x77, 0x3b, 0xa, 0x7d, 0xa, 0xa, 0x23, 0x66, 0x69, 0x6e, 0x64, 0x2d, 0x66, 0x65, 0x65, 0x64, 0x20, 0x3e, 0x20, 0x64, 0x69, 0x76, 0x20, 0x7b, 0xa, 0x9, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x3a, 0x20, 0x66, 0x6c, 0x65, 0x78, 0x3b, 0xa, 0x7d, 0xa, 0xa, 0x23, 0x66, 0x69, 0x6e, 0x64, 0x2d, 0x66, 0x65, 0x65, 0x64, 0x20, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x20, 0x7b, 0xa, 0x9, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x3a, 0x20, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x3b, 0xa, 0x9, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x2d, 0x62, 0x6f, 0x74, 0x74, 0x6f, 0x6d, 0x3a, 0x20, 0x31, 0x65, 0x6d, 0x3b, 0xa, 0x7d, 0xa, 0xa, 0x23, 0x66, 0x69, 0x6e, 0x64, 0x2d, 0x66, 0x65, 0x65, 0x64, 0x20, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x20, 0x7b, 0xa, 0x9, 0x66, 0x6f, 0x6e, 0x74, 0x2d, 0x73, 0x69, 0x7a, 0x65, 0x3a, 0x20, 0x31, 0x65, 0x6d, 0x3b, 0xa, 0x9, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x3a, 0x20, 0x23, 0x66, 0x66, 0x66, 0x3b, 0xa, 0x9, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x3a, 0x20, 0x31, 0x70, 0x78, 0x20, 0x73, 0x6f, 0x6c, 0x69, 0x64, 0x20, 0x72, 0x67, 0x62, 0x61, 0x28, 0x30, 0x2c, 0x20, 0x30, 0x2c, 0x20, 0x30, 0x2c, 0x20, 0x2e, 0x32, 0x29, 0x3b, 0xa, 0x7d, 0xa, 0xa, 0x23, 0x66, 0x69, 0x6e, 0x64, 0x2d, 0x66, 0x65, 0x65, 0x64, 0x20, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5b, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x74, 0x65, 0x78, 0x74, 0x5d, 0x20, 0x7b, 0xa, 0x9, 0x66, 0x6c, 0x65, 0x78, 0x3a, 0x20, 0x31, 0x3b, 0xa, 0x9, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2d, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x3a, 0x20, 0x36, 0x70, 0x78, 0x20, 0x30, 0x20, 0x30, 0x20, 0x36, 0x70, 0x78, 0x3b, 0xa, 0x9, 0x74, 0x65, 0x78, 0x74, 0x2d, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x3a, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x3b, 0xa, 0x7d, 0xa, 0xa, 0x23, 0x66, 0x69, 0x6e, 0x64, 0x2d, 0x66, 0x65, 0x65, 0x64, 0x20, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5b, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x5d, 0x20, 0x7b, 0xa, 0x9, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2d, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x3a, 0x20, 0x30, 0x20, 0x36, 0x70, 0x78, 0x20, 0x36, 0x70, 0x78, 0x20, 0x30, 0x3b, 0xa, 0x9, 0x70, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x20, 0x2e, 0x35, 0x65, 0x6d, 0x20, 0x31, 0x2e, 0x32, 0x65, 0x6d, 0x3b, 0xa, 0x9, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x3a, 0x20, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x3b, 0xa, 0x9, 0x62, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2d, 0x6c, 0x65, 0x66, 0x74, 0x3a, 0x20, 0x6e, 0x6f, 0x6e, 0x65, 0x3b, 0xa, 0x7d, 0xa, 0xa, 0x2e, 0x62, 0x6c, 0x75, 0x72, 0x62, 0x20, 0x7b, 0xa, 0x9, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x3a, 0x20, 0x61, 0x75, 0x74, 0x6f, 0x3b, 0xa, 0x7d, 0xa, 0xa, 0x61, 0x5b, 0x72, 0x65, 0x6c, 0x3d, 0x6e, 0x65, 0x78, 0x74, 0x5d, 0x20, 0x7b, 0xa, 0x9, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x3a, 0x20, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x3b, 0xa, 0x9, 0x6d, 0x61, 0x72, 0x67, 0x69, 0x6e, 0x3a, 0x20, 0x31, 0x65, 0x6d, 0x20, 0x61, 0x75, 0x74, 0x6f, 0x3b, 0xa, 0x9, 0x74, 0x65, 0x78, 0x74, 0x2d, 0x61, 0x6c, 0x69, 0x67, 0x6e, 0x3a, 0x20, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x3b, 0xa, 0x9, 0x70, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x3a, 0x20, 0x31, 0x65, 0x6d, 0x3b, 0xa, 0x9, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x3a, 0x20, 0x23, 0x32, 0x32, 0x32, 0x3b, 0xa, 0x9, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x3a, 0x20, 0x23, 0x66, 0x66, 0x66, 0x3b, 0xa, 0x9, 0x74, 0x65, 0x78, 0x74, 0x2d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x20, 0x75, 0x70, 0x70, 0x65, 0x72, 0x63, 0x61, 0x73, 0x65, 0x3b, 0xa, 0x7d, 0xa}), //++ TODO: optimize? (double allocation) or does compiler already optimize this?
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   ``,
		DirModTime: time.Unix(1491484012, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // style.css

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`public`, &embedded.EmbeddedBox{
		Name: `public`,
		Time: time.Unix(1491620838, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"style.css": file2,
		},
	})
}
