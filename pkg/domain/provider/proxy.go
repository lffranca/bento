package provider

import (
	"context"

	"github.com/lffranca/bento/pkg/domain/entity"
)

type ProxyProvider interface {
	Get(ctx context.Context, id string) (proxy *entity.Proxy, err error)
}
