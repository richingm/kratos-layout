package infrastructure

import (
	"context"
	"github.com/go-kratos/kratos-layout/internal/domain/repository"

	"github.com/go-kratos/kratos/v2/log"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) repository.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) Save(ctx context.Context, g *repository.GreeterPo) (*repository.GreeterPo, error) {
	return g, nil
}

func (r *greeterRepo) Update(ctx context.Context, g *repository.GreeterPo) (*repository.GreeterPo, error) {
	return g, nil
}

func (r *greeterRepo) FindByID(context.Context, int64) (*repository.GreeterPo, error) {
	return nil, nil
}

func (r *greeterRepo) ListByHello(context.Context, string) ([]*repository.GreeterPo, error) {
	return nil, nil
}

func (r *greeterRepo) ListAll(context.Context) ([]*repository.GreeterPo, error) {
	return nil, nil
}
