package provider

import (
	"context"

	"github.com/lffranca/bento/pkg/domain/entity"
)

type OrganizationProvider interface {
	Get(ctx context.Context) (item *entity.Organization, err error)
}
