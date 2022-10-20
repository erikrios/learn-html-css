package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
)

var tmpl *template.Template

func init() {
	initTemplate()
}

func main() {
	port := 3000

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./statics"))
	mux.Handle("/statics/", http.StripPrefix("/statics/", fs))
	mux.HandleFunc("/", helloHandler)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	log.Printf("Server listening on port %d\n", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	initTemplate()
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func initTemplate() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	htmlFiles := path.Join(wd, "assets", "html", "*.html")

	tmpl = template.Must(template.ParseGlob(htmlFiles))
}
