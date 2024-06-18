package handlers

import (
	"bytes"
	"net/http"
	"path/filepath"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index")
}

func renderTemplate(w http.ResponseWriter, tmplName string) {
	templateCache, err := createTemplateCache()
	if err != nil {
		panic(err)
	}

	tmpl, ok := templateCache[tmplName+".page.tmpl"]
	if !ok {
		http.Error(w, "The template doesn't exist.", http.StatusInternalServerError)
		return
	}

	buf := new(bytes.Buffer)
	tmpl.Execute(buf, nil)
	buf.WriteTo(w)
}

func createTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		tmpl := template.Must(template.ParseFiles(page))

		layouts, err := filepath.Glob("./templates/layouts/*.layout.tmpl")
		if err != nil {
			return cache, err
		}

		if len(layouts) > 0 {
			tmpl.ParseGlob("./templates/layouts/*.layout.tmpl")
		}

		cache[name] = tmpl
	}

	return cache, nil
}
