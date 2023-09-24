package files

import (
	"os"
	"sync"

	"github.com/zapsaang/conf_factory/utils/pool"
)

type file struct {
	wg      *sync.WaitGroup
	name    string
	content []byte
}

var getFilePool = pool.MustNewPoolWithFunc(func(i interface{}) {
	args := i.(*file)
	defer args.wg.Done()

	args.content = Get(args.name)
})

func Get(urlOrPath string) (buf []byte) {
	if isURL(urlOrPath) {
		buf, _ = Download(urlOrPath)
		return buf
	}
	buf, _ = os.ReadFile(urlOrPath)
	return buf
}

func GetAll(urlOrPaths []string) [][]byte {
	wg := &sync.WaitGroup{}
	queue := make([]file, len(urlOrPaths))
	wg.Add(len(urlOrPaths))
	for i := range urlOrPaths {
		queue[i].wg = wg
		queue[i].name = urlOrPaths[i]
		getFilePool.Invoke(&queue[i])
	}
	wg.Wait()

	result := make([][]byte, len(urlOrPaths))
	for i := range queue {
		result[i] = queue[i].content
	}
	return result
}
