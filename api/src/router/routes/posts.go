package routes

import (
	"api/src/controllers"
	"net/http"
)

var postsRoute = []Route{
	{
		URI:                    "/posts",
		Method:                 http.MethodPost,
		Function:               controllers.CreatePost,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/posts",
		Method:                 http.MethodGet,
		Function:               controllers.FetchPosts,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/posts/{postID}",
		Method:                 http.MethodGet,
		Function:               controllers.FetchPost,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/posts/{postID}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdatePost,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/posts/{postID}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeletePost,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userID}/posts",
		Method:                 http.MethodGet,
		Function:               controllers.FetchPostByUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/posts/{postID}/like",
		Method:                 http.MethodPost,
		Function:               controllers.LikePost,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/posts/{postID}/dislike",
		Method:                 http.MethodPost,
		Function:               controllers.DislikePost,
		RequiresAuthentication: true,
	},
}
