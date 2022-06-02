package utilstest

import (
	"github.com/lffranca/bento/pkg/domain/entity"
)

func GetDummyGateway() *entity.Gateway {
	return &entity.Gateway{
		ID:   TestFakeID,
		Name: "Gateway Name",
	}
}
