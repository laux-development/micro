package micro

import (
	"html/template"
	"net/http"
)

// templates
// parse and cache templates
var templates = template.Must(template.ParseFiles(
	"templates/components/login/login-element.tmpl",
	"templates/components/login/login-cdn-library.tmpl",
	"templates/components/login/login-script.tmpl",
	"templates/components/login/login-return-script.tmpl",
	"pages/login.tmpl",
	"pages/home.tmpl",
))

func init() {
	http.HandleFunc("/", home)
	http.HandleFunc("/login", login)
}

func home(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "home", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "login", nil)
}
