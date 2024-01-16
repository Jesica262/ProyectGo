package main

import (
	"net/http"
	"app/internal/application"
	"app/internal/handler"
	"app/internal/repository"
	"app/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	// Crear instancias
	personRepository := repository.NewMemoryPersonRepository()
	personService := service.NewPersonService(personRepository)
	personHandler := handler.NewPersonHandler(personService)

	// Configurar rutas
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/people", personHandler.GetAllPeopleHandler)
	router.Get("/people/{id}", personHandler.GetPersonByIDHandler)
	router.Post("/people", personHandler.SavePersonHandler)

	// Crear aplicación
	app := application.NewApplication(router)

	// Configurar y ejecutar la aplicación
	if err := app.Run(); err != nil {
		panic(err)
	}
}