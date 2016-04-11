package util

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"path/filepath"

	"github.com/bborbe/log"
)

var logger = log.DefaultLogger

func IsDirectory(dir string) error {

	file, err := os.Open(dir)
	if err != nil {
		logger.Debugf("open %s failed: %v", dir, err)
		return err
	}
	defer file.Close()
	fileinfo, err := file.Stat()
	if err != nil {
		logger.Debugf("file stat failed: %v", err)
		return err
	}
	if !fileinfo.IsDir() {
		msg := fmt.Sprintf("%s is not a directory", dir)
		logger.Debug(msg)
		return errors.New(msg)
	}
	return nil
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
