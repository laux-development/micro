package micro

import (
	"net/http"
)

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
