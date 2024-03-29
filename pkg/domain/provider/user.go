package provider

import (
	"context"

	"github.com/lffranca/bento/pkg/domain/entity"
)

type UserProvider interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	List(ctx context.Context, pagination entity.Pagination) (items []entity.User, err error)
	Get(ctx context.Context, nameOrID string) (item *entity.User, err error)
}
