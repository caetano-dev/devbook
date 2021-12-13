package routes

import (
	"api/src/controllers"
	"net/http"
)

var UserRoutes = []Route{
	{
		URI:                    "/users",
		Method:                 http.MethodPost,
		Function:               controllers.CreateUser,
		RequiresAuthentication: false,
	},

	{
		URI:                    "/users",
		Method:                 http.MethodGet,
		Function:               controllers.FetchUsers,
		RequiresAuthentication: true,
	},

	{
		URI:                    "/users/{userID}",
		Method:                 http.MethodGet,
		Function:               controllers.FetchUser,
		RequiresAuthentication: true,
	},

	{

		URI:                    "/users/{userID}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdateUser,
		RequiresAuthentication: true,
	},

	{

		URI:                    "/users/{userID}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeleteUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userID}/follow",
		Method:                 http.MethodPost,
		Function:               controllers.FollowUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userID}/unfollow",
		Method:                 http.MethodPost,
		Function:               controllers.UnfollowUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userID}/followers",
		Method:                 http.MethodGet,
		Function:               controllers.FetchFollowers,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userID}/following",
		Method:                 http.MethodGet,
		Function:               controllers.FetchFollowing,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userID}/updatePassword",
		Method:                 http.MethodPost,
		Function:               controllers.UpdatePassword,
		RequiresAuthentication: true,
	},
}
