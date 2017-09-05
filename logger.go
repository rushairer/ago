package ago

import (
	"log"
	"net/http"
	"os"
	"time"
)

//Logger Logger middleware
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf(
			"%d\t%s\t%s\t%s\t%s",
			os.Getpid(),
			r.RemoteAddr,
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	})
}
