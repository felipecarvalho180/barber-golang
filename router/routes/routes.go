package routes

import (
	"barber/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI                string
	Method             string
	Func               func(http.ResponseWriter, *http.Request)
	NeedAuthentication bool
}

func RoutesConfiguration(r *mux.Router) *mux.Router {
	routes := userRouters
	routes = append(routes, haircutRouters...)

	for _, route := range routes {
		if route.NeedAuthentication {
			r.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Authentication(route.Func)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Func)).Methods(route.Method)
		}
	}

	return r
}
