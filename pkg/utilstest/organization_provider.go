package utilstest

import (
	"context"

	"github.com/lffranca/bento/pkg/domain/entity"
	"github.com/stretchr/testify/mock"
)

type OrganizationProviderMock struct {
	mock.Mock
}

func (pkg *OrganizationProviderMock) Get(ctx context.Context) (item *entity.Organization, err error) {
	args := pkg.MethodCalled(
		"Get",
		ctx,
	)

	if len(args) > 0 && args.Get(0) != nil {
		item = args.Get(0).(*entity.Organization)
	}

	if len(args) > 1 && args.Get(1) != nil {
		err = args.Error(1)
	}

	return
}
