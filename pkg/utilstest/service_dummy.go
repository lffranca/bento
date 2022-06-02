package utilstest

import (
	"github.com/lffranca/bento/pkg/domain/entity"
)

func GetDummyService() *entity.Service {
	return &entity.Service{
		ID:   TestFakeID,
		Name: "Service Name",
	}
}

func GetDummyServices(quantity int) (items []entity.Service) {
	for i := 0; i < quantity; i++ {
		items = append(items, *GetDummyService())
	}

	return
}
