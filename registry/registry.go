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

// MembersRepository returns members repository.
func (r *Registry) MembersRepository() repository.MembersInterface {
	return repository.NewMembers(r.db)
}

// MembersInteractor returns members interactor.
func (r *Registry) MembersInteractor() *interactor.MembersInteractor {

	return interactor.NewMembers(r.MembersRepository())
}
