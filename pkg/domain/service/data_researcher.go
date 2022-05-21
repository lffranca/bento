package service

import (
	"context"

	"github.com/lffranca/bento/pkg/domain"
	"github.com/lffranca/bento/pkg/domain/entity"
	"github.com/lffranca/bento/pkg/domain/provider"
)

func NewDataResearcherService(
	gatewayProvider provider.GatewayProvider,
	dataProvider provider.DataResearcherProvider,
) domain.DataResearcherService {
	return &DataResearcherService{gatewayProvider, dataProvider}
}

type DataResearcherService[T entity.SearchResponse] struct {
	gatewayProvider provider.GatewayProvider
	dataProvider    provider.DataResearcherProvider
}

func (pkg *DataResearcherService[T]) List(ctx context.Context, gatewayID string, offset, limit int, tags []string, matchAllTags bool) (response *T, err error) {
	var gateway *entity.Gateway

	gateway, err = pkg.gatewayProvider.Get(ctx, gatewayID)
	if err != nil {
		return nil, err
	}

	result, err := pkg.dataProvider.List(ctx, gateway, offset, limit, tags, matchAllTags)
	if err != nil {
		return nil, err
	}

	response.Items = result.Items
	response.Options = result.Options

	return response, nil
}

func (pkg *DataResearcherService[T]) Get(ctx context.Context, gatewayID string, nameOrID string) (items *T, err error) {
	return
}
