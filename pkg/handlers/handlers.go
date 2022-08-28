package handlers

import (
	"net/http"

	"github.com/blessedmadukoma/modern-go-apps/project/pkg/render"
)

// Home is the  about handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

// About is the  about handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}
