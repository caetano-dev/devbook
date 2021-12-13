package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

//Generate a router
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}
