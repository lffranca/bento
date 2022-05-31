package utilstest

import (
	"github.com/lffranca/bento/pkg/domain/entity"
)

func GetDummyUserAuth() *entity.UserAuthentication {
	return &entity.UserAuthentication{
		AccessToken:  TestFakeJWT,
		IDToken:      TestFakeJWT,
		RefreshToken: TestFakeJWT,
	}
}

func GetDummyUserAuthWithUser() *entity.UserAuthentication {
	return &entity.UserAuthentication{
		AccessToken:  TestFakeJWT,
		IDToken:      TestFakeJWT,
		RefreshToken: TestFakeJWT,
		User:         *GetDummyUser(),
	}
}
