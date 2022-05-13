package provider

import (
	"context"

	"github.com/lffranca/bento/pkg/domain/entity"
)

type UserProvider interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
}
