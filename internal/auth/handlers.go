package auth

import (
	"net/http"
	"text/template"
)

type DataPage struct {
	Email    string
	Password string
}

func HomePage(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "400 Method Not Allowed", http.StatusBadRequest)
		return
	}

	temp, err := template.ParseFiles("templates/home.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	temp.Execute(w, nil)
}

func Register(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/register" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "400 Method Not Allowed", http.StatusBadRequest)
		return
	}

	temp, err := template.ParseFiles("templates/dashboard.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	email := r.FormValue("email")
	password := r.FormValue("password")

	data := DataPage{
		Email:    email,
		Password: password,
	}
	temp.Execute(w, data)
}
