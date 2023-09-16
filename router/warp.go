package router

import (
	"context"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/spf13/cast"
	"github.com/zapsaang/conf_factory/pkg/logs"
	"github.com/zapsaang/conf_factory/utils"
	"github.com/zapsaang/conf_factory/utils/consts"
	"github.com/zapsaang/conf_factory/utils/errors"
)

type WarpOption func(ctx context.Context, c *app.RequestContext, req *Request) error

func Warp(handler Handler, warpOpts ...WarpOption) func(ctx context.Context, c *app.RequestContext) {
	return func(ctx context.Context, c *app.RequestContext) {
		logger := logs.WithContext(ctx)
		logger.Info("get new request")

		var req Request
		var err error

		for _, w := range warpOpts {
			if err = w(ctx, c, &req); err != nil {
				c.String(errors.ErrorToResp(err))
				return
			}
		}

		resp, err := handler(ctx, req)
		if err != nil {
			logger.WithError(err).Error("handle request failed")
			c.String(errors.ErrorToResp(err))
			return
		}
		body, err := cast.ToStringE(resp.Data)
		if err != nil {
			body = utils.MustToIndentedJSONString(resp.Data)
		}
		c.String(consts.StatusOK, body)
		logger.Info("handle request finished")
	}
}

func BindReq(ctx context.Context, c *app.RequestContext, req *Request) error {
	logger := logs.WithContext(ctx)
	if err := c.Bind(&req); err != nil {
		logger.WithError(err).Error("bind request failed")
		return err
	}
	logger.Info("bind request finished")
	return nil
}

func ValidateReq(ctx context.Context, c *app.RequestContext, req *Request) error {
	logger := logs.WithContext(ctx)
	if err := c.Validate(&req); err != nil {
		logger.WithError(err).Error("validate request failed")
		return err
	}
	logger.Info("validate request finished")
	return nil
}

func LogReq(ctx context.Context, c *app.RequestContext, req *Request) error {
	logger := logs.WithContext(ctx)
	reqJSON, err := sonic.MarshalString(req)
	if err != nil {
		logger.WithError(err).Error("marshal request to string failed")
	}
	logger.WithField("request", reqJSON).Info("log request")
	return nil
}
