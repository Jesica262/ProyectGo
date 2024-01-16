package PersonRepository

import "app/internal"

type PersonRepositoryMap struct {
	persons 	map[int]internal.Person
	autoID		int
}

func NewPersonRepository() *PersonRepositoryMap {
	return &PersonRepositoryMap{
		persons: 	make(map[int]internal.Person),
		autoID: 	0,
	}
}

func (r *PersonRepositoryMap) GetAll() ([]internal.Person, error) {

	persons := make([]internal.Person, 0, len(r.persons))

	for _,p := range r.persons {
		persons = append(persons, p)
	}
	return persons, nil
}

func (r *PersonRepositoryMap) GetByID(id int) (person internal.Person,err error) {

	person, ok := r,persons[id]
	if !ok {
		err := errors.New("Person no encontrada")
	}
	return
}

func (r *PersonRepositoryMap) save(person internal.Person) (int error) {

	r.autoID++
	person.ID = r.autoID
	r.persons[r.autoID] = person
	return r.autoID, nil

}