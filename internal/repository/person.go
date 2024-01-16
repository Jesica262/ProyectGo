package repository

import "app/internal"

type PersonRepository interface {
	GetAll() ([]internal.Person, error)
	GetByID(id int)(internal.Person, error)
	Save(person internal.Person)(int, error)
}