package router

import "github.com/gorilla/mux"

//Generate a route
func Generate() *mux.Router {
	return mux.NewRouter()
}
