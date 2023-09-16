package persist

import "context"

type Persist struct{}

func (_ *Persist) Get(ctx context.Context, file string) (content []byte, err error) {

	return
}

func (_ *Persist) GetAll(ctx context.Context, dir string) (contents map[string][]byte, err error) {

	return
}

func (_ *Persist) Set(ctx context.Context, file string, content []byte) (err error) {

	return
}
