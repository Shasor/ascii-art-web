package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Shasor/ascii-art-web/internal/ascii-art"
)

func Index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", "")
}

func Submit(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()

    text := r.Form.Get("text")
	text = strings.ReplaceAll(text, "\r\n", "\\n")

    font := r.Form.Get("font")

    // Process form data (validation, etc.)
    fmt.Println("Text:", text)
    fmt.Println("Font:", font)

    // Generate ASCII art on success
    data := ascii.Ascii(text, font)

    // Render the "index" template with data
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
