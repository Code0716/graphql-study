package repository

import (
	"context"

	"github.com/Code0716/graphql-study/domain"
)

// MembersInterface  is data access methods to Members.
type MembersInterface interface {
	CreateMember(ctx context.Context, member domain.Member) (*domain.Member, error)
	GetAllMembers(ctx context.Context, params domain.Pager) ([]*domain.Member, error)
	GetMember(ctx context.Context, params domain.GetMemberParams) (*domain.Member, error)
}

// MembersRepository is member repository.
type MembersRepository struct {
	SQLHandler SQLHandlerInterface
}

// NewMembers initializes members repository.
func NewMembers(sqlHandler SQLHandlerInterface) *MembersRepository {
	return &MembersRepository{
		sqlHandler,
	}
}

// CreateMember regist member to members db
func (r *MembersRepository) CreateMember(ctx context.Context, member domain.Member) (*domain.Member, error) {
	err := r.SQLHandler.Create(&member)
	if err != nil {
		return nil, err
	}
	return &member, nil
}

// GetMember return members found by params
func (r *MembersRepository) GetMember(ctx context.Context, params domain.GetMemberParams) (*domain.Member, error) {
	var member *domain.Member
	err := r.SQLHandler.First(&member, params)

	if err != nil {
		return nil, err
	}

	return member, nil
}

// GetAllMembers return members found by params
func (r *MembersRepository) GetAllMembers(ctx context.Context, params domain.Pager) ([]*domain.Member, error) {
	var members []*domain.Member
	err := r.SQLHandler.Find(&members, params)

	if err != nil {
		return nil, err
	}

	return members, nil
}
