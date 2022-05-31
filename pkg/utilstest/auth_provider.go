package utilstest

import (
	"context"

	"github.com/lffranca/bento/pkg/domain/entity"
	"github.com/stretchr/testify/mock"
)

type AuthProviderMock struct {
	mock.Mock
}

func (pkg *AuthProviderMock) SignIn(ctx context.Context, email, password string) (item *entity.UserAuthentication, err error) {
	args := pkg.MethodCalled(
		"SignIn",
		ctx,
		email, password,
	)

	if len(args) > 0 && args.Get(0) != nil {
		item = args.Get(0).(*entity.UserAuthentication)
	}

	if len(args) > 1 && args.Get(1) != nil {
		err = args.Error(1)
	}

	return
}
