package domain

import (
	"context"

	"github.com/lffranca/bento/pkg/domain/entity"
)

type OrganizationGetterService interface {
	Get(ctx context.Context) (item *entity.Organization, err error)
}
