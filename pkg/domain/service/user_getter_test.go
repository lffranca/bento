package service

import (
	"context"
	"testing"

	"github.com/lffranca/bento/pkg/utilstest"
	"github.com/stretchr/testify/suite"
)

func TestUserGetterService(t *testing.T) {
	suite.Run(t, new(userGetterServiceTestSuite))
}

type userGetterServiceTestSuite struct {
	suite.Suite
	userProvider *utilstest.UserProviderMock
	service      *UserGetterService
	context      context.Context
}

func (test *userGetterServiceTestSuite) setup() {
	test.userProvider = new(utilstest.UserProviderMock)
	test.service = NewUserGetterService(
		test.userProvider,
	).(*UserGetterService)
	test.context = context.Background()
}

func (test *userGetterServiceTestSuite) TestGet_WhenOk() {
	test.setup()

	test.userProvider.
		On("Get", test.context, utilstest.TestFakeID).
		Return(utilstest.GetDummyUser())

	data, err := test.service.Get(test.context, utilstest.TestFakeID)
	test.Nil(err)
	test.Equal(utilstest.GetDummyUser(), data)
}

func (test *userGetterServiceTestSuite) TestList_WhenOk() {
	test.setup()

	quantity := 10

	test.userProvider.
		On("List", test.context, utilstest.GetDummyPagination()).
		Return(utilstest.GetDummyUsers(quantity))

	data, err := test.service.List(test.context, utilstest.GetDummyPagination())
	test.Nil(err)
	test.Equal(utilstest.GetDummyUsers(quantity), data)
}
