package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templateCache = make(map[string]*template.Template)

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, t string) {
	var template *template.Template
	var err error

	// check if template already in cache - comma ok idiom
	_, ok := templateCache[t]
	if !ok {
		// need to create a template - read from disk
		log.Println("creating template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		// use cached templates - read from memory
		log.Println("using template from cache")
	}
	template = templateCache[t]
	err = template.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

// createTemplateCache creates an in memory caching layer for the templates
func createTemplateCache(t string) error {
	// templates from the templates folder - passed in the handlers
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.gohtml",
	}
	// parse templates from disk
	template, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	// add template to cache - map
	templateCache[t] = template
	return nil
}
