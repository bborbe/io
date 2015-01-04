package util

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"strings"

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
		usr, err := user.Current()
		if err != nil {
			return "", err
		}
		path = fmt.Sprintf("%s/%s", usr.HomeDir, path[2:])
		logger.Debugf("replace ~/ with homedir. new path: %s", path)
	}
	return path, nil
}
