package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func ProcessTemplate(w http.ResponseWriter, tmpl string) {

	// Create a template cache
	templateCache, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// Get requested template from cache
	requestedTemplate, found := templateCache[tmpl]

	if !found {
		log.Fatal(err)
	}

	// To help check for error in the map
	buffer := new(bytes.Buffer)

	err = requestedTemplate.Execute(buffer, nil)
	if err != nil {
		log.Println(err)
	}

	// Render the template
	_, err = buffer.WriteTo(w)
	if err != nil {
		fmt.Println("error rendering template:", err)
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// Get all the files name *.page.tmpl from .templates directory
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	// Loop through all files ending in *.page.tmpl
	for _, page := range pages {
		// get name of the page without the full file path
		name := filepath.Base(page)
		templateSet, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		layouts, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(layouts) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = templateSet
	}

	return myCache, nil
}
