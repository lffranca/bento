package service

import (
	"context"
	"errors"
	"testing"

	"github.com/lffranca/bento/pkg/domain/entity"
	"github.com/lffranca/bento/pkg/utilstest"
	"github.com/stretchr/testify/suite"
)

func TestServiceGetterService(t *testing.T) {
	suite.Run(t, new(serviceGetterServiceTestSuite))
}

type serviceGetterServiceTestSuite struct {
	suite.Suite
	gatewayProvider *utilstest.GatewayProviderMock
	serviceProvider *utilstest.DomainProviderMock[entity.Service]
	service         *serviceGetterService
	context         context.Context
}

func (test *serviceGetterServiceTestSuite) setup() {
	test.gatewayProvider = new(utilstest.GatewayProviderMock)
	test.serviceProvider = new(utilstest.DomainProviderMock[entity.Service])
	test.service = NewServiceGetterService(
		test.gatewayProvider,
		test.serviceProvider,
	).(*serviceGetterService)
	test.context = context.Background()
}

func (test *serviceGetterServiceTestSuite) TestGet_WhenOk() {
	test.setup()

	serviceItem := utilstest.GetDummyService()
	gateway := utilstest.GetDummyGateway()

	test.gatewayProvider.
		On("Get", test.context, gateway.ID).
		Return(gateway)

	test.serviceProvider.
		On("Get", test.context, gateway, serviceItem.ID).
		Return(serviceItem)

	data, err := test.service.Get(test.context, utilstest.TestFakeID, serviceItem.ID)
	test.Nil(err)
	test.Equal(serviceItem, data)
}

func (test *serviceGetterServiceTestSuite) TestGet_WithErrorInGateway() {
	test.setup()

	err := errors.New("")

	test.gatewayProvider.
		On("Get", test.context, utilstest.TestFakeID).
		Return(nil, err)

	data, returnErr := test.service.Get(test.context, utilstest.TestFakeID, utilstest.TestFakeID)
	test.Nil(data)
	test.NotNil(returnErr)
	test.Equal(err, returnErr)
}

func (test *serviceGetterServiceTestSuite) TestList_WhenOk() {
	test.setup()

	quantity := 10

	gateway := utilstest.GetDummyGateway()
	pagination := utilstest.GetDummyPagination()
	tag := utilstest.GetDummyTagOption()

	test.gatewayProvider.
		On("Get", test.context, gateway.ID).
		Return(gateway)

	test.serviceProvider.
		On("List", test.context, gateway, pagination, tag).
		Return(utilstest.GetDummyServices(quantity))

	data, err := test.service.List(test.context, utilstest.TestFakeID, pagination, tag)
	test.Nil(err)
	test.Equal(utilstest.GetDummyServices(quantity), data)
}

func (test *serviceGetterServiceTestSuite) TestList_WithErrorInGateway() {
	test.setup()

	err := errors.New("")

	test.gatewayProvider.
		On("Get", test.context, utilstest.TestFakeID).
		Return(nil, err)

	data, returnErr := test.service.List(test.context, utilstest.TestFakeID, utilstest.GetDummyPagination(), utilstest.GetDummyTagOption())
	test.Nil(data)
	test.NotNil(returnErr)
	test.Equal(err, returnErr)
}
