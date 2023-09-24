package pool

import (
	"net/http"
	"sync"

	"github.com/panjf2000/ants/v2"
	"github.com/zapsaang/conf_factory/utils/consts"
)

// HTTPClientPool 用于管理 HTTP 客户端连接的连接池。
var httpClient = sync.Pool{
	New: func() any {
		return &http.Client{
			Timeout: consts.HTTPRequestTimeout,
		}
	},
}

func GetHTTPClient() *http.Client {
	return httpClient.Get().(*http.Client)
}

func PutHTTPClient(cli *http.Client) {
	httpClient.Put(cli)
}

func MustNewPoolWithFunc(fn func(i interface{})) *ants.PoolWithFunc {
	pool, err := ants.NewPoolWithFunc(consts.AntsPoolLimitation, fn)
	if err != nil {
		panic(err)
	}
	return pool
}
