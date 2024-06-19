package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/Shasor/ascii-art-web/internal/ascii-art"
)

func Index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", "")
}

func Submit(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()

    text := r.Form.Get("text")
    font := r.Form.Get("font")

    // Traiter les données du formulaire (validation etc.)
    fmt.Println("Text:", string('"') + text + string('"'))
    fmt.Println("Font:", font)

    // En cas de succès, définir le message de confirmation
    data := ascii.Ascii(text, font)

    // Rendre le template "index" avec les données
    renderTemplate(w, "index", data)
}

func renderTemplate(w http.ResponseWriter, tmplName string, data string) {
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
	tmpl.Execute(buf, data)
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
