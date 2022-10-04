package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/blessedmadukoma/modern-go-apps/project/pkg/config"
	"github.com/blessedmadukoma/modern-go-apps/project/pkg/handlers"
	"github.com/blessedmadukoma/modern-go-apps/project/pkg/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	// change to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

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
		Addr:    portNumber,
		Handler: routes(&app),
	}

	fmt.Println("Server starting on port", portNumber)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("Error serving the server!")
	}
}
