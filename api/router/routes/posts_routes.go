package routes

import (
	"net/http"

	"github.com/obasajujoshua31/blogos/api/controllers"
)

var PostRoutes = []Route{
	Route{
		URI:     "/posts",
		Handler: controllers.CreatePost,
		Method:  http.MethodPost,
	},
	Route{
		URI:     "/posts",
		Handler: controllers.GetPosts,
		Method:  http.MethodGet,
	},
	Route{
		URI:     "/posts/{id}",
		Handler: controllers.GetPost,
		Method:  http.MethodGet,
	},
	Route{
		URI:     "/posts/{id}",
		Handler: controllers.UpdatePost,
		Method:  http.MethodPut,
	},
}
