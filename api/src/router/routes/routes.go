package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

//Route represents all api routes
type Route struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	RequiresAuthentication bool
}

//Configure adds all routes inside of router
func Configure(r *mux.Router) *mux.Router {
	routes := UserRoutes
	routes = append(routes, loginRoute)
	routes = append(routes, postsRoute...)

	for _, route := range routes {
		if route.RequiresAuthentication {
			r.HandleFunc(route.URI, middlewares.Logger(middlewares.Authenticate(route.Function))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}
	return r
}
