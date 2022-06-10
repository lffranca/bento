package service

import (
	"context"

	"github.com/lffranca/bento/pkg/domain"
	"github.com/lffranca/bento/pkg/domain/entity"
	"github.com/lffranca/bento/pkg/domain/provider"
)

func NewDataResearcherService[T any](
	gatewayProvider provider.GatewayProvider,
	provider provider.DataResearcherProvider[T],
) domain.DataResearcherService[T] {
	return &dataResearcherService[T]{
		gatewayProvider,
		provider,
	}
}

type dataResearcherService[T any] struct {
	gatewayProvider provider.GatewayProvider
	provider        provider.DataResearcherProvider[T]
}

func (pkg *dataResearcherService[T]) List(ctx context.Context, gatewayID string, pagination entity.Pagination, tag entity.TagOption) (items []T, err error) {
	var gateway *entity.Gateway

	gateway, err = pkg.gatewayProvider.Get(ctx, gatewayID)
	if err != nil {
		return nil, err
	}

	return pkg.provider.List(ctx, gateway, pagination, tag)
}

func (pkg *dataResearcherService[T]) Get(ctx context.Context, gatewayID string, nameOrID string) (item *T, err error) {
	var gateway *entity.Gateway

	gateway, err = pkg.gatewayProvider.Get(ctx, gatewayID)
	if err != nil {
		return nil, err
	}

	return pkg.provider.Get(ctx, gateway, nameOrID)
}
