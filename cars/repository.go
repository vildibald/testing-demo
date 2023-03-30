package cars

import (
	"errors"
	"github.com/google/uuid"
)

//go:generate mockgen -source=repository.go -destination=repository_mock.go -package=cars

type Repository interface {
	Create(entity *Car) (*Car, error)
	Get(id string) (*Car, error)
	List() ([]Car, error)
}

var (
	ErrNotFound = errors.New("car not found")
)

type DummyRepository struct {
	storage map[string]Car
}

func (d DummyRepository) Create(entity *Car) (*Car, error) {
	id := uuid.NewString()
	result := *entity
	result.Id = id

	d.storage[id] = result

	return &result, nil
}

func (d DummyRepository) Get(id string) (*Car, error) {
	car, ok := d.storage[id]
	if !ok {
		return nil, ErrNotFound
	}

	return &car, nil
}

func (d DummyRepository) List() ([]Car, error) {
	result := make([]Car, 0, len(d.storage))

	for _, car := range d.storage {
		result = append(result, car)
	}

	return result, nil
}

func NewDummyRepository() *DummyRepository {
	return &DummyRepository{storage: make(map[string]Car)}
}

// Ensure interface implementation
var _ Repository = (*DummyRepository)(nil)
