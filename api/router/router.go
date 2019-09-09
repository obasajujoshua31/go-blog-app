package router

import (
	"github.com/gorilla/mux"
	"github.com/obasajujoshua31/blogos/api/router/routes"
)

func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routes.SetupRoutesWithMiddlewares(r)
}
