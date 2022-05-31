package service

import (
	"context"
	"testing"

	"github.com/lffranca/bento/pkg/utilstest"
	"github.com/stretchr/testify/suite"
)

func TestOrganizationGetterServiceTestSuite(t *testing.T) {
	suite.Run(t, new(organizationGetterServiceTestSuite))
}

type organizationGetterServiceTestSuite struct {
	suite.Suite
	orgProvider *utilstest.OrganizationProviderMock
	service     *OrganizationGetterService
	context     context.Context
}

func (test *organizationGetterServiceTestSuite) setup() {
	test.orgProvider = new(utilstest.OrganizationProviderMock)
	test.service = NewOrganizationGetterService(
		test.orgProvider,
	).(*OrganizationGetterService)
	test.context = context.Background()
}

func (test *organizationGetterServiceTestSuite) TestGet_WhenOk() {
	test.setup()

	test.orgProvider.
		On("Get", test.context).
		Return(utilstest.GetDummyOrg())

	data, err := test.service.Get(test.context)
	test.Nil(err)
	test.Equal(utilstest.GetDummyOrg(), data)
}
