package service

import (
	"context"

	"github.com/lffranca/bento/pkg/domain"
	"github.com/lffranca/bento/pkg/domain/entity"
	"github.com/lffranca/bento/pkg/domain/provider"
)

func NewServiceGetterService(gatewayProvider provider.GatewayProvider, serviceProvider provider.ServiceProvider) domain.ServiceGetterService {
	return &serviceGetterService{gatewayProvider, serviceProvider}
}

type serviceGetterService struct {
	gatewayProvider provider.GatewayProvider
	serviceProvider provider.ServiceProvider
}

func (pkg *serviceGetterService) List(ctx context.Context, gatewayID string, pagination entity.Pagination, tag entity.TagOption) (items []entity.Service, err error) {
	var gateway *entity.Gateway

	gateway, err = pkg.gatewayProvider.Get(ctx, gatewayID)
	if err != nil {
		return nil, err
	}

	return pkg.serviceProvider.List(ctx, gateway, pagination, tag)
}

func (pkg *serviceGetterService) Get(ctx context.Context, gatewayID string, nameOrID string) (item *entity.Service, err error) {
	var gateway *entity.Gateway

	gateway, err = pkg.gatewayProvider.Get(ctx, gatewayID)
	if err != nil {
		return nil, err
	}

	return pkg.serviceProvider.Get(ctx, gateway, nameOrID)
}
