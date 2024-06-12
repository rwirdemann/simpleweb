package main

import (
	"embed"
	"github.com/rwirdemann/simpleweb/pkg/simpleweb"
	"net/http"
)

// Expects all HTML templates in $PROJECTROOT/templates
//
//go:embed all:templates
var templates embed.FS

//go:embed assets
var assets embed.FS

func init() {
	// Required Init call to tell SimpleWeb about its embedded templates, list
	// of base templates (empty) and port
	simpleweb.Init(templates, []string{}, 3030)
}

func main() {
	simpleweb.Static("/assets", assets)
	simpleweb.Register("/", func(w http.ResponseWriter, r *http.Request) {
		simpleweb.Render("templates/index.html", w, struct {
			Name string
		}{Name: "SimpleWeb"})
	}, "GET")
	simpleweb.ShowRoutes()
	simpleweb.Run()
}
