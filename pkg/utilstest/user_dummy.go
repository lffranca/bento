package utilstest

import (
	"github.com/lffranca/bento/pkg/domain/entity"
)

func GetDummyUser() *entity.User {
	return &entity.User{
		ID:           TestFakeID,
		Name:         "Fulano de Souza",
		Email:        TestFakeEmail,
		Organization: *GetDummyOrg(),
	}
}

func GetDummyUsers(quantity int) (items []entity.User) {
	for i := 0; i < quantity; i++ {
		items = append(items, *GetDummyUser())
	}

	return
}
