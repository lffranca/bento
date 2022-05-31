package utilstest

import (
	"context"

	"github.com/lffranca/bento/pkg/domain/entity"
	"github.com/stretchr/testify/mock"
)

type ProxyProviderMock struct {
	mock.Mock
}

func (pkg *ProxyProviderMock) Get(ctx context.Context, nameOrID string) (item *entity.Proxy, err error) {
	args := pkg.MethodCalled(
		"Get",
		ctx,
		nameOrID,
	)

	if len(args) > 0 && args.Get(0) != nil {
		item = args.Get(0).(*entity.Proxy)
	}

	if len(args) > 1 && args.Get(1) != nil {
		err = args.Error(1)
	}

	return
}

func (pkg *ProxyProviderMock) List(ctx context.Context, pagination entity.Pagination, tag entity.TagOption) (items []entity.Proxy, err error) {
	args := pkg.MethodCalled(
		"List",
		ctx,
		pagination,
		tag,
	)

	if len(args) > 0 && args.Get(0) != nil {
		items = args.Get(0).([]entity.Proxy)
	}

	if len(args) > 1 && args.Get(1) != nil {
		err = args.Error(1)
	}

	return
}
