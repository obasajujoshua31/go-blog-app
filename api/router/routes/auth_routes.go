package routes

import (
	"net/http"

	"github.com/obasajujoshua31/blogos/api/controllers"
)

var AuthRoutes = []Route{
	Route{
		URI:     "/auth/login",
		Handler: controllers.LoginUser,
		Method:  http.MethodPost,
	},
}
