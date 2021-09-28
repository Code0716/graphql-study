package interactor

import (
	"context"
	"time"

	"github.com/Code0716/graphql-study/domain"
	"github.com/Code0716/graphql-study/repository"
	uuid "github.com/satori/go.uuid"
)

// TestimonyInteractor is testimony interactor.
type TestimonyInteractor struct {
	TestimonyRepository repository.TestimonyInterface
}

// NewTestimony initializes item interactor.
func NewTestimony(
	testimonysRepo repository.TestimonyInterface,
) *TestimonyInteractor {
	return &TestimonyInteractor{
		TestimonyRepository: testimonysRepo,
	}
}

// GetAllTestimony returns testimony list
// it: testimony interactor
func (it *TestimonyInteractor) GetAllTestimony(ctx context.Context, pager domain.Pager, params domain.GetTestimonyParams) ([]*domain.Testimony, error) {
	testimonyList, err := it.TestimonyRepository.GetAllTestimony(ctx, pager, params)
	if err != nil {
		return nil, err
	}

	return testimonyList, nil
}

// GetTestimony returns testimony
// it: testimonys interactor
func (it *TestimonyInteractor) GetTestimony(ctx context.Context, params domain.GetTestimonyParams) (*domain.Testimony, error) {
	testimony, err := it.TestimonyRepository.GetTestimony(ctx, params)
	if err != nil {
		return nil, err
	}

	return testimony, nil
}

// RegistTestimony regist testimony
// it: testimonys interactor
func (it *TestimonyInteractor) RegistTestimony(ctx context.Context, params domain.Testimony) (*domain.Testimony, error) {

	currentTime := time.Now()
	params.CreatedAt = currentTime
	params.UpdatedAt = currentTime
	params.TestimonyID = uuid.NewV4().String()

	testimony, err := it.TestimonyRepository.CreateTestimony(ctx, params)
	if err != nil {
		return nil, err
	}

	return testimony, nil
}
