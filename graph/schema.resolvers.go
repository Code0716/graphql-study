package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Code0716/graphql-study/domain"
	"github.com/Code0716/graphql-study/graph/generated"
	"github.com/Code0716/graphql-study/graph/model"
)

func (r *mutationResolver) CreateMember(ctx context.Context, input model.CreateMember) (*model.Member, error) {
	membersInteractor := r.Reg.MembersInteractor()

	newMember := domain.Member{
		Name:        input.Name,
		Address:     *input.Address,
		PhoneNumber: *input.PhoneNumber,
		Status:      *input.Status,
		Birthday:    input.Birthday,
	}
	m, err := membersInteractor.RegistMember(ctx, newMember)

	if err != nil {
		return nil, err
	}
	result := model.Member{
		ID:          m.ID,
		MemberID:    m.MemberID,
		Name:        m.Name,
		Address:     m.Address,
		PhoneNumber: m.PhoneNumber,
		Status:      m.Status,
		Birthday:    m.Birthday,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
		DeletedAt:   m.DeletedAt,
	}
	return &result, nil
}

func (r *queryResolver) Members(ctx context.Context, input *model.Pager) ([]*model.Member, error) {
	var limit int
	if input.Limit != nil {
		limit = int(*input.Limit)
	}

	if limit <= 0 || 50 < limit {
		limit = 50
	}

	var offset int
	if input.Offset != nil {
		offset = int(*input.Offset)
	}

	var status string
	if input.Status != nil {
		status = string(*input.Status)
	}

	params := domain.Pager{
		Limit:  &limit,
		Offset: &offset,
		Status: &status,
	}
	membersInteractor := r.Reg.MembersInteractor()
	result, err := membersInteractor.GetAllMembers(ctx, params)

	membersList := make([]*model.Member, 0, len(result))
	if err != nil {
		return nil, err
	}

	for _, item := range result {
		member := model.Member{
			ID:          item.ID,
			MemberID:    item.MemberID,
			Name:        item.Name,
			Address:     item.Address,
			PhoneNumber: item.PhoneNumber,
			Status:      item.Status,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
			DeletedAt:   item.DeletedAt,
		}
		membersList = append(membersList, &member)
	}

	return membersList, nil
}

func (r *queryResolver) Member(ctx context.Context, input model.GetMemberParams) (*model.Member, error) {
	membersInteractor := r.Reg.MembersInteractor()
	params := domain.GetMemberParams{
		ID:   input.ID,
		Name: input.Name,
	}

	m, err := membersInteractor.GetMember(ctx, params)
	if err != nil {
		return nil, err
	}
	result := model.Member{
		ID:          m.ID,
		MemberID:    m.MemberID,
		Name:        m.Name,
		Address:     m.Address,
		PhoneNumber: m.PhoneNumber,
		Status:      m.Status,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
		DeletedAt:   m.DeletedAt,
	}

	return &result, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
