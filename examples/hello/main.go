package main

import (
	"embed"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rwirdemann/simpleweb"
	"log"
	"net/http"
)

//go:embed all:templates
var templates embed.FS

func main() {
	simpleweb.Init([]string{}, templates)

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
		simpleweb.Render("templates/index.html", w, struct {
		}{})
	})
	log.Printf("Listening on :%d...", 3030)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", 3030), router); err != nil {
		log.Fatal(err)
	}

}
