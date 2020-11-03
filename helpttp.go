package helpttp

import (
	"log"
	"net/http"
	"time"
)

// WithLog wraps an HTTP handler and logs the request method and path before
// calling the inner handler.
func WithLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

// WithTimer wraps an HTTP handler and logs the time elapsed while calling the
// inner handler after it completes.
func WithTimer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tstart := time.Now()
		next.ServeHTTP(w, r)
		dur := time.Now().Sub(tstart)
		log.Printf("%s %s done in %s", r.Method, r.URL, dur)
	})
}
