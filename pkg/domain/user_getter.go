package domain

import (
	"context"

	"github.com/lffranca/bento/pkg/domain/entity"
)

type UserGetterService interface {
	List(ctx context.Context, pagination entity.Pagination) (items []entity.User, err error)
	Get(ctx context.Context, nameOrID string) (item *entity.User, err error)
}
