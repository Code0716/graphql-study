package registry

import (
	"github.com/Code0716/graphql-study/interactor"
	"github.com/Code0716/graphql-study/repository"
)

// Registry returns initialized repositories and interactores.
type Registry struct {
	db repository.SQLHandlerInterface
}

// New initializes registry with gorm-repository.
func New(db repository.SQLHandlerInterface) *Registry {
	return &Registry{
		db: db,
	}
}

/*
	以下に具体的な依存性を解決する初期化処理を書く
*/

// PersonsRepository returns persons repository.
func (r *Registry) PersonsRepository() repository.PersonsInterface {
	return repository.NewPersons(r.db)
}

// PersonsInteractor returns persons interactor.
func (r *Registry) PersonsInteractor() *interactor.PersonsInteractor {

	return interactor.NewPersons(r.PersonsRepository())
}
