package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Route represents a route
type Route struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	RequiresAuthentication bool
}

//Config puts all the routes in the router
func Config(router *mux.Router) *mux.Router {
	routes := loginRoutes
	for _, route := range routes {
		router.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}
	return router
}
