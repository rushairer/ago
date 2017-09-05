package ago

import (
	"log"
	"net/http"
)

//VERSION Version of ago
const VERSION = "0.1.0"

//Ago Base class of ago
type Ago struct {
	HTTPServer *http.Server `inject:""`
	Router     *Router      `inject:""`
	Routes     Routes
}

//Run Run ago
func (ago *Ago) Run(addr string) {

	ago.HTTPServer.Addr = addr
	ago.HTTPServer.Handler = ago.Router.NewHandler(ago.Routes)

	log.Printf("Running server %s ...\n", ago.HTTPServer.Addr)

	ago.HTTPServer.ListenAndServe()
}

//AddRoute Add route
func (ago *Ago) AddRoute(controller ControllerInterface, args ...string) Routes {
	var routes Routes
	var pattern, method, methodName string

	if len(args) == 1 {
		pattern = args[0]
	} else if len(args) == 2 {
		pattern = args[0]
		method = args[1]
	} else if len(args) == 3 {
		pattern = args[0]
		method = args[1]
		methodName = args[2]
	}

	if method == "*" || method == "" {
		routes = []*Route{
			&Route{
				Pattern:    pattern,
				Controller: controller,
				Method:     "GET",
			},
			&Route{
				Pattern:    pattern,
				Controller: controller,
				Method:     "POST",
			},
			&Route{
				Pattern:    pattern,
				Controller: controller,
				Method:     "HEAD",
			},
			&Route{
				Pattern:    pattern,
				Controller: controller,
				Method:     "DELETE",
			},
			&Route{
				Pattern:    pattern,
				Controller: controller,
				Method:     "PUT",
			},
			&Route{
				Pattern:    pattern,
				Controller: controller,
				Method:     "PATCH",
			},
			&Route{
				Pattern:    pattern,
				Controller: controller,
				Method:     "OPTIONS",
			},
		}
	} else {
		routes = []*Route{
			&Route{
				Pattern:    pattern,
				Controller: controller,
				Method:     method,
				MethodName: methodName,
			},
		}
	}

	ago.Routes = append(ago.Routes, routes...)

	return routes
}
