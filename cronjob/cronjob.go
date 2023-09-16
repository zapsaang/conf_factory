package cronjob

import (
	"github.com/robfig/cron/v3"

	"github.com/zapsaang/conf_factory/cronjob/surge"
)

func Init() {
	c := cron.New()

	c.AddFunc("@every 1h", surge.Integrator)

	c.Start()
}
