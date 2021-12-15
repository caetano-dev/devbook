package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var userRoutes = []Route{
	{
		URI:                    "/create-user",
		Method:                 http.MethodGet,
		Function:               controllers.LoadSignUpPage,
		RequiresAuthentication: false,
	},
}
