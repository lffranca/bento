package gkong

import (
	"github.com/kong/go-kong/kong"
	"github.com/lffranca/bento/pkg/domain/entity"
	"github.com/lffranca/bento/pkg/utils"
)

func MapperServiceToEntity(value kong.Service) entity.Service {
	return entity.Service{
		ID:        utils.FromPointer[string](value.ID),
		Name:      utils.FromPointer[string](value.Name),
		Host:      utils.FromPointer[string](value.Host),
		Port:      utils.FromPointer[int](value.Port),
		Protocol:  utils.FromPointer[string](value.Protocol),
		URL:       utils.FromPointer[string](value.URL),
		Tags:      utils.FromArrayPointer[string](value.Tags),
		CreatedAt: utils.FromIntToTime(value.CreatedAt),
		UpdatedAt: utils.FromIntToTime(value.UpdatedAt),
	}
}
