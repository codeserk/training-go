package repository

import (
	"api/cat/entity"

	"github.com/google/uuid"
)

type LocalRepository struct {
	cats map[string]entity.Cat
}

func NewLocalRepository() Repository {
	return &LocalRepository{
		cats: map[string]entity.Cat{},
	}
}

func (r *LocalRepository) Find(id string) (*entity.Cat, error) {
	return nil, nil
}

func (r *LocalRepository) Create(name string, color string) (*entity.Cat, error) {
	id := uuid.New().String()
	cat := entity.Cat{
		ID:    id,
		Name:  name,
		Color: color,
	}
	r.cats[id] = cat

	return &cat, nil
}

func (r *LocalRepository) FindAll() ([]*entity.Cat, error) {
	cats := make([]*entity.Cat, len(r.cats))
	index := 0
	for _, value := range r.cats {
		cats[index] = &value
		index++
	}

	return cats, nil
}
