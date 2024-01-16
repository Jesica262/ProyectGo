package handler

import (
	"encoding/json"
	"net/http"
	"app/internal"
	"app/internal/service"
	"app/platform/web/response"
)

type PersonHandler struct {
	
	personService service.PersonService
}

func NewPersonHandler( personService service.PersonService) *PersonHandler {

	return &PersonHandler {
		personService: personService
	}
}

func (h *PersonHandler) GetAllPeopleHandler(w http.ResponseWriter, r *http.Request) {
	
	people, err := h.personService.GetAll()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Error al obtener personas")
		return
	}
	response.JSON(w, http.StatusOK, people)
}

func (h *PersonHandler) GetPersonBYIdHandler(w http.ResponseWriter, r *http.Request) {

	id := 1

	person, err := h.personService.GetByID(id)
	if err != nil {
		response.Error(w, http.StatusNotFound, "Persona no encontrada")
		return
	}
	response.JSON(w, http.StatusOK, person)
}

func (h *PersonHandler) SavePersonHandler(w http.ResponseWriter, r *http.Request) { 

	var person internal.Person

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decoder(&person); err != nil {
		response.Error(w, http.StatusBadRequest, "Error al decodificar la solicitud")
		return
	}

	id, err := h.personService.Save(person)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Error al guardar la persona")
		return
	}

	response.JSON(w, http.StatusCreated, map[string]int{"id": id})

}