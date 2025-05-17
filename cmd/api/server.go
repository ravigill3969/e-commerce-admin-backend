package main

import (
	mw "admin-backend/internal/api/middlerwares"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/net/http2"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the root endpoint\n")
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintln(w, "GET: Returning user info")
	case http.MethodPost:
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		fmt.Fprintf(w, "POST: Received user data: %s\n", string(body))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	port := 8080

	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/user", userHandler)

	cert := "cert.pem"
	key := "key.pem"

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	rl := mw.NewRateLimiter(5, time.Minute)

	server := &http.Server{
		Addr:      fmt.Sprintf(":%d", port),
		Handler:   rl.Middlerware(mw.Compression(mw.ResponseTimeMiddlerware(mw.SecurityHeader(mw.Cors(mux))))),
		TLSConfig: tlsConfig,
	}

	http2.ConfigureServer(server, &http2.Server{})

	fmt.Println("Server is running on port:", port)

	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
