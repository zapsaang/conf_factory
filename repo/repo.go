package repo

import (
	"context"

	"github.com/zapsaang/conf_factory/repo/integrator"
	"github.com/zapsaang/conf_factory/repo/provider"
)

type Repo interface {
	Get(ctx context.Context, file string) (content []byte, err error)
	GetAll(ctx context.Context, dir string) (contents map[string][]byte, err error)
	Set(ctx context.Context, file string, content []byte) (err error)
}

type Role uint8

const (
	RoleIntegrator Role = 1
	RoleProvider   Role = 2
)

func New(role Role) Repo {
	switch role {
	case RoleIntegrator:
		return &integrator.Integrator{}
	case RoleProvider:
		return &provider.Provider{}
	default:
		return nil
	}
}
