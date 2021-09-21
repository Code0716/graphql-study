package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/Code0716/graphql-study/domain"
	"github.com/Code0716/graphql-study/graph/generated"
	"github.com/Code0716/graphql-study/graph/model"
	"github.com/Code0716/graphql-study/util"
)

func (r *mutationResolver) CreatePerson(ctx context.Context, input model.CreatePerson) (*model.Person, error) {
	personsInteractor := r.Reg.PersonsInteractor()

	birthday := new(time.Time)
	if *input.Birthday != "" && input.Birthday != nil {
		*birthday = util.TimeFromStr(*input.Birthday)
	} else {
		birthday = nil
	}

	newPerson := domain.Person{
		Name:        input.Name,
		Address:     *input.Address,
		PhoneNumber: *input.PhoneNumber,
		ClassName:   *input.ClassName,
		Birthday:    birthday,
	}
	m, err := personsInteractor.RegistPerson(ctx, newPerson)

	if err != nil {
		return nil, err
	}
	result := model.Person{
		ID:          m.ID,
		PersonID:    m.PersonID,
		Name:        m.Name,
		Address:     m.Address,
		PhoneNumber: m.PhoneNumber,
		ClassName:   m.ClassName,
		Birthday:    m.Birthday,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
		DeletedAt:   m.DeletedAt,
	}

	return &result, nil
}

func (r *mutationResolver) CreateClassification(ctx context.Context, input model.CreateClassification) (*model.CommonSuccessResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Persons(ctx context.Context, input *model.Pager) ([]*model.Person, error) {
	var limit int
	if input.Limit != nil {
		limit = *input.Limit
	}

	if limit <= 0 || 50 < limit {
		limit = 50
	}

	var offset int
	if input.Offset != nil {
		offset = *input.Offset
	}

	var className string
	if input.ClassName != nil {
		className = *input.ClassName
	}

	params := domain.Pager{
		Limit:     &limit,
		Offset:    &offset,
		ClassName: &className,
	}
	personsInteractor := r.Reg.PersonsInteractor()
	result, err := personsInteractor.GetAllPersons(ctx, params)

	personsList := make([]*model.Person, 0, len(result))
	if err != nil {
		return nil, err
	}

	for _, item := range result {
		person := model.Person{
			ID:          item.ID,
			PersonID:    item.PersonID,
			Name:        item.Name,
			Address:     item.Address,
			PhoneNumber: item.PhoneNumber,
			ClassName:   item.ClassName,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
			DeletedAt:   item.DeletedAt,
		}
		personsList = append(personsList, &person)
	}

	return personsList, nil
}

func (r *queryResolver) Person(ctx context.Context, input model.GetPersonParams) (*model.Person, error) {
	personsInteractor := r.Reg.PersonsInteractor()
	params := domain.GetPersonParams{
		ID:   input.ID,
		Name: input.Name,
	}

	m, err := personsInteractor.GetPerson(ctx, params)
	if err != nil {
		return nil, err
	}
	result := model.Person{
		ID:          m.ID,
		PersonID:    m.PersonID,
		Name:        m.Name,
		Address:     m.Address,
		PhoneNumber: m.PhoneNumber,
		ClassName:   m.ClassName,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
		DeletedAt:   m.DeletedAt,
	}

	return &result, nil
}

func (r *queryResolver) Classifications(ctx context.Context, input *model.Pager) ([]*model.Classification, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
