package middlewares

import (
	"api/src/authentication"
	"api/src/responses"
	"fmt"
	"log"
	"net/http"
)

//Logger logs the requests in the terminal
func Logger(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunction(w, r)
	}
}

// Authenticate checks if user is authenticated
func Authenticate(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("validating")
		if error := authentication.ValidateToken(r); error != nil {
			responses.Error(w, http.StatusUnauthorized, error)
			return
		}
		nextFunction(w, r)
	}

}
