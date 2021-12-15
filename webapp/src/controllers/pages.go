package controllers

import (
	"net/http"
	"webapp/src/utils"
)

//LoadLoginPage loads the login page
func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}

//LoadSignUpPage loads the signup page
func LoadSignUpPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "signup.html", nil)
}
