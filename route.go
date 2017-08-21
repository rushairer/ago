package ago

import (
    "net/http"
)

type Route struct {
    Name            string
	Method          string
	Pattern         string
	HandlerFunc     http.HandlerFunc
    Middlewares     []Middleware
}

func (route Route) Middleware(m Middleware) Route {
    route.Middlewares = append(route.Middlewares, m)
    return route
}
