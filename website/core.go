package website

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var SKey = "JPcT1k6SwyqA2JX-oyGjGfOfHzKsN2BQdI4Cr56KG9M="
var StaticUrl = "/static"

type Paths struct {
	DirPath   string
	StaticUrl string
}

func CleanPath(path string) string {
	path = strings.Replace(path, "//", "/", -1)
	return path
}

func AbsolutePath(path string) string {
	if flag.Lookup("test.v") == nil {
		return CleanPath(filepath.FromSlash(BinPath + "/" + path))
	} else {
		return CleanPath(filepath.FromSlash(path))
	}
}

func BinaryPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal("Error getting file path")
	}
	return dir
}
func FixPathSlash(path string) string {
	return filepath.FromSlash(path)
}
func PagePath(dirPath string, spage string) string {
	return FixPathSlash(dirPath + "/html/" + spage)
}

var BinPath = BinaryPath()
var StaticDir = AbsolutePath("static")
