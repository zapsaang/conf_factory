package cache

import "context"

type Cache struct{}

func (_ *Cache) Get(ctx context.Context, file string) (content []byte, err error) {

	return
}

func (_ *Cache) GetAll(ctx context.Context, dir string) (contents map[string][]byte, err error) {

	return
}

func (_ *Cache) Set(ctx context.Context, file string, content []byte) (err error) {

	return
}
