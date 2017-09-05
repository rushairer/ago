package ago

import (
	"net/http"
)

//Route Include route informaiton
type Route struct {
	Pattern     string
	Controller  ControllerInterface
	Method      string
	MethodName  string
	HandlerFunc http.HandlerFunc
	Middlewares []Middleware
}

//Middleware Add a middleware to this route
func (route *Route) Middleware(m Middleware) *Route {
	route.Middlewares = append(route.Middlewares, m)
	return route
}

//Routes A type of []*Route
type Routes []*Route

//Middleware Add middlewares to Routes
func (routes Routes) Middleware(m Middleware) Routes {
	for key, route := range routes {
		routes[key] = route.Middleware(m)
	}
	return routes
}
