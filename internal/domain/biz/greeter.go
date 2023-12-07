package biz

import (
	"context"
	"github.com/go-kratos/kratos-layout/internal/domain/entity"
	"github.com/go-kratos/kratos-layout/internal/domain/repository"

	v1 "github.com/go-kratos/kratos-layout/api/helloworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// GreeterUsecase is a Greeter usecase.
type GreeterUsecase struct {
	repo repository.GreeterRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewGreeterUsecase(repo repository.GreeterRepo, logger log.Logger) *GreeterUsecase {
	return &GreeterUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *GreeterUsecase) CreateGreeter(ctx context.Context, g *entity.GreeterDo) (*entity.GreeterDo, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)
	greeterPo, err := uc.repo.Save(ctx, &repository.GreeterPo{
		Hello: g.Hello,
	})
	if err != nil {
		return nil, err
	}
	return &entity.GreeterDo{
		Hello: greeterPo.Hello,
	}, nil
}
