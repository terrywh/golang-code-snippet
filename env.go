package snippet

import (
	"os"
	"path/filepath"
	"strings"
)

var AppRoot string

func init() {
	root, _ := filepath.Abs(os.Args[0])
	root = filepath.Dir(filepath.Dir(root))
	if strings.Contains(root, "go-build") {
		root, _ = os.Getwd()
	}
	AppRoot = filepath.Clean(filepath.FromSlash(root))
}
