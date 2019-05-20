package webui

import (
	"net/http"

	"github.com/rakyll/statik/fs"
)

var FileSystem http.FileSystem

func init() {
	data := "\x3c\x21\x64\x6f\x63\x74\x79\x70\x65\x20\x68\x74\x6d\x6c\x3e\x0a\x0a\x3c\x68\x74\x6d\x6c\x20\x6c\x61\x6e\x67\x3d\x22\x65\x6e\x22\x3e\x0a\x3c\x68\x65\x61\x64\x3e\x0a\x20\x20\x3c\x6d\x65\x74\x61\x20\x63\x68\x61\x72\x73\x65\x74\x3d\x22\x75\x74\x66\x2d\x38\x22\x3e\x0a\x0a\x20\x20\x3c\x74\x69\x74\x6c\x65\x3e\x41\x67\x67\x72\x65\x63\x75\x74\x74\x6f\x72\x3c\x2f\x74\x69\x74\x6c\x65\x3e\x0a\x20\x20\x3c\x6d\x65\x74\x61\x20\x6e\x61\x6d\x65\x3d\x22\x64\x65\x73\x63\x72\x69\x70\x74\x69\x6f\x6e\x22\x20\x63\x6f\x6e\x74\x65\x6e\x74\x3d\x22\x54\x68\x65\x20\x48\x54\x4d\x4c\x35\x20\x48\x65\x72\x61\x6c\x64\x22\x3e\x0a\x20\x20\x3c\x6d\x65\x74\x61\x20\x6e\x61\x6d\x65\x3d\x22\x61\x75\x74\x68\x6f\x72\x22\x20\x63\x6f\x6e\x74\x65\x6e\x74\x3d\x22\x53\x69\x74\x65\x50\x6f\x69\x6e\x74\x22\x3e\x0a\x0a\x20\x20\x3c\x6c\x69\x6e\x6b\x20\x72\x65\x6c\x3d\x22\x73\x74\x79\x6c\x65\x73\x68\x65\x65\x74\x22\x20\x68\x72\x65\x66\x3d\x22\x63\x73\x73\x2f\x73\x74\x79\x6c\x65\x73\x2e\x63\x73\x73\x3f\x76\x3d\x31\x2e\x30\x22\x3e\x0a\x0a\x3c\x2f\x68\x65\x61\x64\x3e\x0a\x0a\x3c\x62\x6f\x64\x79\x3e\x0a\x20\x20\x3c\x73\x63\x72\x69\x70\x74\x20\x73\x72\x63\x3d\x22\x6a\x73\x2f\x73\x63\x72\x69\x70\x74\x73\x2e\x6a\x73\x22\x3e\x3c\x2f\x73\x63\x72\x69\x70\x74\x3e\x0a\x3c\x2f\x62\x6f\x64\x79\x3e\x0a\x3c\x2f\x68\x74\x6d\x6c\x3e"
	fs.Register(data)
}
