package provider

import (
	"context"

	"github.com/lffranca/bento/pkg/domain/entity"
)

type DataResearcherProvider[T any] interface {
	List(ctx context.Context, gateway *entity.Gateway, pagination entity.Pagination, tag entity.TagOption) (items []T, err error)
	Get(ctx context.Context, gateway *entity.Gateway, nameOrID string) (item *T, err error)
}
