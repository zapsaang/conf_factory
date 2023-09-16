package files

import (
	"os"
	"path/filepath"
)

type ReadOption func(root, path string, info os.FileInfo) error

type ReadFunc func(path string, content []byte)

func warpRead(root string, readFn ReadFunc, readOpts ...ReadOption) (string, filepath.WalkFunc) {
	return root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		for _, opt := range readOpts {
			if optErr := opt(root, path, info); err != nil {
				return optErr
			}
		}

		if info.IsDir() {
			return nil
		}

		bytes, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		readFn(path, bytes)
		return nil
	}
}

func SkipDir(_, _ string, info os.FileInfo) error {
	if info.IsDir() {
		return filepath.SkipDir
	}
	return nil
}

func SkipSubDir(root, path string, info os.FileInfo) error {
	if info.IsDir() && path != root {
		return filepath.SkipDir
	}
	return nil
}

func SkipAll(_, _ string, _ os.FileInfo) error {
	return filepath.SkipAll
}
