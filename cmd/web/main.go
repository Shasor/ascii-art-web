package main

import (
	"fmt"
	"net/http"

	"github.com/Shasor/ascii-art-web/internal/handlers"
)

func main() {
	const port = ":8080"

	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/submit", handlers.Submit)
	http.Handle("/src/assets/css/", http.StripPrefix("/src/assets/css/", http.FileServer(http.Dir("./src/assets/css"))))
	// http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates"))))
	// http.Handle("/templates/layouts/", http.StripPrefix("/templates/layouts/", http.FileServer(http.Dir("./templates/layouts"))))
	// http.Handle("/internal/handlers/", http.StripPrefix("/internal/handlers/", http.FileServer(http.Dir("./internal/handlers"))))

	fmt.Println("Click here to open the local web page http://localhost" + port)
	http.ListenAndServe(port, nil)
}
