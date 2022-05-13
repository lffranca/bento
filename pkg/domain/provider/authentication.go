package provider

import (
	"context"

	"github.com/lffranca/bento/pkg/domain/entity"
)

type AuthenticationProvider interface {
	SignIn(ctx context.Context, email, password string) (*entity.UserAuthentication, error)
}
