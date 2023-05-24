package main

import (
	"path/filepath"
	"snippetbox/internal/models"
	"text/template"
	"time"
)

type templateData struct {
	CurrentYear int
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
	Form        any
	Flash       string
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 10:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	// the map is the cache
	cache := map[string]*template.Template{}

	// get slice of all template filepaths
	pages, err := filepath.Glob("./ui/html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {

		// get the filename from the full path
		name := filepath.Base(page)

		// Parse base temp into ts
		ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.tmpl.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl.html")
		if err != nil {
			return nil, err
		}

		// parse the templates into a template set
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// add the template set using the name of the page
		// as the key
		cache[name] = ts
	}

	return cache, nil
}
