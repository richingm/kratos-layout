package domain

import (
	"github.com/go-kratos/kratos-layout/internal/domain/biz"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(biz.NewGreeterUsecase)
