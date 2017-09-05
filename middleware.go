package ago

import (
	"net/http"
)

//Middleware Middleware Type
type Middleware func(handler http.Handler) http.Handler
