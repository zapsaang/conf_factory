package files

import (
	"bytes"
	"os"
	"path/filepath"

	"github.com/zapsaang/conf_factory/pkg/logs"
	"github.com/zapsaang/conf_factory/utils/consts"
)

func Read(root string, readFn ReadFunc, readOpts ...ReadOption) {
	root = GetAbsolutePath(root)

	// TODO change to WalkDir
	err := filepath.Walk(warpRead(root, readFn, readOpts...))

	if err != nil {
		logs.WithError(err).WithField("read_dir", root).Error("walk dir failed")
	}
}

func List(root string) []string {
	// TODO support listOpts
	root = GetAbsolutePath(root)

	dirs, err := os.ReadDir(root)
	if err != nil {
		logs.WithError(err).WithField("read_dir", root).Error("read dir failed")
	}
	var files = make([]string, 0, len(dirs))
	for _, entry := range dirs {
		if entry.IsDir() {
			continue
		}
		files = append(files, entry.Name())
	}
	return files
}

func GetAbsolutePath(root string) string {
	// TODO support windows path
	pathLen := len(root)
	if pathLen > 0 && root[0] == consts.ConstUnixAbsolutePathPrefix {
		return root
	}

	wd, err := os.Getwd()
	if err != nil {
		return root
	}

	pathLen += len(wd) + 1
	pathBuf := bytes.Buffer{}
	pathBuf.Grow(pathLen)
	pathBuf.WriteString(wd)
	pathBuf.WriteRune(consts.ConstUnixPathSeparator)
	pathBuf.WriteString(root)
	return pathBuf.String()
}
