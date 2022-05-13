package domain

import (
	"context"

	"github.com/lffranca/bento/pkg/domain/entity"
)

type UserAuthenticator interface {
	AuthenticateUser(ctx context.Context, email, password string) (userAuthentication *entity.UserAuthentication, err error)
}
