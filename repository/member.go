package repository

import (
	"context"

	"github.com/Code0716/graphql-study/domain"
)

// PersonsInterface  is data access methods to Persons.
type PersonsInterface interface {
	CreatePerson(ctx context.Context, person domain.Person) (*domain.Person, error)
	GetAllPersons(ctx context.Context, params domain.Pager) ([]*domain.Person, error)
	GetPerson(ctx context.Context, params domain.GetPersonParams) (*domain.Person, error)
}

// PersonsRepository is person repository.
type PersonsRepository struct {
	SQLHandler SQLHandlerInterface
}

// NewPersons initializes persons repository.
func NewPersons(sqlHandler SQLHandlerInterface) *PersonsRepository {
	return &PersonsRepository{
		sqlHandler,
	}
}

// CreatePerson regist person to persons db
func (r *PersonsRepository) CreatePerson(ctx context.Context, person domain.Person) (*domain.Person, error) {
	err := r.SQLHandler.Create(&person)
	if err != nil {
		return nil, err
	}
	return &person, nil
}

// GetPerson return persons found by params
func (r *PersonsRepository) GetPerson(ctx context.Context, params domain.GetPersonParams) (*domain.Person, error) {
	var person *domain.Person
	err := r.SQLHandler.First(&person, params)

	if err != nil {
		return nil, err
	}

	return person, nil
}

// GetAllPersons return persons found by params
func (r *PersonsRepository) GetAllPersons(ctx context.Context, params domain.Pager) ([]*domain.Person, error) {
	var persons []*domain.Person
	err := r.SQLHandler.Find(&persons, params)

	if err != nil {
		return nil, err
	}

	return persons, nil
}
