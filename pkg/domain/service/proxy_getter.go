package service

import (
	"context"

	"github.com/lffranca/bento/pkg/domain"
	"github.com/lffranca/bento/pkg/domain/entity"
	"github.com/lffranca/bento/pkg/domain/provider"
)

func NewProxyGetterService(proxyProvider provider.ProxyProvider) domain.ProxyGetterService {
	return &ProxyGetterService{proxyProvider}
}

type ProxyGetterService struct {
	proxyProvider provider.ProxyProvider
}

func (pkg *ProxyGetterService) List(ctx context.Context, pagination entity.Pagination, tag entity.TagOption) ([]entity.Proxy, error) {
	return pkg.proxyProvider.List(ctx, pagination, tag)
}

func (pkg *ProxyGetterService) Get(ctx context.Context, nameOrID string) (*entity.Proxy, error) {
	return pkg.proxyProvider.Get(ctx, nameOrID)
}
