package registry

import (
	"github.com/Code0716/graphql-study/interactor"
	"github.com/Code0716/graphql-study/repository"
)

// Getter gets registered instances.
type Getter interface {
	RepositoryGetter
	InteractorGetter
}

// RepositoryGetter gets registered repository instances.
type RepositoryGetter interface {
	PersonsRepository() repository.PersonsInterface
	TestimonyRepository() repository.TestimonyInterface
}

// InteractorGetter gets registered interactor instances.
type InteractorGetter interface {
	PersonsInteractor() *interactor.PersonsInteractor
	TestimonyInteractor() *interactor.TestimonyInteractor
}
