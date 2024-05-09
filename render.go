package simpleweb

import (
	"embed"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

type Flash struct {
	info  string
	error string
}

func (f Flash) Info() string {
	return f.info
}

func (f Flash) Error() string {
	return f.error
}

// baseTemplates contains the list of base templates that are rendered with
// every request.
var baseTemplates []string

// htmlTemplates holds the embedded HTML templates
var htmlTemplates embed.FS

var router *mux.Router

// port of the web server
var port int

var flash = Flash{}

// Init initializes the webapp by providing a collection of embedded html
// templates fs, the names of baseTemplates that are rendered with every request
// (e.g. layout.html) and the port p of the web server.
func Init(ht embed.FS, bt []string, p int) {
	baseTemplates = bt
	htmlTemplates = ht
	router = mux.NewRouter()
	port = p
}

func Register(path string, f http.HandlerFunc, methods ...string) {
	router.HandleFunc(path, f).Methods(methods...)
}

func ShowRoutes() {
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
	tmpls := append(baseTemplates, tmpl)
	t, err := template.ParseFS(htmlTemplates, tmpls...)
	if err != nil {
		return err
	}
	d := struct {
		Flash Flash
		Data  any
	}{
		Flash: flash,
		Data:  data,
	}
	flash = Flash{}
	return t.Execute(w, d)
}

// Render renders tmpl using the provided data.
func Render(tmpl string, w http.ResponseWriter, data any) {
	if err := RenderE(tmpl, w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RenderPartialE(partial string, w http.ResponseWriter, data any) error {
	t, err := template.ParseFS(htmlTemplates, partial)
	if err != nil {
		return err
	}
	d := struct {
		Flash Flash
		Data  any
	}{
		Flash: flash,
		Data:  data,
	}
	flash = Flash{}
	return t.Execute(w, d)
}

// RedirectE redirects to url after setting Flash.error to err.
func RedirectE(w http.ResponseWriter, r *http.Request, url string, err error) {
	Error(err.Error())
	http.Redirect(w, r, url, http.StatusSeeOther)
}

// FormValue returns the value of key for POST and PUT requests.
func FormValue(r *http.Request, key string) (string, error) {
	if err := r.ParseForm(); err != nil {
		return "", err
	}
	return r.FormValue(key), nil
}

func Run() {
	log.Printf("Listening on :%d...", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), router); err != nil {
		log.Fatal(err)
	}
}

func Info(s string) {
	flash.info = s
}

func Error(s string) {
	flash.error = s
}
