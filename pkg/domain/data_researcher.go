package domain

import (
	"context"

	"github.com/lffranca/bento/pkg/domain/entity"
)

type DataResearcherService[T any] interface {
	List(ctx context.Context, gatewayID string, pagination entity.Pagination, tag entity.TagOption) (items []T, err error)
	Get(ctx context.Context, gatewayID string, nameOrID string) (item *T, err error)
}
