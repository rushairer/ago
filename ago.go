package ago

import (
    "log"
    "net/http"
)

const VERSION = "0.1.0"

type Ago struct {
    HttpServer *http.Server `inject:""`
    Router *Router `inject:""`
    Routes []Route
}

func (ago *Ago) Run(addr string) {

    ago.HttpServer.Addr = addr
    ago.HttpServer.Handler = ago.Router.NewHandler(ago.Routes)

    log.Printf("Running server %s ...\n", ago.HttpServer.Addr)

    ago.HttpServer.ListenAndServe()
}
