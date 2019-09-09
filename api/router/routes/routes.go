package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/obasajujoshua31/blogos/api/middlewares"
)

type Route struct {
	URI     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

func LoadRoutes() []Route {
	routes := usersRoutes
	routes = append(routes, PostRoutes...)
	routes = append(routes, AuthRoutes...)
	return routes
}

func SetupRoutes(r *mux.Router) *mux.Router {
	for _, route := range LoadRoutes() {
		r.HandleFunc(route.URI, route.Handler).Methods(route.Method)
	}
	return r
}

func SetupRoutesWithMiddlewares(r *mux.Router) *mux.Router {
	for _, route := range LoadRoutes() {
		r.HandleFunc(route.URI, middlewares.SetMiddlewareLogger(
			middlewares.SetMiddlewareJSON(route.Handler)),
		).Methods(route.Method)
	}
	return r
}
