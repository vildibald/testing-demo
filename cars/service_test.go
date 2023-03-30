package cars

import (
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ServiceTestSuite struct {
	suite.Suite

	r *MockRepository
	s *Service
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

func (s *ServiceTestSuite) SetupTest() {
	mockCtrl := gomock.NewController(s.T())
	defer mockCtrl.Finish()
	s.r = NewMockRepository(mockCtrl)
	s.s = &Service{Repository: s.r}
}

func (s *ServiceTestSuite) TestService_Create() {
	s.Run("should create a car", func() {
		car := &Car{
			Brand: "Ford",
			Model: "Mustang",
		}

		expected := &Car{
			Id:    uuid.NewString(),
			Brand: "Ford",
			Model: "Mustang",
		}

		s.r.EXPECT().Create(car).Return(expected, nil)
		actual, err := s.s.Create(car)

		s.NoError(err)
		s.Equal(expected, actual)
	})

	s.Run("should return error if id is not empty", func() {
		car := &Car{
			Id:    uuid.NewString(),
			Brand: "Ford",
			Model: "Mustang",
		}

		actual, err := s.s.Create(car)

		s.Nil(actual)
		s.Equal(ErrIdNotEmpty, err)
	})
}
