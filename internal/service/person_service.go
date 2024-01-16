package service

import (
	"app/internal"
	"app/internal/repository"
)

type PersonService struct {

	personaRepository repository.PersonRepository
}

func NewPersonService (personaRepository repository.PersonRepository) *PersonService {

	return &PersonService {
		personaRepository: personaRepository,
	}
}

func (s *PersonService) GetAll() ([]internal.Person, error) {

	return s.personaRepository.GetAll()
}

func (s *PersonService) GetByID(id int) (internal.Person, error) {

	return s.personaRepository.GetByID(id)
}

func (s *PersonService) Save(person internal.Person) (int, error) {

	return s.personaRepository.Save(person)
}