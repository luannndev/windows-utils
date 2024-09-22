package main

import (
	"os"
	"path/filepath"
)

func deleteFilesInDirectory(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path == dir {
			return nil
		}

		if info.IsDir() {
			return os.RemoveAll(path)
		} else {
			return os.Remove(path)
		}
	})
}
