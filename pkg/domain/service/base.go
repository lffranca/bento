package service

import (
	"github.com/lffranca/bento/pkg/domain/provider"
)

func NewBaseService(
	log provider.LogProvider,
	metric provider.MetricProvider,
) *BaseService {
	return &BaseService{
		log,
		metric,
	}
}

type BaseService struct {
	log    provider.LogProvider
	metric provider.MetricProvider
}
