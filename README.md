# SimpleWeb

A very simple library for building serverside rendered web apps in Golang.
Example:

```go
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
    })
    simpleweb.Run()
}
```

All required artefacts like static or dynamic HTML pages are embedded. The
compiled application consist of one single binary without any external
dependency on the target machine.

## Basic Templating

The central rendering function `simpleweb.Render` expects a template name
together with an anonymous data struct:

```go
simpleweb.Render("templates/index.html", w, struct {
    Name string
}{Name: "SimpleWeb"})
```

The data attributes are referred inside the template via a dot prefix, e.g.:

```html
<h1>Hello, {{.Name}}</h1>
```

## Flash messages

Flash messages are one-time messages that are rendered by the handler that
generates the final HTML and are deleted afterward:

```go
f := func(w http.ResponseWriter, r *http.Request) {
    simpleweb.Info("info message")
    simpleweb.Render("templates/index.html", w, struct {
        Name string
    }{Name: "SimpleWeb"})
}
```

View helpers are provided for info, warning and error messages. See `hasInfo` and
`info` for showing info messages:

```html
{{if hasInfo}}
<p style="color: green">{{info}}</p>
{{end}}
```

## Limitations

Since SimpleWeb uses go:embed it is required that all templates are accessible
from the projects root directory. Thus, if your project root is `hello` all
templates must live in `hello` or in one of its subdirectories.