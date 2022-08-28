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

	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Server starting on port", portNumber)
	err = http.ListenAndServe(portNumber, nil)
	if err != nil {
		log.Fatal("Error serving the server!")
	}
}
