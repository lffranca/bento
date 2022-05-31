package utilstest

import (
	"context"

	"github.com/lffranca/bento/pkg/domain/entity"
	"github.com/stretchr/testify/mock"
)

type UserProviderMock struct {
	mock.Mock
}

func (pkg *UserProviderMock) GetUserByEmail(ctx context.Context, email string) (item *entity.User, err error) {
	args := pkg.MethodCalled(
		"GetUserByEmail",
		ctx,
		email,
	)

	if len(args) > 0 && args.Get(0) != nil {
		item = args.Get(0).(*entity.User)
	}

	if len(args) > 1 && args.Get(1) != nil {
		err = args.Error(1)
	}

	return
}

func (pkg *UserProviderMock) List(ctx context.Context, pagination entity.Pagination) (items []entity.User, err error) {
	args := pkg.MethodCalled(
		"List",
		ctx,
		pagination,
	)

	if len(args) > 0 && args.Get(0) != nil {
		items = args.Get(0).([]entity.User)
	}

	if len(args) > 1 && args.Get(1) != nil {
		err = args.Error(1)
	}

	return
}

func (pkg *UserProviderMock) Get(ctx context.Context, nameOrID string) (item *entity.User, err error) {
	args := pkg.MethodCalled(
		"Get",
		ctx,
		nameOrID,
	)

	if len(args) > 0 && args.Get(0) != nil {
		item = args.Get(0).(*entity.User)
	}

	if len(args) > 1 && args.Get(1) != nil {
		err = args.Error(1)
	}

	return
}
