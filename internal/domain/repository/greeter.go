package repository

import "context"

type GreeterPo struct {
	Hello string
}

// GreeterRepo is a Greater repo.
type GreeterRepo interface {
	Save(context.Context, *GreeterPo) (*GreeterPo, error)
	Update(context.Context, *GreeterPo) (*GreeterPo, error)
	FindByID(context.Context, int64) (*GreeterPo, error)
	ListByHello(context.Context, string) ([]*GreeterPo, error)
	ListAll(context.Context) ([]*GreeterPo, error)
}
