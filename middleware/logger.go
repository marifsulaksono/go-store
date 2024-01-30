package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		next.ServeHTTP(w, r)

		log.Printf("New Request %s %s by %s with time %v", r.Method, r.RequestURI, r.RemoteAddr, time.Since(startTime))
	})
}
