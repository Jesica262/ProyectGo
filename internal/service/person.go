package service

import "app/internal"

type PersonService interface {
	GetAll() ([]internal.Person, error)
	GetByID(id int) (internal.Person, error)
	Save(person internal.Person) (int, error)
}

