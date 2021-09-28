package repository

import (
	"context"
	"fmt"

	"github.com/Code0716/graphql-study/domain"
)

// TestimonyInterface  is data access methods to Testimony.
type TestimonyInterface interface {
	CreateTestimony(ctx context.Context, testimony domain.Testimony) (*domain.Testimony, error)
	GetAllTestimony(ctx context.Context, pager domain.Pager, params domain.GetTestimonyParams) ([]*domain.Testimony, error)
	GetTestimony(ctx context.Context, params domain.GetTestimonyParams) (*domain.Testimony, error)
}

// TestimonyRepository is testimony repository.
type TestimonyRepository struct {
	SQLHandler SQLHandlerInterface
}

// NewTestimony initializes testimonies repository.
func NewTestimony(sqlHandler SQLHandlerInterface) *TestimonyRepository {
	return &TestimonyRepository{
		sqlHandler,
	}
}

// CreateTestimony regist testimony to testimonies db
func (r *TestimonyRepository) CreateTestimony(ctx context.Context, testimony domain.Testimony) (*domain.Testimony, error) {
	var person domain.Person
	if isExist, _ := r.SQLHandler.IsExist(person.TableName(), domain.Person{PersonID: testimony.PersonID}); !isExist {
		return nil, fmt.Errorf("Person not exist")
	}

	err := r.SQLHandler.Create(&testimony)
	if err != nil {
		return nil, err
	}
	return &testimony, nil
}

// GetTestimony return testimonies found by params
func (r *TestimonyRepository) GetTestimony(ctx context.Context, params domain.GetTestimonyParams) (*domain.Testimony, error) {
	var testimony *domain.Testimony
	err := r.SQLHandler.First(&testimony, params)

	if err != nil {
		return nil, err
	}

	return testimony, nil
}

// GetAllTestimony return testimonies found by params
func (r *TestimonyRepository) GetAllTestimony(ctx context.Context, pager domain.Pager, params domain.GetTestimonyParams) ([]*domain.Testimony, error) {
	var testimonies []*domain.Testimony
	err := r.SQLHandler.Find(
		&testimonies,
		pager,
		domain.Testimony{
			TestimonyID: *params.TestimonyID,
			PersonID:    *params.PersonID,
			Status:      *params.Status,
			CreatedAt:   *params.CreatedAt,
		})

	if err != nil {
		return nil, err
	}

	return testimonies, nil
}
