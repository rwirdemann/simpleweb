# SimpleWeb

A very simple library for building serverside rendered web apps in Golang.
Example:

```go
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

## Limitations

Since SimpleWeb uses go:embed it is required that all templates are accessible
from the projects root directory. Thus, if your project root is `hello` all
templates must live in `hello` or in one of its subdirectories.