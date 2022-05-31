package service

import (
	"context"

	"github.com/lffranca/bento/pkg/domain"
	"github.com/lffranca/bento/pkg/domain/entity"
	"github.com/lffranca/bento/pkg/domain/provider"
)

func NewUserAuthenticatorService(
	authProvider provider.AuthenticationProvider,
	userProvider provider.UserProvider,
) domain.UserAuthenticator {
	return &userAuthenticatorService{
		authProvider,
		userProvider,
	}
}

type userAuthenticatorService struct {
	authProvider provider.AuthenticationProvider
	userProvider provider.UserProvider
}

func (pkg *userAuthenticatorService) AuthenticateUser(ctx context.Context, email, password string) (userAuth *entity.UserAuthentication, err error) {
	var user *entity.User

	userAuth, err = pkg.authProvider.SignIn(ctx, email, password)
	if err != nil {
		return nil, err
	}

	user, err = pkg.userProvider.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if user != nil {
		userAuth.User = *user
	}

	return userAuth, nil
}
