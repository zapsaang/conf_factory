package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

func Register(h *server.Hertz) {
	MustRegValidate()
	h.GET("echo", Warp(echo, BindReq, LogReq))
	h.GET("validate", Warp(echo, BindReq, ValidateReq, LogReq))
}
