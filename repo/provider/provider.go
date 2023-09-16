package provider

import "context"

type Provider struct{}

func (_ *Provider) Get(ctx context.Context, file string) (content []byte, err error) {

	return
}

func (_ *Provider) GetAll(ctx context.Context, dir string) (contents map[string][]byte, err error) {

	return
}

func (_ *Provider) Set(ctx context.Context, file string, content []byte) (err error) {

	return
}
