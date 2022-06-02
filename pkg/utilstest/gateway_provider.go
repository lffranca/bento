package utilstest

import (
	"context"

	"github.com/lffranca/bento/pkg/domain/entity"
	"github.com/stretchr/testify/mock"
)

type GatewayProviderMock struct {
	mock.Mock
}

func (pkg *GatewayProviderMock) Get(ctx context.Context, id string) (item *entity.Gateway, err error) {
	args := pkg.MethodCalled(
		"Get",
		ctx,
		id,
	)

	if len(args) > 0 && args.Get(0) != nil {
		item = args.Get(0).(*entity.Gateway)
	}

	if len(args) > 1 && args.Get(1) != nil {
		err = args.Error(1)
	}

	return
}
