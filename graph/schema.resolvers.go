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

func (r *mutationResolver) CreatePerson(ctx context.Context, input model.CreatePerson) (*model.CommonResponse, error) {
	personsInteractor := r.Reg.PersonsInteractor()

	newPerson := domain.Person{
		Name:        input.Name,
		Address:     *input.Address,
		PhoneNumber: *input.PhoneNumber,
		ClassName:   *input.ClassName,
	}
	m, err := personsInteractor.RegistPerson(ctx, newPerson)

	if err != nil {
		return nil, err
	}
	response := model.CommonResponse{Message: fmt.Sprintf(`Create Success name %s ID %s`, m.Name, m.PersonID)}
	return &response, nil
}

func (r *mutationResolver) UpdatePerson(ctx context.Context, input model.CreatePerson) (*model.Person, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeletePerson(ctx context.Context, personID string) (*model.Person, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateClassification(ctx context.Context, input model.CreateClassification) (*model.Classification, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteClassification(ctx context.Context, classID string) (*model.CommonResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateTestimony(ctx context.Context, input model.CreateTestimony) (*model.CommonResponse, error) {
	testimonyInteractor := r.Reg.TestimonyInteractor()

	if isValid := util.IsValidUUID(input.PersonID); !isValid {
		return nil, fmt.Errorf("invald parson id")
	}

	status := domain.GetGetTestimonyStatus(input.Status)
	params := domain.Testimony{
		PersonID:  input.PersonID,
		Testimony: input.Testimony,
		Status:    status,
	}
	testimony, err := testimonyInteractor.RegistTestimony(ctx, params)
	if err != nil {
		return nil, err
	}

	result := model.CommonResponse{
		Message: fmt.Sprintf("created %s", testimony.TestimonyID),
	}
	return &result, nil
}

func (r *mutationResolver) UpdateTestimony(ctx context.Context, input model.UpdateTestimony) (*model.Testimony, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteTestimony(ctx context.Context, testimonyID string) (*model.CommonResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetPersons(ctx context.Context, input *model.PersonPager) ([]*model.Person, error) {
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

	params := domain.PersonPager{
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

func (r *queryResolver) GetPerson(ctx context.Context, input model.GetPersonParams) (*model.Person, error) {
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

func (r *queryResolver) Classifications(ctx context.Context, input model.Pager) ([]*model.Classification, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetTestimonies(ctx context.Context, params *model.GetTestimonyParams, pager model.Pager) ([]*model.Testimony, error) {
	var limit int
	if pager.Limit != nil {
		limit = *pager.Limit
	}

	if limit <= 0 || 50 < limit {
		limit = 50
	}

	var offset int
	if pager.Offset != nil {
		offset = *pager.Offset
	}

	newPager := domain.Pager{
		Limit:  &limit,
		Offset: &offset,
	}

	testimonyInteractor := r.Reg.TestimonyInteractor()

	var personID string
	if params.PersonID != nil {
		personID = *params.PersonID
	}

	var status string
	if params.Status != nil {
		status = domain.GetGetTestimonyStatus(*params.Status)
	}

	var testimonyID string
	if params.TestimonyID != nil {
		testimonyID = *params.TestimonyID
	}

	var createdAt time.Time
	if params.CreatedAt != nil {
		createdAt = *params.CreatedAt

	}

	newParams := domain.GetTestimonyParams{
		TestimonyID: &testimonyID,
		PersonID:    &personID,
		Status:      &status,
		CreatedAt:   &createdAt,
	}

	result, err := testimonyInteractor.GetAllTestimony(ctx, newPager, newParams)
	if err != nil {
		return nil, err
	}

	testimonyList := make([]*model.Testimony, 0, len(result))

	for _, item := range result {
		person := model.Testimony{
			ID:          item.ID,
			PersonID:    item.PersonID,
			TestimonyID: item.TestimonyID,
			Testimony:   item.Testimony,
			Status:      item.Status,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
			DeletedAt:   item.DeletedAt,
		}
		testimonyList = append(testimonyList, &person)
	}

	return testimonyList, nil
}

func (r *queryResolver) GetTestimony(ctx context.Context, pager model.Pager) (*model.Testimony, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
