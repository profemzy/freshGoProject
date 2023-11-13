package handlers

import (
	"github.com/profemzy/freshGoProject/pkg/render"
	"net/http"
)

// Home is the handler for the home page
func Home(w http.ResponseWriter, _ *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the handler for the about page
func About(w http.ResponseWriter, _ *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
