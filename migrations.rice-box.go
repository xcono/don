package main

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file8 := &embedded.EmbeddedFile{
		Filename:    `000_init.sql`,
		FileModTime: time.Unix(1491616428, 0),
		Content:     string([]byte{0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x20, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x20, 0x28, 0xa, 0x9, 0x69, 0x64, 0x20, 0x74, 0x65, 0x78, 0x74, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x6e, 0x75, 0x6c, 0x6c, 0x20, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x20, 0x6b, 0x65, 0x79, 0x2c, 0xa, 0x9, 0x68, 0x75, 0x62, 0x20, 0x74, 0x65, 0x78, 0x74, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x6e, 0x75, 0x6c, 0x6c, 0x2c, 0xa, 0x9, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x20, 0x74, 0x65, 0x78, 0x74, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x6e, 0x75, 0x6c, 0x6c, 0x2c, 0xa, 0x9, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x75, 0x72, 0x6c, 0x20, 0x74, 0x65, 0x78, 0x74, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x6e, 0x75, 0x6c, 0x6c, 0x2c, 0xa, 0x9, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x20, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x2c, 0xa, 0x9, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x20, 0x28, 0x68, 0x75, 0x62, 0x2c, 0x20, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x29, 0xa, 0x29, 0x3b, 0xa, 0xa, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x20, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x20, 0x28, 0xa, 0x9, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x20, 0x74, 0x65, 0x78, 0x74, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x6e, 0x75, 0x6c, 0x6c, 0x20, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x20, 0x6b, 0x65, 0x79, 0x2c, 0xa, 0x9, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x65, 0x6e, 0x20, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x6e, 0x75, 0x6c, 0x6c, 0x2c, 0xa, 0x9, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x74, 0x65, 0x78, 0x74, 0x2c, 0xa, 0x9, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x74, 0x65, 0x78, 0x74, 0x2c, 0xa, 0x9, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x20, 0x74, 0x65, 0x78, 0x74, 0x2c, 0xa, 0x9, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x20, 0x74, 0x65, 0x78, 0x74, 0x2c, 0xa, 0x9, 0x6e, 0x6f, 0x74, 0x65, 0x20, 0x74, 0x65, 0x78, 0x74, 0xa, 0x29, 0x3b, 0xa, 0xa, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x20, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x20, 0x28, 0xa, 0x9, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x20, 0x74, 0x65, 0x78, 0x74, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x6e, 0x75, 0x6c, 0x6c, 0x2c, 0xa, 0x9, 0x69, 0x64, 0x20, 0x74, 0x65, 0x78, 0x74, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x6e, 0x75, 0x6c, 0x6c, 0x2c, 0xa, 0x9, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x20, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x6e, 0x75, 0x6c, 0x6c, 0x2c, 0xa, 0x9, 0x72, 0x61, 0x77, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x20, 0x74, 0x65, 0x78, 0x74, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x6e, 0x75, 0x6c, 0x6c, 0x2c, 0xa, 0x9, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x20, 0x6b, 0x65, 0x79, 0x20, 0x28, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x2c, 0x20, 0x69, 0x64, 0x29, 0xa, 0x29, 0x3b, 0xa}), //++ TODO: optimize? (double allocation) or does compiler already optimize this?
	}

	// define dirs
	dir7 := &embedded.EmbeddedDir{
		Filename:   ``,
		DirModTime: time.Unix(1491616134, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file8, // 000_init.sql

		},
	}

	// link ChildDirs
	dir7.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`migrations`, &embedded.EmbeddedBox{
		Name: `migrations`,
		Time: time.Unix(1491620838, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir7,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"000_init.sql": file8,
		},
	})
}
