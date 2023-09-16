package router

import (
	"context"
)

type Request struct {
	Token  string `query:"token" vd:"ValidateToken($)"`
	Config string `query:"config"`
}

type Response struct {
	Data interface{}
}

type Handler func(ctx context.Context, req Request) (resp Response, err error)

func echo(ctx context.Context, req Request) (resp Response, err error) {
	resp = Response{
		Data: req,
	}
	return
}
