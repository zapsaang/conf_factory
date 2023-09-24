package files

import (
	"io/fs"
	"os"
)

type WalkDirOption func(root, path string, d fs.DirEntry) error

type ReadFunc func(path string, content []byte)

func warpRead(root string, readFn ReadFunc, readOpts ...WalkDirOption) (string, fs.WalkDirFunc) {
	return root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		for _, opt := range readOpts {
			if optErr := opt(root, path, d); err != nil {
				return optErr
			}
		}

		if d.IsDir() {
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

func SkipDir(_, _ string, d fs.DirEntry) error {
	if d.IsDir() {
		return fs.SkipDir
	}
	return nil
}

func SkipSubDir(root, path string, d fs.DirEntry) error {
	if d.IsDir() && path != root {
		return fs.SkipDir
	}
	return nil
}

func SkipAll(_, _ string, _ fs.DirEntry) error {
	return fs.SkipAll
}
