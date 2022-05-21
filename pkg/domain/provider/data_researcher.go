package provider

import (
	"context"

	"github.com/lffranca/bento/pkg/domain/entity"
)

type DataResearcherProvider[T entity.SearchResponse] interface {
	List(ctx context.Context, gateway *entity.Gateway, offset, limit int, tags []string, matchAllTags bool) (result T, err error)
	Get(ctx context.Context, gateway *entity.Gateway, proxyID string, nameOrID string) (items T, err error)
}
