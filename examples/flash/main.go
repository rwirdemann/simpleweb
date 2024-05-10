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
	// Required Init call to tell SimpleWeb about its embedded templates, list
	// of base templates (empty) and port
	simpleweb.Init(templates, []string{}, 3030)
}

func main() {
	simpleweb.Register("/", func(w http.ResponseWriter, r *http.Request) {
		simpleweb.Render("templates/index.html", w, struct {
			Name string
		}{Name: "SimpleWeb"})
	}, "GET")

	simpleweb.Register("/info", func(w http.ResponseWriter, r *http.Request) {
		simpleweb.Info("info message")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}, "GET")

	simpleweb.ShowRoutes()
	simpleweb.Run()
}
