package service

import (
	"context"
	"errors"
	"testing"

	"github.com/lffranca/bento/pkg/utilstest"
	"github.com/stretchr/testify/suite"
)

func TestUserAuthenticatorTestSuite(t *testing.T) {
	suite.Run(t, new(userAuthenticatorTestSuite))
}

type userAuthenticatorTestSuite struct {
	suite.Suite
	userProvider *utilstest.UserProviderMock
	authProvider *utilstest.AuthProviderMock
	service      *userAuthenticatorService
	context      context.Context
}

func (test *userAuthenticatorTestSuite) setup() {
	test.userProvider = new(utilstest.UserProviderMock)
	test.authProvider = new(utilstest.AuthProviderMock)
	test.service = NewUserAuthenticatorService(
		test.authProvider,
		test.userProvider,
	).(*userAuthenticatorService)
	test.context = context.Background()
}

func (test *userAuthenticatorTestSuite) TestAuthenticateUser_WhenOk() {
	test.setup()

	test.authProvider.
		On("SignIn", test.context, utilstest.TestFakeEmail, utilstest.TestFakePassword).
		Return(utilstest.GetDummyUserAuth())

	test.userProvider.
		On("GetUserByEmail", test.context, utilstest.TestFakeEmail).
		Return(utilstest.GetDummyUser())

	data, err := test.service.AuthenticateUser(test.context, utilstest.TestFakeEmail, utilstest.TestFakePassword)
	test.Nil(err)
	test.Equal(utilstest.GetDummyUserAuthWithUser(), data)
}

func (test *userAuthenticatorTestSuite) TestAuthenticateUser_WithErrorInSign() {
	test.setup()

	err := errors.New("")

	test.authProvider.
		On("SignIn", test.context, utilstest.TestFakeEmail, utilstest.TestFakePassword).
		Return(nil, err)

	data, resultErr := test.service.AuthenticateUser(test.context, utilstest.TestFakeEmail, utilstest.TestFakePassword)
	test.Nil(data)
	test.NotNil(resultErr)
	test.Equal(err, resultErr)
}

func (test *userAuthenticatorTestSuite) TestAuthenticateUser_WithErrorInGetUserByEmail() {
	test.setup()

	err := errors.New("")

	test.authProvider.
		On("SignIn", test.context, utilstest.TestFakeEmail, utilstest.TestFakePassword).
		Return(utilstest.GetDummyUserAuth())

	test.userProvider.
		On("GetUserByEmail", test.context, utilstest.TestFakeEmail).
		Return(nil, err)

	data, resultErr := test.service.AuthenticateUser(test.context, utilstest.TestFakeEmail, utilstest.TestFakePassword)
	test.Nil(data)
	test.NotNil(resultErr)
	test.Equal(err, resultErr)
}
