package util

import (
	"fmt"
	"os"
	"strings"
	"path/filepath"
	"github.com/bborbe/log"
)

var logger = log.DefaultLogger

func IsDirectory(dir string) (bool, error) {
	logger.Debugf("IsDir %s", dir)
	file, err := os.Open(dir)
	defer file.Close()
	if err != nil {
		logger.Debugf("IsDir - open dir %s failed: %v", dir, err)
		return false, nil
	}
	fileinfo, err := file.Stat()
	if err != nil {
		logger.Debugf("IsDir get state for dir %s failed: %v", dir, err)
		return false, err
	}
	return fileinfo.IsDir(), nil
}

func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func NormalizePath(path string) (string, error) {
	if strings.Index(path, "~/") == 0 {
		home := os.Getenv("HOME")
		if len(home) == 0 {
			return "", fmt.Errorf("env HOME not found")
		}
		path = fmt.Sprintf("%s/%s", home, path[2:])
		logger.Debugf("replace ~/ with homedir. new path: %s", path)
	}
	return filepath.Abs(path)
}
