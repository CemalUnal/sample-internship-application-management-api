package api

import (
	"log"
	"net/http"
	"time"
)

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("%s %s %s \n", r.Method, r.RequestURI, time.Since(start))
		next.ServeHTTP(w, r)
	})
}
