package service

import (
	"context"
	"testing"

	"github.com/lffranca/bento/pkg/utilstest"
	"github.com/stretchr/testify/suite"
)

func TestProxyGetterService(t *testing.T) {
	suite.Run(t, new(proxyGetterServiceTestSuite))
}

type proxyGetterServiceTestSuite struct {
	suite.Suite
	proxyProvider *utilstest.ProxyProviderMock
	service       *ProxyGetterService
	context       context.Context
}

func (test *proxyGetterServiceTestSuite) setup() {
	test.proxyProvider = new(utilstest.ProxyProviderMock)
	test.service = NewProxyGetterService(
		test.proxyProvider,
	).(*ProxyGetterService)
	test.context = context.Background()
}

func (test *proxyGetterServiceTestSuite) TestGet_WhenOk() {
	test.setup()

	test.proxyProvider.
		On("Get", test.context, utilstest.TestFakeID).
		Return(utilstest.GetDummyProxy())

	data, err := test.service.Get(test.context, utilstest.TestFakeID)
	test.Nil(err)
	test.Equal(utilstest.GetDummyProxy(), data)
}

func (test *proxyGetterServiceTestSuite) TestList_WhenOk() {
	test.setup()

	quantity := 10

	test.proxyProvider.
		On("List", test.context, utilstest.GetDummyPagination(), utilstest.GetDummyTagOption()).
		Return(utilstest.GetDummyProxies(quantity))

	data, err := test.service.List(test.context, utilstest.GetDummyPagination(), utilstest.GetDummyTagOption())
	test.Nil(err)
	test.Equal(utilstest.GetDummyProxies(quantity), data)
}
