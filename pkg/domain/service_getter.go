package domain

import (
	"context"

	"github.com/lffranca/bento/pkg/domain/entity"
)

type ServiceGetterService interface {
	List(ctx context.Context, gatewayID string, pagination entity.Pagination, tag entity.TagOption) (items []entity.Service, err error)
	Get(ctx context.Context, gatewayID string, nameOrID string) (item *entity.Service, err error)
}
