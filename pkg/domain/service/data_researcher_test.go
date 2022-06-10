package service

import (
	"context"
	"errors"
	"testing"

	"github.com/lffranca/bento/pkg/domain"
	"github.com/lffranca/bento/pkg/domain/entity"
	"github.com/lffranca/bento/pkg/utilstest"
	"github.com/stretchr/testify/suite"
)

func TestDataResearcherService(t *testing.T) {
	suite.Run(t, new(DataResearcherServiceTestSuite[entity.Service]))
}

type DataResearcherServiceTestSuite[T any] struct {
	suite.Suite
	gatewayProvider *utilstest.GatewayProviderMock
	serviceProvider *utilstest.DomainProviderMock[T]
	service         domain.DataResearcherService[T]
	context         context.Context
}

func (test *DataResearcherServiceTestSuite[T]) setup() {
	test.gatewayProvider = new(utilstest.GatewayProviderMock)
	test.serviceProvider = new(utilstest.DomainProviderMock[T])
	test.service = NewDataResearcherService[T](
		test.gatewayProvider,
		test.serviceProvider,
	)
	test.context = context.Background()
}

func (test *DataResearcherServiceTestSuite[T]) TestGet_WhenOk() {
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

func (test *DataResearcherServiceTestSuite[T]) TestGet_WithErrorInGateway() {
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

func (test *DataResearcherServiceTestSuite[T]) TestList_WhenOk() {
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

func (test *DataResearcherServiceTestSuite[T]) TestList_WithErrorInGateway() {
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
