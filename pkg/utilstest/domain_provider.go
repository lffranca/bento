package utilstest

import (
	"context"

	"github.com/lffranca/bento/pkg/domain/entity"
	"github.com/stretchr/testify/mock"
)

type DomainProviderMock[T any] struct {
	mock.Mock
}

func (pkg *DomainProviderMock[T]) List(
	ctx context.Context,
	gateway *entity.Gateway,
	pagination entity.Pagination,
	tag entity.TagOption,
) (items []T, err error) {
	args := pkg.MethodCalled(
		"List",
		ctx,
		gateway,
		pagination,
		tag,
	)

	if len(args) > 0 && args.Get(0) != nil {
		items = args.Get(0).([]T)
	}

	if len(args) > 1 && args.Get(1) != nil {
		err = args.Error(1)
	}

	return
}

func (pkg *DomainProviderMock[T]) Get(ctx context.Context, gateway *entity.Gateway, nameOrID string) (item *T, err error) {
	args := pkg.MethodCalled(
		"Get",
		ctx,
		gateway,
		nameOrID,
	)

	if len(args) > 0 && args.Get(0) != nil {
		item = args.Get(0).(*T)
	}

	if len(args) > 1 && args.Get(1) != nil {
		err = args.Error(1)
	}

	return
}
