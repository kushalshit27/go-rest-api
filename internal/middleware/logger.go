package middleware

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// Logger wraps an HTTP handler and logs the request as necessary.
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		//log.Printf("(%s) \"%s %s %s\" %s", r.RemoteAddr, r.Method, r.RequestURI, r.Proto, time.Since(start))
		log.WithFields(log.Fields{
			"method":     r.Method,
			"path":       r.RequestURI,
			"proto":      r.Proto,
			"latency_ms": time.Since(start).Milliseconds(),
		}).Info("request details")
	})
}
