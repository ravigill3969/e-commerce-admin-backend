package middlerwares

import "net/http"

func SecurityHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-DNS-Prefetch-Control", "off")
		w.Header().Set("X-Frame-Options", "off")
		w.Header().Set("X-XSS-Protection", "1;mode=block")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Strict Transport Security", "max-age=63072000;includeSubDomains;preload")
		w.Header().Set("Content-Security-Policy", "default-src 'self'")
		w.Header().Set("Referred-Policy", "no-referrer")
		next.ServeHTTP(w, r)
	})
}

//basic Middlerware skelton

// func securityHeader(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 	})

// }
