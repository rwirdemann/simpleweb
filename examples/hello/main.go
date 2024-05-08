package main

import (
	"embed"
	"github.com/rwirdemann/simpleweb"
	"net/http"
)

// Expects all HTML templates in $PROJECTROOT/templates
//
//go:embed all:templates
var templates embed.FS

func init() {
	simpleweb.Init(templates, []string{}, 3030)
}

func main() {
	simpleweb.Register("/", func(w http.ResponseWriter, r *http.Request) {
		simpleweb.Render("templates/index.html", w, struct {
			Name string
		}{Name: "SimpleWeb"})
	})
	simpleweb.Run()
}
