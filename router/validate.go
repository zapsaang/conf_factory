package router

import (
	"github.com/cloudwego/hertz/pkg/app/server/binding"
	"github.com/zapsaang/conf_factory/conf/env"
	"github.com/zapsaang/conf_factory/utils/validates"
)

func MustRegValidate() {
	binding.MustRegValidateFunc("ValidateToken", validates.Equals(env.GetToken()))
}
