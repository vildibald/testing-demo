package cars

import (
	"errors"
)

type Service struct {
	Repository Repository
}

var (
	ErrIdNotEmpty = errors.New("id must be empty")
	ErrEmptyField = errors.New("field must not be empty")
)

func (s *Service) Create(car *Car) (*Car, error) {
	if car.Id != "" {
		return nil, ErrIdNotEmpty
	}

	if car.Brand == "" {
		return nil, ErrEmptyField
	}

	if car.Model == "" {
		return nil, ErrEmptyField
	}

	result, err := s.Repository.Create(car)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Service) Get(id string) (*Car, error) {
	return s.Repository.Get(id)
}

func (s *Service) List() ([]Car, error) {
	return s.Repository.List()
}
