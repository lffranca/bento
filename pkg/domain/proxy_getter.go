package domain

import (
	"context"

	"github.com/lffranca/bento/pkg/domain/entity"
)

type ProxyGetterService interface {
	List(ctx context.Context, pagination entity.Pagination, tag entity.TagOption) ([]entity.Proxy, error)
	Get(ctx context.Context, nameOrID string) (*entity.Proxy, error)
}
