package interactor

import (
	"context"
	"time"

	"github.com/Code0716/graphql-study/domain"
	"github.com/Code0716/graphql-study/repository"
	uuid "github.com/satori/go.uuid"
)

// PersonsInteractor is person interactor.
type PersonsInteractor struct {
	PersonsRepository repository.PersonsInterface
}

// NewPersons initializes item interactor.
func NewPersons(
	personsRepo repository.PersonsInterface,
) *PersonsInteractor {
	return &PersonsInteractor{
		PersonsRepository: personsRepo,
	}
}

// GetAllPersons returns person list
// im: persons interactor
func (im *PersonsInteractor) GetAllPersons(ctx context.Context, params domain.Pager) ([]*domain.Person, error) {
	personList, err := im.PersonsRepository.GetAllPersons(ctx, params)
	if err != nil {
		return nil, err
	}

	return personList, nil
}

// GetPerson returns person
// im: persons interactor
func (im *PersonsInteractor) GetPerson(ctx context.Context, params domain.GetPersonParams) (*domain.Person, error) {
	person, err := im.PersonsRepository.GetPerson(ctx, params)
	if err != nil {
		return nil, err
	}

	return person, nil
}

// RegistPerson regist person
// im: persons interactor
func (im *PersonsInteractor) RegistPerson(ctx context.Context, params domain.Person) (*domain.Person, error) {
	currentTime := time.Now()
	params.CreatedAt = currentTime
	params.UpdatedAt = currentTime
	params.PersonID = uuid.NewV4().String()

	if params.Birthday == nil {
		params.Birthday = nil
	}

	person, err := im.PersonsRepository.CreatePerson(ctx, params)
	if err != nil {
		return nil, err
	}
	return person, nil
}
