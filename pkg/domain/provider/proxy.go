package provider

import (
	"context"

	"github.com/lffranca/bento/pkg/domain/entity"
)

type ProxyProvider interface {
	Get(ctx context.Context, nameOrID string) (*entity.Proxy, error)
	List(ctx context.Context, pagination entity.Pagination, tag entity.TagOption) ([]entity.Proxy, error)
}
