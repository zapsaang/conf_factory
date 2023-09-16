package remote

import (
	"context"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/zapsaang/conf_factory/pkg/logs"
)

type Remote struct {
	*http.Client
}

var remoteClient *Remote
var remoteOnce sync.Once

func GetClient() *Remote {
	remoteOnce.Do(
		func() {
			remoteClient = &Remote{
				&http.Client{},
			}
		})
	return remoteClient
}

func (r *Remote) Get(ctx context.Context, file string) (content []byte, err error) {
	logger := logs.WithField("get_file", file)
	req, err := http.NewRequest("GET", file, nil)
	if err != nil {
		logger.WithError(err).Error("new request failed")
		return []byte{}, err
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")
	resp, err := r.Do(req)
	if err != nil {
		logger.WithError(err).Error("get file failed")
		return []byte{}, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.WithError(err).Error("read body failed")
		return []byte{}, err
	}
	return bodyBytes, nil
}

func (r *Remote) GetAll(ctx context.Context, dir string) (contents map[string][]byte, err error) {

	return
}

func (r *Remote) Set(ctx context.Context, file string, content []byte) (err error) {

	return
}
