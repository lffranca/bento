package utilstest

import (
	"github.com/lffranca/bento/pkg/domain/entity"
)

func GetDummyProxy() *entity.Proxy {
	return &entity.Proxy{
		ID:   TestFakeID,
		Name: "Proxy Name",
	}
}

func GetDummyProxies(quantity int) (items []entity.Proxy) {
	for i := 0; i < quantity; i++ {
		items = append(items, *GetDummyProxy())
	}

	return
}
