package provider

import (
	"context"

	"github.com/lffranca/bento/pkg/domain/entity"
)

type GatewayProvider interface {
	Get(ctx context.Context, id string) (gateway *entity.Gateway, err error)
}
