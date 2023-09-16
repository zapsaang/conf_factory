package integrator

import (
	"context"
	"os"
	"strings"

	"github.com/zapsaang/conf_factory/pkg/logs"
	"github.com/zapsaang/conf_factory/repo/remote"
)

type Integrator struct{}

func (i *Integrator) Get(ctx context.Context, file string) (content []byte, err error) {
	logger := logs.WithField("get_file", file)
	if strings.HasPrefix(file, "http") {
		logger.Info("remote file")
		return remote.GetClient().Get(ctx, file)
	}
	content, err = os.ReadFile(file)
	if err != nil {
		logger.WithError(err).Error("read file failed")
	}
	return
}

func (i *Integrator) GetAll(ctx context.Context, dir string) (contents map[string][]byte, err error) {

	return
}

func (i *Integrator) Set(ctx context.Context, file string, content []byte) (err error) {
	logger := logs.WithField("set_file", file)
	err = os.WriteFile(file, content, 0644)
	if err != nil {
		logger.WithError(err).Error("write file failed")
		return
	}
	return
}
