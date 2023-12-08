package controllers

import (
	"fmt"
	"net/http"

	"github.com/gerry-sheva/lenslocked/models"
)

type Users struct {
	Templates struct {
		New Template
	}
	UserService *models.UserService
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}

	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, data)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	email := r.FormValue("email")
	passwordRaw := r.FormValue("password")

	user, err := u.UserService.Create(email, passwordRaw)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "User created: %+v", user)
}
