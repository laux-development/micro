package micro

import (
	"fmt"
	"html/template"
	"net/http"
)

// templates
// parse and cache templates
var templates = template.Must(template.ParseFiles(
	"templates/components/login-element.tmpl",
	"templates/components/login-cdn-library.tmpl",
	"templates/components/login-script.tmpl",
	"templates/pages/login.tmpl",
))

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}
