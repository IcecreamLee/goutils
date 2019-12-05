package goutils

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// GetCurrentPath 返回当前程序运行的路径
func GetCurrentPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1) + "/"
}
