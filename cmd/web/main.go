package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/blessedmadukoma/modern-go-apps/project/pkg/config"
	"github.com/blessedmadukoma/modern-go-apps/project/pkg/handlers"
	"github.com/blessedmadukoma/modern-go-apps/project/pkg/render"
)

const portNumber = ":8080"

func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	fmt.Println("Server starting on port", portNumber)
	
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("Error serving the server!")
	}
}
