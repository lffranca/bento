package gkong

import (
	"github.com/kong/go-kong/kong"
	"github.com/lffranca/bento/pkg/domain/entity"
	"github.com/lffranca/bento/pkg/utils"
)

func MapperRouteToEntity(value kong.Route) entity.Route {
	return entity.Route{
		ID:        utils.FromPointer[string](value.ID),
		Name:      utils.FromPointer[string](value.Name),
		Headers:   value.Headers,
		Methods:   utils.FromArrayPointer[string](value.Methods),
		Paths:     utils.FromArrayPointer[string](value.Paths),
		Protocols: utils.FromArrayPointer[string](value.Protocols),
		Tags:      utils.FromArrayPointer[string](value.Tags),
		CreatedAt: utils.FromIntToTime(value.CreatedAt),
		UpdatedAt: utils.FromIntToTime(value.UpdatedAt),
	}
}

func MapperRoutesToEntities(values []kong.Route) (items []entity.Route) {
	for _, value := range values {
		items = append(items, MapperRouteToEntity(value))
	}

	return
}
