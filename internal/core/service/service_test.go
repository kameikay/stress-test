package service

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	mock "github.com/kameikay/stress-test/internal/mocks"
	"github.com/stretchr/testify/suite"
)

type ServiceTestSuite struct {
	suite.Suite
	ctrl       *gomock.Controller
	repository *mock.MockStressTestRepository
}

func TestServiceStart(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

func (suite *ServiceTestSuite) ServiceTestSuiteDown() {
	defer suite.ctrl.Finish()
}

func (suite *ServiceTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.repository = mock.NewMockStressTestRepository(suite.ctrl)
}

func (suite *ServiceTestSuite) TestNewService() {
	service := New(suite.repository)
	suite.NotNil(service)
}

func (suite *ServiceTestSuite) TestExecute() {
	service := New(suite.repository)

	testCases := []struct {
		name          string
		serverHandler http.HandlerFunc
		expectations  func(repo *mock.MockStressTestRepository)
	}{
		{
			name: "All status 200",
			serverHandler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			},
			expectations: func(repo *mock.MockStressTestRepository) {
				repo.EXPECT().Save(200).Times(10)
				repo.EXPECT().GetTotalRequests().Return(10)
				repo.EXPECT().GetStatus200().Return(10)
				repo.EXPECT().GetOthersStatus().Return(map[int]int{})
			},
		},
		{
			name: "status 404",
			serverHandler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNotFound)
			},
			expectations: func(repo *mock.MockStressTestRepository) {
				repo.EXPECT().Save(404).Times(10)
				repo.EXPECT().GetTotalRequests().Return(10)
				repo.EXPECT().GetStatus200().Return(0)
				repo.EXPECT().GetOthersStatus().Return(map[int]int{404: 10})
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			server := httptest.NewServer(tc.serverHandler)
			defer server.Close()

			tc.expectations(suite.repository)
			service.Execute(server.URL, 10, 1)
		})
	}
}
