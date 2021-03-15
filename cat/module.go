package cat

import "api/cat/repository"

func NewRepository() repository.Repository {
	return repository.NewLocalRepository()
}

func NewLocalService() Service {
	repository := NewRepository()

	return NewService(repository)
}
