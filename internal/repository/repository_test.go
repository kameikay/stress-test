package repository

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type RequestsMemoryTestSuite struct {
	suite.Suite
	ctrl *gomock.Controller
}

func TestRequestsMemoryStart(t *testing.T) {
	suite.Run(t, new(RequestsMemoryTestSuite))
}

func (suite *RequestsMemoryTestSuite) RequestsMemoryTestSuiteDown() {
	defer suite.ctrl.Finish()
}

func (suite *RequestsMemoryTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
}

func (suite *RequestsMemoryTestSuite) TestNewRequestsMemory() {
	requestsMemory := NewRequestsMemory()
	suite.NotNil(requestsMemory)
}

func (suite *RequestsMemoryTestSuite) TestSave() {
	requestsMemory := NewRequestsMemory()
	requestsMemory.Save(200)
	requestsMemory.Save(200)
	requestsMemory.Save(404)

	suite.Equal(2, requestsMemory.Requests[200])
	suite.Equal(1, requestsMemory.Requests[404])
}

func (suite *RequestsMemoryTestSuite) TestGetTotalRequests() {
	requestsMemory := NewRequestsMemory()
	requestsMemory.Save(200)
	requestsMemory.Save(200)
	requestsMemory.Save(404)

	suite.Equal(3, requestsMemory.GetTotalRequests())
}

func (suite *RequestsMemoryTestSuite) TestGetStatus200() {
	requestsMemory := NewRequestsMemory()
	requestsMemory.Save(200)
	requestsMemory.Save(200)
	requestsMemory.Save(404)

	suite.Equal(2, requestsMemory.GetStatus200())
}

func (suite *RequestsMemoryTestSuite) TestGetOthersStatus() {
	requestsMemory := NewRequestsMemory()
	requestsMemory.Save(200)
	requestsMemory.Save(200)
	requestsMemory.Save(404)

	others := requestsMemory.GetOthersStatus()
	suite.Equal(1, others[404])
}

func (suite *RequestsMemoryTestSuite) TestGetErrorRequests() {
	requestsMemory := NewRequestsMemory()
	requestsMemory.Save(0)
	requestsMemory.Save(0)
	requestsMemory.Save(404)

	suite.Equal(2, requestsMemory.GetErrorRequests())
}
