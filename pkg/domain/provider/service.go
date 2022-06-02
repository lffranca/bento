package provider

import (
	"context"

	"github.com/lffranca/bento/pkg/domain/entity"
)

type ServiceProvider interface {
	List(ctx context.Context, gateway *entity.Gateway, pagination entity.Pagination, tag entity.TagOption) (items []entity.Service, err error)
	Get(ctx context.Context, gateway *entity.Gateway, nameOrID string) (item *entity.Service, err error)
}
