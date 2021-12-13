package controllers

import "net/http"

//LoadLoginPage loads the login page
func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login Page"))

}
