package surge

import (
	"context"

	"github.com/zapsaang/conf_factory/conf/conf"
	"github.com/zapsaang/conf_factory/domain/surge"
	"github.com/zapsaang/conf_factory/pkg/logs"
	"github.com/zapsaang/conf_factory/utils/files"
)

func Integrator() {
	ctx := context.Background()
	baseNames := files.List(conf.GetSurgeBaseDir())
	for _, baseName := range baseNames {
		workSpace := surge.New(baseName, surge.Config{})
		for _, processNode := range surge.ProcessFlow {
			if err := processNode(ctx, workSpace); err != nil {
				logs.WithError(err).Error("process failed")
				continue
			}
		}
	}
}
