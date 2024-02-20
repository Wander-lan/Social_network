package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Route represents all routes from API
type Route struct {
	URI                string
	Method             string
	Function           func(http.ResponseWriter, *http.Request)
	NeedAuthentication bool
}

// Config puts all routes inside the router
func Config(r *mux.Router) *mux.Router {
	routes := routesUsers
	routes = append(routes, loginRoute)

	for _, route := range routes {
		if route.NeedAuthentication {
			r.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Authenticate(route.Function)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}

	}

	return r
}
