package application

import (
	"net/http"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Application define la aplicación web
type Application struct {
	server *http.Server
}

// NewApplication crea una nueva instancia de la aplicación web
func NewApplication(router *chi.Mux) *Application {
	addr := ":8080" // Puedes cambiar el puerto según tus preferencias
	server := NewServer(router, addr)
	return &Application{
		server: server,
	}
}

// Run inicia el servidor web
func (app *Application) Run() error {
	fmt.Printf("Servidor escuchando en %s\n", app.server.Addr)
	return app.server.ListenAndServe()
}