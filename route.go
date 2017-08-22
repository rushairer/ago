package ago

import (
    "net/http"
)

type Route struct {
	Pattern         string
    Controller      ControllerInterface
	Method          string
    MethodName      string
	HandlerFunc     http.HandlerFunc
    Middlewares     []Middleware
}

func (route *Route) Middleware(m Middleware) *Route {
    route.Middlewares = append(route.Middlewares, m)
    return route
}

type Routes []*Route

func (routes Routes) Middleware(m Middleware) Routes {
    for key, route := range routes {
        routes[key] = route.Middleware(m)
    }
    return routes
}
