package files

import (
	"fmt"
	"io"
	"net/http"

	"github.com/zapsaang/conf_factory/utils/consts"
	"github.com/zapsaang/conf_factory/utils/pool"
)

func Download(url string) ([]byte, error) {
	client := pool.GetHTTPClient()
	defer pool.PutHTTPClient(client)

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func isURL(str string) bool {
	return consts.RegexpCompiledURL.MatchString(str)
}
