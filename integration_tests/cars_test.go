//go:build integration

package integration_tests

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
	"testing-demo/cars"
)

type CarTestSuite struct {
	suite.Suite

	client *http.Client
}

func TestCarTestSuite(t *testing.T) {
	suite.Run(t, new(CarTestSuite))
}

func (s *CarTestSuite) SetupSuite() {
	s.client = &http.Client{}
}

func (s *CarTestSuite) TestCar_Create() {
	s.Run("should create a car", func() {
		car := &cars.Car{
			Brand: "Ford",
			Model: "Mustang",
		}

		reqPayload, err := json.Marshal(car)
		s.NoError(err)

		req, err := http.NewRequest(http.MethodPut, "http://localhost:8080/cars", bytes.NewBuffer(reqPayload))
		s.NoError(err)

		req.Header.Set("Content-Type", "application/json")

		res, err := s.client.Do(req)
		s.NoError(err)
		s.Equal(http.StatusOK, res.StatusCode)

		actual := &cars.Car{}
		err = json.NewDecoder(res.Body).Decode(actual)
		s.NoError(err)
		s.NotEmpty(actual.Id)
		s.Equal(car.Brand, actual.Brand)
		s.Equal(car.Model, actual.Model)
	})

	s.Run("should return error if id is not empty", func() {
		car := &cars.Car{
			Id:    uuid.NewString(),
			Brand: "Ford",
			Model: "Mustang",
		}

		reqPayload, err := json.Marshal(car)
		s.NoError(err)

		req, err := http.NewRequest(http.MethodPut, "http://localhost:8080/cars", bytes.NewBuffer(reqPayload))
		s.NoError(err)

		req.Header.Set("Content-Type", "application/json")

		res, err := s.client.Do(req)
		s.NoError(err)
		s.Equal(http.StatusBadRequest, res.StatusCode)
	})
}
