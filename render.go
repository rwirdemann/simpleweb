package simpleweb

import (
	"embed"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

// MsgSuccess holds a success messages that is shown on the index page.
var MsgSuccess string

// MsgError holds an error message that is shown on the index page.
var MsgError string

type ViewData struct {
	Title   string
	Message string
	Error   string
}

var templatePatterns []string
var templates embed.FS
var router *mux.Router
var port int

func Init(fs embed.FS, patterns []string, p int) {
	templatePatterns = patterns
	templates = fs
	router = mux.NewRouter()
	port = p
}

func Register(path string, f http.HandlerFunc) {
	router.HandleFunc(path, f)
}

func ShowRoutes() {
	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		tpl, _ := route.GetPathTemplate()
		met, _ := route.GetMethods()
		log.Println(tpl, met)
		return nil
	})
}

// Render renders tmpl embedded in layout.html using the provided data.
func Render(tmpl string, w http.ResponseWriter, data any) {
	if err := RenderE(tmpl, w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// RenderE works the same as Render except returning the error instead of
// handling it.
func RenderE(tmpl string, w http.ResponseWriter, data any) error {
	tmpls := append(templatePatterns, tmpl)
	t, err := template.ParseFS(templates, tmpls...)
	if err != nil {
		return err
	}
	return t.Execute(w, data)
}

// RenderS renders tmpl embedded in layout.html and inserts title.
func RenderS(tmpl string, w http.ResponseWriter, title string) {
	if err := RenderE(tmpl, w, ViewData{
		Title:   title,
		Message: "",
		Error:   "",
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// RedirectE redirects to url after setting the global msgError to err.
func RedirectE(w http.ResponseWriter, r *http.Request, url string, err error) {
	MsgError = err.Error()
	http.Redirect(w, r, url, http.StatusSeeOther)
}

func ClearMessages() (string, string) {
	e := MsgError
	m := MsgSuccess
	MsgError = ""
	MsgSuccess = ""
	return m, e
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
