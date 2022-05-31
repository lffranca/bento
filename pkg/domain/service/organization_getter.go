package service

import (
	"context"

	"github.com/lffranca/bento/pkg/domain"
	"github.com/lffranca/bento/pkg/domain/entity"
	"github.com/lffranca/bento/pkg/domain/provider"
)

func NewOrganizationGetterService(orgProvider provider.OrganizationProvider) domain.OrganizationGetterService {
	return &OrganizationGetterService{orgProvider}
}

type OrganizationGetterService struct {
	orgProvider provider.OrganizationProvider
}

func (pkg *OrganizationGetterService) Get(ctx context.Context) (*entity.Organization, error) {
	return pkg.orgProvider.Get(ctx)
}
