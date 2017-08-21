package ago

import (
    "net/http"
    "github.com/gorilla/mux"
)

type Router struct {
    Router *mux.Router `inject:""`
}

func (router *Router) NewHandler(routes []Route) *mux.Router{
    router.Router.StrictSlash(true)

    for _, route := range routes {

        route = route.Middleware(Logger)

        var handler http.Handler
        handler = route.HandlerFunc

        if (len(route.Middlewares) > 0) {
            for _, m := range route.Middlewares {
                handler = m(handler)
            }
        }

        if (route.Method == "WS") {
            router.Router.
                Path(route.Pattern).
                Name(route.Name).
                Handler(handler)
        } else {
            router.Router.
                Methods(route.Method).
                Path(route.Pattern).
                Name(route.Name).
                Handler(handler)
        }
    }

    return router.Router
}
