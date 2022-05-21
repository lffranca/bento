package domain

import (
	"context"

	"github.com/lffranca/bento/pkg/domain/entity"
)

type Item interface {
	comparable
}

type DataResearcherService[T entity.SearchResponse] interface {
	List(ctx context.Context, gatewayID string, offset, limit int, tags []string, matchAllTags bool) (items []T, err error)
	Get(ctx context.Context, gatewayID string, nameOrID string) (items *T, err error)
}
