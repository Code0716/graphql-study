package interactor

import (
	"context"
	"time"

	"github.com/Code0716/graphql-study/domain"
	"github.com/Code0716/graphql-study/repository"
	uuid "github.com/satori/go.uuid"
)

// MembersInteractor is member interactor.
type MembersInteractor struct {
	MembersRepository repository.MembersInterface
}

// NewMembers initializes item interactor.
func NewMembers(
	membersRepo repository.MembersInterface,
) *MembersInteractor {
	return &MembersInteractor{
		MembersRepository: membersRepo,
	}
}

// GetAllMembers returns member list
// im: members interactor
func (im *MembersInteractor) GetAllMembers(ctx context.Context, params domain.Pager) ([]*domain.Member, error) {
	memberList, err := im.MembersRepository.GetAllMembers(ctx, params)
	if err != nil {
		return nil, err
	}

	return memberList, nil
}

// GetMember returns member
// im: members interactor
func (im *MembersInteractor) GetMember(ctx context.Context, params domain.GetMemberParams) (*domain.Member, error) {
	member, err := im.MembersRepository.GetMember(ctx, params)
	if err != nil {
		return nil, err
	}

	return member, nil
}

// RegistMember regist member
// im: members interactor
func (im *MembersInteractor) RegistMember(ctx context.Context, params domain.Member) (*domain.Member, error) {
	currentTime := time.Now()
	params.CreatedAt = currentTime
	params.UpdatedAt = currentTime
	params.MemberID = uuid.NewV4().String()

	if params.Birthday == nil {
		params.Birthday = nil
	}

	if params.Status == "" {
		params.Status = domain.StatusMap[domain.StatusCodeOther]
	}

	member, err := im.MembersRepository.CreateMember(ctx, params)
	if err != nil {
		return nil, err
	}
	return member, nil
}
