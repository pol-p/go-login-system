package middleware

import (
	"log"
	"net/http"
	"time"
)

type statusCodeResponse struct {
	http.ResponseWriter
	Status int
}

func (r *statusCodeResponse) WriteHeader(code int) {
	r.Status = code
	r.ResponseWriter.WriteHeader(code)
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writer := &statusCodeResponse{ResponseWriter: w, Status: 0}
		start := time.Now()
		next.ServeHTTP(writer, r)
		clock := time.Since(start)
		log.Printf("%s %s %s | Code: %d | Time: %v", r.Method, r.URL.Path, r.RemoteAddr, writer.Status, clock)
	})
}
