package simpleweb

import (
	"embed"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"path"
)

// baseTemplates contains the list of base templates that are rendered with
// every request.
var baseTemplates []string

// htmlTemplates holds the embedded HTML templates
var htmlTemplates embed.FS

var router *mux.Router

// port of the web server
var port = -1

// Init initializes the webapp by providing a collection of embedded html
// templates fs, the names of baseTemplates that are rendered with every request
// (e.g. layout.html) and the port p of the web server.
func Init(ht embed.FS, bt []string, p int) {
	baseTemplates = bt
	htmlTemplates = ht
	router = mux.NewRouter()

	port = p
}

// Static registers path as root directory for serving static content like css, js or images.
func Static(path string, fs embed.FS) {
	router.PathPrefix(path).Handler(http.FileServer(http.FS(fs)))
}

func Register(path string, f http.HandlerFunc, methods ...string) {
	if router == nil {
		log.Fatal("simpleweb.Init needs to be called first")
	}
	router.HandleFunc(path, f).Methods(methods...)
}

func ShowRoutes() {
	if router == nil {
		log.Fatal("simpleweb.Init needs to be called first")
	}
	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		tpl, _ := route.GetPathTemplate()
		met, _ := route.GetMethods()
		log.Println(tpl, met)
		return nil
	})
}

// RenderE renders tmpl using the provided data and returns. Returns any errors
// that have occurred.
func RenderE(tmpl string, w http.ResponseWriter, data any) error {
	files := append(baseTemplates, tmpl)
	t, err := template.New(path.Base(files[0])).Funcs(flashHelper).ParseFS(htmlTemplates, files...)
	if err != nil {
		return err
	}
	return t.Execute(w, data)
}

// Render renders tmpl using the provided data.
func Render(tmpl string, w http.ResponseWriter, data any) {
	if err := RenderE(tmpl, w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// RenderPartialE renders the partial without any surrounding templates.
func RenderPartialE(partial string, w http.ResponseWriter, data any) error {
	t, err := template.ParseFS(htmlTemplates, partial)
	if err != nil {
		return err
	}
	return t.Execute(w, data)
}

// RedirectE redirects to url after setting Flash.error to err.
func RedirectE(w http.ResponseWriter, r *http.Request, url string, err error) {
	Error(err.Error())
	http.Redirect(w, r, url, http.StatusSeeOther)
}

// Redirect redirects to url.
func Redirect(w http.ResponseWriter, r *http.Request, url string) {
	http.Redirect(w, r, url, http.StatusSeeOther)
}

// FormValue returns the value of key for POST and PUT requests.
func FormValue(r *http.Request, key string) (string, error) {
	if err := r.ParseForm(); err != nil {
		return "", err
	}
	return r.FormValue(key), nil
}

// Run starts the web server on the given port.
func Run() {
	if port == -1 {
		log.Fatal("simpleweb.Init needs to be called first")
	}
	log.Printf("Listening on :%d...", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), router); err != nil {
		log.Fatal(err)
	}
}
