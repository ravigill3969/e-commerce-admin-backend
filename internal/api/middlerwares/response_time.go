package middlerwares

import (
	"fmt"
	"net/http"
	"time"
)

func ResponseTimeMiddlerware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received req in response writer")
		start := time.Now()
		wrappedWrite := &responseWriter{ResponseWriter: w, status: http.StatusOK}

		duration := time.Since(start)
		
		w.Header().Set("X-Response-Time", duration.String())
		next.ServeHTTP(wrappedWrite, r)
		
		duration = time.Since(start)
		//log the req details
		fmt.Printf(" Method %s , URL %s . status %d , duration %v \n", r.Method, r.URL, wrappedWrite.status, duration.String())
		fmt.Println("send response from response time middleware")
	})
}

//response writer

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)

}
