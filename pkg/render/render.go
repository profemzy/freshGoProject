package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templateMap = make(map[string]*template.Template)

func ProcessTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check to see if template already exist in our cache
	_, inMap := templateMap[t]

	if !inMap {
		// need to create the template
		err := createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	}

	tmpl = templateMap[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	// parse the templates
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		log.Println(err)
	}

	// add template to cache(map)
	templateMap[t] = tmpl

	return nil
}
