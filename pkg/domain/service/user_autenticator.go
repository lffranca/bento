package service

import (
	"context"

	"github.com/lffranca/bento/pkg/domain"
	"github.com/lffranca/bento/pkg/domain/entity"
	"github.com/lffranca/bento/pkg/domain/provider"
)

func NewUserAuthenticatorService(
	base *BaseService,
	authProvider provider.AuthenticationProvider,
	userProvider provider.UserProvider,
) domain.UserAuthenticator {
	return &userAuthenticatorService{
		base,
		authProvider,
		userProvider,
	}
}

type userAuthenticatorService struct {
	*BaseService
	authProvider provider.AuthenticationProvider
	userProvider provider.UserProvider
}

func (pkg *userAuthenticatorService) AuthenticateUser(ctx context.Context, email, password string) (userAuthentication *entity.UserAuthentication, err error) {
	var user *entity.User

	userAuthentication, err = pkg.authProvider.SignIn(ctx, email, password)
	if err != nil {
		return nil, err
	}

	user, err = pkg.userProvider.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if user != nil {
		userAuthentication.User = *user
	}

	return userAuthentication, nil
}
