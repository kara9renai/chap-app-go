package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles((filepath.Join("templates", t.filename))))
	})
	t.templ.Execute(w, r)
}

func main() {
	const port = ":8080"
	http.Handle("/", &templateHandler{filename: "chat.html"})

	log.Println("Starting web server on", nil)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
