package service

import (
	"context"

	"github.com/lffranca/bento/pkg/domain"
	"github.com/lffranca/bento/pkg/domain/entity"
	"github.com/lffranca/bento/pkg/domain/provider"
)

func NewUserGetterService(userProvider provider.UserProvider) domain.UserGetterService {
	return &UserGetterService{userProvider}
}

type UserGetterService struct {
	userProvider provider.UserProvider
}

func (pkg *UserGetterService) List(ctx context.Context, pagination entity.Pagination) ([]entity.User, error) {
	return pkg.userProvider.List(ctx, pagination)
}

func (pkg *UserGetterService) Get(ctx context.Context, nameOrID string) (*entity.User, error) {
	return pkg.userProvider.Get(ctx, nameOrID)
}
