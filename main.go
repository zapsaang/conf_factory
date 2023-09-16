package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"

	"github.com/zapsaang/conf_factory/conf/env"
	"github.com/zapsaang/conf_factory/cronjob"
	"github.com/zapsaang/conf_factory/router"
)

func main() {
	cronjob.Init()

	h := server.Default(server.WithHostPorts(env.GetServerAddress()))
	router.Register(h)
	h.Spin()
}
