package ago

import (
    "net/http"
)

type Middleware func(handler http.Handler) http.Handler
