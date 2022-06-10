package gkong

import (
	"context"

	"github.com/lffranca/bento/pkg/domain/entity"
	"github.com/lffranca/bento/pkg/domain/provider"
)

func NewServiceAdapter() provider.DataResearcherProvider[entity.Service] {
	return &serviceAdapter{}
}

type serviceAdapter struct{}

func (pkg *serviceAdapter) List(ctx context.Context, gateway *entity.Gateway, pagination entity.Pagination, tag entity.TagOption) (items []entity.Service, err error) {
	return
}

func (pkg *serviceAdapter) Get(ctx context.Context, gateway *entity.Gateway, nameOrID string) (item *entity.Service, err error) {
	return
}
