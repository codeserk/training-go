package repository

import "api/cat/entity"

type Reader interface {
	Find(id string) (*entity.Cat, error)
	FindAll() ([]*entity.Cat, error)
}

type Writer interface {
	Create(name string, color string) (*entity.Cat, error)
}

type Repository interface {
	Reader
	Writer
}
