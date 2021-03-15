package cat

import (
	"api/cat/entity"
	"api/cat/repository"
)

type Service interface {
	Find(id string) (*entity.Cat, error)
	FindAll() ([]*entity.Cat, error)

	Create(name string, color string) (*entity.Cat, error)
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) Service {
	return &service{
		repository: repository,
	}
}

func (service *service) Find(id string) (*entity.Cat, error) {
	return service.repository.Find(id)
}

func (s *service) FindAll() ([]*entity.Cat, error) {
	return s.repository.FindAll()
}

func (service *service) Create(name string, color string) (*entity.Cat, error) {
	return service.repository.Create(name, color)
}
