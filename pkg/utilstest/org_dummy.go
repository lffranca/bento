package utilstest

import (
	"github.com/lffranca/bento/pkg/domain/entity"
)

func GetDummyOrg() *entity.Organization {
	return &entity.Organization{
		ID:   TestFakeID,
		Name: "Organization",
	}
}
